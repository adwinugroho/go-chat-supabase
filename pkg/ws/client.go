package ws

import (
	"go-chat-supabase/entity"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan *Message
	ID      string `json:"id"`
	RoomID  string `json:"roomId"`
}

type Message struct {
	Content string `json:"content"`
	RoomID  string `json:"roomId"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	var messageToDB []entity.Message
	var now = time.Now().Local().Format("2006-01-02 15:04:05.0000")
	for {
		for message := range c.Message {
			c.Conn.WriteJSON(message)
			var eachMessage = entity.Message{
				Content:   message.Content,
				CreatedAt: now,
			}
			messageToDB = append(messageToDB, eachMessage)
		}
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Leave <- c
		c.Conn.Close()
	}()

	for {
		_, messageWS, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v\n", err)
			}
			break
		}

		initMessage := &Message{
			Content: string(messageWS),
			RoomID:  c.RoomID,
		}

		hub.ForwardMessage <- initMessage
	}
}
