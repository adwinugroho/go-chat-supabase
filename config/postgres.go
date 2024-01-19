package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type (
	EnvPostgres struct {
		PG_HOST     string `mapstructure:"pg_host"`
		PG_USERNAME string `mapstructure:"pg_username"`
		PG_PASSWORD string `mapstructure:"pg_password"`
		PG_PORT     string `mapstructure:"pg_port"`
		PG_DBNAME   string `mapstructure:"pg_dbname"`
	}
)

var (
	PostgreSQLConfig EnvPostgres
	// SQL Encryption or else
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
	// err = db.Ping()
	// if err != nil {
	// 	log.Printf("Error cause:%+v\n", err)
	// 	return nil
	// }

	fmt.Println("Connected!")

	return db
}
