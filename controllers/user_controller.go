package controllers

import (
	"finalproject/database"
	"finalproject/helpers"
	"finalproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{

		"id":       User.ID,
		"username": User.Username,
		"email":    User.Email,
		"age":      User.Age,
		"role":     User.Role,
	})

}

func UserLogin(c *gin.Context) {
	User := models.User{}
	db := database.GetDB()
	contentType := helpers.GetHeaderValue(c)
	_, _ = db, contentType
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)

	}
	password := User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid email",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid password",
		})
		return
	}
	token := helpers.GenerateToken(User.ID, User.Username, User.Role)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
