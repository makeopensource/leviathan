package jobs

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/google/uuid"
	"github.com/makeopensource/leviathan/service/message_store"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

const jobQueueTopicName = "jobqueue.topic"
const totalJobs = 5

func setupJobQueue() {
	// setup job store
	_ = message_store.NewMessageStore()
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
