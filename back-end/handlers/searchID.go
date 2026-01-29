package handlers

import (
	"encoding/json"
	"net/http"
)

func SearchID(w http.ResponseWriter, r *http.Request) {
	var book IDRequest
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Erro ao ler JSON. Verifique a formatação.", http.StatusBadRequest)
		return
	}
	if book.ID == "" {
		http.Error(w, "Campo 'ID' é obrigatórios.", http.StatusBadRequest)
		return
	}
	searchQuery := `SELECT * FROM books WHERE id = $1`
	rows, err := DB.Query(searchQuery, book.ID)
	if err != nil {
		http.Error(w, "Erro ao buscar por ID", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var books []BookResponse
	for rows.Next() {
		var b BookResponse
		if err := rows.Scan(&b.ID, &b.Name, &b.Author, &b.Price); err != nil {
			http.Error(w, "Erro ao processar dados do banco", http.StatusInternalServerError)
			return
		}
		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro na leitura das linhas", http.StatusInternalServerError)
		return
	}
	if len(books) == 0 {
		http.Error(w, "Nenhum livro com este ID encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
