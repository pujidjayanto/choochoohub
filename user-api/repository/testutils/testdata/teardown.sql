begin;

-- Disable foreign key checks
set session_replication_role = replica;

-- Truncate all tables
truncate table
  purchases,
  products,
  users
restart identity CASCADE;

-- Enable foreign key checks
set session_replication_role = default;

commit;
