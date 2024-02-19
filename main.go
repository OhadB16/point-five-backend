package main

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/rs/cors"
    "point-five-backend/internal/controller"
    "point-five-backend/internal/util" 
    "point-five-backend/internal/config" 
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open(config.SQLite3Str, config.EventsDBPath)
	if err != nil {
		util.LogError(err) 
		return
	}
    defer db.Close()

    // Set up the HTTP server routing with CORS
    mux := http.NewServeMux()
    mux.HandleFunc(config.EventsPath, controller.EventsHandler(db))

    // Configure CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{config.GetStr},
        AllowedHeaders:   []string{config.AcceptStr, config.ContentTypeStr, config.AuthorizationStr},
        AllowCredentials: true,
        Debug:            false,
    })

    // Start the HTTP server with CORS handler
    handler := c.Handler(mux)
    log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		util.LogError(err) 
	}
}
