package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Travel struct {
	gorm.Model

	Destination string
	StartDate   string
	EndDate     string
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
