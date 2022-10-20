package models

type Livreur struct {
	ID        int    `json:"id_livreur" gorm:"primaryKey; not null"`
	ID_User   int    `json:"id_user"`
	Pseudo    string `json:"Pseudo"`
	Telephone string `json:"Telephone"`
}

var Livreurs []Livreur

func (b *Livreur) TableName() string {
	return "livreur"
}
