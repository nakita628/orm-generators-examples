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

CREATE TABLE public.auth_group (
    id integer NOT NULL,
    name character varying(150) NOT NULL
);

CREATE SEQUENCE public.auth_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.auth_group_id_seq OWNED BY public.auth_group.id;

CREATE TABLE public.auth_group_permissions (
    id integer NOT NULL,
    group_id integer NOT NULL,
    permission_id integer NOT NULL
);

CREATE SEQUENCE public.auth_group_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.auth_group_permissions_id_seq OWNED BY public.auth_group_permissions.id;

CREATE TABLE public.auth_permission (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    content_type_id integer NOT NULL,
    codename character varying(100) NOT NULL
);

CREATE SEQUENCE public.auth_permission_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.auth_permission_id_seq OWNED BY public.auth_permission.id;

CREATE TABLE public.auth_user (
    id integer NOT NULL,
    password character varying(128) NOT NULL,
    last_login timestamp with time zone,
    is_superuser boolean NOT NULL,
    username character varying(150) NOT NULL,
    first_name character varying(150) NOT NULL,
    last_name character varying(150) NOT NULL,
    email character varying(254) NOT NULL,
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
    date_joined timestamp with time zone NOT NULL
);

CREATE TABLE public.auth_user_groups (
    id integer NOT NULL,
    user_id integer NOT NULL,
    group_id integer NOT NULL
);

CREATE SEQUENCE public.auth_user_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.auth_user_groups_id_seq OWNED BY public.auth_user_groups.id;

CREATE SEQUENCE public.auth_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.auth_user_id_seq OWNED BY public.auth_user.id;

CREATE TABLE public.auth_user_user_permissions (
    id integer NOT NULL,
    user_id integer NOT NULL,
    permission_id integer NOT NULL
);

CREATE SEQUENCE public.auth_user_user_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.auth_user_user_permissions_id_seq OWNED BY public.auth_user_user_permissions.id;

CREATE TABLE public.django_admin_log (
    id integer NOT NULL,
    action_time timestamp with time zone NOT NULL,
    user_id integer NOT NULL,
    content_type_id integer,
    object_id character varying,
    object_repr character varying(200) NOT NULL,
    action_flag smallint NOT NULL,
    change_message character varying NOT NULL
);

CREATE SEQUENCE public.django_admin_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_admin_log_id_seq OWNED BY public.django_admin_log.id;

CREATE TABLE public.django_content_type (
    id integer NOT NULL,
    app_label character varying(100) NOT NULL,
    model character varying(100) NOT NULL
);

CREATE SEQUENCE public.django_content_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_content_type_id_seq OWNED BY public.django_content_type.id;

CREATE TABLE public.django_session (
    session_key character varying(40) NOT NULL,
    session_data character varying NOT NULL,
    expire_date timestamp with time zone NOT NULL
);

CREATE TABLE public.oidc_provider_client (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    owner_id integer,
    client_type character varying(30) NOT NULL,
    client_id character varying(255) NOT NULL,
    client_secret character varying(255) NOT NULL,
    jwt_alg character varying(10) NOT NULL,
    date_created date NOT NULL,
    website_url character varying(255) NOT NULL,
    terms_url character varying(255) NOT NULL,
    contact_email character varying(255) NOT NULL,
    logo character varying(100) NOT NULL,
    reuse_consent boolean NOT NULL,
    require_consent boolean NOT NULL,
    _redirect_uris character varying NOT NULL,
    _post_logout_redirect_uris character varying NOT NULL,
    _scope character varying NOT NULL
);

CREATE SEQUENCE public.oidc_provider_client_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.oidc_provider_client_id_seq OWNED BY public.oidc_provider_client.id;

CREATE TABLE public.oidc_provider_client_response_types (
    id bigint NOT NULL,
    client_id bigint NOT NULL,
    responsetype_id bigint NOT NULL
);

CREATE SEQUENCE public.oidc_provider_client_response_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.oidc_provider_client_response_types_id_seq OWNED BY public.oidc_provider_client_response_types.id;

CREATE TABLE public.oidc_provider_code (
    id bigint NOT NULL,
    client_id bigint NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    _scope character varying NOT NULL,
    user_id integer NOT NULL,
    code character varying(255) NOT NULL,
    nonce character varying(255) NOT NULL,
    is_authentication boolean NOT NULL,
    code_challenge character varying(255),
    code_challenge_method character varying(255)
);

CREATE SEQUENCE public.oidc_provider_code_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.oidc_provider_code_id_seq OWNED BY public.oidc_provider_code.id;

CREATE TABLE public.oidc_provider_responsetype (
    id bigint NOT NULL,
    value character varying(30) NOT NULL,
    description character varying(50) NOT NULL
);

CREATE SEQUENCE public.oidc_provider_responsetype_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.oidc_provider_responsetype_id_seq OWNED BY public.oidc_provider_responsetype.id;

CREATE TABLE public.oidc_provider_rsakey (
    id bigint NOT NULL,
    key character varying NOT NULL
);

CREATE SEQUENCE public.oidc_provider_rsakey_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.oidc_provider_rsakey_id_seq OWNED BY public.oidc_provider_rsakey.id;

CREATE TABLE public.oidc_provider_token (
    id bigint NOT NULL,
    client_id bigint NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    _scope character varying NOT NULL,
    user_id integer,
    access_token character varying(255) NOT NULL,
    refresh_token character varying(255) NOT NULL,
    _id_token character varying NOT NULL
);

CREATE SEQUENCE public.oidc_provider_token_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.oidc_provider_token_id_seq OWNED BY public.oidc_provider_token.id;

CREATE TABLE public.oidc_provider_userconsent (
    id bigint NOT NULL,
    client_id bigint NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    _scope character varying NOT NULL,
    user_id integer NOT NULL,
    date_given timestamp with time zone NOT NULL
);

CREATE SEQUENCE public.oidc_provider_userconsent_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.oidc_provider_userconsent_id_seq OWNED BY public.oidc_provider_userconsent.id;

ALTER TABLE ONLY public.auth_group ALTER COLUMN id SET DEFAULT nextval('public.auth_group_id_seq'::regclass);

ALTER TABLE ONLY public.auth_group_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_group_permissions_id_seq'::regclass);

ALTER TABLE ONLY public.auth_permission ALTER COLUMN id SET DEFAULT nextval('public.auth_permission_id_seq'::regclass);

ALTER TABLE ONLY public.auth_user ALTER COLUMN id SET DEFAULT nextval('public.auth_user_id_seq'::regclass);

ALTER TABLE ONLY public.auth_user_groups ALTER COLUMN id SET DEFAULT nextval('public.auth_user_groups_id_seq'::regclass);

ALTER TABLE ONLY public.auth_user_user_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_user_user_permissions_id_seq'::regclass);

ALTER TABLE ONLY public.django_admin_log ALTER COLUMN id SET DEFAULT nextval('public.django_admin_log_id_seq'::regclass);

ALTER TABLE ONLY public.django_content_type ALTER COLUMN id SET DEFAULT nextval('public.django_content_type_id_seq'::regclass);

ALTER TABLE ONLY public.oidc_provider_client ALTER COLUMN id SET DEFAULT nextval('public.oidc_provider_client_id_seq'::regclass);

ALTER TABLE ONLY public.oidc_provider_client_response_types ALTER COLUMN id SET DEFAULT nextval('public.oidc_provider_client_response_types_id_seq'::regclass);

ALTER TABLE ONLY public.oidc_provider_code ALTER COLUMN id SET DEFAULT nextval('public.oidc_provider_code_id_seq'::regclass);

ALTER TABLE ONLY public.oidc_provider_responsetype ALTER COLUMN id SET DEFAULT nextval('public.oidc_provider_responsetype_id_seq'::regclass);

ALTER TABLE ONLY public.oidc_provider_rsakey ALTER COLUMN id SET DEFAULT nextval('public.oidc_provider_rsakey_id_seq'::regclass);

ALTER TABLE ONLY public.oidc_provider_token ALTER COLUMN id SET DEFAULT nextval('public.oidc_provider_token_id_seq'::regclass);

ALTER TABLE ONLY public.oidc_provider_userconsent ALTER COLUMN id SET DEFAULT nextval('public.oidc_provider_userconsent_id_seq'::regclass);

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_name_key UNIQUE (name);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_username_key UNIQUE (username);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_session
    ADD CONSTRAINT django_session_pkey PRIMARY KEY (session_key);

ALTER TABLE ONLY public.oidc_provider_client
    ADD CONSTRAINT oidc_provider_client_client_id_key UNIQUE (client_id);

ALTER TABLE ONLY public.oidc_provider_client
    ADD CONSTRAINT oidc_provider_client_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.oidc_provider_client_response_types
    ADD CONSTRAINT oidc_provider_client_response_types_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.oidc_provider_code
    ADD CONSTRAINT oidc_provider_code_code_key UNIQUE (code);

ALTER TABLE ONLY public.oidc_provider_code
    ADD CONSTRAINT oidc_provider_code_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.oidc_provider_responsetype
    ADD CONSTRAINT oidc_provider_responsetype_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.oidc_provider_responsetype
    ADD CONSTRAINT oidc_provider_responsetype_value_key UNIQUE (value);

ALTER TABLE ONLY public.oidc_provider_rsakey
    ADD CONSTRAINT oidc_provider_rsakey_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.oidc_provider_token
    ADD CONSTRAINT oidc_provider_token_access_token_key UNIQUE (access_token);

ALTER TABLE ONLY public.oidc_provider_token
    ADD CONSTRAINT oidc_provider_token_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.oidc_provider_token
    ADD CONSTRAINT oidc_provider_token_refresh_token_key UNIQUE (refresh_token);

ALTER TABLE ONLY public.oidc_provider_userconsent
    ADD CONSTRAINT oidc_provider_userconsent_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT "fk-auth_group_permissions-group_id" FOREIGN KEY (group_id) REFERENCES public.auth_group(id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT "fk-auth_group_permissions-permission_id" FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT "fk-auth_permission-content_type_id" FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT "fk-auth_user_groups-group_id" FOREIGN KEY (group_id) REFERENCES public.auth_group(id);

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT "fk-auth_user_groups-user_id" FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT "fk-auth_user_user_permissions-permission_id" FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id);

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT "fk-auth_user_user_permissions-user_id" FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT "fk-django_admin_log-content_type_id" FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT "fk-django_admin_log-user_id" FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.oidc_provider_client
    ADD CONSTRAINT "fk-oidc_provider_client-owner_id" FOREIGN KEY (owner_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.oidc_provider_client_response_types
    ADD CONSTRAINT "fk-oidc_provider_client_response_types-client_id" FOREIGN KEY (client_id) REFERENCES public.oidc_provider_client(id);

ALTER TABLE ONLY public.oidc_provider_client_response_types
    ADD CONSTRAINT "fk-oidc_provider_client_response_types-responsetype_id" FOREIGN KEY (responsetype_id) REFERENCES public.oidc_provider_responsetype(id);

ALTER TABLE ONLY public.oidc_provider_code
    ADD CONSTRAINT "fk-oidc_provider_code-client_id" FOREIGN KEY (client_id) REFERENCES public.oidc_provider_client(id);

ALTER TABLE ONLY public.oidc_provider_code
    ADD CONSTRAINT "fk-oidc_provider_code-user_id" FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.oidc_provider_token
    ADD CONSTRAINT "fk-oidc_provider_token-client_id" FOREIGN KEY (client_id) REFERENCES public.oidc_provider_client(id);

ALTER TABLE ONLY public.oidc_provider_token
    ADD CONSTRAINT "fk-oidc_provider_token-user_id" FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.oidc_provider_userconsent
    ADD CONSTRAINT "fk-oidc_provider_userconsent-client_id" FOREIGN KEY (client_id) REFERENCES public.oidc_provider_client(id);

ALTER TABLE ONLY public.oidc_provider_userconsent
    ADD CONSTRAINT "fk-oidc_provider_userconsent-user_id" FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

