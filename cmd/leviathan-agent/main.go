package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/google/uuid"
	api "github.com/makeopensource/leviathan/cmd/api"
	"github.com/makeopensource/leviathan/internal/dockerclient"
	store "github.com/makeopensource/leviathan/internal/messagestore"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

const jobQueueTopicName = "jobqueue.topic"
const totalJobs = 5

// test functions
//func main() {
//	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
//
//	client, err := dockerclient.NewLocalClient()
//	if err != nil {
//		log.Fatal().Msg("Failed to setup local docker client")
//	}
//
//	//client, err := dockerclient.NewSSHClient("r334@192.168.50.123")
//	//if err != nil {
//	//	log.Fatal().Msg("Failed to setup docker client")
//	//}
//
//	log.Info().Msg("Connected to remote client")
//
//	err = dockerclient.BuildImageFromDockerfile(client, "example/ex-Dockerfile", "testimage:latest")
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to build image")
//		return
//	}
//
//	images, err := dockerclient.ListImages(client)
//	if err != nil {
//		log.Error().Msg("Failed to build image")
//		return
//	}
//
//	for _, image := range images {
//		log.Info().Msgf("Container names: %v", image.RepoTags)
//	}
//
//	newContainerId, err := dockerclient.CreateNewContainer(
//		client,
//		"92912992939",
//		"testimage:latest",
//		[]string{"py", "/home/autolab/student.py"},
//		container.Resources{
//			Memory:   512 * 1000000,
//			NanoCPUs: 2 * 1000000000,
//		},
//	)
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to create container")
//		return
//	}
//
//	err = dockerclient.CopyToContainer(client, newContainerId, "example/student/test.py")
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to copy to container")
//	}
//
//	err = dockerclient.StartContainer(client, newContainerId)
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to start container")
//		return
//	}
//
//	err = dockerclient.TailContainerLogs(context.Background(), client, newContainerId)
//	if err != nil {
//		log.Error().Err(err).Msg("Failed to tail logs")
//		return
//	}
//
//	data, err := dockerclient.ListContainers(client)
//	if err != nil {
//		log.Error().Msg("Failed to build image")
//		return
//	}
//
//	for _, info := range data {
//		log.Info().Msgf("Container names: %v", info.Names)
//	}
//
//}

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

	port := "9221"
	srvAddr := fmt.Sprintf(":%s", port)
	srv := api.SetupPaths()
	err = srv.Run(srvAddr)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to start address on %s", srvAddr)
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
