package models

type Admin struct {
	ID      int `json:"id_admin" gorm:"primaryKey; not null"`
	ID_User int `json:"id_user"`
}

var Admins []Admin

func (b *Admin) TableName() string {
	return "admin"
}
