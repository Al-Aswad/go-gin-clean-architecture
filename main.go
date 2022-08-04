package main

import (
	"gin-note-app/config"
	"gin-note-app/controllers"
	"gin-note-app/middleware"
	"gin-note-app/repositories"
	"gin-note-app/services"
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

	routesAuth := r.Group("v1")
	{
		routesAuth.POST("/login", authController.Login)
		routesAuth.POST("/register", authController.Register)
	}

	routes := r.Group("v1")
	routes.Use(middleware.IsUser(jwtService))
	{
		routes.POST("/users", userController.Create)
		routes.POST("/notes", noteController.Create)
		routes.GET("/notes", noteController.All)
		routes.GET("/notes/:id", noteController.FindNoteByID)
		routes.PUT("/notes/:id", noteController.UpdateNoteByID)
		routes.PUT("/notes/archive/:id", noteController.UpdateArchive)
		routes.DELETE("/notes/:id", noteController.DeteleNoteByID)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	r.Run(":" + port)

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
