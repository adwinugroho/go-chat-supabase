package service

import (
	"go-chat-supabase/model"
	"go-chat-supabase/repository"
)

type MessageImp struct {
	repoMessage repository.MessageInterface
}

func NewMessageService(repoMessage repository.MessageInterface) MessageInterface {
	return &MessageImp{repoMessage: repoMessage}
}

func (m *MessageImp) HandlerFetch() error {
	return nil
}

func (m *MessageImp) HandlerSend(body *model.NewSendMessageRequest) error {
	return nil
}
