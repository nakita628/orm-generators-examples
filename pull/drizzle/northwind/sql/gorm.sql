SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

CREATE EXTENSION IF NOT EXISTS btree_gin WITH SCHEMA public;

COMMENT ON EXTENSION btree_gin IS 'support for indexing common datatypes in GIN';

CREATE EXTENSION IF NOT EXISTS btree_gist WITH SCHEMA public;

COMMENT ON EXTENSION btree_gist IS 'support for indexing common datatypes in GiST';

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';

CREATE EXTENSION IF NOT EXISTS hstore WITH SCHEMA public;

COMMENT ON EXTENSION hstore IS 'data type for storing sets of (key, value) pairs';

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';

CREATE EXTENSION IF NOT EXISTS unaccent WITH SCHEMA public;

COMMENT ON EXTENSION unaccent IS 'text search dictionary that removes accents';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

CREATE EXTENSION IF NOT EXISTS vector WITH SCHEMA public;

COMMENT ON EXTENSION vector IS 'vector data type and ivfflat and hnsw access methods';

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.customers (
    id character varying(5) NOT NULL,
    company_name text NOT NULL,
    contact_name text NOT NULL,
    contact_title text NOT NULL,
    address text NOT NULL,
    city text NOT NULL,
    postal_code text,
    region text,
    country text NOT NULL,
    phone text NOT NULL,
    fax text
);

CREATE TABLE public.employees (
    id text NOT NULL,
    last_name text NOT NULL,
    first_name text,
    title text NOT NULL,
    title_of_courtesy text NOT NULL,
    birth_date date NOT NULL,
    hire_date date NOT NULL,
    address text NOT NULL,
    city text NOT NULL,
    postal_code text NOT NULL,
    country text NOT NULL,
    home_phone text NOT NULL,
    extension bigint NOT NULL,
    notes text NOT NULL,
    recipient_id text
);

CREATE TABLE public.orders (
    id text NOT NULL,
    order_date date NOT NULL,
    required_date date NOT NULL,
    shipped_date date,
    ship_via bigint NOT NULL,
    freight numeric NOT NULL,
    ship_name text NOT NULL,
    ship_city text NOT NULL,
    ship_region text,
    ship_postal_code text,
    ship_country text NOT NULL,
    customer_id character varying(5) NOT NULL,
    employee_id text NOT NULL
);

CREATE TABLE public.products (
    id text NOT NULL,
    name text NOT NULL,
    qt_per_unit text NOT NULL,
    unit_price numeric NOT NULL,
    units_in_stock bigint NOT NULL,
    units_on_order bigint NOT NULL,
    reorder_level bigint NOT NULL,
    discontinued bigint NOT NULL,
    supplier_id text NOT NULL
);

CREATE TABLE public.suppliers (
    id text NOT NULL,
    company_name text NOT NULL,
    contact_name text NOT NULL,
    contact_title text NOT NULL,
    address text NOT NULL,
    city text NOT NULL,
    region text,
    postal_code text NOT NULL,
    country text NOT NULL,
    phone text NOT NULL
);

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_customers_orders FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_employees_orders FOREIGN KEY (employee_id) REFERENCES public.employees(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_employees_other_employees FOREIGN KEY (recipient_id) REFERENCES public.employees(id);

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_suppliers_products FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) ON DELETE CASCADE;

