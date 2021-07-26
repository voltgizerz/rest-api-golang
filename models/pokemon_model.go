package models

import (
	"rest-api-golang/config"
	"gorm.io/gorm/clause"
 )


//GetAllPokemons Fetch all pokemon data
func GetAllPokemons(pokemon *[]Pokemon) (err error) {
	if err = config.DB.Preload(clause.Associations).Find(pokemon).Error; err != nil {
	 	return err
	}
	return nil
}

//GetSinglePokemon singel pokemon data
func GetSinglePokemon(pokemon *[]Pokemon, id string) (err error) {
	if err = config.DB.Preload(clause.Associations).Where("id = ?", id).Find(pokemon).Error; err != nil {
	 	return err
	}
	return nil
}