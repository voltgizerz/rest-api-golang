package main

import (
	"log"

	"rest-api-golang/config"
	"rest-api-golang/routes"
)

func main() {
	// initialize database
	config.InitDB()

	// setup router
	r := routes.SetupRouter()

	err := r.Run()
	if err != nil {
		log.Println(err)
	}
}
