-- +goose Up
-- +goose StatementBegin
CREATE TABLE stations (
  id UUID PRIMARY KEY,
  code VARCHAR(10) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL,
  city VARCHAR(100) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE trains (
  id UUID PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  category VARCHAR(50) NOT NULL,
  capacity INT NOT NULL,
  code VARCHAR(20) NOT NULL,
  class_name VARCHAR(50) NOT NULL, -- Economy, Executive, Business
  class_code VARCHAR(10) NOT NULL, -- ECO, EXE, BIS
  subclass_name VARCHAR(100), -- AC, Non-AC, Premium, etc.
  subclass_code VARCHAR(20),
  created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPZ NOT NULL DEFAULT NOW()
);

CREATE TABLE train_stops (
  id UUID PRIMARY KEY,
  train_id INT REFERENCES trains(id) ON DELETE CASCADE,
  station_id INT REFERENCES stations(id),
  stop_order INT NOT NULL,          -- e.g. 1, 2, 3, ...
  arrival_time TIME,                -- nullable for first station
  departure_time TIME,              -- nullable for last station
  created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  UNIQUE(train_id, stop_order)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS train_stops;
DROP TABLE IF EXISTS trains;
DROP TABLE IF EXISTS stations;
-- +goose StatementEnd
