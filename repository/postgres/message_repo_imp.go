package postgres

import (
	"database/sql"
	"fmt"
	"go-chat-supabase/entity"
	"go-chat-supabase/repository"
	"log"

	"github.com/google/uuid"
	"github.com/lib/pq"
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
	model.MessageID = uuid.New().String()
	// defer c.DB.Close()
	query := `INSERT INTO table_message (message_id, content, description, created_at, user_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := c.DB.Exec(query, model.MessageID, pq.Array(model.Content), model.Description, model.CreatedAt, pq.Array(model.UserID))
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return "", err
	}

	fmt.Printf("Insert new message to DB with ID:%s\n", model.MessageID)
	return model.MessageID, nil
}

func (c *MessageImp) ListAll(filters map[string]interface{}) ([]entity.Message, error) {
	var results []entity.Message
	query := `SELECT * FROM table_message`
	if filters != nil {
		_, ok := filters["search_text"]
		if ok {
			searchByMessage := "%" + filters["search_text"].(string) + "%"
			query = fmt.Sprintf(`SELECT * FROM tb_message WHERE "content" LIKE '%s'`, searchByMessage)
		}
	}
	// log.Println("query:", query)
	rows, err := c.DB.Query(query)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result entity.Message

		err = rows.Scan(&result.MessageID, pq.Array(&result.Content), &result.Description, &result.CreatedAt, pq.Array(&result.UserID))
		if err != nil {
			log.Printf("Error cause:%+v\n", err)
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
