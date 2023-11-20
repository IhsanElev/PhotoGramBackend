package middleware

import (
	"finalproject/database"
	"finalproject/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errror":  "unauthorized",
				"message": "you are not allowed to acces this data",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		Photo := models.UserPhoto{}
		err = db.Select("user_id").First(&Photo, uint(photoId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesnot exist",
			})
			return
		}
		fmt.Println("userID: ", userId)
		fmt.Println("Photo.UserID: ", Photo.UserID)
		if Photo.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to accessed this data",
			})
			return
		}
		c.Next()
	}
}
func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errror":  "unauthorized",
				"message": "you are not allowed to acces this data",
			})
			return
		}
		fmt.Println(socialMediaId)
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		SocialMedia := models.UserSocialMedia{}
		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesnot exist",
			})
			return
		}
		fmt.Println("userID: ", userId)
		fmt.Println("SocialMedia.UserID: ", SocialMedia.UserID)
		if SocialMedia.UserID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to accessed this data",
			})
			return
		}
		c.Next()
	}
}
func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errror":  "unauthorized",
				"message": "you are not allowed to acces this data",
			})
			return
		}
		fmt.Println(commentId)
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		Comments := models.Comment{}
		err = db.Select("user_id").First(&Comments, uint(commentId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesnot exist",
			})
			return
		}
		fmt.Println("userID: ", userId)
		fmt.Println("Comments.UserID: ", Comments.UserId)
		if Comments.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to accessed this data",
			})
			return
		}
		c.Next()
	}
}
