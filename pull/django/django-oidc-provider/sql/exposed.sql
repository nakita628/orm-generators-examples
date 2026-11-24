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
    last_login timestamp(6) with time zone,
    is_superuser boolean NOT NULL,
    username character varying(150) NOT NULL,
    first_name character varying(150) NOT NULL,
    last_name character varying(150) NOT NULL,
    email character varying(254) NOT NULL,
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
    date_joined timestamp(6) with time zone NOT NULL
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
    action_time timestamp(6) with time zone NOT NULL,
    user_id integer NOT NULL,
    content_type_id integer,
    object_id text,
    object_repr character varying(200) NOT NULL,
    action_flag smallint NOT NULL,
    change_message text NOT NULL
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
    session_data text NOT NULL,
    expire_date timestamp(6) with time zone NOT NULL
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
    _redirect_uris text NOT NULL,
    _post_logout_redirect_uris text NOT NULL,
    _scope text NOT NULL
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
    expires_at timestamp(6) with time zone NOT NULL,
    _scope text NOT NULL,
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
    key text NOT NULL
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
    expires_at timestamp(6) with time zone NOT NULL,
    _scope text NOT NULL,
    user_id integer,
    access_token character varying(255) NOT NULL,
    refresh_token character varying(255) NOT NULL,
    _id_token text NOT NULL
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
    expires_at timestamp(6) with time zone NOT NULL,
    _scope text NOT NULL,
    user_id integer NOT NULL,
    date_given timestamp(6) with time zone NOT NULL
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
    ADD CONSTRAINT auth_group_permissions_group_id_permission_id_0cd325b0_uniq UNIQUE (group_id, permission_id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_codename_01ab375a_uniq UNIQUE (content_type_id, codename);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_user_id_group_id_94350c0c_uniq UNIQUE (user_id, group_id);

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_user_id_permission_id_14a6b632_uniq UNIQUE (user_id, permission_id);

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_username_key UNIQUE (username);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_app_label_model_76bd3d3b_uniq UNIQUE (app_label, model);

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_session
    ADD CONSTRAINT django_session_pkey PRIMARY KEY (session_key);

ALTER TABLE ONLY public.oidc_provider_client
    ADD CONSTRAINT oidc_provider_client_client_id_key UNIQUE (client_id);

ALTER TABLE ONLY public.oidc_provider_client
    ADD CONSTRAINT oidc_provider_client_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.oidc_provider_client_response_types
    ADD CONSTRAINT oidc_provider_client_res_client_id_responsetype_i_0d116294_uniq UNIQUE (client_id, responsetype_id);

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

ALTER TABLE ONLY public.oidc_provider_userconsent
    ADD CONSTRAINT oidc_provider_userconsent_user_id_client_id_1c2032f6_uniq UNIQUE (user_id, client_id);

CREATE INDEX auth_group_name_a6ea08ec_like ON public.auth_group USING btree (name);

CREATE INDEX auth_group_permissions_group_id_b120cbf9 ON public.auth_group_permissions USING btree (group_id);

CREATE INDEX auth_group_permissions_permission_id_84c5c92e ON public.auth_group_permissions USING btree (permission_id);

CREATE INDEX auth_permission_content_type_id_2f476e4b ON public.auth_permission USING btree (content_type_id);

CREATE INDEX auth_user_groups_group_id_97559544 ON public.auth_user_groups USING btree (group_id);

CREATE INDEX auth_user_groups_user_id_6a12ed8b ON public.auth_user_groups USING btree (user_id);

CREATE INDEX auth_user_user_permissions_permission_id_1fbb5f2c ON public.auth_user_user_permissions USING btree (permission_id);

CREATE INDEX auth_user_user_permissions_user_id_a95ead1b ON public.auth_user_user_permissions USING btree (user_id);

CREATE INDEX auth_user_username_6821ab7c_like ON public.auth_user USING btree (username);

CREATE INDEX django_admin_log_content_type_id_c4bce8eb ON public.django_admin_log USING btree (content_type_id);

CREATE INDEX django_admin_log_user_id_c564eba6 ON public.django_admin_log USING btree (user_id);

CREATE INDEX django_session_expire_date_a5c62663 ON public.django_session USING btree (expire_date);

CREATE INDEX django_session_session_key_c0390e0f_like ON public.django_session USING btree (session_key);

CREATE INDEX oidc_provider_client_client_id_47991f0f_like ON public.oidc_provider_client USING btree (client_id);

CREATE INDEX oidc_provider_client_owner_id_735e6178 ON public.oidc_provider_client USING btree (owner_id);

CREATE INDEX oidc_provider_client_response_types_client_id_576aac8c ON public.oidc_provider_client_response_types USING btree (client_id);

CREATE INDEX oidc_provider_client_response_types_responsetype_id_c90751ad ON public.oidc_provider_client_response_types USING btree (responsetype_id);

CREATE INDEX oidc_provider_code_client_id_2449f20c ON public.oidc_provider_code USING btree (client_id);

CREATE INDEX oidc_provider_code_code_74b41245_like ON public.oidc_provider_code USING btree (code);

CREATE INDEX oidc_provider_code_user_id_14775b3f ON public.oidc_provider_code USING btree (user_id);

CREATE INDEX oidc_provider_responsetype_value_3a4af356_like ON public.oidc_provider_responsetype USING btree (value);

CREATE INDEX oidc_provider_token_access_token_6cfc5a80_like ON public.oidc_provider_token USING btree (access_token);

CREATE INDEX oidc_provider_token_client_id_86a72d16 ON public.oidc_provider_token USING btree (client_id);

CREATE INDEX oidc_provider_token_refresh_token_41436ea1_like ON public.oidc_provider_token USING btree (refresh_token);

CREATE INDEX oidc_provider_token_user_id_c61045ea ON public.oidc_provider_token USING btree (user_id);

CREATE INDEX oidc_provider_userconsent_client_id_7837dc61 ON public.oidc_provider_userconsent USING btree (client_id);

CREATE INDEX oidc_provider_userconsent_user_id_ebaddfd3 ON public.oidc_provider_userconsent USING btree (user_id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissio_permission_id_84c5c92e_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_b120cbf9_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_2f476e4b_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_group_id_97559544_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id);

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_user_id_6a12ed8b_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id);

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_content_type_id_c4bce8eb_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_user_id_c564eba6_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.oidc_provider_client_response_types
    ADD CONSTRAINT oidc_provider_client_client_id_576aac8c_fk_oidc_prov FOREIGN KEY (client_id) REFERENCES public.oidc_provider_client(id);

ALTER TABLE ONLY public.oidc_provider_client
    ADD CONSTRAINT oidc_provider_client_owner_id_735e6178_fk_auth_user_id FOREIGN KEY (owner_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.oidc_provider_client_response_types
    ADD CONSTRAINT oidc_provider_client_responsetype_id_c90751ad_fk_oidc_prov FOREIGN KEY (responsetype_id) REFERENCES public.oidc_provider_responsetype(id);

ALTER TABLE ONLY public.oidc_provider_code
    ADD CONSTRAINT oidc_provider_code_client_id_2449f20c_fk_oidc_prov FOREIGN KEY (client_id) REFERENCES public.oidc_provider_client(id);

ALTER TABLE ONLY public.oidc_provider_code
    ADD CONSTRAINT oidc_provider_code_user_id_14775b3f_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.oidc_provider_token
    ADD CONSTRAINT oidc_provider_token_client_id_86a72d16_fk_oidc_prov FOREIGN KEY (client_id) REFERENCES public.oidc_provider_client(id);

ALTER TABLE ONLY public.oidc_provider_token
    ADD CONSTRAINT oidc_provider_token_user_id_c61045ea_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

ALTER TABLE ONLY public.oidc_provider_userconsent
    ADD CONSTRAINT oidc_provider_userco_client_id_7837dc61_fk_oidc_prov FOREIGN KEY (client_id) REFERENCES public.oidc_provider_client(id);

ALTER TABLE ONLY public.oidc_provider_userconsent
    ADD CONSTRAINT oidc_provider_userconsent_user_id_ebaddfd3_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

