package main

import (
	"database/sql"
	"fmt"
	"golang/handlers"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func startHandlers() {
	http.HandleFunc("/", dashBoardHandler)
	http.HandleFunc("POST /insert", handlers.Insert)
	http.HandleFunc("POST /delete", handlers.Delete)
	http.HandleFunc("GET /list", handlers.List)
	http.HandleFunc("POST /search", handlers.SearchID)
	http.HandleFunc("PATCH /update", handlers.Update)
	http.HandleFunc("POST /author", handlers.SearchAuthor)
	http.HandleFunc("POST /name", handlers.SearchName)
}
func main() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Println("Arquivo .env inexistente")
	}
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Erro ao abrir conexão: ", err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}
	handlers.DB = db
	startHandlers()
	fmt.Println("Servidor rodando em " + os.Getenv("BASE_LINK"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
func dashBoardHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, os.Getenv("DASHBOARD_URL"), http.StatusFound)
}
