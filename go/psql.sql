CREATE TABLE public.users(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	address TEXT NOT NULL,
	avatar TEXT  NOT NULL
);