package domain

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
