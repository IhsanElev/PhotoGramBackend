package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type UserPhoto struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" binding:"required"`
	UserID   uint
	User     *User
	Comments []Comment `gorm:"ForeignKey:UserPhotoID" json:"comments"`
	Likes    []Like    `gorm:"ForeignKey:UserPhotoID" json:"likes"`
}

func (u *UserPhoto) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (u *UserPhoto) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
