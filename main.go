package main

import (
	"Scheduler/server"
	"log"
	"os"
)

func main() {

	log.Println("Starting Scheduler App")

	log.Println("Initializing configuration")
	initConfig := InitConfig(getConfigFileName())

	log.Println("Initializing database")

	log.Println("Initializing HTTP sever")
	httpServer := server.InitHttpServer(initConfig)

	httpServer.Start()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "scheduler-" + env
	}

	return "scheduler"
}
