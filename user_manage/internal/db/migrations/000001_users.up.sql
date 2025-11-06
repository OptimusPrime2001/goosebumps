create extension if not exists "pgcrypto";
create table if not  exists users (
	user_id UUID  default gen_random_uuid() primary key,
	username varchar(50) not null,
	email varchar(100) unique not null,
	created_at timestamptz not null default now()
);
