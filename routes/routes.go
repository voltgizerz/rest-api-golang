package routes

import (
	"rest-api-golang/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(controllerPokemon controllers.PokemonController) *gin.Engine {
	r := gin.New()

	pokemon := r.Group("/api/v1/")
	{
		pokemon.GET("pokemons/", controllerPokemon.GetPokemons)
		pokemon.GET("pokemons/:id", controllerPokemon.GetPokemons)
		pokemon.POST("pokemons/", controllerPokemon.PostPokemons)
		pokemon.PATCH("pokemons/:id", controllerPokemon.PatchPokemons)
		pokemon.DELETE("pokemons/:id", controllerPokemon.DeletePokemon)
	}

	// set no route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	return r
}
