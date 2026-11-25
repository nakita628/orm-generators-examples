# activerecord/mastodon

The tables Active Record creates from [Mastodon](https://github.com/mastodon/mastodon)'s db/schema.rb, as SQL: [source.sql](source.sql). `prisma db pull` writes them up as [schema.prisma](schema.prisma), and from that write-up Prisma and every ORM hekireki generates create tables of their own; their SQL is in [sql/](sql), dumped the way source.sql is, and held to it here.

`=` identical · `≈` the same but for constraint and index names, and the order of columns and statements · `≠` tables written differently, or statements one side has and the other does not. The last column holds each ORM's SQL to what Prisma itself creates from the same write-up: a difference there is the ORM's or hekireki's, one only against source.sql is something the write-up does not say.

| From the write-up | Creates the tables with | SQL | Against source.sql | Against prisma.sql |
| --- | --- | --- | --- | --- |
| Prisma: schema.prisma | `prisma db push` | [prisma.sql](sql/prisma.sql) | ≠ 15 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql ([diff](sql/prisma.diff)) |  |
| hekireki-drizzle | `drizzle-kit export` | [drizzle.sql](sql/drizzle.sql) | ≠ 15 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql ([diff](sql/drizzle.diff)) | ≈ |
| hekireki-kysely | — | — | writes types only | — |
| hekireki-sqlalchemy | `metadata.create_all()` | [sqlalchemy.sql](sql/sqlalchemy.sql) | ≠ 15 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql ([diff](sql/sqlalchemy.diff)) | ≈ |
| hekireki-django | `makemigrations`, `migrate` | [django.sql](sql/django.sql) | ≠ 122 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql ([diff](sql/django.diff)) | ≠ 122 |
| hekireki-gorm | `AutoMigrate` | — | ❌ could not create its tables: ERROR: multiple default values specified for column "id" of table "accounts" (SQLSTATE 42601) | — |
| hekireki-sea-orm | `Schema::create_table_from_entity` | — | ❌ could not create its tables: "invites": Execution Error: error returned from database: relation "users" does not exist | — |
| hekireki-ecto | — | — | writes schemas only (no migrations) | — |
| hekireki-activerecord | — | — | writes models only (no migrations) | — |
| hekireki-eloquent | — | — | writes models only (no migrations) | — |
| hekireki-efcore | `EnsureCreated()` | — | ❌ could not create its tables: Property 'AccountConversations.ParticipantAccountIds' cannot be used as a key because it has type 'List<long>' which does not implement 'IComparable<T>', 'IComparable' or 'IStructuralComparable'. Use 'HasConversion' in 'OnModelCreating' to wrap 'List<long>' with a type that can be compared. | — |
| hekireki-exposed | `SchemaUtils.create` | [exposed.sql](sql/exposed.sql) | ≠ 15 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql ([diff](sql/exposed.diff)) | ≈ |

124 models, postgresql. hekireki 0.55.0, Prisma 8.1.0-dev.9. The same code run against source.sql's tables: [results/mastodon.md](../results/mastodon.md).

<details><summary>Prisma: schema.prisma: 15 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql</summary>

```diff
@@ public.account_conversations
- participant_account_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ participant_account_ids bigint[] DEFAULT ARRAY[]
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] DEFAULT ARRAY[]
@@ public.account_stats
- INDEX USING btree (last_status_at DESC NULLS LAST, account_id)
+ INDEX USING btree (last_status_at DESC, account_id)
@@ public.accounts
- INDEX USING btree (url text_pattern_ops) WHERE (url IS NOT NULL)
- UNIQUE INDEX USING btree (lower((username)::text), COALESCE(lower((domain)::text), ''::text))
- INDEX USING gin ((((setweight(to_tsvector('simple'::regconfig, (display_name)::text), 'A'::"char") || setweight(to_tsvector('simple'::regconfig, (username)::text), 'B'::"char")) || setweight(to_tsvector('simple'::regconfig, (COALESCE(domain, ''::character varying))::text), 'C'::"char"))))
+ INDEX USING btree (url) WHERE (url IS NOT NULL)
@@ public.conversations
- UNIQUE INDEX USING btree (uri text_pattern_ops) WHERE (uri IS NOT NULL)
+ UNIQUE INDEX USING btree (uri) WHERE (uri IS NOT NULL)
@@ public.custom_filters
- context character varying[] DEFAULT ARRAY[] NOT NULL
+ context character varying[] DEFAULT ARRAY[]
@@ public.global_follow_recommendations
- reason character varying[] NOT NULL
+ reason character varying[]
- public.instances (the table is only in source.sql)
@@ public.media_attachments
- UNIQUE INDEX USING btree (shortcode text_pattern_ops) WHERE (shortcode IS NOT NULL)
+ UNIQUE INDEX USING btree (shortcode) WHERE (shortcode IS NOT NULL)
@@ public.oauth_access_tokens
- UNIQUE INDEX USING btree (refresh_token text_pattern_ops) WHERE (refresh_token IS NOT NULL)
+ UNIQUE INDEX USING btree (refresh_token) WHERE (refresh_token IS NOT NULL)
@@ public.polls
- cached_tallies bigint[] DEFAULT ARRAY[] NOT NULL
+ cached_tallies bigint[] DEFAULT ARRAY[]
- options character varying[] DEFAULT ARRAY[] NOT NULL
+ options character varying[] DEFAULT ARRAY[]
@@ public.reports
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] DEFAULT ARRAY[]
@@ public.statuses
- UNIQUE INDEX USING btree (uri text_pattern_ops) WHERE (uri IS NOT NULL)
+ UNIQUE INDEX USING btree (uri) WHERE (uri IS NOT NULL)
@@ public.tags
- UNIQUE INDEX USING btree (lower((name)::text) text_pattern_ops)
@@ public.username_blocks
- UNIQUE INDEX USING btree (lower((username)::text))
@@ public.users
- UNIQUE INDEX USING btree (reset_password_token text_pattern_ops) WHERE (reset_password_token IS NOT NULL)
+ UNIQUE INDEX USING btree (reset_password_token) WHERE (reset_password_token IS NOT NULL)
@@ public.webhooks
- events character varying[] DEFAULT ARRAY[] NOT NULL
+ events character varying[] DEFAULT ARRAY[]
@@ only in source.sql
- CREATE SEQUENCE public.accounts_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.collection_items_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.collections_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE MATERIALIZED VIEW public.instances AS WITH domain_counts(domain, accounts_count) AS (SELECT accounts.domain, count(*) AS accounts_count FROM public.accounts WHERE (accounts.domain IS NOT NULL) GROUP BY accounts.domain) SELECT domain_counts.domain, domain_counts.accounts_count FROM domain_coun …
- CREATE SEQUENCE public.media_attachments_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.notification_requests_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.quotes_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.statuses_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE VIEW public.user_ips AS SELECT user_id, ip, max(used_at) AS used_at FROM (SELECT users.id AS user_id, users.sign_up_ip AS ip, users.created_at AS used_at FROM public.users WHERE (users.sign_up_ip IS NOT NULL) UNION ALL SELECT session_activations.user_id, session_activations.ip, session_activa …
```

</details>

<details><summary>hekireki-drizzle: 15 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql</summary>

```diff
@@ public.account_conversations
- participant_account_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ participant_account_ids bigint[] DEFAULT ARRAY[]
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] DEFAULT ARRAY[]
@@ public.account_stats
- INDEX USING btree (last_status_at DESC NULLS LAST, account_id)
+ INDEX USING btree (last_status_at DESC, account_id)
@@ public.accounts
- INDEX USING btree (url text_pattern_ops) WHERE (url IS NOT NULL)
- UNIQUE INDEX USING btree (lower((username)::text), COALESCE(lower((domain)::text), ''::text))
- INDEX USING gin ((((setweight(to_tsvector('simple'::regconfig, (display_name)::text), 'A'::"char") || setweight(to_tsvector('simple'::regconfig, (username)::text), 'B'::"char")) || setweight(to_tsvector('simple'::regconfig, (COALESCE(domain, ''::character varying))::text), 'C'::"char"))))
+ INDEX USING btree (url) WHERE (url IS NOT NULL)
@@ public.conversations
- UNIQUE INDEX USING btree (uri text_pattern_ops) WHERE (uri IS NOT NULL)
+ UNIQUE INDEX USING btree (uri) WHERE (uri IS NOT NULL)
@@ public.custom_filters
- context character varying[] DEFAULT ARRAY[] NOT NULL
+ context character varying[] DEFAULT ARRAY[]
@@ public.global_follow_recommendations
- reason character varying[] NOT NULL
+ reason character varying[]
- public.instances (the table is only in source.sql)
@@ public.media_attachments
- UNIQUE INDEX USING btree (shortcode text_pattern_ops) WHERE (shortcode IS NOT NULL)
+ UNIQUE INDEX USING btree (shortcode) WHERE (shortcode IS NOT NULL)
@@ public.oauth_access_tokens
- UNIQUE INDEX USING btree (refresh_token text_pattern_ops) WHERE (refresh_token IS NOT NULL)
+ UNIQUE INDEX USING btree (refresh_token) WHERE (refresh_token IS NOT NULL)
@@ public.polls
- cached_tallies bigint[] DEFAULT ARRAY[] NOT NULL
+ cached_tallies bigint[] DEFAULT ARRAY[]
- options character varying[] DEFAULT ARRAY[] NOT NULL
+ options character varying[] DEFAULT ARRAY[]
@@ public.reports
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] DEFAULT ARRAY[]
@@ public.statuses
- UNIQUE INDEX USING btree (uri text_pattern_ops) WHERE (uri IS NOT NULL)
+ UNIQUE INDEX USING btree (uri) WHERE (uri IS NOT NULL)
@@ public.tags
- UNIQUE INDEX USING btree (lower((name)::text) text_pattern_ops)
@@ public.username_blocks
- UNIQUE INDEX USING btree (lower((username)::text))
@@ public.users
- UNIQUE INDEX USING btree (reset_password_token text_pattern_ops) WHERE (reset_password_token IS NOT NULL)
+ UNIQUE INDEX USING btree (reset_password_token) WHERE (reset_password_token IS NOT NULL)
@@ public.webhooks
- events character varying[] DEFAULT ARRAY[] NOT NULL
+ events character varying[] DEFAULT ARRAY[]
@@ only in source.sql
- CREATE SEQUENCE public.accounts_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.collection_items_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.collections_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE MATERIALIZED VIEW public.instances AS WITH domain_counts(domain, accounts_count) AS (SELECT accounts.domain, count(*) AS accounts_count FROM public.accounts WHERE (accounts.domain IS NOT NULL) GROUP BY accounts.domain) SELECT domain_counts.domain, domain_counts.accounts_count FROM domain_coun …
- CREATE SEQUENCE public.media_attachments_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.notification_requests_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.quotes_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.statuses_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE VIEW public.user_ips AS SELECT user_id, ip, max(used_at) AS used_at FROM (SELECT users.id AS user_id, users.sign_up_ip AS ip, users.created_at AS used_at FROM public.users WHERE (users.sign_up_ip IS NOT NULL) UNION ALL SELECT session_activations.user_id, session_activations.ip, session_activa …
```

</details>

<details><summary>hekireki-sqlalchemy: 15 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql</summary>

```diff
@@ public.account_conversations
- participant_account_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ participant_account_ids bigint[] DEFAULT ARRAY[]
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] DEFAULT ARRAY[]
@@ public.account_stats
- INDEX USING btree (last_status_at DESC NULLS LAST, account_id)
+ INDEX USING btree (last_status_at DESC, account_id)
@@ public.accounts
- INDEX USING btree (url text_pattern_ops) WHERE (url IS NOT NULL)
- UNIQUE INDEX USING btree (lower((username)::text), COALESCE(lower((domain)::text), ''::text))
- INDEX USING gin ((((setweight(to_tsvector('simple'::regconfig, (display_name)::text), 'A'::"char") || setweight(to_tsvector('simple'::regconfig, (username)::text), 'B'::"char")) || setweight(to_tsvector('simple'::regconfig, (COALESCE(domain, ''::character varying))::text), 'C'::"char"))))
+ INDEX USING btree (url) WHERE (url IS NOT NULL)
@@ public.conversations
- UNIQUE INDEX USING btree (uri text_pattern_ops) WHERE (uri IS NOT NULL)
+ UNIQUE INDEX USING btree (uri) WHERE (uri IS NOT NULL)
@@ public.custom_filters
- context character varying[] DEFAULT ARRAY[] NOT NULL
+ context character varying[] DEFAULT ARRAY[]
@@ public.global_follow_recommendations
- reason character varying[] NOT NULL
+ reason character varying[]
- public.instances (the table is only in source.sql)
@@ public.media_attachments
- UNIQUE INDEX USING btree (shortcode text_pattern_ops) WHERE (shortcode IS NOT NULL)
+ UNIQUE INDEX USING btree (shortcode) WHERE (shortcode IS NOT NULL)
@@ public.oauth_access_tokens
- UNIQUE INDEX USING btree (refresh_token text_pattern_ops) WHERE (refresh_token IS NOT NULL)
+ UNIQUE INDEX USING btree (refresh_token) WHERE (refresh_token IS NOT NULL)
@@ public.polls
- cached_tallies bigint[] DEFAULT ARRAY[] NOT NULL
+ cached_tallies bigint[] DEFAULT ARRAY[]
- options character varying[] DEFAULT ARRAY[] NOT NULL
+ options character varying[] DEFAULT ARRAY[]
@@ public.reports
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] DEFAULT ARRAY[]
@@ public.statuses
- UNIQUE INDEX USING btree (uri text_pattern_ops) WHERE (uri IS NOT NULL)
+ UNIQUE INDEX USING btree (uri) WHERE (uri IS NOT NULL)
@@ public.tags
- UNIQUE INDEX USING btree (lower((name)::text) text_pattern_ops)
@@ public.username_blocks
- UNIQUE INDEX USING btree (lower((username)::text))
@@ public.users
- UNIQUE INDEX USING btree (reset_password_token text_pattern_ops) WHERE (reset_password_token IS NOT NULL)
+ UNIQUE INDEX USING btree (reset_password_token) WHERE (reset_password_token IS NOT NULL)
@@ public.webhooks
- events character varying[] DEFAULT ARRAY[] NOT NULL
+ events character varying[] DEFAULT ARRAY[]
@@ only in source.sql
- CREATE SEQUENCE public.accounts_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.collection_items_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.collections_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE MATERIALIZED VIEW public.instances AS WITH domain_counts(domain, accounts_count) AS (SELECT accounts.domain, count(*) AS accounts_count FROM public.accounts WHERE (accounts.domain IS NOT NULL) GROUP BY accounts.domain) SELECT domain_counts.domain, domain_counts.accounts_count FROM domain_coun …
- CREATE SEQUENCE public.media_attachments_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.notification_requests_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.quotes_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.statuses_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE VIEW public.user_ips AS SELECT user_id, ip, max(used_at) AS used_at FROM (SELECT users.id AS user_id, users.sign_up_ip AS ip, users.created_at AS used_at FROM public.users WHERE (users.sign_up_ip IS NOT NULL) UNION ALL SELECT session_activations.user_id, session_activations.ip, session_activa …
```

</details>

<details><summary>hekireki-django: 122 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql</summary>

```diff
@@ public.account_aliases
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- acct character varying DEFAULT ''::character varying NOT NULL
+ acct character varying NOT NULL
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- uri character varying DEFAULT ''::character varying NOT NULL
+ uri character varying NOT NULL
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
@@ public.account_conversations
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- lock_version integer DEFAULT 0 NOT NULL
+ lock_version integer NOT NULL
- participant_account_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ participant_account_ids bigint[] NOT NULL
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] NOT NULL
- unread boolean DEFAULT false NOT NULL
+ unread boolean NOT NULL
- FOREIGN KEY (conversation_id) REFERENCES public.conversations(id) ON DELETE CASCADE
+ FOREIGN KEY (conversation_id) REFERENCES public.conversations(id) DEFERRABLE INITIALLY DEFERRED
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
@@ public.account_deletion_requests
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
@@ public.account_domain_blocks
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
@@ public.account_migrations
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- acct character varying DEFAULT ''::character varying NOT NULL
+ acct character varying NOT NULL
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- followers_count bigint DEFAULT 0 NOT NULL
+ followers_count bigint NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
- FOREIGN KEY (target_account_id) REFERENCES public.accounts(id) ON DELETE SET NULL
+ FOREIGN KEY (target_account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
- INDEX USING btree (target_account_id) WHERE (target_account_id IS NOT NULL)
@@ public.account_moderation_notes
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
- FOREIGN KEY (target_account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (target_account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
@@ public.account_notes
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- FOREIGN KEY (target_account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (target_account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
@@ public.account_pins
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- FOREIGN KEY (target_account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (target_account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
@@ public.account_reach_filters
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- saturated boolean DEFAULT false NOT NULL
+ saturated boolean NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
@@ public.account_relationship_severance_events
- id bigint NOT NULL DEFAULT nextval(<its sequence>)
+ id bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY
- created_at timestamp without time zone NOT NULL
+ created_at timestamp with time zone NOT NULL
- followers_count integer DEFAULT 0 NOT NULL
+ followers_count integer NOT NULL
- following_count integer DEFAULT 0 NOT NULL
+ following_count integer NOT NULL
- updated_at timestamp without time zone NOT NULL
+ updated_at timestamp with time zone NOT NULL
- FOREIGN KEY (account_id) REFERENCES public.accounts(id) ON DELETE CASCADE
+ FOREIGN KEY (account_id) REFERENCES public.accounts(id) DEFERRABLE INITIALLY DEFERRED
- FOREIGN KEY (relationship_severance_event_id) REFERENCES public.relationship_severance_events(id) ON DELETE CASCADE
@@ … 1669 more lines in sql/django.diff
```

</details>

<details><summary>hekireki-exposed: 15 tables written differently, 1 table only in source.sql, 9 other statements only in source.sql</summary>

```diff
@@ public.account_conversations
- participant_account_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ participant_account_ids bigint[] DEFAULT ARRAY[]
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] DEFAULT ARRAY[]
@@ public.account_stats
- INDEX USING btree (last_status_at DESC NULLS LAST, account_id)
+ INDEX USING btree (last_status_at DESC, account_id)
@@ public.accounts
- INDEX USING btree (url text_pattern_ops) WHERE (url IS NOT NULL)
- UNIQUE INDEX USING btree (lower((username)::text), COALESCE(lower((domain)::text), ''::text))
- INDEX USING gin ((((setweight(to_tsvector('simple'::regconfig, (display_name)::text), 'A'::"char") || setweight(to_tsvector('simple'::regconfig, (username)::text), 'B'::"char")) || setweight(to_tsvector('simple'::regconfig, (COALESCE(domain, ''::character varying))::text), 'C'::"char"))))
+ INDEX USING btree (url) WHERE (url IS NOT NULL)
@@ public.conversations
- UNIQUE INDEX USING btree (uri text_pattern_ops) WHERE (uri IS NOT NULL)
+ UNIQUE INDEX USING btree (uri) WHERE (uri IS NOT NULL)
@@ public.custom_filters
- context character varying[] DEFAULT ARRAY[] NOT NULL
+ context character varying[] DEFAULT ARRAY[]
@@ public.global_follow_recommendations
- reason character varying[] NOT NULL
+ reason character varying[]
- public.instances (the table is only in source.sql)
@@ public.media_attachments
- UNIQUE INDEX USING btree (shortcode text_pattern_ops) WHERE (shortcode IS NOT NULL)
+ UNIQUE INDEX USING btree (shortcode) WHERE (shortcode IS NOT NULL)
@@ public.oauth_access_tokens
- UNIQUE INDEX USING btree (refresh_token text_pattern_ops) WHERE (refresh_token IS NOT NULL)
+ UNIQUE INDEX USING btree (refresh_token) WHERE (refresh_token IS NOT NULL)
@@ public.polls
- cached_tallies bigint[] DEFAULT ARRAY[] NOT NULL
+ cached_tallies bigint[] DEFAULT ARRAY[]
- options character varying[] DEFAULT ARRAY[] NOT NULL
+ options character varying[] DEFAULT ARRAY[]
@@ public.reports
- status_ids bigint[] DEFAULT ARRAY[] NOT NULL
+ status_ids bigint[] DEFAULT ARRAY[]
@@ public.statuses
- UNIQUE INDEX USING btree (uri text_pattern_ops) WHERE (uri IS NOT NULL)
+ UNIQUE INDEX USING btree (uri) WHERE (uri IS NOT NULL)
@@ public.tags
- UNIQUE INDEX USING btree (lower((name)::text) text_pattern_ops)
@@ public.username_blocks
- UNIQUE INDEX USING btree (lower((username)::text))
@@ public.users
- UNIQUE INDEX USING btree (reset_password_token text_pattern_ops) WHERE (reset_password_token IS NOT NULL)
+ UNIQUE INDEX USING btree (reset_password_token) WHERE (reset_password_token IS NOT NULL)
@@ public.webhooks
- events character varying[] DEFAULT ARRAY[] NOT NULL
+ events character varying[] DEFAULT ARRAY[]
@@ only in source.sql
- CREATE SEQUENCE public.accounts_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.collection_items_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.collections_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE MATERIALIZED VIEW public.instances AS WITH domain_counts(domain, accounts_count) AS (SELECT accounts.domain, count(*) AS accounts_count FROM public.accounts WHERE (accounts.domain IS NOT NULL) GROUP BY accounts.domain) SELECT domain_counts.domain, domain_counts.accounts_count FROM domain_coun …
- CREATE SEQUENCE public.media_attachments_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.notification_requests_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.quotes_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE SEQUENCE public.statuses_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
- CREATE VIEW public.user_ips AS SELECT user_id, ip, max(used_at) AS used_at FROM (SELECT users.id AS user_id, users.sign_up_ip AS ip, users.created_at AS used_at FROM public.users WHERE (users.sign_up_ip IS NOT NULL) UNION ALL SELECT session_activations.user_id, session_activations.ip, session_activa …
```

</details>

Written by `node pull/pull.mjs`; do not edit.
