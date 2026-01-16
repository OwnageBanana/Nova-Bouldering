package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"

	svc "novabouldering.ca/backend/api/service"
	"novabouldering.ca/backend/api/storage"
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
	// publicBaseURL := os.Getenv("NOVA_R2_PUBLIC_URL")
	// if publicBaseURL == "" {
	// 	log.Fatal("NOVA_R2_PUBLIC_URL environment variable is not set")
	// }

	// properly format for connection. it escapes special characters
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

	dbPool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()

	storageService, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	service := svc.NBService{
		Postgres:       dbPool,
		WriteAccessKey: writeAcessKey,
		Storage:        storageService,
	}

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
	mux.HandleFunc("POST /climbs/images", service.UploadClimbImage)
	mux.HandleFunc("PUT /climbs/{id}/image", service.UpdateClimbImage)
	// mux.HandleFunc("POST /climbs", service.UpdateClimbsBatch)

	mux.HandleFunc("GET /tags", service.GetAllTags)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080", // Your frontend dev port
			"http://localhost:8085", // Your frontend dev port
			"https://novabouldering.com",
			"https://www.novabouldering.com",
			"https://novabouldering.ca",
			"https://www.novabouldering.ca",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)
	log.Println("Server starting on :8085")
	if err := http.ListenAndServe(":8085", handler); err != nil {
		log.Fatal(err)
	}
}
