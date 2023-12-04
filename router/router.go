package router

import (
	"finalproject/controllers"
	"finalproject/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	swagger := r.Group("/swagger")
	{
		swagger.StaticFile("/swagger.json", "./docs/swagger.json")
	}
	url := ginSwagger.URL("http://localhost:8080/finalproject/docs.json")
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.GET("/photo", controllers.GetPhotos)
		photoRouter.GET("/:photoId", controllers.GetPhotoByID)
		photoRouter.POST("/create", controllers.CreatePhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controllers.DeletePhoto)
	}
	socialMediaRouter := r.Group("/socialmedia")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/create", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/socialmedia", controllers.GetSocialMedia)
		socialMediaRouter.GET("/:socialmediaId", controllers.GetSocialMediaById)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/create", controllers.CreateComment)
		commentRouter.GET("/comments", controllers.GetComments)
		commentRouter.GET("/:commentId", controllers.GetCommentById)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controllers.DeleteComment)
	}
	likeRouter := r.Group("/likes")
	{
		likeRouter.Use(middleware.Authentication())
		likeRouter.POST("/create", controllers.CreateLike)
		likeRouter.DELETE("/:likeId", middleware.LikeAuthorization(), controllers.Unlike)
	}
	return r
}
