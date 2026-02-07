package models

import "time"

type Page struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Route     string    `json:"route"`
	IsHome    bool      `json:"is_home"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
