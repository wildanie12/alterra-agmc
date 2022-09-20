package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"agmc_d6/models"
	"agmc_d6/services"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bs *services.BookService
}

func NewBook(bs *services.BookService) *BookController {
	return &BookController{
		bs: bs,
	}
}

func (bc *BookController) Index(c echo.Context) error {
	books, _ := bc.bs.FindAll()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Success getting a list of books",
		"data":    books,
	})
}

func (bc *BookController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("bookID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "Book ID is invalid",
		})
	}
	
	book, err := bc.bs.Find(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{} {
			"status": "error",
			"code": http.StatusBadGateway,
			"message": err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "success updating a book",
		"data":    book,
	})
}

func (bc *BookController) Store(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	book, _ = bc.bs.Insert(book)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Success storing a books",
		"data":    book,
	})
}

func (bc *BookController) Update(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("bookID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "Book ID is invalid",
		})
	}
	
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "Cannot parse payload",
		})
	}
	bookInput := models.Book{}
	json.Unmarshal(body, &bookInput)

	book, err := bc.bs.Update(bookInput, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{} {
			"status": "error",
			"code": http.StatusBadGateway,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "success updating a book",
		"data":    book,
	})
}

func (bc *BookController) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("bookID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "Book ID is invalid",
		})
	}
	err = bc.bs.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{} {
			"status": "error",
			"code": http.StatusBadGateway,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "success deleting a book",
		"data": map[string]int{
			"id": id,
		},
	})
}
