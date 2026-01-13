package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	// "novabouldering.com/data/pkg/database"

	svc "novabouldering.ca/backend/api/service"
)

func main() {
	dbUser := os.Getenv("NOVA_DATABASE_USER")
	if dbUser == "" {
		log.Fatal("NOVA_DATABASE_USER environment variable is not set")
	}
	dbPass := os.Getenv("NOVA_DATABASE_PASS")
	if dbPass == "" {
		log.Fatal("NOVA_DATABASE_PASS environment variable is not set")
	}
	dbHost := os.Getenv("NOVA_DATABASE_HOST")
	if dbHost == "" {
		log.Fatal("NOVA_DATABASE_HOST environment variable is not set")
	}
	dbName := os.Getenv("NOVA_DATABASE_DB_NAME")
	if dbName == "" {
		log.Fatal("NOVA_DATABASE_DB_NAME environment variable is not set")
	}
	writeAcessKey := os.Getenv("NOVA_WRITE_ACESS_KEY")
	if writeAcessKey == "" {
		log.Fatal("NOVA_WRITE_ACESS_KEY environment variable is not set")
	}

	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(dbUser, dbPass),
		Host:   dbHost,
		Path:   dbName,
	}
	q := u.Query()
	q.Set("sslmode", "disable")
	u.RawQuery = q.Encode()

	connectionString := u.String()

	fmt.Println("Safe Connection String:")
	fmt.Println(connectionString)
	dbPool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	service := svc.NBService{Postgres: dbPool, WriteAccessKey: writeAcessKey }

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
	mux.HandleFunc("POST /climbs", service.CreateClimb)
	mux.HandleFunc("GET /climbs/{id}", service.GetClimb)
	mux.HandleFunc("DELETE /climbs/{id}", service.DeleteClimb)
	mux.HandleFunc("POST /climbs/{id}", service.UpdateClimb)
	mux.HandleFunc("GET /climbs/{id}/tags", service.GetAllClimbTags)
	// mux.HandleFunc("POST /climbs", service.UpdateClimbsBatch)

	mux.HandleFunc("GET /tags", service.GetAllTags)

	log.Println("Server starting on :8085")
	if err := http.ListenAndServe(":8085", mux); err != nil {
		log.Fatal(err)
	}
}
