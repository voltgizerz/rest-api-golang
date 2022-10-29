package routes

import (
	"net/http"
	"os"
	"rest-api-golang/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(controllerPokemon controllers.PokemonController) *gin.Engine {
	r := gin.New()

	mode, _ := os.LookupEnv("GIN_MODE")
	gin.SetMode(mode)

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
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found",
		})
	})

	return r
}
