#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';
    CREATE DATABASE $POSTGRES_DB;
    GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER
    \connect $POSTGRES_DB
    BEGIN;
        CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

        CREATE TABLE IF NOT EXISTS user_data
        (
            id SERIAL PRIMARY KEY,
            name VARCHAR NOT NULL,
            age int NOT NULL,
            user_id uuid DEFAULT uuid_generate_v4()
        );
        INSERT INTO user_data (name, age) VALUES
                                ('Tolib', 19),
                                ('Anvar', 13),
                                ('Otabek', 24);
    COMMIT;
EOSQL