package models

import (
	"finalproject/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string            `gorm:"not null" json:"username" form:"username" binding:"required"`
	Email        string            `gorm:"not null;unique" json:"email" form:"email" binding:"required,email"`
	Password     string            `gorm:"not null" json:"password" form:"password" binding:"required,min=6"`
	Age          int               `gorm:"not null" json:"age" form:"age" binding:"required,gt=7"`
	Role         string            `gorm:"default:user" json:"role,omitempty"`
	Comments     []Comment         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Photos       []UserPhoto       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	SocialMedias []UserSocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialmedias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}
