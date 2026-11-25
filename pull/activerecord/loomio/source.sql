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

CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA public;

COMMENT ON EXTENSION pg_stat_statements IS 'track planning and execution statistics of all SQL statements executed';

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

CREATE TABLE public.account_completion_proofs (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    name_managed boolean DEFAULT false NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    user_id bigint NOT NULL
);

CREATE SEQUENCE public.account_completion_proofs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.account_completion_proofs_id_seq OWNED BY public.account_completion_proofs.id;

CREATE TABLE public.action_mailbox_inbound_emails (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    message_checksum character varying NOT NULL,
    message_id character varying NOT NULL,
    processing_completed_at timestamp(6) without time zone,
    status integer DEFAULT 0 NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.action_mailbox_inbound_emails_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.action_mailbox_inbound_emails_id_seq OWNED BY public.action_mailbox_inbound_emails.id;

CREATE TABLE public.active_storage_attachments (
    id bigint NOT NULL,
    blob_id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    group_id integer,
    name character varying NOT NULL,
    record_id bigint NOT NULL,
    record_type character varying NOT NULL
);

CREATE SEQUENCE public.active_storage_attachments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.active_storage_attachments_id_seq OWNED BY public.active_storage_attachments.id;

CREATE TABLE public.active_storage_blobs (
    id bigint NOT NULL,
    byte_size bigint NOT NULL,
    checksum character varying,
    content_type character varying,
    created_at timestamp without time zone NOT NULL,
    filename character varying NOT NULL,
    key character varying NOT NULL,
    metadata text,
    service_name character varying NOT NULL
);

CREATE SEQUENCE public.active_storage_blobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.active_storage_blobs_id_seq OWNED BY public.active_storage_blobs.id;

CREATE TABLE public.active_storage_variant_records (
    id bigint NOT NULL,
    blob_id bigint NOT NULL,
    variation_digest character varying NOT NULL
);

CREATE SEQUENCE public.active_storage_variant_records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.active_storage_variant_records_id_seq OWNED BY public.active_storage_variant_records.id;

CREATE TABLE public.anonymous_ballot_choices (
    anonymous_ballot_id uuid NOT NULL,
    poll_option_id bigint NOT NULL,
    score integer DEFAULT 1 NOT NULL
);

CREATE TABLE public.anonymous_ballots (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    none_of_the_above boolean DEFAULT false NOT NULL,
    poll_id bigint NOT NULL
);

CREATE TABLE public.anonymous_poll_voters (
    id bigint NOT NULL,
    ballot_submitted boolean DEFAULT false NOT NULL,
    group_member boolean DEFAULT false,
    invited_at timestamp(6) without time zone,
    inviter_id bigint,
    poll_id bigint NOT NULL,
    voter_id bigint NOT NULL
);

CREATE SEQUENCE public.anonymous_poll_voters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.anonymous_poll_voters_id_seq OWNED BY public.anonymous_poll_voters.id;

CREATE TABLE public.ar_internal_metadata (
    key character varying NOT NULL,
    value character varying,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE TABLE public.bookmarks (
    id bigint NOT NULL,
    bookmarkable_id bigint NOT NULL,
    bookmarkable_type character varying NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    discarded_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone NOT NULL,
    user_id bigint NOT NULL
);

CREATE SEQUENCE public.bookmarks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.bookmarks_id_seq OWNED BY public.bookmarks.id;

CREATE TABLE public.chatbots (
    id bigint NOT NULL,
    access_token character varying,
    author_id integer,
    channel character varying,
    created_at timestamp(6) without time zone NOT NULL,
    event_kinds character varying[],
    group_id integer,
    kind character varying,
    name character varying,
    notification_only boolean DEFAULT false NOT NULL,
    server character varying,
    updated_at timestamp(6) without time zone NOT NULL,
    webhook_kind character varying
);

CREATE SEQUENCE public.chatbots_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.chatbots_id_seq OWNED BY public.chatbots.id;

CREATE TABLE public.comments (
    id integer NOT NULL,
    attachments jsonb DEFAULT '[]'::jsonb NOT NULL,
    body text DEFAULT ''::text,
    body_format character varying(10) DEFAULT 'md'::character varying NOT NULL,
    content_locale character varying,
    created_at timestamp without time zone,
    discarded_at timestamp without time zone,
    discarded_by integer,
    edited_at timestamp without time zone,
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    parent_id integer,
    parent_type character varying NOT NULL,
    updated_at timestamp without time zone,
    user_id integer DEFAULT 0,
    versions_count integer DEFAULT 0
);

CREATE SEQUENCE public.comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;

CREATE TABLE public.discussion_templates (
    id bigint NOT NULL,
    allow_comments boolean DEFAULT true NOT NULL,
    allow_reactions boolean DEFAULT true NOT NULL,
    attachments jsonb DEFAULT '[]'::jsonb NOT NULL,
    attachments_admin_only boolean DEFAULT false NOT NULL,
    author_id integer,
    comment_length_max integer,
    content_locale character varying,
    created_at timestamp(6) without time zone NOT NULL,
    default_to_direct_discussion boolean DEFAULT false NOT NULL,
    description text,
    description_format character varying(10) DEFAULT 'html'::character varying NOT NULL,
    discarded_at timestamp(6) without time zone,
    discarded_by integer,
    group_id integer,
    hidden_at timestamp without time zone,
    hider_id integer,
    key character varying,
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    max_depth integer DEFAULT 3 NOT NULL,
    newest_first boolean DEFAULT false NOT NULL,
    poll_template_keys_or_ids jsonb DEFAULT '[]'::jsonb NOT NULL,
    polls_one_at_a_time boolean DEFAULT false NOT NULL,
    "position" integer,
    process_introduction character varying,
    process_introduction_format character varying DEFAULT 'html'::character varying NOT NULL,
    process_name character varying,
    process_subtitle character varying,
    public boolean DEFAULT false NOT NULL,
    recipient_audience character varying,
    source_discussion_id integer,
    tags character varying[] DEFAULT '{}'::character varying[],
    title character varying,
    title_placeholder character varying,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.discussion_templates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.discussion_templates_id_seq OWNED BY public.discussion_templates.id;

CREATE TABLE public.discussions (
    id integer NOT NULL,
    attachments jsonb DEFAULT '[]'::jsonb NOT NULL,
    author_id integer,
    content_locale character varying,
    created_at timestamp without time zone,
    description text,
    description_format character varying(10) DEFAULT 'md'::character varying NOT NULL,
    discarded_at timestamp without time zone,
    discarded_by integer,
    discussion_template_id integer,
    discussion_template_key character varying,
    info jsonb DEFAULT '{}'::jsonb NOT NULL,
    key character varying(255),
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    tags character varying[] DEFAULT '{}'::character varying[],
    template boolean DEFAULT false NOT NULL,
    title character varying(255),
    topic_id integer,
    updated_at timestamp without time zone,
    versions_count integer DEFAULT 0,
    context_changed_at timestamp(6) without time zone
);

CREATE SEQUENCE public.discussions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.discussions_id_seq OWNED BY public.discussions.id;

CREATE TABLE public.email_suppressions (
    id bigint NOT NULL,
    clear_method character varying,
    cleared_at timestamp(6) without time zone,
    cleared_by_id bigint,
    created_at timestamp(6) without time zone NOT NULL,
    email_sha256 character varying NOT NULL,
    notification_delivery_event_id bigint,
    occurred_at timestamp(6) without time zone,
    reason character varying NOT NULL,
    source_key character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.email_suppressions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.email_suppressions_id_seq OWNED BY public.email_suppressions.id;

CREATE TABLE public.forward_email_rules (
    id bigint NOT NULL,
    email character varying,
    handle public.citext NOT NULL
);

CREATE SEQUENCE public.forward_email_rules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.forward_email_rules_id_seq OWNED BY public.forward_email_rules.id;

CREATE TABLE public.group_follows (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    group_id bigint NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    user_id bigint NOT NULL
);

CREATE SEQUENCE public.group_follows_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.group_follows_id_seq OWNED BY public.group_follows.id;

CREATE TABLE public.group_handle_redirects (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    group_id bigint NOT NULL,
    handle public.citext NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.group_handle_redirects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.group_handle_redirects_id_seq OWNED BY public.group_handle_redirects.id;

CREATE TABLE public.groups (
    id integer NOT NULL,
    admin_memberships_count integer DEFAULT 0 NOT NULL,
    admin_tags character varying,
    admins_can_edit_user_content boolean DEFAULT true NOT NULL,
    attachments jsonb DEFAULT '[]'::jsonb NOT NULL,
    category character varying,
    city character varying,
    content_locale character varying,
    country character varying,
    created_at timestamp without time zone,
    creator_id integer,
    description text,
    description_format character varying(10) DEFAULT 'md'::character varying NOT NULL,
    discarded_at timestamp without time zone,
    discarded_by integer,
    discussion_privacy_options character varying DEFAULT 'private_only'::character varying NOT NULL,
    discussions_count integer DEFAULT 0 NOT NULL,
    full_name character varying(255),
    handle public.citext,
    info jsonb DEFAULT '{}'::jsonb NOT NULL,
    is_visible_to_parent_members boolean DEFAULT false NOT NULL,
    is_visible_to_public boolean DEFAULT true NOT NULL,
    key character varying(255),
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    listed_in_explore boolean DEFAULT false NOT NULL,
    members_can_add_guests boolean DEFAULT true NOT NULL,
    members_can_add_members boolean DEFAULT false NOT NULL,
    members_can_announce boolean DEFAULT true NOT NULL,
    members_can_create_subgroups boolean DEFAULT false NOT NULL,
    members_can_create_tags boolean DEFAULT true NOT NULL,
    members_can_create_templates boolean DEFAULT false NOT NULL,
    members_can_delete_comments boolean DEFAULT true NOT NULL,
    members_can_edit_comments boolean DEFAULT true,
    members_can_edit_discussions boolean DEFAULT true NOT NULL,
    members_can_raise_motions boolean DEFAULT true NOT NULL,
    members_can_start_discussions boolean DEFAULT true NOT NULL,
    members_can_vote boolean DEFAULT true NOT NULL,
    membership_granted_upon character varying DEFAULT 'approval'::character varying NOT NULL,
    memberships_count integer DEFAULT 0 NOT NULL,
    name character varying(255),
    non_members_can_start_discussions boolean DEFAULT false NOT NULL,
    org_members_count integer DEFAULT 0 NOT NULL,
    parent_id integer,
    parent_members_can_see_discussions boolean DEFAULT false NOT NULL,
    pending_memberships_count integer DEFAULT 0 NOT NULL,
    poll_templates_count integer DEFAULT 0 NOT NULL,
    polls_count integer DEFAULT 0 NOT NULL,
    region character varying,
    request_to_join_prompt character varying,
    subscription_id integer,
    theme_id integer,
    token character varying,
    updated_at timestamp without time zone,
    is_root boolean GENERATED ALWAYS AS ((parent_id IS NULL)) STORED,
    parent_is_root boolean GENERATED ALWAYS AS (
CASE
    WHEN (parent_id IS NOT NULL) THEN true
    ELSE NULL::boolean
END) STORED
);

CREATE SEQUENCE public.groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;

CREATE TABLE public.legacy_anonymous_vote_reasons (
    anonymous_ballot_id uuid NOT NULL,
    body text NOT NULL
);

CREATE TABLE public.login_tokens (
    id integer NOT NULL,
    code integer NOT NULL,
    created_at timestamp without time zone,
    failed_attempts integer DEFAULT 0 NOT NULL,
    is_reactivation boolean DEFAULT false NOT NULL,
    redirect character varying,
    token character varying,
    updated_at timestamp without time zone,
    user_id integer,
    used_at timestamp(6) without time zone
);

CREATE SEQUENCE public.login_tokens_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.login_tokens_id_seq OWNED BY public.login_tokens.id;

CREATE TABLE public.member_email_aliases (
    id bigint NOT NULL,
    author_id integer NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    email public.citext NOT NULL,
    group_id integer NOT NULL,
    require_dkim boolean DEFAULT true NOT NULL,
    require_spf boolean DEFAULT true NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    user_id integer
);

CREATE SEQUENCE public.member_email_aliases_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.member_email_aliases_id_seq OWNED BY public.member_email_aliases.id;

CREATE TABLE public.membership_requests (
    id integer NOT NULL,
    approved_at timestamp(6) without time zone,
    created_at timestamp without time zone NOT NULL,
    decline_reason text,
    declined_at timestamp(6) without time zone,
    email character varying(255),
    group_id integer NOT NULL,
    introduction text,
    name character varying(255),
    requestor_id integer,
    responder_id integer,
    updated_at timestamp without time zone NOT NULL
);

CREATE SEQUENCE public.membership_requests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.membership_requests_id_seq OWNED BY public.membership_requests.id;

CREATE TABLE public.memberships (
    id integer NOT NULL,
    accepted_at timestamp without time zone,
    admin boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone,
    delegate boolean DEFAULT false NOT NULL,
    experiences jsonb DEFAULT '{}'::jsonb NOT NULL,
    group_id integer NOT NULL,
    inviter_id integer,
    revoked_at timestamp without time zone,
    revoker_id integer,
    title character varying,
    token character varying,
    updated_at timestamp without time zone,
    user_id integer NOT NULL,
    volume_email integer DEFAULT 2 NOT NULL,
    volume_push integer DEFAULT 2 NOT NULL,
    weight numeric(12,3) DEFAULT 1.0 NOT NULL,
    CONSTRAINT memberships_volume_email CHECK ((volume_email = ANY (ARRAY[1, 2, 3]))),
    CONSTRAINT memberships_volume_push CHECK ((volume_push = ANY (ARRAY[1, 2, 3]))),
    CONSTRAINT memberships_weight_nonnegative CHECK ((weight >= (0)::numeric))
);

CREATE SEQUENCE public.memberships_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.memberships_id_seq OWNED BY public.memberships.id;

CREATE TABLE public.mobile_access_tokens (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    mobile_device_id uuid NOT NULL,
    revoked_at timestamp(6) without time zone,
    scopes text NOT NULL,
    token_digest character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.mobile_access_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.mobile_access_tokens_id_seq OWNED BY public.mobile_access_tokens.id;

CREATE TABLE public.mobile_authorization_codes (
    id bigint NOT NULL,
    client_id character varying NOT NULL,
    code_challenge character varying NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    redirect_uri character varying NOT NULL,
    token_digest character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    used_at timestamp(6) without time zone,
    user_id bigint NOT NULL
);

CREATE SEQUENCE public.mobile_authorization_codes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.mobile_authorization_codes_id_seq OWNED BY public.mobile_authorization_codes.id;

CREATE TABLE public.mobile_devices (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    last_seen_at timestamp(6) without time zone NOT NULL,
    name character varying NOT NULL,
    platform character varying DEFAULT 'ios'::character varying NOT NULL,
    protocol_version integer DEFAULT 1 NOT NULL,
    refresh_family_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    revoked_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone NOT NULL,
    user_id bigint NOT NULL,
    CONSTRAINT mobile_devices_protocol_version CHECK ((protocol_version > 0))
);

CREATE TABLE public.mobile_push_registrations (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    delivery_key_ciphertext text NOT NULL,
    mobile_device_id uuid NOT NULL,
    registration_id uuid NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.mobile_push_registrations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.mobile_push_registrations_id_seq OWNED BY public.mobile_push_registrations.id;

CREATE TABLE public.mobile_refresh_tokens (
    id bigint NOT NULL,
    consumed_at timestamp(6) without time zone,
    created_at timestamp(6) without time zone NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    family_id uuid NOT NULL,
    idle_expires_at timestamp(6) without time zone NOT NULL,
    mobile_device_id uuid NOT NULL,
    parent_id bigint,
    revoked_at timestamp(6) without time zone,
    token_digest character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.mobile_refresh_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.mobile_refresh_tokens_id_seq OWNED BY public.mobile_refresh_tokens.id;

CREATE TABLE public.mobile_relay_authorizations (
    id bigint NOT NULL,
    consumed_at timestamp(6) without time zone,
    created_at timestamp(6) without time zone NOT NULL,
    delivery_key_ciphertext text NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    mobile_device_id uuid NOT NULL,
    registration_id uuid NOT NULL,
    token_digest character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.mobile_relay_authorizations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.mobile_relay_authorizations_id_seq OWNED BY public.mobile_relay_authorizations.id;

CREATE TABLE public.mobile_web_session_tickets (
    id bigint NOT NULL,
    consumed_at timestamp(6) without time zone,
    created_at timestamp(6) without time zone NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    mobile_device_id uuid NOT NULL,
    token_digest character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.mobile_web_session_tickets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.mobile_web_session_tickets_id_seq OWNED BY public.mobile_web_session_tickets.id;

CREATE TABLE public.notification_deliveries (
    id bigint NOT NULL,
    channel character varying NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    delivered_at timestamp(6) without time zone,
    email_message_id character varying,
    notification_id bigint NOT NULL,
    provider_message_id character varying,
    recipient_id bigint NOT NULL,
    recipient_type character varying NOT NULL,
    translation_values jsonb DEFAULT '{}'::jsonb NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    viewed_at timestamp(6) without time zone,
    CONSTRAINT notification_deliveries_channel CHECK (((channel)::text = ANY (ARRAY[('in_app'::character varying)::text, ('email'::character varying)::text, ('push'::character varying)::text, ('chatbot'::character varying)::text])))
);

CREATE SEQUENCE public.notification_deliveries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.notification_deliveries_id_seq OWNED BY public.notification_deliveries.id;

CREATE TABLE public.notification_delivery_events (
    id bigint NOT NULL,
    bounce_type character varying,
    created_at timestamp(6) without time zone NOT NULL,
    detail text,
    kind character varying NOT NULL,
    notification_delivery_id bigint NOT NULL,
    occurred_at timestamp(6) without time zone NOT NULL,
    source_key character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.notification_delivery_events_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.notification_delivery_events_id_seq OWNED BY public.notification_delivery_events.id;

CREATE TABLE public.notifications (
    id bigint NOT NULL,
    actor_id integer,
    created_at timestamp(6) without time zone NOT NULL,
    deliveries_generated_at timestamp(6) without time zone,
    kind character varying NOT NULL,
    recipient_audience character varying,
    recipient_chatbot_ids integer[] DEFAULT '{}'::integer[] NOT NULL,
    recipient_context jsonb DEFAULT '{}'::jsonb NOT NULL,
    recipient_message text,
    recipient_user_ids integer[] DEFAULT '{}'::integer[] NOT NULL,
    subject_id bigint NOT NULL,
    subject_type character varying NOT NULL,
    translation_values jsonb DEFAULT '{}'::jsonb NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.notifications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.notifications_id_seq OWNED BY public.notifications.id;

CREATE TABLE public.omniauth_identities (
    id integer NOT NULL,
    access_token character varying DEFAULT ''::character varying,
    created_at timestamp without time zone NOT NULL,
    custom_fields jsonb DEFAULT '{}'::jsonb NOT NULL,
    email character varying(255),
    identity_type character varying(255),
    logo character varying,
    name character varying(255),
    uid character varying(255),
    updated_at timestamp without time zone NOT NULL,
    user_id integer
);

CREATE SEQUENCE public.omniauth_identities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.omniauth_identities_id_seq OWNED BY public.omniauth_identities.id;

CREATE TABLE public.outcomes (
    id integer NOT NULL,
    attachments jsonb DEFAULT '[]'::jsonb NOT NULL,
    author_id integer NOT NULL,
    content_locale character varying,
    created_at timestamp without time zone,
    custom_fields jsonb DEFAULT '{}'::jsonb NOT NULL,
    latest boolean DEFAULT true NOT NULL,
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    poll_id integer NOT NULL,
    poll_option_id integer,
    review_on date,
    statement text NOT NULL,
    statement_format character varying(10) DEFAULT 'md'::character varying NOT NULL,
    updated_at timestamp without time zone,
    versions_count integer DEFAULT 0 NOT NULL
);

CREATE SEQUENCE public.outcomes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.outcomes_id_seq OWNED BY public.outcomes.id;

CREATE TABLE public.partition_sequences (
    counter integer DEFAULT 0,
    id integer NOT NULL,
    key text NOT NULL
);

CREATE TABLE public.passkey_challenges (
    id bigint NOT NULL,
    ceremony character varying NOT NULL,
    challenge_digest character varying NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    user_id bigint
);

CREATE SEQUENCE public.passkey_challenges_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.passkey_challenges_id_seq OWNED BY public.passkey_challenges.id;

CREATE TABLE public.passkey_credentials (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    external_id character varying NOT NULL,
    last_used_at timestamp(6) without time zone,
    name character varying NOT NULL,
    public_key bytea NOT NULL,
    sign_count bigint DEFAULT 0 NOT NULL,
    transports jsonb DEFAULT '[]'::jsonb NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    user_handle character varying NOT NULL,
    user_id bigint NOT NULL
);

CREATE SEQUENCE public.passkey_credentials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.passkey_credentials_id_seq OWNED BY public.passkey_credentials.id;

CREATE TABLE public.pending_emails (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    group_id bigint NOT NULL,
    inbound_email_id bigint NOT NULL,
    sender_email character varying,
    sender_name character varying,
    status integer DEFAULT 0 NOT NULL,
    subject character varying,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.pending_emails_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.pending_emails_id_seq OWNED BY public.pending_emails.id;

CREATE TABLE public.pg_search_documents (
    id bigint NOT NULL,
    author_id bigint,
    authored_at timestamp(6) without time zone,
    content text,
    created_at timestamp(6) without time zone NOT NULL,
    discussion_id bigint,
    group_id bigint,
    poll_id bigint,
    searchable_id bigint,
    searchable_type character varying,
    tags character varying[] DEFAULT '{}'::character varying[],
    topic_id bigint,
    ts_content tsvector,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.pg_search_documents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.pg_search_documents_id_seq OWNED BY public.pg_search_documents.id;

CREATE TABLE public.pg_search_words (
    document_count integer NOT NULL,
    word text NOT NULL
);

CREATE TABLE public.poll_options (
    id integer NOT NULL,
    created_at timestamp without time zone,
    icon character varying,
    meaning character varying,
    name character varying NOT NULL,
    poll_id integer NOT NULL,
    priority integer DEFAULT 0 NOT NULL,
    prompt character varying,
    scores_by_voter jsonb DEFAULT '{}'::jsonb NOT NULL,
    test_against character varying,
    test_operator character varying,
    test_percent integer,
    total_score numeric(30,3) DEFAULT 0.0 NOT NULL,
    unweighted_score bigint DEFAULT 0 NOT NULL,
    updated_at timestamp without time zone,
    vote_count integer DEFAULT 0 NOT NULL,
    vote_counts_by_score jsonb DEFAULT '{}'::jsonb NOT NULL,
    vote_weight numeric(30,3) DEFAULT 0.0 NOT NULL
);

CREATE SEQUENCE public.poll_options_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.poll_options_id_seq OWNED BY public.poll_options.id;

CREATE TABLE public.poll_templates (
    id bigint NOT NULL,
    agree_target integer,
    allow_comments boolean DEFAULT true NOT NULL,
    allow_reactions boolean DEFAULT true NOT NULL,
    anonymous boolean DEFAULT false NOT NULL,
    attachments_admin_only boolean DEFAULT false NOT NULL,
    author_id integer NOT NULL,
    can_respond_maybe boolean DEFAULT true NOT NULL,
    chart_type character varying,
    choice_count_max integer,
    choice_count_min integer,
    comment_length_max integer,
    content_locale character varying,
    created_at timestamp(6) without time zone NOT NULL,
    default_duration_in_days integer DEFAULT 7 NOT NULL,
    details text,
    details_format character varying(10) DEFAULT 'md'::character varying NOT NULL,
    discarded_at timestamp(6) without time zone,
    discarded_by integer,
    group_id integer NOT NULL,
    hidden_at timestamp(6) without time zone,
    hide_results integer DEFAULT 0 NOT NULL,
    hider_id integer,
    key character varying,
    limit_reason_length boolean DEFAULT true NOT NULL,
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    meeting_duration integer,
    notify_on_closing_soon integer DEFAULT 0 NOT NULL,
    notify_on_open boolean DEFAULT true NOT NULL,
    outcome_review_due_in_days integer,
    outcome_statement character varying,
    outcome_statement_format character varying DEFAULT 'html'::character varying NOT NULL,
    points_per_voter integer,
    poll_option_name_format character varying DEFAULT 'plain'::character varying NOT NULL,
    poll_options jsonb DEFAULT '[]'::jsonb NOT NULL,
    poll_type character varying NOT NULL,
    "position" integer DEFAULT 0 NOT NULL,
    process_introduction character varying,
    process_introduction_format character varying DEFAULT 'md'::character varying NOT NULL,
    process_name character varying,
    process_subtitle character varying,
    public boolean DEFAULT false NOT NULL,
    quorum_pct integer,
    reason_prompt character varying,
    score_max integer,
    score_min integer,
    show_none_of_the_above boolean DEFAULT false NOT NULL,
    shuffle_options boolean DEFAULT false NOT NULL,
    specified_voters_only boolean DEFAULT false NOT NULL,
    stance_reason_required integer DEFAULT 1 NOT NULL,
    stv_method character varying DEFAULT 'scottish'::character varying NOT NULL,
    stv_quota character varying DEFAULT 'droop'::character varying NOT NULL,
    stv_seats integer DEFAULT 1 NOT NULL,
    tags character varying[] DEFAULT '{}'::character varying[],
    title character varying,
    title_placeholder character varying,
    updated_at timestamp(6) without time zone NOT NULL,
    weighted_voting boolean DEFAULT false NOT NULL
);

CREATE SEQUENCE public.poll_templates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.poll_templates_id_seq OWNED BY public.poll_templates.id;

CREATE TABLE public.polls (
    id integer NOT NULL,
    agree_target integer,
    anonymous boolean DEFAULT false NOT NULL,
    attachments jsonb DEFAULT '[]'::jsonb NOT NULL,
    author_id integer NOT NULL,
    chart_type character varying,
    choice_count_max integer,
    choice_count_min integer,
    closed_at timestamp without time zone,
    closing_at timestamp without time zone,
    content_locale character varying,
    created_at timestamp without time zone,
    custom_fields jsonb DEFAULT '{}'::jsonb NOT NULL,
    default_duration_in_days integer,
    details text,
    details_format character varying(10) DEFAULT 'md'::character varying NOT NULL,
    discarded_at timestamp without time zone,
    discarded_by integer,
    hide_results integer DEFAULT 0 NOT NULL,
    key character varying NOT NULL,
    limit_reason_length boolean DEFAULT true NOT NULL,
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    notify_on_closing_soon integer DEFAULT 0 NOT NULL,
    notify_on_open boolean DEFAULT true NOT NULL,
    opened_at timestamp(6) without time zone,
    opening_at timestamp(6) without time zone,
    points_per_voter integer,
    poll_option_name_format character varying,
    poll_template_id integer,
    poll_template_key character varying,
    poll_type character varying NOT NULL,
    process_name character varying,
    process_subtitle character varying,
    process_url character varying,
    quorum_pct integer,
    reason_prompt character varying,
    score_max integer,
    score_min integer,
    scores_by_option jsonb DEFAULT '[]'::jsonb NOT NULL,
    show_none_of_the_above boolean DEFAULT false NOT NULL,
    shuffle_options boolean DEFAULT false NOT NULL,
    specified_voters_only boolean DEFAULT false NOT NULL,
    stance_reason_required integer DEFAULT 1 NOT NULL,
    stv_method character varying,
    stv_quota character varying,
    stv_seats integer,
    tags character varying[] DEFAULT '{}'::character varying[],
    template boolean DEFAULT false NOT NULL,
    title character varying NOT NULL,
    topic_id integer,
    updated_at timestamp without time zone,
    versions_count integer DEFAULT 0,
    vote_count_none_of_the_above integer DEFAULT 0 NOT NULL,
    vote_count_eligible integer DEFAULT 0 NOT NULL,
    vote_count_uncast integer DEFAULT 0 NOT NULL,
    weighted_voting boolean DEFAULT false NOT NULL
);

CREATE SEQUENCE public.polls_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.polls_id_seq OWNED BY public.polls.id;

CREATE TABLE public.push_subscriptions (
    id bigint NOT NULL,
    auth_key character varying NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    endpoint text NOT NULL,
    endpoint_digest character varying NOT NULL,
    expires_at timestamp(6) without time zone,
    failure_count integer DEFAULT 0 NOT NULL,
    last_seen_at timestamp(6) without time zone,
    name character varying,
    p256dh_key character varying NOT NULL,
    revoked_at timestamp(6) without time zone,
    session_id bigint NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    user_agent character varying,
    user_id bigint NOT NULL,
    CONSTRAINT push_subscriptions_failure_count CHECK ((failure_count >= 0))
);

CREATE SEQUENCE public.push_subscriptions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.push_subscriptions_id_seq OWNED BY public.push_subscriptions.id;

CREATE TABLE public.reactions (
    id integer NOT NULL,
    created_at timestamp without time zone,
    reactable_id integer,
    reactable_type character varying DEFAULT 'Comment'::character varying NOT NULL,
    reaction character varying DEFAULT '+1'::character varying NOT NULL,
    updated_at timestamp without time zone,
    user_id integer
);

CREATE SEQUENCE public.reactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.reactions_id_seq OWNED BY public.reactions.id;

CREATE TABLE public.schema_migrations (
    version character varying NOT NULL
);

CREATE TABLE public.sessions (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    ip_address character varying,
    revoked_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone NOT NULL,
    user_agent character varying,
    user_id bigint NOT NULL
);

CREATE SEQUENCE public.sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.sessions_id_seq OWNED BY public.sessions.id;

CREATE TABLE public.solid_cable_messages (
    id bigint NOT NULL,
    channel bytea NOT NULL,
    channel_hash bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    payload bytea NOT NULL
);

CREATE SEQUENCE public.solid_cable_messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_cable_messages_id_seq OWNED BY public.solid_cable_messages.id;

CREATE TABLE public.solid_cache_entries (
    id bigint NOT NULL,
    byte_size integer NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    key bytea NOT NULL,
    key_hash bigint NOT NULL,
    value bytea NOT NULL
);

CREATE SEQUENCE public.solid_cache_entries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_cache_entries_id_seq OWNED BY public.solid_cache_entries.id;

CREATE TABLE public.solid_queue_batch_executions (
    id bigint NOT NULL,
    batch_id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    job_id bigint NOT NULL
);

CREATE SEQUENCE public.solid_queue_batch_executions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_batch_executions_id_seq OWNED BY public.solid_queue_batch_executions.id;

CREATE TABLE public.solid_queue_batches (
    id bigint NOT NULL,
    active_job_batch_id character varying,
    completed_jobs integer DEFAULT 0 NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    description character varying,
    enqueued_at timestamp(6) without time zone,
    failed_at timestamp(6) without time zone,
    failed_jobs integer DEFAULT 0 NOT NULL,
    finished_at timestamp(6) without time zone,
    metadata text,
    on_failure text,
    on_finish text,
    on_success text,
    total_jobs integer DEFAULT 0 NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.solid_queue_batches_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_batches_id_seq OWNED BY public.solid_queue_batches.id;

CREATE TABLE public.solid_queue_blocked_executions (
    id bigint NOT NULL,
    concurrency_key character varying NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    job_id bigint NOT NULL,
    priority integer DEFAULT 0 NOT NULL,
    queue_name character varying NOT NULL
);

CREATE SEQUENCE public.solid_queue_blocked_executions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_blocked_executions_id_seq OWNED BY public.solid_queue_blocked_executions.id;

CREATE TABLE public.solid_queue_claimed_executions (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    job_id bigint NOT NULL,
    process_id bigint
);

CREATE SEQUENCE public.solid_queue_claimed_executions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_claimed_executions_id_seq OWNED BY public.solid_queue_claimed_executions.id;

CREATE TABLE public.solid_queue_failed_executions (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    error text,
    job_id bigint NOT NULL
);

CREATE SEQUENCE public.solid_queue_failed_executions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_failed_executions_id_seq OWNED BY public.solid_queue_failed_executions.id;

CREATE TABLE public.solid_queue_jobs (
    id bigint NOT NULL,
    active_job_id character varying,
    arguments text,
    batch_id bigint,
    class_name character varying NOT NULL,
    concurrency_key character varying,
    created_at timestamp(6) without time zone NOT NULL,
    finished_at timestamp(6) without time zone,
    priority integer DEFAULT 0 NOT NULL,
    queue_name character varying NOT NULL,
    scheduled_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.solid_queue_jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_jobs_id_seq OWNED BY public.solid_queue_jobs.id;

CREATE TABLE public.solid_queue_pauses (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    queue_name character varying NOT NULL
);

CREATE SEQUENCE public.solid_queue_pauses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_pauses_id_seq OWNED BY public.solid_queue_pauses.id;

CREATE TABLE public.solid_queue_processes (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    hostname character varying,
    kind character varying NOT NULL,
    last_heartbeat_at timestamp(6) without time zone NOT NULL,
    metadata text,
    name character varying NOT NULL,
    pid integer NOT NULL,
    supervisor_id bigint
);

CREATE SEQUENCE public.solid_queue_processes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_processes_id_seq OWNED BY public.solid_queue_processes.id;

CREATE TABLE public.solid_queue_ready_executions (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    job_id bigint NOT NULL,
    priority integer DEFAULT 0 NOT NULL,
    queue_name character varying NOT NULL
);

CREATE SEQUENCE public.solid_queue_ready_executions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_ready_executions_id_seq OWNED BY public.solid_queue_ready_executions.id;

CREATE TABLE public.solid_queue_recurring_executions (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    job_id bigint NOT NULL,
    run_at timestamp(6) without time zone NOT NULL,
    task_key character varying NOT NULL
);

CREATE SEQUENCE public.solid_queue_recurring_executions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_recurring_executions_id_seq OWNED BY public.solid_queue_recurring_executions.id;

CREATE TABLE public.solid_queue_recurring_tasks (
    id bigint NOT NULL,
    arguments text,
    class_name character varying,
    command character varying(2048),
    created_at timestamp(6) without time zone NOT NULL,
    description text,
    key character varying NOT NULL,
    priority integer DEFAULT 0,
    queue_name character varying,
    schedule character varying NOT NULL,
    static boolean DEFAULT true NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.solid_queue_recurring_tasks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_recurring_tasks_id_seq OWNED BY public.solid_queue_recurring_tasks.id;

CREATE TABLE public.solid_queue_scheduled_executions (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    job_id bigint NOT NULL,
    priority integer DEFAULT 0 NOT NULL,
    queue_name character varying NOT NULL,
    scheduled_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.solid_queue_scheduled_executions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_scheduled_executions_id_seq OWNED BY public.solid_queue_scheduled_executions.id;

CREATE TABLE public.solid_queue_semaphores (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    expires_at timestamp(6) without time zone NOT NULL,
    key character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    value integer DEFAULT 1 NOT NULL
);

CREATE SEQUENCE public.solid_queue_semaphores_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.solid_queue_semaphores_id_seq OWNED BY public.solid_queue_semaphores.id;

CREATE TABLE public.spam_rule_hits (
    id bigint NOT NULL,
    spam_rule_id bigint NOT NULL,
    action character varying NOT NULL,
    bucket_at timestamp(6) without time zone NOT NULL,
    count bigint DEFAULT 0 NOT NULL,
    last_matched_at timestamp(6) without time zone NOT NULL,
    CONSTRAINT spam_rule_hits_positive_count CHECK ((count > 0))
);

CREATE SEQUENCE public.spam_rule_hits_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.spam_rule_hits_id_seq OWNED BY public.spam_rule_hits.id;

CREATE TABLE public.spam_rules (
    id bigint NOT NULL,
    kind character varying NOT NULL,
    pattern text NOT NULL,
    reason text,
    expires_at timestamp(6) without time zone,
    discarded_at timestamp(6) without time zone,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    CONSTRAINT spam_rules_kind CHECK (((kind)::text = ANY (ARRAY[('ip'::character varying)::text, ('email_domain'::character varying)::text, ('regex'::character varying)::text]))),
    CONSTRAINT spam_rules_pattern_length CHECK (((length(btrim(pattern)) >= 1) AND (length(btrim(pattern)) <= 4096)))
);

CREATE SEQUENCE public.spam_rules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.spam_rules_id_seq OWNED BY public.spam_rules.id;

CREATE TABLE public.stance_choices (
    id integer NOT NULL,
    created_at timestamp without time zone,
    poll_option_id integer NOT NULL,
    score integer DEFAULT 1 NOT NULL,
    stance_id integer NOT NULL,
    updated_at timestamp without time zone
);

CREATE SEQUENCE public.stance_choices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.stance_choices_id_seq OWNED BY public.stance_choices.id;

CREATE TABLE public.stances (
    id integer NOT NULL,
    accepted_at timestamp without time zone,
    attachments jsonb DEFAULT '[]'::jsonb NOT NULL,
    cast_at timestamp without time zone,
    content_locale character varying,
    created_at timestamp without time zone,
    inviter_id integer,
    latest boolean DEFAULT true NOT NULL,
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    none_of_the_above boolean DEFAULT false NOT NULL,
    option_scores jsonb DEFAULT '{}'::jsonb NOT NULL,
    participant_id integer NOT NULL,
    poll_id integer NOT NULL,
    reason character varying,
    reason_format character varying(10) DEFAULT 'md'::character varying NOT NULL,
    redacted_at timestamp(6) without time zone,
    redactor_id integer,
    revoked_at timestamp without time zone,
    revoker_id integer,
    token character varying,
    updated_at timestamp without time zone,
    versions_count integer DEFAULT 0,
    weight numeric(12,3) DEFAULT 1.0 NOT NULL,
    CONSTRAINT stances_weight_nonnegative CHECK ((weight >= (0)::numeric))
);

CREATE SEQUENCE public.stances_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.stances_id_seq OWNED BY public.stances.id;

CREATE TABLE public.subscription_update_receipts (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    event_id character varying NOT NULL,
    payload_digest character varying NOT NULL,
    processed_at timestamp(6) without time zone,
    subscription_id bigint,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.subscription_update_receipts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.subscription_update_receipts_id_seq OWNED BY public.subscription_update_receipts.id;

CREATE TABLE public.subscriptions (
    id integer NOT NULL,
    activated_at timestamp without time zone,
    allow_guests boolean DEFAULT true NOT NULL,
    allow_subgroups boolean DEFAULT true NOT NULL,
    billing_service_price_point_id character varying,
    billing_service_product_id character varying,
    billing_service_subscription_id character varying,
    billing_service_updated_at timestamp(6) without time zone,
    canceled_at timestamp without time zone,
    chargify_subscription_id integer,
    created_at timestamp without time zone,
    expires_at timestamp without time zone,
    info jsonb,
    lead_status character varying,
    max_members integer,
    max_orgs integer,
    max_threads integer,
    members_count integer,
    owner_id integer,
    payment_method character varying DEFAULT 'none'::character varying NOT NULL,
    plan character varying DEFAULT 'free'::character varying,
    renewed_at timestamp without time zone,
    renews_at timestamp without time zone,
    state character varying DEFAULT 'active'::character varying NOT NULL,
    updated_at timestamp without time zone
);

CREATE SEQUENCE public.subscriptions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.subscriptions_id_seq OWNED BY public.subscriptions.id;

CREATE TABLE public.taggings (
    id integer NOT NULL,
    created_at timestamp without time zone,
    tag_id integer NOT NULL,
    taggable_id integer NOT NULL,
    taggable_type character varying NOT NULL,
    updated_at timestamp without time zone
);

CREATE SEQUENCE public.taggings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.taggings_id_seq OWNED BY public.taggings.id;

CREATE TABLE public.tags (
    id integer NOT NULL,
    color character varying,
    created_at timestamp without time zone,
    group_id integer NOT NULL,
    name public.citext NOT NULL,
    priority integer DEFAULT 0 NOT NULL,
    updated_at timestamp without time zone,
    used_group_ids integer[] DEFAULT '{}'::integer[] NOT NULL
);

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;

CREATE TABLE public.tasks (
    id bigint NOT NULL,
    author_id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    discarded_at timestamp without time zone,
    doer_id integer,
    done_at timestamp without time zone,
    due_on date,
    name character varying NOT NULL,
    record_id bigint,
    record_type character varying,
    remind integer,
    remind_at timestamp without time zone,
    uid integer NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.tasks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.tasks_id_seq OWNED BY public.tasks.id;

CREATE TABLE public.tasks_users (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    task_id bigint NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    user_id bigint NOT NULL
);

CREATE SEQUENCE public.tasks_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.tasks_users_id_seq OWNED BY public.tasks_users.id;

CREATE TABLE public.topic_items (
    id integer NOT NULL,
    child_count integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone,
    depth integer DEFAULT 0 NOT NULL,
    itemable_id integer NOT NULL,
    itemable_type character varying(255) NOT NULL,
    kind character varying(255) NOT NULL,
    parent_id integer,
    pinned boolean DEFAULT false NOT NULL,
    pinned_title text,
    "position" integer DEFAULT 0 NOT NULL,
    position_key character varying,
    sequence_id integer,
    tags character varying[] DEFAULT '{}'::character varying[] NOT NULL,
    topic_id integer NOT NULL,
    updated_at timestamp without time zone,
    user_id integer,
    CONSTRAINT topic_items_itemable_type_present CHECK ((btrim((itemable_type)::text) <> ''::text)),
    CONSTRAINT topic_items_kind_present CHECK ((btrim((kind)::text) <> ''::text)),
    CONSTRAINT topic_items_tags_supported CHECK (((cardinality(tags) = 0) OR (((kind)::text = 'new_discussion'::text) AND ((itemable_type)::text = 'Discussion'::text)) OR (((kind)::text = 'poll_created'::text) AND ((itemable_type)::text = 'Poll'::text)) OR (((kind)::text = 'outcome_created'::text) AND ((itemable_type)::text = 'Outcome'::text)) OR (((kind)::text = 'new_comment'::text) AND ((itemable_type)::text = 'Comment'::text))))
);

CREATE SEQUENCE public.topic_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.topic_items_id_seq OWNED BY public.topic_items.id;

CREATE TABLE public.topic_readers (
    id integer NOT NULL,
    accepted_at timestamp without time zone,
    admin boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone,
    dismissed_at timestamp without time zone,
    guest boolean DEFAULT false NOT NULL,
    inviter_id integer,
    last_read_at timestamp without time zone,
    read_ranges_string character varying,
    revoked_at timestamp without time zone,
    revoker_id integer,
    token character varying,
    topic_id integer NOT NULL,
    updated_at timestamp without time zone,
    user_id integer NOT NULL,
    volume_email integer DEFAULT 2 NOT NULL,
    volume_push integer DEFAULT 2 NOT NULL,
    CONSTRAINT topic_readers_volume_email CHECK ((volume_email = ANY (ARRAY[1, 2, 3]))),
    CONSTRAINT topic_readers_volume_push CHECK ((volume_push = ANY (ARRAY[1, 2, 3])))
);

CREATE SEQUENCE public.topic_readers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.topic_readers_id_seq OWNED BY public.topic_readers.id;

CREATE TABLE public.topics (
    id bigint NOT NULL,
    active_polls_count integer DEFAULT 0 NOT NULL,
    allow_comments boolean DEFAULT true NOT NULL,
    allow_reactions boolean DEFAULT true NOT NULL,
    attachments_admin_only boolean DEFAULT false NOT NULL,
    comment_length_max integer,
    created_at timestamp(6) without time zone NOT NULL,
    discarded_at timestamp(6) without time zone,
    discarded_by integer,
    group_id integer,
    has_tagged_replies boolean DEFAULT false NOT NULL,
    items_count integer DEFAULT 0 NOT NULL,
    last_activity_at timestamp without time zone,
    locked_at timestamp without time zone,
    locker_id integer,
    max_depth integer DEFAULT 3 NOT NULL,
    newest_first boolean DEFAULT false NOT NULL,
    pinned_at timestamp without time zone,
    polls_one_at_a_time boolean DEFAULT false NOT NULL,
    private boolean DEFAULT true NOT NULL,
    ranges_string character varying,
    seen_by_count integer DEFAULT 0 NOT NULL,
    tags character varying[] DEFAULT '{}'::character varying[],
    topicable_id integer NOT NULL,
    topicable_type character varying NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.topics_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.topics_id_seq OWNED BY public.topics.id;

CREATE TABLE public.translations (
    id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    fields public.hstore,
    language character varying(255),
    translatable_id integer,
    translatable_type character varying(255),
    updated_at timestamp without time zone NOT NULL
);

CREATE SEQUENCE public.translations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.translations_id_seq OWNED BY public.translations.id;

CREATE TABLE public.trial_invitation_recipients (
    id bigint NOT NULL,
    ip_address inet NOT NULL,
    actor_id integer NOT NULL,
    recipient_id integer NOT NULL,
    invited_at timestamp(6) without time zone NOT NULL
);

CREATE SEQUENCE public.trial_invitation_recipients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.trial_invitation_recipients_id_seq OWNED BY public.trial_invitation_recipients.id;

CREATE TABLE public.users (
    id integer NOT NULL,
    api_key character varying NOT NULL,
    attachments jsonb DEFAULT '[]'::jsonb NOT NULL,
    auto_translate boolean DEFAULT false NOT NULL,
    autodetect_time_zone boolean DEFAULT true NOT NULL,
    avatar_initials character varying(255),
    avatar_kind character varying(255) DEFAULT 'initials'::character varying NOT NULL,
    bot boolean DEFAULT false NOT NULL,
    city character varying,
    content_locale character varying,
    country character varying,
    created_at timestamp without time zone,
    current_sign_in_at timestamp without time zone,
    current_sign_in_ip inet,
    date_time_pref character varying,
    deactivated_at timestamp without time zone,
    deactivator_id integer,
    detected_locale character varying(255),
    email public.citext,
    email_api_key character varying(255) NOT NULL,
    email_catch_up_day integer DEFAULT 7,
    email_newsletter boolean DEFAULT false NOT NULL,
    email_sha256 character varying,
    email_verified boolean DEFAULT false NOT NULL,
    experiences jsonb DEFAULT '{}'::jsonb NOT NULL,
    failed_attempts integer DEFAULT 0 NOT NULL,
    is_admin boolean DEFAULT false,
    key character varying(255),
    last_seen_at timestamp without time zone,
    last_sign_in_at timestamp without time zone,
    last_sign_in_ip inet,
    legal_accepted_at timestamp without time zone,
    link_previews jsonb DEFAULT '[]'::jsonb NOT NULL,
    location character varying DEFAULT ''::character varying NOT NULL,
    locked_at timestamp without time zone,
    memberships_count integer DEFAULT 0 NOT NULL,
    name character varying(255),
    password_digest character varying(128) DEFAULT ''::character varying,
    region character varying,
    secret_token character varying DEFAULT public.gen_random_uuid() NOT NULL,
    selected_locale character varying(255),
    short_bio character varying DEFAULT ''::character varying NOT NULL,
    short_bio_format character varying(10) DEFAULT 'md'::character varying NOT NULL,
    sign_in_count integer DEFAULT 0,
    time_zone character varying(255),
    unsubscribe_token character varying(255) NOT NULL,
    updated_at timestamp without time zone,
    username character varying(255),
    volume_email_default integer DEFAULT 2 NOT NULL,
    volume_push_default integer DEFAULT 2 NOT NULL,
    webauthn_id character varying,
    CONSTRAINT users_volume_email_default CHECK ((volume_email_default = ANY (ARRAY[1, 2, 3]))),
    CONSTRAINT users_volume_push_default CHECK ((volume_push_default = ANY (ARRAY[1, 2, 3])))
);

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;

CREATE TABLE public.versions (
    id integer NOT NULL,
    created_at timestamp without time zone,
    event character varying(255) NOT NULL,
    item_id integer NOT NULL,
    item_type character varying(255) NOT NULL,
    object_changes jsonb,
    whodunnit integer
);

CREATE SEQUENCE public.versions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.versions_id_seq OWNED BY public.versions.id;

ALTER TABLE ONLY public.account_completion_proofs ALTER COLUMN id SET DEFAULT nextval('public.account_completion_proofs_id_seq'::regclass);

ALTER TABLE ONLY public.action_mailbox_inbound_emails ALTER COLUMN id SET DEFAULT nextval('public.action_mailbox_inbound_emails_id_seq'::regclass);

ALTER TABLE ONLY public.active_storage_attachments ALTER COLUMN id SET DEFAULT nextval('public.active_storage_attachments_id_seq'::regclass);

ALTER TABLE ONLY public.active_storage_blobs ALTER COLUMN id SET DEFAULT nextval('public.active_storage_blobs_id_seq'::regclass);

ALTER TABLE ONLY public.active_storage_variant_records ALTER COLUMN id SET DEFAULT nextval('public.active_storage_variant_records_id_seq'::regclass);

ALTER TABLE ONLY public.anonymous_poll_voters ALTER COLUMN id SET DEFAULT nextval('public.anonymous_poll_voters_id_seq'::regclass);

ALTER TABLE ONLY public.bookmarks ALTER COLUMN id SET DEFAULT nextval('public.bookmarks_id_seq'::regclass);

ALTER TABLE ONLY public.chatbots ALTER COLUMN id SET DEFAULT nextval('public.chatbots_id_seq'::regclass);

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);

ALTER TABLE ONLY public.discussion_templates ALTER COLUMN id SET DEFAULT nextval('public.discussion_templates_id_seq'::regclass);

ALTER TABLE ONLY public.discussions ALTER COLUMN id SET DEFAULT nextval('public.discussions_id_seq'::regclass);

ALTER TABLE ONLY public.email_suppressions ALTER COLUMN id SET DEFAULT nextval('public.email_suppressions_id_seq'::regclass);

ALTER TABLE ONLY public.forward_email_rules ALTER COLUMN id SET DEFAULT nextval('public.forward_email_rules_id_seq'::regclass);

ALTER TABLE ONLY public.group_follows ALTER COLUMN id SET DEFAULT nextval('public.group_follows_id_seq'::regclass);

ALTER TABLE ONLY public.group_handle_redirects ALTER COLUMN id SET DEFAULT nextval('public.group_handle_redirects_id_seq'::regclass);

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);

ALTER TABLE ONLY public.login_tokens ALTER COLUMN id SET DEFAULT nextval('public.login_tokens_id_seq'::regclass);

ALTER TABLE ONLY public.member_email_aliases ALTER COLUMN id SET DEFAULT nextval('public.member_email_aliases_id_seq'::regclass);

ALTER TABLE ONLY public.membership_requests ALTER COLUMN id SET DEFAULT nextval('public.membership_requests_id_seq'::regclass);

ALTER TABLE ONLY public.memberships ALTER COLUMN id SET DEFAULT nextval('public.memberships_id_seq'::regclass);

ALTER TABLE ONLY public.mobile_access_tokens ALTER COLUMN id SET DEFAULT nextval('public.mobile_access_tokens_id_seq'::regclass);

ALTER TABLE ONLY public.mobile_authorization_codes ALTER COLUMN id SET DEFAULT nextval('public.mobile_authorization_codes_id_seq'::regclass);

ALTER TABLE ONLY public.mobile_push_registrations ALTER COLUMN id SET DEFAULT nextval('public.mobile_push_registrations_id_seq'::regclass);

ALTER TABLE ONLY public.mobile_refresh_tokens ALTER COLUMN id SET DEFAULT nextval('public.mobile_refresh_tokens_id_seq'::regclass);

ALTER TABLE ONLY public.mobile_relay_authorizations ALTER COLUMN id SET DEFAULT nextval('public.mobile_relay_authorizations_id_seq'::regclass);

ALTER TABLE ONLY public.mobile_web_session_tickets ALTER COLUMN id SET DEFAULT nextval('public.mobile_web_session_tickets_id_seq'::regclass);

ALTER TABLE ONLY public.notification_deliveries ALTER COLUMN id SET DEFAULT nextval('public.notification_deliveries_id_seq'::regclass);

ALTER TABLE ONLY public.notification_delivery_events ALTER COLUMN id SET DEFAULT nextval('public.notification_delivery_events_id_seq'::regclass);

ALTER TABLE ONLY public.notifications ALTER COLUMN id SET DEFAULT nextval('public.notifications_id_seq'::regclass);

ALTER TABLE ONLY public.omniauth_identities ALTER COLUMN id SET DEFAULT nextval('public.omniauth_identities_id_seq'::regclass);

ALTER TABLE ONLY public.outcomes ALTER COLUMN id SET DEFAULT nextval('public.outcomes_id_seq'::regclass);

ALTER TABLE ONLY public.passkey_challenges ALTER COLUMN id SET DEFAULT nextval('public.passkey_challenges_id_seq'::regclass);

ALTER TABLE ONLY public.passkey_credentials ALTER COLUMN id SET DEFAULT nextval('public.passkey_credentials_id_seq'::regclass);

ALTER TABLE ONLY public.pending_emails ALTER COLUMN id SET DEFAULT nextval('public.pending_emails_id_seq'::regclass);

ALTER TABLE ONLY public.pg_search_documents ALTER COLUMN id SET DEFAULT nextval('public.pg_search_documents_id_seq'::regclass);

ALTER TABLE ONLY public.poll_options ALTER COLUMN id SET DEFAULT nextval('public.poll_options_id_seq'::regclass);

ALTER TABLE ONLY public.poll_templates ALTER COLUMN id SET DEFAULT nextval('public.poll_templates_id_seq'::regclass);

ALTER TABLE ONLY public.polls ALTER COLUMN id SET DEFAULT nextval('public.polls_id_seq'::regclass);

ALTER TABLE ONLY public.push_subscriptions ALTER COLUMN id SET DEFAULT nextval('public.push_subscriptions_id_seq'::regclass);

ALTER TABLE ONLY public.reactions ALTER COLUMN id SET DEFAULT nextval('public.reactions_id_seq'::regclass);

ALTER TABLE ONLY public.sessions ALTER COLUMN id SET DEFAULT nextval('public.sessions_id_seq'::regclass);

ALTER TABLE ONLY public.solid_cable_messages ALTER COLUMN id SET DEFAULT nextval('public.solid_cable_messages_id_seq'::regclass);

ALTER TABLE ONLY public.solid_cache_entries ALTER COLUMN id SET DEFAULT nextval('public.solid_cache_entries_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_batch_executions ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_batch_executions_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_batches ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_batches_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_blocked_executions ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_blocked_executions_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_claimed_executions ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_claimed_executions_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_failed_executions ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_failed_executions_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_jobs ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_jobs_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_pauses ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_pauses_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_processes ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_processes_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_ready_executions ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_ready_executions_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_recurring_executions ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_recurring_executions_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_recurring_tasks ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_recurring_tasks_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_scheduled_executions ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_scheduled_executions_id_seq'::regclass);

ALTER TABLE ONLY public.solid_queue_semaphores ALTER COLUMN id SET DEFAULT nextval('public.solid_queue_semaphores_id_seq'::regclass);

ALTER TABLE ONLY public.spam_rule_hits ALTER COLUMN id SET DEFAULT nextval('public.spam_rule_hits_id_seq'::regclass);

ALTER TABLE ONLY public.spam_rules ALTER COLUMN id SET DEFAULT nextval('public.spam_rules_id_seq'::regclass);

ALTER TABLE ONLY public.stance_choices ALTER COLUMN id SET DEFAULT nextval('public.stance_choices_id_seq'::regclass);

ALTER TABLE ONLY public.stances ALTER COLUMN id SET DEFAULT nextval('public.stances_id_seq'::regclass);

ALTER TABLE ONLY public.subscription_update_receipts ALTER COLUMN id SET DEFAULT nextval('public.subscription_update_receipts_id_seq'::regclass);

ALTER TABLE ONLY public.subscriptions ALTER COLUMN id SET DEFAULT nextval('public.subscriptions_id_seq'::regclass);

ALTER TABLE ONLY public.taggings ALTER COLUMN id SET DEFAULT nextval('public.taggings_id_seq'::regclass);

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);

ALTER TABLE ONLY public.tasks ALTER COLUMN id SET DEFAULT nextval('public.tasks_id_seq'::regclass);

ALTER TABLE ONLY public.tasks_users ALTER COLUMN id SET DEFAULT nextval('public.tasks_users_id_seq'::regclass);

ALTER TABLE ONLY public.topic_items ALTER COLUMN id SET DEFAULT nextval('public.topic_items_id_seq'::regclass);

ALTER TABLE ONLY public.topic_readers ALTER COLUMN id SET DEFAULT nextval('public.topic_readers_id_seq'::regclass);

ALTER TABLE ONLY public.topics ALTER COLUMN id SET DEFAULT nextval('public.topics_id_seq'::regclass);

ALTER TABLE ONLY public.translations ALTER COLUMN id SET DEFAULT nextval('public.translations_id_seq'::regclass);

ALTER TABLE ONLY public.trial_invitation_recipients ALTER COLUMN id SET DEFAULT nextval('public.trial_invitation_recipients_id_seq'::regclass);

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);

ALTER TABLE ONLY public.versions ALTER COLUMN id SET DEFAULT nextval('public.versions_id_seq'::regclass);

ALTER TABLE ONLY public.account_completion_proofs
    ADD CONSTRAINT account_completion_proofs_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.action_mailbox_inbound_emails
    ADD CONSTRAINT action_mailbox_inbound_emails_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.active_storage_attachments
    ADD CONSTRAINT active_storage_attachments_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.active_storage_blobs
    ADD CONSTRAINT active_storage_blobs_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.active_storage_variant_records
    ADD CONSTRAINT active_storage_variant_records_pkey PRIMARY KEY (id);

ALTER TABLE public.anonymous_ballot_choices
    ADD CONSTRAINT anonymous_ballot_choices_score_nonnegative CHECK ((score >= 0)) NOT VALID;

ALTER TABLE ONLY public.anonymous_ballots
    ADD CONSTRAINT anonymous_ballots_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.anonymous_poll_voters
    ADD CONSTRAINT anonymous_poll_voters_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.ar_internal_metadata
    ADD CONSTRAINT ar_internal_metadata_pkey PRIMARY KEY (key);

ALTER TABLE ONLY public.bookmarks
    ADD CONSTRAINT bookmarks_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.chatbots
    ADD CONSTRAINT chatbots_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.discussion_templates
    ADD CONSTRAINT discussion_templates_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.discussions
    ADD CONSTRAINT discussions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.email_suppressions
    ADD CONSTRAINT email_suppressions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.forward_email_rules
    ADD CONSTRAINT forward_email_rules_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.group_follows
    ADD CONSTRAINT group_follows_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.group_handle_redirects
    ADD CONSTRAINT group_handle_redirects_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.legacy_anonymous_vote_reasons
    ADD CONSTRAINT legacy_anonymous_vote_reasons_pkey PRIMARY KEY (anonymous_ballot_id);

ALTER TABLE ONLY public.login_tokens
    ADD CONSTRAINT login_tokens_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.member_email_aliases
    ADD CONSTRAINT member_email_aliases_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.membership_requests
    ADD CONSTRAINT membership_requests_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.memberships
    ADD CONSTRAINT memberships_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.mobile_access_tokens
    ADD CONSTRAINT mobile_access_tokens_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.mobile_authorization_codes
    ADD CONSTRAINT mobile_authorization_codes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.mobile_devices
    ADD CONSTRAINT mobile_devices_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.mobile_push_registrations
    ADD CONSTRAINT mobile_push_registrations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.mobile_refresh_tokens
    ADD CONSTRAINT mobile_refresh_tokens_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.mobile_relay_authorizations
    ADD CONSTRAINT mobile_relay_authorizations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.mobile_web_session_tickets
    ADD CONSTRAINT mobile_web_session_tickets_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.notification_deliveries
    ADD CONSTRAINT notification_deliveries_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.notification_delivery_events
    ADD CONSTRAINT notification_delivery_events_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.omniauth_identities
    ADD CONSTRAINT omniauth_identities_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.outcomes
    ADD CONSTRAINT outcomes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.partition_sequences
    ADD CONSTRAINT partition_sequences_pkey PRIMARY KEY (key, id);

ALTER TABLE ONLY public.passkey_challenges
    ADD CONSTRAINT passkey_challenges_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.passkey_credentials
    ADD CONSTRAINT passkey_credentials_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.pending_emails
    ADD CONSTRAINT pending_emails_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.pg_search_documents
    ADD CONSTRAINT pg_search_documents_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.poll_options
    ADD CONSTRAINT poll_options_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.poll_templates
    ADD CONSTRAINT poll_templates_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.polls
    ADD CONSTRAINT polls_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.push_subscriptions
    ADD CONSTRAINT push_subscriptions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.reactions
    ADD CONSTRAINT reactions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_cable_messages
    ADD CONSTRAINT solid_cable_messages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_cache_entries
    ADD CONSTRAINT solid_cache_entries_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_batch_executions
    ADD CONSTRAINT solid_queue_batch_executions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_batches
    ADD CONSTRAINT solid_queue_batches_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_blocked_executions
    ADD CONSTRAINT solid_queue_blocked_executions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_claimed_executions
    ADD CONSTRAINT solid_queue_claimed_executions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_failed_executions
    ADD CONSTRAINT solid_queue_failed_executions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_jobs
    ADD CONSTRAINT solid_queue_jobs_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_pauses
    ADD CONSTRAINT solid_queue_pauses_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_processes
    ADD CONSTRAINT solid_queue_processes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_ready_executions
    ADD CONSTRAINT solid_queue_ready_executions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_recurring_executions
    ADD CONSTRAINT solid_queue_recurring_executions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_recurring_tasks
    ADD CONSTRAINT solid_queue_recurring_tasks_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_scheduled_executions
    ADD CONSTRAINT solid_queue_scheduled_executions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.solid_queue_semaphores
    ADD CONSTRAINT solid_queue_semaphores_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.spam_rule_hits
    ADD CONSTRAINT spam_rule_hits_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.spam_rules
    ADD CONSTRAINT spam_rules_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.stance_choices
    ADD CONSTRAINT stance_choices_pkey PRIMARY KEY (id);

ALTER TABLE public.stance_choices
    ADD CONSTRAINT stance_choices_score_nonnegative CHECK ((score >= 0)) NOT VALID;

ALTER TABLE ONLY public.stances
    ADD CONSTRAINT stances_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.subscription_update_receipts
    ADD CONSTRAINT subscription_update_receipts_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.subscriptions
    ADD CONSTRAINT subscriptions_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.taggings
    ADD CONSTRAINT taggings_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.tasks_users
    ADD CONSTRAINT tasks_users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.topic_items
    ADD CONSTRAINT topic_items_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.topic_readers
    ADD CONSTRAINT topic_readers_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT topics_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.translations
    ADD CONSTRAINT translations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.trial_invitation_recipients
    ADD CONSTRAINT trial_invitation_recipients_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.versions
    ADD CONSTRAINT versions_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX active_storage_blobs_idx ON public.active_storage_blobs USING btree (id);

CREATE INDEX discussion_readers_guests ON public.topic_readers USING btree (guest) WHERE (guest = true);

CREATE INDEX groups_subscription_id_idx ON public.groups USING btree (subscription_id);

CREATE INDEX index_account_completion_proofs_on_expires_at ON public.account_completion_proofs USING btree (expires_at);

CREATE INDEX index_account_completion_proofs_on_user_id ON public.account_completion_proofs USING btree (user_id);

CREATE UNIQUE INDEX index_action_mailbox_inbound_emails_uniqueness ON public.action_mailbox_inbound_emails USING btree (message_id, message_checksum);

CREATE UNIQUE INDEX index_active_push_subscriptions_on_endpoint_digest ON public.push_subscriptions USING btree (endpoint_digest) WHERE (revoked_at IS NULL);

CREATE INDEX index_active_storage_attachments_on_blob_id ON public.active_storage_attachments USING btree (blob_id);

CREATE INDEX index_active_storage_attachments_on_group_id ON public.active_storage_attachments USING btree (group_id);

CREATE UNIQUE INDEX index_active_storage_attachments_uniqueness ON public.active_storage_attachments USING btree (record_type, record_id, name, blob_id);

CREATE UNIQUE INDEX index_active_storage_blobs_on_key ON public.active_storage_blobs USING btree (key);

CREATE UNIQUE INDEX index_active_storage_variant_records_uniqueness ON public.active_storage_variant_records USING btree (blob_id, variation_digest);

CREATE INDEX index_anonymous_ballot_choices_on_anonymous_ballot_id ON public.anonymous_ballot_choices USING btree (anonymous_ballot_id);

CREATE UNIQUE INDEX index_anonymous_ballot_choices_on_ballot_and_option ON public.anonymous_ballot_choices USING btree (anonymous_ballot_id, poll_option_id);

CREATE INDEX index_anonymous_ballot_choices_on_poll_option_id ON public.anonymous_ballot_choices USING btree (poll_option_id);

CREATE INDEX index_anonymous_ballots_on_poll_id ON public.anonymous_ballots USING btree (poll_id);

CREATE INDEX index_anonymous_poll_voters_on_inviter_id ON public.anonymous_poll_voters USING btree (inviter_id);

CREATE INDEX index_anonymous_poll_voters_on_poll_id ON public.anonymous_poll_voters USING btree (poll_id);

CREATE UNIQUE INDEX index_anonymous_poll_voters_on_poll_id_and_voter_id ON public.anonymous_poll_voters USING btree (poll_id, voter_id);

CREATE INDEX index_anonymous_poll_voters_on_voter_id ON public.anonymous_poll_voters USING btree (voter_id);

CREATE INDEX index_bookmarks_on_bookmarkable ON public.bookmarks USING btree (bookmarkable_type, bookmarkable_id);

CREATE INDEX index_bookmarks_on_discarded_at ON public.bookmarks USING btree (discarded_at);

CREATE UNIQUE INDEX index_bookmarks_on_user_and_bookmarkable ON public.bookmarks USING btree (user_id, bookmarkable_type, bookmarkable_id);

CREATE INDEX index_bookmarks_on_user_id ON public.bookmarks USING btree (user_id);

CREATE INDEX index_chatbots_on_group_id ON public.chatbots USING btree (group_id);

CREATE INDEX index_comments_on_created_at ON public.comments USING btree (created_at);

CREATE INDEX index_comments_on_discarded_by ON public.comments USING btree (discarded_by);

CREATE INDEX index_comments_on_parent_type_and_parent_id ON public.comments USING btree (parent_type, parent_id);

CREATE INDEX index_comments_on_user_id ON public.comments USING btree (user_id);

CREATE UNIQUE INDEX index_discussion_readers_on_token ON public.topic_readers USING btree (token);

CREATE INDEX index_discussion_templates_on_author_id ON public.discussion_templates USING btree (author_id);

CREATE INDEX index_discussion_templates_on_discarded_at ON public.discussion_templates USING btree (discarded_at);

CREATE INDEX index_discussion_templates_on_discarded_by ON public.discussion_templates USING btree (discarded_by);

CREATE INDEX index_discussion_templates_on_hidden_at ON public.discussion_templates USING btree (hidden_at);

CREATE INDEX index_discussion_templates_on_hider_id ON public.discussion_templates USING btree (hider_id);

CREATE INDEX index_discussions_on_author_id ON public.discussions USING btree (author_id);

CREATE INDEX index_discussions_on_created_at ON public.discussions USING btree (created_at);

CREATE INDEX index_discussions_on_discarded_at ON public.discussions USING btree (discarded_at) WHERE (discarded_at IS NULL);

CREATE INDEX index_discussions_on_discarded_by ON public.discussions USING btree (discarded_by);

CREATE UNIQUE INDEX index_discussions_on_key ON public.discussions USING btree (key);

CREATE INDEX index_discussions_on_tags ON public.discussions USING gin (tags);

CREATE INDEX index_discussions_on_template ON public.discussions USING btree (template) WHERE (template IS TRUE);

CREATE INDEX index_discussions_on_topic_id ON public.discussions USING btree (topic_id);

CREATE INDEX index_email_suppressions_active_address ON public.email_suppressions USING btree (email_sha256) WHERE (cleared_at IS NULL);

CREATE INDEX index_email_suppressions_on_cleared_by_id ON public.email_suppressions USING btree (cleared_by_id);

CREATE INDEX index_email_suppressions_on_email_sha256_and_id ON public.email_suppressions USING btree (email_sha256, id);

CREATE INDEX index_email_suppressions_on_notification_delivery_event_id ON public.email_suppressions USING btree (notification_delivery_event_id);

CREATE UNIQUE INDEX index_email_suppressions_on_source_key ON public.email_suppressions USING btree (source_key);

CREATE INDEX index_group_follows_on_group_id ON public.group_follows USING btree (group_id);

CREATE INDEX index_group_follows_on_user_id ON public.group_follows USING btree (user_id);

CREATE UNIQUE INDEX index_group_follows_on_user_id_and_group_id ON public.group_follows USING btree (user_id, group_id);

CREATE INDEX index_group_handle_redirects_on_group_id ON public.group_handle_redirects USING btree (group_id);

CREATE UNIQUE INDEX index_group_handle_redirects_on_group_id_and_handle ON public.group_handle_redirects USING btree (group_id, handle);

CREATE UNIQUE INDEX index_group_handle_redirects_on_handle ON public.group_handle_redirects USING btree (handle);

CREATE INDEX index_groups_on_created_at ON public.groups USING btree (created_at);

CREATE INDEX index_groups_on_creator_id ON public.groups USING btree (creator_id);

CREATE INDEX index_groups_on_discarded_at ON public.groups USING btree (discarded_at) WHERE (discarded_at IS NULL);

CREATE INDEX index_groups_on_full_name ON public.groups USING btree (full_name);

CREATE UNIQUE INDEX index_groups_on_handle ON public.groups USING btree (handle);

CREATE UNIQUE INDEX index_groups_on_id_and_is_root ON public.groups USING btree (id, is_root);

CREATE UNIQUE INDEX index_groups_on_key ON public.groups USING btree (key);

CREATE INDEX index_groups_on_name ON public.groups USING btree (name);

CREATE INDEX index_groups_on_parent_id ON public.groups USING btree (parent_id);

CREATE UNIQUE INDEX index_groups_on_token ON public.groups USING btree (token);

CREATE INDEX index_inbound_emails_on_message_id_and_id ON public.action_mailbox_inbound_emails USING btree (message_id, id);

CREATE INDEX index_login_tokens_on_token ON public.login_tokens USING btree (token);

CREATE UNIQUE INDEX index_member_email_aliases_on_email_and_group_id ON public.member_email_aliases USING btree (email, group_id);

CREATE INDEX index_membership_requests_on_group_id ON public.membership_requests USING btree (group_id);

CREATE INDEX index_membership_requests_on_requestor_id ON public.membership_requests USING btree (requestor_id);

CREATE INDEX index_membership_requests_on_responder_id ON public.membership_requests USING btree (responder_id);

CREATE INDEX index_memberships_on_created_at ON public.memberships USING btree (created_at);

CREATE UNIQUE INDEX index_memberships_on_group_id_and_user_id ON public.memberships USING btree (group_id, user_id);

CREATE INDEX index_memberships_on_inviter_id ON public.memberships USING btree (inviter_id);

CREATE INDEX index_memberships_on_revoked_at_and_id_for_relay ON public.memberships USING btree (revoked_at, id) WHERE (revoked_at IS NOT NULL);

CREATE INDEX index_memberships_on_revoker_id ON public.memberships USING btree (revoker_id);

CREATE UNIQUE INDEX index_memberships_on_token ON public.memberships USING btree (token);

CREATE INDEX index_memberships_on_user_id_and_volume_email ON public.memberships USING btree (user_id, volume_email);

CREATE INDEX index_memberships_on_user_id_and_volume_push ON public.memberships USING btree (user_id, volume_push);

CREATE INDEX index_memberships_on_volume_email ON public.memberships USING btree (volume_email);

CREATE INDEX index_memberships_on_volume_push ON public.memberships USING btree (volume_push);

CREATE INDEX index_mobile_access_tokens_on_expires_at ON public.mobile_access_tokens USING btree (expires_at);

CREATE INDEX index_mobile_access_tokens_on_mobile_device_id ON public.mobile_access_tokens USING btree (mobile_device_id);

CREATE UNIQUE INDEX index_mobile_access_tokens_on_token_digest ON public.mobile_access_tokens USING btree (token_digest);

CREATE INDEX index_mobile_authorization_codes_on_expires_at ON public.mobile_authorization_codes USING btree (expires_at);

CREATE UNIQUE INDEX index_mobile_authorization_codes_on_token_digest ON public.mobile_authorization_codes USING btree (token_digest);

CREATE INDEX index_mobile_authorization_codes_on_user_id ON public.mobile_authorization_codes USING btree (user_id);

CREATE INDEX index_mobile_devices_on_user_id ON public.mobile_devices USING btree (user_id);

CREATE INDEX index_mobile_devices_on_user_id_and_revoked_at ON public.mobile_devices USING btree (user_id, revoked_at);

CREATE UNIQUE INDEX index_mobile_push_registrations_on_mobile_device_id ON public.mobile_push_registrations USING btree (mobile_device_id);

CREATE UNIQUE INDEX index_mobile_push_registrations_on_registration_id ON public.mobile_push_registrations USING btree (registration_id);

CREATE INDEX index_mobile_refresh_tokens_on_family_id_and_revoked_at ON public.mobile_refresh_tokens USING btree (family_id, revoked_at);

CREATE INDEX index_mobile_refresh_tokens_on_mobile_device_id ON public.mobile_refresh_tokens USING btree (mobile_device_id);

CREATE INDEX index_mobile_refresh_tokens_on_parent_id ON public.mobile_refresh_tokens USING btree (parent_id);

CREATE UNIQUE INDEX index_mobile_refresh_tokens_on_token_digest ON public.mobile_refresh_tokens USING btree (token_digest);

CREATE INDEX index_mobile_relay_authorizations_on_expires_at ON public.mobile_relay_authorizations USING btree (expires_at);

CREATE INDEX index_mobile_relay_authorizations_on_mobile_device_id ON public.mobile_relay_authorizations USING btree (mobile_device_id);

CREATE UNIQUE INDEX index_mobile_relay_authorizations_on_token_digest ON public.mobile_relay_authorizations USING btree (token_digest);

CREATE INDEX index_mobile_web_session_tickets_on_expires_at ON public.mobile_web_session_tickets USING btree (expires_at);

CREATE INDEX index_mobile_web_session_tickets_on_mobile_device_id ON public.mobile_web_session_tickets USING btree (mobile_device_id);

CREATE UNIQUE INDEX index_mobile_web_session_tickets_on_token_digest ON public.mobile_web_session_tickets USING btree (token_digest);

CREATE UNIQUE INDEX index_notification_deliveries_on_email_message_id ON public.notification_deliveries USING btree (email_message_id);

CREATE INDEX index_notification_deliveries_on_email_send_cursor ON public.notification_deliveries USING btree (delivered_at, id) WHERE (((channel)::text = 'email'::text) AND (delivered_at IS NOT NULL));

CREATE UNIQUE INDEX index_notification_deliveries_on_identity ON public.notification_deliveries USING btree (notification_id, channel, recipient_type, recipient_id);

CREATE INDEX index_notification_deliveries_on_notification_id ON public.notification_deliveries USING btree (notification_id);

CREATE INDEX index_notification_deliveries_on_provider_message_id ON public.notification_deliveries USING btree (provider_message_id);

CREATE INDEX index_notification_deliveries_on_recipient ON public.notification_deliveries USING btree (recipient_type, recipient_id);

CREATE INDEX index_notification_delivery_events_on_notification_delivery_id ON public.notification_delivery_events USING btree (notification_delivery_id);

CREATE UNIQUE INDEX index_notification_delivery_events_on_source ON public.notification_delivery_events USING btree (notification_delivery_id, source_key);

CREATE INDEX index_notifications_on_actor_id ON public.notifications USING btree (actor_id);

CREATE INDEX index_notifications_on_pending_delivery_resolution ON public.notifications USING btree (id) WHERE (deliveries_generated_at IS NULL);

CREATE INDEX index_notifications_on_recipient_user_ids ON public.notifications USING gin (recipient_user_ids);

CREATE INDEX index_notifications_on_subject ON public.notifications USING btree (subject_type, subject_id);

CREATE UNIQUE INDEX index_omniauth_identities_on_identity_type_and_uid ON public.omniauth_identities USING btree (identity_type, uid);

CREATE INDEX index_outcomes_on_author_id ON public.outcomes USING btree (author_id);

CREATE INDEX index_outcomes_on_created_at ON public.outcomes USING btree (created_at);

CREATE INDEX index_outcomes_on_poll_id ON public.outcomes USING btree (poll_id);

CREATE UNIQUE INDEX index_passkey_challenges_on_challenge_digest ON public.passkey_challenges USING btree (challenge_digest);

CREATE INDEX index_passkey_challenges_on_expires_at ON public.passkey_challenges USING btree (expires_at);

CREATE INDEX index_passkey_challenges_on_user_id ON public.passkey_challenges USING btree (user_id);

CREATE UNIQUE INDEX index_passkey_credentials_on_external_id ON public.passkey_credentials USING btree (external_id);

CREATE INDEX index_passkey_credentials_on_user_id ON public.passkey_credentials USING btree (user_id);

CREATE INDEX index_pending_emails_on_group_id ON public.pending_emails USING btree (group_id);

CREATE UNIQUE INDEX index_pending_emails_on_inbound_email_id ON public.pending_emails USING btree (inbound_email_id);

CREATE INDEX index_personas_on_email ON public.omniauth_identities USING btree (email);

CREATE INDEX index_personas_on_user_id ON public.omniauth_identities USING btree (user_id);

CREATE INDEX index_pg_search_documents_on_author_id ON public.pg_search_documents USING btree (author_id);

CREATE INDEX index_pg_search_documents_on_discussion_id ON public.pg_search_documents USING btree (discussion_id);

CREATE INDEX index_pg_search_documents_on_group_id ON public.pg_search_documents USING btree (group_id);

CREATE INDEX index_pg_search_documents_on_poll_id ON public.pg_search_documents USING btree (poll_id);

CREATE INDEX index_pg_search_documents_on_searchable ON public.pg_search_documents USING btree (searchable_type, searchable_id);

CREATE INDEX index_pg_search_documents_on_tags ON public.pg_search_documents USING gin (tags);

CREATE INDEX index_pg_search_documents_on_topic_id ON public.pg_search_documents USING btree (topic_id);

CREATE UNIQUE INDEX index_pg_search_words_on_word ON public.pg_search_words USING btree (word);

CREATE INDEX index_pg_search_words_on_word_trigram ON public.pg_search_words USING gin (word public.gin_trgm_ops);

CREATE INDEX index_poll_options_on_poll_id ON public.poll_options USING btree (poll_id);

CREATE INDEX index_poll_templates_on_discarded_at ON public.poll_templates USING btree (discarded_at);

CREATE INDEX index_poll_templates_on_discarded_by ON public.poll_templates USING btree (discarded_by);

CREATE INDEX index_poll_templates_on_hidden_at ON public.poll_templates USING btree (hidden_at);

CREATE INDEX index_poll_templates_on_hider_id ON public.poll_templates USING btree (hider_id);

CREATE INDEX index_polls_on_author_id ON public.polls USING btree (author_id);

CREATE INDEX index_polls_on_closed_at_and_closing_at ON public.polls USING btree (closed_at, closing_at);

CREATE INDEX index_polls_on_closed_at_and_topic_id ON public.polls USING btree (closed_at, topic_id);

CREATE INDEX index_polls_on_discarded_by ON public.polls USING btree (discarded_by);

CREATE UNIQUE INDEX index_polls_on_key ON public.polls USING btree (key);

CREATE INDEX index_polls_on_pending_opening_at ON public.polls USING btree (opening_at) WHERE ((opened_at IS NULL) AND (closed_at IS NULL) AND (discarded_at IS NULL) AND (opening_at IS NOT NULL));

CREATE INDEX index_polls_on_tags ON public.polls USING gin (tags);

CREATE INDEX index_polls_on_topic_id ON public.polls USING btree (topic_id);

CREATE INDEX index_push_subscriptions_on_session_id ON public.push_subscriptions USING btree (session_id);

CREATE INDEX index_push_subscriptions_on_user_id ON public.push_subscriptions USING btree (user_id);

CREATE INDEX index_push_subscriptions_on_user_id_and_revoked_at ON public.push_subscriptions USING btree (user_id, revoked_at);

CREATE INDEX index_reactions_on_created_at ON public.reactions USING btree (created_at);

CREATE INDEX index_reactions_on_reactable_id_and_reactable_type ON public.reactions USING btree (reactable_id, reactable_type);

CREATE INDEX index_reactions_on_user_id ON public.reactions USING btree (user_id);

CREATE INDEX index_sessions_on_revoked_at ON public.sessions USING btree (revoked_at) WHERE (revoked_at IS NOT NULL);

CREATE INDEX index_sessions_on_user_id ON public.sessions USING btree (user_id);

CREATE INDEX index_solid_cable_messages_on_channel ON public.solid_cable_messages USING btree (channel);

CREATE INDEX index_solid_cable_messages_on_channel_hash ON public.solid_cable_messages USING btree (channel_hash);

CREATE INDEX index_solid_cable_messages_on_created_at ON public.solid_cable_messages USING btree (created_at);

CREATE INDEX index_solid_cache_entries_on_byte_size ON public.solid_cache_entries USING btree (byte_size);

CREATE UNIQUE INDEX index_solid_cache_entries_on_key_hash ON public.solid_cache_entries USING btree (key_hash);

CREATE INDEX index_solid_cache_entries_on_key_hash_and_byte_size ON public.solid_cache_entries USING btree (key_hash, byte_size);

CREATE INDEX index_solid_queue_batch_executions_on_batch_id ON public.solid_queue_batch_executions USING btree (batch_id);

CREATE UNIQUE INDEX index_solid_queue_batch_executions_on_job_id ON public.solid_queue_batch_executions USING btree (job_id);

CREATE UNIQUE INDEX index_solid_queue_batches_on_active_job_batch_id ON public.solid_queue_batches USING btree (active_job_batch_id);

CREATE INDEX index_solid_queue_batches_on_finished_at ON public.solid_queue_batches USING btree (finished_at);

CREATE INDEX index_solid_queue_blocked_executions_for_maintenance ON public.solid_queue_blocked_executions USING btree (expires_at, concurrency_key);

CREATE INDEX index_solid_queue_blocked_executions_for_release ON public.solid_queue_blocked_executions USING btree (concurrency_key, priority, job_id);

CREATE UNIQUE INDEX index_solid_queue_blocked_executions_on_job_id ON public.solid_queue_blocked_executions USING btree (job_id);

CREATE UNIQUE INDEX index_solid_queue_claimed_executions_on_job_id ON public.solid_queue_claimed_executions USING btree (job_id);

CREATE INDEX index_solid_queue_claimed_executions_on_process_id_and_job_id ON public.solid_queue_claimed_executions USING btree (process_id, job_id);

CREATE INDEX index_solid_queue_dispatch_all ON public.solid_queue_scheduled_executions USING btree (scheduled_at, priority, job_id);

CREATE UNIQUE INDEX index_solid_queue_failed_executions_on_job_id ON public.solid_queue_failed_executions USING btree (job_id);

CREATE INDEX index_solid_queue_jobs_for_alerting ON public.solid_queue_jobs USING btree (scheduled_at, finished_at);

CREATE INDEX index_solid_queue_jobs_for_filtering ON public.solid_queue_jobs USING btree (queue_name, finished_at);

CREATE INDEX index_solid_queue_jobs_on_active_job_id ON public.solid_queue_jobs USING btree (active_job_id);

CREATE INDEX index_solid_queue_jobs_on_batch_id ON public.solid_queue_jobs USING btree (batch_id);

CREATE INDEX index_solid_queue_jobs_on_class_name ON public.solid_queue_jobs USING btree (class_name);

CREATE INDEX index_solid_queue_jobs_on_finished_at ON public.solid_queue_jobs USING btree (finished_at);

CREATE UNIQUE INDEX index_solid_queue_pauses_on_queue_name ON public.solid_queue_pauses USING btree (queue_name);

CREATE INDEX index_solid_queue_poll_all ON public.solid_queue_ready_executions USING btree (priority, job_id);

CREATE INDEX index_solid_queue_poll_by_queue ON public.solid_queue_ready_executions USING btree (queue_name, priority, job_id);

CREATE INDEX index_solid_queue_processes_on_last_heartbeat_at ON public.solid_queue_processes USING btree (last_heartbeat_at);

CREATE UNIQUE INDEX index_solid_queue_processes_on_name_and_supervisor_id ON public.solid_queue_processes USING btree (name, supervisor_id);

CREATE INDEX index_solid_queue_processes_on_supervisor_id ON public.solid_queue_processes USING btree (supervisor_id);

CREATE UNIQUE INDEX index_solid_queue_ready_executions_on_job_id ON public.solid_queue_ready_executions USING btree (job_id);

CREATE UNIQUE INDEX index_solid_queue_recurring_executions_on_job_id ON public.solid_queue_recurring_executions USING btree (job_id);

CREATE UNIQUE INDEX index_solid_queue_recurring_executions_on_task_key_and_run_at ON public.solid_queue_recurring_executions USING btree (task_key, run_at);

CREATE UNIQUE INDEX index_solid_queue_recurring_tasks_on_key ON public.solid_queue_recurring_tasks USING btree (key);

CREATE INDEX index_solid_queue_recurring_tasks_on_static ON public.solid_queue_recurring_tasks USING btree (static);

CREATE UNIQUE INDEX index_solid_queue_scheduled_executions_on_job_id ON public.solid_queue_scheduled_executions USING btree (job_id);

CREATE INDEX index_solid_queue_semaphores_on_expires_at ON public.solid_queue_semaphores USING btree (expires_at);

CREATE UNIQUE INDEX index_solid_queue_semaphores_on_key ON public.solid_queue_semaphores USING btree (key);

CREATE INDEX index_solid_queue_semaphores_on_key_and_value ON public.solid_queue_semaphores USING btree (key, value);

CREATE INDEX index_spam_rule_hits_on_bucket_at ON public.spam_rule_hits USING btree (bucket_at);

CREATE INDEX index_spam_rule_hits_on_spam_rule_id ON public.spam_rule_hits USING btree (spam_rule_id);

CREATE UNIQUE INDEX index_spam_rule_hits_on_spam_rule_id_and_action_and_bucket_at ON public.spam_rule_hits USING btree (spam_rule_id, action, bucket_at);

CREATE INDEX index_stance_choices_on_poll_option_id ON public.stance_choices USING btree (poll_option_id);

CREATE INDEX index_stance_choices_on_stance_id ON public.stance_choices USING btree (stance_id);

CREATE UNIQUE INDEX index_stance_choices_on_stance_id_and_poll_option_id ON public.stance_choices USING btree (stance_id, poll_option_id);

CREATE INDEX index_stances_on_cast_at_and_id_for_relay ON public.stances USING btree (cast_at, id) WHERE ((cast_at IS NOT NULL) AND (redacted_at IS NULL));

CREATE INDEX index_stances_on_created_at ON public.stances USING btree (created_at);

CREATE INDEX index_stances_on_inviter_id ON public.stances USING btree (inviter_id);

CREATE INDEX index_stances_on_participant_id ON public.stances USING btree (participant_id);

CREATE INDEX index_stances_on_poll_id ON public.stances USING btree (poll_id);

CREATE INDEX index_stances_on_poll_id_and_cast_at ON public.stances USING btree (poll_id NULLS FIRST, cast_at NULLS FIRST);

CREATE UNIQUE INDEX index_stances_on_poll_id_and_participant_id_and_latest ON public.stances USING btree (poll_id, participant_id, latest) WHERE (latest = true);

CREATE INDEX index_stances_on_redactor_id ON public.stances USING btree (redactor_id);

CREATE INDEX index_stances_on_revoker_id ON public.stances USING btree (revoker_id);

CREATE UNIQUE INDEX index_stances_on_token ON public.stances USING btree (token);

CREATE UNIQUE INDEX index_subscription_update_receipts_on_event_id ON public.subscription_update_receipts USING btree (event_id);

CREATE INDEX index_subscription_update_receipts_on_subscription_id ON public.subscription_update_receipts USING btree (subscription_id);

CREATE UNIQUE INDEX index_subscriptions_on_billing_service_id ON public.subscriptions USING btree (billing_service_subscription_id) WHERE (billing_service_subscription_id IS NOT NULL);

CREATE INDEX index_subscriptions_on_owner_id ON public.subscriptions USING btree (owner_id);

CREATE INDEX index_subscriptions_on_plan ON public.subscriptions USING btree (plan);

CREATE INDEX index_subscriptions_on_trial_expiry_for_relay ON public.subscriptions USING btree (expires_at, id) WHERE (((plan)::text = 'trial'::text) AND (expires_at IS NOT NULL));

CREATE INDEX index_taggings_on_tag_id ON public.taggings USING btree (tag_id);

CREATE INDEX index_taggings_on_taggable_type_and_taggable_id ON public.taggings USING btree (taggable_type, taggable_id);

CREATE INDEX index_tags_on_group_id ON public.tags USING btree (group_id);

CREATE UNIQUE INDEX index_tags_on_group_id_and_name ON public.tags USING btree (group_id, name);

CREATE INDEX index_tags_on_name ON public.tags USING btree (name);

CREATE INDEX index_tags_on_used_group_ids ON public.tags USING gin (used_group_ids);

CREATE INDEX index_tasks_on_author_id ON public.tasks USING btree (author_id);

CREATE INDEX index_tasks_on_discarded_at ON public.tasks USING btree (discarded_at);

CREATE INDEX index_tasks_on_doer_id ON public.tasks USING btree (doer_id);

CREATE INDEX index_tasks_on_due_on ON public.tasks USING btree (due_on);

CREATE INDEX index_tasks_on_record_type_and_record_id ON public.tasks USING btree (record_type, record_id);

CREATE INDEX index_tasks_on_remind_at ON public.tasks USING btree (remind_at);

CREATE INDEX index_tasks_users_on_task_id ON public.tasks_users USING btree (task_id);

CREATE INDEX index_tasks_users_on_user_id ON public.tasks_users USING btree (user_id);

CREATE INDEX index_topic_items_on_created_at ON public.topic_items USING btree (created_at);

CREATE UNIQUE INDEX index_topic_items_on_id_and_topic_id ON public.topic_items USING btree (id, topic_id);

CREATE INDEX index_topic_items_on_itemable_id_and_kind ON public.topic_items USING btree (itemable_id, kind);

CREATE INDEX index_topic_items_on_itemable_type_and_itemable_id ON public.topic_items USING btree (itemable_type, itemable_id);

CREATE INDEX index_topic_items_on_parent_id ON public.topic_items USING btree (parent_id);

CREATE INDEX index_topic_items_on_parent_id_and_topic_id ON public.topic_items USING btree (parent_id, topic_id) WHERE (topic_id IS NOT NULL);

CREATE INDEX index_topic_items_on_position_key ON public.topic_items USING btree (position_key);

CREATE INDEX index_topic_items_on_tags ON public.topic_items USING gin (tags);

CREATE INDEX index_topic_items_on_topic_id ON public.topic_items USING btree (topic_id);

CREATE UNIQUE INDEX index_topic_items_on_topic_id_and_sequence_id ON public.topic_items USING btree (topic_id, sequence_id);

CREATE INDEX index_topic_items_on_topic_id_depth_sequence_id ON public.topic_items USING btree (topic_id, depth, sequence_id);

CREATE INDEX index_topic_items_on_topic_id_sequence_id_pinned ON public.topic_items USING btree (topic_id, sequence_id) WHERE (pinned = true);

CREATE UNIQUE INDEX index_topic_items_on_unique_discussion_root ON public.topic_items USING btree (itemable_type, itemable_id, kind) WHERE (((itemable_type)::text = 'Discussion'::text) AND ((kind)::text = 'new_discussion'::text));

CREATE UNIQUE INDEX index_topic_items_on_unique_poll_root ON public.topic_items USING btree (itemable_type, itemable_id, kind) WHERE (((itemable_type)::text = 'Poll'::text) AND ((kind)::text = 'poll_created'::text));

CREATE INDEX index_topic_items_on_user_id ON public.topic_items USING btree (user_id);

CREATE INDEX index_topic_items_with_tags_on_topic_id ON public.topic_items USING btree (topic_id) WHERE (cardinality(tags) > 0);

CREATE INDEX index_topic_readers_guest_user_id ON public.topic_readers USING btree (user_id, topic_id) WHERE (guest = true);

CREATE INDEX index_topic_readers_on_revoker_id ON public.topic_readers USING btree (revoker_id);

CREATE UNIQUE INDEX index_topic_readers_on_topic_id_and_user_id ON public.topic_readers USING btree (topic_id, user_id);

CREATE INDEX index_topic_readers_on_user_id ON public.topic_readers USING btree (user_id);

CREATE INDEX index_topics_on_discarded_at_null ON public.topics USING btree (discarded_at) WHERE (discarded_at IS NULL);

CREATE INDEX index_topics_on_discarded_by ON public.topics USING btree (discarded_by);

CREATE INDEX index_topics_on_group_id ON public.topics USING btree (group_id);

CREATE INDEX index_topics_on_group_last_activity_inbox ON public.topics USING btree (group_id, last_activity_at DESC) WHERE (discarded_at IS NULL);

CREATE INDEX index_topics_on_last_activity_at ON public.topics USING btree (last_activity_at DESC);

CREATE INDEX index_topics_on_locked_at ON public.topics USING btree (locked_at);

CREATE INDEX index_topics_on_locker_id ON public.topics USING btree (locker_id);

CREATE INDEX index_topics_on_tags ON public.topics USING gin (tags);

CREATE UNIQUE INDEX index_topics_on_topicable_type_and_topicable_id ON public.topics USING btree (topicable_type, topicable_id);

CREATE INDEX index_translations_on_language ON public.translations USING btree (language);

CREATE INDEX index_translations_on_translatable_type_and_translatable_id ON public.translations USING btree (translatable_type, translatable_id);

CREATE UNIQUE INDEX index_trial_invitation_recipients_identity ON public.trial_invitation_recipients USING btree (ip_address, actor_id, recipient_id);

CREATE INDEX index_trial_invitation_recipients_on_invited_at ON public.trial_invitation_recipients USING btree (invited_at);

CREATE INDEX index_trial_invitation_recipients_on_ip_address_and_invited_at ON public.trial_invitation_recipients USING btree (ip_address, invited_at);

CREATE INDEX index_users_on_api_key ON public.users USING btree (api_key);

CREATE INDEX index_users_on_created_at_and_id_for_abuse_review ON public.users USING btree (created_at, id);

CREATE INDEX index_users_on_deactivator_id ON public.users USING btree (deactivator_id);

CREATE UNIQUE INDEX index_users_on_email ON public.users USING btree (email);

CREATE INDEX index_users_on_email_verified ON public.users USING btree (email_verified);

CREATE UNIQUE INDEX index_users_on_key ON public.users USING btree (key);

CREATE UNIQUE INDEX index_users_on_unsubscribe_token ON public.users USING btree (unsubscribe_token);

CREATE UNIQUE INDEX index_users_on_username ON public.users USING btree (username);

CREATE UNIQUE INDEX index_users_on_webauthn_id ON public.users USING btree (webauthn_id);

CREATE INDEX index_versions_on_item_type_and_item_id ON public.versions USING btree (item_type, item_id);

CREATE INDEX inviter_id_not_null ON public.topic_readers USING btree (inviter_id) WHERE (inviter_id IS NOT NULL);

CREATE INDEX pg_search_documents_searchable_index ON public.pg_search_documents USING gin (ts_content);

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT fk_groups_single_level_subgroup FOREIGN KEY (parent_id, parent_is_root) REFERENCES public.groups(id, is_root);

ALTER TABLE ONLY public.spam_rule_hits
    ADD CONSTRAINT fk_rails_05fc51451f FOREIGN KEY (spam_rule_id) REFERENCES public.spam_rules(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_rails_06a39bb8cc FOREIGN KEY (actor_id) REFERENCES public.users(id) ON DELETE SET NULL;

ALTER TABLE ONLY public.anonymous_poll_voters
    ADD CONSTRAINT fk_rails_12044b32f9 FOREIGN KEY (poll_id) REFERENCES public.polls(id);

ALTER TABLE ONLY public.trial_invitation_recipients
    ADD CONSTRAINT fk_rails_122c3d29a2 FOREIGN KEY (actor_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.mobile_relay_authorizations
    ADD CONSTRAINT fk_rails_1dc193b29c FOREIGN KEY (mobile_device_id) REFERENCES public.mobile_devices(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.mobile_push_registrations
    ADD CONSTRAINT fk_rails_20073d55f7 FOREIGN KEY (mobile_device_id) REFERENCES public.mobile_devices(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.passkey_credentials
    ADD CONSTRAINT fk_rails_225250860c FOREIGN KEY (user_id) REFERENCES public.users(id);

ALTER TABLE ONLY public.solid_queue_recurring_executions
    ADD CONSTRAINT fk_rails_318a5533ed FOREIGN KEY (job_id) REFERENCES public.solid_queue_jobs(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.pending_emails
    ADD CONSTRAINT fk_rails_37076ce2c9 FOREIGN KEY (inbound_email_id) REFERENCES public.action_mailbox_inbound_emails(id);

ALTER TABLE ONLY public.solid_queue_failed_executions
    ADD CONSTRAINT fk_rails_39bbc7a631 FOREIGN KEY (job_id) REFERENCES public.solid_queue_jobs(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.stances
    ADD CONSTRAINT fk_rails_40f679246d FOREIGN KEY (revoker_id) REFERENCES public.users(id) ON DELETE SET NULL;

ALTER TABLE ONLY public.tasks_users
    ADD CONSTRAINT fk_rails_438e319f2f FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.push_subscriptions
    ADD CONSTRAINT fk_rails_43d43720fc FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.outcomes
    ADD CONSTRAINT fk_rails_4960b8adaa FOREIGN KEY (poll_id) REFERENCES public.polls(id);

ALTER TABLE ONLY public.solid_queue_blocked_executions
    ADD CONSTRAINT fk_rails_4cd34e2228 FOREIGN KEY (job_id) REFERENCES public.solid_queue_jobs(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.anonymous_ballot_choices
    ADD CONSTRAINT fk_rails_578b8af196 FOREIGN KEY (poll_option_id) REFERENCES public.poll_options(id);

ALTER TABLE ONLY public.discussions
    ADD CONSTRAINT fk_rails_5e49c0df8c FOREIGN KEY (topic_id) REFERENCES public.topics(id) DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE ONLY public.mobile_access_tokens
    ADD CONSTRAINT fk_rails_5e75bf068b FOREIGN KEY (mobile_device_id) REFERENCES public.mobile_devices(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.stance_choices
    ADD CONSTRAINT fk_rails_5fe105c062 FOREIGN KEY (poll_option_id) REFERENCES public.poll_options(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.topic_readers
    ADD CONSTRAINT fk_rails_6e2f364f27 FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.membership_requests
    ADD CONSTRAINT fk_rails_70e5b09e14 FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.group_handle_redirects
    ADD CONSTRAINT fk_rails_7456fe8392 FOREIGN KEY (group_id) REFERENCES public.groups(id);

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT fk_rails_758836b4f0 FOREIGN KEY (user_id) REFERENCES public.users(id);

ALTER TABLE ONLY public.solid_queue_batch_executions
    ADD CONSTRAINT fk_rails_7c5e073422 FOREIGN KEY (batch_id) REFERENCES public.solid_queue_batches(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.topic_readers
    ADD CONSTRAINT fk_rails_7ce9e0ce13 FOREIGN KEY (topic_id) REFERENCES public.topics(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.email_suppressions
    ADD CONSTRAINT fk_rails_7d37329afe FOREIGN KEY (cleared_by_id) REFERENCES public.users(id) ON DELETE SET NULL;

ALTER TABLE ONLY public.email_suppressions
    ADD CONSTRAINT fk_rails_7dcbb9c021 FOREIGN KEY (notification_delivery_event_id) REFERENCES public.notification_delivery_events(id) ON DELETE SET NULL;

ALTER TABLE ONLY public.notification_deliveries
    ADD CONSTRAINT fk_rails_7f4a998ea3 FOREIGN KEY (notification_id) REFERENCES public.notifications(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.solid_queue_ready_executions
    ADD CONSTRAINT fk_rails_81fcbd66af FOREIGN KEY (job_id) REFERENCES public.solid_queue_jobs(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.stances
    ADD CONSTRAINT fk_rails_832b51543a FOREIGN KEY (redactor_id) REFERENCES public.users(id) ON DELETE SET NULL;

ALTER TABLE ONLY public.push_subscriptions
    ADD CONSTRAINT fk_rails_855264bb6c FOREIGN KEY (session_id) REFERENCES public.sessions(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.polls
    ADD CONSTRAINT fk_rails_861a79e923 FOREIGN KEY (topic_id) REFERENCES public.topics(id) DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE ONLY public.mobile_web_session_tickets
    ADD CONSTRAINT fk_rails_8a8d6cc141 FOREIGN KEY (mobile_device_id) REFERENCES public.mobile_devices(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.memberships
    ADD CONSTRAINT fk_rails_99326fb65d FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.active_storage_variant_records
    ADD CONSTRAINT fk_rails_993965df05 FOREIGN KEY (blob_id) REFERENCES public.active_storage_blobs(id);

ALTER TABLE ONLY public.mobile_authorization_codes
    ADD CONSTRAINT fk_rails_9c7f847bfa FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.solid_queue_claimed_executions
    ADD CONSTRAINT fk_rails_9cfe4d4944 FOREIGN KEY (job_id) REFERENCES public.solid_queue_jobs(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.tasks_users
    ADD CONSTRAINT fk_rails_9e3a948eaa FOREIGN KEY (task_id) REFERENCES public.tasks(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.taggings
    ADD CONSTRAINT fk_rails_9fcd2e236b FOREIGN KEY (tag_id) REFERENCES public.tags(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.group_follows
    ADD CONSTRAINT fk_rails_9fd22d074b FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.poll_options
    ADD CONSTRAINT fk_rails_aa85becb42 FOREIGN KEY (poll_id) REFERENCES public.polls(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.memberships
    ADD CONSTRAINT fk_rails_aaf389f138 FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.mobile_refresh_tokens
    ADD CONSTRAINT fk_rails_adff67f1ba FOREIGN KEY (parent_id) REFERENCES public.mobile_refresh_tokens(id) ON DELETE SET NULL;

ALTER TABLE ONLY public.group_follows
    ADD CONSTRAINT fk_rails_aeb40bad55 FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.anonymous_ballots
    ADD CONSTRAINT fk_rails_b45f0b3905 FOREIGN KEY (poll_id) REFERENCES public.polls(id);

ALTER TABLE ONLY public.anonymous_poll_voters
    ADD CONSTRAINT fk_rails_b858c4c7e5 FOREIGN KEY (inviter_id) REFERENCES public.users(id);

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT fk_rails_b9878cbc48 FOREIGN KEY (group_id) REFERENCES public.groups(id);

ALTER TABLE ONLY public.solid_queue_batch_executions
    ADD CONSTRAINT fk_rails_bc9f981155 FOREIGN KEY (job_id) REFERENCES public.solid_queue_jobs(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.trial_invitation_recipients
    ADD CONSTRAINT fk_rails_bd9e4b490b FOREIGN KEY (recipient_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT fk_rails_be49f097d1 FOREIGN KEY (parent_id) REFERENCES public.groups(id);

ALTER TABLE ONLY public.account_completion_proofs
    ADD CONSTRAINT fk_rails_c02f9f4bf7 FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.active_storage_attachments
    ADD CONSTRAINT fk_rails_c3b3935057 FOREIGN KEY (blob_id) REFERENCES public.active_storage_blobs(id);

ALTER TABLE ONLY public.solid_queue_scheduled_executions
    ADD CONSTRAINT fk_rails_c4316f352d FOREIGN KEY (job_id) REFERENCES public.solid_queue_jobs(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.passkey_challenges
    ADD CONSTRAINT fk_rails_c58e583d84 FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.subscription_update_receipts
    ADD CONSTRAINT fk_rails_d0a7a8c82a FOREIGN KEY (subscription_id) REFERENCES public.subscriptions(id);

ALTER TABLE ONLY public.legacy_anonymous_vote_reasons
    ADD CONSTRAINT fk_rails_d23ea4dbb3 FOREIGN KEY (anonymous_ballot_id) REFERENCES public.anonymous_ballots(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.notification_delivery_events
    ADD CONSTRAINT fk_rails_d464e6f3e4 FOREIGN KEY (notification_delivery_id) REFERENCES public.notification_deliveries(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.anonymous_ballot_choices
    ADD CONSTRAINT fk_rails_d5a8965893 FOREIGN KEY (anonymous_ballot_id) REFERENCES public.anonymous_ballots(id);

ALTER TABLE ONLY public.pending_emails
    ADD CONSTRAINT fk_rails_d651d810f5 FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.anonymous_poll_voters
    ADD CONSTRAINT fk_rails_d9f669b313 FOREIGN KEY (voter_id) REFERENCES public.users(id);

ALTER TABLE ONLY public.stance_choices
    ADD CONSTRAINT fk_rails_dd52b0505a FOREIGN KEY (stance_id) REFERENCES public.stances(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT fk_rails_e94b1a3eb4 FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.topic_items
    ADD CONSTRAINT fk_rails_e9a6de7263 FOREIGN KEY (topic_id) REFERENCES public.topics(id);

ALTER TABLE ONLY public.stances
    ADD CONSTRAINT fk_rails_f580fa74ad FOREIGN KEY (poll_id) REFERENCES public.polls(id);

ALTER TABLE ONLY public.mobile_devices
    ADD CONSTRAINT fk_rails_f61b19cc7b FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.stances
    ADD CONSTRAINT fk_rails_fa118674d1 FOREIGN KEY (inviter_id) REFERENCES public.users(id) ON DELETE SET NULL;

ALTER TABLE ONLY public.mobile_refresh_tokens
    ADD CONSTRAINT fk_rails_fa71c6f511 FOREIGN KEY (mobile_device_id) REFERENCES public.mobile_devices(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.stances
    ADD CONSTRAINT fk_rails_fa7c1370b3 FOREIGN KEY (participant_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.topic_items
    ADD CONSTRAINT topic_items_parent_same_topic FOREIGN KEY (parent_id, topic_id) REFERENCES public.topic_items(id, topic_id) ON DELETE CASCADE DEFERRABLE;

