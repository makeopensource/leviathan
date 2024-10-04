package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/google/uuid"
	api "github.com/makeopensource/leviathan/cmd/api"
	"github.com/makeopensource/leviathan/internal/config"
	"github.com/makeopensource/leviathan/internal/dockerclient"
	store "github.com/makeopensource/leviathan/internal/message-store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"os"
	"time"
)

const jobQueueTopicName = "jobqueue.topic"
const totalJobs = 5

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	config.InitConfig()

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
	mux := api.SetupPaths()

	log.Info().Msgf("Started server on %s", srvAddr)
	err = http.ListenAndServe(
		srvAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

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
