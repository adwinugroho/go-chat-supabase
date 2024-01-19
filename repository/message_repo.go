package repository

import "go-chat-supabase/entity"

type (
	MessageInterface interface {
		Insert(model entity.Message) error
		ListAll() ([]entity.Message, error)
	}
)
