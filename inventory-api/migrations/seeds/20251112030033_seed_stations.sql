-- +goose Up
-- +goose StatementBegin
-- Ensure uuid-ossp extension exists
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Insert initial stations
INSERT INTO stations (id, code, name, city)
VALUES
  (uuid_generate_v4(), 'SBY', 'Stasiun Surabaya Gubeng', 'Surabaya'),
  (uuid_generate_v4(), 'ML', 'Stasiun Malang', 'Malang'),
  (uuid_generate_v4(), 'PWK', 'Stasiun Purwokerto', 'Purwokerto'),
  (uuid_generate_v4(), 'YK', 'Stasiun Yogyakarta', 'Yogyakarta'),
  (uuid_generate_v4(), 'SMT', 'Stasiun Semarang Tawang', 'Semarang'),
  (uuid_generate_v4(), 'BD', 'Stasiun Bandung', 'Bandung'),
  (uuid_generate_v4(), 'GMR', 'Stasiun Gambir', 'Jakarta');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM stations
WHERE code IN ('SBY', 'ML', 'PWK', 'YK', 'SMT', 'BD', 'GMR');
-- +goose StatementEnd
