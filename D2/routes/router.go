package routes

import (
	"agmc_d2/controllers"

	"github.com/labstack/echo/v4"
)

var bookController controllers.BookController

func init() {
	bookController = *controllers.NewBook()
}

func SetRouter(e *echo.Echo) {
	book := e.Group("/v1/books")
	book.GET("", bookController.Index)

}