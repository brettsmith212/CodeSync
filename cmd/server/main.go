/**
 * @file Main entry point for the CodeSync web server
 * @description
 * This file initializes the Go web server using the Chi router, sets up middleware,
 * loads environment variables, registers routes, and serves static files and templates.
 * It acts as the central hub for starting the application.
 *
 * Key features:
 * - Environment variable loading via godotenv
 * - Basic route handling with HTMX and Go templates
 * - Static file serving from the public directory
 *
 * @dependencies
 * - github.com/go-chi/chi/v5: HTTP router
 * - github.com/joho/godotenv: Environment variable loader
 * - html/template: Go template rendering
 * - net/http: HTTP server and client functionality
 *
 * @notes
 * - The server uses a default port of 8080 if PORT is not specified in .env
 * - Future routes for file, clipboard, and XML handling will be added later
 */

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize router
	r := chi.NewRouter()

	// Middleware for logging, recovery, and request context
	r.Use(middleware.Logger)      // Logs incoming requests
	r.Use(middleware.Recoverer)   // Recovers from panics
	r.Use(middleware.RequestID)   // Adds a unique ID to each request
	r.Use(middleware.RealIP)      // Ensures real client IP is captured

	// Load all templates from internal/templates directory
	tmpl, err := template.ParseGlob("internal/templates/**/*.html")
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}

	// Define root route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Data for the template
		data := map[string]interface{}{
			"Title": "Home",
		}
		// Render the base template with the data
		if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// TODO: Add routes for file handlers, clipboard handlers, and XML handlers

	// Serve static files from the public directory
	fileServer := http.FileServer(http.Dir("public"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}