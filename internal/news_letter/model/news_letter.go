package model

import (
	edition_model "go_newsletter_api/internal/edition/model"
)

type Newsletter struct {
	ID       uint                    `json:"id"`
	Name     string                  `json:"name"`
	AdminID  uint                    `json:"admin_id"`
	Editions []edition_model.Edition `gorm:"foreignkey:NewsletterID;constraint:OnDelete:CASCADE"`
}

type NewsletterSubscriber struct {
	ID           uint `gorm:"primary_key" json:"id"`
	NewsletterID uint `json:"newsletter_id"`
	UserID       uint `json:"user_id"`
}
