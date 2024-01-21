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
migrate create -ext sql -dir ./pkg/migration -seq init_schema
```
On this `000001_init_schema.up.sql` file in `pkg/migration` folder add this:
```sql
CREATE TABLE table_message (
    message_id VARCHAR(255) PRIMARY KEY,
    content TEXT,
    user_id TEXT,
    description VARCHAR(255),
    created_at VARCHAR(255)
);
```

```shell
migrate create -ext sql -dir ./pkg/migration -seq insert_dummy_table_message_1
```
On this `000002_insert_dummy_table_message_1.up.sql` file in `pkg/migration` folder add this:
```sql
INSERT INTO table_message (message_id,"content",description,created_at,user_id) VALUES
	 ('0652acf7-6a6f-4306-9a7f-1185beefa3cf','{"Ayo makan","Dimana?","Terserah ikam aja","Oke di Soto Banjar Nyaman Banar ya!"}','','2024-01-21 12:42:26','{"user2","user1","user2","user1"}'),
	 ('6c2b8769-23e7-4568-b984-9313fdb02576','{"Ayo makan","Makan dimana cuyyy?","Di Soto Banjar Nyaman Banar","Kuyy gas lahhhh"}','','2024-01-21 13:31:19','{"user2","user1","user2","user1"}');
```

### Run Migration

```shell
migrate -path ./pkg/migration -database "postgres://postgres:p@ssw0rd@localhost:5432/postgres?sslmode=disable" up
```

## Run Application

### Run web server

```bash
go run main.go
```
