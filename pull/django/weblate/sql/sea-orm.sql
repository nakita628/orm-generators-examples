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
    params json NOT NULL,
    address character varying,
    user_agent character varying(200) NOT NULL,
    "timestamp" timestamp with time zone NOT NULL
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
    listing_columns json NOT NULL,
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
    workspace_id uuid,
    actor_id integer,
    event character varying(20) NOT NULL,
    details json NOT NULL,
    "timestamp" timestamp with time zone NOT NULL
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
    configuration json NOT NULL,
    state json NOT NULL,
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
    created timestamp with time zone NOT NULL,
    details json NOT NULL,
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
    created timestamp with time zone NOT NULL
);

CREATE TABLE public.checks_check (
    id integer NOT NULL,
    unit_id integer NOT NULL,
    name character varying(50) NOT NULL,
    dismissed boolean NOT NULL,
    metadata json NOT NULL
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
    value json NOT NULL
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

CREATE TABLE public.django_celery_beat_clockedschedule (
    id integer NOT NULL,
    clocked_time timestamp with time zone NOT NULL
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
    args character varying NOT NULL,
    kwargs character varying NOT NULL,
    queue character varying(200),
    exchange character varying(200),
    routing_key character varying(200),
    headers character varying NOT NULL,
    priority integer,
    expires timestamp with time zone,
    expire_seconds integer,
    one_off boolean NOT NULL,
    start_time timestamp with time zone,
    enabled boolean NOT NULL,
    last_run_at timestamp with time zone,
    total_run_count integer NOT NULL,
    date_changed timestamp with time zone NOT NULL,
    description character varying NOT NULL
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
    ident smallint NOT NULL,
    last_update timestamp with time zone NOT NULL
);

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
    created_at timestamp with time zone,
    last_used_at timestamp with time zone,
    credential_type character varying(32) NOT NULL,
    credential_id bytea NOT NULL,
    public_key bytea NOT NULL,
    transports json NOT NULL,
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

CREATE TABLE public.django_session (
    session_key character varying(40) NOT NULL,
    session_data character varying NOT NULL,
    expire_date timestamp with time zone NOT NULL
);

CREATE TABLE public.fonts_font (
    id integer NOT NULL,
    family character varying(100) NOT NULL,
    style character varying(100) NOT NULL,
    font character varying(100) NOT NULL,
    project_id integer NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
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
    unreferenced timestamp with time zone
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
    metadata json NOT NULL,
    pending_manifest json NOT NULL,
    digest character varying(64) NOT NULL,
    created timestamp with time zone NOT NULL,
    retired timestamp with time zone,
    status character varying(16) NOT NULL,
    error character varying NOT NULL,
    published timestamp with time zone
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
    path character varying NOT NULL,
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
    formula character varying NOT NULL,
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
    "timestamp" timestamp with time zone NOT NULL,
    error character varying NOT NULL
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
    workspace_id uuid,
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
    configuration json NOT NULL,
    enabled boolean NOT NULL,
    workspace_id uuid,
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
    source character varying NOT NULL,
    target character varying NOT NULL,
    origin character varying NOT NULL,
    context character varying NOT NULL,
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
    workspace_id uuid,
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
    updated timestamp with time zone NOT NULL
);

CREATE TABLE public.metrics_metric (
    id bigint NOT NULL,
    date date NOT NULL,
    scope smallint NOT NULL,
    relation integer NOT NULL,
    secondary integer NOT NULL,
    changes integer NOT NULL,
    data json
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
    throttling_failure_timestamp timestamp with time zone,
    throttling_failure_count integer NOT NULL,
    created_at timestamp with time zone,
    last_used_at timestamp with time zone
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
    throttling_failure_timestamp timestamp with time zone,
    throttling_failure_count integer NOT NULL,
    created_at timestamp with time zone,
    last_used_at timestamp with time zone,
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
    "timestamp" timestamp with time zone NOT NULL,
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
    signature json NOT NULL,
    generation uuid NOT NULL,
    data json,
    task_id uuid,
    requested_at timestamp with time zone,
    error character varying(16) NOT NULL
);

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
    "timestamp" timestamp with time zone NOT NULL
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
    data json NOT NULL,
    "timestamp" timestamp with time zone NOT NULL
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
    extra_data json NOT NULL,
    created timestamp with time zone NOT NULL,
    modified timestamp with time zone NOT NULL
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
    "timestamp" timestamp with time zone NOT NULL,
    updated timestamp with time zone NOT NULL,
    name character varying(150) NOT NULL,
    dismissed_at timestamp with time zone,
    dismissed_by_id integer,
    dismissal_reason character varying(500) NOT NULL,
    dismissal_fingerprint character varying(64) NOT NULL,
    severity smallint NOT NULL,
    details json NOT NULL
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
    message character varying NOT NULL,
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
    invitation_id uuid,
    accepted_at timestamp with time zone
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
    check_flags character varying NOT NULL,
    enforced_checks json NOT NULL,
    inherit_enforced_checks boolean NOT NULL,
    license character varying(150) NOT NULL,
    inherit_license boolean NOT NULL,
    agreement character varying NOT NULL,
    inherit_agreement boolean NOT NULL,
    new_lang character varying(10) NOT NULL,
    inherit_new_lang boolean NOT NULL,
    language_code_style character varying(20) NOT NULL,
    inherit_language_code_style boolean NOT NULL,
    secondary_language_id integer,
    inherit_secondary_language boolean NOT NULL,
    commit_message character varying NOT NULL,
    inherit_commit_message boolean NOT NULL,
    add_message character varying NOT NULL,
    inherit_add_message boolean NOT NULL,
    delete_message character varying NOT NULL,
    inherit_delete_message boolean NOT NULL,
    merge_message character varying NOT NULL,
    inherit_merge_message boolean NOT NULL,
    addon_message character varying NOT NULL,
    inherit_addon_message boolean NOT NULL,
    pull_message character varying NOT NULL,
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
    workspace_id uuid,
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
    "timestamp" timestamp with time zone NOT NULL,
    action integer NOT NULL,
    target character varying NOT NULL,
    old character varying NOT NULL,
    details json NOT NULL
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
    comment character varying NOT NULL,
    user_id integer,
    "timestamp" timestamp with time zone NOT NULL,
    resolved boolean NOT NULL,
    userdetails json NOT NULL
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
    vcs_params json NOT NULL,
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
    file_format_params json NOT NULL,
    locked boolean NOT NULL,
    hide_glossary_matches boolean NOT NULL,
    allow_translation_propagation boolean NOT NULL,
    contribute_project_tm boolean NOT NULL,
    enable_suggestions boolean NOT NULL,
    suggestion_voting boolean NOT NULL,
    suggestion_autoaccept smallint NOT NULL,
    check_flags character varying NOT NULL,
    enforced_checks json NOT NULL,
    inherit_enforced_checks boolean NOT NULL,
    license character varying(150) NOT NULL,
    inherit_license boolean NOT NULL,
    agreement character varying NOT NULL,
    inherit_agreement boolean NOT NULL,
    new_lang character varying(10) NOT NULL,
    inherit_new_lang boolean NOT NULL,
    language_code_style character varying(20) NOT NULL,
    inherit_language_code_style boolean NOT NULL,
    manage_units boolean NOT NULL,
    merge_style character varying(10) NOT NULL,
    commit_message character varying NOT NULL,
    inherit_commit_message boolean NOT NULL,
    add_message character varying NOT NULL,
    inherit_add_message boolean NOT NULL,
    delete_message character varying NOT NULL,
    inherit_delete_message boolean NOT NULL,
    merge_message character varying NOT NULL,
    inherit_merge_message boolean NOT NULL,
    addon_message character varying NOT NULL,
    inherit_addon_message boolean NOT NULL,
    pull_message character varying NOT NULL,
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
    workspace_id uuid,
    "timestamp" timestamp with time zone NOT NULL
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
    target character varying NOT NULL,
    explanation character varying NOT NULL,
    source_unit_explanation character varying NOT NULL,
    state integer NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
    add_unit boolean NOT NULL,
    automatically_translated boolean NOT NULL,
    metadata json NOT NULL
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
    instructions character varying NOT NULL,
    workspace_id uuid,
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
    language_aliases character varying NOT NULL,
    secondary_language_id integer,
    check_flags character varying NOT NULL,
    enforced_checks json NOT NULL,
    inherit_enforced_checks boolean NOT NULL,
    license character varying(150) NOT NULL,
    inherit_license boolean NOT NULL,
    agreement character varying NOT NULL,
    inherit_agreement boolean NOT NULL,
    new_lang character varying(10) NOT NULL,
    inherit_new_lang boolean NOT NULL,
    language_code_style character varying(20) NOT NULL,
    inherit_language_code_style boolean NOT NULL,
    inherit_secondary_language boolean NOT NULL,
    commit_message character varying NOT NULL,
    inherit_commit_message boolean NOT NULL,
    add_message character varying NOT NULL,
    inherit_add_message boolean NOT NULL,
    delete_message character varying NOT NULL,
    inherit_delete_message boolean NOT NULL,
    merge_message character varying NOT NULL,
    inherit_merge_message boolean NOT NULL,
    addon_message character varying NOT NULL,
    inherit_addon_message boolean NOT NULL,
    pull_message character varying NOT NULL,
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
    created timestamp with time zone NOT NULL,
    kind character varying(32) NOT NULL,
    version smallint NOT NULL,
    parameters json NOT NULL,
    data json NOT NULL,
    workspace_id uuid,
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
    query character varying NOT NULL,
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
    target character varying NOT NULL,
    user_id integer,
    userdetails json NOT NULL,
    "timestamp" timestamp with time zone NOT NULL
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
    check_flags character varying NOT NULL
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
    location character varying NOT NULL,
    context character varying NOT NULL,
    note character varying NOT NULL,
    flags character varying NOT NULL,
    source character varying NOT NULL,
    previous_source character varying NOT NULL,
    target character varying NOT NULL,
    state integer NOT NULL,
    original_state integer NOT NULL,
    details json NOT NULL,
    "position" integer NOT NULL,
    num_words integer NOT NULL,
    priority integer NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
    last_updated timestamp with time zone NOT NULL,
    extra_flags character varying NOT NULL,
    explanation character varying NOT NULL,
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
    key character varying NOT NULL
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
    private_key character varying NOT NULL,
    client_id character varying(255) NOT NULL,
    client_secret character varying(255) NOT NULL,
    webhook_secret character varying(255) NOT NULL,
    webhook_token uuid NOT NULL,
    html_url character varying(200) NOT NULL,
    created timestamp with time zone NOT NULL,
    updated timestamp with time zone NOT NULL
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
    workspace_id uuid NOT NULL,
    hostname character varying(255) NOT NULL,
    enabled boolean NOT NULL,
    created timestamp with time zone NOT NULL,
    repositories json NOT NULL,
    repositories_updated timestamp with time zone
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
    payload json NOT NULL,
    created timestamp with time zone NOT NULL,
    updated timestamp with time zone NOT NULL
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
    defining_workspace_id uuid,
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
    uuid uuid NOT NULL,
    backup_authorship boolean NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
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
    invitation_id uuid NOT NULL,
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
    last_login timestamp with time zone,
    username character varying(150) NOT NULL,
    full_name character varying(150) NOT NULL,
    email character varying(190),
    is_superuser boolean NOT NULL,
    is_active boolean NOT NULL,
    is_bot boolean NOT NULL,
    date_expires timestamp with time zone,
    date_joined timestamp with time zone NOT NULL
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
    expiry timestamp with time zone,
    note character varying NOT NULL
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
    "timestamp" timestamp with time zone NOT NULL,
    event character varying(100) NOT NULL,
    warning boolean NOT NULL,
    log character varying NOT NULL
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
    "timestamp" timestamp with time zone NOT NULL,
    passphrase character varying(100) NOT NULL,
    paperkey character varying NOT NULL
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
    message character varying NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
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
    evidence json NOT NULL,
    status character varying(10) NOT NULL,
    first_seen timestamp with time zone NOT NULL,
    last_seen timestamp with time zone NOT NULL,
    handled_at timestamp with time zone
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
    scan_id uuid NOT NULL,
    actor_id integer,
    action character varying(20) NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
    reason character varying(50) NOT NULL,
    evidence json NOT NULL
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
    id uuid NOT NULL,
    actor_id integer,
    filename character varying(255) NOT NULL,
    operation character varying(20) NOT NULL,
    finding_type character varying(50) NOT NULL,
    state character varying(20) NOT NULL,
    created timestamp with time zone NOT NULL,
    completed timestamp with time zone,
    processed integer NOT NULL,
    total integer NOT NULL,
    summary json NOT NULL,
    error character varying NOT NULL
);

CREATE TABLE public.wladmin_supportstatus (
    id integer NOT NULL,
    name character varying(150) NOT NULL,
    secret character varying(400) NOT NULL,
    expiry timestamp with time zone,
    in_limits boolean NOT NULL,
    discoverable boolean NOT NULL,
    limits json NOT NULL,
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
    id uuid NOT NULL,
    metric_id integer NOT NULL,
    name character varying(100) NOT NULL,
    propagation_scope character varying(20) NOT NULL,
    name_managed boolean NOT NULL,
    use_workspace_tm boolean NOT NULL,
    contribute_workspace_tm boolean NOT NULL,
    license character varying(150) NOT NULL,
    agreement character varying NOT NULL,
    new_lang character varying(10) NOT NULL,
    language_code_style character varying(20) NOT NULL,
    secondary_language_id integer,
    check_flags character varying NOT NULL,
    enforced_checks json NOT NULL,
    inherit_enforced_checks boolean NOT NULL,
    commit_message character varying NOT NULL,
    add_message character varying NOT NULL,
    delete_message character varying NOT NULL,
    merge_message character varying NOT NULL,
    addon_message character varying NOT NULL,
    pull_message character varying NOT NULL
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

ALTER TABLE ONLY public.django_celery_beat_solarschedule ALTER COLUMN id SET DEFAULT nextval('public.django_celery_beat_solarschedule_id_seq'::regclass);

ALTER TABLE ONLY public.django_content_type ALTER COLUMN id SET DEFAULT nextval('public.django_content_type_id_seq'::regclass);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnattestation ALTER COLUMN id SET DEFAULT nextval('public.django_otp_webauthn_webauthnattestation_id_seq'::regclass);

ALTER TABLE ONLY public.django_otp_webauthn_webauthncredential ALTER COLUMN id SET DEFAULT nextval('public.django_otp_webauthn_webauthncredential_id_seq'::regclass);

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

ALTER TABLE ONLY public.accounts_profile
    ADD CONSTRAINT accounts_profile_user_id_key UNIQUE (user_id);

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

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_name_key UNIQUE (name);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.authtoken_token
    ADD CONSTRAINT authtoken_token_pkey PRIMARY KEY (key);

ALTER TABLE ONLY public.authtoken_token
    ADD CONSTRAINT authtoken_token_user_id_key UNIQUE (user_id);

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
    ADD CONSTRAINT django_celery_beat_periodictask_name_key UNIQUE (name);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT django_celery_beat_periodictask_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_celery_beat_periodictasks
    ADD CONSTRAINT django_celery_beat_periodictasks_pkey PRIMARY KEY (ident);

ALTER TABLE ONLY public.django_celery_beat_solarschedule
    ADD CONSTRAINT django_celery_beat_solarschedule_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnattestation
    ADD CONSTRAINT django_otp_webauthn_webauthnattestation_credential_id_key UNIQUE (credential_id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnattestation
    ADD CONSTRAINT django_otp_webauthn_webauthnattestation_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthncredential
    ADD CONSTRAINT django_otp_webauthn_webauthncredential_credential_id_sha256_key UNIQUE (credential_id_sha256);

ALTER TABLE ONLY public.django_otp_webauthn_webauthncredential
    ADD CONSTRAINT django_otp_webauthn_webauthncredential_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnuserhandle
    ADD CONSTRAINT django_otp_webauthn_webauthnuserhandle_handle_hex_key UNIQUE (handle_hex);

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
    ADD CONSTRAINT kotlin_sdk_kotlinsdkcleanup_build_id_key UNIQUE (build_id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkcleanup
    ADD CONSTRAINT kotlin_sdk_kotlinsdkcleanup_path_key UNIQUE (path);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkcleanup
    ADD CONSTRAINT kotlin_sdk_kotlinsdkcleanup_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.lang_language
    ADD CONSTRAINT lang_language_code_key UNIQUE (code);

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
    ADD CONSTRAINT trans_backupauthorclaim_invitation_id_key UNIQUE (invitation_id);

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
    ADD CONSTRAINT trans_componentlist_name_key UNIQUE (name);

ALTER TABLE ONLY public.trans_componentlist
    ADD CONSTRAINT trans_componentlist_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_componentlist
    ADD CONSTRAINT trans_componentlist_slug_key UNIQUE (slug);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT trans_contributoragreement_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_label
    ADD CONSTRAINT trans_label_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_pendingunitchange
    ADD CONSTRAINT trans_pendingunitchange_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_project
    ADD CONSTRAINT trans_project_name_key UNIQUE (name);

ALTER TABLE ONLY public.trans_project
    ADD CONSTRAINT trans_project_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trans_project
    ADD CONSTRAINT trans_project_slug_key UNIQUE (slug);

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
    ADD CONSTRAINT vcs_githubappcredentials_hostname_key UNIQUE (hostname);

ALTER TABLE ONLY public.vcs_githubappcredentials
    ADD CONSTRAINT vcs_githubappcredentials_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.vcs_githubappcredentials
    ADD CONSTRAINT vcs_githubappcredentials_webhook_token_key UNIQUE (webhook_token);

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
    ADD CONSTRAINT weblate_auth_permission_codename_key UNIQUE (codename);

ALTER TABLE ONLY public.weblate_auth_permission
    ADD CONSTRAINT weblate_auth_permission_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_role
    ADD CONSTRAINT weblate_auth_role_name_key UNIQUE (name);

ALTER TABLE ONLY public.weblate_auth_role_permissions
    ADD CONSTRAINT weblate_auth_role_permissions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_role
    ADD CONSTRAINT weblate_auth_role_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_user
    ADD CONSTRAINT weblate_auth_user_email_key UNIQUE (email);

ALTER TABLE ONLY public.weblate_auth_user_groups_limit_languages
    ADD CONSTRAINT weblate_auth_user_groups_limit_languages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_user_groups
    ADD CONSTRAINT weblate_auth_user_groups_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_user
    ADD CONSTRAINT weblate_auth_user_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.weblate_auth_user
    ADD CONSTRAINT weblate_auth_user_username_key UNIQUE (username);

ALTER TABLE ONLY public.weblate_auth_userblock
    ADD CONSTRAINT weblate_auth_userblock_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_backuplog
    ADD CONSTRAINT wladmin_backuplog_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_backupservice
    ADD CONSTRAINT wladmin_backupservice_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wladmin_configurationerror
    ADD CONSTRAINT wladmin_configurationerror_name_key UNIQUE (name);

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
    ADD CONSTRAINT workspaces_workspace_metric_id_key UNIQUE (metric_id);

ALTER TABLE ONLY public.workspaces_workspace
    ADD CONSTRAINT workspaces_workspace_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts_auditlog
    ADD CONSTRAINT "fk-accounts_auditlog-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.accounts_profile
    ADD CONSTRAINT "fk-accounts_profile-dashboard_component_list_id" FOREIGN KEY (dashboard_component_list_id) REFERENCES public.trans_componentlist(id);

ALTER TABLE ONLY public.accounts_profile
    ADD CONSTRAINT "fk-accounts_profile-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.accounts_profile_languages
    ADD CONSTRAINT "fk-accounts_profile_languages-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.accounts_profile_languages
    ADD CONSTRAINT "fk-accounts_profile_languages-profile_id" FOREIGN KEY (profile_id) REFERENCES public.accounts_profile(id);

ALTER TABLE ONLY public.accounts_profile_secondary_languages
    ADD CONSTRAINT "fk-accounts_profile_secondary_languages-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.accounts_profile_secondary_languages
    ADD CONSTRAINT "fk-accounts_profile_secondary_languages-profile_id" FOREIGN KEY (profile_id) REFERENCES public.accounts_profile(id);

ALTER TABLE ONLY public.accounts_profile_watched
    ADD CONSTRAINT "fk-accounts_profile_watched-profile_id" FOREIGN KEY (profile_id) REFERENCES public.accounts_profile(id);

ALTER TABLE ONLY public.accounts_profile_watched
    ADD CONSTRAINT "fk-accounts_profile_watched-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.accounts_subscription
    ADD CONSTRAINT "fk-accounts_subscription-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.accounts_subscription
    ADD CONSTRAINT "fk-accounts_subscription-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.accounts_subscription
    ADD CONSTRAINT "fk-accounts_subscription-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT "fk-accounts_teamauditlog-actor_id" FOREIGN KEY (actor_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT "fk-accounts_teamauditlog-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT "fk-accounts_teamauditlog-team_id" FOREIGN KEY (team_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.accounts_teamauditlog
    ADD CONSTRAINT "fk-accounts_teamauditlog-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.accounts_verifiedemail
    ADD CONSTRAINT "fk-accounts_verifiedemail-social_id" FOREIGN KEY (social_id) REFERENCES public.social_auth_usersocialauth(id);

ALTER TABLE ONLY public.addons_addon
    ADD CONSTRAINT "fk-addons_addon-category_id" FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.addons_addon
    ADD CONSTRAINT "fk-addons_addon-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.addons_addon
    ADD CONSTRAINT "fk-addons_addon-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.addons_addonactivitylog
    ADD CONSTRAINT "fk-addons_addonactivitylog-addon_id" FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.addons_addonactivitylog
    ADD CONSTRAINT "fk-addons_addonactivitylog-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.addons_event
    ADD CONSTRAINT "fk-addons_event-addon_id" FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT "fk-auth_group_permissions-group_id" FOREIGN KEY (group_id) REFERENCES public.auth_group(id);

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT "fk-auth_group_permissions-permission_id" FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id);

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT "fk-auth_permission-content_type_id" FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.authtoken_token
    ADD CONSTRAINT "fk-authtoken_token-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.checks_check
    ADD CONSTRAINT "fk-checks_check-unit_id" FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT "fk-django_admin_log-content_type_id" FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id);

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT "fk-django_admin_log-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT "fk-django_celery_beat_periodictask-clocked_id" FOREIGN KEY (clocked_id) REFERENCES public.django_celery_beat_clockedschedule(id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT "fk-django_celery_beat_periodictask-crontab_id" FOREIGN KEY (crontab_id) REFERENCES public.django_celery_beat_crontabschedule(id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT "fk-django_celery_beat_periodictask-interval_id" FOREIGN KEY (interval_id) REFERENCES public.django_celery_beat_intervalschedule(id);

ALTER TABLE ONLY public.django_celery_beat_periodictask
    ADD CONSTRAINT "fk-django_celery_beat_periodictask-solar_id" FOREIGN KEY (solar_id) REFERENCES public.django_celery_beat_solarschedule(id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnattestation
    ADD CONSTRAINT "fk-django_otp_webauthn_webauthnattestation-credential_id" FOREIGN KEY (credential_id) REFERENCES public.django_otp_webauthn_webauthncredential(id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthncredential
    ADD CONSTRAINT "fk-django_otp_webauthn_webauthncredential-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.django_otp_webauthn_webauthnuserhandle
    ADD CONSTRAINT "fk-django_otp_webauthn_webauthnuserhandle-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.fonts_font
    ADD CONSTRAINT "fk-fonts_font-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.fonts_font
    ADD CONSTRAINT "fk-fonts_font-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.fonts_fontgroup
    ADD CONSTRAINT "fk-fonts_fontgroup-font_id" FOREIGN KEY (font_id) REFERENCES public.fonts_font(id);

ALTER TABLE ONLY public.fonts_fontgroup
    ADD CONSTRAINT "fk-fonts_fontgroup-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.fonts_fontoverride
    ADD CONSTRAINT "fk-fonts_fontoverride-font_id" FOREIGN KEY (font_id) REFERENCES public.fonts_font(id);

ALTER TABLE ONLY public.fonts_fontoverride
    ADD CONSTRAINT "fk-fonts_fontoverride-group_id" FOREIGN KEY (group_id) REFERENCES public.fonts_fontgroup(id);

ALTER TABLE ONLY public.fonts_fontoverride
    ADD CONSTRAINT "fk-fonts_fontoverride-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact
    ADD CONSTRAINT "fk-kotlin_sdk_kotlinsdkartifact-addon_id" FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact_builds
    ADD CONSTRAINT "fk-kotlin_sdk_kotlinsdkartifact_builds-kotlinsdkartifact_id" FOREIGN KEY (kotlinsdkartifact_id) REFERENCES public.kotlin_sdk_kotlinsdkartifact(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkartifact_builds
    ADD CONSTRAINT "fk-kotlin_sdk_kotlinsdkartifact_builds-kotlinsdkbuild_id" FOREIGN KEY (kotlinsdkbuild_id) REFERENCES public.kotlin_sdk_kotlinsdkbuild(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkbuild
    ADD CONSTRAINT "fk-kotlin_sdk_kotlinsdkbuild-addon_id" FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.kotlin_sdk_kotlinsdkcleanup
    ADD CONSTRAINT "fk-kotlin_sdk_kotlinsdkcleanup-build_id" FOREIGN KEY (build_id) REFERENCES public.kotlin_sdk_kotlinsdkbuild(id);

ALTER TABLE ONLY public.lang_plural
    ADD CONSTRAINT "fk-lang_plural-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.machinery_machineryerror
    ADD CONSTRAINT "fk-machinery_machineryerror-installation_id" FOREIGN KEY (installation_id) REFERENCES public.machinery_machineryinstallation(id);

ALTER TABLE ONLY public.machinery_machineryexclusion
    ADD CONSTRAINT "fk-machinery_machineryexclusion-installation_id" FOREIGN KEY (installation_id) REFERENCES public.machinery_machineryinstallation(id);

ALTER TABLE ONLY public.machinery_machineryexclusion
    ADD CONSTRAINT "fk-machinery_machineryexclusion-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.machinery_machineryexclusion
    ADD CONSTRAINT "fk-machinery_machineryexclusion-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.machinery_machineryinstallation
    ADD CONSTRAINT "fk-machinery_machineryinstallation-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.machinery_machineryinstallation
    ADD CONSTRAINT "fk-machinery_machineryinstallation-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT "fk-memory_memory-legacy_project_id" FOREIGN KEY (legacy_project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT "fk-memory_memory-legacy_user_id" FOREIGN KEY (legacy_user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT "fk-memory_memory-source_language_id" FOREIGN KEY (source_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.memory_memory
    ADD CONSTRAINT "fk-memory_memory-target_language_id" FOREIGN KEY (target_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT "fk-memory_memoryscope-memory_id" FOREIGN KEY (memory_id) REFERENCES public.memory_memory(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT "fk-memory_memoryscope-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT "fk-memory_memoryscope-source_component_id" FOREIGN KEY (source_component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT "fk-memory_memoryscope-source_project_id" FOREIGN KEY (source_project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT "fk-memory_memoryscope-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.memory_memoryscope
    ADD CONSTRAINT "fk-memory_memoryscope-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.otp_static_staticdevice
    ADD CONSTRAINT "fk-otp_static_staticdevice-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.otp_static_statictoken
    ADD CONSTRAINT "fk-otp_static_statictoken-device_id" FOREIGN KEY (device_id) REFERENCES public.otp_static_staticdevice(id);

ALTER TABLE ONLY public.otp_totp_totpdevice
    ADD CONSTRAINT "fk-otp_totp_totpdevice-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.screenshots_screenshot
    ADD CONSTRAINT "fk-screenshots_screenshot-backup_author_claim_id" FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.screenshots_screenshot
    ADD CONSTRAINT "fk-screenshots_screenshot-translation_id" FOREIGN KEY (translation_id) REFERENCES public.trans_translation(id);

ALTER TABLE ONLY public.screenshots_screenshot
    ADD CONSTRAINT "fk-screenshots_screenshot-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.screenshots_screenshot_units
    ADD CONSTRAINT "fk-screenshots_screenshot_units-screenshot_id" FOREIGN KEY (screenshot_id) REFERENCES public.screenshots_screenshot(id);

ALTER TABLE ONLY public.screenshots_screenshot_units
    ADD CONSTRAINT "fk-screenshots_screenshot_units-unit_id" FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.screenshots_screenshotocr
    ADD CONSTRAINT "fk-screenshots_screenshotocr-screenshot_id" FOREIGN KEY (screenshot_id) REFERENCES public.screenshots_screenshot(id);

ALTER TABLE ONLY public.social_auth_usersocialauth
    ADD CONSTRAINT "fk-social_auth_usersocialauth-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_alert
    ADD CONSTRAINT "fk-trans_alert-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_alert
    ADD CONSTRAINT "fk-trans_alert-dismissed_by_id" FOREIGN KEY (dismissed_by_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT "fk-trans_announcement-category_id" FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT "fk-trans_announcement-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT "fk-trans_announcement-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_announcement
    ADD CONSTRAINT "fk-trans_announcement-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_autocomponentlist
    ADD CONSTRAINT "fk-trans_autocomponentlist-componentlist_id" FOREIGN KEY (componentlist_id) REFERENCES public.trans_componentlist(id);

ALTER TABLE ONLY public.trans_backupauthorclaim
    ADD CONSTRAINT "fk-trans_backupauthorclaim-invitation_id" FOREIGN KEY (invitation_id) REFERENCES public.weblate_auth_invitation(uuid);

ALTER TABLE ONLY public.trans_backupauthorclaim
    ADD CONSTRAINT "fk-trans_backupauthorclaim-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_backupauthorclaim
    ADD CONSTRAINT "fk-trans_backupauthorclaim-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_backupauthorvote
    ADD CONSTRAINT "fk-trans_backupauthorvote-claim_id" FOREIGN KEY (claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_backupauthorvote
    ADD CONSTRAINT "fk-trans_backupauthorvote-placeholder_id" FOREIGN KEY (placeholder_id) REFERENCES public.trans_vote(id);

ALTER TABLE ONLY public.trans_backupauthorvote
    ADD CONSTRAINT "fk-trans_backupauthorvote-suggestion_id" FOREIGN KEY (suggestion_id) REFERENCES public.trans_suggestion(id);

ALTER TABLE ONLY public.trans_category
    ADD CONSTRAINT "fk-trans_category-category_id" FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_category
    ADD CONSTRAINT "fk-trans_category-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_category
    ADD CONSTRAINT "fk-trans_category-secondary_language_id" FOREIGN KEY (secondary_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-addon_id" FOREIGN KEY (addon_id) REFERENCES public.addons_addon(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-alert_id" FOREIGN KEY (alert_id) REFERENCES public.trans_alert(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-announcement_id" FOREIGN KEY (announcement_id) REFERENCES public.trans_announcement(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-author_id" FOREIGN KEY (author_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-backup_author_claim_id" FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-category_id" FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-comment_id" FOREIGN KEY (comment_id) REFERENCES public.trans_comment(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-screenshot_id" FOREIGN KEY (screenshot_id) REFERENCES public.screenshots_screenshot(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-suggestion_id" FOREIGN KEY (suggestion_id) REFERENCES public.trans_suggestion(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-translation_id" FOREIGN KEY (translation_id) REFERENCES public.trans_translation(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-unit_id" FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_change
    ADD CONSTRAINT "fk-trans_change-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.trans_comment
    ADD CONSTRAINT "fk-trans_comment-backup_author_claim_id" FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_comment
    ADD CONSTRAINT "fk-trans_comment-unit_id" FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_comment
    ADD CONSTRAINT "fk-trans_comment-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT "fk-trans_component-category_id" FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT "fk-trans_component-linked_component_id" FOREIGN KEY (linked_component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT "fk-trans_component-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT "fk-trans_component-secondary_language_id" FOREIGN KEY (secondary_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_component
    ADD CONSTRAINT "fk-trans_component-source_language_id" FOREIGN KEY (source_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_component_links
    ADD CONSTRAINT "fk-trans_component_links-category_id" FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_component_links
    ADD CONSTRAINT "fk-trans_component_links-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_component_links
    ADD CONSTRAINT "fk-trans_component_links-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_componentlist_components
    ADD CONSTRAINT "fk-trans_componentlist_components-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_componentlist_components
    ADD CONSTRAINT "fk-trans_componentlist_components-componentlist_id" FOREIGN KEY (componentlist_id) REFERENCES public.trans_componentlist(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT "fk-trans_contributoragreement-category_id" FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT "fk-trans_contributoragreement-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT "fk-trans_contributoragreement-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT "fk-trans_contributoragreement-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_contributoragreement
    ADD CONSTRAINT "fk-trans_contributoragreement-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.trans_label
    ADD CONSTRAINT "fk-trans_label-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_pendingunitchange
    ADD CONSTRAINT "fk-trans_pendingunitchange-author_id" FOREIGN KEY (author_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_pendingunitchange
    ADD CONSTRAINT "fk-trans_pendingunitchange-backup_author_claim_id" FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_pendingunitchange
    ADD CONSTRAINT "fk-trans_pendingunitchange-unit_id" FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_project
    ADD CONSTRAINT "fk-trans_project-secondary_language_id" FOREIGN KEY (secondary_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_project
    ADD CONSTRAINT "fk-trans_project-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT "fk-trans_report-category_id" FOREIGN KEY (category_id) REFERENCES public.trans_category(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT "fk-trans_report-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT "fk-trans_report-creator_id" FOREIGN KEY (creator_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT "fk-trans_report-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_report
    ADD CONSTRAINT "fk-trans_report-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.trans_savedsearch
    ADD CONSTRAINT "fk-trans_savedsearch-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_savedsearch
    ADD CONSTRAINT "fk-trans_savedsearch-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_suggestion
    ADD CONSTRAINT "fk-trans_suggestion-backup_author_claim_id" FOREIGN KEY (backup_author_claim_id) REFERENCES public.trans_backupauthorclaim(id);

ALTER TABLE ONLY public.trans_suggestion
    ADD CONSTRAINT "fk-trans_suggestion-unit_id" FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_suggestion
    ADD CONSTRAINT "fk-trans_suggestion-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_translation
    ADD CONSTRAINT "fk-trans_translation-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_translation
    ADD CONSTRAINT "fk-trans_translation-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_translation
    ADD CONSTRAINT "fk-trans_translation-plural_id" FOREIGN KEY (plural_id) REFERENCES public.lang_plural(id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT "fk-trans_unit-source_unit_id" FOREIGN KEY (source_unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT "fk-trans_unit-translation_id" FOREIGN KEY (translation_id) REFERENCES public.trans_translation(id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT "fk-trans_unit-translation_parent_id" FOREIGN KEY (translation_parent_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_unit
    ADD CONSTRAINT "fk-trans_unit-variant_id" FOREIGN KEY (variant_id) REFERENCES public.trans_variant(id);

ALTER TABLE ONLY public.trans_unit_labels
    ADD CONSTRAINT "fk-trans_unit_labels-label_id" FOREIGN KEY (label_id) REFERENCES public.trans_label(id);

ALTER TABLE ONLY public.trans_unit_labels
    ADD CONSTRAINT "fk-trans_unit_labels-unit_id" FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_variant
    ADD CONSTRAINT "fk-trans_variant-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.trans_variant_defining_units
    ADD CONSTRAINT "fk-trans_variant_defining_units-unit_id" FOREIGN KEY (unit_id) REFERENCES public.trans_unit(id);

ALTER TABLE ONLY public.trans_variant_defining_units
    ADD CONSTRAINT "fk-trans_variant_defining_units-variant_id" FOREIGN KEY (variant_id) REFERENCES public.trans_variant(id);

ALTER TABLE ONLY public.trans_vote
    ADD CONSTRAINT "fk-trans_vote-suggestion_id" FOREIGN KEY (suggestion_id) REFERENCES public.trans_suggestion(id);

ALTER TABLE ONLY public.trans_vote
    ADD CONSTRAINT "fk-trans_vote-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.trans_workflowsetting
    ADD CONSTRAINT "fk-trans_workflowsetting-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.trans_workflowsetting
    ADD CONSTRAINT "fk-trans_workflowsetting-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.trans_workflowsetting
    ADD CONSTRAINT "fk-trans_workflowsetting-source_language_id" FOREIGN KEY (source_language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.vcs_installation
    ADD CONSTRAINT "fk-vcs_installation-workspace_id" FOREIGN KEY (workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.weblate_auth_autogroup
    ADD CONSTRAINT "fk-weblate_auth_autogroup-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group
    ADD CONSTRAINT "fk-weblate_auth_group-defining_project_id" FOREIGN KEY (defining_project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.weblate_auth_group
    ADD CONSTRAINT "fk-weblate_auth_group-defining_workspace_id" FOREIGN KEY (defining_workspace_id) REFERENCES public.workspaces_workspace(id);

ALTER TABLE ONLY public.weblate_auth_group_admins
    ADD CONSTRAINT "fk-weblate_auth_group_admins-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_admins
    ADD CONSTRAINT "fk-weblate_auth_group_admins-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.weblate_auth_group_componentlists
    ADD CONSTRAINT "fk-weblate_auth_group_componentlists-componentlist_id" FOREIGN KEY (componentlist_id) REFERENCES public.trans_componentlist(id);

ALTER TABLE ONLY public.weblate_auth_group_componentlists
    ADD CONSTRAINT "fk-weblate_auth_group_componentlists-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_components
    ADD CONSTRAINT "fk-weblate_auth_group_components-component_id" FOREIGN KEY (component_id) REFERENCES public.trans_component(id);

ALTER TABLE ONLY public.weblate_auth_group_components
    ADD CONSTRAINT "fk-weblate_auth_group_components-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_languages
    ADD CONSTRAINT "fk-weblate_auth_group_languages-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_languages
    ADD CONSTRAINT "fk-weblate_auth_group_languages-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.weblate_auth_group_projects
    ADD CONSTRAINT "fk-weblate_auth_group_projects-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_projects
    ADD CONSTRAINT "fk-weblate_auth_group_projects-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.weblate_auth_group_roles
    ADD CONSTRAINT "fk-weblate_auth_group_roles-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_group_roles
    ADD CONSTRAINT "fk-weblate_auth_group_roles-role_id" FOREIGN KEY (role_id) REFERENCES public.weblate_auth_role(id);

ALTER TABLE ONLY public.weblate_auth_invitation
    ADD CONSTRAINT "fk-weblate_auth_invitation-author_id" FOREIGN KEY (author_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.weblate_auth_invitation
    ADD CONSTRAINT "fk-weblate_auth_invitation-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_invitation
    ADD CONSTRAINT "fk-weblate_auth_invitation-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.weblate_auth_invitation_limit_languages
    ADD CONSTRAINT "fk-weblate_auth_invitation_limit_languages-invitation_id" FOREIGN KEY (invitation_id) REFERENCES public.weblate_auth_invitation(uuid);

ALTER TABLE ONLY public.weblate_auth_invitation_limit_languages
    ADD CONSTRAINT "fk-weblate_auth_invitation_limit_languages-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.weblate_auth_role_permissions
    ADD CONSTRAINT "fk-weblate_auth_role_permissions-permission_id" FOREIGN KEY (permission_id) REFERENCES public.weblate_auth_permission(id);

ALTER TABLE ONLY public.weblate_auth_role_permissions
    ADD CONSTRAINT "fk-weblate_auth_role_permissions-role_id" FOREIGN KEY (role_id) REFERENCES public.weblate_auth_role(id);

ALTER TABLE ONLY public.weblate_auth_user_groups
    ADD CONSTRAINT "fk-weblate_auth_user_groups-group_id" FOREIGN KEY (group_id) REFERENCES public.weblate_auth_group(id);

ALTER TABLE ONLY public.weblate_auth_user_groups
    ADD CONSTRAINT "fk-weblate_auth_user_groups-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.weblate_auth_user_groups_limit_languages
    ADD CONSTRAINT "fk-weblate_auth_user_groups_limit_languages-language_id" FOREIGN KEY (language_id) REFERENCES public.lang_language(id);

ALTER TABLE ONLY public.weblate_auth_user_groups_limit_languages
    ADD CONSTRAINT "fk-weblate_auth_user_groups_limit_languages-teammembership_id" FOREIGN KEY (teammembership_id) REFERENCES public.weblate_auth_user_groups(id);

ALTER TABLE ONLY public.weblate_auth_userblock
    ADD CONSTRAINT "fk-weblate_auth_userblock-project_id" FOREIGN KEY (project_id) REFERENCES public.trans_project(id);

ALTER TABLE ONLY public.weblate_auth_userblock
    ADD CONSTRAINT "fk-weblate_auth_userblock-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.wladmin_backuplog
    ADD CONSTRAINT "fk-wladmin_backuplog-service_id" FOREIGN KEY (service_id) REFERENCES public.wladmin_backupservice(id);

ALTER TABLE ONLY public.wladmin_securityfinding
    ADD CONSTRAINT "fk-wladmin_securityfinding-user_id" FOREIGN KEY (user_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.wladmin_securityfindingevent
    ADD CONSTRAINT "fk-wladmin_securityfindingevent-actor_id" FOREIGN KEY (actor_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.wladmin_securityfindingevent
    ADD CONSTRAINT "fk-wladmin_securityfindingevent-finding_id" FOREIGN KEY (finding_id) REFERENCES public.wladmin_securityfinding(id);

ALTER TABLE ONLY public.wladmin_securityfindingevent
    ADD CONSTRAINT "fk-wladmin_securityfindingevent-scan_id" FOREIGN KEY (scan_id) REFERENCES public.wladmin_securityscan(id);

ALTER TABLE ONLY public.wladmin_securityscan
    ADD CONSTRAINT "fk-wladmin_securityscan-actor_id" FOREIGN KEY (actor_id) REFERENCES public.weblate_auth_user(id);

ALTER TABLE ONLY public.workspaces_workspace
    ADD CONSTRAINT "fk-workspaces_workspace-secondary_language_id" FOREIGN KEY (secondary_language_id) REFERENCES public.lang_language(id);

