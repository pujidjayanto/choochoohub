-- +goose Up
-- +goose StatementBegin
alter table users
add column is_verified boolean not null default false;

create table user_otps (
  id uuid primary key default gen_random_uuid(),
  user_id uuid references users(id) on delete cascade,
  channel varchar(20) not null,              -- 'email' | 'sms'
  destination varchar(255) not null,         -- e.g. user@example.com or phone number
  otp_hash text not null,
  purpose varchar(50) not null,              -- 'signup' | 'signin' | 'password_reset'
  status varchar(20) not null default 'pending', -- pending|verified|max_attempted|expired|user_invalidated
  send_attempts int not null default 1,
  expires_at timestamptz not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

-- indexes for faster OTP lookup
create index idx_user_otps_destination_purpose on user_otps(destination, purpose);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_otps;
alter table users drop column if exists is_verified;
-- +goose StatementEnd
