FROM golang:1.21.3

ARG APP_PORT

RUN apt-get update && \
    apt-get install -y make

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.0

WORKDIR /app

EXPOSE ${APP_PORT}
