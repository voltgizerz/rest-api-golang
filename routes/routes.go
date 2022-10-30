package routes

import (
	"net/http"
	"os"
	"rest-api-golang/controllers"

	"github.com/gin-gonic/gin"
)

// RouterController - all router controllers.
type RouterController struct {
	Pokemon controllers.PokemonController
}

// NewRouter - init router.
func NewRouter(controllers RouterController) *gin.Engine {
	r := gin.New()

	mode, _ := os.LookupEnv("GIN_MODE")
	gin.SetMode(mode)

	pokemon := r.Group("/api/v1/")
	{
		pokemon.GET("pokemons/", controllers.Pokemon.GetPokemons)
		pokemon.GET("pokemons/:id", controllers.Pokemon.GetPokemons)
		pokemon.POST("pokemons/", controllers.Pokemon.PostPokemons)
		pokemon.PATCH("pokemons/:id", controllers.Pokemon.PatchPokemons)
		pokemon.DELETE("pokemons/:id", controllers.Pokemon.DeletePokemon)
	}

	// set no route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found!",
		})
	})

	return r
}
