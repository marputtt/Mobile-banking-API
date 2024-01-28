// router/router.go
package router

import (
	"github.com/gin-gonic/gin"
	"MOBILEBANKINGAPI/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// User Endpoints
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
		userRoutes.PUT("/", controllers.UpdateUserInfo)
		userRoutes.DELETE("/", controllers.DeleteUser)
	}

	// Photo Endpoints
	photoRoutes := r.Group("/photos")
	{
		photoRoutes.POST("/", controllers.UploadPhoto)
		photoRoutes.GET("/", controllers.RetrievePhotos)
		photoRoutes.PUT("/:id", controllers.UpdatePhoto)
		photoRoutes.DELETE("/:id", controllers.DeletePhoto)
	}

	return r
}
