package controllers

import (
	"finalproject/database"
	"finalproject/helpers"
	"finalproject/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.UserSocialMedia{}
	contentType := helpers.GetHeaderValue(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.UserID = userId
	err := db.Debug().Create(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)

}
func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := []models.UserSocialMedia{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	err := db.Debug().Where("user_id = ?", userID).Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)

}
func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.UserSocialMedia{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	socialMediaId := c.Param("socialMediaId")
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	err := db.Debug().Where("user_id = ? AND socialmedia_id", userID, socialMediaId).Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)

}
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.UserSocialMedia{}
	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)
	err := db.Debug().Where("id = ?", socialMediaID).Updates(models.UserSocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.UserSocialMedia{}
	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)
	err := db.Debug().Where("id = ?", socialMediaID).Delete(SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("id with %d berhasil dihapus", socialMediaID)})
}
