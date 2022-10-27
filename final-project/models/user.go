package models

import (
	"final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"email~invalid email format"`
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Password     string        `gorm:"not null" json:"password,omitempty" form:"password" valid:"required,minstringlength(6)~Your password is required and has to have a minimum length of 6 characters"`
	Age          int           `gorm:"not null" json:"age,omitempty" form:"age" valid:"range(8|100)~Minimum age is 8 years old"`
	CreatedAt    *time.Time    `json:"created_at,omitempty"`
	UpdatedAt    *time.Time    `json:"updated_at,omitempty"`
	Photos       []Photo       `json:"-"`
	Comments     []Comment     `json:"-"`
	Socialmedias []Socialmedia `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
