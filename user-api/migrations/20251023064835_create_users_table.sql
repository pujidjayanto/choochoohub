-- +goose Up
-- +goose StatementBegin
create table public.users (
  id uuid primary key,
  email varchar(255) not null unique,
  password_hash text not null,
  user_type varchar(20) not null default 'starter',
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create index idx_users_email on users(email);

create table public.user_profiles (
  id uuid primary key,
  user_id uuid not null references users(id) on delete cascade,
  phone varchar(20) unique,
  name varchar(255) not null,
  dob date not null,
  gender char(1) not null,
  identity_number varchar(50) not null,
  identity_type varchar(10) not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create unique index idx_user_profiles_user_id on user_profiles(user_id);

create unique index idx_unique_non_null_phone
on user_profiles(phone) where phone is not null;

create unique index idx_unique_non_null_identity
on user_profiles(identity_type, identity_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_profiles;
drop table if exists users;
-- +goose StatementEnd
