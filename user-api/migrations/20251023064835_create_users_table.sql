-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  user_type VARCHAR(20) NOT NULL DEFAULT 'unverified',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_users_email ON users (email);

CREATE TABLE user_profiles (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  phone VARCHAR(20) UNIQUE,
  name VARCHAR(255) NOT NULL,
  dob DATE NOT NULL,
  gender CHAR(1) NOT NULL,
  identity_number VARCHAR(50) NOT NULL,
  identity_type VARCHAR(10) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX idx_user_profiles_user_id ON user_profiles(user_id);

CREATE UNIQUE INDEX idx_unique_non_null_phone
ON user_profiles (phone) WHERE phone IS NOT NULL;

CREATE UNIQUE INDEX idx_unique_non_null_identity
ON user_profiles (identity_type, identity_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
