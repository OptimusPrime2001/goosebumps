create table if not  exists users (
	user_id UUID  default gen_random_uuid() primary key,
	name varchar(50) not null,
	email varchar(100) unique not null
);
