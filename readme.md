# Simple API Chat

## Description

Simple API Chat peer-to-peer and integrate with Supabase to store chat and fetch chat in realtime.

## Architecture
* config: pengaturan API, pengaturan env, tempat dimana membuka koneksi database dll.

* controller: routing API dan validasi dari depan (body dll.)

* entity: tabel dan isi tabel di database

* model: objek request dan response

* pkg: tempat 3rd party API, library dari luar, pecahan fungsi-fungsi untuk digunakan di service/repository/controller

* repository: fungsi-fungsi pada database seperti insert, update, delete dibuat disini

* service: fungsi-fungsi dimana logika aplikasi ditempatkan

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

`/api/chat/fetch:` Fetch all messages and save to database postgres and supabase
`/api/chat/list-all:` Retrieve all data or search filtering by message

### Request
```json
{
    "filters": null
}
```
### Response
```json
{
    "code": 200,
    "data": [
        {
            "message_id": "0652acf7-6a6f-4306-9a7f-1185beefa3cf",
            "content": [
                "Ayo makan",
                "Dimana?",
                "Terserah ikam aja",
                "Oke di Soto Banjar Nyaman Banar ya!"
            ],
            "user_id": [
                "user2",
                "user1",
                "user2",
                "user1"
            ],
            "created_at": "2024-01-21 12:42:26"
        },
        {
            "message_id": "6c2b8769-23e7-4568-b984-9313fdb02576",
            "content": [
                "Ayo makan",
                "Makan dimana cuyyy?",
                "Di Soto Banjar Nyaman Banar",
                "Kuyy gas lahhhh"
            ],
            "user_id": [
                "user2",
                "user1",
                "user2",
                "user1"
            ],
            "created_at": "2024-01-21 13:31:19"
        },
        {
            "message_id": "0652acf7-6a6f-4306-9a7f-1185beefa3cf",
            "content": [
                "Ayo makan",
                "Dimana?",
                "Terserah ikam aja",
                "Oke di Soto Banjar Nyaman Banar ya!"
            ],
            "user_id": [
                "user2",
                "user1",
                "user2",
                "user1"
            ],
            "created_at": "2024-01-21 12:42:26"
        }
    ],
    "status": true
}
```

`/api/chat/room/new:` Create a new room for chat with another user
### Request
```json
{
    "name": "Room No. 1"
}
```
### Response
```json
{
    "code": 200,
    "message": "Room successfully created with ID 293554",
    "status": true
}
```

`/api/chat/send:` Send broadcast messages for all user
### Request
```json
{
    "content": "Ini adalah pemberitahuan"
}
```
### Response
```json
{
    "code": 200,
    "message": "Message successfully sent",
    "status": true
}
```

`ws/[param_roomId]?clientId=[query_clientId]:` Websocket server for chat

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
