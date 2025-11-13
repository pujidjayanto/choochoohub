-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Example routes for Doho Express (Surabaya - Malang)
INSERT INTO train_routes (id, train_id, station_id, route_order, arrival_time, departure_time, cumulative_price, distance_km)
SELECT
  uuid_generate_v4(),
  t.id,
  s.id,
  r.route_order,
  r.arrival_time,
  r.departure_time,
  r.cumulative_price,
  r.distance_km
FROM trains t
JOIN (
  VALUES
    (1, 'SBY', NULL::TIME, '06:00'::TIME, 0, 0),
    (2, 'ML', '08:00'::TIME, NULL::TIME, 50000, 120)
) AS r(route_order, code, arrival_time, departure_time, cumulative_price, distance_km)
JOIN stations s ON s.code = r.code
WHERE t.code = 'DOH01';

-- Routes for Argo Parahyangan (Jakarta - Bandung)
INSERT INTO train_routes (id, train_id, station_id, route_order, arrival_time, departure_time, cumulative_price, distance_km)
SELECT
  uuid_generate_v4(),
  t.id,
  s.id,
  r.route_order,
  r.arrival_time,
  r.departure_time,
  r.cumulative_price,
  r.distance_km
FROM trains t
JOIN (
  VALUES
    (1, 'GMR', NULL::TIME, '07:00'::TIME, 0, 0),
    (2, 'BD', '09:30'::TIME, NULL::TIME, 150000, 150)
) AS r(route_order, code, arrival_time, departure_time, cumulative_price, distance_km)
JOIN stations s ON s.code = r.code
WHERE t.code = 'ARG01';

-- Routes for Tawang Jaya (Semarang - Jakarta)
INSERT INTO train_routes (id, train_id, station_id, route_order, arrival_time, departure_time, cumulative_price, distance_km)
SELECT
  uuid_generate_v4(),
  t.id,
  s.id,
  r.route_order,
  r.arrival_time,
  r.departure_time,
  r.cumulative_price,
  r.distance_km
FROM trains t
JOIN (
  VALUES
    (1, 'SMT', NULL::TIME, '06:00'::TIME, 0, 0),
    (2, 'YK', '07:30'::TIME, '07:40'::TIME, 50000, 100),
    (3, 'GMR', '10:00'::TIME, NULL::TIME, 100000, 200)
) AS r(route_order, code, arrival_time, departure_time, cumulative_price, distance_km)
JOIN stations s ON s.code = r.code
WHERE t.code = 'TWJ01';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM train_routes
WHERE train_id IN (SELECT id FROM trains WHERE code IN ('DOH01', 'ARG01', 'TWJ01'));
-- +goose StatementEnd
