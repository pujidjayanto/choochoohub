-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Example daily schedules
INSERT INTO train_schedules (id, train_id, departure_station_id, destination_station_id, schedule_date)
SELECT
  uuid_generate_v4(),
  t.id,
  s1.id,
  s2.id,
  CURRENT_DATE
FROM trains t
JOIN stations s1 ON (s1.code = 'SBY' AND t.code = 'DOH01')
JOIN stations s2 ON (s2.code = 'ML');

INSERT INTO train_schedules (id, train_id, departure_station_id, destination_station_id, schedule_date)
SELECT
  uuid_generate_v4(),
  t.id,
  s1.id,
  s2.id,
  CURRENT_DATE
FROM trains t
JOIN stations s1 ON (s1.code = 'GMR' AND t.code = 'ARG01')
JOIN stations s2 ON (s2.code = 'BD');

INSERT INTO train_schedules (id, train_id, departure_station_id, destination_station_id, schedule_date)
SELECT
  uuid_generate_v4(),
  t.id,
  s1.id,
  s2.id,
  CURRENT_DATE
FROM trains t
JOIN stations s1 ON (s1.code = 'SMT' AND t.code = 'TWJ01')
JOIN stations s2 ON (s2.code = 'GMR');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM train_schedules
WHERE train_id IN (SELECT id FROM trains WHERE code IN ('DOH01', 'ARG01', 'TWJ01'));
-- +goose StatementEnd
