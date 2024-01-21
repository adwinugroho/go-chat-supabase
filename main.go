package main

import (
	"go-chat-supabase/config"
	"go-chat-supabase/controller"
	"go-chat-supabase/pkg/ws"
	"go-chat-supabase/repository/postgres"
	"go-chat-supabase/service"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	// test load config with viper YAML
	config.LoadConfig()
	config.InitRedisClient()
	connPostgres := config.InitPostgresConnection(
		config.PostgreSQLConfig.PG_HOST,
		config.PostgreSQLConfig.PG_PORT,
		config.PostgreSQLConfig.PG_USERNAME,
		config.PostgreSQLConfig.PG_PASSWORD,
		config.PostgreSQLConfig.PG_DBNAME,
	)
	log.Println("successfully connect", connPostgres.Ping())

	chSupabase := config.InitSupabaseConnection(config.SupabaseConfig.SB_URL, config.SupabaseConfig.SB_API_KEY, "")
	clientSupabase, err := config.InitSupabaseConnectionV2(config.SupabaseConfig.SB_URL, config.SupabaseConfig.SB_API_KEY, config.SupabaseConfig.SB_PASSWORD)
	// log.Println("successfully test supabase realtime")

	// init repo
	initRepoMessage := postgres.NewMessageRepository(connPostgres)
	h := ws.NewHub()
	// init service
	initService := service.NewMessageService(initRepoMessage, chSupabase, clientSupabase, h)
	// setup fiber
	app := fiber.New()
	app.Use(recover.New())
	app.Use("/api/chat/ws", AllowUpgrade)
	// init controller
	initController := controller.NewChatController(&initService)
	// setup routing
	initController.RouteChat(app)
	go h.Run()
	// start app
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("error cause: ", err)
	}
}

func AllowUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	return nil
}
