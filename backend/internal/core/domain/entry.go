package domain

import "time"

type EventType int

const (
	Personal EventType = iota
	Statement
	Event
	Exercise
)

type Entry struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	EventType   EventType     `json:"eventType"`
	EventDate   time.Time     `json:"eventDate"`
	Origin      string        `json:"origin"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	Score       float32       `json:"score"`
	Positive    *bool         `json:"positive"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}
