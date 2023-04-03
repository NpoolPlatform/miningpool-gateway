package pubsub

import (
	"github.com/google/uuid"
)

type MessageBase struct {
	MessageID   string
	UniqueID    uuid.UUID
	Sender      string
	RespondToID *uuid.UUID
}

type Message struct {
	MessageBase
	Body []byte
}