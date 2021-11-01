package main

import (
	"dopi-quickstart/router"
	"dopi-quickstart/service"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load()

	client, err := service.NewClient()
	if err != nil {
		log.Fatalln("unable to connect to mongodb", err)
	}
	defer client.Close()

	quickstartService := service.NewItemService(client)
	itemRouter := router.NewItemRouter(quickstartService)

	port := os.Getenv("PORT")
	log.Println("Starting server on the port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, itemRouter.MuxRouter))
}
