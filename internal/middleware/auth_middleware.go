package middleware

import (
	"context"
	//"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	// ContextKeyUserID is the key used to store the user ID in the request context.
	UserIDKey contextKey = "userID"
)

// AuthMiddleware is a middleware that checks for a valid JWT token in the Authorization header.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		// fmt.Println("HEADER =", header)
		tokenString := strings.TrimPrefix(header)

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			//fmt.Println("JWT parse error:", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, " token is not valid", http.StatusUnauthorized)
			return
		}

		// Extract the user ID from the token claims
		claims := token.Claims.(jwt.MapClaims)

		userID := uint(claims["user_id"].(float64))

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
