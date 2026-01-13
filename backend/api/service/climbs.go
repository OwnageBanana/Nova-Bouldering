package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	database "novabouldering.ca/backend/api/pkg"
)

type PutClimbRequest struct {
	database.Climb
	tags []*database.Tag
}

func (svc *NBService) CreateClimb(w http.ResponseWriter, r *http.Request) {

	if err := svc.ValidateWriteAccess(r); err != nil {
		log.Printf("Failed auth validation: %#v ", err)
		http.Error(w, "401 not authorized", http.StatusUnauthorized)
		return
	}

	var req PutClimbRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	err := database.CreateClimb(ctx, svc.Postgres, &req.Climb)
	if err != nil {
		log.Printf("failed query on CreateClimb: %#v %v", req.Climb, err.Error())
		http.Error(w, "500 internal Server error: Failed to get database info", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&req.Climb)
}

func (svc *NBService) UpdateClimb(w http.ResponseWriter, r *http.Request) {

	if err := svc.ValidateWriteAccess(r); err != nil {
		log.Printf("Failed auth validation: %#v ", err)
		http.Error(w, "401 not authorized", http.StatusUnauthorized)
		return
	}

	var req PutClimbRequest
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

func (svc *NBService) DeleteClimb(w http.ResponseWriter, r *http.Request) {

	if err := svc.ValidateWriteAccess(r); err != nil {
		log.Printf("Failed auth validation: %#v ", err)
		http.Error(w, "401 not authorized", http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = database.DeleteClimb(ctx, svc.Postgres, (int32(id)))
	if err != nil {
		log.Printf("failed query on get all: %v", err.Error())
		http.Error(w, "500 internal Server error: Failed to get database info", http.StatusInternalServerError)
		return
	}
}

func (svc *NBService) GetAllClimbs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := database.GetAllClimbs(ctx, svc.Postgres)
	if err != nil {
		log.Printf("failed query on get all climbs: %v", err.Error())
		http.Error(w, "500 internal Server error: Failed to get database info", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (svc *NBService) GetClimb(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	if id == 0 {
		http.Error(w, "Invalid request, no Id provided", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	data, err := database.GetClimb(ctx, svc.Postgres, (int32(id)))
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		log.Printf("failed query on get climb: %v", err.Error())
		http.Error(w, "500 internal Server error: Failed to get database info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
