package models

import (
	"rest-api-golang/config"
 )


//GetAllPokemons Fetch all pokemon data
func GetAllPokemons(pokemon *[]Pokemon) (err error) {
	if err = config.DB.Find(pokemon).Error; err != nil {
	 	return err
	}
	return nil
}