package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"agmc_d4/models"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	books []models.Book
}

func NewBook() *BookController {
	return &BookController{
		books: []models.Book{
			{
				ID:       0,
				Title:    "Book #1",
				Author:   "Wawan",
				Category: "Romance",
				Year:     "2021",
				Stock:    228,
			},
			{
				ID:       1,
				Title:    "Book #2",
				Author:   "Cahyo",
				Category: "Action",
				Year:     "2021",
				Stock:    337,
			},
		},
	}
}

func (bc *BookController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Success getting a list of books",
		"data":    bc.books,
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
	if id >= len(bc.books) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "Book ID is invalid",
		})
	}
	book := bc.books[id]
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
	book.ID = uint(len(bc.books))

	bc.books = append(bc.books, book)

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
	if id >= len(bc.books) {
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

	book := bc.books[id]
	if bookInput.Title != "" {
		book.Title = bookInput.Title
	}
	if bookInput.Author != "" {
		book.Author = bookInput.Author
	}
	if bookInput.Category != "" {
		book.Category = bookInput.Category
	}
	if bookInput.Year != "" {
		book.Year = bookInput.Year
	}
	if bookInput.Stock != 0 {
		book.Stock = bookInput.Stock
	}
	bc.books[id] = book

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
	if id >= len(bc.books) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "Book ID is invalid",
		})
	}
	bc.books = append(bc.books[:id], bc.books[id+1:]...)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "success deleting a book",
		"data": map[string]int{
			"id": id,
		},
	})
}
