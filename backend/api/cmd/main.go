package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	// "novabouldering.com/data/pkg/database"
	svc "novabouldering.ca/backend/api/service"
)

func main() {
	// Connect using standard database/sql
	dbURL := os.Getenv("NOVA_DATABASE_URL")
	if dbURL == "" {
		log.Fatal("NOVA_DATABASE_URL environment variable is not set")
	}
	// Connect using standard database/sql
	writeAccessKey := os.Getenv("NOVA_DATABASE_WRITE_ACCESS_KEY")
	if writeAccessKey == "" {
		log.Fatal("NOVA_DATABASE_WRITE_ACCESS_KEY environment variable is not set")
	}
	// postgres://username:password@localhost:5432/database_name
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	service := svc.NBService{Postgres: dbPool, WriteAccessKey: writeAccessKey}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	mux.HandleFunc("POST /AuthWriteAccess", service.AuthWriteAccess)
	mux.HandleFunc("GET /all", service.GetAll)
	// mux.HandleFunc("GET /zones", service.GetAllZones)
	// mux.HandleFunc("GET /areas", service.GetAllAreas)
	// mux.HandleFunc("GET /crags", service.GetAllCrags)
	// mux.HandleFunc("GET /boulders", service.GetAllBoulders)
	// mux.HandleFunc("GET /boulders/{id}", service.GetBoulder)

	mux.HandleFunc("GET /climbs", service.GetAllClimbs)
	mux.HandleFunc("GET /climbs/{id}/tags", service.GetAllClimbTags)
	// mux.HandleFunc("POST /climbs", service.UpdateClimbsBatch)
	mux.HandleFunc("POST /climbs/{id}", service.UpdateClimb)

	mux.HandleFunc("GET /tags", service.GetAllTags)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
