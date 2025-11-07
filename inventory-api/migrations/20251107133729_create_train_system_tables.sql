-- +goose Up
-- +goose StatementBegin
CREATE TABLE stations (
  id UUID PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  city VARCHAR(100) NOT NULL,
  created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPZ NOT NULL DEFAULT NOW()
);

CREATE TABLE trains (
  id UUID PRIMARY KEY DEFAULT,
  name VARCHAR(100) NOT NULL,
  type VARCHAR(50) NOT NULL,
  capacity INT NOT NULL,
  created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS trains;
DROP TABLE IF EXISTS stations;
-- +goose StatementEnd
