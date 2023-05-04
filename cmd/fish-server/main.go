package main

import (
	"github.com/isaro/fish/internal/config"
	"github.com/isaro/fish/internal/router"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	httpServer := router.HttpRouter{}.CreateRouter()

	log.Println("Starting server on port 8080 in environment: ", config.GetConfig().Environment())
	err := http.ListenAndServe(":8080", httpServer)
	if err != nil {

	}
}
