package service

import (
	"go-chat-supabase/entity"
	"go-chat-supabase/model"

	"github.com/gofiber/contrib/websocket"
)

type (
	MessageInterface interface {
		CreateRoom(body *model.NewRoomRequest) error
		HandlerFetch() error
		HandlerSend(body *model.NewSendMessageRequest) error
		HandleServerRooom() func(*websocket.Conn)
		ListMessage(body *model.ListAllMessageRequest) ([]entity.Message, error)
	}
)
