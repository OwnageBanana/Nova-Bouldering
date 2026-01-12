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

type UpdateClimbRequest struct {
	database.Climb
	tags []*database.Tag
}

func (svc *NBService) UpdateClimb(w http.ResponseWriter, r *http.Request) {
	if err := svc.ValidateWriteAccess(r); err != nil {
		log.Printf("Failed auth validation: %#v ", err)
		http.Error(w, "401 not authorized", http.StatusUnauthorized)
		return
	}
	var req UpdateClimbRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	err := database.UpdateClimb(ctx, svc.Postgres, &req.Climb)
		if err != nil {
		log.Printf("failed query on get all: %v", err.Error())
		http.Error(w, "500 internal Server error: Failed to get database info", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&req.Climb)
}

func (svc *NBService) GetAllClimbs(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Imlemented", http.StatusNotImplemented)
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
