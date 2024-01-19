package main

import (
	"go-chat-supabase/config"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	// test load config with viper YAML
	loadConfig := config.LoadConfig()
	log.Println("load with viper:", loadConfig)
	connPostgres := config.InitPostgresConnection(
		loadConfig.PostgresConfig.PG_HOST,
		loadConfig.PostgresConfig.PG_USERNAME,
		loadConfig.PostgresConfig.PG_PASSWORD,
		loadConfig.PostgresConfig.PG_DBNAME,
		loadConfig.PostgresConfig.PG_PORT,
	)
	log.Println("successfully connect", connPostgres.Ping())
}
