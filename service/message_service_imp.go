package service

import (
	"encoding/json"
	"go-chat-supabase/config"
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
	bytesWSessage, err := config.ReadCache("2024-01-21")
	if err != nil {
		log.Printf("error when read cache redis cause:%v\n", err)
		return err
	} else if bytesWSessage == nil {
		log.Printf("result cache nil %s", string(bytesWSessage))
		return err
	}

	var getWSMessage []ws.Message
	json.Unmarshal(bytesWSessage, &getWSMessage)
	log.Printf("ws message in service:%v\n", getWSMessage)
	var messageToDB entity.Message
	for _, dataInWs := range getWSMessage {
		messageToDB.Content = append(messageToDB.Content, dataInWs.Content)
		messageToDB.UserID = append(messageToDB.UserID, dataInWs.ClientID)
	}
	messageToDB.CreatedAt = now
	// insert to DB
	idMessage, err := m.repoMessage.Insert(messageToDB)
	if err != nil {
		log.Println("error cause:", err)
		return err
	}

	messageToDB.MessageID = idMessage
	// insert to supabase
	var results []entity.Message
	err = m.clientSupabase.DB.From("tb_message").Insert(messageToDB).Execute(&results)
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

func (m *MessageImp) ListMessage(body *model.ListAllMessageRequest) ([]entity.Message, error) {
	list, err := m.repoMessage.ListAll(nil)
	if err != nil {
		log.Printf("error when get data from DB cause:%v\n", err)
		return nil, err
	} else if list == nil {
		log.Printf("result data DB nil %v", list)
		return nil, nil
	}
	return list, nil
}

func (m *MessageImp) HandleServerRooom() func(*websocket.Conn) {
	return ws.HandleServer(m.hub)
}
