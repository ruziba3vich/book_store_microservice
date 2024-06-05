CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

DB_URL=postgres://postgres:Dost0n1k@localhost:5432/book_store?sslmode=disable

run:
	go run cmd/main.go

migrate_up:
	migrate -path migrations -database ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -database ${DB_URL} -verbose down

migrate_force:
	migrate -path migrations -database ${DB_URL} -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq lirary

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1
