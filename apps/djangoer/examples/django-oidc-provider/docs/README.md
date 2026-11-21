# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [ResponseType](#responsetype) · [Client](#client) · [Code](#code) · [Token](#token) · [UserConsent](#userconsent) · [RSAKey](#rsakey)

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

CREATE TABLE "oidc_provider_responsetype" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "value" varchar(30) NOT NULL UNIQUE, "description" varchar(50) NOT NULL);

CREATE TABLE "oidc_provider_client" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(100) NOT NULL, "owner_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "client_type" varchar(30) NOT NULL, "client_id" varchar(255) NOT NULL UNIQUE, "client_secret" varchar(255) NOT NULL, "jwt_alg" varchar(10) NOT NULL, "date_created" date NOT NULL, "website_url" varchar(255) NOT NULL, "terms_url" varchar(255) NOT NULL, "contact_email" varchar(255) NOT NULL, "logo" varchar(100) NOT NULL, "reuse_consent" bool NOT NULL, "require_consent" bool NOT NULL, "_redirect_uris" text NOT NULL, "_post_logout_redirect_uris" text NOT NULL, "_scope" text NOT NULL);

CREATE TABLE "oidc_provider_client_response_types" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "client_id" bigint NOT NULL REFERENCES "oidc_provider_client" ("id") DEFERRABLE INITIALLY DEFERRED, "responsetype_id" bigint NOT NULL REFERENCES "oidc_provider_responsetype" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "oidc_provider_code" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "client_id" bigint NOT NULL REFERENCES "oidc_provider_client" ("id") DEFERRABLE INITIALLY DEFERRED, "expires_at" datetime NOT NULL, "_scope" text NOT NULL, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "code" varchar(255) NOT NULL UNIQUE, "nonce" varchar(255) NOT NULL, "is_authentication" bool NOT NULL, "code_challenge" varchar(255) NULL, "code_challenge_method" varchar(255) NULL);

CREATE TABLE "oidc_provider_token" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "client_id" bigint NOT NULL REFERENCES "oidc_provider_client" ("id") DEFERRABLE INITIALLY DEFERRED, "expires_at" datetime NOT NULL, "_scope" text NOT NULL, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "access_token" varchar(255) NOT NULL UNIQUE, "refresh_token" varchar(255) NOT NULL UNIQUE, "_id_token" text NOT NULL);

CREATE TABLE "oidc_provider_userconsent" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "client_id" bigint NOT NULL REFERENCES "oidc_provider_client" ("id") DEFERRABLE INITIALLY DEFERRED, "expires_at" datetime NOT NULL, "_scope" text NOT NULL, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "date_given" datetime NOT NULL);

CREATE TABLE "oidc_provider_rsakey" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "key" text NOT NULL);

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

CREATE INDEX "oidc_provider_client_owner_id_735e6178" ON "oidc_provider_client" ("owner_id");

CREATE UNIQUE INDEX "oidc_provider_client_response_types_client_id_responsetype_id_0d116294_uniq" ON "oidc_provider_client_response_types" ("client_id", "responsetype_id");

CREATE INDEX "oidc_provider_client_response_types_client_id_576aac8c" ON "oidc_provider_client_response_types" ("client_id");

CREATE INDEX "oidc_provider_client_response_types_responsetype_id_c90751ad" ON "oidc_provider_client_response_types" ("responsetype_id");

CREATE INDEX "oidc_provider_code_client_id_2449f20c" ON "oidc_provider_code" ("client_id");

CREATE INDEX "oidc_provider_code_user_id_14775b3f" ON "oidc_provider_code" ("user_id");

CREATE INDEX "oidc_provider_token_client_id_86a72d16" ON "oidc_provider_token" ("client_id");

CREATE INDEX "oidc_provider_token_user_id_c61045ea" ON "oidc_provider_token" ("user_id");

CREATE UNIQUE INDEX "oidc_provider_userconsent_user_id_client_id_1c2032f6_uniq" ON "oidc_provider_userconsent" ("user_id", "client_id");

CREATE INDEX "oidc_provider_userconsent_client_id_7837dc61" ON "oidc_provider_userconsent" ("client_id");

CREATE INDEX "oidc_provider_userconsent_user_id_ebaddfd3" ON "oidc_provider_userconsent" ("user_id");
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
- [Client](#client).`owner_id` → `id`, on delete set null
- [Code](#code).`user_id` → `id`, on delete cascade
- [Token](#token).`user_id` → `id`, on delete cascade
- [UserConsent](#userconsent).`user_id` → `id`, on delete cascade

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

## ResponseType

Table `oidc_provider_responsetype`, app `oidc_provider`.

![ResponseType and the models it relates to](tables/ResponseType.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `value` | `ResponseTypeValue` | `varchar(30)` |  | unique |  |  |
| `description` | `CharField` | `varchar(50)` |  |  |  |  |

Relations:

- `client_set` — many to many with [Client](#client)

## Client

Table `oidc_provider_client`, app `oidc_provider`.

![Client and the models it relates to](tables/Client.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(100)` |  |  | `''` |  |
| `owner_id` | `ForeignKey` | `integer` | yes | FK | `None` |  |
| `client_type` | `ClientClientType` | `varchar(30)` |  |  | `'confidential'` | <b>Confidential</b> clients are capable of maintaining the confidentiality of their credentials. <b>Public</b> clients are incapable. |
| `client_id` | `CharField` | `varchar(255)` |  | unique |  |  |
| `client_secret` | `CharField` | `varchar(255)` |  |  |  |  |
| `jwt_alg` | `ClientJwtAlg` | `varchar(10)` |  |  | `'RS256'` | Algorithm used to encode ID Tokens. |
| `date_created` | `DateField` | `date` |  |  | `now()` |  |
| `website_url` | `CharField` | `varchar(255)` |  |  | `''` |  |
| `terms_url` | `CharField` | `varchar(255)` |  |  | `''` | External reference to the privacy policy of the client. |
| `contact_email` | `CharField` | `varchar(255)` |  |  | `''` |  |
| `logo` | `FileField` | `varchar(100)` |  |  | `''` |  |
| `reuse_consent` | `BooleanField` | `bool` |  |  | `True` | If enabled, server will save the user consent given to a specific client, so that user won't be prompted for the same authorization multiple times. |
| `require_consent` | `BooleanField` | `bool` |  |  | `True` | If disabled, the Server will NEVER ask the user for consent. |
| `_redirect_uris` | `TextField` | `text` |  |  | `''` | Enter each URI on a new line. |
| `_post_logout_redirect_uris` | `TextField` | `text` |  |  | `''` | Enter each URI on a new line. |
| `_scope` | `TextField` | `text` |  |  | `''` | Specifies the authorized scope values for the client app. |

Indexes:

- index `oidc_provider_client_owner_id_735e6178` on `owner_id`

Relations:

- `owner_id` → [User](#user).`id`, on delete set null
- `response_types` — many to many with [ResponseType](#responsetype)
- [Code](#code).`client_id` → `id`, on delete cascade
- [Token](#token).`client_id` → `id`, on delete cascade
- [UserConsent](#userconsent).`client_id` → `id`, on delete cascade

## Code

Table `oidc_provider_code`, app `oidc_provider`.

Columns from `BaseCodeTokenModel` (abstract): `client_id`, `expires_at`, `_scope`.

![Code and the models it relates to](tables/Code.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `client_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `expires_at` | `DateTimeField` | `datetime` |  |  |  |  |
| `_scope` | `TextField` | `text` |  |  | `''` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `code` | `CharField` | `varchar(255)` |  | unique |  |  |
| `nonce` | `CharField` | `varchar(255)` |  |  | `''` |  |
| `is_authentication` | `BooleanField` | `bool` |  |  | `False` |  |
| `code_challenge` | `CharField` | `varchar(255)` | yes |  |  |  |
| `code_challenge_method` | `CharField` | `varchar(255)` | yes |  |  |  |

Indexes:

- index `oidc_provider_code_client_id_2449f20c` on `client_id`
- index `oidc_provider_code_user_id_14775b3f` on `user_id`

Relations:

- `client_id` → [Client](#client).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## Token

Table `oidc_provider_token`, app `oidc_provider`.

Columns from `BaseCodeTokenModel` (abstract): `client_id`, `expires_at`, `_scope`.

![Token and the models it relates to](tables/Token.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `client_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `expires_at` | `DateTimeField` | `datetime` |  |  |  |  |
| `_scope` | `TextField` | `text` |  |  | `''` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `access_token` | `CharField` | `varchar(255)` |  | unique |  |  |
| `refresh_token` | `CharField` | `varchar(255)` |  | unique |  |  |
| `_id_token` | `TextField` | `text` |  |  |  |  |

Indexes:

- index `oidc_provider_token_client_id_86a72d16` on `client_id`
- index `oidc_provider_token_user_id_c61045ea` on `user_id`

Relations:

- `client_id` → [Client](#client).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## UserConsent

Table `oidc_provider_userconsent`, app `oidc_provider`.

Columns from `BaseCodeTokenModel` (abstract): `client_id`, `expires_at`, `_scope`.

![UserConsent and the models it relates to](tables/UserConsent.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `client_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `expires_at` | `DateTimeField` | `datetime` |  |  |  |  |
| `_scope` | `TextField` | `text` |  |  | `''` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `date_given` | `DateTimeField` | `datetime` |  |  |  |  |

Indexes:

- index `oidc_provider_userconsent_client_id_7837dc61` on `client_id`
- index `oidc_provider_userconsent_user_id_ebaddfd3` on `user_id`
- unique `oidc_provider_userconsent_user_id_client_id_1c2032f6_uniq` on `user_id`, `client_id`

Relations:

- `client_id` → [Client](#client).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## RSAKey

Table `oidc_provider_rsakey`, app `oidc_provider`.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `key` | `TextField` | `text` |  |  |  | Paste your private RSA Key here. |

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## ResponseTypeValue

| Name | Stored as |
| --- | --- |
| `code (Authorization Code Flow)` | `code` |
| `id_token (Implicit Flow)` | `id_token` |
| `id_token token (Implicit Flow)` | `id_token token` |
| `code token (Hybrid Flow)` | `code token` |
| `code id_token (Hybrid Flow)` | `code id_token` |
| `code id_token token (Hybrid Flow)` | `code id_token token` |

## ClientClientType

| Name | Stored as |
| --- | --- |
| `Confidential` | `confidential` |
| `Public` | `public` |

## ClientJwtAlg

| Name | Stored as |
| --- | --- |
| `HS256` | `HS256` |
| `RS256` | `RS256` |
