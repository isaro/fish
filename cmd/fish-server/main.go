package main

import (
	"github.com/isaro/fish/internal/config"
	"github.com/isaro/fish/internal/router"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	httpServer := router.HttpRouter{}.CreateRouter()
	var dg = config.GetConfig().DiscordConfig().DiscordSession

	log.Println("Starting healthcheck on port 8080 in environment: ", config.GetConfig().Environment())
	//Used for Kubernetes Health check
	go func() {
		err := http.ListenAndServe(":8080", httpServer)
		if err != nil {
			log.Fatal(err)
		}
	}()

	//Wait for SIGINT or SIGTERM
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	//Close the Discord websocket connection
	err := dg.Close()
	if err != nil {
		log.Fatal("Error closing Discord websocket connection: ", err)
	}

}
