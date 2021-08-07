package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api-golang/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api/v1/")
	{
		grp1.GET("pokemons/", controllers.GetPokemons)
	 	grp1.GET("pokemons/:id", controllers.GetPokemons)
		grp1.POST("pokemons/", controllers.PostPokemons)
		grp1.PATCH("pokemons/:id", controllers.PatchPokemons)
		grp1.DELETE("pokemons/:id", controllers.DeletePokemon)
	}
	
	return r
 }