package model

type Newsletter struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	AdminID uint   `json:"admin_id"`

	// Define a one-to-many relationship with NewsletterSubscriber
	Subscribers []NewsletterSubscriber `gorm:"foreignkey:NewsletterID"`
}

type NewsletterSubscriber struct {
	ID           uint `gorm:"primary_key" json:"id"`
	NewsletterID uint `json:"newsletter_id"` // Foreign key to Newsletter
	UserID       uint `json:"user_id"`

	// Define a reference to the Newsletter model
	Newsletter Newsletter `gorm:"foreignkey:NewsletterID"`
}
