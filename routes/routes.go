package routes

import (
	"rest-api-golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	pokemon := r.Group("/api/v1/")
	{
		pokemon.GET("pokemons/", controllers.GetPokemons)
		pokemon.GET("pokemons/:id", controllers.GetPokemons)
		pokemon.POST("pokemons/", controllers.PostPokemons)
		pokemon.PATCH("pokemons/:id", controllers.PatchPokemons)
		pokemon.DELETE("pokemons/:id", controllers.DeletePokemon)
	}

	// set no route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	return r
}
