package service

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"point-five-backend/internal/config"
	"point-five-backend/internal/model"
	"point-five-backend/internal/repository"
	"point-five-backend/internal/util"
)

func FetchGitHubEvents() ([]model.GitHubEvent, error) {
	req, err := http.NewRequest(config.GetStr, config.ApiURL, nil)
	if err = util.LogError(err); err != nil {
		return nil, err
	}

	req.Header.Set(config.AuthorizationStr, config.BearerStr+config.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err = util.LogError(err); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var events []model.GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func EnrichReposWithStarCount(client *http.Client, events []model.GitHubEvent) ([]model.GitHubEvent, error) {
	for i, event := range events {
		repoURL := event.Repo.URL // Ensure this URL points to the repository details API endpoint
		req, err := http.NewRequest(config.GetStr, repoURL, nil)
		if err = util.LogError(err); err != nil {
			continue
		}

		req.Header.Set(config.AuthorizationStr, config.BearerStr+config.Token)
		resp, err := client.Do(req)
		if err = util.LogError(err); err != nil {
			continue
		}
		defer resp.Body.Close()

		var repoDetails struct {
			StargazersCount int `json:"stargazers_count"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&repoDetails); err != nil {
			continue
		}

		events[i].Repo.StarCount = repoDetails.StargazersCount
	}

	return events, nil
}

func StoreEvents(db *sql.DB, events []model.GitHubEvent) error {
    err := repository.StoreEvents(db, events)
	if err = util.LogError(err); err != nil {
		return err
	}
    // Return nil if no error occurs
    return nil
}

func CreateEventTable(db *sql.DB) {
	repository.CreateEventTable(db)
}

func RetrieveEvents(db *sql.DB) ([]model.GitHubEvent, error) {
	return repository.RetrieveEvents(db)
}