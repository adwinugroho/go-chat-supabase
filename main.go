package main

import (
	"go-chat-supabase/config"
	"go-chat-supabase/controller"
	"go-chat-supabase/repository/postgres"
	"go-chat-supabase/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	// test load config with viper YAML
	config.LoadConfig()
	connPostgres := config.InitPostgresConnection(
		config.PostgreSQLConfig.PG_HOST,
		config.PostgreSQLConfig.PG_PORT,
		config.PostgreSQLConfig.PG_USERNAME,
		config.PostgreSQLConfig.PG_PASSWORD,
		config.PostgreSQLConfig.PG_DBNAME,
	)
	log.Println("successfully connect", connPostgres.Ping())

	// config.InitSupabaseConnection(config.SupabaseConfig.SB_URL, config.SupabaseConfig.SB_API_KEY, "")
	// log.Println("successfully test supabase realtime")

	// init repo
	initRepoMessage := postgres.NewMessageRepository(connPostgres)

	// init service
	initService := service.NewMessageService(initRepoMessage)
	// setup fiber
	app := fiber.New()
	app.Use(recover.New())
	// init controller
	initController := controller.NewChatController(&initService)
	// setup routing
	initController.RouteChat(app)
	// start app
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("error cause: ", err)
	}
}
