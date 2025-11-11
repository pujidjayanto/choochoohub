package model

import "time"

type Train struct {
	ID           string    `db:"id"`
	Name         string    `db:"name"`
	Category     string    `db:"category"`
	Capacity     int       `db:"capacity"`
	Code         string    `db:"code"`
	ClassName    string    `db:"class_name"`
	ClassCode    string    `db:"class_code"`
	SubclassName string    `db:"subclass_name"`
	SubclassCode string    `db:"subclass_code"`
	Direction    string    `db:"direction"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
