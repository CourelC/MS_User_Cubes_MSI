package models

type Customer struct {
	ID        int    `json:"id_customer" gorm:"primaryKey; not null"`
	ID_User   int    `json:"id_user"`
	Pseudo    string `json:"Pseudo"`
	Telephone string `json:"Telephone"`
}

var Customers []Customer

func (b *Customer) TableName() string {
	return "customer"
}
