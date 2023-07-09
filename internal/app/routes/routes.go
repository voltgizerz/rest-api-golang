package routes

import (
	"net/http"
	"os"
	"rest-api-golang/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

// RouterController - all router controllers.
type RouterController struct {
	Pokemon controllers.PokemonController
}

// NewRouter - init router.
func NewRouter(controller RouterController) *gin.Engine {
	r := gin.New()

	mode, _ := os.LookupEnv("GIN_MODE")
	gin.SetMode(mode)

	pokemon := r.Group("/api/v1/")
	{
		pokemon.GET("pokemons/", controller.Pokemon.GetPokemons)
		pokemon.GET("pokemons/:id", controller.Pokemon.GetPokemons)
		pokemon.POST("pokemons/", controller.Pokemon.PostPokemons)
		pokemon.PATCH("pokemons/:id", controller.Pokemon.PatchPokemons)
		pokemon.DELETE("pokemons/:id", controller.Pokemon.DeletePokemon)
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
