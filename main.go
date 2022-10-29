package main

import (
	"rest-api-golang/config"
	"rest-api-golang/controllers"
	"rest-api-golang/logger"
	"rest-api-golang/repository"
	"rest-api-golang/routes"
	"rest-api-golang/service"
)

func main() {
	config.LoadENV()

	// initialize database
	db := config.InitDB()

	// init repository
	repoPokemon := repository.NewPokemonRepository(db)

	// init service
	servicePokemon := service.NewPokemonService(repoPokemon)

	// init controller
	controllerPokemon := controllers.NewPokemonController(servicePokemon)

	// setup router
	r := routes.NewRouter(controllerPokemon)

	err := r.Run()
	if err != nil {
		logger.Log.Error(err)
	}
}
