package model

import "time"

type GitHubEvent struct {
	ID        string        `json:"id"`
	Type      string        `json:"type"`
	Actor     GitHubActor   `json:"actor"`
	Repo      GitHubRepo    `json:"repo"`
	Payload   GitHubPayload `json:"payload"` // New field
	Public    bool          `json:"public"`  // New field
	CreatedAt time.Time     `json:"created_at"`
}

type GitHubActor struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"` // New field
}

type GitHubRepo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"` // New field
	StarCount int    `json:"stargazers_count,omitempty"` // Add this line to include star count
}

// Define a new struct for the payload, adjust according to the actual structure
type GitHubPayload struct {
	Action string `json:"action"` // Example field, adjust based on actual payload data
}
