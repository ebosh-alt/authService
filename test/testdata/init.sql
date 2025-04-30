\c postgres
CREATE EXTENSION IF NOT EXISTS dblink;
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'auth') THEN
            PERFORM dblink_exec('dbname=postgres user=' || current_user, 'CREATE DATABASE auth');
        END IF;
    END
$$;
\c auth
DO
$$
    BEGIN
        CREATE TABLE IF NOT EXISTS users
        (
            id            SERIAL PRIMARY KEY,
            telegram_id   BIGINT UNIQUE NOT NULL,
            first_name    VARCHAR(255),
            access_token  VARCHAR(255),
            refresh_token VARCHAR(255),
            created_at    TIMESTAMP DEFAULT NOW(),
            updated_at    TIMESTAMP DEFAULT NOW()
        );


        RAISE NOTICE 'Таблицы успешно созданы.';
    EXCEPTION
        WHEN OTHERS THEN
            RAISE EXCEPTION 'Ошибка при создании таблиц: %', SQLERRM;
    END
$$;