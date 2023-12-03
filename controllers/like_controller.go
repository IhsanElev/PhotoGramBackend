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

func CreateLike(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	like := models.Like{}
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&like)
	} else {
		c.ShouldBind(&like)
	}

	like.UserId = userID

	if err := db.First(&models.UserPhoto{}, like.UserPhotoID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid like"})
		return
	}
	result := db.Create(&like)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, like)
}
func Unlike(c *gin.Context) {
	db := database.GetDB()
	Like := models.Like{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	likeID, _ := strconv.Atoi(c.Param("likeId"))
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType

	if contentType == appJSON {
		c.ShouldBindJSON(&Like)
	} else {
		c.ShouldBind(&Like)
	}
	Like.UserId = userID
	Like.ID = uint(likeID)
	err := db.Debug().Where("id = ? AND user_id = ?", likeID, userID).Delete(&Like).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Requeest", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("id with %d berhasil dihapus", likeID)})
}
