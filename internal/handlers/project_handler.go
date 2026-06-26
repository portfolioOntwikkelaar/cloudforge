package handlers

import (
	"cloudforge/internal/middleware"
	"cloudforge/internal/models"
	"cloudforge/internal/services"
	"encoding/json"
	"fmt"

	//"fmt"
	"net/http"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var req models.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	value := r.Context().Value(middleware.UserIDKey)

	//fmt.Printf("context value=%#v\n", value)
	//fmt.Printf("context type=%T\n", value)

	userID, ok := value.(uint)
	if !ok {
		http.Error(w, "unauthorized user ID", http.StatusUnauthorized)
		return
	}

	err := services.CreateProject(userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func GetProjects(w http.ResponseWriter, r *http.Request) {
	value := r.Context().Value(middleware.UserIDKey)

	userID, ok := value.(uint)
	if !ok {
		http.Error(w, "unauthorized user ID", http.StatusUnauthorized)
		return
	}

	fmt.Println("Calling GetProjects with userID=", userID)
	fmt.Println("Returned from repository")
	projects, err := services.GetProjects(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
