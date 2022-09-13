package models

type Book struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Year     string `json:"year"`
	Stock    int    `json:"stock"`
}
