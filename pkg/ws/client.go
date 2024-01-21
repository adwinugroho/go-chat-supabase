package ws

import (
	"encoding/json"
	"go-chat-supabase/config"
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
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	ClientID string `json:"clientId,omitempty"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		for message := range c.Message {
			c.Conn.WriteJSON(message)
		}
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Leave <- c
		c.Conn.Close()
	}()

	var now = time.Now().Local().Format("2006-01-02")

	for {
		_, messageWS, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error cause: %v\n", err)
			}
			break
		}

		initMessage := &Message{
			Content:  string(messageWS),
			RoomID:   c.RoomID,
			ClientID: c.ID,
		}
		var arrMessage []Message
		msgRedis, err := config.ReadCache(now)
		if err != nil {
			log.Println("error when get message from redis", err)
		}

		if msgRedis == nil {
			arrMessage = append(arrMessage, *initMessage)
			err = config.WriteCache(now, arrMessage)
			if err != nil {
				log.Println("error when set message to redis", err)
			}
		} else {
			// var getMsgRedis []Message
			json.Unmarshal(msgRedis, &arrMessage)
			arrMessage = append(arrMessage, *initMessage)
			err = config.DeleteCache(now)
			if err != nil {
				log.Println("error when delete message redis", err)
			}
			err = config.WriteCache(now, arrMessage)
			if err != nil {
				log.Println("error when set message to redis", err)
			}
		}

		log.Printf("results in read ws:%v\n", arrMessage)
		hub.ForwardMessage <- initMessage
	}
}
