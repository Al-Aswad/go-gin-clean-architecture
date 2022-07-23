package main

import (
	"go-gin-clean-architecture/app/config"
	"go-gin-clean-architecture/app/controllers"
	"go-gin-clean-architecture/app/repositories"
	"go-gin-clean-architecture/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = config.DBConnect()
	userRepo       repositories.UserRepository = repositories.CreateUserRepo(db)
	userService    services.UserService        = services.CreateUserService(userRepo)
	userController controllers.UserController  = controllers.CreateUserController(userService)
)

func main() {
	// defer config.CloseDatabaseConnection(db)
	r := setupRouter()

	routes := r.Group("v1")
	{
		routes.POST("/users", userController.Create)
	}

	r.Run(":5000")

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
