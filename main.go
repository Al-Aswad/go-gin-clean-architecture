package main

import (
	"go-gin-clean-architecture/app/config"
	"go-gin-clean-architecture/app/controllers"
	"go-gin-clean-architecture/app/repositories"
	"go-gin-clean-architecture/app/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB            = config.DBConnect()
	authSerive services.AuthSevice = services.CreateAuthService(userRepo)
	jwtService services.JWTservice = services.CreateJwtService()

	userRepo    repositories.UserRepository = repositories.CreateUserRepo(db)
	userService services.UserService        = services.CreateUserService(userRepo)

	noteRepo    repositories.NoteRepository = repositories.CreateNoteRepository(db)
	noteService services.NoteService        = services.CreateNoteService(noteRepo)

	authController controllers.AuthController = controllers.CreateAuthController(authSerive, userService, jwtService)
	userController controllers.UserController = controllers.CreateUserController(userService)
	noteController controllers.NoteController = controllers.CreateNoteController(noteService, jwtService)
)

func main() {
	// defer config.CloseDatabaseConnection(db)
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	r := setupRouter()

	routes := r.Group("v1")
	{
		routes.POST("/login", authController.Login)
		routes.POST("/register", authController.Register)
		routes.POST("/users", userController.Create)
		routes.POST("/notes", noteController.Create)
		routes.PUT("/notes/:id", noteController.UpdateNoteByID)
	}

	r.Run(":" + os.Getenv("PORT"))

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
