package handlers

import (
	"encoding/json"
	"net/http"

	"cloudforge/internal/middleware"
)

func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)

	json.NewEncoder(w).Encode(map[string]any{
		"user_id": userID,
	})
}
