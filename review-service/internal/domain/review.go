package domain

import "time"

type Review struct {
	ID        uint64
	ProductID uint64
	UserID    uint64
	Rating    float64
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
