package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	database "novabouldering.ca/backend/api/pkg"
	"novabouldering.ca/backend/api/storage"
)

type NBService struct {
	Postgres       *pgxpool.Pool
	Storage        *storage.StorageService
	WriteAccessKey string
	ServerSecret   string
	// PublicBaseURL  string // Base URL for public R2 bucket access (e.g., "https://pub-xxx.r2.dev")
}

func (svc *NBService) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	all, err := database.GetFullHierarchy(ctx, svc.Postgres)
	if err != nil {
		log.Printf("failed query on get all: %v", err.Error())
		http.Error(w, "500 internal Server error: Failed to get database info", http.StatusInternalServerError)
		return
	}
	svc.addImageURLsToHierarchy(all)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all)
}

func (svc *NBService) addImageURLsToHierarchy(zones []*database.ZoneNode) {
	for _, zone := range zones {
		svc.prefixImageURL(zone.ImageURL)
		for _, crag := range zone.Crags {
			svc.prefixImageURL(crag.ImageURL)
			for _, area := range crag.Areas {
				svc.prefixImageURL(area.ImageURL)
				for _, boulder := range area.Boulders {
					svc.prefixImageURL(boulder.ImageURL)
					for i := range boulder.Climbs {
						svc.prefixImageURL(boulder.Climbs[i].ImageURL)
					}
				}
			}
		}
	}
}

func (svc *NBService) prefixImageURL(url *string) {
		if url != nil {*url = "https://images.novabouldering.ca/" + *url}
	// if url != nil && *url != "" && svc.PublicBaseURL != "" {
	// 	*url = svc.PublicBaseURL + "/" + *url
	// }
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
