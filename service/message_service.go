package service

import "go-chat-supabase/model"

type (
	MessageInterface interface {
		HandlerFetch() error
		HandlerSend(body *model.NewSendMessageRequest) error
	}
)
