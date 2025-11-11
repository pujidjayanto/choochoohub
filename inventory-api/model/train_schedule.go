package model

import "time"

type TrainSchedule struct {
	ID                   string    `db:"id"`
	TrainID              string    `db:"train_id"`
	DepartureStationID   string    `db:"departure_station_id"`
	DestinationStationID string    `db:"destination_station_id"`
	ScheduleDate         time.Time `db:"schedule_date"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
