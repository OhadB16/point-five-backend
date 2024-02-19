package model

import "time"

type GitHubEvent struct {
	ID        string        `json:"id"`
	Type      string        `json:"type"`
	Actor     GitHubActor   `json:"actor"`
	Repo      GitHubRepo    `json:"repo"`
	Payload   GitHubPayload `json:"payload"` 
	Public    bool          `json:"public"`  
	CreatedAt time.Time     `json:"created_at"`
}

type GitHubActor struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"` 
}

type GitHubRepo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"` 
	StarCount int    `json:"stargazers_count,omitempty"` 
}

type GitHubPayload struct {
	Action string `json:"action"` 
}
