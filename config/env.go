package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		PostgresConfig PostgresConfig
	}
)

func LoadConfig() AppConfig {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error when load config, cause: ", err)
		return AppConfig{}
	}

	var cfg AppConfig
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Println("error when unmarshall config, cause: ", err)
		return cfg
	}
	// log.Println(cfg)
	return cfg
}
