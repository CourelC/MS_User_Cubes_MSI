package models

type Restaurateur struct {
	ID          int    `json:"id_restaurateur" gorm:"primaryKey; not null"`
	ID_User     int    `json:"id_user"`
	ID_Address  int    `json:"id_address"`
	Description string `json:"Description"`
	Telephone   string `json:"Telephone"`
	SIRET       string `json:"SIRET"`
	SIREN       string `json:"SIREN"`
}

var Restaurateurs []Restaurateur

func (b *Restaurateur) TableName() string {
	return "restaurateur"
}

func CreateRestaurateur(user *User) (err error) {
	return
}
