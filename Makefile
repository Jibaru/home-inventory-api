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

dev-start:
	docker-compose --env-file app.env up

dev-down:
	docker-compose down

dev-run:
	docker exec -it home-inventory-api-workspace-1 /bin/bash -c "make run"

dev-bash:
	docker exec -it home-inventory-api-workspace-1 /bin/bash
