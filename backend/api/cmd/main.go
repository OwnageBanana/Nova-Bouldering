package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Connect using standard database/sql
	dbURL := os.Getenv("NOVA_DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	// "postgres://user:pass@localhost:5432/dbname"
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Modern routing (Go 1.22+)
	mux.HandleFunc("GET /{}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		// Example data return
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"user_id": id, "status": "active"})
	})

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", mux)
}