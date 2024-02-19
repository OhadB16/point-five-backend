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
        events, err := service.RetrieveEvents(db)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set(config.ContentTypeStr, config.ApplicationJsonStr)
        json.NewEncoder(w).Encode(events)
    }
}
