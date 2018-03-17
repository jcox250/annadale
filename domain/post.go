package domain

import "time"

type PostRepo interface{}

type Post struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	AddedBy     string    `json:"added_by,omitempty"`
	AddedDate   time.Time `json:"added_date,omitempty"`
	UpdatedBy   string    `json:"updated_by,omitempty"`
	UpdatedDate time.Time `json:"updated_date,omitempty"`
	Archive     bool      `json:"archive,omitempty"`
}
