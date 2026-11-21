# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Application](#application) · [Grant](#grant) · [AccessToken](#accesstoken) · [RefreshToken](#refreshtoken) · [IDToken](#idtoken) · [DeviceGrant](#devicegrant)

<details>
<summary>The whole schema as SQL</summary>

```sql
CREATE TABLE "django_admin_log" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "action_time" datetime NOT NULL, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "object_id" text NULL, "object_repr" varchar(200) NOT NULL, "action_flag" smallint unsigned NOT NULL CHECK ("action_flag" >= 0), "change_message" text NOT NULL);

CREATE TABLE "auth_permission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(255) NOT NULL, "content_type_id" integer NOT NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "codename" varchar(100) NOT NULL);

CREATE TABLE "auth_group" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(150) NOT NULL UNIQUE);

CREATE TABLE "auth_group_permissions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED, "permission_id" integer NOT NULL REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "auth_user" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "password" varchar(128) NOT NULL, "last_login" datetime NULL, "is_superuser" bool NOT NULL, "username" varchar(150) NOT NULL UNIQUE, "first_name" varchar(150) NOT NULL, "last_name" varchar(150) NOT NULL, "email" varchar(254) NOT NULL, "is_staff" bool NOT NULL, "is_active" bool NOT NULL, "date_joined" datetime NOT NULL);

CREATE TABLE "auth_user_groups" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "auth_user_user_permissions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "permission_id" integer NOT NULL REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "django_content_type" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "app_label" varchar(100) NOT NULL, "model" varchar(100) NOT NULL);

CREATE TABLE "django_session" ("session_key" varchar(40) NOT NULL PRIMARY KEY, "session_data" text NOT NULL, "expire_date" datetime NOT NULL);

CREATE TABLE "oauth2_provider_application" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "client_id" varchar(255) NOT NULL UNIQUE, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "redirect_uris" text NOT NULL, "post_logout_redirect_uris" text NOT NULL, "client_type" varchar(32) NOT NULL, "authorization_grant_type" varchar(44) NOT NULL, "client_secret" varchar(255) NOT NULL, "hash_client_secret" bool NOT NULL, "name" varchar(255) NOT NULL, "skip_authorization" bool NOT NULL, "created" datetime NOT NULL, "updated" datetime NOT NULL, "algorithm" varchar(5) NOT NULL, "allowed_origins" text NOT NULL, "registration_source" varchar(32) NOT NULL, "cimd_expires_at" datetime NULL);

CREATE TABLE "oauth2_provider_grant" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "code" varchar(255) NOT NULL UNIQUE, "application_id" bigint NOT NULL REFERENCES "oauth2_provider_application" ("id") DEFERRABLE INITIALLY DEFERRED, "expires" datetime NOT NULL, "redirect_uri" text NOT NULL, "scope" text NOT NULL, "created" datetime NOT NULL, "updated" datetime NOT NULL, "code_challenge" varchar(128) NOT NULL, "code_challenge_method" varchar(10) NOT NULL, "nonce" varchar(255) NOT NULL, "claims" text NOT NULL, "resource" text NOT NULL CHECK ((JSON_VALID("resource") OR "resource" IS NULL)));

CREATE TABLE "oauth2_provider_accesstoken" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "source_refresh_token_id" bigint NULL UNIQUE REFERENCES "oauth2_provider_refreshtoken" ("id") DEFERRABLE INITIALLY DEFERRED, "token" text NOT NULL, "token_checksum" varchar(64) NOT NULL UNIQUE, "id_token_id" bigint NULL UNIQUE REFERENCES "oauth2_provider_idtoken" ("id") DEFERRABLE INITIALLY DEFERRED, "application_id" bigint NULL REFERENCES "oauth2_provider_application" ("id") DEFERRABLE INITIALLY DEFERRED, "expires" datetime NOT NULL, "scope" text NOT NULL, "resource" text NOT NULL CHECK ((JSON_VALID("resource") OR "resource" IS NULL)), "created" datetime NOT NULL, "updated" datetime NOT NULL);

CREATE TABLE "oauth2_provider_refreshtoken" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "token" text NOT NULL, "token_checksum" varchar(64) NOT NULL, "application_id" bigint NOT NULL REFERENCES "oauth2_provider_application" ("id") DEFERRABLE INITIALLY DEFERRED, "access_token_id" bigint NULL UNIQUE REFERENCES "oauth2_provider_accesstoken" ("id") DEFERRABLE INITIALLY DEFERRED, "token_family" char(32) NULL, "resource" text NOT NULL CHECK ((JSON_VALID("resource") OR "resource" IS NULL)), "created" datetime NOT NULL, "updated" datetime NOT NULL, "revoked" datetime NULL);

CREATE TABLE "oauth2_provider_idtoken" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "jti" char(32) NOT NULL UNIQUE, "application_id" bigint NULL REFERENCES "oauth2_provider_application" ("id") DEFERRABLE INITIALLY DEFERRED, "expires" datetime NOT NULL, "scope" text NOT NULL, "created" datetime NOT NULL, "updated" datetime NOT NULL);

CREATE TABLE "oauth2_provider_devicegrant" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "device_code" varchar(100) NOT NULL, "user_code" varchar(100) NOT NULL, "scope" text NOT NULL, "interval" integer NOT NULL, "expires" datetime NOT NULL, "status" varchar(64) NOT NULL, "client_id" varchar(100) NOT NULL, "last_checked" datetime NOT NULL, CONSTRAINT "oauth2_provider_devicegrant_unique_device_code" UNIQUE ("device_code"));

CREATE INDEX "django_admin_log_user_id_c564eba6" ON "django_admin_log" ("user_id");

CREATE INDEX "django_admin_log_content_type_id_c4bce8eb" ON "django_admin_log" ("content_type_id");

CREATE UNIQUE INDEX "auth_permission_content_type_id_codename_01ab375a_uniq" ON "auth_permission" ("content_type_id", "codename");

CREATE INDEX "auth_permission_content_type_id_2f476e4b" ON "auth_permission" ("content_type_id");

CREATE UNIQUE INDEX "auth_group_permissions_group_id_permission_id_0cd325b0_uniq" ON "auth_group_permissions" ("group_id", "permission_id");

CREATE INDEX "auth_group_permissions_group_id_b120cbf9" ON "auth_group_permissions" ("group_id");

CREATE INDEX "auth_group_permissions_permission_id_84c5c92e" ON "auth_group_permissions" ("permission_id");

CREATE UNIQUE INDEX "auth_user_groups_user_id_group_id_94350c0c_uniq" ON "auth_user_groups" ("user_id", "group_id");

CREATE INDEX "auth_user_groups_user_id_6a12ed8b" ON "auth_user_groups" ("user_id");

CREATE INDEX "auth_user_groups_group_id_97559544" ON "auth_user_groups" ("group_id");

CREATE UNIQUE INDEX "auth_user_user_permissions_user_id_permission_id_14a6b632_uniq" ON "auth_user_user_permissions" ("user_id", "permission_id");

CREATE INDEX "auth_user_user_permissions_user_id_a95ead1b" ON "auth_user_user_permissions" ("user_id");

CREATE INDEX "auth_user_user_permissions_permission_id_1fbb5f2c" ON "auth_user_user_permissions" ("permission_id");

CREATE UNIQUE INDEX "django_content_type_app_label_model_76bd3d3b_uniq" ON "django_content_type" ("app_label", "model");

CREATE INDEX "django_session_expire_date_a5c62663" ON "django_session" ("expire_date");

CREATE INDEX "oauth2_provider_application_user_id_79829054" ON "oauth2_provider_application" ("user_id");

CREATE INDEX "oauth2_provider_application_client_secret_53133678" ON "oauth2_provider_application" ("client_secret");

CREATE INDEX "oauth2_provider_grant_user_id_e8f62af8" ON "oauth2_provider_grant" ("user_id");

CREATE INDEX "oauth2_provider_grant_application_id_81923564" ON "oauth2_provider_grant" ("application_id");

CREATE INDEX "oauth2_provider_accesstoken_user_id_6e4c9a65" ON "oauth2_provider_accesstoken" ("user_id");

CREATE INDEX "oauth2_provider_accesstoken_application_id_b22886e1" ON "oauth2_provider_accesstoken" ("application_id");

CREATE UNIQUE INDEX "oauth2_provider_refreshtoken_token_checksum_revoked_46c81d1d_uniq" ON "oauth2_provider_refreshtoken" ("token_checksum", "revoked");

CREATE INDEX "oauth2_provider_refreshtoken_user_id_da837fce" ON "oauth2_provider_refreshtoken" ("user_id");

CREATE INDEX "oauth2_provider_refreshtoken_application_id_2d1c311b" ON "oauth2_provider_refreshtoken" ("application_id");

CREATE INDEX "oauth2_prov_token_f_996e8a_idx" ON "oauth2_provider_refreshtoken" ("token_family");

CREATE INDEX "oauth2_provider_idtoken_user_id_dd512b59" ON "oauth2_provider_idtoken" ("user_id");

CREATE INDEX "oauth2_provider_idtoken_application_id_08c5ff4f" ON "oauth2_provider_idtoken" ("application_id");

CREATE INDEX "oauth2_provider_devicegrant_user_id_1cec5156" ON "oauth2_provider_devicegrant" ("user_id");

CREATE INDEX "oauth2_provider_devicegrant_client_id_229dd06d" ON "oauth2_provider_devicegrant" ("client_id");
```

</details>

## LogEntry

Table `django_admin_log`, app `admin`.

![LogEntry and the models it relates to](tables/LogEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `action_time` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `object_id` | `TextField` | `text` | yes |  |  |  |
| `object_repr` | `CharField` | `varchar(200)` |  |  |  |  |
| `action_flag` | `LogEntryActionFlag` | `smallint unsigned` |  |  |  |  |
| `change_message` | `TextField` | `text` |  |  |  |  |

Indexes:

- index `django_admin_log_user_id_c564eba6` on `user_id`
- index `django_admin_log_content_type_id_c4bce8eb` on `content_type_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `content_type_id` → [ContentType](#contenttype).`id`, on delete set null

## Permission

Table `auth_permission`, app `auth`.

The permissions system provides a way to assign permissions to specific
users and groups of users.

The permission system is used by the Django admin site, but may also be
useful in your own code. The Django admin site uses permissions as follows:

    - The "add" permission limits the user's ability to view the "add" form
      and add an object.
    - The "change" permission limits a user's ability to view the change
      list, view the "change" form and change an object.
    - The "delete" permission limits the ability to delete an object.
    - The "view" permission limits the ability to view an object.

Permissions are set globally per type of object, not per specific object
instance. It is possible to say "Mary may change news stories," but it's
not currently possible to say "Mary may change news stories, but only the
ones she created herself" or "Mary may only change news stories that have a
certain status or publication date."

The permissions listed above are automatically created for each model.

![Permission and the models it relates to](tables/Permission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `content_type_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `codename` | `CharField` | `varchar(100)` |  |  |  |  |

Indexes:

- index `auth_permission_content_type_id_2f476e4b` on `content_type_id`
- unique `auth_permission_content_type_id_codename_01ab375a_uniq` on `content_type_id`, `codename`

Relations:

- `content_type_id` → [ContentType](#contenttype).`id`, on delete cascade
- `group_set` — many to many with [Group](#group)
- `user_set` — many to many with [User](#user)

## Group

Table `auth_group`, app `auth`.

Groups are a generic way of categorizing users to apply permissions, or
some other label, to those users. A user can belong to any number of
groups.

A user in a group automatically has all the permissions granted to that
group. For example, if the group 'Site editors' has the permission
can_edit_home_page, any user in that group will have that permission.

Beyond permissions, groups are a convenient way to categorize users to
apply some label, or extended functionality, to them. For example, you
could create a group 'Special users', and you could write code that would
do special things to those users -- such as giving them access to a
members-only portion of your site, or sending them members-only email
messages.

![Group and the models it relates to](tables/Group.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(150)` |  | unique |  |  |

Relations:

- `permissions` — many to many with [Permission](#permission)
- `user_set` — many to many with [User](#user)

## User

Table `auth_user`, app `auth`.

Columns from `AbstractUser` (abstract): `username`, `first_name`, `last_name`, `email`, `is_staff`, `is_active`, `date_joined`.

Columns from `AbstractBaseUser` (abstract): `password`, `last_login`.

Columns from `PermissionsMixin` (abstract): `is_superuser`.

Users within the Django authentication system are represented by this
model.

Username and password are required. Other fields are optional.

![User and the models it relates to](tables/User.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `password` | `CharField` | `varchar(128)` |  |  |  |  |
| `last_login` | `DateTimeField` | `datetime` | yes |  |  |  |
| `is_superuser` | `BooleanField` | `bool` |  |  | `False` | Designates that this user has all permissions without explicitly assigning them. |
| `username` | `CharField` | `varchar(150)` |  | unique |  | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. |
| `first_name` | `CharField` | `varchar(150)` |  |  |  |  |
| `last_name` | `CharField` | `varchar(150)` |  |  |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `is_staff` | `BooleanField` | `bool` |  |  | `False` | Designates whether the user can log into this admin site. |
| `is_active` | `BooleanField` | `bool` |  |  | `True` | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. |
| `date_joined` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Relations:

- [LogEntry](#logentry).`user_id` → `id`, on delete cascade
- `groups` — many to many with [Group](#group)
- `user_permissions` — many to many with [Permission](#permission)
- [Application](#application).`user_id` → `id`, on delete cascade
- [Grant](#grant).`user_id` → `id`, on delete cascade
- [AccessToken](#accesstoken).`user_id` → `id`, on delete cascade
- [RefreshToken](#refreshtoken).`user_id` → `id`, on delete cascade
- [IDToken](#idtoken).`user_id` → `id`, on delete cascade
- [DeviceGrant](#devicegrant).`user_id` → `id`, on delete cascade

## ContentType

Table `django_content_type`, app `contenttypes`.

![ContentType and the models it relates to](tables/ContentType.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `app_label` | `CharField` | `varchar(100)` |  |  |  |  |
| `model` | `CharField` | `varchar(100)` |  |  |  |  |

Indexes:

- unique `django_content_type_app_label_model_76bd3d3b_uniq` on `app_label`, `model`

Relations:

- [LogEntry](#logentry).`content_type_id` → `id`, on delete set null
- [Permission](#permission).`content_type_id` → `id`, on delete cascade

## Session

Table `django_session`, app `sessions`.

Columns from `AbstractBaseSession` (abstract): `session_key`, `session_data`, `expire_date`.

Django provides full support for anonymous sessions. The session
framework lets you store and retrieve arbitrary data on a
per-site-visitor basis. It stores data on the server side and
abstracts the sending and receiving of cookies. Cookies contain a
session ID -- not the data itself.

The Django sessions framework is entirely cookie-based. It does
not fall back to putting session IDs in URLs. This is an intentional
design decision. Not only does that behavior make URLs ugly, it makes
your site vulnerable to session-ID theft via the "Referer" header.

For complete documentation on using Sessions in your code, consult
the sessions documentation that is shipped with Django (also available
on the Django web site).

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `session_key` | `CharField` | `varchar(40)` |  | PK |  |  |
| `session_data` | `TextField` | `text` |  |  |  |  |
| `expire_date` | `DateTimeField` | `datetime` |  |  |  |  |

Indexes:

- index `django_session_expire_date_a5c62663` on `expire_date`

## Application

Table `oauth2_provider_application`, app `oauth2_provider`.

Columns from `AbstractApplication` (abstract): `id`, `client_id`, `user_id`, `redirect_uris`, `post_logout_redirect_uris`, `client_type`, `authorization_grant_type`, `client_secret`, `hash_client_secret`, `name`, `skip_authorization`, `created`, `updated`, `algorithm`, `allowed_origins`, `registration_source`, `cimd_expires_at`.

![Application and the models it relates to](tables/Application.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `client_id` | `CharField` | `varchar(255)` |  | unique | `generate_client_id()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `redirect_uris` | `TextField` | `text` |  |  |  | Allowed URIs list, space separated |
| `post_logout_redirect_uris` | `TextField` | `text` |  |  | `''` | Allowed Post Logout URIs list, space separated |
| `client_type` | `ApplicationClientType` | `varchar(32)` |  |  |  |  |
| `authorization_grant_type` | `ApplicationAuthorizationGrantType` | `varchar(44)` |  |  |  |  |
| `client_secret` | `ClientSecretField` | `varchar(255)` |  |  | `generate_client_secret()` | Client secret for authentication |
| `hash_client_secret` | `BooleanField` | `bool` |  |  | `True` |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `skip_authorization` | `BooleanField` | `bool` |  |  | `False` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `algorithm` | `ApplicationAlgorithm` | `varchar(5)` |  |  | `''` |  |
| `allowed_origins` | `TextField` | `text` |  |  | `''` | Allowed origins list to enable CORS, space separated |
| `registration_source` | `ApplicationRegistrationSource` | `varchar(32)` |  |  | `'manual'` | How this application was registered (manual, DCR per RFC 7591, or CIMD) |
| `cimd_expires_at` | `DateTimeField` | `datetime` | yes |  | `None` | When the cached Client ID Metadata Document should be re-fetched |

Indexes:

- index `oauth2_provider_application_user_id_79829054` on `user_id`
- index `oauth2_provider_application_client_secret_53133678` on `client_secret`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [Grant](#grant).`application_id` → `id`, on delete cascade
- [AccessToken](#accesstoken).`application_id` → `id`, on delete cascade
- [RefreshToken](#refreshtoken).`application_id` → `id`, on delete cascade
- [IDToken](#idtoken).`application_id` → `id`, on delete cascade

## Grant

Table `oauth2_provider_grant`, app `oauth2_provider`.

Columns from `AbstractGrant` (abstract): `id`, `user_id`, `code`, `application_id`, `expires`, `redirect_uri`, `scope`, `created`, `updated`, `code_challenge`, `code_challenge_method`, `nonce`, `claims`, `resource`.

![Grant and the models it relates to](tables/Grant.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `code` | `CharField` | `varchar(255)` |  | unique |  |  |
| `application_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `expires` | `DateTimeField` | `datetime` |  |  |  |  |
| `redirect_uri` | `TextField` | `text` |  |  |  |  |
| `scope` | `TextField` | `text` |  |  |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `code_challenge` | `CharField` | `varchar(128)` |  |  | `''` |  |
| `code_challenge_method` | `GrantCodeChallengeMethod` | `varchar(10)` |  |  | `''` |  |
| `nonce` | `CharField` | `varchar(255)` |  |  | `''` |  |
| `claims` | `TextField` | `text` |  |  |  |  |
| `resource` | `ResourceJSONField` | `text` |  |  | `list()` |  |

Indexes:

- index `oauth2_provider_grant_user_id_e8f62af8` on `user_id`
- index `oauth2_provider_grant_application_id_81923564` on `application_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `application_id` → [Application](#application).`id`, on delete cascade

## AccessToken

Table `oauth2_provider_accesstoken`, app `oauth2_provider`.

Columns from `AbstractAccessToken` (abstract): `id`, `user_id`, `source_refresh_token_id`, `token`, `token_checksum`, `id_token_id`, `application_id`, `expires`, `scope`, `resource`, `created`, `updated`.

![AccessToken and the models it relates to](tables/AccessToken.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `source_refresh_token_id` | `OneToOneField` | `bigint` | yes | FK, unique |  |  |
| `token` | `TextField` | `text` |  |  |  |  |
| `token_checksum` | `TokenChecksumField` | `varchar(64)` |  | unique |  |  |
| `id_token_id` | `OneToOneField` | `bigint` | yes | FK, unique |  |  |
| `application_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `expires` | `DateTimeField` | `datetime` |  |  |  |  |
| `scope` | `TextField` | `text` |  |  |  |  |
| `resource` | `ResourceJSONField` | `text` |  |  | `list()` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `oauth2_provider_accesstoken_user_id_6e4c9a65` on `user_id`
- index `oauth2_provider_accesstoken_application_id_b22886e1` on `application_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `source_refresh_token_id` → [RefreshToken](#refreshtoken).`id`, on delete set null
- `id_token_id` → [IDToken](#idtoken).`id`, on delete cascade
- `application_id` → [Application](#application).`id`, on delete cascade
- [RefreshToken](#refreshtoken).`access_token_id` → `id`, on delete set null

## RefreshToken

Table `oauth2_provider_refreshtoken`, app `oauth2_provider`.

Columns from `AbstractRefreshToken` (abstract): `id`, `user_id`, `token`, `token_checksum`, `application_id`, `access_token_id`, `token_family`, `resource`, `created`, `updated`, `revoked`.

![RefreshToken and the models it relates to](tables/RefreshToken.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `token` | `TextField` | `text` |  |  |  |  |
| `token_checksum` | `TokenChecksumField` | `varchar(64)` |  |  |  |  |
| `application_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `access_token_id` | `OneToOneField` | `bigint` | yes | FK, unique |  |  |
| `token_family` | `UUIDField` | `char(32)` | yes |  |  |  |
| `resource` | `ResourceJSONField` | `text` |  |  | `list()` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `revoked` | `DateTimeField` | `datetime` | yes |  |  |  |

Indexes:

- index `oauth2_provider_refreshtoken_user_id_da837fce` on `user_id`
- index `oauth2_provider_refreshtoken_application_id_2d1c311b` on `application_id`
- unique `oauth2_provider_refreshtoken_token_checksum_revoked_46c81d1d_uniq` on `token_checksum`, `revoked`
- index `oauth2_prov_token_f_996e8a_idx` on `token_family`

Relations:

- [AccessToken](#accesstoken).`source_refresh_token_id` → `id`, on delete set null
- `user_id` → [User](#user).`id`, on delete cascade
- `application_id` → [Application](#application).`id`, on delete cascade
- `access_token_id` → [AccessToken](#accesstoken).`id`, on delete set null

## IDToken

Table `oauth2_provider_idtoken`, app `oauth2_provider`.

Columns from `AbstractIDToken` (abstract): `id`, `user_id`, `jti`, `application_id`, `expires`, `scope`, `created`, `updated`.

![IDToken and the models it relates to](tables/IDToken.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `jti` | `UUIDField` | `char(32)` |  | unique | `uuid4()` |  |
| `application_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `expires` | `DateTimeField` | `datetime` |  |  |  |  |
| `scope` | `TextField` | `text` |  |  |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `oauth2_provider_idtoken_user_id_dd512b59` on `user_id`
- index `oauth2_provider_idtoken_application_id_08c5ff4f` on `application_id`

Relations:

- [AccessToken](#accesstoken).`id_token_id` → `id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade
- `application_id` → [Application](#application).`id`, on delete cascade

## DeviceGrant

Table `oauth2_provider_devicegrant`, app `oauth2_provider`.

Columns from `AbstractDeviceGrant` (abstract): `id`, `user_id`, `device_code`, `user_code`, `scope`, `interval`, `expires`, `status`, `client_id`, `last_checked`.

![DeviceGrant and the models it relates to](tables/DeviceGrant.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `device_code` | `CharField` | `varchar(100)` |  |  |  |  |
| `user_code` | `CharField` | `varchar(100)` |  |  |  |  |
| `scope` | `TextField` | `text` |  |  |  |  |
| `interval` | `IntegerField` | `integer` |  |  | `5` |  |
| `expires` | `DateTimeField` | `datetime` |  |  |  |  |
| `status` | `DeviceGrantStatus` | `varchar(64)` |  |  | `'authorization-pending'` |  |
| `client_id` | `CharField` | `varchar(100)` |  |  |  |  |
| `last_checked` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `oauth2_provider_devicegrant_user_id_1cec5156` on `user_id`
- index `oauth2_provider_devicegrant_client_id_229dd06d` on `client_id`
- unique `oauth2_provider_devicegrant_unique_device_code` on `device_code`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## ApplicationClientType

| Name | Stored as |
| --- | --- |
| `Confidential` | `confidential` |
| `Public` | `public` |

## ApplicationAuthorizationGrantType

| Name | Stored as |
| --- | --- |
| `Authorization code` | `authorization-code` |
| `Device Code` | `urn:ietf:params:oauth:grant-type:device_code` |
| `Implicit` | `implicit` |
| `Resource owner password-based` | `password` |
| `Client credentials` | `client-credentials` |
| `OpenID connect hybrid` | `openid-hybrid` |

## ApplicationAlgorithm

| Name | Stored as |
| --- | --- |
| `No OIDC support` |  |
| `RSA with SHA-2 256` | `RS256` |
| `HMAC with SHA-2 256` | `HS256` |

## ApplicationRegistrationSource

| Name | Stored as |
| --- | --- |
| `Manual` | `manual` |
| `Dynamic Client Registration` | `dcr` |
| `Client ID Metadata Document` | `cimd` |

## GrantCodeChallengeMethod

| Name | Stored as |
| --- | --- |
| `plain` | `plain` |
| `S256` | `S256` |

## DeviceGrantStatus

| Name | Stored as |
| --- | --- |
| `Authorized` | `authorized` |
| `Authorization pending` | `authorization-pending` |
| `Expired` | `expired` |
| `Denied` | `denied` |
