package main

import (
	config "rest-api-golang/configs"
	"rest-api-golang/internal/app/controllers"
	"rest-api-golang/internal/app/repository"
	"rest-api-golang/internal/app/routes"
	"rest-api-golang/internal/app/service"
	"rest-api-golang/logger"
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
	controller := routes.RouterController{
		Pokemon: controllerPokemon,
	}

	r := routes.NewRouter(controller)

	err := r.Run()
	if err != nil {
		logger.Log.Error(err)
	}
}
