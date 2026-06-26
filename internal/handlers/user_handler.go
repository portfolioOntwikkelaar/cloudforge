package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloudforge/internal/middleware"
)

func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)

	fmt.Printf("userID=%v (%T)\n", userID, userID)

	json.NewEncoder(w).Encode(map[string]any{
		"user_id": userID,
	})
}
