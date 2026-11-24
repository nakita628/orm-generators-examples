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

CREATE TABLE public.cas_server_federatediendityprovider (
    id bigint NOT NULL,
    suffix character varying(30) NOT NULL,
    server_url character varying(255) NOT NULL,
    cas_protocol_version character varying(30) NOT NULL,
    verbose_name character varying(255) NOT NULL,
    pos integer NOT NULL,
    display boolean NOT NULL
);

CREATE SEQUENCE public.cas_server_federatediendityprovider_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_federatediendityprovider_id_seq OWNED BY public.cas_server_federatediendityprovider.id;

CREATE TABLE public.cas_server_federateduser (
    id bigint NOT NULL,
    _attributs character varying,
    username character varying(124) NOT NULL,
    provider_id bigint NOT NULL,
    ticket character varying(255) NOT NULL,
    last_update timestamp with time zone NOT NULL
);

CREATE SEQUENCE public.cas_server_federateduser_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_federateduser_id_seq OWNED BY public.cas_server_federateduser.id;

CREATE TABLE public.cas_server_federateslo (
    id bigint NOT NULL,
    username character varying(30) NOT NULL,
    session_key character varying(40),
    ticket character varying(255) NOT NULL
);

CREATE SEQUENCE public.cas_server_federateslo_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_federateslo_id_seq OWNED BY public.cas_server_federateslo.id;

CREATE TABLE public.cas_server_filterattributvalue (
    id bigint NOT NULL,
    attribut character varying(255) NOT NULL,
    pattern character varying(255) NOT NULL,
    service_pattern_id bigint NOT NULL
);

CREATE SEQUENCE public.cas_server_filterattributvalue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_filterattributvalue_id_seq OWNED BY public.cas_server_filterattributvalue.id;

CREATE TABLE public.cas_server_newversionwarning (
    id bigint NOT NULL,
    version character varying(255) NOT NULL
);

CREATE SEQUENCE public.cas_server_newversionwarning_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_newversionwarning_id_seq OWNED BY public.cas_server_newversionwarning.id;

CREATE TABLE public.cas_server_proxy (
    id bigint NOT NULL,
    url character varying(255) NOT NULL,
    proxy_ticket_id bigint NOT NULL
);

CREATE SEQUENCE public.cas_server_proxy_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_proxy_id_seq OWNED BY public.cas_server_proxy.id;

CREATE TABLE public.cas_server_proxygrantingticket (
    id bigint NOT NULL,
    _attributs character varying,
    user_id bigint NOT NULL,
    validate boolean NOT NULL,
    service character varying NOT NULL,
    service_pattern_id bigint NOT NULL,
    creation timestamp with time zone NOT NULL,
    renew boolean NOT NULL,
    single_log_out boolean NOT NULL,
    value character varying(255) NOT NULL
);

CREATE SEQUENCE public.cas_server_proxygrantingticket_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_proxygrantingticket_id_seq OWNED BY public.cas_server_proxygrantingticket.id;

CREATE TABLE public.cas_server_proxyticket (
    id bigint NOT NULL,
    _attributs character varying,
    user_id bigint NOT NULL,
    validate boolean NOT NULL,
    service character varying NOT NULL,
    service_pattern_id bigint NOT NULL,
    creation timestamp with time zone NOT NULL,
    renew boolean NOT NULL,
    single_log_out boolean NOT NULL,
    value character varying(255) NOT NULL
);

CREATE SEQUENCE public.cas_server_proxyticket_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_proxyticket_id_seq OWNED BY public.cas_server_proxyticket.id;

CREATE TABLE public.cas_server_replaceattributname (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    replace character varying(255) NOT NULL,
    service_pattern_id bigint NOT NULL
);

CREATE SEQUENCE public.cas_server_replaceattributname_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_replaceattributname_id_seq OWNED BY public.cas_server_replaceattributname.id;

CREATE TABLE public.cas_server_replaceattributvalue (
    id bigint NOT NULL,
    attribut character varying(255) NOT NULL,
    pattern character varying(255) NOT NULL,
    replace character varying(255) NOT NULL,
    service_pattern_id bigint NOT NULL
);

CREATE SEQUENCE public.cas_server_replaceattributvalue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_replaceattributvalue_id_seq OWNED BY public.cas_server_replaceattributvalue.id;

CREATE TABLE public.cas_server_servicepattern (
    id bigint NOT NULL,
    pos integer NOT NULL,
    name character varying(255),
    pattern character varying(255) NOT NULL,
    user_field character varying(255) NOT NULL,
    restrict_users boolean NOT NULL,
    proxy boolean NOT NULL,
    proxy_callback boolean NOT NULL,
    single_log_out boolean NOT NULL,
    single_log_out_callback character varying(255) NOT NULL
);

CREATE SEQUENCE public.cas_server_servicepattern_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_servicepattern_id_seq OWNED BY public.cas_server_servicepattern.id;

CREATE TABLE public.cas_server_serviceticket (
    id bigint NOT NULL,
    _attributs character varying,
    user_id bigint NOT NULL,
    validate boolean NOT NULL,
    service character varying NOT NULL,
    service_pattern_id bigint NOT NULL,
    creation timestamp with time zone NOT NULL,
    renew boolean NOT NULL,
    single_log_out boolean NOT NULL,
    value character varying(255) NOT NULL
);

CREATE SEQUENCE public.cas_server_serviceticket_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_serviceticket_id_seq OWNED BY public.cas_server_serviceticket.id;

CREATE TABLE public.cas_server_user (
    id bigint NOT NULL,
    session_key character varying(40),
    username character varying(250) NOT NULL,
    date timestamp with time zone NOT NULL,
    last_login timestamp with time zone NOT NULL
);

CREATE SEQUENCE public.cas_server_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_user_id_seq OWNED BY public.cas_server_user.id;

CREATE TABLE public.cas_server_userattributes (
    id bigint NOT NULL,
    _attributs character varying,
    username character varying(155) NOT NULL
);

CREATE SEQUENCE public.cas_server_userattributes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_userattributes_id_seq OWNED BY public.cas_server_userattributes.id;

CREATE TABLE public.cas_server_username (
    id bigint NOT NULL,
    value character varying(255) NOT NULL,
    service_pattern_id bigint NOT NULL
);

CREATE SEQUENCE public.cas_server_username_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cas_server_username_id_seq OWNED BY public.cas_server_username.id;

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

ALTER TABLE ONLY public.auth_group ALTER COLUMN id SET DEFAULT nextval('public.auth_group_id_seq'::regclass);

ALTER TABLE ONLY public.auth_group_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_group_permissions_id_seq'::regclass);

ALTER TABLE ONLY public.auth_permission ALTER COLUMN id SET DEFAULT nextval('public.auth_permission_id_seq'::regclass);

ALTER TABLE ONLY public.auth_user ALTER COLUMN id SET DEFAULT nextval('public.auth_user_id_seq'::regclass);

ALTER TABLE ONLY public.auth_user_groups ALTER COLUMN id SET DEFAULT nextval('public.auth_user_groups_id_seq'::regclass);

ALTER TABLE ONLY public.auth_user_user_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_user_user_permissions_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_federatediendityprovider ALTER COLUMN id SET DEFAULT nextval('public.cas_server_federatediendityprovider_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_federateduser ALTER COLUMN id SET DEFAULT nextval('public.cas_server_federateduser_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_federateslo ALTER COLUMN id SET DEFAULT nextval('public.cas_server_federateslo_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_filterattributvalue ALTER COLUMN id SET DEFAULT nextval('public.cas_server_filterattributvalue_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_newversionwarning ALTER COLUMN id SET DEFAULT nextval('public.cas_server_newversionwarning_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_proxy ALTER COLUMN id SET DEFAULT nextval('public.cas_server_proxy_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_proxygrantingticket ALTER COLUMN id SET DEFAULT nextval('public.cas_server_proxygrantingticket_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_proxyticket ALTER COLUMN id SET DEFAULT nextval('public.cas_server_proxyticket_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_replaceattributname ALTER COLUMN id SET DEFAULT nextval('public.cas_server_replaceattributname_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_replaceattributvalue ALTER COLUMN id SET DEFAULT nextval('public.cas_server_replaceattributvalue_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_servicepattern ALTER COLUMN id SET DEFAULT nextval('public.cas_server_servicepattern_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_serviceticket ALTER COLUMN id SET DEFAULT nextval('public.cas_server_serviceticket_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_user ALTER COLUMN id SET DEFAULT nextval('public.cas_server_user_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_userattributes ALTER COLUMN id SET DEFAULT nextval('public.cas_server_userattributes_id_seq'::regclass);

ALTER TABLE ONLY public.cas_server_username ALTER COLUMN id SET DEFAULT nextval('public.cas_server_username_id_seq'::regclass);

ALTER TABLE ONLY public.django_admin_log ALTER COLUMN id SET DEFAULT nextval('public.django_admin_log_id_seq'::regclass);

ALTER TABLE ONLY public.django_content_type ALTER COLUMN id SET DEFAULT nextval('public.django_content_type_id_seq'::regclass);

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

ALTER TABLE ONLY public.cas_server_federatediendityprovider
    ADD CONSTRAINT cas_server_federatediendityprovider_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_federatediendityprovider
    ADD CONSTRAINT cas_server_federatediendityprovider_suffix_key UNIQUE (suffix);

ALTER TABLE ONLY public.cas_server_federateduser
    ADD CONSTRAINT cas_server_federateduser_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_federateslo
    ADD CONSTRAINT cas_server_federateslo_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_filterattributvalue
    ADD CONSTRAINT cas_server_filterattributvalue_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_newversionwarning
    ADD CONSTRAINT cas_server_newversionwarning_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_proxy
    ADD CONSTRAINT cas_server_proxy_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_proxygrantingticket
    ADD CONSTRAINT cas_server_proxygrantingticket_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_proxygrantingticket
    ADD CONSTRAINT cas_server_proxygrantingticket_value_key UNIQUE (value);

ALTER TABLE ONLY public.cas_server_proxyticket
    ADD CONSTRAINT cas_server_proxyticket_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_proxyticket
    ADD CONSTRAINT cas_server_proxyticket_value_key UNIQUE (value);

ALTER TABLE ONLY public.cas_server_replaceattributname
    ADD CONSTRAINT cas_server_replaceattributname_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_replaceattributvalue
    ADD CONSTRAINT cas_server_replaceattributvalue_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_servicepattern
    ADD CONSTRAINT cas_server_servicepattern_name_key UNIQUE (name);

ALTER TABLE ONLY public.cas_server_servicepattern
    ADD CONSTRAINT cas_server_servicepattern_pattern_key UNIQUE (pattern);

ALTER TABLE ONLY public.cas_server_servicepattern
    ADD CONSTRAINT cas_server_servicepattern_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_serviceticket
    ADD CONSTRAINT cas_server_serviceticket_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_serviceticket
    ADD CONSTRAINT cas_server_serviceticket_value_key UNIQUE (value);

ALTER TABLE ONLY public.cas_server_user
    ADD CONSTRAINT cas_server_user_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_userattributes
    ADD CONSTRAINT cas_server_userattributes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.cas_server_userattributes
    ADD CONSTRAINT cas_server_userattributes_username_key UNIQUE (username);

ALTER TABLE ONLY public.cas_server_username
    ADD CONSTRAINT cas_server_username_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_session
    ADD CONSTRAINT django_session_pkey PRIMARY KEY (session_key);

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

ALTER TABLE ONLY public.cas_server_federateduser
    ADD CONSTRAINT "fk-cas_server_federateduser-provider_id" FOREIGN KEY (provider_id) REFERENCES public.cas_server_federatediendityprovider(id);

ALTER TABLE ONLY public.cas_server_filterattributvalue
    ADD CONSTRAINT "fk-cas_server_filterattributvalue-service_pattern_id" FOREIGN KEY (service_pattern_id) REFERENCES public.cas_server_servicepattern(id);

ALTER TABLE ONLY public.cas_server_proxy
    ADD CONSTRAINT "fk-cas_server_proxy-proxy_ticket_id" FOREIGN KEY (proxy_ticket_id) REFERENCES public.cas_server_proxyticket(id);

ALTER TABLE ONLY public.cas_server_proxygrantingticket
    ADD CONSTRAINT "fk-cas_server_proxygrantingticket-service_pattern_id" FOREIGN KEY (service_pattern_id) REFERENCES public.cas_server_servicepattern(id);

ALTER TABLE ONLY public.cas_server_proxygrantingticket
    ADD CONSTRAINT "fk-cas_server_proxygrantingticket-user_id" FOREIGN KEY (user_id) REFERENCES public.cas_server_user(id);

ALTER TABLE ONLY public.cas_server_proxyticket
    ADD CONSTRAINT "fk-cas_server_proxyticket-service_pattern_id" FOREIGN KEY (service_pattern_id) REFERENCES public.cas_server_servicepattern(id);

ALTER TABLE ONLY public.cas_server_proxyticket
    ADD CONSTRAINT "fk-cas_server_proxyticket-user_id" FOREIGN KEY (user_id) REFERENCES public.cas_server_user(id);

ALTER TABLE ONLY public.cas_server_replaceattributname
    ADD CONSTRAINT "fk-cas_server_replaceattributname-service_pattern_id" FOREIGN KEY (service_pattern_id) REFERENCES public.cas_server_servicepattern(id);

ALTER TABLE ONLY public.cas_server_replaceattributvalue
    ADD CONSTRAINT "fk-cas_server_replaceattributvalue-service_pattern_id" FOREIGN KEY (service_pattern_id) REFERENCES public.cas_server_servicepattern(id);

ALTER TABLE ONLY public.cas_server_serviceticket
    ADD CONSTRAINT "fk-cas_server_serviceticket-service_pattern_id" FOREIGN KEY (service_pattern_id) REFERENCES public.cas_server_servicepattern(id);

ALTER TABLE ONLY public.cas_server_serviceticket
    ADD CONSTRAINT "fk-cas_server_serviceticket-user_id" FOREIGN KEY (user_id) REFERENCES public.cas_server_user(id);

ALTER TABLE ONLY public.cas_server_username
    ADD CONSTRAINT "fk-cas_server_username-service_pattern_id" FOREIGN KEY (service_pattern_id) REFERENCES public.cas_server_servicepattern(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT "fk-django_admin_log-content_type_id" FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT "fk-django_admin_log-user_id" FOREIGN KEY (user_id) REFERENCES public.auth_user(id);

