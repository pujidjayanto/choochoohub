package repository

import (
	"context"

	"github.com/pujidjayanto/choochoohub/inventory-api/model"
	"github.com/pujidjayanto/choochoohub/inventory-api/pkg/db"
)

type StationRepository interface {
	List(ctx context.Context) ([]*model.Station, error)
}

type stationRepository struct {
	db db.DatabaseHandler
}

func NewStationRepository(db db.DatabaseHandler) StationRepository {
	return &stationRepository{db: db}
}

func (r *stationRepository) List(ctx context.Context) ([]*model.Station, error) {
	query := `
		SELECT id, code, name, city, created_at, updated_at
		FROM stations
		ORDER BY name
	`

	rows, err := r.db.GetPool(ctx).Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stations []*model.Station
	for rows.Next() {
		s := &model.Station{}
		if err := rows.Scan(&s.ID, &s.Code, &s.Name, &s.City, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		stations = append(stations, s)
	}

	return stations, nil
}
