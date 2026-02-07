package models

import "time"

type Widget struct {
	ID        string                 `json:"id"`
	PageID    string                 `json:"page_id"`
	Type      string                 `json:"type"`
	Position  int                    `json:"position"`
	Config    map[string]interface{} `json:"config"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
