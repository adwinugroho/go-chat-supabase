package controller

import (
	"fmt"
	"go-chat-supabase/config"
	"go-chat-supabase/model"
	"go-chat-supabase/pkg/helper"
	"go-chat-supabase/service"
	"log"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ChatController struct {
	MessageService service.MessageInterface
}

func NewChatController(messageService *service.MessageInterface) ChatController {
	return ChatController{MessageService: *messageService}
}

func (controller *ChatController) RouteChat(app *fiber.App) {
	routeChat := app.Group("/api/chat")
	routeChat.Use(CheckAPIKey)
	routeChat.Post("/fetch", controller.FetchMessage)
	routeChat.Post("/list-all", controller.ListMessage)
	routeChat.Post("/send", controller.SendMessage)
	routeChat.Post("/room/new", controller.NewRoom)

	initWS := websocket.New(controller.MessageService.HandleServerRooom())
	routeChat.Use("/ws/:roomId", initWS)
}

func (controller *ChatController) FetchMessage(c *fiber.Ctx) error {
	err := controller.MessageService.HandlerFetch()
	if err != nil {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]interface{}{
				"code":         500,
				"status":       false,
				"errorMessage": "Internal Server Error",
			})
	}
	return c.Status(fiber.StatusOK).
		JSON(map[string]interface{}{
			"code":    200,
			"status":  true,
			"message": "Ok!",
		})
}

func (controller *ChatController) ListMessage(c *fiber.Ctx) error {
	var body model.ListAllMessageRequest
	err := c.BodyParser(&body)
	if err != nil {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"code":         400,
				"status":       false,
				"errorMessage": "Invalid Data",
			})
	}

	result, err := controller.MessageService.ListMessage(&body)
	if err != nil {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]interface{}{
				"code":         500,
				"status":       false,
				"errorMessage": "Internal Server Error",
			})
	} else if result == nil {
		return c.Status(fiber.StatusNotFound).
			JSON(map[string]interface{}{
				"code":         404,
				"status":       false,
				"errorMessage": "Data Not Found",
			})
	}
	return c.Status(fiber.StatusOK).
		JSON(map[string]interface{}{
			"code":   200,
			"status": true,
			"data":   result,
		})
}

func (controller *ChatController) NewRoom(c *fiber.Ctx) error {
	var body model.NewRoomRequest
	err := c.BodyParser(&body)
	if err != nil {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"code":         400,
				"status":       false,
				"errorMessage": "Invalid Data",
			})
	}

	if body.Name == "" {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"code":         400,
				"status":       false,
				"errorMessage": "Bad Request",
			})
	}
	body.RoomID = fmt.Sprintf("%d", helper.GenerateRandomNumber(6))
	err = controller.MessageService.CreateRoom(&body)
	if err != nil {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]interface{}{
				"code":         500,
				"status":       false,
				"errorMessage": "Internal Server Error",
			})
	}
	return c.Status(fiber.StatusOK).
		JSON(map[string]interface{}{
			"code":    200,
			"status":  true,
			"message": fmt.Sprintf("Room successfully created with ID %s", body.RoomID),
		})
}

func (controller *ChatController) SendMessage(c *fiber.Ctx) error {
	var body model.NewSendMessageRequest
	err := c.BodyParser(&body)
	if err != nil {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"code":         400,
				"status":       false,
				"errorMessage": "Invalid Data",
			})
	}

	if body.Content == "" {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"code":         400,
				"status":       false,
				"errorMessage": "Bad Request",
			})
	}

	err = controller.MessageService.HandlerSend(&body)
	if err != nil {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]interface{}{
				"code":         500,
				"status":       false,
				"errorMessage": "Internal Server Error",
			})
	}
	return c.Status(fiber.StatusOK).
		JSON(map[string]interface{}{
			"code":    200,
			"status":  true,
			"message": "Message successfully sent!",
		})
}

func customKeyFunc() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// Always check the signing method
		if t.Method.Alg() != jwtware.HS256 {
			return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
		}

		// TODO custom implementation of loading signing key like from a database
		signingKey := config.GeneralConfig.SecretKey

		return []byte(signingKey), nil
	}
}
