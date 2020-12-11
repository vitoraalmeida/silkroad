\c silkroad

CREATE TABLE IF NOT EXISTS category (
    id serial CONSTRAINT pk_id_category PRIMARY KEY,
    name text NOT NULL,
    created_at  timestamp with time zone NOT NULL DEFAULT now(),
    updated_at  timestamp with time zone NOT NULL DEFAULT now()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON category
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

