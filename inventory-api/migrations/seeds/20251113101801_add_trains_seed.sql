-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO trains (
  id, name, category, capacity, code, class_name, class_code, subclass_name, subclass_code, direction
)
VALUES
  (uuid_generate_v4(), 'Doho Express', 'Local', 400, 'DOH01', 'Economy', 'ECO', 'AC', 'AC', 'Surabaya - Malang'),
  (uuid_generate_v4(), 'Argo Parahyangan', 'Intercity', 500, 'ARG01', 'Executive', 'EXE', 'Premium', 'PRM', 'Jakarta - Bandung'),
  (uuid_generate_v4(), 'Tawang Jaya', 'Intercity', 600, 'TWJ01', 'Business', 'BIS', 'AC', 'AC', 'Semarang - Jakarta');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM trains
WHERE code IN ('DOH01', 'ARG01', 'TWJ01');
-- +goose StatementEnd
