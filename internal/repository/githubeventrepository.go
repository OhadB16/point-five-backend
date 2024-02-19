package repository

import (
	"database/sql"
	"point-five-backend/internal/model"
	"point-five-backend/internal/util"
	"fmt"
)

func CreateEventTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS events (
		"id" TEXT NOT NULL PRIMARY KEY,
		"type" TEXT,
		"actor_id" INTEGER,
		"actor_login" TEXT,
		"actor_avatar_url" TEXT,    
		"repo_id" INTEGER,
		"repo_name" TEXT,
		"repo_url" TEXT,
		"repo_star_count" INTEGER DEFAULT 0,
		"payload_action" TEXT,
		"public" BOOLEAN,
		"created_at" DATETIME
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		util.LogError(err, "Failed to create events table")
	}
}

func StoreEvents(db *sql.DB, events []model.GitHubEvent) error {
    for _, event := range events {
        _, err := db.Exec(`INSERT OR IGNORE INTO events 
            (id, type, actor_id, actor_login, actor_avatar_url, repo_id, repo_name, repo_url, repo_star_count, payload_action, public, created_at) 
            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, // Ensure there are 12 placeholders
            event.ID, event.Type, event.Actor.ID, event.Actor.Login, event.Actor.AvatarURL,
            event.Repo.ID, event.Repo.Name, event.Repo.URL, event.Repo.StarCount, 
            event.Payload.Action, event.Public, event.CreatedAt)

        if err != nil {
            util.LogError(err, fmt.Sprintf("Failed to store event %s", event.ID))
            return err
        }
    }
    return nil
}
func RetrieveEvents(db *sql.DB) ([]model.GitHubEvent, error) {
    var events []model.GitHubEvent

    query := `SELECT id, type, actor_id, actor_login, actor_avatar_url, repo_id, repo_name, repo_url, repo_star_count, payload_action, public, created_at FROM events`
    rows, err := db.Query(query)
    if err != nil {
        util.LogError(err, "Error querying events from database")
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var event model.GitHubEvent
        var actor model.GitHubActor
        var repo model.GitHubRepo
        var payload model.GitHubPayload
        var repoStarCount int  // Local variable to hold the repo star count

        // Adjust the scanning based on your database schema
        err := rows.Scan(&event.ID, &event.Type, &actor.ID, &actor.Login, &actor.AvatarURL, 
                         &repo.ID, &repo.Name, &repo.URL, &repoStarCount,  // Scan into repoStarCount
                         &payload.Action, &event.Public, &event.CreatedAt)
        if err != nil {
            util.LogError(err, "Error scanning event from database row")
            return nil, err
        }

        // Assign nested structs
        repo.StarCount = repoStarCount  // Assign repoStarCount to repo.StarCount
        event.Actor = actor
        event.Repo = repo
        event.Payload = payload

        events = append(events, event)
    }

    if err = rows.Err(); err != nil {
        util.LogError(err, "Error iterating over database rows")
        return nil, err
    }

    return events, nil
}


func scanEvent(rows *sql.Rows, event *model.GitHubEvent) error {
    var (
        actorID      int
        actorLogin   string
        actorAvatar  string
        repoID       int
        repoName     string
        repoURL      string
        repoStarCount int  
        payloadAction string
    )

    err := rows.Scan(
        &event.ID,
        &event.Type,
        &actorID,
        &actorLogin,
        &actorAvatar,
        &repoID,
        &repoName,
        &repoURL,
        &repoStarCount,  
        &payloadAction,
        &event.Public,
        &event.CreatedAt,
    )
    if err != nil {
        return err
    }

    // Assign the nested struct fields after scanning
    event.Actor = model.GitHubActor{
        ID:        actorID,
        Login:     actorLogin,
        AvatarURL: actorAvatar,
    }
    event.Repo = model.GitHubRepo{
        ID:        repoID,
        Name:      repoName,
        URL:       repoURL,
        StarCount: repoStarCount,  // Assign repoStarCount to StarCount
    }
    event.Payload = model.GitHubPayload{
        Action: payloadAction,
    }

    return nil
}

