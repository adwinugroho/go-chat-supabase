package config

import (
	"log"

	"github.com/spf13/viper"
)

var configStruct = map[string]interface{}{
	"postgres-config": &PostgreSQLConfig,
}

func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error when load config, cause: ", err)
		return
	}

	for key, value := range configStruct {
		log.Println("Loading config: ", key)
		if err := viper.Unmarshal(value); err != nil {
			log.Printf("Error loading config %s, cause: %+v\n", key, err)
			log.Fatal(err)
		}
		log.Printf("%s: %+v\n", key, value)
		log.Printf("Config %s loaded successfully", key)
	}
}
