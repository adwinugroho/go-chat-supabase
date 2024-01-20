package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms          map[string]*Room
	Join           chan *Client
	Leave          chan *Client
	ForwardMessage chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:          make(map[string]*Room),
		Join:           make(chan *Client),
		Leave:          make(chan *Client),
		ForwardMessage: make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Join:
			log.Printf("client in ch join:%+v\n", &client)
			_, ok := h.Rooms[client.RoomID]
			log.Printf("ok in join:%v\n", ok)
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
					if len(h.Rooms[client.RoomID].Clients) != 0 {
						h.ForwardMessage <- &Message{
							Content: "user left the chat",
							RoomID:  client.RoomID,
						}
					}

					delete(h.Rooms[client.RoomID].Clients, client.ID)
					close(client.Message)
				}
			}

		case objMessage := <-h.ForwardMessage:
			_, ok := h.Rooms[objMessage.RoomID]
			log.Printf("ok in forward message:%v\n", ok)
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
		// log.Println("roomId:", roomID)
		client := &Client{
			Conn:    conn,
			Message: make(chan *Message),
			ID:      clientID,
			RoomID:  roomID,
		}
		h.Join <- client
		// log.Printf("cek client who joined:%+v\n", &h.Join)

		m := &Message{
			Content: "A new user has joined the room",
			RoomID:  roomID,
		}
		h.ForwardMessage <- m
		// log.Printf("cek message forwarded:%+v\n", &h.ForwardMessage)
		go client.WriteMessage()
		client.ReadMessage(h)
	}

}
