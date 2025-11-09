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
  direction VARCHAR(100) NOT NULL, -- Surabaya - Malang
  created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPZ NOT NULL DEFAULT NOW()
);

CREATE TABLE train_routes (
  id UUID PRIMARY KEY,
  train_id INT REFERENCES trains(id) ON DELETE CASCADE,
  station_id INT REFERENCES stations(id),
  route_order INT NOT NULL,          -- e.g. 1, 2, 3, ...
  arrival_time TIME,                -- nullable for first station
  departure_time TIME,              -- nullable for last station
  cumulative_price NUMERIC(12,2) NOT NULL,
  distance_km NUMERIC(8,2) NOT NULL DEFAULT 0, -- distance from previous station in km
  created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  UNIQUE(train_id, route_order)
);

CREATE TABLE train_schedules (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  train_id UUID NOT NULL REFERENCES trains(id) ON DELETE CASCADE,
  departure_station_id UUID NOT NULL REFERENCES stations(id),
  destination_station_id UUID NOT NULL REFERENCES stations(id),
  schedule_date DATE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS train_schedules;
DROP TABLE IF EXISTS train_routes;
DROP TABLE IF EXISTS trains;
DROP TABLE IF EXISTS stations;
-- +goose StatementEnd
