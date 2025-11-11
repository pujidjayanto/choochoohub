package model

import "time"

type TrainRoute struct {
	ID              string    `db:"id"`
	TrainID         string    `db:"train_id"`
	StationID       string    `db:"station_id"`
	RouteOrder      int       `db:"route_order"`
	ArrivalTime     string    `db:"arrival_time"`   // store as "HH:MM:SS"
	DepartureTime   string    `db:"departure_time"` // store as "HH:MM:SS"
	CumulativePrice float64   `db:"cumulative_price"`
	DistanceKm      float64   `db:"distance_km"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
