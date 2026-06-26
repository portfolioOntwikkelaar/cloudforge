package handlers

import (
	"cloudforge/internal/middleware"
	"cloudforge/internal/models"
	"cloudforge/internal/services"
	"encoding/json"
	"net/http"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var req models.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(uint)

	err := services.CreateProject(userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
