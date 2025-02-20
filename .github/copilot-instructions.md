# Project Instructions

Use specification and guidelines as you build the app.

Write the complete code for every step. Do not get lazy.

Your goal is to completely finish whatever I ask for.

You will see `<ai_context>` tags in the code. These are context tags that you should use to help you understand the codebase.

## Overview

This is a web app template built with HTMX and Go.

## Tech Stack

- Frontend: HTMX, Tailwind CSS, Alpine.js (optional for client-side interactivity)
- Backend: Go, Chi (router), Postgres, sqlc (SQL queries), Go Templates
- Auth: Custom JWT-based authentication (or optional third-party like Auth0)
- Payments: Stripe
- Deployment: Docker, Fly.io (or any Go-compatible platform)

## Project Structure

- `cmd` - Application entry points
  - `server` - Main server application
- `internal` - Internal application code
  - `db` - Database logic and sqlc-generated code
  - `handlers` - HTTP handlers for routes
  - `models` - Data models/structs
  - `services` - Business logic
  - `templates` - Go HTML templates
  - `utils` - Utility functions
- `migrations` - Database migration files
- `public` - Static assets (CSS, JS, images)
- `queries` - SQL query files for sqlc
- `scripts` - Build and utility scripts

## Rules

Follow these rules when building the app.

### General Rules

- Use Go module imports with the project module path (e.g., `github.com/username/project/internal/handlers`)
- Use kebab-case for file and folder names unless otherwise specified
- Keep all internal code in the `internal` directory to enforce encapsulation
- Use Go’s standard error handling (`if err != nil`) consistently

#### Env Rules

- Use a `.env` file for environment variables, loaded with `github.com/joho/godotenv`
- Update `.env.example` if you add new environment variables
- Access environment variables using `os.Getenv("VARIABLE_NAME")`
- Do not expose sensitive environment variables in client-side responses
- Example `.env` usage in Go:

```go
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    dbURL := os.Getenv("DATABASE_URL")
}
```

### Frontend Rules
Follow these rules when working on the frontend.

It uses HTMX, Tailwind CSS, and Go HTML templates.

#### General Rules
- Use HTMX for dynamic interactions instead of heavy client-side JavaScript
- Serve HTML fragments from the server via Go templates
- Use Tailwind CSS for styling, included via CDN or compiled locally
- Use lucide-icons (via CDN) for icons
- Keep JavaScript minimal; use Alpine.js only for lightweight client-side interactivity when neededFrontend Rules
- Follow these rules when working on the frontend.

It uses HTMX, Tailwind CSS, and Go HTML templates.

#### Templates
- Store templates in internal/templates
- Name files like base.html, index.html, partials/todo-item.html
- Use Go’s html/template package for rendering
- Organize partials in a partials subdirectory for reusable snippets
- Always include a base template (base.html) with common layout structure
- Use HTMX attributes (hx-get, hx-post, hx-target, etc.) for interactivity

Example base template:
```html
<!-- internal/templates/base.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }}</title>
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
    <script src="https://unpkg.com/htmx.org@1.9.6"></script>
    <script src="https://unpkg.com/lucide@0.0.6/dist/umd/lucide.min.js"></script>
</head>
<body class="bg-gray-100">
    {{ block "content" . }}{{ end }}
</body>
</html>
```

Example page template with HTMX:
```html
<!-- internal/templates/index.html -->
{{ define "content" }}
<div class="container mx-auto p-4">
    <h1 class="text-2xl">Todos</h1>
    <form hx-post="/todos" hx-target="#todo-list" hx-swap="innerHTML">
        <input type="text" name="content" class="border p-2">
        <button type="submit" class="bg-blue-500 text-white p-2">Add Todo</button>
    </form>
    <div id="todo-list">
        {{ range .Todos }}
            {{ template "todo-item" . }}
        {{ end }}
    </div>
</div>
{{ end }}
```

Example partial template:
```html
<!-- internal/templates/partials/todo-item.html -->
{{ define "todo-item" }}
<div class="flex items-center p-2 border-b" id="todo-{{ .ID }}">
    <span>{{ .Content }}</span>
    <button hx-delete="/todos/{{ .ID }}" hx-target="#todo-{{ .ID }}" hx-swap="outerHTML" class="ml-auto text-red-500">
        Delete
    </button>
</div>
{{ end }}
```

#### Data Fetching
- Fetch data server-side in Go handlers and render it into templates
- Use HTMX to request HTML fragments from the server for dynamic updates
- Avoid client-side data fetching unless absolutely necessary (e.g., via Alpine.js)

### Backend Rules
Follow these rules when working on the backend.

It uses Go, Chi router, Postgres, and sqlc for database queries.

#### General Rules
- Use `github.com/go-chi/chi` for routing
- Structure handlers in `internal/handlers`
- Keep business logic in `internal/services`
- Use `context.Context` for request-scoped data and cancellation

#### Organization
- Handlers: `internal/handlers/example-handler.go`
- Services: `internal/services/example-service.go`
- Models: `internal/models/example.go`
- Database: `internal/db/db.go` and `queries/*.sql`

#### Database
- Use sqlc to generate type-safe Go code from SQL queries
- Store SQL queries in queries/ (e.g., queries/todos.sql)
- Use migrations for schema changes (e.g., with github.com/golang-migrate/migrate)
- Include created_at and updated_at fields in all tables
- Use UUIDs for primary keys (via github.com/google/uuid)

Example sqlc query:

```sql
-- queries/todos.sql
-- name: CreateTodo :one
INSERT INTO todos (id, user_id, content, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;

-- name: GetTodosByUser :many
SELECT * FROM todos WHERE user_id = $1;
```

Generated Go code (via sqlc):

```go
// internal/db/todos.sql.go (auto-generated)
type Queries struct {
    db *sql.DB
}

type Todo struct {
    ID        uuid.UUID
    UserID    string
    Content   string
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (q *Queries) CreateTodo(ctx context.Context, id uuid.UUID, userID, content string) (Todo, error) {
    // Generated implementation
}

func (q *Queries) GetTodosByUser(ctx context.Context, userID string) ([]Todo, error) {
    // Generated implementation
}
```

Example model:

```go
// internal/models/todo.go
package models

import "github.com/google/uuid"

type Todo struct {
    ID        uuid.UUID
    UserID    string
    Content   string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

#### Handlers
- Name handler files like `todo-handler.go`
- Use HTMX-friendly responses (HTML fragments)
- Return errors as HTTP status codes with simple HTML error messages
- Use Go structs for request/response data

Example handler:
```go
// internal/handlers/todo-handler.go
package handlers

import (
    "html/template"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/google/uuid"
    "github.com/username/project/internal/db"
    "github.com/username/project/internal/models"
)

type TodoHandler struct {
    Queries *db.Queries
    Tmpl    *template.Template
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
    content := r.FormValue("content")
    userID := "user123" // Replace with real auth logic
    id := uuid.New()

    todo, err := h.Queries.CreateTodo(r.Context(), id, userID, content)
    if err != nil {
        http.Error(w, "Failed to create todo", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    h.Tmpl.ExecuteTemplate(w, "todo-item", todo)
}

func (h *TodoHandler) RegisterRoutes(r chi.Router) {
    r.Post("/todos", h.CreateTodo)
}
```

#### Services
- Encapsulate business logic in services
- Name files like todo-service.go

Example service:

```go
// internal/services/todo-service.go
package services

import (
    "context"

    "github.com/username/project/internal/db"
    "github.com/username/project/internal/models"
    "github.com/google/uuid"
)

type TodoService struct {
    Queries *db.Queries
}

func (s *TodoService) CreateTodo(ctx context.Context, userID, content string) (models.Todo, error) {
    id := uuid.New()
    return s.Queries.CreateTodo(ctx, id, userID, content)
}
```

### Auth Rules
Follow these rules when working on auth.

It uses custom JWT-based authentication (or an optional third-party like Auth0).

#### General Rules
- Implement JWT middleware with github.com/golang-jwt/jwt/v5
- Store JWT secrets in environment variables
- Validate tokens in a Chi middleware
- Pass user info via context.Context

Example middleware:

```go
// internal/handlers/auth-middleware.go
package handlers

import (
    "context"
    "net/http"
    "os"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), "userID", claims["userID"])
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

### Payments Rules
Follow these rules when working on payments.

It uses Stripe for payments.

- Use github.com/stripe/stripe-go/v76 for Stripe integration
- Handle payment logic in internal/services/stripe-service.go
- Return HTML fragments for HTMX-driven payment flows