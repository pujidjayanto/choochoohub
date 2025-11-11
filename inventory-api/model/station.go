package model

import "time"

type Station struct {
	ID        string    `db:"id"`
	Code      string    `db:"code"`
	Name      string    `db:"name"`
	City      string    `db:"city"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
