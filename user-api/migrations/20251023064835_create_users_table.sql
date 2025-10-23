-- +goose Up
-- +goose StatementBegin
create table users (
  id uuid primary key,
  email varchar(255) not null unique,
  password text not null,
  phone varchar(20) not null unique,
  name varchar(255) not null,
  dob date,
  gender char(1) not null,
  identity_number varchar(50),
  identity_type varchar(20),
  created_at timestamptz not null,
  updated_at timestamptz not null
);

create index idx_users_name on users (name);
create index idx_users_phone on users (phone);
create index idx_users_identity_number on users (identity_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
