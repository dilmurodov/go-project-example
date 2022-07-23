CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_data
(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    age int NOT NULL,
    user_id uuid DEFAULT uuid_generate_v4()
);