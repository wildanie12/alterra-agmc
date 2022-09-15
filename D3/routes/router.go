package routes

import (
	"agmc_d3/controllers"
	"agmc_d3/database"
	"agmc_d3/lib"
	"agmc_d3/models"

	"github.com/labstack/echo/v4"
)

var bookController *controllers.BookController
var userController *controllers.UserController
var userRepository *lib.UserRepository

func init() {

	db := database.NewMySQL()
	db.AutoMigrate(&models.User{})
	userRepository = lib.NewUserRepository(db)

	bookController = controllers.NewBook()
	userController = controllers.NewUser(userRepository)
}

func SetRouter(e *echo.Echo) {
	book := e.Group("/v1/books")
	book.GET("", bookController.Index)
	book.POST("", bookController.Store)
	book.GET("/:bookID", bookController.Show)
	book.PUT("/:bookID", bookController.Update)
	book.DELETE("/:bookID", bookController.Delete)

	user := e.Group("/v1/users")
	user.GET("", userController.Index)
	user.POST("", userController.Store)
	user.GET("/:userID", userController.Show)
	user.PUT("/:userID", userController.Update)
	user.DELETE("/:userID", userController.Delete)
}