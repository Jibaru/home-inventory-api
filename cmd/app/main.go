package main

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/http"
	"log"
	"strconv"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	http.RunServer(config.AppHost, strconv.Itoa(config.AppPort))
}
