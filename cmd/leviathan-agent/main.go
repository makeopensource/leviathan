package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/UBAutograding/leviathan/cmd/routes"
	"github.com/UBAutograding/leviathan/internal/dockerclient"
	store "github.com/UBAutograding/leviathan/internal/messagestore"
	"github.com/UBAutograding/leviathan/internal/rpc/V1/docker"
	"github.com/UBAutograding/leviathan/internal/rpc/V1/jobs"
	"github.com/UBAutograding/leviathan/internal/rpc/V1/labs"
	"github.com/UBAutograding/leviathan/internal/rpc/V1/stats"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"time"
)

const jobQueueTopicName = "jobqueue.topic"
const totalJobs = 5

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	_, err := dockerclient.NewSSHClient("r334@192.168.50.123")
	if err != nil {
		log.Fatal().Msg("Failed to setup docker client")
	}

	// setup job store
	_ = store.NewMessageStore()
	// setup job queue
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	// job handler
	messages, err := pubSub.Subscribe(context.Background(), jobQueueTopicName)
	if err != nil {
		panic(err)
	}
	// setup job processors
	for i := 1; i < totalJobs; i++ {
		go process(messages, i)
	}

	// setup grpc

	//grpcSrv := setupGrpcPaths()
	grpcSrv := grpc.NewServer()

	dockerSrv := &routes.DockerServiceSrv{}
	docker.RegisterDockerServiceServer(grpcSrv, dockerSrv)

	jobSrv := &routes.JobsServiceSrv{}
	jobs.RegisterJobServiceServer(grpcSrv, jobSrv)

	labSrv := &routes.LabServiceSrv{}
	labs.RegisterLabServiceServer(grpcSrv, labSrv)

	statsSrv := &routes.StatsServiceSrv{}
	stats.RegisterStatsServiceServer(grpcSrv, statsSrv)

	grpcPort := fmt.Sprintf(":%s", "9221")

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Error().Err(err).Msgf("failed to start server on %s", grpcPort)
		return
	}

	log.Info().Msgf("Grpc server started on %s", grpcPort)
	err = grpcSrv.Serve(listen)
	if err != nil {
		log.Error().Err(err).Msgf("failed to start server on %s", grpcPort)
		return
	}
}

// dummy process
func process(messages <-chan *message.Message, processer int) {
	for msg := range messages {
		fmt.Printf("Processor %d, received message: %s, payload: %s\n", processer, msg.UUID, string(msg.Payload))
		msg.Ack()
		time.Sleep(10 * time.Second)
		fmt.Println("Processor ", processer, "processed", msg.UUID)
	}
}

func handleNewJobRequest(w http.ResponseWriter, r *http.Request, publisher message.Publisher) {
	urlPath := r.URL.Path
	jobStr := urlPath[len("/job/set/"):]

	msgId := watermill.NewUUID()
	msg := message.NewMessage(msgId, []byte(jobStr))
	err := publisher.Publish(jobQueueTopicName, msg)
	if err != nil {
		panic(err)
	}

	_, err = w.Write([]byte("You accessed the path: " + urlPath + "UUID" + msgId))
	if err != nil {
		log.Error().Err(err)
		return
	}
}

func handleJobStatus(w http.ResponseWriter, r *http.Request, publisher message.Publisher) {
	urlPath := r.URL.Path
	jobIDStr := urlPath[len("/job/query/"):]

	jobID, err := uuid.Parse(jobIDStr)
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	_, err2 := fmt.Fprintf(w, "Querying job with ID: %s", jobID)
	if err2 != nil {
		return
	}
}

//func cleanup(c *client.Client, containerID string) {
//	dockerclient.StopContainer(c, containerID)
//	err := dockerclient.RemoveContainer(c, containerID, false, false)
//	if err != nil {
//		os.Exit(-1)
//	}
//	os.Exit(0)
//}

//func main() {
//
//	log.SetLevel(log.DebugLevel)
//	if log.GetLevel() == log.TraceLevel {
//		log.SetReportCaller(true)
//	}
//	// TODO: For prod ensure logs are json formatted for ingest
//	// log.SetFormatter(&log.JSONFormatter{})
//
//	cli, err := client.NewEnvClient()
//	// cli, err := dockerclient.NewSSHClient("yeager")
//	if err != nil {
//		log.Fatal("Failed to setup docker client")
//	}
//
//	// err = dockerclient.PullImage(cli, "ubautograding/autograding_image_2004")
//
//	id, err := dockerclient.CreateNewContainer(cli, "ubautograding/autograding_image_2004")
//	if err != nil {
//		log.Fatal("Failed to create image")
//	}
//
//	err = dockerclient.CopyToContainer(cli, id, fmt.Sprintf("%s/tmp/sanitycheck/tmp/", util.UserHomeDir()))
//	if err != nil {
//		cleanup(cli, id)
//	}
//
//	err = dockerclient.StartContainer(cli, id)
//	if err != nil {
//		cleanup(cli, id)
//	}
//
//	ctx, cancel := context.WithCancel(context.Background())
//	go func() {
//		err := dockerclient.TailContainerLogs(ctx, cli, id)
//		if err != nil {
//			log.Fatal("")
//		}
//	}()
//
//	time.Sleep(10 * time.Second)
//	cancel()
//
//	cleanup(cli, id)
//	dockerclient.ListContainer(cli)
//}
//list, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
//if err != nil {
//	return
//}
//go func() {
//	logs, err := cli.ContainerLogs(context.Background(), "fe2c5534dba7d7a18cd802d2d5133cd5763f166c0464a828c6d929512744a338", types.ContainerLogsOptions{ShowStdout: true,
//		ShowStderr: true,
//		Follow:     true,
//	})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer func(logs io.ReadCloser) {
//		err := logs.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(logs)
//
//	// Copy the log stream to stdout
//	_, err = io.Copy(os.Stdout, logs)
//	if err != nil && err != io.EOF {
//		panic(err)
//	}
//}()
//for _, container := range list {
//	fmt.Print(container.Names)
//}
