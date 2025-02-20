package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Load templates
	tmpl, err := template.ParseGlob("internal/templates/**/*.html")
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Title": "Home",
		}
		if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Serve static files
	fileServer := http.FileServer(http.Dir("public"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
