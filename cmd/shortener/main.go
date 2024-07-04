package main

import (
	"log"
	"net/http"

	"github.com/vokinneberg/go-url-shortener-ddd/internal/api"
	"github.com/vokinneberg/go-url-shortener-ddd/internal/repository"
)

func main() {
	storage := repository.NewMemStore()
	httpHandler := api.NewHandler(storage)
	log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
