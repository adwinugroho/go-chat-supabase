# Simple API Chat

## Description

Simple API Chat peer-to-peer and integrate with Supabase to store chat and fetch chat in realtime.

## Architecture
* config
init

* controller
init

* entity
init

* model
init

* pkg
init

* repository
init

* service
init

## Tech Stack

- Golang 
- PostgreSQL (Database)
- Supabase
- Redis

## Framework & Library

- GoFiber (HTTP Framework)
- GoFiber Websocket
- Database/sql
- Viper (Configuration)
- Golang Migrate (Database Migration)


## Configuration

Example configuration is in `example.config.yaml` file.

## API Spec

Init

## Database Migration

All database migration is in `pkg/migration` folder.

### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration


## Run Application

### Run web server

```bash
go run main.go
```
