package postgres

import (
	"database/sql"
	"go-chat-supabase/entity"
	"go-chat-supabase/repository"
)

type MessageImp struct {
	DB *sql.DB
}

func NewMessageRepository(conn *sql.DB) repository.MessageInterface {
	return &MessageImp{
		DB: conn,
	}
}

func (c *MessageImp) Insert(model entity.Message) error {
	return nil
}

func (c *MessageImp) ListAll() ([]entity.Message, error) {
	var results []entity.Message
	return results, nil
}
