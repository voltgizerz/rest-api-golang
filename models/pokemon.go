package models

type Pokemon struct {
	ID   								int 					 `json:"id"`
	Name  							string				 `json:"name"`
	Height 							int						 `json:"height"`
	Weight 							int		 				 `json:"weight"`
	BaseExperience 			int  					 `json:"base_experience"`
	Types 							[]*Type 			 `gorm:"many2many:pokemon_types" json:"types"`
}
