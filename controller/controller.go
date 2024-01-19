package controller

import (
	"go-chat-supabase/model"
	"go-chat-supabase/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ChatController struct {
	MessageService service.MessageInterface
}

func NewChatController(messageService *service.MessageInterface) ChatController {
	return ChatController{MessageService: *messageService}
}

func (controller *ChatController) RouteChat(app *fiber.App) {
	routeOrder := app.Group("/api/chat")
	routeOrder.Post("/send", controller.SendMessage)
}

func (controller *ChatController) SendMessage(c *fiber.Ctx) error {
	var request model.NewSendMessageRequest
	err := c.BodyParser(&request)
	if err != nil {
		log.Println("error cause:", err)
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"code":         400,
				"status":       false,
				"errorMessage": "Invalid Data",
			})
	}
	err = controller.MessageService.HandlerSend(&request)
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
