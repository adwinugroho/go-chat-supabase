package config

import (
	"log"

	supa "github.com/nedpals/supabase-go"

	realtimego "github.com/overseedio/realtime-go"
)

type (
	EnvSupabase struct {
		SB_URL      string `mapstructure:"supabase_url"`
		SB_API_KEY  string `mapstructure:"supabase_api_key"`
		SB_PASSWORD string `mapstructure:"supabase_password"`
	}

	PostgresChanges struct {
		Event  string `json:"event"`
		Schema string `json:"schema"`
		Table  string `json:"table"`
	}

	Config struct {
		PostgresChanges []PostgresChanges `json:"postgres_changes"`
	}

	Payload struct {
		Config Config `json:"config"`
	}

	Message struct {
		Topic   string      `json:"topic"`
		Event   string      `json:"event"`
		Payload interface{} `json:"payload"`
		Ref     string      `json:"ref"`
	}
)

var (
	SupabaseConfig EnvSupabase
)

// RLS Token optional
func InitSupabaseConnection(endpoint, apiKey, rlsToken string) *realtimego.Channel {
	supabaseClient, err := realtimego.NewClient(endpoint, apiKey)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil
	}

	// connect to server
	err = supabaseClient.Connect()
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil
	}

	// create and subscribe to channel
	db := "realtime"
	schema := "public"
	table := "coba"
	ch, err := supabaseClient.Channel(realtimego.WithTable(&db, &schema, &table))
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil
	}

	return ch
}

func InitSupabaseConnectionV2(url, key, password string) (*supa.Client, error) {
	supabase := supa.CreateClient(url, key)

	// ctx := context.Background()
	// user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
	// 	Email:    "adwinnugroho16@gmail.com",
	// 	Password: password,
	// })
	// if err != nil {
	// 	log.Printf("Error cause:%+v\n", err)
	// 	return nil, err
	// }
	log.Println("client supabase:", supabase)
	return supabase, nil
}
