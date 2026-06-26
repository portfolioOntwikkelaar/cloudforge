package main

import (
	"cloudforge/internal/handlers"
	"cloudforge/internal/middleware"
	"fmt"
	"net/http"
	"os"

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

	r.Get("/projects", handlers.GetProjects)
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Get("/me", handlers.Me)
		r.Post("/projects", handlers.CreateProject)
	})

	r.Post(
		"/register",
		handlers.Register,
	)
	r.Post("/login", handlers.Login)

	fmt.Println("Server running on port 8080")
	fmt.Println("JWT_SECRET =", os.Getenv("JWT_SECRET"))
	http.ListenAndServe(":8080", r)
}
