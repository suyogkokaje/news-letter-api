package model

type Newsletter struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	AdminID uint   `json:"admin_id"`
}

type NewsletterSubscriber struct {
	ID           uint `gorm:"primary_key" json:"id"`
	NewsletterID uint `json:"newsletter_id"`
	UserID       uint `json:"user_id"`
}
