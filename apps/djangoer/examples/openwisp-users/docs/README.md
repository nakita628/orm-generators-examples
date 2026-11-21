# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [ContentType](#contenttype) · [Session](#session) · [Site](#site) · [EmailAddress](#emailaddress) · [EmailConfirmation](#emailconfirmation) · [SocialApp](#socialapp) · [SocialAccount](#socialaccount) · [SocialToken](#socialtoken) · [User](#user) · [Organization](#organization) · [OrganizationUser](#organizationuser) · [OrganizationOwner](#organizationowner) · [OrganizationInvitation](#organizationinvitation) · [Token](#token)

<details>
<summary>The whole schema as SQL</summary>

```sql
CREATE TABLE "django_admin_log" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "action_time" datetime NOT NULL, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "object_id" text NULL, "object_repr" varchar(200) NOT NULL, "action_flag" smallint unsigned NOT NULL CHECK ("action_flag" >= 0), "change_message" text NOT NULL);

CREATE TABLE "auth_permission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(255) NOT NULL, "content_type_id" integer NOT NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "codename" varchar(100) NOT NULL);

CREATE TABLE "auth_group" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(150) NOT NULL UNIQUE);

CREATE TABLE "auth_group_permissions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED, "permission_id" integer NOT NULL REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "django_content_type" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "app_label" varchar(100) NOT NULL, "model" varchar(100) NOT NULL);

CREATE TABLE "django_session" ("session_key" varchar(40) NOT NULL PRIMARY KEY, "session_data" text NOT NULL, "expire_date" datetime NOT NULL);

CREATE TABLE "django_site" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "domain" varchar(100) NOT NULL UNIQUE, "name" varchar(50) NOT NULL);

CREATE TABLE "account_emailaddress" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "email" varchar(254) NOT NULL, "verified" bool NOT NULL, "primary" bool NOT NULL);

CREATE TABLE "account_emailconfirmation" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "email_address_id" integer NOT NULL REFERENCES "account_emailaddress" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL, "sent" datetime NULL, "key" varchar(64) NOT NULL UNIQUE);

CREATE TABLE "socialaccount_socialapp" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "provider" varchar(30) NOT NULL, "provider_id" varchar(200) NOT NULL, "name" varchar(40) NOT NULL, "client_id" varchar(191) NOT NULL, "secret" varchar(191) NOT NULL, "key" varchar(191) NOT NULL, "settings" text NOT NULL CHECK ((JSON_VALID("settings") OR "settings" IS NULL)));

CREATE TABLE "socialaccount_socialapp_sites" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "socialapp_id" integer NOT NULL REFERENCES "socialaccount_socialapp" ("id") DEFERRABLE INITIALLY DEFERRED, "site_id" integer NOT NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "socialaccount_socialaccount" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "provider" varchar(200) NOT NULL, "uid" varchar(191) NOT NULL, "last_login" datetime NOT NULL, "date_joined" datetime NOT NULL, "extra_data" text NOT NULL CHECK ((JSON_VALID("extra_data") OR "extra_data" IS NULL)));

CREATE TABLE "socialaccount_socialtoken" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "app_id" integer NULL REFERENCES "socialaccount_socialapp" ("id") DEFERRABLE INITIALLY DEFERRED, "account_id" integer NOT NULL REFERENCES "socialaccount_socialaccount" ("id") DEFERRABLE INITIALLY DEFERRED, "token" text NOT NULL, "token_secret" text NOT NULL, "expires_at" datetime NULL);

CREATE TABLE "openwisp_users_user" ("password" varchar(128) NOT NULL, "last_login" datetime NULL, "is_superuser" bool NOT NULL, "username" varchar(150) NOT NULL UNIQUE, "first_name" varchar(150) NOT NULL, "last_name" varchar(150) NOT NULL, "is_staff" bool NOT NULL, "is_active" bool NOT NULL, "date_joined" datetime NOT NULL, "id" char(32) NOT NULL PRIMARY KEY, "email" varchar(254) NULL UNIQUE, "bio" text NOT NULL, "url" varchar(200) NOT NULL, "company" varchar(30) NOT NULL, "location" varchar(256) NOT NULL, "phone_number" varchar(128) NULL UNIQUE, "birth_date" date NULL, "notes" text NOT NULL, "language" varchar(8) NOT NULL, "password_updated" date NULL, "password_based_token" bool NULL, "expiration_date" date NULL);

CREATE TABLE "openwisp_users_user_groups" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "openwisp_users_user_user_permissions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "permission_id" integer NOT NULL REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "openwisp_users_organization" ("name" varchar(200) NOT NULL, "is_active" bool NOT NULL, "created" datetime NOT NULL, "modified" datetime NOT NULL, "slug" varchar(200) NOT NULL UNIQUE, "id" char(32) NOT NULL PRIMARY KEY, "description" text NOT NULL, "email" varchar(254) NOT NULL, "url" varchar(200) NOT NULL);

CREATE TABLE "openwisp_users_organizationuser" ("created" datetime NOT NULL, "modified" datetime NOT NULL, "id" char(32) NOT NULL PRIMARY KEY, "is_admin" bool NOT NULL, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "openwisp_users_organizationowner" ("created" datetime NOT NULL, "modified" datetime NOT NULL, "id" char(32) NOT NULL PRIMARY KEY, "organization_user_id" char(32) NOT NULL UNIQUE REFERENCES "openwisp_users_organizationuser" ("id") DEFERRABLE INITIALLY DEFERRED, "organization_id" char(32) NOT NULL UNIQUE REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "openwisp_users_organizationinvitation" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "guid" char(32) NOT NULL, "invitee_identifier" varchar(1000) NOT NULL, "created" datetime NOT NULL, "modified" datetime NOT NULL, "invited_by_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "invitee_id" char(32) NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "authtoken_token" ("key" varchar(40) NOT NULL PRIMARY KEY, "user_id" char(32) NOT NULL UNIQUE REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL);

CREATE INDEX "django_admin_log_user_id_c564eba6" ON "django_admin_log" ("user_id");

CREATE INDEX "django_admin_log_content_type_id_c4bce8eb" ON "django_admin_log" ("content_type_id");

CREATE UNIQUE INDEX "auth_permission_content_type_id_codename_01ab375a_uniq" ON "auth_permission" ("content_type_id", "codename");

CREATE INDEX "auth_permission_content_type_id_2f476e4b" ON "auth_permission" ("content_type_id");

CREATE UNIQUE INDEX "auth_group_permissions_group_id_permission_id_0cd325b0_uniq" ON "auth_group_permissions" ("group_id", "permission_id");

CREATE INDEX "auth_group_permissions_group_id_b120cbf9" ON "auth_group_permissions" ("group_id");

CREATE INDEX "auth_group_permissions_permission_id_84c5c92e" ON "auth_group_permissions" ("permission_id");

CREATE UNIQUE INDEX "django_content_type_app_label_model_76bd3d3b_uniq" ON "django_content_type" ("app_label", "model");

CREATE INDEX "django_session_expire_date_a5c62663" ON "django_session" ("expire_date");

CREATE UNIQUE INDEX "account_emailaddress_user_id_email_987c8728_uniq" ON "account_emailaddress" ("user_id", "email");

CREATE UNIQUE INDEX "unique_primary_email" ON "account_emailaddress" ("user_id", "primary") WHERE "primary" = 1;

CREATE UNIQUE INDEX "unique_verified_email" ON "account_emailaddress" ("email") WHERE "verified" = 1;

CREATE INDEX "account_emailaddress_user_id_2c513194" ON "account_emailaddress" ("user_id");

CREATE INDEX "account_emailaddress_email_03be32b2" ON "account_emailaddress" ("email");

CREATE INDEX "account_emailconfirmation_email_address_id_5b7f8c58" ON "account_emailconfirmation" ("email_address_id");

CREATE UNIQUE INDEX "socialaccount_socialapp_sites_socialapp_id_site_id_71a9a768_uniq" ON "socialaccount_socialapp_sites" ("socialapp_id", "site_id");

CREATE INDEX "socialaccount_socialapp_sites_socialapp_id_97fb6e7d" ON "socialaccount_socialapp_sites" ("socialapp_id");

CREATE INDEX "socialaccount_socialapp_sites_site_id_2579dee5" ON "socialaccount_socialapp_sites" ("site_id");

CREATE UNIQUE INDEX "socialaccount_socialaccount_provider_uid_fc810c6e_uniq" ON "socialaccount_socialaccount" ("provider", "uid");

CREATE INDEX "socialaccount_socialaccount_user_id_8146e70c" ON "socialaccount_socialaccount" ("user_id");

CREATE UNIQUE INDEX "socialaccount_socialtoken_app_id_account_id_fca4e0ac_uniq" ON "socialaccount_socialtoken" ("app_id", "account_id");

CREATE INDEX "socialaccount_socialtoken_app_id_636a42d7" ON "socialaccount_socialtoken" ("app_id");

CREATE INDEX "socialaccount_socialtoken_account_id_951f210e" ON "socialaccount_socialtoken" ("account_id");

CREATE INDEX "user_id_email_idx" ON "openwisp_users_user" ("id", "email");

CREATE INDEX "user_active_expiry_idx" ON "openwisp_users_user" ("is_active", "expiration_date");

CREATE UNIQUE INDEX "openwisp_users_user_groups_user_id_group_id_7b92002e_uniq" ON "openwisp_users_user_groups" ("user_id", "group_id");

CREATE INDEX "openwisp_users_user_groups_user_id_2371794b" ON "openwisp_users_user_groups" ("user_id");

CREATE INDEX "openwisp_users_user_groups_group_id_e9ab4901" ON "openwisp_users_user_groups" ("group_id");

CREATE UNIQUE INDEX "openwisp_users_user_user_permissions_user_id_permission_id_24ba4ccf_uniq" ON "openwisp_users_user_user_permissions" ("user_id", "permission_id");

CREATE INDEX "openwisp_users_user_user_permissions_user_id_f3610345" ON "openwisp_users_user_user_permissions" ("user_id");

CREATE INDEX "openwisp_users_user_user_permissions_permission_id_f00dc354" ON "openwisp_users_user_user_permissions" ("permission_id");

CREATE UNIQUE INDEX "openwisp_users_organizationuser_user_id_organization_id_d40e96a2_uniq" ON "openwisp_users_organizationuser" ("user_id", "organization_id");

CREATE INDEX "openwisp_users_organizationuser_user_id_a5799aa4" ON "openwisp_users_organizationuser" ("user_id");

CREATE INDEX "openwisp_users_organizationuser_organization_id_a1c43bbd" ON "openwisp_users_organizationuser" ("organization_id");

CREATE INDEX "openwisp_users_organizationinvitation_invited_by_id_69a7098a" ON "openwisp_users_organizationinvitation" ("invited_by_id");

CREATE INDEX "openwisp_users_organizationinvitation_invitee_id_d3c1b419" ON "openwisp_users_organizationinvitation" ("invitee_id");

CREATE INDEX "openwisp_users_organizationinvitation_organization_id_17466f8f" ON "openwisp_users_organizationinvitation" ("organization_id");
```

</details>

## LogEntry

Table `django_admin_log`, app `admin`.

![LogEntry and the models it relates to](tables/LogEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `action_time` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
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

Proxy models, on this table: `Group`.

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

## Site

Table `django_site`, app `sites`.

![Site and the models it relates to](tables/Site.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `domain` | `CharField` | `varchar(100)` |  | unique |  |  |
| `name` | `CharField` | `varchar(50)` |  |  |  |  |

Relations:

- `socialapp_set` — many to many with [SocialApp](#socialapp)

## EmailAddress

Table `account_emailaddress`, app `account`.

![EmailAddress and the models it relates to](tables/EmailAddress.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `verified` | `BooleanField` | `bool` |  |  | `False` |  |
| `primary` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `account_emailaddress_user_id_2c513194` on `user_id`
- index `account_emailaddress_email_03be32b2` on `email`
- unique `account_emailaddress_user_id_email_987c8728_uniq` on `user_id`, `email`
- unique `unique_primary_email` on `user_id`, `primary` where `"primary" = 1`
- unique `unique_verified_email` on `email` where `"verified" = 1`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [EmailConfirmation](#emailconfirmation).`email_address_id` → `id`, on delete cascade

## EmailConfirmation

Table `account_emailconfirmation`, app `account`.

![EmailConfirmation and the models it relates to](tables/EmailConfirmation.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `email_address_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `sent` | `DateTimeField` | `datetime` | yes |  |  |  |
| `key` | `CharField` | `varchar(64)` |  | unique |  |  |

Indexes:

- index `account_emailconfirmation_email_address_id_5b7f8c58` on `email_address_id`

Relations:

- `email_address_id` → [EmailAddress](#emailaddress).`id`, on delete cascade

## SocialApp

Table `socialaccount_socialapp`, app `socialaccount`.

![SocialApp and the models it relates to](tables/SocialApp.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `provider` | `CharField` | `varchar(30)` |  |  |  |  |
| `provider_id` | `CharField` | `varchar(200)` |  |  |  |  |
| `name` | `CharField` | `varchar(40)` |  |  |  |  |
| `client_id` | `CharField` | `varchar(191)` |  |  |  | App ID, or consumer key |
| `secret` | `CharField` | `varchar(191)` |  |  |  | API secret, client secret, or consumer secret |
| `key` | `CharField` | `varchar(191)` |  |  |  | Key |
| `settings` | `JSONField` | `text` |  |  | `dict()` |  |

Relations:

- `sites` — many to many with [Site](#site)
- [SocialToken](#socialtoken).`app_id` → `id`, on delete set null

## SocialAccount

Table `socialaccount_socialaccount`, app `socialaccount`.

![SocialAccount and the models it relates to](tables/SocialAccount.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `provider` | `CharField` | `varchar(200)` |  |  |  |  |
| `uid` | `CharField` | `varchar(191)` |  |  |  |  |
| `last_login` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `date_joined` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `extra_data` | `JSONField` | `text` |  |  | `dict()` |  |

Indexes:

- index `socialaccount_socialaccount_user_id_8146e70c` on `user_id`
- unique `socialaccount_socialaccount_provider_uid_fc810c6e_uniq` on `provider`, `uid`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [SocialToken](#socialtoken).`account_id` → `id`, on delete cascade

## SocialToken

Table `socialaccount_socialtoken`, app `socialaccount`.

![SocialToken and the models it relates to](tables/SocialToken.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `app_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `account_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `token` | `TextField` | `text` |  |  |  | "oauth_token" (OAuth1) or access token (OAuth2) |
| `token_secret` | `TextField` | `text` |  |  |  | "oauth_token_secret" (OAuth1) or refresh token (OAuth2) |
| `expires_at` | `DateTimeField` | `datetime` | yes |  |  |  |

Indexes:

- index `socialaccount_socialtoken_app_id_636a42d7` on `app_id`
- index `socialaccount_socialtoken_account_id_951f210e` on `account_id`
- unique `socialaccount_socialtoken_app_id_account_id_fca4e0ac_uniq` on `app_id`, `account_id`

Relations:

- `app_id` → [SocialApp](#socialapp).`id`, on delete set null
- `account_id` → [SocialAccount](#socialaccount).`id`, on delete cascade

## User

Table `openwisp_users_user`, app `openwisp_users`.

Columns from `AbstractUser` (abstract): `id`, `bio`, `url`, `company`, `location`, `phone_number`, `birth_date`, `notes`, `language`, `password_updated`, `password_based_token`, `expiration_date`.

Columns from `AbstractUser` (abstract): `username`, `first_name`, `last_name`, `email`, `is_staff`, `is_active`, `date_joined`.

Columns from `AbstractBaseUser` (abstract): `password`, `last_login`.

Columns from `PermissionsMixin` (abstract): `is_superuser`.

![User and the models it relates to](tables/User.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `password` | `CharField` | `varchar(128)` |  |  |  |  |
| `last_login` | `DateTimeField` | `datetime` | yes |  |  |  |
| `is_superuser` | `BooleanField` | `bool` |  |  | `False` | Designates that this user has all permissions without explicitly assigning them. |
| `username` | `CharField` | `varchar(150)` |  | unique |  | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. |
| `first_name` | `CharField` | `varchar(150)` |  |  |  |  |
| `last_name` | `CharField` | `varchar(150)` |  |  |  |  |
| `is_staff` | `BooleanField` | `bool` |  |  | `False` | Designates whether the user can log into this admin site. |
| `is_active` | `BooleanField` | `bool` |  |  | `True` | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. |
| `date_joined` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `email` | `EmailField` | `varchar(254)` | yes | unique |  |  |
| `bio` | `TextField` | `text` |  |  |  |  |
| `url` | `URLField` | `varchar(200)` |  |  |  |  |
| `company` | `CharField` | `varchar(30)` |  |  |  |  |
| `location` | `CharField` | `varchar(256)` |  |  |  |  |
| `phone_number` | `PhoneNumberField` | `varchar(128)` | yes | unique |  |  |
| `birth_date` | `DateField` | `date` | yes |  |  |  |
| `notes` | `TextField` | `text` |  |  |  | notes for internal usage |
| `language` | `UserLanguage` | `varchar(8)` |  |  | `'en-us'` |  |
| `password_updated` | `DateField` | `date` | yes |  |  |  |
| `password_based_token` | `BooleanField` | `bool` | yes |  | `None` | Indicates whether the last authentication token was obtained using the local password. When false, the token came from an external method (eg: SSO, SAML) and password expiration is not enforced for it. None means no token has been issued for this user since this feature was introduced. |
| `expiration_date` | `DateField` | `date` | yes |  |  | If set, the account will be deactivated after this date passes and the user will no longer be able to log in. |

Indexes:

- index `user_id_email_idx` on `id`, `email`
- index `user_active_expiry_idx` on `is_active`, `expiration_date`

Relations:

- `openwisp_users_organization` — many to many with [Organization](#organization), through [OrganizationUser](#organizationuser)
- [LogEntry](#logentry).`user_id` → `id`, on delete cascade
- [EmailAddress](#emailaddress).`user_id` → `id`, on delete cascade
- [SocialAccount](#socialaccount).`user_id` → `id`, on delete cascade
- `groups` — many to many with [Group](#group)
- `user_permissions` — many to many with [Permission](#permission)
- [OrganizationUser](#organizationuser).`user_id` → `id`, on delete cascade
- [OrganizationInvitation](#organizationinvitation).`invited_by_id` → `id`, on delete cascade
- [OrganizationInvitation](#organizationinvitation).`invitee_id` → `id`, on delete cascade
- [Token](#token).`user_id` → `id`, on delete cascade

## Organization

Table `openwisp_users_organization`, app `openwisp_users`.

Columns from `BaseOrganization` (abstract): `id`, `description`, `email`, `url`.

Columns from `AbstractOrganization` (abstract): `slug`.

Columns from `SharedBaseModel` (abstract): `created`, `modified`.

Columns from `AbstractBaseOrganization` (abstract): `name`, `is_active`.

![Organization and the models it relates to](tables/Organization.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `name` | `CharField` | `varchar(200)` |  |  |  | The name of the organization |
| `is_active` | `BooleanField` | `bool` |  |  | `True` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `slug` | `SlugField` | `varchar(200)` |  | unique |  | The name in all lowercase, suitable for URL identification |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `url` | `URLField` | `varchar(200)` |  |  |  |  |

Relations:

- `users` — many to many with [User](#user), through [OrganizationUser](#organizationuser)
- [OrganizationUser](#organizationuser).`organization_id` → `id`, on delete cascade
- [OrganizationOwner](#organizationowner).`organization_id` → `id`, on delete cascade
- [OrganizationInvitation](#organizationinvitation).`organization_id` → `id`, on delete cascade

## OrganizationUser

Table `openwisp_users_organizationuser`, app `openwisp_users`.

Columns from `BaseOrganizationUser` (abstract): `id`.

Columns from `AbstractOrganizationUser` (abstract): `is_admin`.

Columns from `SharedBaseModel` (abstract): `created`, `modified`.

![OrganizationUser and the models it relates to](tables/OrganizationUser.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `is_admin` | `BooleanField` | `bool` |  |  | `False` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |

Indexes:

- index `openwisp_users_organizationuser_user_id_a5799aa4` on `user_id`
- index `openwisp_users_organizationuser_organization_id_a1c43bbd` on `organization_id`
- unique `openwisp_users_organizationuser_user_id_organization_id_d40e96a2_uniq` on `user_id`, `organization_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `organization_id` → [Organization](#organization).`id`, on delete cascade
- [OrganizationOwner](#organizationowner).`organization_user_id` → `id`, on delete cascade

## OrganizationOwner

Table `openwisp_users_organizationowner`, app `openwisp_users`.

Columns from `BaseOrganizationOwner` (abstract): `id`.

Columns from `SharedBaseModel` (abstract): `created`, `modified`.

![OrganizationOwner and the models it relates to](tables/OrganizationOwner.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `organization_user_id` | `OneToOneField` | `char(32)` |  | FK, unique |  |  |
| `organization_id` | `OneToOneField` | `char(32)` |  | FK, unique |  |  |

Relations:

- `organization_user_id` → [OrganizationUser](#organizationuser).`id`, on delete cascade
- `organization_id` → [Organization](#organization).`id`, on delete cascade

## OrganizationInvitation

Table `openwisp_users_organizationinvitation`, app `openwisp_users`.

Columns from `SharedBaseModel` (abstract): `created`, `modified`.

Columns from `AbstractBaseInvitation` (abstract): `guid`, `invitee_identifier`.

![OrganizationInvitation and the models it relates to](tables/OrganizationInvitation.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `guid` | `UUIDField` | `char(32)` |  |  |  |  |
| `invitee_identifier` | `CharField` | `varchar(1000)` |  |  |  | The contact identifier for the invitee, email, phone number, social media handle, etc. |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `invited_by_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `invitee_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |

Indexes:

- index `openwisp_users_organizationinvitation_invited_by_id_69a7098a` on `invited_by_id`
- index `openwisp_users_organizationinvitation_invitee_id_d3c1b419` on `invitee_id`
- index `openwisp_users_organizationinvitation_organization_id_17466f8f` on `organization_id`

Relations:

- `invited_by_id` → [User](#user).`id`, on delete cascade
- `invitee_id` → [User](#user).`id`, on delete cascade
- `organization_id` → [Organization](#organization).`id`, on delete cascade

## Token

Table `authtoken_token`, app `authtoken`.

Proxy models, on this table: `ApiKey`, `TokenProxy`.

The default authorization token model.

![Token and the models it relates to](tables/Token.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `key` | `CharField` | `varchar(40)` |  | PK |  |  |
| `user_id` | `OneToOneField` | `char(32)` |  | FK, unique |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## UserLanguage

| Name | Stored as |
| --- | --- |
| `Afrikaans` | `af` |
| `Arabic` | `ar` |
| `Algerian Arabic` | `ar-dz` |
| `Asturian` | `ast` |
| `Azerbaijani` | `az` |
| `Bulgarian` | `bg` |
| `Belarusian` | `be` |
| `Bengali` | `bn` |
| `Breton` | `br` |
| `Bosnian` | `bs` |
| `Catalan` | `ca` |
| `Central Kurdish (Sorani)` | `ckb` |
| `Czech` | `cs` |
| `Welsh` | `cy` |
| `Danish` | `da` |
| `German` | `de` |
| `Lower Sorbian` | `dsb` |
| `Greek` | `el` |
| `English` | `en` |
| `Australian English` | `en-au` |
| `British English` | `en-gb` |
| `Esperanto` | `eo` |
| `Spanish` | `es` |
| `Argentinian Spanish` | `es-ar` |
| `Colombian Spanish` | `es-co` |
| `Mexican Spanish` | `es-mx` |
| `Nicaraguan Spanish` | `es-ni` |
| `Venezuelan Spanish` | `es-ve` |
| `Estonian` | `et` |
| `Basque` | `eu` |
| `Persian` | `fa` |
| `Finnish` | `fi` |
| `French` | `fr` |
| `Frisian` | `fy` |
| `Irish` | `ga` |
| `Scottish Gaelic` | `gd` |
| `Galician` | `gl` |
| `Hebrew` | `he` |
| `Hindi` | `hi` |
| `Croatian` | `hr` |
| `Upper Sorbian` | `hsb` |
| `Haitian Creole` | `ht` |
| `Hungarian` | `hu` |
| `Armenian` | `hy` |
| `Interlingua` | `ia` |
| `Indonesian` | `id` |
| `Igbo` | `ig` |
| `Ido` | `io` |
| `Icelandic` | `is` |
| `Italian` | `it` |
| `Japanese` | `ja` |
| `Georgian` | `ka` |
| `Kabyle` | `kab` |
| `Kazakh` | `kk` |
| `Khmer` | `km` |
| `Kannada` | `kn` |
| `Korean` | `ko` |
| `Kyrgyz` | `ky` |
| `Luxembourgish` | `lb` |
| `Lithuanian` | `lt` |
| `Latvian` | `lv` |
| `Macedonian` | `mk` |
| `Malayalam` | `ml` |
| `Mongolian` | `mn` |
| `Marathi` | `mr` |
| `Malay` | `ms` |
| `Burmese` | `my` |
| `Norwegian Bokmål` | `nb` |
| `Nepali` | `ne` |
| `Dutch` | `nl` |
| `Norwegian Nynorsk` | `nn` |
| `Ossetic` | `os` |
| `Punjabi` | `pa` |
| `Polish` | `pl` |
| `Portuguese` | `pt` |
| `Brazilian Portuguese` | `pt-br` |
| `Romanian` | `ro` |
| `Russian` | `ru` |
| `Slovak` | `sk` |
| `Slovenian` | `sl` |
| `Albanian` | `sq` |
| `Serbian` | `sr` |
| `Serbian Latin` | `sr-latn` |
| `Swedish` | `sv` |
| `Swahili` | `sw` |
| `Tamil` | `ta` |
| `Telugu` | `te` |
| `Tajik` | `tg` |
| `Thai` | `th` |
| `Turkmen` | `tk` |
| `Turkish` | `tr` |
| `Tatar` | `tt` |
| `Udmurt` | `udm` |
| `Uyghur` | `ug` |
| `Ukrainian` | `uk` |
| `Urdu` | `ur` |
| `Uzbek` | `uz` |
| `Vietnamese` | `vi` |
| `Simplified Chinese` | `zh-hans` |
| `Traditional Chinese` | `zh-hant` |
