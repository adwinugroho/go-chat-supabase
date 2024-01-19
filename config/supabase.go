package config

import (
	"log"

	realtimego "github.com/overseedio/realtime-go"
)

type (
	EnvSupabase struct {
		SB_URL     string `mapstructure:"supabase_url"`
		SB_API_KEY string `mapstructure:"supabase_api_key"`
	}
)

var (
	SupabaseConfig EnvSupabase
)

// RLS Token optional
func InitSupabaseConnection(endpoint, apiKey, rlsToken string) {
	c, err := realtimego.NewClient(endpoint, apiKey)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return
	}

	// connect to server
	err = c.Connect()
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return
	}

	// create and subscribe to channel
	db := "realtime"
	schema := "public"
	table := "coba"
	ch, err := c.Channel(realtimego.WithTable(&db, &schema, &table))
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return
	}

	// setup hooks
	ch.OnInsert = func(m realtimego.Message) {
		log.Println("***ON INSERT....", m)
	}
	ch.OnDelete = func(m realtimego.Message) {
		log.Println("***ON DELETE....", m)
	}
	ch.OnUpdate = func(m realtimego.Message) {
		log.Println("***ON UPDATE....", m)
	}

	// subscribe to channel
	err = ch.Subscribe()
	if err != nil {
		return
	}
}
