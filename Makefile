.PHONY: run lint

run:
	go run cmd/app/*

lint:
	golangci-lint run
