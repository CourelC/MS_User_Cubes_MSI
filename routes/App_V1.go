package routes

import (
	"MS_User_Cubes_MSI/controllers/user"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "MS_User_Cubes_MSI/docs"
)

// @host localhost:3000
// @BasePath /v1
func AppV1Router(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		// Documentation Swagger
		{
			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		// Users
		v1Users := v1.Group("/user")
		{
			v1Users.POST("/", user.NewUser)
			v1Users.GET("/:id", user.GetUserById)
			v1Users.GET("/", user.GetUsers)
			v1Users.PUT("/:id/inactive", user.InactiveUser)
			v1Users.PUT("/:id/valided", user.ValidedUser)
			v1Users.GET("/auth", user.AuthUser)
		}
	}
}
