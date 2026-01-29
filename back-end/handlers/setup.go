package handlers

import "database/sql"

var DB *sql.DB

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}
type IDRequest struct {
	ID string `json:"id"`
}
type AuthorRequest struct {
	Author string `json:"author"`
}
type NameRequest struct {
	Name string `json:"name"`
}
type BookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Price  string `json:"price"`
}
type BookResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Price  string `json:"price"`
}
type BookUpdate struct {
	ID        string `json:"id"`
	NewName   string `json:"newName"`
	NewAuthor string `json:"newAuthor"`
	NewPrice  string `json:"newPrice"`
}
