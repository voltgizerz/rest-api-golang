package models

type Pokemon struct {
	Id    							int 			`gorm:"column:pok_id",json:"id"`
	Name  							string		`json:"name"`
	Height 							int				`json:"height"`
	Weight 							int		 		`json:"weight"`
	BaseExperience 			int  			`json:"base_experience"`
}