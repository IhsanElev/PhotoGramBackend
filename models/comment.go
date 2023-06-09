package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserId      uint `gorm:"not null" json:"user_id" form:"user_id" binding:"required"`
	User        *User
	UserPhoto   *UserPhoto
	UserPhotoID uint   `gorm:"not null" json:"user_photo_id" form:"user_photo_id" constraint:"OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Message     string `gorm:"not null" json:"message" form:"message" binding:"required"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
