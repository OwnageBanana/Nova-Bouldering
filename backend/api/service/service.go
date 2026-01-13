package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	database "novabouldering.ca/backend/api/pkg"
)

type NBService struct {
	Postgres       *pgxpool.Pool
	WriteAccessKey string
	ServerSecret   string
}

func (svc *NBService) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	all, err := database.GetFullHierarchy(ctx, svc.Postgres)
	if err != nil {
		log.Printf("failed query on get all: %v", err.Error())
		http.Error(w, "500 internal Server error: Failed to get database info", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all)
}

// mux.HandleFunc("GET /climbs/{id}/tags", service.GetAllClimbTags)
func (svc *NBService) GetAllClimbTags(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Imlemented", http.StatusNotImplemented)
}

// mux.HandleFunc("GET /tags", service.GetAllTags)
func (svc *NBService) GetAllTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tags, err := database.GetAllTags(ctx, svc.Postgres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}
