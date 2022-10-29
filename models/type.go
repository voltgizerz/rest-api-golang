package models

type Type struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DamageTypeId int    `json:"damage_type_id"`
}
