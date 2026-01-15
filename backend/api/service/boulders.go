package servicepackage service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	database "novabouldering.ca/backend/api/pkg"
)

type PutBoulderRequest struct {
	database.boulder
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
	for _, climb := range data {
		svc.prefixImageURL(climb.ImageURL)
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
	data, err := database.GetClimb(ctx, svc.Postgres, int32(id))
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		log.Printf("failed query on get climb: %v", err.Error())
		http.Error(w, "500 internal Server error: Failed to get database info", http.StatusInternalServerError)
		return
	}
	svc.prefixImageURL(data.ImageURL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (svc *NBService) UploadClimbImage(w http.ResponseWriter, r *http.Request) {
	if err := svc.ValidateWriteAccess(r); err != nil {
		log.Printf("Failed auth validation: %#v ", err)
		http.Error(w, "401 not authorized", http.StatusUnauthorized)
		return
	}

	// Parse multipart form (10MB max)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Missing image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Build the R2 key: climbs/{filename}
	r2Key := fmt.Sprintf("climbs/%s", header.Filename)

	// Upload to R2
	ctx := r.Context()
	if err := svc.Storage.Write(ctx, r2Key, file, header.Header.Get("Content-Type")); err != nil {
		log.Printf("Failed to upload to R2: %v", err)
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}

	// Create the image record in the database
	img := &database.Image{
		R2Key:       r2Key,
		Filename:    header.Filename,
		ContentType: header.Header.Get("Content-Type"),
		SizeBytes:   int32(header.Size),
	}

	if err := database.CreateImage(ctx, svc.Postgres, img); err != nil {
		log.Printf("Failed to create image record: %v", err)
		http.Error(w, "Failed to save image record", http.StatusInternalServerError)
		return
	}

	// Return the image with the full public URL
	response := struct {
		Id       int32  `json:"id"`
		ImageURL string `json:"image_url"`
		Filename string `json:"filename"`
	}{
		Id:       img.Id,
		ImageURL: svc.PublicBaseURL + "/" + r2Key,
		Filename: img.Filename,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (svc *NBService) UpdateClimbImage(w http.ResponseWriter, r *http.Request) {
	if err := svc.ValidateWriteAccess(r); err != nil {
		log.Printf("Failed auth validation: %#v ", err)
		http.Error(w, "401 not authorized", http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	climbId, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid climb ID format", http.StatusBadRequest)
		return
	}

	var req struct {
		ImageId int32 `json:"image_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := database.UpdateClimbImage(ctx, svc.Postgres, int32(climbId), req.ImageId); err != nil {
		log.Printf("Failed to update climb image: %v", err)
		http.Error(w, "Failed to update climb image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
