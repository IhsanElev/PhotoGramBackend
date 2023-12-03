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

// GetPhotos retrieves photos for a user.
//
// It takes a gin.Context parameter, which represents the HTTP request and response.
// It does not take any other parameters.
// It returns an array of models.UserPhoto objects.
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
	likes, err := GetPhotoLikes(photo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
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
	photo.Likes = likes
	c.JSON(http.StatusOK, photo)
}

// GetPhotoComments retrieves the comments for a given photo ID.
//
// It takes a uint parameter `photoID` representing the ID of the photo.
// It returns a slice of models.Comment and an error.
func GetPhotoComments(photoID uint) ([]models.Comment, error) {
	db := database.GetDB()
	var comments []models.Comment
	err := db.Debug().Where("user_photo_id = ?", photoID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetPhotoLikes returns the likes for a given photo ID.
// k
// Parameters:
// - photoID: the ID of the photo to get the likes for.
//
// Returns:
// - []models.Like: a slice of Like models representing the likes for the photo.
// - error: an error if there was a problem retrieving the likes.
func GetPhotoLikes(photoID uint) ([]models.Like, error) {
	db := database.GetDB()
	var likes []models.Like
	err := db.Debug().Where("user_photo_id = ?", photoID).Find(&likes).Error
	if err != nil {
		return nil, err
	}
	return likes, nil
}

// CreatePhoto is a Go function that handles the creation of a user photo.
//
// It takes a Gin context as a parameter and retrieves the database connection.
// It also retrieves the user's data from the context and extracts the user ID.
//
// The function then checks the content type of the request and binds the header
// and body of the request accordingly. It sets the user ID of the photo to the
// extracted user ID.
//
// Finally, it creates the photo record in the database and returns the created
// photo as a JSON response.
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

// UpdatePhoto updates a user photo.
//
// This function takes a gin.Context as a parameter and updates a user photo in the database. It retrieves the database connection,
// the user photo ID, the user ID from the context, and the content type from the request header. It then binds the request body to
// the UserPhoto struct based on the content type. After that, it updates the corresponding user photo record in the database with
// the provided title, caption, and photo URL. If the update is successful, it returns the updated photo in the response. If there
// is an error during the update, it returns a JSON response with a 400 status code and an error message.
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

// DeletePhoto deletes a user photo.
//
// c: The gin context.
// Returns nothing.
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
