package message_store

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"sync"
)

type MessageStore struct {
	messages map[string]*message.Message
	mu       sync.RWMutex
}

func NewMessageStore() *MessageStore {
	return &MessageStore{
		messages: make(map[string]*message.Message),
	}
}

func (s *MessageStore) Store(msg *message.Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages[msg.UUID] = msg
}

func (s *MessageStore) Get(id string) (*message.Message, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	msg, ok := s.messages[id]
	return msg, ok
}
