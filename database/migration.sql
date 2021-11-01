CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    password TEXT NOT NULL
);