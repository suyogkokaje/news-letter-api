package model

import (
	"time"
)

type Edition struct {
	ID            uint      `json:"id"`
	EditionNumber int       `json:"edition_number"`
	Name          string    `json:"name"`
	NewsletterID  uint      `json:"newsletter_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	PublishAt     time.Time `json:"publish_at"`
}
