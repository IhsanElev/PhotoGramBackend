package models

type Like struct {
	GormModel
	UserId      uint `gorm:"not null" json:"user_id" form:"user_id" binding:"required"`
	User        *User
	UserPhoto   *UserPhoto
	UserPhotoID uint `gorm:"not null" json:"user_photo_id" form:"user_photo_id" constraint:"OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
