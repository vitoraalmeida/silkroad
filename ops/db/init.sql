\c silkroad

CREATE TABLE IF NOT EXISTS category (
    id serial CONSTRAINT pk_id_category PRIMARY KEY,
    name text NOT NULL,
    created_at  timestamp with time zone NOT NULL DEFAULT now(),
    updated_at  timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS product (
    id serial CONSTRAINT pk_id_product PRIMARY KEY,
    name text NOT NULL,
    category_id INT, 
    price numeric(9,2) NOT NULL,
    stock INT NOT NULL,
    available BOOLEAN NOT NULL,
    created_at  timestamp with time zone NOT NULL DEFAULT now(),
    updated_at  timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT fk_category
        FOREIGN KEY (category_id)
        REFERENCES category (id)
);

CREATE TABLE IF NOT EXISTS customer (
    id serial CONSTRAINT pk_id_customer PRIMARY KEY,
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    cpf text UNIQUE NOT NULL,
    password text NOT NULL,
    created_at  timestamp with time zone NOT NULL DEFAULT now(),
    updated_at  timestamp with time zone NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS sale (
    id serial CONSTRAINT pk_id_sale PRIMARY KEY,
    customer_id INT, 
    total_amount numeric(9,2) NOT NULL,
    created_at  timestamp with time zone NOT NULL DEFAULT now(),
    updated_at  timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT fk_customer
        FOREIGN KEY (customer_id)
        REFERENCES customer (id)
);

CREATE OR REPLACE FUNCTION trigger_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_timestamp_category
BEFORE UPDATE ON category
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();

CREATE TRIGGER update_timestamp_product
BEFORE UPDATE ON product
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();

CREATE TRIGGER update_timestamp_customer
BEFORE UPDATE ON customer
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();

CREATE TRIGGER update_timestamp_sale
BEFORE UPDATE ON sale
FOR EACH ROW
EXECUTE PROCEDURE trigger_update_timestamp();
