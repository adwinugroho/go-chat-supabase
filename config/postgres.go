package config

import (
	"database/sql"
	"fmt"
	"log"
)

type (
	PostgresConfig struct {
		PG_HOST     string `mapstructure:"pghost"`
		PG_USERNAME string `mapstructure:"pgusername"`
		PG_PASSWORD string `mapstructure:"pgpassword"`
		PG_PORT     string `mapstructure:"pgport"`
		PG_DBNAME   string `mapstructure:"pgdbname"`
	}
)

func InitPostgresConnection(host, port, user, password, dbName string) *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil
	}

	// check db
	err = db.Ping()
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil
	}

	fmt.Println("Connected!")

	return db
}
