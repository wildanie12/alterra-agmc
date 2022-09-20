package services

import (
	"agmc_d6/models"
	"errors"
)

type BookService struct {
	books []models.Book	
}

func NewBook() *BookService {
	return &BookService{
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

func (bs *BookService) FindAll() ([]models.Book, error) {
	return bs.books, nil
}

func (bs *BookService) Find(id int) (models.Book, error) {
	if id >= len(bs.books) {
		return models.Book{}, errors.New("Book ID is invalid")
	}
	return bs.books[id], nil
}

func (bs *BookService) Insert(book models.Book) (models.Book, error) {
	book.ID = uint(len(bs.books))
	bs.books = append(bs.books, book)
	return book, nil
}

func (bs *BookService) Update(bookInput models.Book, id int) (models.Book, error) {
	if id >= len(bs.books) {
		return models.Book{}, errors.New("Book ID is invalid")
	}
	book := bs.books[id]
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
	bs.books[id] = book
	return book, nil
}

func (bs *BookService) Delete(id int) error {
	if id >= len(bs.books) {
		return errors.New("Book ID is invalid")
	}
	return nil
}

