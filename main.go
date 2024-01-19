package main

import (
	"go-chat-supabase/config"
	"log"
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
}
