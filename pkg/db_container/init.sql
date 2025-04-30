\c postgres
CREATE EXTENSION IF NOT EXISTS dblink;
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'db_auth') THEN
            PERFORM dblink_exec('dbname=postgres user=' || current_user, 'CREATE DATABASE db_auth');
            RAISE NOTICE 'Created database: %', 'db_auth';
        END IF;
    END;
$$;
\c db_auth
