package domain

import "time"

type EntryType string

const (
	Exercise  EntryType = "exercise"
	Statement EntryType = "statement"
	Event     EntryType = "event"
)

type Entry struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	EntryType   EntryType     `json:"entryType"`
	EventDate   time.Time     `json:"eventDate"`
	Origin      string        `json:"origin"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	Score       float32       `json:"score"`
	Positive    *bool         `json:"positive"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}
