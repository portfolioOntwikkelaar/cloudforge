package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cloudforge/internal/services"

	"github.com/go-chi/chi/v5"
)

func GetProjectBuilds(w http.ResponseWriter, r *http.Request) {
	projectID64, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}
	builds, err := services.GetProjectBuilds(uint(projectID64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(builds)
}

func GetBuild(w http.ResponseWriter, r *http.Request) {
	buildID64, err := strconv.ParseUint(chi.URLParam(r, "buildID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid build ID", http.StatusBadRequest)
		return
	}
	build, err := services.GetBuild(uint(buildID64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(build)
}
