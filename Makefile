.PHONY: run lint

include app.env
export $(shell sed 's/=.*//' app.env)

run:
	go run cmd/app/*

lint:
	golangci-lint run

migrate:
	 cd migrations && goose mysql "$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?parseTime=true" up

test:
	go test ./...
