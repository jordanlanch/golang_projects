package dbmodels

import "time"

//Model default db model struct
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"-"`
}

// Server db struct
type Server struct {
	Model
	VIN        string `gorm:"type:varchar(30);unique_index" json:"vin"`
	Plate      string `gorm:"not null" json:"plate"`
	Status     string `gorm:"not null" json:"status"`
	WebhookURL string `gorm:"not null" json:"webhook_url"`
	Active     bool   `gorm:"not null;default:true" json:"active"`
}
