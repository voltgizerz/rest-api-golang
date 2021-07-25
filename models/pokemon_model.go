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