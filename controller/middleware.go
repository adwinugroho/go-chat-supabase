package controller

import (
	"go-chat-supabase/config"

	"github.com/gofiber/fiber/v2"
)

func CheckAPIKey(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	// fmt.Println(headers)
	APIKeyFromHeader, ok := headers["Api-Key"]
	// fmt.Println(APIKeyFromHeader)
	if ok && APIKeyFromHeader[0] == config.GeneralConfig.SecretKey {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).
		JSON(map[string]interface{}{
			"code":         401,
			"status":       false,
			"errorMessage": "Unauthorized",
		})
}
