package domain

import "time"

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       float32
	Quantity    uint16
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
