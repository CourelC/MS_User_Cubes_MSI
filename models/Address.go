package models

type Address struct {
	ID          int    `json:"id_address" gorm:"primaryKey; not null"`
	Address     string `json:"Address"`
	CodePostale string `json:"CodePostale"`
	Ville       string `json:"Ville"`
}

func (b *Address) TableName() string {
	return "address"
}

func AddAddress() (err error) {
	return
}

func GetAddressById(id int) (err error) {
	return
}
