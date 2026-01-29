package handlers

import (
	"encoding/json"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, name, author, price FROM books ORDER BY id ASC")
	if err != nil {
		http.Error(w, "Erro ao buscar livros", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var books []BookResponse
	for rows.Next() {
		var b BookResponse
		if err := rows.Scan(&b.ID, &b.Name, &b.Author, &b.Price); err != nil {
			continue
		}
		books = append(books, b)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
