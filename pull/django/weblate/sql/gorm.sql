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

CREATE TABLE public.accounts_auditlog (
    id integer NOT NULL,
    user_id integer,
    activity character varying(20) NOT NULL,
    params jsonb NOT NULL,
    address text,
    user_agent character varying(200) NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.accounts_auditlog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.accounts_auditlog_id_seq OWNED BY public.accounts_auditlog.id;

CREATE TABLE public.accounts_profile (
    id integer NOT NULL,
    user_id integer NOT NULL,
    language character varying(10) NOT NULL,
    suggested integer NOT NULL,
    translated integer NOT NULL,
    uploaded integer NOT NULL,
    commented integer NOT NULL,
    theme character varying(10) NOT NULL,
    hide_completed boolean NOT NULL,
    secondary_in_zen boolean NOT NULL,
    hide_source_secondary boolean NOT NULL,
    wide_tables boolean NOT NULL,
    listing_columns jsonb NOT NULL,
    editor_link character varying(200) NOT NULL,
    translate_mode integer NOT NULL,
    zen_mode integer NOT NULL,
    special_chars character varying(30) NOT NULL,
    nearby_strings smallint NOT NULL,
    auto_watch boolean NOT NULL,
    contribute_personal_tm boolean NOT NULL,
    dashboard_view integer NOT NULL,
    dashboard_component_list_id integer,
    website character varying(200) NOT NULL,
    contact character varying(200) NOT NULL,
    liberapay character varying(50) NOT NULL,
    fediverse character varying(200) NOT NULL,
    codesite character varying(200) NOT NULL,
    github character varying(50) NOT NULL,
    twitter character varying(50) NOT NULL,
    linkedin character varying(50) NOT NULL,
    location character varying(100) NOT NULL,
    company character varying(100) NOT NULL,
    public_email character varying(190) NOT NULL,
    commit_email character varying(190) NOT NULL,
    commit_name integer NOT NULL,
    last_2fa character varying(15) NOT NULL
);

CREATE SEQUENCE public.accounts_profile_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.accounts_profile_id_seq OWNED BY public.accounts_profile.id;

CREATE TABLE public.accounts_profile_languages (
    id integer NOT NULL,
    profile_id integer NOT NULL,
    language_id integer NOT NULL
);

CREATE SEQUENCE public.accounts_profile_languages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.accounts_profile_languages_id_seq OWNED BY public.accounts_profile_languages.id;

CREATE TABLE public.accounts_profile_secondary_languages (
    id integer NOT NULL,
    profile_id integer NOT NULL,
    language_id integer NOT NULL
);

CREATE SEQUENCE public.accounts_profile_secondary_languages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.accounts_profile_secondary_languages_id_seq OWNED BY public.accounts_profile_secondary_languages.id;

CREATE TABLE public.accounts_profile_watched (
    id integer NOT NULL,
    profile_id integer NOT NULL,
    project_id integer NOT NULL
);

CREATE SEQUENCE public.accounts_profile_watched_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.accounts_profile_watched_id_seq OWNED BY public.accounts_profile_watched.id;

CREATE TABLE public.accounts_subscription (
    id integer NOT NULL,
    user_id integer NOT NULL,
    notification character varying(100) NOT NULL,
    scope integer NOT NULL,
    frequency integer NOT NULL,
    project_id integer,
    component_id integer,
    onetime boolean NOT NULL
);

CREATE SEQUENCE public.accounts_subscription_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.accounts_subscription_id_seq OWNED BY public.accounts_subscription.id;

CREATE TABLE public.accounts_teamauditlog (
    id integer NOT NULL,
    team_id integer,
    team_name character varying(150) NOT NULL,
    scope character varying(10) NOT NULL,
    scope_label character varying(200) NOT NULL,
    project_id integer,
    workspace_id character(36),
    actor_id integer,
    event character varying(20) NOT NULL,
    details jsonb NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.accounts_teamauditlog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.accounts_teamauditlog_id_seq OWNED BY public.accounts_teamauditlog.id;

CREATE TABLE public.accounts_verifiedemail (
    id integer NOT NULL,
    is_deliverable boolean NOT NULL,
    social_id bigint NOT NULL,
    email character varying(190) NOT NULL
);

CREATE SEQUENCE public.accounts_verifiedemail_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.accounts_verifiedemail_id_seq OWNED BY public.accounts_verifiedemail.id;

CREATE TABLE public.addons_addon (
    id integer NOT NULL,
    component_id integer,
    project_id integer,
    category_id integer,
    name character varying(100) NOT NULL,
    configuration jsonb NOT NULL,
    state jsonb NOT NULL,
    repo_scope boolean NOT NULL
);

CREATE SEQUENCE public.addons_addon_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.addons_addon_id_seq OWNED BY public.addons_addon.id;

CREATE TABLE public.addons_addonactivitylog (
    id integer NOT NULL,
    addon_id integer NOT NULL,
    component_id integer,
    event integer NOT NULL,
    created timestamp(6) with time zone NOT NULL,
    details jsonb NOT NULL,
    status integer NOT NULL
);

CREATE SEQUENCE public.addons_addonactivitylog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.addons_addonactivitylog_id_seq OWNED BY public.addons_addonactivitylog.id;

CREATE TABLE public.addons_event (
    id integer NOT NULL,
    addon_id integer NOT NULL,
    event integer NOT NULL
);

CREATE SEQUENCE public.addons_event_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.addons_event_id_seq OWNED BY public.addons_event.id;

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

CREATE TABLE public.authtoken_token (
    key character varying(40) NOT NULL,
    user_id integer NOT NULL,
    created timestamp(6) with time zone NOT NULL
);

CREATE TABLE public.checks_check (
    id integer NOT NULL,
    unit_id integer NOT NULL,
    name character varying(50) NOT NULL,
    dismissed boolean NOT NULL,
    metadata jsonb NOT NULL
);

CREATE SEQUENCE public.checks_check_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.checks_check_id_seq OWNED BY public.checks_check.id;

CREATE TABLE public.configuration_setting (
    id integer NOT NULL,
    category integer NOT NULL,
    name character varying(100) NOT NULL,
    value jsonb NOT NULL
);

CREATE SEQUENCE public.configuration_setting_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.configuration_setting_id_seq OWNED BY public.configuration_setting.id;

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

CREATE TABLE public.django_celery_beat_clockedschedule (
    id integer NOT NULL,
    clocked_time timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.django_celery_beat_clockedschedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_celery_beat_clockedschedule_id_seq OWNED BY public.django_celery_beat_clockedschedule.id;

CREATE TABLE public.django_celery_beat_crontabschedule (
    id integer NOT NULL,
    minute character varying(240) NOT NULL,
    hour character varying(96) NOT NULL,
    day_of_month character varying(124) NOT NULL,
    month_of_year character varying(64) NOT NULL,
    day_of_week character varying(64) NOT NULL,
    timezone character varying(63) NOT NULL
);

CREATE SEQUENCE public.django_celery_beat_crontabschedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_celery_beat_crontabschedule_id_seq OWNED BY public.django_celery_beat_crontabschedule.id;

CREATE TABLE public.django_celery_beat_intervalschedule (
    id integer NOT NULL,
    every integer NOT NULL,
    period character varying(24) NOT NULL
);

CREATE SEQUENCE public.django_celery_beat_intervalschedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_celery_beat_intervalschedule_id_seq OWNED BY public.django_celery_beat_intervalschedule.id;

CREATE TABLE public.django_celery_beat_periodictask (
    id integer NOT NULL,
    name character varying(200) NOT NULL,
    task character varying(200) NOT NULL,
    interval_id integer,
    crontab_id integer,
    solar_id integer,
    clocked_id integer,
    args text NOT NULL,
    kwargs text NOT NULL,
    queue character varying(200),
    exchange character varying(200),
    routing_key character varying(200),
    headers text NOT NULL,
    priority integer,
    expires timestamp(6) with time zone,
    expire_seconds integer,
    one_off boolean NOT NULL,
    start_time timestamp(6) with time zone,
    enabled boolean NOT NULL,
    last_run_at timestamp(6) with time zone,
    total_run_count integer NOT NULL,
    date_changed timestamp(6) with time zone NOT NULL,
    description text NOT NULL
);

CREATE SEQUENCE public.django_celery_beat_periodictask_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_celery_beat_periodictask_id_seq OWNED BY public.django_celery_beat_periodictask.id;

CREATE TABLE public.django_celery_beat_periodictasks (
    ident bigint NOT NULL,
    last_update timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.django_celery_beat_periodictasks_ident_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_celery_beat_periodictasks_ident_seq OWNED BY public.django_celery_beat_periodictasks.ident;

CREATE TABLE public.django_celery_beat_solarschedule (
    id integer NOT NULL,
    event character varying(24) NOT NULL,
    latitude numeric(9,6) NOT NULL,
    longitude numeric(9,6) NOT NULL
);

CREATE SEQUENCE public.django_celery_beat_solarschedule_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_celery_beat_solarschedule_id_seq OWNED BY public.django_celery_beat_solarschedule.id;

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

CREATE TABLE public.django_otp_webauthn_webauthnattestation (
    id bigint NOT NULL,
    fmt character varying(255) NOT NULL,
    data bytea NOT NULL,
    client_data_json bytea NOT NULL,
    credential_id bigint NOT NULL
);

CREATE SEQUENCE public.django_otp_webauthn_webauthnattestation_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_otp_webauthn_webauthnattestation_id_seq OWNED BY public.django_otp_webauthn_webauthnattestation.id;

CREATE TABLE public.django_otp_webauthn_webauthncredential (
    id bigint NOT NULL,
    user_id integer NOT NULL,
    name character varying(64) NOT NULL,
    confirmed boolean NOT NULL,
    created_at timestamp(6) with time zone,
    last_used_at timestamp(6) with time zone,
    credential_type character varying(32) NOT NULL,
    credential_id bytea NOT NULL,
    public_key bytea NOT NULL,
    transports jsonb NOT NULL,
    sign_count integer NOT NULL,
    backup_eligible boolean NOT NULL,
    backup_state boolean NOT NULL,
    aaguid character varying(36) NOT NULL,
    credential_id_sha256 character varying(64) NOT NULL,
    discoverable boolean
);

CREATE SEQUENCE public.django_otp_webauthn_webauthncredential_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_otp_webauthn_webauthncredential_id_seq OWNED BY public.django_otp_webauthn_webauthncredential.id;

CREATE TABLE public.django_otp_webauthn_webauthnuserhandle (
    handle_hex character varying(128) NOT NULL,
    user_id integer NOT NULL
);

CREATE SEQUENCE public.django_otp_webauthn_webauthnuserhandle_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.django_otp_webauthn_webauthnuserhandle_user_id_seq OWNED BY public.django_otp_webauthn_webauthnuserhandle.user_id;

CREATE TABLE public.django_session (
    session_key character varying(40) NOT NULL,
    session_data text NOT NULL,
    expire_date timestamp(6) with time zone NOT NULL
);

CREATE TABLE public.fonts_font (
    id integer NOT NULL,
    family character varying(100) NOT NULL,
    style character varying(100) NOT NULL,
    font character varying(100) NOT NULL,
    project_id integer NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    user_id integer
);

CREATE SEQUENCE public.fonts_font_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.fonts_font_id_seq OWNED BY public.fonts_font.id;

CREATE TABLE public.fonts_fontgroup (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    font_id integer NOT NULL,
    project_id integer NOT NULL
);

CREATE SEQUENCE public.fonts_fontgroup_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.fonts_fontgroup_id_seq OWNED BY public.fonts_fontgroup.id;

CREATE TABLE public.fonts_fontoverride (
    id integer NOT NULL,
    group_id integer NOT NULL,
    font_id integer NOT NULL,
    language_id integer NOT NULL
);

CREATE SEQUENCE public.fonts_fontoverride_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.fonts_fontoverride_id_seq OWNED BY public.fonts_fontoverride.id;

CREATE TABLE public.kotlin_sdk_kotlinsdkartifact (
    id integer NOT NULL,
    addon_id integer NOT NULL,
    digest character varying(64) NOT NULL,
    unreferenced timestamp(6) with time zone
);

CREATE TABLE public.kotlin_sdk_kotlinsdkartifact_builds (
    id integer NOT NULL,
    kotlinsdkartifact_id integer NOT NULL,
    kotlinsdkbuild_id integer NOT NULL
);

CREATE SEQUENCE public.kotlin_sdk_kotlinsdkartifact_builds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.kotlin_sdk_kotlinsdkartifact_builds_id_seq OWNED BY public.kotlin_sdk_kotlinsdkartifact_builds.id;

CREATE SEQUENCE public.kotlin_sdk_kotlinsdkartifact_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.kotlin_sdk_kotlinsdkartifact_id_seq OWNED BY public.kotlin_sdk_kotlinsdkartifact.id;

CREATE TABLE public.kotlin_sdk_kotlinsdkbuild (
    id integer NOT NULL,
    addon_id integer NOT NULL,
    package_name character varying(127) NOT NULL,
    version_code bigint NOT NULL,
    metadata jsonb NOT NULL,
    pending_manifest jsonb NOT NULL,
    digest character varying(64) NOT NULL,
    created timestamp(6) with time zone NOT NULL,
    retired timestamp(6) with time zone,
    status character varying(16) NOT NULL,
    error text NOT NULL,
    published timestamp(6) with time zone
);

CREATE SEQUENCE public.kotlin_sdk_kotlinsdkbuild_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.kotlin_sdk_kotlinsdkbuild_id_seq OWNED BY public.kotlin_sdk_kotlinsdkbuild.id;

CREATE TABLE public.kotlin_sdk_kotlinsdkcleanup (
    id integer NOT NULL,
    path text NOT NULL,
    build_id integer
);

CREATE SEQUENCE public.kotlin_sdk_kotlinsdkcleanup_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.kotlin_sdk_kotlinsdkcleanup_id_seq OWNED BY public.kotlin_sdk_kotlinsdkcleanup.id;

CREATE TABLE public.lang_language (
    id integer NOT NULL,
    code character varying(50) NOT NULL,
    name character varying(100) NOT NULL,
    direction character varying(3) NOT NULL,
    population bigint NOT NULL
);

CREATE SEQUENCE public.lang_language_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.lang_language_id_seq OWNED BY public.lang_language.id;

CREATE TABLE public.lang_plural (
    id integer NOT NULL,
    source smallint NOT NULL,
    number smallint NOT NULL,
    formula text NOT NULL,
    type integer NOT NULL,
    language_id integer NOT NULL
);

CREATE SEQUENCE public.lang_plural_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.lang_plural_id_seq OWNED BY public.lang_plural.id;

CREATE TABLE public.machinery_machineryerror (
    id integer NOT NULL,
    installation_id integer NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    error text NOT NULL
);

CREATE SEQUENCE public.machinery_machineryerror_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.machinery_machineryerror_id_seq OWNED BY public.machinery_machineryerror.id;

CREATE TABLE public.machinery_machineryexclusion (
    id integer NOT NULL,
    workspace_id character(36),
    project_id integer,
    installation_id integer,
    engine character varying(100) NOT NULL
);

CREATE SEQUENCE public.machinery_machineryexclusion_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.machinery_machineryexclusion_id_seq OWNED BY public.machinery_machineryexclusion.id;

CREATE TABLE public.machinery_machineryinstallation (
    id integer NOT NULL,
    engine character varying(100) NOT NULL,
    slug character varying(100) NOT NULL,
    configuration jsonb NOT NULL,
    enabled boolean NOT NULL,
    workspace_id character(36),
    project_id integer
);

CREATE SEQUENCE public.machinery_machineryinstallation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.machinery_machineryinstallation_id_seq OWNED BY public.machinery_machineryinstallation.id;

CREATE TABLE public.memory_memory (
    id integer NOT NULL,
    source_language_id integer NOT NULL,
    target_language_id integer NOT NULL,
    source text NOT NULL,
    target text NOT NULL,
    origin text NOT NULL,
    context text NOT NULL,
    legacy_user_id integer,
    legacy_project_id integer,
    legacy_from_file boolean NOT NULL,
    legacy_shared boolean NOT NULL,
    status integer NOT NULL
);

CREATE SEQUENCE public.memory_memory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.memory_memory_id_seq OWNED BY public.memory_memory.id;

CREATE TABLE public.memory_memoryscope (
    id integer NOT NULL,
    memory_id integer NOT NULL,
    scope smallint NOT NULL,
    project_id integer,
    workspace_id character(36),
    source_project_id integer,
    source_component_id integer,
    user_id integer
);

CREATE SEQUENCE public.memory_memoryscope_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.memory_memoryscope_id_seq OWNED BY public.memory_memoryscope.id;

CREATE TABLE public.memory_memoryscopemigrationstate (
    name character varying(50) NOT NULL,
    last_memory_id integer NOT NULL,
    completed boolean NOT NULL,
    updated timestamp(6) with time zone NOT NULL
);

CREATE TABLE public.metrics_metric (
    id bigint NOT NULL,
    date date NOT NULL,
    scope smallint NOT NULL,
    relation integer NOT NULL,
    secondary integer NOT NULL,
    changes integer NOT NULL,
    data jsonb
);

CREATE SEQUENCE public.metrics_metric_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.metrics_metric_id_seq OWNED BY public.metrics_metric.id;

CREATE TABLE public.otp_static_staticdevice (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying(64) NOT NULL,
    confirmed boolean NOT NULL,
    throttling_failure_timestamp timestamp(6) with time zone,
    throttling_failure_count integer NOT NULL,
    created_at timestamp(6) with time zone,
    last_used_at timestamp(6) with time zone
);

CREATE SEQUENCE public.otp_static_staticdevice_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.otp_static_staticdevice_id_seq OWNED BY public.otp_static_staticdevice.id;

CREATE TABLE public.otp_static_statictoken (
    id integer NOT NULL,
    device_id integer NOT NULL,
    token character varying(16) NOT NULL
);

CREATE SEQUENCE public.otp_static_statictoken_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.otp_static_statictoken_id_seq OWNED BY public.otp_static_statictoken.id;

CREATE TABLE public.otp_totp_totpdevice (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying(64) NOT NULL,
    confirmed boolean NOT NULL,
    throttling_failure_timestamp timestamp(6) with time zone,
    throttling_failure_count integer NOT NULL,
    created_at timestamp(6) with time zone,
    last_used_at timestamp(6) with time zone,
    key character varying(80) NOT NULL,
    step smallint NOT NULL,
    t0 bigint NOT NULL,
    digits smallint NOT NULL,
    tolerance smallint NOT NULL,
    drift smallint NOT NULL,
    last_t bigint NOT NULL
);

CREATE SEQUENCE public.otp_totp_totpdevice_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.otp_totp_totpdevice_id_seq OWNED BY public.otp_totp_totpdevice.id;

CREATE TABLE public.screenshots_screenshot (
    id integer NOT NULL,
    backup_author_claim_id integer,
    name character varying(200) NOT NULL,
    repository_filename character varying(200) NOT NULL,
    image character varying(100) NOT NULL,
    translation_id integer NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    user_id integer
);

CREATE SEQUENCE public.screenshots_screenshot_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.screenshots_screenshot_id_seq OWNED BY public.screenshots_screenshot.id;

CREATE TABLE public.screenshots_screenshot_units (
    id integer NOT NULL,
    screenshot_id integer NOT NULL,
    unit_id integer NOT NULL
);

CREATE SEQUENCE public.screenshots_screenshot_units_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.screenshots_screenshot_units_id_seq OWNED BY public.screenshots_screenshot_units.id;

CREATE TABLE public.screenshots_screenshotocr (
    screenshot_id integer NOT NULL,
    signature jsonb NOT NULL,
    generation character(36) NOT NULL,
    data jsonb,
    task_id character(36),
    requested_at timestamp(6) with time zone,
    error character varying(16) NOT NULL
);

CREATE SEQUENCE public.screenshots_screenshotocr_screenshot_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.screenshots_screenshotocr_screenshot_id_seq OWNED BY public.screenshots_screenshotocr.screenshot_id;

CREATE TABLE public.social_auth_association (
    id bigint NOT NULL,
    server_url character varying(255) NOT NULL,
    handle character varying(255) NOT NULL,
    secret character varying(255) NOT NULL,
    issued integer NOT NULL,
    lifetime integer NOT NULL,
    assoc_type character varying(64) NOT NULL
);

CREATE SEQUENCE public.social_auth_association_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.social_auth_association_id_seq OWNED BY public.social_auth_association.id;

CREATE TABLE public.social_auth_code (
    id bigint NOT NULL,
    email character varying(254) NOT NULL,
    code character varying(32) NOT NULL,
    verified boolean NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.social_auth_code_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.social_auth_code_id_seq OWNED BY public.social_auth_code.id;

CREATE TABLE public.social_auth_nonce (
    id bigint NOT NULL,
    server_url character varying(255) NOT NULL,
    "timestamp" integer NOT NULL,
    salt character varying(65) NOT NULL
);

CREATE SEQUENCE public.social_auth_nonce_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.social_auth_nonce_id_seq OWNED BY public.social_auth_nonce.id;

CREATE TABLE public.social_auth_partial (
    id bigint NOT NULL,
    token character varying(32) NOT NULL,
    next_step smallint NOT NULL,
    backend character varying(32) NOT NULL,
    data jsonb NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.social_auth_partial_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.social_auth_partial_id_seq OWNED BY public.social_auth_partial.id;

CREATE TABLE public.social_auth_usersocialauth (
    id bigint NOT NULL,
    user_id integer NOT NULL,
    provider character varying(32) NOT NULL,
    id_key character varying(255) NOT NULL,
    uid character varying(255) NOT NULL,
    extra_data jsonb NOT NULL,
    created timestamp(6) with time zone NOT NULL,
    modified timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.social_auth_usersocialauth_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.social_auth_usersocialauth_id_seq OWNED BY public.social_auth_usersocialauth.id;

CREATE TABLE public.trans_alert (
    id integer NOT NULL,
    component_id integer NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    updated timestamp(6) with time zone NOT NULL,
    name character varying(150) NOT NULL,
    dismissed_at timestamp(6) with time zone,
    dismissed_by_id integer,
    dismissal_reason character varying(500) NOT NULL,
    dismissal_fingerprint character varying(64) NOT NULL,
    severity smallint NOT NULL,
    details jsonb NOT NULL
);

CREATE SEQUENCE public.trans_alert_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_alert_id_seq OWNED BY public.trans_alert.id;

CREATE TABLE public.trans_announcement (
    id integer NOT NULL,
    message text NOT NULL,
    project_id integer,
    category_id integer,
    component_id integer,
    language_id integer,
    severity character varying(25) NOT NULL,
    expiry date,
    notify boolean NOT NULL
);

CREATE SEQUENCE public.trans_announcement_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_announcement_id_seq OWNED BY public.trans_announcement.id;

CREATE TABLE public.trans_autocomponentlist (
    id integer NOT NULL,
    project_match character varying(200) NOT NULL,
    component_match character varying(200) NOT NULL,
    componentlist_id integer NOT NULL
);

CREATE SEQUENCE public.trans_autocomponentlist_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_autocomponentlist_id_seq OWNED BY public.trans_autocomponentlist.id;

CREATE TABLE public.trans_backupauthorclaim (
    id integer NOT NULL,
    project_id integer NOT NULL,
    user_id integer NOT NULL,
    invitation_id character(36),
    accepted_at timestamp(6) with time zone
);

CREATE SEQUENCE public.trans_backupauthorclaim_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_backupauthorclaim_id_seq OWNED BY public.trans_backupauthorclaim.id;

CREATE TABLE public.trans_backupauthorvote (
    id integer NOT NULL,
    claim_id integer NOT NULL,
    suggestion_id integer NOT NULL,
    placeholder_id integer,
    value smallint NOT NULL
);

CREATE SEQUENCE public.trans_backupauthorvote_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_backupauthorvote_id_seq OWNED BY public.trans_backupauthorvote.id;

CREATE TABLE public.trans_category (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    slug character varying(100) NOT NULL,
    project_id integer NOT NULL,
    category_id integer,
    check_flags text NOT NULL,
    enforced_checks jsonb NOT NULL,
    inherit_enforced_checks boolean NOT NULL,
    license character varying(150) NOT NULL,
    inherit_license boolean NOT NULL,
    agreement text NOT NULL,
    inherit_agreement boolean NOT NULL,
    new_lang character varying(10) NOT NULL,
    inherit_new_lang boolean NOT NULL,
    language_code_style character varying(20) NOT NULL,
    inherit_language_code_style boolean NOT NULL,
    secondary_language_id integer,
    inherit_secondary_language boolean NOT NULL,
    commit_message text NOT NULL,
    inherit_commit_message boolean NOT NULL,
    add_message text NOT NULL,
    inherit_add_message boolean NOT NULL,
    delete_message text NOT NULL,
    inherit_delete_message boolean NOT NULL,
    merge_message text NOT NULL,
    inherit_merge_message boolean NOT NULL,
    addon_message text NOT NULL,
    inherit_addon_message boolean NOT NULL,
    pull_message text NOT NULL,
    inherit_pull_message boolean NOT NULL
);

CREATE SEQUENCE public.trans_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_category_id_seq OWNED BY public.trans_category.id;

CREATE TABLE public.trans_change (
    id integer NOT NULL,
    backup_author_claim_id integer,
    unit_id integer,
    language_id integer,
    project_id integer,
    workspace_id character(36),
    category_id integer,
    component_id integer,
    translation_id integer,
    comment_id integer,
    suggestion_id integer,
    announcement_id integer,
    screenshot_id integer,
    alert_id integer,
    addon_id integer,
    user_id integer,
    author_id integer,
    "timestamp" timestamp(6) with time zone NOT NULL,
    action integer NOT NULL,
    target text NOT NULL,
    old text NOT NULL,
    details jsonb NOT NULL
);

CREATE SEQUENCE public.trans_change_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_change_id_seq OWNED BY public.trans_change.id;

CREATE TABLE public.trans_comment (
    id integer NOT NULL,
    backup_author_claim_id integer,
    unit_id integer NOT NULL,
    comment text NOT NULL,
    user_id integer,
    "timestamp" timestamp(6) with time zone NOT NULL,
    resolved boolean NOT NULL,
    userdetails jsonb NOT NULL
);

CREATE SEQUENCE public.trans_comment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_comment_id_seq OWNED BY public.trans_comment.id;

CREATE TABLE public.trans_component (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    slug character varying(100) NOT NULL,
    project_id integer NOT NULL,
    category_id integer,
    vcs character varying(20) NOT NULL,
    vcs_params jsonb NOT NULL,
    repo character varying(300) NOT NULL,
    linked_component_id integer,
    push character varying(300) NOT NULL,
    repoweb character varying(200) NOT NULL,
    repoweb_translations character varying(200) NOT NULL,
    git_export character varying(220) NOT NULL,
    report_source_bugs character varying(190) NOT NULL,
    branch character varying(200) NOT NULL,
    push_branch character varying(200) NOT NULL,
    filemask character varying(400) NOT NULL,
    screenshot_filemask character varying(400) NOT NULL,
    template character varying(400) NOT NULL,
    edit_template boolean NOT NULL,
    intermediate character varying(400) NOT NULL,
    new_base character varying(400) NOT NULL,
    file_format character varying(50) NOT NULL,
    file_format_params jsonb NOT NULL,
    locked boolean NOT NULL,
    hide_glossary_matches boolean NOT NULL,
    allow_translation_propagation boolean NOT NULL,
    contribute_project_tm boolean NOT NULL,
    enable_suggestions boolean NOT NULL,
    suggestion_voting boolean NOT NULL,
    suggestion_autoaccept smallint NOT NULL,
    check_flags text NOT NULL,
    enforced_checks jsonb NOT NULL,
    inherit_enforced_checks boolean NOT NULL,
    license character varying(150) NOT NULL,
    inherit_license boolean NOT NULL,
    agreement text NOT NULL,
    inherit_agreement boolean NOT NULL,
    new_lang character varying(10) NOT NULL,
    inherit_new_lang boolean NOT NULL,
    language_code_style character varying(20) NOT NULL,
    inherit_language_code_style boolean NOT NULL,
    manage_units boolean NOT NULL,
    merge_style character varying(10) NOT NULL,
    commit_message text NOT NULL,
    inherit_commit_message boolean NOT NULL,
    add_message text NOT NULL,
    inherit_add_message boolean NOT NULL,
    delete_message text NOT NULL,
    inherit_delete_message boolean NOT NULL,
    merge_message text NOT NULL,
    inherit_merge_message boolean NOT NULL,
    addon_message text NOT NULL,
    inherit_addon_message boolean NOT NULL,
    pull_message text NOT NULL,
    inherit_pull_message boolean NOT NULL,
    push_on_commit boolean NOT NULL,
    push_on_update boolean NOT NULL,
    commit_pending_age smallint NOT NULL,
    auto_lock_error boolean NOT NULL,
    source_language_id integer NOT NULL,
    language_regex character varying(500) NOT NULL,
    variant_regex character varying(190) NOT NULL,
    priority integer NOT NULL,
    restricted boolean NOT NULL,
    is_glossary boolean NOT NULL,
    glossary_color character varying(30) NOT NULL,
    remote_revision character varying(200) NOT NULL,
    local_revision character varying(200) NOT NULL,
    processed_revision character varying(200) NOT NULL,
    pull_request_url character varying(2048) NOT NULL,
    key_filter character varying(500) NOT NULL,
    secondary_language_id integer,
    inherit_secondary_language boolean NOT NULL
);

CREATE SEQUENCE public.trans_component_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_component_id_seq OWNED BY public.trans_component.id;

CREATE TABLE public.trans_component_links (
    id integer NOT NULL,
    component_id integer NOT NULL,
    project_id integer NOT NULL,
    category_id integer
);

CREATE SEQUENCE public.trans_component_links_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_component_links_id_seq OWNED BY public.trans_component_links.id;

CREATE TABLE public.trans_componentlist (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    slug character varying(100) NOT NULL,
    show_dashboard boolean NOT NULL
);

CREATE TABLE public.trans_componentlist_components (
    id integer NOT NULL,
    componentlist_id integer NOT NULL,
    component_id integer NOT NULL
);

CREATE SEQUENCE public.trans_componentlist_components_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_componentlist_components_id_seq OWNED BY public.trans_componentlist_components.id;

CREATE SEQUENCE public.trans_componentlist_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_componentlist_id_seq OWNED BY public.trans_componentlist.id;

CREATE TABLE public.trans_contributoragreement (
    id integer NOT NULL,
    user_id integer NOT NULL,
    component_id integer,
    project_id integer,
    category_id integer,
    workspace_id character(36),
    "timestamp" timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.trans_contributoragreement_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_contributoragreement_id_seq OWNED BY public.trans_contributoragreement.id;

CREATE TABLE public.trans_label (
    id integer NOT NULL,
    project_id integer NOT NULL,
    name character varying(190) NOT NULL,
    color character varying(30) NOT NULL,
    description character varying(250) NOT NULL
);

CREATE SEQUENCE public.trans_label_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_label_id_seq OWNED BY public.trans_label.id;

CREATE TABLE public.trans_pendingunitchange (
    id integer NOT NULL,
    backup_author_claim_id integer,
    unit_id integer NOT NULL,
    author_id integer NOT NULL,
    target text NOT NULL,
    explanation text NOT NULL,
    source_unit_explanation text NOT NULL,
    state integer NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    add_unit boolean NOT NULL,
    automatically_translated boolean NOT NULL,
    metadata jsonb NOT NULL
);

CREATE SEQUENCE public.trans_pendingunitchange_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_pendingunitchange_id_seq OWNED BY public.trans_pendingunitchange.id;

CREATE TABLE public.trans_project (
    id integer NOT NULL,
    name character varying(60) NOT NULL,
    slug character varying(60) NOT NULL,
    web character varying(200) NOT NULL,
    instructions text NOT NULL,
    workspace_id character(36),
    use_shared_tm boolean NOT NULL,
    contribute_shared_tm boolean NOT NULL,
    use_workspace_tm boolean NOT NULL,
    contribute_workspace_tm boolean NOT NULL,
    autoclean_tm boolean NOT NULL,
    access_control integer NOT NULL,
    public_sharing boolean NOT NULL,
    enforced_2fa boolean NOT NULL,
    translation_review boolean NOT NULL,
    source_review boolean NOT NULL,
    commit_policy integer NOT NULL,
    enable_hooks boolean NOT NULL,
    language_aliases text NOT NULL,
    secondary_language_id integer,
    check_flags text NOT NULL,
    enforced_checks jsonb NOT NULL,
    inherit_enforced_checks boolean NOT NULL,
    license character varying(150) NOT NULL,
    inherit_license boolean NOT NULL,
    agreement text NOT NULL,
    inherit_agreement boolean NOT NULL,
    new_lang character varying(10) NOT NULL,
    inherit_new_lang boolean NOT NULL,
    language_code_style character varying(20) NOT NULL,
    inherit_language_code_style boolean NOT NULL,
    inherit_secondary_language boolean NOT NULL,
    commit_message text NOT NULL,
    inherit_commit_message boolean NOT NULL,
    add_message text NOT NULL,
    inherit_add_message boolean NOT NULL,
    delete_message text NOT NULL,
    inherit_delete_message boolean NOT NULL,
    merge_message text NOT NULL,
    inherit_merge_message boolean NOT NULL,
    addon_message text NOT NULL,
    inherit_addon_message boolean NOT NULL,
    pull_message text NOT NULL,
    inherit_pull_message boolean NOT NULL
);

CREATE SEQUENCE public.trans_project_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_project_id_seq OWNED BY public.trans_project.id;

CREATE TABLE public.trans_report (
    id integer NOT NULL,
    creator_id integer NOT NULL,
    created timestamp(6) with time zone NOT NULL,
    kind character varying(32) NOT NULL,
    version smallint NOT NULL,
    parameters jsonb NOT NULL,
    data jsonb NOT NULL,
    workspace_id character(36),
    project_id integer,
    category_id integer,
    component_id integer
);

CREATE SEQUENCE public.trans_report_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_report_id_seq OWNED BY public.trans_report.id;

CREATE TABLE public.trans_savedsearch (
    id integer NOT NULL,
    project_id integer,
    user_id integer,
    name character varying(190) NOT NULL,
    query text NOT NULL,
    description character varying(250) NOT NULL
);

CREATE SEQUENCE public.trans_savedsearch_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_savedsearch_id_seq OWNED BY public.trans_savedsearch.id;

CREATE TABLE public.trans_suggestion (
    id integer NOT NULL,
    backup_author_claim_id integer,
    unit_id integer NOT NULL,
    target text NOT NULL,
    user_id integer,
    userdetails jsonb NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.trans_suggestion_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_suggestion_id_seq OWNED BY public.trans_suggestion.id;

CREATE TABLE public.trans_translation (
    id integer NOT NULL,
    component_id integer NOT NULL,
    language_id integer NOT NULL,
    plural_id integer NOT NULL,
    revision character varying(200) NOT NULL,
    filename character varying(400) NOT NULL,
    language_code character varying(50) NOT NULL,
    check_flags text NOT NULL
);

CREATE SEQUENCE public.trans_translation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_translation_id_seq OWNED BY public.trans_translation.id;

CREATE TABLE public.trans_unit (
    id integer NOT NULL,
    translation_id integer NOT NULL,
    id_hash bigint NOT NULL,
    location text NOT NULL,
    context text NOT NULL,
    note text NOT NULL,
    flags text NOT NULL,
    source text NOT NULL,
    previous_source text NOT NULL,
    target text NOT NULL,
    state integer NOT NULL,
    original_state integer NOT NULL,
    details jsonb NOT NULL,
    "position" integer NOT NULL,
    num_words integer NOT NULL,
    priority integer NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    last_updated timestamp(6) with time zone NOT NULL,
    extra_flags text NOT NULL,
    explanation text NOT NULL,
    variant_id integer,
    automatically_translated boolean NOT NULL,
    source_unit_id integer,
    translation_parent_id integer
);

CREATE SEQUENCE public.trans_unit_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_unit_id_seq OWNED BY public.trans_unit.id;

CREATE TABLE public.trans_unit_labels (
    id integer NOT NULL,
    unit_id integer NOT NULL,
    label_id integer NOT NULL
);

CREATE SEQUENCE public.trans_unit_labels_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_unit_labels_id_seq OWNED BY public.trans_unit_labels.id;

CREATE TABLE public.trans_variant (
    id integer NOT NULL,
    component_id integer NOT NULL,
    variant_regex character varying(190) NOT NULL,
    key text NOT NULL
);

CREATE TABLE public.trans_variant_defining_units (
    id integer NOT NULL,
    variant_id integer NOT NULL,
    unit_id integer NOT NULL
);

CREATE SEQUENCE public.trans_variant_defining_units_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_variant_defining_units_id_seq OWNED BY public.trans_variant_defining_units.id;

CREATE SEQUENCE public.trans_variant_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_variant_id_seq OWNED BY public.trans_variant.id;

CREATE TABLE public.trans_vote (
    id integer NOT NULL,
    suggestion_id integer NOT NULL,
    user_id integer NOT NULL,
    value smallint NOT NULL
);

CREATE SEQUENCE public.trans_vote_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_vote_id_seq OWNED BY public.trans_vote.id;

CREATE TABLE public.trans_workflowsetting (
    id integer NOT NULL,
    project_id integer,
    language_id integer NOT NULL,
    source_language_id integer,
    translation_review boolean NOT NULL,
    enable_suggestions boolean NOT NULL,
    restrict_direct_editing boolean NOT NULL,
    suggestion_voting boolean NOT NULL,
    suggestion_autoaccept smallint NOT NULL
);

CREATE SEQUENCE public.trans_workflowsetting_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trans_workflowsetting_id_seq OWNED BY public.trans_workflowsetting.id;

CREATE TABLE public.vcs_githubappcredentials (
    id integer NOT NULL,
    hostname character varying(255) NOT NULL,
    app_id character varying(50) NOT NULL,
    app_slug character varying(255) NOT NULL,
    private_key text NOT NULL,
    client_id character varying(255) NOT NULL,
    client_secret character varying(255) NOT NULL,
    webhook_secret character varying(255) NOT NULL,
    webhook_token character(36) NOT NULL,
    html_url character varying(200) NOT NULL,
    created timestamp(6) with time zone NOT NULL,
    updated timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.vcs_githubappcredentials_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.vcs_githubappcredentials_id_seq OWNED BY public.vcs_githubappcredentials.id;

CREATE TABLE public.vcs_installation (
    id integer NOT NULL,
    provider character varying(20) NOT NULL,
    installation_id character varying(50) NOT NULL,
    target_type character varying(20) NOT NULL,
    target_login character varying(255) NOT NULL,
    workspace_id character(36) NOT NULL,
    hostname character varying(255) NOT NULL,
    enabled boolean NOT NULL,
    created timestamp(6) with time zone NOT NULL,
    repositories jsonb NOT NULL,
    repositories_updated timestamp(6) with time zone
);

CREATE SEQUENCE public.vcs_installation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.vcs_installation_id_seq OWNED BY public.vcs_installation.id;

CREATE TABLE public.vcs_pendinginstallation (
    id integer NOT NULL,
    provider character varying(20) NOT NULL,
    hostname character varying(255) NOT NULL,
    installation_id character varying(50) NOT NULL,
    payload jsonb NOT NULL,
    created timestamp(6) with time zone NOT NULL,
    updated timestamp(6) with time zone NOT NULL
);

CREATE SEQUENCE public.vcs_pendinginstallation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.vcs_pendinginstallation_id_seq OWNED BY public.vcs_pendinginstallation.id;

CREATE TABLE public.weblate_auth_autogroup (
    id integer NOT NULL,
    match character varying(200) NOT NULL,
    group_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_autogroup_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_autogroup_id_seq OWNED BY public.weblate_auth_autogroup.id;

CREATE TABLE public.weblate_auth_group (
    id integer NOT NULL,
    name character varying(150) NOT NULL,
    defining_project_id integer,
    defining_workspace_id character(36),
    project_selection integer NOT NULL,
    language_selection integer NOT NULL,
    internal boolean NOT NULL,
    enforced_2fa boolean NOT NULL
);

CREATE TABLE public.weblate_auth_group_admins (
    id integer NOT NULL,
    group_id integer NOT NULL,
    user_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_group_admins_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_group_admins_id_seq OWNED BY public.weblate_auth_group_admins.id;

CREATE TABLE public.weblate_auth_group_componentlists (
    id integer NOT NULL,
    group_id integer NOT NULL,
    componentlist_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_group_componentlists_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_group_componentlists_id_seq OWNED BY public.weblate_auth_group_componentlists.id;

CREATE TABLE public.weblate_auth_group_components (
    id integer NOT NULL,
    group_id integer NOT NULL,
    component_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_group_components_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_group_components_id_seq OWNED BY public.weblate_auth_group_components.id;

CREATE SEQUENCE public.weblate_auth_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_group_id_seq OWNED BY public.weblate_auth_group.id;

CREATE TABLE public.weblate_auth_group_languages (
    id integer NOT NULL,
    group_id integer NOT NULL,
    language_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_group_languages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_group_languages_id_seq OWNED BY public.weblate_auth_group_languages.id;

CREATE TABLE public.weblate_auth_group_projects (
    id integer NOT NULL,
    group_id integer NOT NULL,
    project_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_group_projects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_group_projects_id_seq OWNED BY public.weblate_auth_group_projects.id;

CREATE TABLE public.weblate_auth_group_roles (
    id integer NOT NULL,
    group_id integer NOT NULL,
    role_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_group_roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_group_roles_id_seq OWNED BY public.weblate_auth_group_roles.id;

CREATE TABLE public.weblate_auth_invitation (
    uuid character(36) NOT NULL,
    backup_authorship boolean NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    author_id integer NOT NULL,
    user_id integer,
    username character varying(150) NOT NULL,
    full_name character varying(150) NOT NULL,
    group_id integer NOT NULL,
    email character varying(190) NOT NULL,
    automatic boolean NOT NULL,
    team_admin boolean NOT NULL,
    is_superuser boolean NOT NULL
);

CREATE TABLE public.weblate_auth_invitation_limit_languages (
    id integer NOT NULL,
    invitation_id character(36) NOT NULL,
    language_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_invitation_limit_languages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_invitation_limit_languages_id_seq OWNED BY public.weblate_auth_invitation_limit_languages.id;

CREATE TABLE public.weblate_auth_permission (
    id integer NOT NULL,
    codename character varying(100) NOT NULL,
    name character varying(200) NOT NULL
);

CREATE SEQUENCE public.weblate_auth_permission_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_permission_id_seq OWNED BY public.weblate_auth_permission.id;

CREATE TABLE public.weblate_auth_role (
    id integer NOT NULL,
    name character varying(200) NOT NULL
);

CREATE SEQUENCE public.weblate_auth_role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_role_id_seq OWNED BY public.weblate_auth_role.id;

CREATE TABLE public.weblate_auth_role_permissions (
    id integer NOT NULL,
    role_id integer NOT NULL,
    permission_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_role_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_role_permissions_id_seq OWNED BY public.weblate_auth_role_permissions.id;

CREATE TABLE public.weblate_auth_user (
    id integer NOT NULL,
    password character varying(128) NOT NULL,
    last_login timestamp(6) with time zone,
    username character varying(150) NOT NULL,
    full_name character varying(150) NOT NULL,
    email character varying(190),
    is_superuser boolean NOT NULL,
    is_active boolean NOT NULL,
    is_bot boolean NOT NULL,
    date_expires timestamp(6) with time zone,
    date_joined timestamp(6) with time zone NOT NULL
);

CREATE TABLE public.weblate_auth_user_groups (
    id integer NOT NULL,
    user_id integer NOT NULL,
    group_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_user_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_user_groups_id_seq OWNED BY public.weblate_auth_user_groups.id;

CREATE TABLE public.weblate_auth_user_groups_limit_languages (
    id integer NOT NULL,
    teammembership_id integer NOT NULL,
    language_id integer NOT NULL
);

CREATE SEQUENCE public.weblate_auth_user_groups_limit_languages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_user_groups_limit_languages_id_seq OWNED BY public.weblate_auth_user_groups_limit_languages.id;

CREATE SEQUENCE public.weblate_auth_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_user_id_seq OWNED BY public.weblate_auth_user.id;

CREATE TABLE public.weblate_auth_userblock (
    id integer NOT NULL,
    user_id integer NOT NULL,
    project_id integer NOT NULL,
    expiry timestamp(6) with time zone,
    note text NOT NULL
);

CREATE SEQUENCE public.weblate_auth_userblock_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.weblate_auth_userblock_id_seq OWNED BY public.weblate_auth_userblock.id;

CREATE TABLE public.wladmin_backuplog (
    id integer NOT NULL,
    service_id integer NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    event character varying(100) NOT NULL,
    warning boolean NOT NULL,
    log text NOT NULL
);

CREATE SEQUENCE public.wladmin_backuplog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.wladmin_backuplog_id_seq OWNED BY public.wladmin_backuplog.id;

CREATE TABLE public.wladmin_backupservice (
    id integer NOT NULL,
    repository character varying(500) NOT NULL,
    enabled boolean NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    passphrase character varying(100) NOT NULL,
    paperkey text NOT NULL
);

CREATE SEQUENCE public.wladmin_backupservice_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.wladmin_backupservice_id_seq OWNED BY public.wladmin_backupservice.id;

CREATE TABLE public.wladmin_configurationerror (
    id integer NOT NULL,
    name character varying(150) NOT NULL,
    message text NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    ignored boolean NOT NULL
);

CREATE SEQUENCE public.wladmin_configurationerror_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.wladmin_configurationerror_id_seq OWNED BY public.wladmin_configurationerror.id;

CREATE TABLE public.wladmin_securityfinding (
    id integer NOT NULL,
    user_id integer NOT NULL,
    finding_type character varying(50) NOT NULL,
    key character varying(64) NOT NULL,
    evidence jsonb NOT NULL,
    status character varying(10) NOT NULL,
    first_seen timestamp(6) with time zone NOT NULL,
    last_seen timestamp(6) with time zone NOT NULL,
    handled_at timestamp(6) with time zone
);

CREATE SEQUENCE public.wladmin_securityfinding_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.wladmin_securityfinding_id_seq OWNED BY public.wladmin_securityfinding.id;

CREATE TABLE public.wladmin_securityfindingevent (
    id integer NOT NULL,
    finding_id integer NOT NULL,
    scan_id character(36) NOT NULL,
    actor_id integer,
    action character varying(20) NOT NULL,
    "timestamp" timestamp(6) with time zone NOT NULL,
    reason character varying(50) NOT NULL,
    evidence jsonb NOT NULL
);

CREATE SEQUENCE public.wladmin_securityfindingevent_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.wladmin_securityfindingevent_id_seq OWNED BY public.wladmin_securityfindingevent.id;

CREATE TABLE public.wladmin_securityscan (
    id character(36) NOT NULL,
    actor_id integer,
    filename character varying(255) NOT NULL,
    operation character varying(20) NOT NULL,
    finding_type character varying(50) NOT NULL,
    state character varying(20) NOT NULL,
    created timestamp(6) with time zone NOT NULL,
    completed timestamp(6) with time zone,
    processed integer NOT NULL,
    total integer NOT NULL,
    summary jsonb NOT NULL,
    error text NOT NULL
);

CREATE TABLE public.wladmin_supportstatus (
    id integer NOT NULL,
    name character varying(150) NOT NULL,
    secret character varying(400) NOT NULL,
    expiry timestamp(6) with time zone,
    in_limits boolean NOT NULL,
    discoverable boolean NOT NULL,
    limits jsonb NOT NULL,
    has_subscription boolean NOT NULL,
    backup_repository character varying(500) NOT NULL,
    enabled boolean NOT NULL
);

CREATE SEQUENCE public.wladmin_supportstatus_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.wladmin_supportstatus_id_seq OWNED BY public.wladmin_supportstatus.id;

CREATE TABLE public.workspaces_workspace (
    id character(36) NOT NULL,
    metric_id integer NOT NULL,
    name character varying(100) NOT NULL,
    propagation_scope character varying(20) NOT NULL,
    name_managed boolean NOT NULL,
    use_workspace_tm boolean NOT NULL,
    contribute_workspace_tm boolean NOT NULL,
    license character varying(150) NOT NULL,
    agreement text NOT NULL,
    new_lang character varying(10) NOT NULL,
    language_code_style character varying(20) NOT NULL,
    secondary_language_id integer,
    check_flags text NOT NULL,
    enforced_checks jsonb NOT NULL,
    inherit_enforced_checks boolean NOT NULL,
    commit_message text NOT NULL,
    add_message text NOT NULL,
    delete_message text NOT NULL,
    merge_message text NOT NULL,
    addon_message text NOT NULL,
    pull_message text NOT NULL
);

ALTER TABLE ONLY public.accounts_auditlog ALTER COLUMN id SET DEFAULT nextval('public.accounts_auditlog_id_seq'::regclass);

ALTER TABLE ONLY public.accounts_profile ALTER COLUMN id SET DEFAULT nextval('public.accounts_profile_id_seq'::regclass);

ALTER TABLE ONLY public.accounts_profile_languages ALTER COLUMN id SET DEFAULT nextval('public.accounts_profile_languages_id_seq'::regclass);

ALTER TABLE ONLY public.accounts_profile_secondary_languages ALTER COLUMN id SET DEFAULT nextval('public.accounts_profile_secondary_languages_id_seq'::regclass);

ALTER TABLE ONLY public.accounts_profile_watched ALTER COLUMN id SET DEFAULT nextval('public.accounts_profile_watched_id_seq'::regclass);

ALTER TABLE ONLY public.accounts_subscription ALTER COLUMN id SET DEFAULT nextval('public.accounts_subscription_id_seq'::regclass);

ALTER TABLE ONLY public.accounts_teamauditlog ALTER COLUMN id SET DEFAULT nextval('public.accounts_teamauditlog_id_seq'::regclass);

ALTER TABLE ONLY public.accounts_verifiedemail ALTER COLUMN id SET DEFAULT nextval('public.accounts_verifiedemail_id_seq'::regclass);

ALTER TABLE ONLY public.addons_addon ALTER COLUMN id SET DEFAULT nextval('public.addons_addon_id_seq'::regclass);

ALTER TABLE ONLY public.addons_addonactivitylog ALTER COLUMN id SET DEFAULT nextval('public.addons_addonactivitylog_id_seq'::regclass);

ALTER TABLE ONLY public.addons_event ALTER COLUMN id SET DEFAULT nextval('public.addons_event_id_seq'::regclass);

ALTER TABLE ONLY public.auth_group ALTER COLUMN id SET DEFAULT nextval('public.auth_group_id_seq'::regclass);

ALTER TABLE ONLY public.auth_group_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_group_permissions_id_seq'::regclass);

ALTER TABLE ONLY public.auth_permission ALTER COLUMN id SET DEFAULT nextval('public.auth_permission_id_seq'::regclass);

ALTER TABLE ONLY public.checks_check ALTER COLUMN id SET DEFAULT nextval('public.checks_check_id_seq'::regclass);

ALTER TABLE ONLY public.configuration_setting ALTER COLUMN id SET DEFAULT nextval('public.configuration_setting_id_seq'::regclass);

ALTER TABLE ONLY public.django_admin_log ALTER COLUMN id SET DEFAULT nextval('public.django_admin_log_id_seq'::regclass);

ALTER TABLE ONLY public.django_celery_beat_clockedschedule ALTER COLUMN id SET DEFAULT nextval('public.django_celery_beat_clockedschedule_id_seq'::regclass);

ALTER TABLE ONLY public.django_celery_beat_crontabschedule ALTER COLUMN id SET DEFAULT nextval('public.django_celery_beat_crontabschedule_id_seq'::regclass);

ALTER TABLE ONLY public.django_celery_beat_intervalschedule ALTER COLUMN id SET DEFAULT nextval('public.django_celery_beat_intervalschedule_id_seq'::regclass);

ALTER TABLE ONLY public.django_celery_beat_periodictask ALTER COLUMN id SET DEFAULT nextval('public.django_celery_beat_periodictask_id_seq'::regclass);

ALTER TABLE ONLY public.django_celery_beat_periodictasks ALTER COLUMN ident SET DEFAULT nextval('public.django_celery_beat_periodictasks_ident_seq'::regclass);

ALTER TABLE ONLY public.django_celery_beat_solarschedule ALTER COLUMN id SET DEFAULT nextval('public.django_celery_beat_solarschedule_id_seq'::regclass);

ALTER TABLE ONLY public.django_content_type ALTER COLUMN id SET DEFAULT nextval('public.django_content_type_id_seq'::regclass);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnattestation ALTER COLUMN id SET DEFAULT nextval('public.django_otp_webauthn_webauthnattestation_id_seq'::regclass);

ALTER TABLE ONLY public.django_otp_webauthn_webauthncredential ALTER COLUMN id SET DEFAULT nextval('public.django_otp_webauthn_webauthncredential_id_seq'::regclass);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnuserhandle ALTER COLUMN user_id SET DEFAULT nextval('public.django_otp_webauthn_webauthnuserhandle_user_id_seq'::regclass);

ALTER TABLE ONLY public.fonts_font ALTER COLUMN id SET DEFAULT nextval('public.fonts_font_id_seq'::regclass);

ALTER TABLE ONLY public.fonts_fontgroup ALTER COLUMN id SET DEFAULT nextval('public.fonts_fontgroup_id_seq'::regclass);

ALTER TABLE ONLY public.fonts_fontoverride ALTER COLUMN id SET DEFAULT nextval('public.fonts_fontoverride_id_seq'::regclass);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact ALTER COLUMN id SET DEFAULT nextval('public.kotlin_sdk_kotlinsdkartifact_id_seq'::regclass);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact_builds ALTER COLUMN id SET DEFAULT nextval('public.kotlin_sdk_kotlinsdkartifact_builds_id_seq'::regclass);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkbuild ALTER COLUMN id SET DEFAULT nextval('public.kotlin_sdk_kotlinsdkbuild_id_seq'::regclass);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkcleanup ALTER COLUMN id SET DEFAULT nextval('public.kotlin_sdk_kotlinsdkcleanup_id_seq'::regclass);

ALTER TABLE ONLY public.lang_language ALTER COLUMN id SET DEFAULT nextval('public.lang_language_id_seq'::regclass);

ALTER TABLE ONLY public.lang_plural ALTER COLUMN id SET DEFAULT nextval('public.lang_plural_id_seq'::regclass);

ALTER TABLE ONLY public.machinery_machineryerror ALTER COLUMN id SET DEFAULT nextval('public.machinery_machineryerror_id_seq'::regclass);

ALTER TABLE ONLY public.machinery_machineryexclusion ALTER COLUMN id SET DEFAULT nextval('public.machinery_machineryexclusion_id_seq'::regclass);

ALTER TABLE ONLY public.machinery_machineryinstallation ALTER COLUMN id SET DEFAULT nextval('public.machinery_machineryinstallation_id_seq'::regclass);

ALTER TABLE ONLY public.memory_memory ALTER COLUMN id SET DEFAULT nextval('public.memory_memory_id_seq'::regclass);

ALTER TABLE ONLY public.memory_memoryscope ALTER COLUMN id SET DEFAULT nextval('public.memory_memoryscope_id_seq'::regclass);

ALTER TABLE ONLY public.metrics_metric ALTER COLUMN id SET DEFAULT nextval('public.metrics_metric_id_seq'::regclass);

ALTER TABLE ONLY public.otp_static_staticdevice ALTER COLUMN id SET DEFAULT nextval('public.otp_static_staticdevice_id_seq'::regclass);

ALTER TABLE ONLY public.otp_static_statictoken ALTER COLUMN id SET DEFAULT nextval('public.otp_static_statictoken_id_seq'::regclass);

ALTER TABLE ONLY public.otp_totp_totpdevice ALTER COLUMN id SET DEFAULT nextval('public.otp_totp_totpdevice_id_seq'::regclass);

ALTER TABLE ONLY public.screenshots_screenshot ALTER COLUMN id SET DEFAULT nextval('public.screenshots_screenshot_id_seq'::regclass);

ALTER TABLE ONLY public.screenshots_screenshot_units ALTER COLUMN id SET DEFAULT nextval('public.screenshots_screenshot_units_id_seq'::regclass);

ALTER TABLE ONLY public.screenshots_screenshotocr ALTER COLUMN screenshot_id SET DEFAULT nextval('public.screenshots_screenshotocr_screenshot_id_seq'::regclass);

ALTER TABLE ONLY public.social_auth_association ALTER COLUMN id SET DEFAULT nextval('public.social_auth_association_id_seq'::regclass);

ALTER TABLE ONLY public.social_auth_code ALTER COLUMN id SET DEFAULT nextval('public.social_auth_code_id_seq'::regclass);

ALTER TABLE ONLY public.social_auth_nonce ALTER COLUMN id SET DEFAULT nextval('public.social_auth_nonce_id_seq'::regclass);

ALTER TABLE ONLY public.social_auth_partial ALTER COLUMN id SET DEFAULT nextval('public.social_auth_partial_id_seq'::regclass);

ALTER TABLE ONLY public.social_auth_usersocialauth ALTER COLUMN id SET DEFAULT nextval('public.social_auth_usersocialauth_id_seq'::regclass);

ALTER TABLE ONLY public.trans_alert ALTER COLUMN id SET DEFAULT nextval('public.trans_alert_id_seq'::regclass);

ALTER TABLE ONLY public.trans_announcement ALTER COLUMN id SET DEFAULT nextval('public.trans_announcement_id_seq'::regclass);

ALTER TABLE ONLY public.trans_autocomponentlist ALTER COLUMN id SET DEFAULT nextval('public.trans_autocomponentlist_id_seq'::regclass);

ALTER TABLE ONLY public.trans_backupauthorclaim ALTER COLUMN id SET DEFAULT nextval('public.trans_backupauthorclaim_id_seq'::regclass);

ALTER TABLE ONLY public.trans_backupauthorvote ALTER COLUMN id SET DEFAULT nextval('public.trans_backupauthorvote_id_seq'::regclass);

ALTER TABLE ONLY public.trans_category ALTER COLUMN id SET DEFAULT nextval('public.trans_category_id_seq'::regclass);

ALTER TABLE ONLY public.trans_change ALTER COLUMN id SET DEFAULT nextval('public.trans_change_id_seq'::regclass);

ALTER TABLE ONLY public.trans_comment ALTER COLUMN id SET DEFAULT nextval('public.trans_comment_id_seq'::regclass);

ALTER TABLE ONLY public.trans_component ALTER COLUMN id SET DEFAULT nextval('public.trans_component_id_seq'::regclass);

ALTER TABLE ONLY public.trans_component_links ALTER COLUMN id SET DEFAULT nextval('public.trans_component_links_id_seq'::regclass);

ALTER TABLE ONLY public.trans_componentlist ALTER COLUMN id SET DEFAULT nextval('public.trans_componentlist_id_seq'::regclass);

ALTER TABLE ONLY public.trans_componentlist_components ALTER COLUMN id SET DEFAULT nextval('public.trans_componentlist_components_id_seq'::regclass);

ALTER TABLE ONLY public.trans_contributoragreement ALTER COLUMN id SET DEFAULT nextval('public.trans_contributoragreement_id_seq'::regclass);

ALTER TABLE ONLY public.trans_label ALTER COLUMN id SET DEFAULT nextval('public.trans_label_id_seq'::regclass);

ALTER TABLE ONLY public.trans_pendingunitchange ALTER COLUMN id SET DEFAULT nextval('public.trans_pendingunitchange_id_seq'::regclass);

ALTER TABLE ONLY public.trans_project ALTER COLUMN id SET DEFAULT nextval('public.trans_project_id_seq'::regclass);

ALTER TABLE ONLY public.trans_report ALTER COLUMN id SET DEFAULT nextval('public.trans_report_id_seq'::regclass);

ALTER TABLE ONLY public.trans_savedsearch ALTER COLUMN id SET DEFAULT nextval('public.trans_savedsearch_id_seq'::regclass);

ALTER TABLE ONLY public.trans_suggestion ALTER COLUMN id SET DEFAULT nextval('public.trans_suggestion_id_seq'::regclass);

ALTER TABLE ONLY public.trans_translation ALTER COLUMN id SET DEFAULT nextval('public.trans_translation_id_seq'::regclass);

ALTER TABLE ONLY public.trans_unit ALTER COLUMN id SET DEFAULT nextval('public.trans_unit_id_seq'::regclass);

ALTER TABLE ONLY public.trans_unit_labels ALTER COLUMN id SET DEFAULT nextval('public.trans_unit_labels_id_seq'::regclass);

ALTER TABLE ONLY public.trans_variant ALTER COLUMN id SET DEFAULT nextval('public.trans_variant_id_seq'::regclass);

ALTER TABLE ONLY public.trans_variant_defining_units ALTER COLUMN id SET DEFAULT nextval('public.trans_variant_defining_units_id_seq'::regclass);

ALTER TABLE ONLY public.trans_vote ALTER COLUMN id SET DEFAULT nextval('public.trans_vote_id_seq'::regclass);

ALTER TABLE ONLY public.trans_workflowsetting ALTER COLUMN id SET DEFAULT nextval('public.trans_workflowsetting_id_seq'::regclass);

ALTER TABLE ONLY public.vcs_githubappcredentials ALTER COLUMN id SET DEFAULT nextval('public.vcs_githubappcredentials_id_seq'::regclass);

ALTER TABLE ONLY public.vcs_installation ALTER COLUMN id SET DEFAULT nextval('public.vcs_installation_id_seq'::regclass);

ALTER TABLE ONLY public.vcs_pendinginstallation ALTER COLUMN id SET DEFAULT nextval('public.vcs_pendinginstallation_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_autogroup ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_autogroup_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_group ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_group_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_group_admins ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_group_admins_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_group_componentlists ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_group_componentlists_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_group_components ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_group_components_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_group_languages ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_group_languages_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_group_projects ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_group_projects_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_group_roles ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_group_roles_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_invitation_limit_languages ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_invitation_limit_languages_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_permission ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_permission_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_role ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_role_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_role_permissions ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_role_permissions_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_user ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_user_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_user_groups ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_user_groups_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_user_groups_limit_languages ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_user_groups_limit_languages_id_seq'::regclass);

ALTER TABLE ONLY public.weblate_auth_userblock ALTER COLUMN id SET DEFAULT nextval('public.weblate_auth_userblock_id_seq'::regclass);

ALTER TABLE ONLY public.wladmin_backuplog ALTER COLUMN id SET DEFAULT nextval('public.wladmin_backuplog_id_seq'::regclass);

ALTER TABLE ONLY public.wladmin_backupservice ALTER COLUMN id SET DEFAULT nextval('public.wladmin_backupservice_id_seq'::regclass);

ALTER TABLE ONLY public.wladmin_configurationerror ALTER COLUMN id SET DEFAULT nextval('public.wladmin_configurationerror_id_seq'::regclass);

ALTER TABLE ONLY public.wladmin_securityfinding ALTER COLUMN id SET DEFAULT nextval('public.wladmin_securityfinding_id_seq'::regclass);

ALTER TABLE ONLY public.wladmin_securityfindingevent ALTER COLUMN id SET DEFAULT nextval('public.wladmin_securityfindingevent_id_seq'::regclass);

ALTER TABLE ONLY public.wladmin_supportstatus ALTER COLUMN id SET DEFAULT nextval('public.wladmin_supportstatus_id_seq'::regclass);

ALTER TABLE ONLY public.accounts_auditlog
    ADD CONSTRAINT accounts_auditlog_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts_profile_languages
    ADD CONSTRAINT accounts_profile_languages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts_profile
    ADD CONSTRAINT accounts_profile_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts_profile_secondary_languages
    ADD CONSTRAINT accounts_profile_secondary_languages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts_profile_watched
    ADD CONSTRAINT accounts_profile_watched_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts_subscription
    ADD CONSTRAINT accounts_subscription_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT accounts_teamauditlog_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts_verifiedemail
    ADD CONSTRAINT accounts_verifiedemail_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.addons_addon
    ADD CONSTRAINT addons_addon_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.addons_addonactivitylog
    ADD CONSTRAINT addons_addonactivitylog_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.addons_event
    ADD CONSTRAINT addons_event_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.authtoken_token
    ADD CONSTRAINT authtoken_token_pkey PRIMARY KEY (key);

ALTER TABLE ONLY public.checks_check
    ADD CONSTRAINT checks_check_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.configuration_setting
    ADD CONSTRAINT configuration_setting_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_celery_beat_clockedschedule
    ADD CONSTRAINT django_celery_beat_clockedschedule_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_celery_beat_crontabschedule
    ADD CONSTRAINT django_celery_beat_crontabschedule_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_celery_beat_intervalschedule
    ADD CONSTRAINT django_celery_beat_intervalschedule_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT django_celery_beat_periodictask_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_celery_beat_periodictasks
    ADD CONSTRAINT django_celery_beat_periodictasks_pkey PRIMARY KEY (ident);

ALTER TABLE ONLY public.django_celery_beat_solarschedule
    ADD CONSTRAINT django_celery_beat_solarschedule_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnattestation
    ADD CONSTRAINT django_otp_webauthn_webauthnattestation_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthncredential
    ADD CONSTRAINT django_otp_webauthn_webauthncredential_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnuserhandle
    ADD CONSTRAINT django_otp_webauthn_webauthnuserhandle_pkey PRIMARY KEY (user_id);

ALTER TABLE ONLY public.django_session
    ADD CONSTRAINT django_session_pkey PRIMARY KEY (session_key);

ALTER TABLE ONLY public.fonts_font
    ADD CONSTRAINT fonts_font_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.fonts_fontgroup
    ADD CONSTRAINT fonts_fontgroup_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.fonts_fontoverride
    ADD CONSTRAINT fonts_fontoverride_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact_builds
    ADD CONSTRAINT kotlin_sdk_kotlinsdkartifact_builds_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact
    ADD CONSTRAINT kotlin_sdk_kotlinsdkartifact_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkbuild
    ADD CONSTRAINT kotlin_sdk_kotlinsdkbuild_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkcleanup
    ADD CONSTRAINT kotlin_sdk_kotlinsdkcleanup_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.lang_language
    ADD CONSTRAINT lang_language_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.lang_plural
    ADD CONSTRAINT lang_plural_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.machinery_machineryerror
    ADD CONSTRAINT machinery_machineryerror_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.machinery_machineryexclusion
    ADD CONSTRAINT machinery_machineryexclusion_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.machinery_machineryinstallation
    ADD CONSTRAINT machinery_machineryinstallation_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT memory_memory_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT memory_memoryscope_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.memory_memoryscopemigrationstate
    ADD CONSTRAINT memory_memoryscopemigrationstate_pkey PRIMARY KEY (name);

ALTER TABLE ONLY public.metrics_metric
    ADD CONSTRAINT metrics_metric_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.otp_static_staticdevice
    ADD CONSTRAINT otp_static_staticdevice_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.otp_static_statictoken
    ADD CONSTRAINT otp_static_statictoken_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.otp_totp_totpdevice
    ADD CONSTRAINT otp_totp_totpdevice_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.screenshots_screenshot
    ADD CONSTRAINT screenshots_screenshot_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.screenshots_screenshot_units
    ADD CONSTRAINT screenshots_screenshot_units_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.screenshots_screenshotocr
    ADD CONSTRAINT screenshots_screenshotocr_pkey PRIMARY KEY (screenshot_id);

ALTER TABLE ONLY public.social_auth_association
    ADD CONSTRAINT social_auth_association_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.social_auth_code
    ADD CONSTRAINT social_auth_code_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.social_auth_nonce
    ADD CONSTRAINT social_auth_nonce_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.social_auth_partial
    ADD CONSTRAINT social_auth_partial_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.social_auth_usersocialauth
    ADD CONSTRAINT social_auth_usersocialauth_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_alert
    ADD CONSTRAINT trans_alert_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT trans_announcement_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_autocomponentlist
    ADD CONSTRAINT trans_autocomponentlist_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_backupauthorclaim
    ADD CONSTRAINT trans_backupauthorclaim_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_backupauthorvote
    ADD CONSTRAINT trans_backupauthorvote_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_category
    ADD CONSTRAINT trans_category_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_comment
    ADD CONSTRAINT trans_comment_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_component_links
    ADD CONSTRAINT trans_component_links_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT trans_component_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_componentlist_components
    ADD CONSTRAINT trans_componentlist_components_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_componentlist
    ADD CONSTRAINT trans_componentlist_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT trans_contributoragreement_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_label
    ADD CONSTRAINT trans_label_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_pendingunitchange
    ADD CONSTRAINT trans_pendingunitchange_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_project
    ADD CONSTRAINT trans_project_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT trans_report_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_savedsearch
    ADD CONSTRAINT trans_savedsearch_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_suggestion
    ADD CONSTRAINT trans_suggestion_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_translation
    ADD CONSTRAINT trans_translation_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_unit_labels
    ADD CONSTRAINT trans_unit_labels_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT trans_unit_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_variant_defining_units
    ADD CONSTRAINT trans_variant_defining_units_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_variant
    ADD CONSTRAINT trans_variant_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_vote
    ADD CONSTRAINT trans_vote_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_workflowsetting
    ADD CONSTRAINT trans_workflowsetting_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.vcs_githubappcredentials
    ADD CONSTRAINT vcs_githubappcredentials_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.vcs_installation
    ADD CONSTRAINT vcs_installation_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.vcs_pendinginstallation
    ADD CONSTRAINT vcs_pendinginstallation_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_autogroup
    ADD CONSTRAINT weblate_auth_autogroup_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_group_admins
    ADD CONSTRAINT weblate_auth_group_admins_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_group_componentlists
    ADD CONSTRAINT weblate_auth_group_componentlists_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_group_components
    ADD CONSTRAINT weblate_auth_group_components_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_group_languages
    ADD CONSTRAINT weblate_auth_group_languages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_group
    ADD CONSTRAINT weblate_auth_group_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_group_projects
    ADD CONSTRAINT weblate_auth_group_projects_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_group_roles
    ADD CONSTRAINT weblate_auth_group_roles_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_invitation_limit_languages
    ADD CONSTRAINT weblate_auth_invitation_limit_languages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_invitation
    ADD CONSTRAINT weblate_auth_invitation_pkey PRIMARY KEY (uuid);

ALTER TABLE ONLY public.weblate_auth_permission
    ADD CONSTRAINT weblate_auth_permission_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_role_permissions
    ADD CONSTRAINT weblate_auth_role_permissions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_role
    ADD CONSTRAINT weblate_auth_role_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_user_groups_limit_languages
    ADD CONSTRAINT weblate_auth_user_groups_limit_languages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_user_groups
    ADD CONSTRAINT weblate_auth_user_groups_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_user
    ADD CONSTRAINT weblate_auth_user_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_userblock
    ADD CONSTRAINT weblate_auth_userblock_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_backuplog
    ADD CONSTRAINT wladmin_backuplog_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_backupservice
    ADD CONSTRAINT wladmin_backupservice_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_configurationerror
    ADD CONSTRAINT wladmin_configurationerror_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_securityfinding
    ADD CONSTRAINT wladmin_securityfinding_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_securityfindingevent
    ADD CONSTRAINT wladmin_securityfindingevent_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_securityscan
    ADD CONSTRAINT wladmin_securityscan_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_supportstatus
    ADD CONSTRAINT wladmin_supportstatus_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.workspaces_workspace
    ADD CONSTRAINT workspaces_workspace_pkey PRIMARY KEY (id);

CREATE INDEX accounts_auditlog_activity_8f58593c ON public.accounts_auditlog USING btree (activity);

CREATE INDEX accounts_auditlog_activity_8f58593c_like ON public.accounts_auditlog USING btree (activity);

CREATE INDEX accounts_auditlog_timestamp_fc82620d ON public.accounts_auditlog USING btree ("timestamp");

CREATE INDEX accounts_auditlog_user_id_1dfbca79 ON public.accounts_auditlog USING btree (user_id);

CREATE INDEX accounts_profile_commented_8e3a8722 ON public.accounts_profile USING btree (commented);

CREATE INDEX accounts_profile_dashboard_component_list_id_4730151f ON public.accounts_profile USING btree (dashboard_component_list_id);

CREATE INDEX accounts_profile_languages_language_id_194609ba ON public.accounts_profile_languages USING btree (language_id);

CREATE INDEX accounts_profile_languages_profile_id_eabf6019 ON public.accounts_profile_languages USING btree (profile_id);

CREATE UNIQUE INDEX accounts_profile_languages_profile_id_language_id_48036091_uniq ON public.accounts_profile_languages USING btree (profile_id, language_id);

CREATE UNIQUE INDEX accounts_profile_seconda_profile_id_language_id_cb800d1e_uniq ON public.accounts_profile_secondary_languages USING btree (profile_id, language_id);

CREATE INDEX accounts_profile_secondary_languages_language_id_e0f1e325 ON public.accounts_profile_secondary_languages USING btree (language_id);

CREATE INDEX accounts_profile_secondary_languages_profile_id_acf6974e ON public.accounts_profile_secondary_languages USING btree (profile_id);

CREATE INDEX accounts_profile_suggested_8e0fb145 ON public.accounts_profile USING btree (suggested);

CREATE INDEX accounts_profile_translated_45d8856d ON public.accounts_profile USING btree (translated);

CREATE INDEX accounts_profile_uploaded_f5223324 ON public.accounts_profile USING btree (uploaded);

CREATE UNIQUE INDEX accounts_profile_user_id_key ON public.accounts_profile USING btree (user_id);

CREATE INDEX accounts_profile_watched_profile_id_adce7487 ON public.accounts_profile_watched USING btree (profile_id);

CREATE UNIQUE INDEX accounts_profile_watched_profile_id_project_id_678d8754_uniq ON public.accounts_profile_watched USING btree (profile_id, project_id);

CREATE INDEX accounts_profile_watched_project_id_b8975f00 ON public.accounts_profile_watched USING btree (project_id);

CREATE INDEX accounts_subscription_component_id_f17fa0c7 ON public.accounts_subscription USING btree (component_id);

CREATE UNIQUE INDEX accounts_subscription_notification_unique ON public.accounts_subscription USING btree (user_id, notification, scope, project_id, component_id);

CREATE INDEX accounts_subscription_project_id_0ac0474e ON public.accounts_subscription USING btree (project_id);

CREATE INDEX accounts_subscription_user_id_980c85f5 ON public.accounts_subscription USING btree (user_id);

CREATE INDEX accounts_teamauditlog_actor_id_96cc352d ON public.accounts_teamauditlog USING btree (actor_id);

CREATE INDEX accounts_teamauditlog_project_id_c3db3056 ON public.accounts_teamauditlog USING btree (project_id);

CREATE INDEX accounts_teamauditlog_team_id_b59809a3 ON public.accounts_teamauditlog USING btree (team_id);

CREATE INDEX accounts_teamauditlog_timestamp_184b340e ON public.accounts_teamauditlog USING btree ("timestamp");

CREATE INDEX accounts_teamauditlog_workspace_id_43f2ba4f ON public.accounts_teamauditlog USING btree (workspace_id);

CREATE INDEX accounts_verifiedemail_social_id_c6d8a262 ON public.accounts_verifiedemail USING btree (social_id);

CREATE INDEX addons_addon_category_id_99fed34f ON public.addons_addon USING btree (category_id);

CREATE INDEX addons_addon_component_id_87dbd2fc ON public.addons_addon USING btree (component_id);

CREATE INDEX addons_addon_project_id_d2181050 ON public.addons_addon USING btree (project_id);

CREATE INDEX addons_addon_repo_scope_20c74be8 ON public.addons_addon USING btree (repo_scope);

CREATE INDEX addons_addonactivitylog_addon_id_44e38f98 ON public.addons_addonactivitylog USING btree (addon_id);

CREATE INDEX addons_addonactivitylog_component_id_baafffdc ON public.addons_addonactivitylog USING btree (component_id);

CREATE UNIQUE INDEX addons_event_addon_id_event_2407a5ef_uniq ON public.addons_event USING btree (addon_id, event);

CREATE INDEX auth_group_name_a6ea08ec_like ON public.auth_group USING btree (name);

CREATE UNIQUE INDEX auth_group_name_key ON public.auth_group USING btree (name);

CREATE INDEX auth_group_permissions_group_id_b120cbf9 ON public.auth_group_permissions USING btree (group_id);

CREATE UNIQUE INDEX auth_group_permissions_group_id_permission_id_0cd325b0_uniq ON public.auth_group_permissions USING btree (group_id, permission_id);

CREATE INDEX auth_group_permissions_permission_id_84c5c92e ON public.auth_group_permissions USING btree (permission_id);

CREATE INDEX auth_permission_content_type_id_2f476e4b ON public.auth_permission USING btree (content_type_id);

CREATE UNIQUE INDEX auth_permission_content_type_id_codename_01ab375a_uniq ON public.auth_permission USING btree (content_type_id, codename);

CREATE INDEX authtoken_token_key_10f0b77e_like ON public.authtoken_token USING btree (key);

CREATE UNIQUE INDEX authtoken_token_user_id_key ON public.authtoken_token USING btree (user_id);

CREATE UNIQUE INDEX category_name_unique ON public.trans_category USING btree (name, project_id, category_id);

CREATE UNIQUE INDEX category_slug_unique ON public.trans_category USING btree (slug, project_id, category_id);

CREATE INDEX checks_active_unit_idx ON public.checks_check USING btree (unit_id);

CREATE INDEX checks_check_dismissed_a11836a3 ON public.checks_check USING btree (dismissed);

CREATE UNIQUE INDEX checks_check_unit_id_name_ea73a317_uniq ON public.checks_check USING btree (unit_id, name);

CREATE UNIQUE INDEX component_name_unique ON public.trans_component USING btree (name, project_id, category_id);

CREATE UNIQUE INDEX component_slug_unique ON public.trans_component USING btree (slug, project_id, category_id);

CREATE UNIQUE INDEX configuration_setting_category_name_5144ffcf_uniq ON public.configuration_setting USING btree (category, name);

CREATE INDEX django_admin_log_content_type_id_c4bce8eb ON public.django_admin_log USING btree (content_type_id);

CREATE INDEX django_admin_log_user_id_c564eba6 ON public.django_admin_log USING btree (user_id);

CREATE INDEX django_celery_beat_periodictask_clocked_id_47a69f82 ON public.django_celery_beat_periodictask USING btree (clocked_id);

CREATE INDEX django_celery_beat_periodictask_crontab_id_d3cba168 ON public.django_celery_beat_periodictask USING btree (crontab_id);

CREATE INDEX django_celery_beat_periodictask_interval_id_a8ca27da ON public.django_celery_beat_periodictask USING btree (interval_id);

CREATE INDEX django_celery_beat_periodictask_name_265a36b7_like ON public.django_celery_beat_periodictask USING btree (name);

CREATE UNIQUE INDEX django_celery_beat_periodictask_name_key ON public.django_celery_beat_periodictask USING btree (name);

CREATE INDEX django_celery_beat_periodictask_solar_id_a87ce72c ON public.django_celery_beat_periodictask USING btree (solar_id);

CREATE UNIQUE INDEX django_celery_beat_solar_event_latitude_longitude_ba64999a_uniq ON public.django_celery_beat_solarschedule USING btree (event, latitude, longitude);

CREATE UNIQUE INDEX django_content_type_app_label_model_76bd3d3b_uniq ON public.django_content_type USING btree (app_label, model);

CREATE INDEX django_otp_webauthn_weba_credential_id_sha256_1a385252_like ON public.django_otp_webauthn_webauthncredential USING btree (credential_id_sha256);

CREATE UNIQUE INDEX django_otp_webauthn_webauthnattestation_credential_id_key ON public.django_otp_webauthn_webauthnattestation USING btree (credential_id);

CREATE UNIQUE INDEX django_otp_webauthn_webauthncredential_credential_id_sha256_key ON public.django_otp_webauthn_webauthncredential USING btree (credential_id_sha256);

CREATE INDEX django_otp_webauthn_webauthncredential_user_id_023391d8 ON public.django_otp_webauthn_webauthncredential USING btree (user_id);

CREATE INDEX django_otp_webauthn_webauthnuserhandle_handle_hex_ebcf77a1_like ON public.django_otp_webauthn_webauthnuserhandle USING btree (handle_hex);

CREATE UNIQUE INDEX django_otp_webauthn_webauthnuserhandle_handle_hex_key ON public.django_otp_webauthn_webauthnuserhandle USING btree (handle_hex);

CREATE INDEX django_session_expire_date_a5c62663 ON public.django_session USING btree (expire_date);

CREATE INDEX django_session_session_key_c0390e0f_like ON public.django_session USING btree (session_key);

CREATE UNIQUE INDEX fonts_font_family_style_project_id_9f3996b3_uniq ON public.fonts_font USING btree (family, style, project_id);

CREATE INDEX fonts_font_project_id_855704e0 ON public.fonts_font USING btree (project_id);

CREATE INDEX fonts_font_user_id_5897e323 ON public.fonts_font USING btree (user_id);

CREATE INDEX fonts_fontgroup_font_id_e4c04e1d ON public.fonts_fontgroup USING btree (font_id);

CREATE INDEX fonts_fontgroup_name_cf198a4d ON public.fonts_fontgroup USING btree (name);

CREATE INDEX fonts_fontgroup_name_cf198a4d_like ON public.fonts_fontgroup USING btree (name);

CREATE UNIQUE INDEX fonts_fontgroup_project_id_name_0b0fa938_uniq ON public.fonts_fontgroup USING btree (name, project_id);

CREATE INDEX fonts_fontoverride_font_id_ae703c52 ON public.fonts_fontoverride USING btree (font_id);

CREATE UNIQUE INDEX fonts_fontoverride_group_id_language_id_92a46b82_uniq ON public.fonts_fontoverride USING btree (group_id, language_id);

CREATE INDEX fonts_fontoverride_language_id_ccb8c819 ON public.fonts_fontoverride USING btree (language_id);

CREATE UNIQUE INDEX kotlin_sdk_artifact_unique ON public.kotlin_sdk_kotlinsdkartifact USING btree (addon_id, digest);

CREATE UNIQUE INDEX kotlin_sdk_build_unique ON public.kotlin_sdk_kotlinsdkbuild USING btree (addon_id, package_name, version_code);

CREATE UNIQUE INDEX kotlin_sdk_kotlinsdkarti_kotlinsdkartifact_id_kot_6c5e8d5f_uniq ON public.kotlin_sdk_kotlinsdkartifact_builds USING btree (kotlinsdkartifact_id, kotlinsdkbuild_id);

CREATE INDEX kotlin_sdk_kotlinsdkartifa_kotlinsdkartifact_id_ebd2906e ON public.kotlin_sdk_kotlinsdkartifact_builds USING btree (kotlinsdkartifact_id);

CREATE INDEX kotlin_sdk_kotlinsdkartifact_addon_id_ebe8e551 ON public.kotlin_sdk_kotlinsdkartifact USING btree (addon_id);

CREATE INDEX kotlin_sdk_kotlinsdkartifact_builds_kotlinsdkbuild_id_070b729b ON public.kotlin_sdk_kotlinsdkartifact_builds USING btree (kotlinsdkbuild_id);

CREATE INDEX kotlin_sdk_kotlinsdkbuild_addon_id_7ab4618d ON public.kotlin_sdk_kotlinsdkbuild USING btree (addon_id);

CREATE UNIQUE INDEX kotlin_sdk_kotlinsdkcleanup_build_id_key ON public.kotlin_sdk_kotlinsdkcleanup USING btree (build_id);

CREATE INDEX kotlin_sdk_kotlinsdkcleanup_path_eac6eb18_like ON public.kotlin_sdk_kotlinsdkcleanup USING btree (path);

CREATE UNIQUE INDEX kotlin_sdk_kotlinsdkcleanup_path_key ON public.kotlin_sdk_kotlinsdkcleanup USING btree (path);

CREATE INDEX lang_language_code_764a7a1d_like ON public.lang_language USING btree (code);

CREATE UNIQUE INDEX lang_language_code_key ON public.lang_language USING btree (code);

CREATE INDEX lang_plural_language_id_b5f4646a ON public.lang_plural USING btree (language_id);

CREATE INDEX machinery_m_install_ec7c4c_idx ON public.machinery_machineryerror USING btree (id, installation_id, "timestamp");

CREATE INDEX machinery_machineryerror_installation_id_721886a4 ON public.machinery_machineryerror USING btree (installation_id);

CREATE INDEX machinery_machineryerror_timestamp_07d8c763 ON public.machinery_machineryerror USING btree ("timestamp");

CREATE INDEX machinery_machineryexclusion_installation_id_f938fc8b ON public.machinery_machineryexclusion USING btree (installation_id);

CREATE INDEX machinery_machineryexclusion_project_id_d1dfc73d ON public.machinery_machineryexclusion USING btree (project_id);

CREATE INDEX machinery_machineryexclusion_workspace_id_7aa4f8fd ON public.machinery_machineryexclusion USING btree (workspace_id);

CREATE INDEX machinery_machineryinstallation_project_id_7821aa7d ON public.machinery_machineryinstallation USING btree (project_id);

CREATE INDEX machinery_machineryinstallation_slug_620b370d ON public.machinery_machineryinstallation USING btree (slug);

CREATE INDEX machinery_machineryinstallation_slug_620b370d_like ON public.machinery_machineryinstallation USING btree (slug);

CREATE INDEX machinery_machineryinstallation_workspace_id_071d1875 ON public.machinery_machineryinstallation USING btree (workspace_id);

CREATE UNIQUE INDEX machinery_project_installation_exclusion ON public.machinery_machineryexclusion USING btree (project_id, installation_id);

CREATE UNIQUE INDEX machinery_workspace_installation_exclusion ON public.machinery_machineryexclusion USING btree (workspace_id, installation_id);

CREATE INDEX memory_from_file ON public.memory_memory USING btree (legacy_from_file);

CREATE INDEX memory_memory_legacy_project_id_13aeb4a1 ON public.memory_memory USING btree (legacy_project_id);

CREATE INDEX memory_memory_legacy_user_id_aedef6a6 ON public.memory_memory USING btree (legacy_user_id);

CREATE INDEX memory_memory_source_language_id_c53ef4ff ON public.memory_memory USING btree (source_language_id);

CREATE INDEX memory_memory_target_language_id_c44dca3c ON public.memory_memory USING btree (target_language_id);

CREATE INDEX memory_memoryscope_memory_id_59a416c6 ON public.memory_memoryscope USING btree (memory_id);

CREATE INDEX memory_memoryscope_project_id_0f24de49 ON public.memory_memoryscope USING btree (project_id);

CREATE INDEX memory_memoryscope_source_component_id_bc705407 ON public.memory_memoryscope USING btree (source_component_id);

CREATE INDEX memory_memoryscope_source_project_id_080b6f0a ON public.memory_memoryscope USING btree (source_project_id);

CREATE INDEX memory_memoryscope_user_id_bccb55f8 ON public.memory_memoryscope USING btree (user_id);

CREATE INDEX memory_memoryscope_workspace_id_6d209c88 ON public.memory_memoryscope USING btree (workspace_id);

CREATE INDEX memory_memoryscopemigrationstate_name_d3f4ce7e_like ON public.memory_memoryscopemigrationstate USING btree (name);

CREATE INDEX memory_scope_project ON public.memory_memoryscope USING btree (memory_id, scope, project_id);

CREATE INDEX memory_scope_scope_memory ON public.memory_memoryscope USING btree (memory_id, scope);

CREATE INDEX memory_scope_source_component ON public.memory_memoryscope USING btree (memory_id, scope, source_component_id);

CREATE INDEX memory_scope_source_project ON public.memory_memoryscope USING btree (memory_id, scope, source_project_id);

CREATE INDEX memory_scope_user ON public.memory_memoryscope USING btree (memory_id, scope, user_id);

CREATE INDEX memory_scope_workspace ON public.memory_memoryscope USING btree (memory_id, scope, workspace_id);

CREATE INDEX memory_scope_workspace_source ON public.memory_memoryscope USING btree (memory_id, scope, workspace_id, source_project_id);

CREATE UNIQUE INDEX metrics_metric_scope_relation_secondary_date_f0ac693e_uniq ON public.metrics_metric USING btree (date, scope, relation, secondary);

CREATE INDEX otp_static_staticdevice_user_id_7f9cff2b ON public.otp_static_staticdevice USING btree (user_id);

CREATE INDEX otp_static_statictoken_device_id_74b7c7d1 ON public.otp_static_statictoken USING btree (device_id);

CREATE INDEX otp_static_statictoken_token_d0a51866 ON public.otp_static_statictoken USING btree (token);

CREATE INDEX otp_static_statictoken_token_d0a51866_like ON public.otp_static_statictoken USING btree (token);

CREATE INDEX otp_totp_totpdevice_user_id_0fb18292 ON public.otp_totp_totpdevice USING btree (user_id);

CREATE INDEX screenshots_screenshot_backup_author_claim_id_19b10d15 ON public.screenshots_screenshot USING btree (backup_author_claim_id);

CREATE INDEX screenshots_screenshot_translation_id_3b239850 ON public.screenshots_screenshot USING btree (translation_id);

CREATE UNIQUE INDEX screenshots_screenshot_u_screenshot_id_unit_id_c8f65eb6_uniq ON public.screenshots_screenshot_units USING btree (screenshot_id, unit_id);

CREATE INDEX screenshots_screenshot_units_screenshot_id_03db0a11 ON public.screenshots_screenshot_units USING btree (screenshot_id);

CREATE INDEX screenshots_screenshot_units_unit_id_2b24914f ON public.screenshots_screenshot_units USING btree (unit_id);

CREATE INDEX screenshots_screenshot_user_id_9350e149 ON public.screenshots_screenshot USING btree (user_id);

CREATE UNIQUE INDEX security_finding_event_identity ON public.wladmin_securityfindingevent USING btree (finding_id, scan_id, action);

CREATE UNIQUE INDEX security_finding_identity ON public.wladmin_securityfinding USING btree (user_id, finding_type, key);

CREATE UNIQUE INDEX social_auth_association_server_url_handle_078befa2_uniq ON public.social_auth_association USING btree (server_url, handle);

CREATE INDEX social_auth_code_code_a2393167 ON public.social_auth_code USING btree (code);

CREATE INDEX social_auth_code_code_a2393167_like ON public.social_auth_code USING btree (code);

CREATE UNIQUE INDEX social_auth_code_email_code_801b2d02_uniq ON public.social_auth_code USING btree (email, code);

CREATE INDEX social_auth_code_timestamp_176b341f ON public.social_auth_code USING btree ("timestamp");

CREATE UNIQUE INDEX social_auth_nonce_server_url_timestamp_salt_f6284463_uniq ON public.social_auth_nonce USING btree (server_url, "timestamp", salt);

CREATE INDEX social_auth_nonce_timestamp_bc4d078a ON public.social_auth_nonce USING btree ("timestamp");

CREATE INDEX social_auth_partial_timestamp_50f2119f ON public.social_auth_partial USING btree ("timestamp");

CREATE INDEX social_auth_partial_token_3017fea3 ON public.social_auth_partial USING btree (token);

CREATE INDEX social_auth_partial_token_3017fea3_like ON public.social_auth_partial USING btree (token);

CREATE UNIQUE INDEX social_auth_usersocialauth_provider_uid_e6b5e668_uniq ON public.social_auth_usersocialauth USING btree (provider, uid);

CREATE INDEX social_auth_usersocialauth_uid_796e51dc ON public.social_auth_usersocialauth USING btree (uid);

CREATE INDEX social_auth_usersocialauth_uid_796e51dc_like ON public.social_auth_usersocialauth USING btree (uid);

CREATE INDEX social_auth_usersocialauth_user_id_17d28448 ON public.social_auth_usersocialauth USING btree (user_id);

CREATE INDEX trans_alert_active_idx ON public.trans_alert USING btree (component_id, severity);

CREATE UNIQUE INDEX trans_alert_component_id_name_4a5a9b1d_uniq ON public.trans_alert USING btree (component_id, name);

CREATE INDEX trans_alert_dismissed_by_id_7fb54cb8 ON public.trans_alert USING btree (dismissed_by_id);

CREATE INDEX trans_alert_severity_3a81ac0d ON public.trans_alert USING btree (severity);

CREATE INDEX trans_announcement_category_id_913f87b9 ON public.trans_announcement USING btree (category_id);

CREATE INDEX trans_announcement_component_id_72695410 ON public.trans_announcement USING btree (component_id);

CREATE INDEX trans_announcement_expiry_a06d62a1 ON public.trans_announcement USING btree (expiry);

CREATE INDEX trans_announcement_language_id_f5f7a357 ON public.trans_announcement USING btree (language_id);

CREATE INDEX trans_announcement_project_id_3713e8a9 ON public.trans_announcement USING btree (project_id);

CREATE INDEX trans_autocomponentlist_componentlist_id_9ff87f21 ON public.trans_autocomponentlist USING btree (componentlist_id);

CREATE UNIQUE INDEX trans_backupauthorclaim_invitation_id_key ON public.trans_backupauthorclaim USING btree (invitation_id);

CREATE INDEX trans_backupauthorclaim_project_id_9f198408 ON public.trans_backupauthorclaim USING btree (project_id);

CREATE INDEX trans_backupauthorclaim_user_id_7d26dd3f ON public.trans_backupauthorclaim USING btree (user_id);

CREATE INDEX trans_backupauthorvote_claim_id_69890c5f ON public.trans_backupauthorvote USING btree (claim_id);

CREATE INDEX trans_backupauthorvote_placeholder_id_cda3b1bd ON public.trans_backupauthorvote USING btree (placeholder_id);

CREATE INDEX trans_backupauthorvote_suggestion_id_bdad2c90 ON public.trans_backupauthorvote USING btree (suggestion_id);

CREATE INDEX trans_category_category_id_0f21e824 ON public.trans_category USING btree (category_id);

CREATE INDEX trans_category_project_id_7dc717ed ON public.trans_category USING btree (project_id);

CREATE INDEX trans_category_secondary_language_id_3b074b19 ON public.trans_category USING btree (secondary_language_id);

CREATE INDEX trans_category_slug_52f13beb ON public.trans_category USING btree (slug);

CREATE INDEX trans_category_slug_52f13beb_like ON public.trans_category USING btree (slug);

CREATE INDEX trans_change_action_idx ON public.trans_change USING btree ("timestamp", action);

CREATE INDEX trans_change_addon_id_84220e1c ON public.trans_change USING btree (addon_id);

CREATE INDEX trans_change_alert_id_34808078 ON public.trans_change USING btree (alert_id);

CREATE INDEX trans_change_announce_idx ON public.trans_change USING btree ("timestamp");

CREATE INDEX trans_change_announcement_id_1ad27981 ON public.trans_change USING btree (announcement_id);

CREATE INDEX trans_change_backup_author_claim_id_69f41b5a ON public.trans_change USING btree (backup_author_claim_id);

CREATE INDEX trans_change_category_idx ON public.trans_change USING btree (category_id, "timestamp", action);

CREATE INDEX trans_change_comment_id_9702b436 ON public.trans_change USING btree (comment_id);

CREATE INDEX trans_change_component_idx ON public.trans_change USING btree (component_id, "timestamp", action);

CREATE INDEX trans_change_lang_cat_idx ON public.trans_change USING btree (language_id, category_id, "timestamp");

CREATE INDEX trans_change_lang_comp_idx ON public.trans_change USING btree (language_id, component_id, "timestamp");

CREATE INDEX trans_change_language_idx ON public.trans_change USING btree (language_id, "timestamp", action);

CREATE INDEX trans_change_prj_language_idx ON public.trans_change USING btree (language_id, project_id, "timestamp", action);

CREATE INDEX trans_change_proj_addon_idx ON public.trans_change USING btree (project_id, "timestamp");

CREATE INDEX trans_change_project_idx ON public.trans_change USING btree (project_id, "timestamp", action);

CREATE INDEX trans_change_screenshot_id_ebbee35e ON public.trans_change USING btree (screenshot_id);

CREATE INDEX trans_change_site_addon_idx ON public.trans_change USING btree ("timestamp");

CREATE INDEX trans_change_suggestion_id_0e61f1cb ON public.trans_change USING btree (suggestion_id);

CREATE INDEX trans_change_translation_idx ON public.trans_change USING btree (translation_id, "timestamp", action);

CREATE INDEX trans_change_unit_idx ON public.trans_change USING btree (unit_id, "timestamp", action);

CREATE INDEX trans_change_user_id_8d1225f9 ON public.trans_change USING btree (user_id);

CREATE INDEX trans_change_user_idx ON public.trans_change USING btree (user_id, "timestamp", action);

CREATE INDEX trans_change_workspace_idx ON public.trans_change USING btree (workspace_id, "timestamp", action);

CREATE INDEX trans_comment_backup_author_claim_id_3052bed5 ON public.trans_comment USING btree (backup_author_claim_id);

CREATE INDEX trans_comment_resolved_81cae7f4 ON public.trans_comment USING btree (resolved);

CREATE INDEX trans_comment_timestamp_1148a5ab ON public.trans_comment USING btree ("timestamp");

CREATE INDEX trans_comment_unit_id_73ca7c89 ON public.trans_comment USING btree (unit_id);

CREATE INDEX trans_comment_user_id_c2f200eb ON public.trans_comment USING btree (user_id);

CREATE INDEX trans_comp_repo_branch_idx ON public.trans_component USING btree (repo, branch);

CREATE INDEX trans_compo_project_d044e1_idx ON public.trans_component USING btree (project_id, allow_translation_propagation);

CREATE INDEX trans_component_category_id_7c9c5306 ON public.trans_component USING btree (category_id);

CREATE INDEX trans_component_is_glossary_b4b922b7 ON public.trans_component USING btree (is_glossary);

CREATE INDEX trans_component_linked_component_id_82b0de9e ON public.trans_component USING btree (linked_component_id);

CREATE INDEX trans_component_links_category_id_28d2bcf2 ON public.trans_component_links USING btree (category_id);

CREATE INDEX trans_component_links_component_id_2342af45 ON public.trans_component_links USING btree (component_id);

CREATE UNIQUE INDEX trans_component_links_component_id_project_id_8e996e95_uniq ON public.trans_component_links USING btree (component_id, project_id);

CREATE INDEX trans_component_links_project_id_840f1e11 ON public.trans_component_links USING btree (project_id);

CREATE INDEX trans_component_restricted_ec682d06 ON public.trans_component USING btree (restricted);

CREATE INDEX trans_component_secondary_language_id_496a8d45 ON public.trans_component USING btree (secondary_language_id);

CREATE INDEX trans_component_slug_5bd48be1 ON public.trans_component USING btree (slug);

CREATE INDEX trans_component_slug_5bd48be1_like ON public.trans_component USING btree (slug);

CREATE INDEX trans_component_source_language_id_2610b49e ON public.trans_component USING btree (source_language_id);

CREATE UNIQUE INDEX trans_componentlist_comp_componentlist_id_compone_99528c41_uniq ON public.trans_componentlist_components USING btree (componentlist_id, component_id);

CREATE INDEX trans_componentlist_components_component_id_5642ff5c ON public.trans_componentlist_components USING btree (component_id);

CREATE INDEX trans_componentlist_components_componentlist_id_44d9ba8f ON public.trans_componentlist_components USING btree (componentlist_id);

CREATE INDEX trans_componentlist_name_93196759_like ON public.trans_componentlist USING btree (name);

CREATE UNIQUE INDEX trans_componentlist_name_key ON public.trans_componentlist USING btree (name);

CREATE INDEX trans_componentlist_show_dashboard_1394f78f ON public.trans_componentlist USING btree (show_dashboard);

CREATE INDEX trans_componentlist_slug_05ac37c2_like ON public.trans_componentlist USING btree (slug);

CREATE UNIQUE INDEX trans_componentlist_slug_key ON public.trans_componentlist USING btree (slug);

CREATE INDEX trans_contributoragreement_category_id_16151d5d ON public.trans_contributoragreement USING btree (category_id);

CREATE INDEX trans_contributoragreement_component_id_31bafbef ON public.trans_contributoragreement USING btree (component_id);

CREATE INDEX trans_contributoragreement_project_id_5761269c ON public.trans_contributoragreement USING btree (project_id);

CREATE INDEX trans_contributoragreement_workspace_id_34926ae5 ON public.trans_contributoragreement USING btree (workspace_id);

CREATE UNIQUE INDEX trans_label_project_id_name_ef6f7cca_uniq ON public.trans_label USING btree (project_id, name);

CREATE INDEX trans_pendingunitchange_author_id_38073ef3 ON public.trans_pendingunitchange USING btree (author_id);

CREATE INDEX trans_pendingunitchange_backup_author_claim_id_96ae0409 ON public.trans_pendingunitchange USING btree (backup_author_claim_id);

CREATE INDEX trans_pendingunitchange_state_bbbe001f ON public.trans_pendingunitchange USING btree (state);

CREATE INDEX trans_pendingunitchange_timestamp_1f1271fd ON public.trans_pendingunitchange USING btree ("timestamp");

CREATE INDEX trans_pendingunitchange_unit_id_eda769c3 ON public.trans_pendingunitchange USING btree (unit_id);

CREATE INDEX trans_project_access_control_c17e787a ON public.trans_project USING btree (access_control);

CREATE INDEX trans_project_name_a3285a23_like ON public.trans_project USING btree (name);

CREATE UNIQUE INDEX trans_project_name_key ON public.trans_project USING btree (name);

CREATE INDEX trans_project_secondary_language_id_57e84b09 ON public.trans_project USING btree (secondary_language_id);

CREATE INDEX trans_project_slug_77cd973b_like ON public.trans_project USING btree (slug);

CREATE UNIQUE INDEX trans_project_slug_key ON public.trans_project USING btree (slug);

CREATE INDEX trans_project_workspace_id_a1945651 ON public.trans_project USING btree (workspace_id);

CREATE INDEX trans_report_category_id_fb4530ac ON public.trans_report USING btree (category_id);

CREATE INDEX trans_report_component_id_894b3313 ON public.trans_report USING btree (component_id);

CREATE INDEX trans_report_created_04e774e5 ON public.trans_report USING btree (created);

CREATE INDEX trans_report_creator_id_2b66af65 ON public.trans_report USING btree (creator_id);

CREATE INDEX trans_report_kind_82833535 ON public.trans_report USING btree (kind);

CREATE INDEX trans_report_kind_82833535_like ON public.trans_report USING btree (kind);

CREATE INDEX trans_report_project_id_026b7e0f ON public.trans_report USING btree (project_id);

CREATE INDEX trans_report_workspace_id_f914bdb5 ON public.trans_report USING btree (workspace_id);

CREATE UNIQUE INDEX trans_saved_search_project_name ON public.trans_savedsearch USING btree (project_id, name);

CREATE UNIQUE INDEX trans_saved_search_user_name ON public.trans_savedsearch USING btree (user_id, name);

CREATE INDEX trans_suggestion_backup_author_claim_id_f18fec6a ON public.trans_suggestion USING btree (backup_author_claim_id);

CREATE INDEX trans_suggestion_unit_id_eb46fceb ON public.trans_suggestion USING btree (unit_id);

CREATE INDEX trans_suggestion_user_id_69ce0c2b ON public.trans_suggestion USING btree (user_id);

CREATE UNIQUE INDEX trans_translation_component_id_language_id_331fa2a7_uniq ON public.trans_translation USING btree (component_id, language_id);

CREATE INDEX trans_translation_language_id_030f0b30 ON public.trans_translation USING btree (language_id);

CREATE INDEX trans_translation_plural_id_5cf36dc7 ON public.trans_translation USING btree (plural_id);

CREATE INDEX trans_unit_automatically_translated_96d2b31d ON public.trans_unit USING btree (automatically_translated);

CREATE INDEX trans_unit_labels_label_id_4ca92ebd ON public.trans_unit_labels USING btree (label_id);

CREATE INDEX trans_unit_labels_unit_id_a3c2ddc5 ON public.trans_unit_labels USING btree (unit_id);

CREATE UNIQUE INDEX trans_unit_labels_unit_id_label_id_6ab4c307_uniq ON public.trans_unit_labels USING btree (unit_id, label_id);

CREATE INDEX trans_unit_source_tm_idx ON public.trans_unit USING btree (source);

CREATE INDEX trans_unit_source_unit_id_7a735f87 ON public.trans_unit USING btree (source_unit_id);

CREATE UNIQUE INDEX trans_unit_translation_id_id_hash_ab945589_uniq ON public.trans_unit USING btree (translation_id, id_hash);

CREATE INDEX trans_unit_translation_parent_id_bad2a48e ON public.trans_unit USING btree (translation_parent_id);

CREATE INDEX trans_unit_variant_id_c3315c70 ON public.trans_unit USING btree (variant_id);

CREATE INDEX trans_variant_defining_units_unit_id_f692977b ON public.trans_variant_defining_units USING btree (unit_id);

CREATE INDEX trans_variant_defining_units_variant_id_20848c13 ON public.trans_variant_defining_units USING btree (variant_id);

CREATE UNIQUE INDEX trans_variant_defining_units_variant_id_unit_id_cf32c9b3_uniq ON public.trans_variant_defining_units USING btree (variant_id, unit_id);

CREATE UNIQUE INDEX trans_vote_suggestion_id_user_id_fe8bdce2_uniq ON public.trans_vote USING btree (suggestion_id, user_id);

CREATE INDEX trans_vote_user_id_86a644fd ON public.trans_vote USING btree (user_id);

CREATE INDEX trans_workflowsetting_language_id_fcb1fa15 ON public.trans_workflowsetting USING btree (language_id);

CREATE INDEX trans_workflowsetting_project_id_32daa6c8 ON public.trans_workflowsetting USING btree (project_id);

CREATE INDEX trans_workflowsetting_source_language_id_6df86150 ON public.trans_workflowsetting USING btree (source_language_id);

CREATE INDEX vcs_githubappcredentials_hostname_71ed5837_like ON public.vcs_githubappcredentials USING btree (hostname);

CREATE UNIQUE INDEX vcs_githubappcredentials_hostname_key ON public.vcs_githubappcredentials USING btree (hostname);

CREATE UNIQUE INDEX vcs_githubappcredentials_webhook_token_key ON public.vcs_githubappcredentials USING btree (webhook_token);

CREATE UNIQUE INDEX vcs_installation_provider_hostname_instal_ab8ed013_uniq ON public.vcs_installation USING btree (provider, installation_id, workspace_id, hostname);

CREATE INDEX vcs_installation_workspace_id_1a97507d ON public.vcs_installation USING btree (workspace_id);

CREATE UNIQUE INDEX vcs_pendinginstallation_provider_hostname_instal_1c297e27_uniq ON public.vcs_pendinginstallation USING btree (provider, hostname, installation_id);

CREATE INDEX webauthncredential_sha256_idx ON public.django_otp_webauthn_webauthncredential USING btree (credential_id_sha256);

CREATE INDEX weblate_auth_autogroup_group_id_37b7ff0b ON public.weblate_auth_autogroup USING btree (group_id);

CREATE INDEX weblate_auth_group_admins_group_id_fa29ae29 ON public.weblate_auth_group_admins USING btree (group_id);

CREATE UNIQUE INDEX weblate_auth_group_admins_group_id_user_id_8c69bc69_uniq ON public.weblate_auth_group_admins USING btree (group_id, user_id);

CREATE INDEX weblate_auth_group_admins_user_id_f5851c1b ON public.weblate_auth_group_admins USING btree (user_id);

CREATE UNIQUE INDEX weblate_auth_group_compo_group_id_component_id_57e2f006_uniq ON public.weblate_auth_group_components USING btree (group_id, component_id);

CREATE UNIQUE INDEX weblate_auth_group_compo_group_id_componentlist_i_4a7f6bad_uniq ON public.weblate_auth_group_componentlists USING btree (group_id, componentlist_id);

CREATE INDEX weblate_auth_group_componentlists_componentlist_id_b59f8fe9 ON public.weblate_auth_group_componentlists USING btree (componentlist_id);

CREATE INDEX weblate_auth_group_componentlists_group_id_7af40eb9 ON public.weblate_auth_group_componentlists USING btree (group_id);

CREATE INDEX weblate_auth_group_components_component_id_2b53d012 ON public.weblate_auth_group_components USING btree (component_id);

CREATE INDEX weblate_auth_group_components_group_id_a07e20aa ON public.weblate_auth_group_components USING btree (group_id);

CREATE INDEX weblate_auth_group_defining_project_id_b2ad81a7 ON public.weblate_auth_group USING btree (defining_project_id);

CREATE INDEX weblate_auth_group_defining_workspace_id_84f40bfd ON public.weblate_auth_group USING btree (defining_workspace_id);

CREATE INDEX weblate_auth_group_languages_group_id_aa12d0a6 ON public.weblate_auth_group_languages USING btree (group_id);

CREATE UNIQUE INDEX weblate_auth_group_languages_group_id_language_id_c34f29db_uniq ON public.weblate_auth_group_languages USING btree (group_id, language_id);

CREATE INDEX weblate_auth_group_languages_language_id_4b7f5d81 ON public.weblate_auth_group_languages USING btree (language_id);

CREATE INDEX weblate_auth_group_projects_group_id_f713a220 ON public.weblate_auth_group_projects USING btree (group_id);

CREATE UNIQUE INDEX weblate_auth_group_projects_group_id_project_id_2a91e8e2_uniq ON public.weblate_auth_group_projects USING btree (group_id, project_id);

CREATE INDEX weblate_auth_group_projects_project_id_26b8e2b6 ON public.weblate_auth_group_projects USING btree (project_id);

CREATE INDEX weblate_auth_group_roles_group_id_ed18f8d0 ON public.weblate_auth_group_roles USING btree (group_id);

CREATE UNIQUE INDEX weblate_auth_group_roles_group_id_role_id_9662132c_uniq ON public.weblate_auth_group_roles USING btree (group_id, role_id);

CREATE INDEX weblate_auth_group_roles_role_id_de6300ca ON public.weblate_auth_group_roles USING btree (role_id);

CREATE UNIQUE INDEX weblate_auth_invitation__invitation_id_language_i_f7270cfe_uniq ON public.weblate_auth_invitation_limit_languages USING btree (invitation_id, language_id);

CREATE INDEX weblate_auth_invitation_author_id_a713ef92 ON public.weblate_auth_invitation USING btree (author_id);

CREATE INDEX weblate_auth_invitation_group_id_07a52aff ON public.weblate_auth_invitation USING btree (group_id);

CREATE INDEX weblate_auth_invitation_limit_languages_invitation_id_ed3a1e77 ON public.weblate_auth_invitation_limit_languages USING btree (invitation_id);

CREATE INDEX weblate_auth_invitation_limit_languages_language_id_26454a1a ON public.weblate_auth_invitation_limit_languages USING btree (language_id);

CREATE INDEX weblate_auth_invitation_user_id_6a0c6dec ON public.weblate_auth_invitation USING btree (user_id);

CREATE INDEX weblate_auth_permission_codename_b551931f_like ON public.weblate_auth_permission USING btree (codename);

CREATE UNIQUE INDEX weblate_auth_permission_codename_key ON public.weblate_auth_permission USING btree (codename);

CREATE INDEX weblate_auth_role_name_97854f8d_like ON public.weblate_auth_role USING btree (name);

CREATE UNIQUE INDEX weblate_auth_role_name_key ON public.weblate_auth_role USING btree (name);

CREATE UNIQUE INDEX weblate_auth_role_permis_role_id_permission_id_58e25e9f_uniq ON public.weblate_auth_role_permissions USING btree (role_id, permission_id);

CREATE INDEX weblate_auth_role_permissions_permission_id_d0650152 ON public.weblate_auth_role_permissions USING btree (permission_id);

CREATE INDEX weblate_auth_role_permissions_role_id_fa163124 ON public.weblate_auth_role_permissions USING btree (role_id);

CREATE INDEX weblate_auth_user_email_aaabd769_like ON public.weblate_auth_user USING btree (email);

CREATE UNIQUE INDEX weblate_auth_user_email_key ON public.weblate_auth_user USING btree (email);

CREATE INDEX weblate_auth_user_groups_group_id_48c0a977 ON public.weblate_auth_user_groups USING btree (group_id);

CREATE INDEX weblate_auth_user_groups_l_teammembership_id_8da6367c ON public.weblate_auth_user_groups_limit_languages USING btree (teammembership_id);

CREATE INDEX weblate_auth_user_groups_limit_languages_language_id_52ecbc9e ON public.weblate_auth_user_groups_limit_languages USING btree (language_id);

CREATE UNIQUE INDEX weblate_auth_user_groups_teammembership_id_langua_73451555_uniq ON public.weblate_auth_user_groups_limit_languages USING btree (teammembership_id, language_id);

CREATE INDEX weblate_auth_user_groups_user_id_37e61ff1 ON public.weblate_auth_user_groups USING btree (user_id);

CREATE UNIQUE INDEX weblate_auth_user_groups_user_id_group_id_16cfc05b_uniq ON public.weblate_auth_user_groups USING btree (user_id, group_id);

CREATE INDEX weblate_auth_user_is_bot_f4b1c6fb ON public.weblate_auth_user USING btree (is_bot);

CREATE INDEX weblate_auth_user_username_d226e4a0_like ON public.weblate_auth_user USING btree (username);

CREATE UNIQUE INDEX weblate_auth_user_username_key ON public.weblate_auth_user USING btree (username);

CREATE INDEX weblate_auth_userblock_project_id_44aae0be ON public.weblate_auth_userblock USING btree (project_id);

CREATE UNIQUE INDEX weblate_auth_userblock_user_id_project_id_7e74a707_uniq ON public.weblate_auth_userblock USING btree (user_id, project_id);

CREATE INDEX wladmin_backuplog_event_fbf23e08 ON public.wladmin_backuplog USING btree (event);

CREATE INDEX wladmin_backuplog_event_fbf23e08_like ON public.wladmin_backuplog USING btree (event);

CREATE INDEX wladmin_backuplog_service_id_10292c2d ON public.wladmin_backuplog USING btree (service_id);

CREATE INDEX wladmin_con_ignored_fb498d_idx ON public.wladmin_configurationerror USING btree ("timestamp", ignored);

CREATE INDEX wladmin_configurationerror_ignored_8a22575f ON public.wladmin_configurationerror USING btree (ignored);

CREATE INDEX wladmin_configurationerror_name_99ad2168_like ON public.wladmin_configurationerror USING btree (name);

CREATE UNIQUE INDEX wladmin_configurationerror_name_key ON public.wladmin_configurationerror USING btree (name);

CREATE INDEX wladmin_sec_status_ec922b_idx ON public.wladmin_securityfinding USING btree (finding_type, status);

CREATE INDEX wladmin_securityfinding_user_id_6c31bc00 ON public.wladmin_securityfinding USING btree (user_id);

CREATE INDEX wladmin_securityfindingevent_actor_id_ddbbf0da ON public.wladmin_securityfindingevent USING btree (actor_id);

CREATE INDEX wladmin_securityfindingevent_finding_id_fb30cf88 ON public.wladmin_securityfindingevent USING btree (finding_id);

CREATE INDEX wladmin_securityfindingevent_scan_id_573ba79c ON public.wladmin_securityfindingevent USING btree (scan_id);

CREATE INDEX wladmin_securityscan_actor_id_0aebe956 ON public.wladmin_securityscan USING btree (actor_id);

CREATE INDEX wladmin_securityscan_state_135c0fa8 ON public.wladmin_securityscan USING btree (state);

CREATE INDEX wladmin_securityscan_state_135c0fa8_like ON public.wladmin_securityscan USING btree (state);

CREATE INDEX wladmin_supportstatus_enabled_fbe08b47 ON public.wladmin_supportstatus USING btree (enabled);

CREATE INDEX wladmin_supportstatus_expiry_2855a549 ON public.wladmin_supportstatus USING btree (expiry);

CREATE UNIQUE INDEX workspaces_workspace_metric_id_key ON public.workspaces_workspace USING btree (metric_id);

CREATE INDEX workspaces_workspace_secondary_language_id_b8ff77e6 ON public.workspaces_workspace USING btree (secondary_language_id);

ALTER TABLE ONLY public.accounts_auditlog
    ADD CONSTRAINT accounts_auditlog_user_id_1dfbca79_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.accounts_profile
    ADD CONSTRAINT accounts_profile_dashboard_component__4730151f_fk_trans_com FOREIGN KEY (dashboard_component_list_id) REFERENCES public.trans_componentlist(id);

ALTER TABLE ONLY public.accounts_profile_languages
    ADD CONSTRAINT accounts_profile_lan_language_id_194609ba_fk_lang_lang FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.accounts_profile_languages
    ADD CONSTRAINT accounts_profile_lan_profile_id_eabf6019_fk_accounts_ FOREIGN KEY (profile_id) REFERENCES public.accounts_profile(id);

ALTER TABLE ONLY public.accounts_profile_secondary_languages
    ADD CONSTRAINT accounts_profile_sec_language_id_e0f1e325_fk_lang_lang FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.accounts_profile_secondary_languages
    ADD CONSTRAINT accounts_profile_sec_profile_id_acf6974e_fk_accounts_ FOREIGN KEY (profile_id) REFERENCES public.accounts_profile(id);

ALTER TABLE ONLY public.accounts_profile
    ADD CONSTRAINT accounts_profile_user_id_49a85d32_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.accounts_profile_watched
    ADD CONSTRAINT accounts_profile_wat_profile_id_adce7487_fk_accounts_ FOREIGN KEY (profile_id) REFERENCES public.accounts_profile(id);

ALTER TABLE ONLY public.accounts_profile_watched
    ADD CONSTRAINT accounts_profile_wat_project_id_b8975f00_fk_trans_pro FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.accounts_subscription
    ADD CONSTRAINT accounts_subscriptio_component_id_f17fa0c7_fk_trans_com FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.accounts_subscription
    ADD CONSTRAINT accounts_subscription_project_id_0ac0474e_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.accounts_subscription
    ADD CONSTRAINT accounts_subscription_user_id_980c85f5_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT accounts_teamauditlo_workspace_id_43f2ba4f_fk_workspace FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT accounts_teamauditlog_actor_id_96cc352d_fk_weblate_auth_user_id FOREIGN KEY (actor_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT accounts_teamauditlog_project_id_c3db3056_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT accounts_teamauditlog_team_id_b59809a3_fk_weblate_auth_group_id FOREIGN KEY (team_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.accounts_verifiedemail
    ADD CONSTRAINT accounts_verifiedema_social_id_c6d8a262_fk_social_au FOREIGN KEY (social_id) REFERENCES public.social_auth_usersocialauth(id);

ALTER TABLE ONLY public.addons_addon
    ADD CONSTRAINT addons_addon_category_id_99fed34f_fk_trans_category_id FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.addons_addon
    ADD CONSTRAINT addons_addon_component_id_87dbd2fc_fk_trans_component_id FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.addons_addon
    ADD CONSTRAINT addons_addon_project_id_d2181050_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.addons_addonactivitylog
    ADD CONSTRAINT addons_addonactivity_component_id_baafffdc_fk_trans_com FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.addons_addonactivitylog
    ADD CONSTRAINT addons_addonactivitylog_addon_id_44e38f98_fk_addons_addon_id FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.addons_event
    ADD CONSTRAINT addons_event_addon_id_66d8d06c_fk_addons_addon_id FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissio_permission_id_84c5c92e_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_b120cbf9_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_2f476e4b_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.authtoken_token
    ADD CONSTRAINT authtoken_token_user_id_35299eff_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.checks_check
    ADD CONSTRAINT checks_check_unit_id_c312ecef_fk_trans_unit_id FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_content_type_id_c4bce8eb_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_user_id_c564eba6_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT django_celery_beat_p_clocked_id_47a69f82_fk_django_ce FOREIGN KEY (clocked_id) REFERENCES public.django_celery_beat_clockedschedule(id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT django_celery_beat_p_crontab_id_d3cba168_fk_django_ce FOREIGN KEY (crontab_id) REFERENCES public.django_celery_beat_crontabschedule(id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT django_celery_beat_p_interval_id_a8ca27da_fk_django_ce FOREIGN KEY (interval_id) REFERENCES public.django_celery_beat_intervalschedule(id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT django_celery_beat_p_solar_id_a87ce72c_fk_django_ce FOREIGN KEY (solar_id) REFERENCES public.django_celery_beat_solarschedule(id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnattestation
    ADD CONSTRAINT django_otp_webauthn__credential_id_b2e6f7e8_fk_django_ot FOREIGN KEY (credential_id) REFERENCES public.django_otp_webauthn_webauthncredential(id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthncredential
    ADD CONSTRAINT django_otp_webauthn__user_id_023391d8_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnuserhandle
    ADD CONSTRAINT django_otp_webauthn__user_id_47f479ce_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.fonts_font
    ADD CONSTRAINT fonts_font_project_id_855704e0_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.fonts_font
    ADD CONSTRAINT fonts_font_user_id_5897e323_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.fonts_fontgroup
    ADD CONSTRAINT fonts_fontgroup_font_id_e4c04e1d_fk_fonts_font_id FOREIGN KEY (font_id) REFERENCES public.fonts_font(id);

ALTER TABLE ONLY public.fonts_fontgroup
    ADD CONSTRAINT fonts_fontgroup_project_id_5d47adf5_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.fonts_fontoverride
    ADD CONSTRAINT fonts_fontoverride_font_id_ae703c52_fk_fonts_font_id FOREIGN KEY (font_id) REFERENCES public.fonts_font(id);

ALTER TABLE ONLY public.fonts_fontoverride
    ADD CONSTRAINT fonts_fontoverride_group_id_cab65c09_fk_fonts_fontgroup_id FOREIGN KEY (group_id) REFERENCES public.fonts_fontgroup(id);

ALTER TABLE ONLY public.fonts_fontoverride
    ADD CONSTRAINT fonts_fontoverride_language_id_ccb8c819_fk_lang_language_id FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact
    ADD CONSTRAINT kotlin_sdk_kotlinsdk_addon_id_ebe8e551_fk_addons_ad FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkcleanup
    ADD CONSTRAINT kotlin_sdk_kotlinsdk_build_id_5cb8cade_fk_kotlin_sd FOREIGN KEY (build_id) REFERENCES public.kotlin_sdk_kotlinsdkbuild(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact_builds
    ADD CONSTRAINT kotlin_sdk_kotlinsdk_kotlinsdkartifact_id_ebd2906e_fk_kotlin_sd FOREIGN KEY (kotlinsdkartifact_id) REFERENCES public.kotlin_sdk_kotlinsdkartifact(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact_builds
    ADD CONSTRAINT kotlin_sdk_kotlinsdk_kotlinsdkbuild_id_070b729b_fk_kotlin_sd FOREIGN KEY (kotlinsdkbuild_id) REFERENCES public.kotlin_sdk_kotlinsdkbuild(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkbuild
    ADD CONSTRAINT kotlin_sdk_kotlinsdkbuild_addon_id_7ab4618d_fk_addons_addon_id FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.lang_plural
    ADD CONSTRAINT lang_plural_language_id_b5f4646a_fk_lang_language_id FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.machinery_machineryerror
    ADD CONSTRAINT machinery_machinerye_installation_id_721886a4_fk_machinery FOREIGN KEY (installation_id) REFERENCES public.machinery_machineryinstallation(id);

ALTER TABLE ONLY public.machinery_machineryexclusion
    ADD CONSTRAINT machinery_machinerye_installation_id_f938fc8b_fk_machinery FOREIGN KEY (installation_id) REFERENCES public.machinery_machineryinstallation(id);

ALTER TABLE ONLY public.machinery_machineryexclusion
    ADD CONSTRAINT machinery_machinerye_project_id_d1dfc73d_fk_trans_pro FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.machinery_machineryexclusion
    ADD CONSTRAINT machinery_machinerye_workspace_id_7aa4f8fd_fk_workspace FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.machinery_machineryinstallation
    ADD CONSTRAINT machinery_machineryi_project_id_7821aa7d_fk_trans_pro FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.machinery_machineryinstallation
    ADD CONSTRAINT machinery_machineryi_workspace_id_071d1875_fk_workspace FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT memory_memory_legacy_project_id_13aeb4a1_fk_trans_project_id FOREIGN KEY (legacy_project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT memory_memory_legacy_user_id_aedef6a6_fk_weblate_auth_user_id FOREIGN KEY (legacy_user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT memory_memory_source_language_id_c53ef4ff_fk_lang_language_id FOREIGN KEY (source_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT memory_memory_target_language_id_c44dca3c_fk_lang_language_id FOREIGN KEY (target_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT memory_memoryscope_memory_id_59a416c6_fk_memory_memory_id FOREIGN KEY (memory_id) REFERENCES public.memory_memory(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT memory_memoryscope_project_id_0f24de49_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT memory_memoryscope_source_component_id_bc705407_fk_trans_com FOREIGN KEY (source_component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT memory_memoryscope_source_project_id_080b6f0a_fk_trans_pro FOREIGN KEY (source_project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT memory_memoryscope_user_id_bccb55f8_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT memory_memoryscope_workspace_id_6d209c88_fk_workspace FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.otp_static_staticdevice
    ADD CONSTRAINT otp_static_staticdev_user_id_7f9cff2b_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.otp_static_statictoken
    ADD CONSTRAINT otp_static_statictok_device_id_74b7c7d1_fk_otp_stati FOREIGN KEY (device_id) REFERENCES public.otp_static_staticdevice(id);

ALTER TABLE ONLY public.otp_totp_totpdevice
    ADD CONSTRAINT otp_totp_totpdevice_user_id_0fb18292_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.screenshots_screenshot
    ADD CONSTRAINT screenshots_screensh_backup_author_claim__19b10d15_fk_trans_bac FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.screenshots_screenshot_units
    ADD CONSTRAINT screenshots_screensh_screenshot_id_03db0a11_fk_screensho FOREIGN KEY (screenshot_id) REFERENCES public.screenshots_screenshot(id);

ALTER TABLE ONLY public.screenshots_screenshotocr
    ADD CONSTRAINT screenshots_screensh_screenshot_id_845ecce4_fk_screensho FOREIGN KEY (screenshot_id) REFERENCES public.screenshots_screenshot(id);

ALTER TABLE ONLY public.screenshots_screenshot
    ADD CONSTRAINT screenshots_screensh_translation_id_3b239850_fk_trans_tra FOREIGN KEY (translation_id) REFERENCES public.trans_translation(id);

ALTER TABLE ONLY public.screenshots_screenshot_units
    ADD CONSTRAINT screenshots_screenshot_units_unit_id_2b24914f_fk_trans_unit_id FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.screenshots_screenshot
    ADD CONSTRAINT screenshots_screenshot_user_id_9350e149_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.social_auth_usersocialauth
    ADD CONSTRAINT social_auth_usersoci_user_id_17d28448_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_alert
    ADD CONSTRAINT trans_alert_component_id_e9bdfacb_fk_trans_component_id FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_alert
    ADD CONSTRAINT trans_alert_dismissed_by_id_7fb54cb8_fk_weblate_auth_user_id FOREIGN KEY (dismissed_by_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT trans_announcement_category_id_913f87b9_fk_trans_category_id FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT trans_announcement_component_id_72695410_fk_trans_component_id FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT trans_announcement_language_id_f5f7a357_fk_lang_language_id FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT trans_announcement_project_id_3713e8a9_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_autocomponentlist
    ADD CONSTRAINT trans_autocomponentl_componentlist_id_9ff87f21_fk_trans_com FOREIGN KEY (componentlist_id) REFERENCES public.trans_componentlist(id);

ALTER TABLE ONLY public.trans_backupauthorclaim
    ADD CONSTRAINT trans_backupauthorcl_invitation_id_d0b51998_fk_weblate_a FOREIGN KEY (invitation_id) REFERENCES public.weblate_auth_invitation(uuid);

ALTER TABLE ONLY public.trans_backupauthorclaim
    ADD CONSTRAINT trans_backupauthorcl_user_id_7d26dd3f_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_backupauthorclaim
    ADD CONSTRAINT trans_backupauthorclaim_project_id_9f198408_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_backupauthorvote
    ADD CONSTRAINT trans_backupauthorvo_claim_id_69890c5f_fk_trans_bac FOREIGN KEY (claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_backupauthorvote
    ADD CONSTRAINT trans_backupauthorvo_suggestion_id_bdad2c90_fk_trans_sug FOREIGN KEY (suggestion_id) REFERENCES public.trans_suggestion(id);

ALTER TABLE ONLY public.trans_backupauthorvote
    ADD CONSTRAINT trans_backupauthorvote_placeholder_id_cda3b1bd_fk_trans_vote_id FOREIGN KEY (placeholder_id) REFERENCES public.trans_vote(id);

ALTER TABLE ONLY public.trans_category
    ADD CONSTRAINT trans_category_category_id_0f21e824_fk_trans_category_id FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_category
    ADD CONSTRAINT trans_category_project_id_7dc717ed_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_category
    ADD CONSTRAINT trans_category_secondary_language_i_3b074b19_fk_lang_lang FOREIGN KEY (secondary_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_addon_id_84220e1c_fk_addons_addon_id FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_alert_id_34808078_fk_trans_alert_id FOREIGN KEY (alert_id) REFERENCES public.trans_alert(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_announcement_id_1ad27981_fk_trans_announcement_id FOREIGN KEY (announcement_id) REFERENCES public.trans_announcement(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_author_id_09329d48_fk_weblate_auth_user_id FOREIGN KEY (author_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_backup_author_claim__69f41b5a_fk_trans_bac FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_category_id_436958e0_fk_trans_category_id FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_comment_id_9702b436_fk_trans_comment_id FOREIGN KEY (comment_id) REFERENCES public.trans_comment(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_component_id_288ed925_fk_trans_component_id FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_language_id_73bd2ce4_fk_lang_language_id FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_project_id_b2db51ee_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_screenshot_id_ebbee35e_fk_screensho FOREIGN KEY (screenshot_id) REFERENCES public.screenshots_screenshot(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_suggestion_id_0e61f1cb_fk_trans_suggestion_id FOREIGN KEY (suggestion_id) REFERENCES public.trans_suggestion(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_translation_id_c5aa927b_fk_trans_translation_id FOREIGN KEY (translation_id) REFERENCES public.trans_translation(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_unit_id_76b0c6c0_fk_trans_unit_id FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_user_id_8d1225f9_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT trans_change_workspace_id_a1fbe04c_fk_workspaces_workspace_id FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.trans_comment
    ADD CONSTRAINT trans_comment_backup_author_claim__3052bed5_fk_trans_bac FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_comment
    ADD CONSTRAINT trans_comment_unit_id_73ca7c89_fk_trans_unit_id FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_comment
    ADD CONSTRAINT trans_comment_user_id_c2f200eb_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT trans_component_category_id_7c9c5306_fk_trans_category_id FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_component_links
    ADD CONSTRAINT trans_component_link_component_id_2342af45_fk_trans_com FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT trans_component_linked_component_id_82b0de9e_fk_trans_com FOREIGN KEY (linked_component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_component_links
    ADD CONSTRAINT trans_component_links_category_id_28d2bcf2_fk_trans_category_id FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_component_links
    ADD CONSTRAINT trans_component_links_project_id_840f1e11_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT trans_component_project_id_04a8b52c_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT trans_component_secondary_language_i_496a8d45_fk_lang_lang FOREIGN KEY (secondary_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT trans_component_source_language_id_2610b49e_fk_lang_language_id FOREIGN KEY (source_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_componentlist_components
    ADD CONSTRAINT trans_componentlist__component_id_5642ff5c_fk_trans_com FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_componentlist_components
    ADD CONSTRAINT trans_componentlist__componentlist_id_44d9ba8f_fk_trans_com FOREIGN KEY (componentlist_id) REFERENCES public.trans_componentlist(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT trans_contributoragr_category_id_16151d5d_fk_trans_cat FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT trans_contributoragr_component_id_31bafbef_fk_trans_com FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT trans_contributoragr_project_id_5761269c_fk_trans_pro FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT trans_contributoragr_user_id_220dedfc_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT trans_contributoragr_workspace_id_34926ae5_fk_workspace FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.trans_label
    ADD CONSTRAINT trans_label_project_id_8320ed70_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_pendingunitchange
    ADD CONSTRAINT trans_pendingunitcha_author_id_38073ef3_fk_weblate_a FOREIGN KEY (author_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_pendingunitchange
    ADD CONSTRAINT trans_pendingunitcha_backup_author_claim__96ae0409_fk_trans_bac FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_pendingunitchange
    ADD CONSTRAINT trans_pendingunitchange_unit_id_eda769c3_fk_trans_unit_id FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_project
    ADD CONSTRAINT trans_project_secondary_language_i_57e84b09_fk_lang_lang FOREIGN KEY (secondary_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_project
    ADD CONSTRAINT trans_project_workspace_id_a1945651_fk_workspaces_workspace_id FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT trans_report_category_id_fb4530ac_fk_trans_category_id FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT trans_report_component_id_894b3313_fk_trans_component_id FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT trans_report_creator_id_2b66af65_fk_weblate_auth_user_id FOREIGN KEY (creator_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT trans_report_project_id_026b7e0f_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT trans_report_workspace_id_f914bdb5_fk_workspaces_workspace_id FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.trans_savedsearch
    ADD CONSTRAINT trans_savedsearch_project_id_6559d9bb_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_savedsearch
    ADD CONSTRAINT trans_savedsearch_user_id_8c4baec1_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_suggestion
    ADD CONSTRAINT trans_suggestion_backup_author_claim__f18fec6a_fk_trans_bac FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_suggestion
    ADD CONSTRAINT trans_suggestion_unit_id_eb46fceb_fk_trans_unit_id FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_suggestion
    ADD CONSTRAINT trans_suggestion_user_id_69ce0c2b_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_translation
    ADD CONSTRAINT trans_translation_component_id_c4b70a26_fk_trans_component_id FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_translation
    ADD CONSTRAINT trans_translation_language_id_030f0b30_fk_lang_language_id FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_translation
    ADD CONSTRAINT trans_translation_plural_id_5cf36dc7_fk_lang_plural_id FOREIGN KEY (plural_id) REFERENCES public.lang_plural(id);

ALTER TABLE ONLY public.trans_unit_labels
    ADD CONSTRAINT trans_unit_labels_label_id_4ca92ebd_fk_trans_label_id FOREIGN KEY (label_id) REFERENCES public.trans_label(id);

ALTER TABLE ONLY public.trans_unit_labels
    ADD CONSTRAINT trans_unit_labels_unit_id_a3c2ddc5_fk_trans_unit_id FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT trans_unit_source_unit_id_7a735f87_fk_trans_unit_id FOREIGN KEY (source_unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT trans_unit_translation_id_513bb910_fk_trans_translation_id FOREIGN KEY (translation_id) REFERENCES public.trans_translation(id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT trans_unit_translation_parent_id_bad2a48e_fk_trans_unit_id FOREIGN KEY (translation_parent_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT trans_unit_variant_id_c3315c70_fk_trans_variant_id FOREIGN KEY (variant_id) REFERENCES public.trans_variant(id);

ALTER TABLE ONLY public.trans_variant
    ADD CONSTRAINT trans_variant_component_id_0915cbfc_fk_trans_component_id FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_variant_defining_units
    ADD CONSTRAINT trans_variant_defini_variant_id_20848c13_fk_trans_var FOREIGN KEY (variant_id) REFERENCES public.trans_variant(id);

ALTER TABLE ONLY public.trans_variant_defining_units
    ADD CONSTRAINT trans_variant_defining_units_unit_id_f692977b_fk_trans_unit_id FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_vote
    ADD CONSTRAINT trans_vote_suggestion_id_0df2bda2_fk_trans_suggestion_id FOREIGN KEY (suggestion_id) REFERENCES public.trans_suggestion(id);

ALTER TABLE ONLY public.trans_vote
    ADD CONSTRAINT trans_vote_user_id_86a644fd_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_workflowsetting
    ADD CONSTRAINT trans_workflowsettin_source_language_id_6df86150_fk_lang_lang FOREIGN KEY (source_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_workflowsetting
    ADD CONSTRAINT trans_workflowsetting_language_id_fcb1fa15_fk_lang_language_id FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_workflowsetting
    ADD CONSTRAINT trans_workflowsetting_project_id_32daa6c8_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.vcs_installation
    ADD CONSTRAINT vcs_installation_workspace_id_1a97507d_fk_workspace FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.weblate_auth_autogroup
    ADD CONSTRAINT weblate_auth_autogro_group_id_37b7ff0b_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_admins
    ADD CONSTRAINT weblate_auth_group_a_group_id_fa29ae29_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_admins
    ADD CONSTRAINT weblate_auth_group_a_user_id_f5851c1b_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.weblate_auth_group_components
    ADD CONSTRAINT weblate_auth_group_c_component_id_2b53d012_fk_trans_com FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.weblate_auth_group_componentlists
    ADD CONSTRAINT weblate_auth_group_c_componentlist_id_b59f8fe9_fk_trans_com FOREIGN KEY (componentlist_id) REFERENCES public.trans_componentlist(id);

ALTER TABLE ONLY public.weblate_auth_group_componentlists
    ADD CONSTRAINT weblate_auth_group_c_group_id_7af40eb9_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_components
    ADD CONSTRAINT weblate_auth_group_c_group_id_a07e20aa_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group
    ADD CONSTRAINT weblate_auth_group_defining_project_id_b2ad81a7_fk_trans_pro FOREIGN KEY (defining_project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.weblate_auth_group
    ADD CONSTRAINT weblate_auth_group_defining_workspace_i_84f40bfd_fk_workspace FOREIGN KEY (defining_workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.weblate_auth_group_languages
    ADD CONSTRAINT weblate_auth_group_l_group_id_aa12d0a6_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_languages
    ADD CONSTRAINT weblate_auth_group_l_language_id_4b7f5d81_fk_lang_lang FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.weblate_auth_group_projects
    ADD CONSTRAINT weblate_auth_group_p_group_id_f713a220_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_projects
    ADD CONSTRAINT weblate_auth_group_p_project_id_26b8e2b6_fk_trans_pro FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.weblate_auth_group_roles
    ADD CONSTRAINT weblate_auth_group_r_group_id_ed18f8d0_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_roles
    ADD CONSTRAINT weblate_auth_group_r_role_id_de6300ca_fk_weblate_a FOREIGN KEY (role_id) REFERENCES public.weblate_auth_role(id);

ALTER TABLE ONLY public.weblate_auth_invitation
    ADD CONSTRAINT weblate_auth_invitat_author_id_a713ef92_fk_weblate_a FOREIGN KEY (author_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.weblate_auth_invitation
    ADD CONSTRAINT weblate_auth_invitat_group_id_07a52aff_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_invitation_limit_languages
    ADD CONSTRAINT weblate_auth_invitat_invitation_id_ed3a1e77_fk_weblate_a FOREIGN KEY (invitation_id) REFERENCES public.weblate_auth_invitation(uuid);

ALTER TABLE ONLY public.weblate_auth_invitation_limit_languages
    ADD CONSTRAINT weblate_auth_invitat_language_id_26454a1a_fk_lang_lang FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.weblate_auth_invitation
    ADD CONSTRAINT weblate_auth_invitat_user_id_6a0c6dec_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.weblate_auth_role_permissions
    ADD CONSTRAINT weblate_auth_role_pe_permission_id_d0650152_fk_weblate_a FOREIGN KEY (permission_id) REFERENCES public.weblate_auth_permission(id);

ALTER TABLE ONLY public.weblate_auth_role_permissions
    ADD CONSTRAINT weblate_auth_role_pe_role_id_fa163124_fk_weblate_a FOREIGN KEY (role_id) REFERENCES public.weblate_auth_role(id);

ALTER TABLE ONLY public.weblate_auth_user_groups
    ADD CONSTRAINT weblate_auth_user_gr_group_id_48c0a977_fk_weblate_a FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_user_groups_limit_languages
    ADD CONSTRAINT weblate_auth_user_gr_language_id_52ecbc9e_fk_lang_lang FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.weblate_auth_user_groups_limit_languages
    ADD CONSTRAINT weblate_auth_user_gr_teammembership_id_8da6367c_fk_weblate_a FOREIGN KEY (teammembership_id) REFERENCES public.weblate_auth_user_groups(id);

ALTER TABLE ONLY public.weblate_auth_user_groups
    ADD CONSTRAINT weblate_auth_user_gr_user_id_37e61ff1_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.weblate_auth_userblock
    ADD CONSTRAINT weblate_auth_userblock_project_id_44aae0be_fk_trans_project_id FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.weblate_auth_userblock
    ADD CONSTRAINT weblate_auth_userblock_user_id_06173205_fk_weblate_auth_user_id FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.wladmin_backuplog
    ADD CONSTRAINT wladmin_backuplog_service_id_10292c2d_fk_wladmin_b FOREIGN KEY (service_id) REFERENCES public.wladmin_backupservice(id);

ALTER TABLE ONLY public.wladmin_securityfindingevent
    ADD CONSTRAINT wladmin_securityfind_actor_id_ddbbf0da_fk_weblate_a FOREIGN KEY (actor_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.wladmin_securityfindingevent
    ADD CONSTRAINT wladmin_securityfind_finding_id_fb30cf88_fk_wladmin_s FOREIGN KEY (finding_id) REFERENCES public.wladmin_securityfinding(id);

ALTER TABLE ONLY public.wladmin_securityfindingevent
    ADD CONSTRAINT wladmin_securityfind_scan_id_573ba79c_fk_wladmin_s FOREIGN KEY (scan_id) REFERENCES public.wladmin_securityscan(id);

ALTER TABLE ONLY public.wladmin_securityfinding
    ADD CONSTRAINT wladmin_securityfind_user_id_6c31bc00_fk_weblate_a FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.wladmin_securityscan
    ADD CONSTRAINT wladmin_securityscan_actor_id_0aebe956_fk_weblate_auth_user_id FOREIGN KEY (actor_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.workspaces_workspace
    ADD CONSTRAINT workspaces_workspace_secondary_language_i_b8ff77e6_fk_lang_lang FOREIGN KEY (secondary_language_id) REFERENCES public.lang_language(id);

