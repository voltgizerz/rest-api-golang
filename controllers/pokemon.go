package controllers

import (
	"rest-api-golang/models"

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
			"data": pokemon, 
			"status": http.StatusOK,
		})
	}
 }