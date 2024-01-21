package ws

import (
	"errors"
	"fmt"
	"go-chat-supabase/config"
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms            map[string]*Room
	Join             chan *Client
	Leave            chan *Client
	ForwardMessage   chan *Message
	BroadcastMessage chan *config.BroadcastMessageSupabase
}

func NewHub() *Hub {
	return &Hub{
		Rooms:            make(map[string]*Room),
		Join:             make(chan *Client),
		Leave:            make(chan *Client),
		ForwardMessage:   make(chan *Message),
		BroadcastMessage: make(chan *config.BroadcastMessageSupabase),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Join:
			// log.Printf("client in ch join:%+v\n", &client)
			_, ok := h.Rooms[client.RoomID]
			// log.Printf("ok in join:%v\n", ok)
			if ok {
				r := h.Rooms[client.RoomID]

				_, ok := r.Clients[client.ID]
				if !ok {
					r.Clients[client.ID] = client
				}
			}

		case client := <-h.Leave:
			_, ok := h.Rooms[client.RoomID]
			if ok {
				_, ok := h.Rooms[client.RoomID].Clients[client.ID]
				if ok {
					delete(h.Rooms[client.RoomID].Clients, client.ID)
					close(client.Message)
				}
			}

		case objMessage := <-h.ForwardMessage:
			_, ok := h.Rooms[objMessage.RoomID]
			// log.Printf("ok in forward message:%v\n", ok)
			if ok {
				for _, cl := range h.Rooms[objMessage.RoomID].Clients {
					cl.Message <- objMessage
				}
			}

		}
	}
}

func HandleServer(h *Hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		clientID := conn.Query("clientId")
		// log.Println("clientId:", clientID)
		roomID := conn.Params("roomId")
		if _, ok := h.Rooms[roomID]; !ok {
			log.Printf("error cause: %v\n", errors.New("roomId not found"))
			return
		}
		// log.Println("roomId:", roomID)
		client := &Client{
			Conn:    conn,
			Message: make(chan *Message),
			ID:      clientID,
			RoomID:  roomID,
		}
		h.Join <- client
		// if len(h.Rooms[roomID].Clients) >= 2 {
		// 	log.Printf("error cause: %v\n", errors.New("room is already full"))
		// 	return
		// }
		// log.Printf("cek client who joined:%+v\n", &h.Join)

		m := &Message{
			Content: fmt.Sprintf("User %s has joined the room", clientID),
			RoomID:  roomID,
		}
		h.ForwardMessage <- m
		// log.Printf("cek message forwarded:%+v\n", &h.ForwardMessage)
		go client.WriteMessage()
		client.ReadMessage(h)
	}

}
