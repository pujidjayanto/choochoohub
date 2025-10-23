-- +goose Up
-- +goose StatementBegin
create table users (
  id uuid primary key,
  email text not null unique,
  password text not null,
  phone varchar(20) not null unique,
  name text not null,
  created_at timestamptz not null,
  updated_at timestamptz not null
);

create index idx_users_name on users (name);
create index idx_users_phone on users (phone);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_users_name;
drop index if exists idx_users_phone;
drop table if exists users;
-- +goose StatementEnd
