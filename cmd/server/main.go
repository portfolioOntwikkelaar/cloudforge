package main

import (
	"fmt"
	"net/http"

	"cloudforge/internal/database"

	"github.com/go-chi/chi/v5"
)

func main() {
	if err := database.Connect(); err != nil {
		panic(err)
	}

	if err := database.Migrate(); err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Cloudforge running")
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
