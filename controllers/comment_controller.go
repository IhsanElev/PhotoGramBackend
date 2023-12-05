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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	comment := models.Comment{}
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}
	comment.UserId = userId
	if err := db.First(&models.UserPhoto{}, comment.UserPhotoID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid photo ID"})
		return
	}

	result := db.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, comment)
}
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	Comments := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&Comments)
	} else {
		c.ShouldBind(&Comments)
	}
	Comments.UserId = userId
	Comments.ID = uint(commentId)
	err := db.Debug().Where("id = ?", commentId).Model(&models.Comment{}).Where("id = ?", commentId).Updates(&Comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comments)
}
func GetCommentById(c *gin.Context) {
	db := database.GetDB()
	Comments := models.Comment{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&Comments)
	} else {
		c.ShouldBind(&Comments)
	}
	Comments.UserId = userId
	Comments.ID = uint(commentId)

	err := db.Debug().Where("user_id = ? AND id = ?", Comments.UserId, commentId).Find(&Comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comments)

}
func GetComments(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&Comments)
	} else {
		c.ShouldBind(&Comments)
	}
	err := db.Debug().Where("user_id = ?", userID).Find(&Comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comments)

}
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comments := models.Comment{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	commentID, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType

	if contentType == appJSON {
		c.ShouldBindJSON(&Comments)
	} else {
		c.ShouldBind(&Comments)
	}
	Comments.UserId = userID
	Comments.ID = uint(commentID)
	err := db.Debug().Where("id = ? AND user_id = ?", commentID, userID).Delete(&Comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Requeest", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("id with %d berhasil dihapus", commentID)})
}
