CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) NOT NULL PRIMARY KEY UNIQUE,
    name VARCHAR(12) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    access_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE FUNCTION set_update_time() RETURNS TRIGGER AS '
  BEGIN
    NEW.updated_at := ''now'';
    return NEW;
  END;
' LANGUAGE 'plpgsql';

create trigger update_tri before update on users for each row
  execute procedure set_update_time();
