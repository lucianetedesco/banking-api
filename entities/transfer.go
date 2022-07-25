package entities

import "time"

type Transfer struct {
	ID                   uint      `json:"id"`
	AccountOriginId      uint      `json:"account_origin_id" gorm:"type:integer;not null"`
	AccountDestinationId uint      `json:"account_destination_id" gorm:"type:integer;not null"`
	Amount               float64   `json:"amount" gorm:"not null"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}
