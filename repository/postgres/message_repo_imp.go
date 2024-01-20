package postgres

import (
	"database/sql"
	"fmt"
	"go-chat-supabase/entity"
	"go-chat-supabase/repository"
	"log"

	"github.com/google/uuid"
)

type MessageImp struct {
	DB *sql.DB
}

func NewMessageRepository(conn *sql.DB) repository.MessageInterface {
	return &MessageImp{
		DB: conn,
	}
}

func (c *MessageImp) Insert(model entity.Message) (string, error) {
	defer c.DB.Close()
	query := `INSERT INTO tb_message (content, description, createdAt) VALUES ($1, $2, $3) RETURNING id`

	model.ID = uuid.New().String()
	err := c.DB.QueryRow(query, model.Content, model.Description, model.CreatedAt).Scan(&model.ID)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return "", err
	}

	fmt.Printf("Insert new message to DB with ID:%s\n", model.ID)
	return model.ID, nil
}

func (c *MessageImp) ListAll(filters map[string]interface{}) ([]entity.Message, error) {
	var results []entity.Message
	sqlStatement := `SELECT * FROM tb_message`

	rows, err := c.DB.Query(sqlStatement)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result entity.Message

		err = rows.Scan(&result.ID, &result.Content, &result.Description, &result.CreatedAt)
		if err != nil {
			log.Printf("Error cause:%+v\n", err)
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
