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

//CreateSinglePokemon create singel pokemon data
func CreatePokemon(pokemon *Pokemon) (err error) {
	if err = config.DB.Create(&pokemon).Error; err != nil {
	 	return err
	}
	return nil
}

//DeleteSinglePokemon create singel pokemon data
func DeletePokemon(id string) (err error) {
	var pokemon *Pokemon
	if err = config.DB.Delete(&pokemon, id).Error; err != nil {
	 	return err
	}
	return nil
}

//UpdatePokemon to update pokemon data
func UpdatePokemon(pokemon *Pokemon, id string) (err error) {
	if err = config.DB.Model(&pokemon).Where("id = ?", id).Updates(pokemon).Error; err != nil {
	 	return err
	}
	return nil
}