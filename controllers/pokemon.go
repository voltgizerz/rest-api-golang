package controllers

import (
	"encoding/json"
	"rest-api-golang/models"
	"io/ioutil"
	"time"

	"net/http"
  "github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetPokemons(c *gin.Context) {
	id := c.Param("id")
	var err error
	var pokemon []models.Pokemon

	if id == "" {
		err = models.GetAllPokemons(&pokemon)
	} else {
		err = models.GetSinglePokemon(&pokemon, id)
	}

	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, gin.H{
		  "status": http.StatusOK,
			"data": pokemon, 
		})
	}
}


//PostPokemons - Create new pokemon
func PostPokemons(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var pokemon models.Pokemon

	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "request body not valid",
			"data": pokemon,
		})
	} else {
		result := models.CreatePokemon(&pokemon)
		if result != nil {
			// Failed insert pokemon to database
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"message": "failed insert data to database",
				"data": pokemon,
			})
		} else {
			// success created new pokemon
			c.JSON(http.StatusCreated, gin.H{
				"status": http.StatusCreated,
				"message": "created",
				"created_at": time.Now().Format(time.RFC850),
				"data": pokemon,
			})
		}
	}
}

//DeletePokemon delete pokemon
func DeletePokemon(c *gin.Context) {
	id := c.Param("id")
	result :=  models.DeletePokemon(id)
	if result != nil {
		// Failed insert pokemon to database
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "failed insert data to database",
		})
	} else {
		// success deleted new pokemon
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "deleted",
			"deleted_at": time.Now().Format(time.RFC850),
		})
	}
}