package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	var book BookRequest
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Erro ao ler JSON. Verifique a formatação.", http.StatusBadRequest)
		return
	}
	if book.Name == "" || book.Author == "" || book.Price == "" {
		http.Error(w, "Campos 'name', 'author' e 'price' são obrigatórios.", http.StatusBadRequest)
		return
	}
	queryInsert := `INSERT INTO books (name, author, price) VALUES($1, $2, $3) RETURNING id`
	db := DB
	var novoID int
	err = db.QueryRow(queryInsert, book.Name, book.Author, book.Price).Scan(&novoID)
	if err != nil {
		log.Println("Erro ", err)
		http.Error(w, "Erro interno ao processar dados", http.StatusInternalServerError)
		return
	}
	response := Message{
		Status:  201,
		Message: "success",
		Details: fmt.Sprintf("Informações inseridas: Id: %d, name: %s, author: %s, price: %s", novoID, book.Name, book.Author, book.Price),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
