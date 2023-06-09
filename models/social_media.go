package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type UserSocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" binding:"required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" binding:"required"`
	UserID         uint
	User           *User
}

func (s *UserSocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
func (s *UserSocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
