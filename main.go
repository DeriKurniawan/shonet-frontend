package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"

	"github.com/jacky-htg/shonet-frontend/config"
	"github.com/jacky-htg/shonet-frontend/routing"
)

func main() {
	router := routing.NewRouter()
	allowedHeaders  := handlers.AllowedHeaders([]string{"*"})
	allowedOrigins  := handlers.AllowedOrigins([]string{"*"})
	allowedMethods  := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	err := http.ListenAndServe(config.GetString("server.address"), handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
