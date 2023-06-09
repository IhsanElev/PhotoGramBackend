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

func GetPhotos(c *gin.Context) {
	photo := []models.UserPhoto{}
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	err := db.Debug().Where("user_id = ?", userData["id"]).Find(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, photo)
}
func GetPhotoByID(c *gin.Context) {
	photo := models.UserPhoto{}
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoID := c.Param("photoId")
	err := db.Debug().Where("user_id = ? AND id = ?", userData["id"], photoID).Find(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	comments, err := GetPhotoComments(photo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}
	photo.Comments = comments
	c.JSON(http.StatusOK, photo)
}
func GetPhotoComments(photoID uint) ([]models.Comment, error) {
	db := database.GetDB()
	var comments []models.Comment
	err := db.Debug().Where("user_photo_id = ?", photoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.UserPhoto{}
	contentType := helpers.GetHeaderValue(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindHeader(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserID = userID
	err := db.Debug().Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
}
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.UserPhoto{}
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	photoID, _ := strconv.Atoi(c.Param("photoId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindHeader(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoID)

	err := db.Model(&Photo).Where("id = ?", photoID).Updates(models.UserPhoto{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.UserPhoto{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoID, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserID = userID
	Photo.ID = uint(photoID)
	err := db.Debug().Where("id = ? AND user_id = ?", photoID, userID).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Requeest", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("id with %d berhasil dihapus", photoID)})
}
