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
    company_name character varying NOT NULL,
    contact_name character varying NOT NULL,
    contact_title character varying NOT NULL,
    address character varying NOT NULL,
    city character varying NOT NULL,
    postal_code character varying,
    region character varying,
    country character varying NOT NULL,
    phone character varying NOT NULL,
    fax character varying
);

CREATE TABLE public.employees (
    id character varying NOT NULL,
    last_name character varying NOT NULL,
    first_name character varying,
    title character varying NOT NULL,
    title_of_courtesy character varying NOT NULL,
    birth_date date NOT NULL,
    hire_date date NOT NULL,
    address character varying NOT NULL,
    city character varying NOT NULL,
    postal_code character varying NOT NULL,
    country character varying NOT NULL,
    home_phone character varying NOT NULL,
    extension integer NOT NULL,
    notes text NOT NULL,
    recipient_id character varying
);

CREATE TABLE public.orders (
    id character varying NOT NULL,
    order_date date NOT NULL,
    required_date date NOT NULL,
    shipped_date date,
    ship_via integer NOT NULL,
    freight double precision NOT NULL,
    ship_name character varying NOT NULL,
    ship_city character varying NOT NULL,
    ship_region character varying,
    ship_postal_code character varying,
    ship_country character varying NOT NULL,
    customer_id character varying NOT NULL,
    employee_id character varying NOT NULL
);

CREATE TABLE public.products (
    id character varying NOT NULL,
    name character varying NOT NULL,
    qt_per_unit character varying NOT NULL,
    unit_price double precision NOT NULL,
    units_in_stock integer NOT NULL,
    units_on_order integer NOT NULL,
    reorder_level integer NOT NULL,
    discontinued integer NOT NULL,
    supplier_id character varying NOT NULL
);

CREATE TABLE public.suppliers (
    id character varying NOT NULL,
    company_name character varying NOT NULL,
    contact_name character varying NOT NULL,
    contact_title character varying NOT NULL,
    address character varying NOT NULL,
    city character varying NOT NULL,
    region character varying,
    postal_code character varying NOT NULL,
    country character varying NOT NULL,
    phone character varying NOT NULL
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

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_recipient_id_fkey FOREIGN KEY (recipient_id) REFERENCES public.employees(id);

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES public.employees(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_supplier_id_fkey FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) ON DELETE CASCADE;

