package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"point-five-backend/internal/config"
	"point-five-backend/internal/model"
	"point-five-backend/internal/repository"
	"point-five-backend/internal/service"
	"point-five-backend/internal/util"
    "fmt"
)

// Global variable to hold events
var eventsCache []model.GitHubEvent

func FetchAndStoreEvents(db *sql.DB) error {
    
    repository.CreateEventTable(db)
    events, err := service.FetchGitHubEvents()
    if err != nil {
        util.LogError(err)
        return err 
    }

    client := &http.Client{} // Create an HTTP client for making requests
    enrichedEvents, err := service.EnrichReposWithStarCount(client, events) // Enrich events with star count
    if err != nil {
        util.LogError(err)
        return err 
    }
    fmt.Println(len(enrichedEvents)) 
    
    err = service.StoreEvents(db, enrichedEvents) 
    if err != nil {
        util.LogError(err)
        return err 
    }
    eventsCache = enrichedEvents 

    return nil 
}

func EventsHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Fetch and store GitHub events every time the /events endpoint is called.
        if err := FetchAndStoreEvents(db); err != nil {
            util.LogError(err)
            http.Error(w, "Failed to fetch and store events", http.StatusInternalServerError)
            return
        }

        // Retrieve the latest events from the database after updating
        events, err := service.RetrieveEvents(db)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        w.Header().Set(config.ContentTypeStr, config.ApplicationJsonStr)
        json.NewEncoder(w).Encode(events)
    }
}
