package routes

import (
	"agmc_d6/controllers"
	"agmc_d6/database"
	"agmc_d6/middlewares"
	"agmc_d6/models"
	"agmc_d6/repositories"

	"github.com/labstack/echo/v4"
)

var authController *controllers.AuthController
var bookController *controllers.BookController
var userController *controllers.UserController
var userRepository *repositories.UserRepository

func init() {

	db := database.NewMySQL()
	db.AutoMigrate(&models.User{})
	userRepository = repositories.NewUserRepository(db)

	authController = controllers.NewAuth(userRepository)
	bookController = controllers.NewBook()
	userController = controllers.NewUser(userRepository)
}

func SetRouter(e *echo.Echo) {
	auth := e.Group("/v1/auth")
	auth.POST("/login", authController.Login)

	book := e.Group("/v1/books")
	book.GET("", bookController.Index)
	book.POST("", bookController.Store, middlewares.VerifyJWT)
	book.GET("/:bookID", bookController.Show)
	book.PUT("/:bookID", bookController.Update, middlewares.VerifyJWT)
	book.DELETE("/:bookID", bookController.Delete, middlewares.VerifyJWT)

	user := e.Group("/v1/users")
	user.GET("", userController.Index, middlewares.VerifyJWT)
	user.POST("", userController.Store)
	user.GET("/:userID", userController.Show, middlewares.VerifyJWT)
	user.PUT("/:userID", userController.Update, middlewares.VerifyJWT)
	user.DELETE("/:userID", userController.Delete, middlewares.VerifyJWT)
}
