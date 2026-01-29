package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
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
	queryDelete := "DELETE FROM books WHERE id=$1"

	result, err := DB.Exec(queryDelete, book.ID)
	if err != nil {
		http.Error(w, "Erro ao deletar", http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Nenhum livro com essas especificações encontrado", http.StatusNotFound)
		return
	}
	response := Message{
		Status:  200,
		Message: "success",
		Details: fmt.Sprintf("%d livros deletados", rowsAffected),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
