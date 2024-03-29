package entity

// PokemonType - ...
type PokemonType struct {
	ID        int `json:"id"`
	PokemonID int `gorm:"primaryKey" json:"pokemon_id"`
	TypeID    int `gorm:"primaryKey" json:"type_id"`
	Slot      int `json:"slot"`
}
