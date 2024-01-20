package repository

import "go-chat-supabase/entity"

type (
	MessageInterface interface {
		Insert(model entity.Message) (string, error)
		ListAll(filters map[string]interface{}) ([]entity.Message, error)
	}
)
