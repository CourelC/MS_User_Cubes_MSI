package models

import (
	"MS_User_Cubes_MSI/config"
	modelErrors "MS_User_Cubes_MSI/errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int       `json:"id_user" gorm:"primaryKey; not null"`
	Name         string    `json:"Name"`
	Mail         string    `json:"Mail"`
	Password     string    `json:"Password"`
	CreationDate time.Time `json:"CreationDate,omitempty" gorm:"autoCreateTime:mili"`
	IsActive     bool      `json:"IsActive"`
	IsValided    bool      `json:"IsValided"`
}

var Users []User

func (b *User) TableName() string {
	return "user"
}

func CreateUser(user *User) (err error) {
	return
}

// Filtrer sur la distance
func GetUsers() (err error) {
	return
}

//Récup détail user
func GetUserByID(id int) (err error) {
 // Gérer avec champ type
	
	err = config.DB.Where("id = ?", id).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
		default:
			err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}
	return
}

// ValidedUser
func ValidedUser(id int) (err error) {

	return
}

// InactiveUser
func InactiveUser(id int) (err error) {

	return
}

func GetAdressUser(id int) (err error) {
	return
}

// AuthUser
func AuthUser(mail string, password string) (err error) {
	return
}
