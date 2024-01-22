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
		SB_WS_URL   string `mapstructure:"sb_ws_url"`
	}
)

var (
	SupabaseConfig EnvSupabase
)

type BroadcastMessageSupabase struct {
	Event   string                           `json:"event"`
	Topic   string                           `json:"topic"`
	Payload *PayloadBroadcastMessageSupabase `json:"payload"`
	Ref     interface{}                      `json:"ref"`
}

type PayloadBroadcastMessageSupabase struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
	Type    string      `json:"ty"`
}

// RLS Token optional
func InitSupabaseRealtimeChannel(endpoint, apiKey, rlsToken string) *realtimego.Channel {
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
	// log.Println("client supabase:", supabase)
	return supabase, nil
}
