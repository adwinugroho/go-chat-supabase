package service

import (
	"go-chat-supabase/entity"
	"go-chat-supabase/model"
	"go-chat-supabase/pkg/ws"
	"go-chat-supabase/repository"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	supa "github.com/nedpals/supabase-go"

	realtimego "github.com/overseedio/realtime-go"
)

type MessageImp struct {
	repoMessage     repository.MessageInterface
	channelSupabase *realtimego.Channel
	clientSupabase  *supa.Client
	hub             *ws.Hub
}

func NewMessageService(repoMessage repository.MessageInterface, chSupabase *realtimego.Channel, clientSupabase *supa.Client, hub *ws.Hub) MessageInterface {
	return &MessageImp{
		repoMessage:     repoMessage,
		channelSupabase: chSupabase,
		clientSupabase:  clientSupabase,
		hub:             hub,
	}
}

func (m *MessageImp) CreateRoom(body *model.NewRoomRequest) error {
	m.hub.Rooms[body.RoomID] = &ws.Room{
		ID:      body.RoomID,
		Name:    body.Name,
		Clients: make(map[string]*ws.Client),
	}
	return nil
}

func (m *MessageImp) HandlerFetch() error {
	now := time.Now().Format("2006-01-02 15:04:05")
	// row := map[string]interface{}{
	// 	"created_at": now,
	// 	"message":    "ini pesan",
	// }
	getWSessage := <-m.hub.ForwardMessage
	var messageToDB = entity.Message{
		Content:   getWSessage.Content,
		CreatedAt: now,
	}

	var results []entity.Message
	// insert to supabase
	err := m.clientSupabase.DB.From("coba").Insert(messageToDB).Execute(&results)
	if err != nil {
		log.Println("error cause:", err)
		return err
	}
	// insert to DB
	_, err = m.repoMessage.Insert(messageToDB)
	if err != nil {
		log.Println("error cause:", err)
		return err
	}

	// fmt.Println(results)
	return nil
}

func (m *MessageImp) HandlerSend(body *model.NewSendMessageRequest) error {
	return nil
}

func (m *MessageImp) HandleServerRooom() func(*websocket.Conn) {
	return ws.HandleServer(m.hub)
}
