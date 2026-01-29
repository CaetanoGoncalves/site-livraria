package handlers

import (
	"encoding/json"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var book BookUpdate
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Erro ao ler JSON. Verifique a formatação.", http.StatusBadRequest)
		return
	}
	if book.ID == "" {
		http.Error(w, "Campos 'ID' é obrigatórios.", http.StatusBadRequest)
		return
	}
	queryUpdate := `
	UPDATE books 
	SET name = $2, author = $3, price = $4
	WHERE id = $1`
	result, err := DB.Exec(
		queryUpdate,
		book.ID,
		book.NewName,
		book.NewAuthor,
		book.NewPrice,
	)
	if err != nil {
		http.Error(w, "Erro ao atualizar: \n"+err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffect, _ := result.RowsAffected()
	if rowsAffect == 0 {
		http.Error(w, "Nenhum livro com esses parametros encontrado", http.StatusNotFound)
		return
	}
	response := Message{
		Status:  200,
		Message: "success",
		Details: "none",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
