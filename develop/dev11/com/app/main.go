package main

import (
	"net/http"

	"github.com/rixagis/wb-level-2/develop/dev11/internal/app/handlers"
	"github.com/rixagis/wb-level-2/develop/dev11/internal/app/storage"
)

func main() {
	storage := storage.NewEventStorage()
	handler := handlers.NewHandler(storage)

	mux := handler.InitRoutes()
	http.ListenAndServe(":8080", mux)
}