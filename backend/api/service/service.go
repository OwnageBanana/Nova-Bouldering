package service

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	database "novabouldering.ca/backend/api/pkg"
)

type NBService struct {
	Postgres *pgxpool.Pool
	WriteAccessKey string
}

type AuthWriteAccessRequeset struct {
	Key string `json:"key"`
}

func (svc *NBService) AuthWriteAccess (w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	http.Error(w,"Not Imlemented", http.StatusNotImplemented)
	// var req AuthWriteAccessRequeset
}

func (svc *NBService) GetAll(w http.ResponseWriter, r *http.Request) {
	http.Error(w,"Not Imlemented", http.StatusNotImplemented)
}

func (svc *NBService) UpdateClimb(w http.ResponseWriter, r *http.Request) {
	http.Error(w,"Not Imlemented", http.StatusNotImplemented)
}

func (svc *NBService) GetAllClimbs(w http.ResponseWriter, r *http.Request) {
	http.Error(w,"Not Imlemented", http.StatusNotImplemented)
}


// mux.HandleFunc("GET /climbs/{id}/tags", service.GetAllClimbTags)
func (svc *NBService) GetAllClimbTags(w http.ResponseWriter, r *http.Request) {
	http.Error(w,"Not Imlemented", http.StatusNotImplemented)
}

	// mux.HandleFunc("GET /tags", service.GetAllTags)
func (svc *NBService) GetAllTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tags, err := database.GetAllTags(ctx, svc.Postgres)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}