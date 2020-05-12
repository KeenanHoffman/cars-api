CREATE TABLE cars (
	id serial PRIMARY KEY,
	make text,
	model text,
	vin text
);
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO postgres;
