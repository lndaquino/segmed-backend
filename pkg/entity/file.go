package entity

import "time"

// File struct models a file entity
type File struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	URL       string    `json:"url"`
	Status    *bool     `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}
