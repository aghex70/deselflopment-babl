package ports

import "time"

type CreateEntryRequest struct {
	Name        string    `json:"name"`
	EventType   int       `json:"eventType"`
	EventDate   time.Time `json:"eventDate"`
	Origin      string    `json:"origin"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Score       float32   `json:"score"`
	Positive    *bool     `json:"positive"`
}

type UpdateEntryRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type UpdateUserRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
