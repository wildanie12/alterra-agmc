package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


type BookController struct {}

func NewBook() *BookController {
	return &BookController{}
}

func (bc *BookController) Index(c echo.Context) error {

	return c.String(http.StatusOK ,"ngok")
}