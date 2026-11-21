# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [DatabaseConnection](#databaseconnection) · [PromptLog](#promptlog) · [TableDescription](#tabledescription) · [Query](#query) · [QueryLog](#querylog) · [QueryFavorite](#queryfavorite) · [ExplorerValue](#explorervalue)

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

CREATE TABLE "explorer_databaseconnection" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "alias" varchar(255) NOT NULL UNIQUE, "engine" varchar(255) NOT NULL, "name" varchar(255) NULL, "user" BLOB NULL, "password" BLOB NULL, "host" BLOB NOT NULL, "port" varchar(255) NULL, "extras" text NULL CHECK ((JSON_VALID("extras") OR "extras" IS NULL)), "upload_fingerprint" varchar(255) NULL, "default" bool NOT NULL);

CREATE TABLE "explorer_promptlog" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "prompt" text NOT NULL, "user_request" text NOT NULL, "response" text NOT NULL, "run_by_user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "run_at" datetime NOT NULL, "duration" real NULL, "model" varchar(128) NOT NULL, "error" text NULL, "database_connection_id" integer NULL REFERENCES "explorer_databaseconnection" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "explorer_tabledescription" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "database_connection_id" integer NOT NULL REFERENCES "explorer_databaseconnection" ("id") DEFERRABLE INITIALLY DEFERRED, "table_name" varchar(512) NOT NULL, "description" text NOT NULL);

CREATE TABLE "explorer_query" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "title" varchar(255) NOT NULL, "sql" text NOT NULL, "description" text NOT NULL, "created_by_user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created_at" datetime NOT NULL, "last_run_date" datetime NOT NULL, "snapshot" bool NOT NULL, "connection" varchar(128) NOT NULL, "database_connection_id" integer NULL REFERENCES "explorer_databaseconnection" ("id") DEFERRABLE INITIALLY DEFERRED, "few_shot" bool NOT NULL);

CREATE TABLE "explorer_querylog" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "sql" text NOT NULL, "query_id" integer NULL REFERENCES "explorer_query" ("id") DEFERRABLE INITIALLY DEFERRED, "run_by_user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "run_at" datetime NOT NULL, "duration" real NULL, "connection" varchar(128) NOT NULL, "database_connection_id" integer NULL REFERENCES "explorer_databaseconnection" ("id") DEFERRABLE INITIALLY DEFERRED, "success" bool NOT NULL, "error" text NULL);

CREATE TABLE "explorer_queryfavorite" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "query_id" integer NOT NULL REFERENCES "explorer_query" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "explorer_explorervalue" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "key" varchar(5) NOT NULL UNIQUE, "value" text NULL);

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

CREATE INDEX "explorer_promptlog_run_by_user_id_ec301ea7" ON "explorer_promptlog" ("run_by_user_id");

CREATE INDEX "explorer_promptlog_database_connection_id_eb15d329" ON "explorer_promptlog" ("database_connection_id");

CREATE UNIQUE INDEX "explorer_tabledescription_database_connection_id_table_name_73ac361d_uniq" ON "explorer_tabledescription" ("database_connection_id", "table_name");

CREATE INDEX "explorer_tabledescription_database_connection_id_16e24fab" ON "explorer_tabledescription" ("database_connection_id");

CREATE INDEX "explorer_query_created_by_user_id_182dd868" ON "explorer_query" ("created_by_user_id");

CREATE INDEX "explorer_query_database_connection_id_64440483" ON "explorer_query" ("database_connection_id");

CREATE INDEX "explorer_querylog_query_id_1635a6b4" ON "explorer_querylog" ("query_id");

CREATE INDEX "explorer_querylog_run_by_user_id_cf26a49f" ON "explorer_querylog" ("run_by_user_id");

CREATE INDEX "explorer_querylog_database_connection_id_cf8ddd6a" ON "explorer_querylog" ("database_connection_id");

CREATE UNIQUE INDEX "explorer_queryfavorite_query_id_user_id_602b82de_uniq" ON "explorer_queryfavorite" ("query_id", "user_id");

CREATE INDEX "explorer_queryfavorite_query_id_1f4e3cf7" ON "explorer_queryfavorite" ("query_id");

CREATE INDEX "explorer_queryfavorite_user_id_5ef2792c" ON "explorer_queryfavorite" ("user_id");
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
- [PromptLog](#promptlog).`run_by_user_id` → `id`, on delete cascade
- [Query](#query).`created_by_user_id` → `id`, on delete cascade
- [QueryLog](#querylog).`run_by_user_id` → `id`, on delete cascade
- [QueryFavorite](#queryfavorite).`user_id` → `id`, on delete cascade

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

## DatabaseConnection

Table `explorer_databaseconnection`, app `explorer`.

![DatabaseConnection and the models it relates to](tables/DatabaseConnection.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `alias` | `CharField` | `varchar(255)` |  | unique |  |  |
| `engine` | `DatabaseConnectionEngine` | `varchar(255)` |  |  |  |  |
| `name` | `CharField` | `varchar(255)` | yes |  |  |  |
| `user` | `EncryptedCharField` | `BLOB` | yes |  |  |  |
| `password` | `EncryptedCharField` | `BLOB` | yes |  |  |  |
| `host` | `EncryptedCharField` | `BLOB` |  |  |  |  |
| `port` | `CharField` | `varchar(255)` | yes |  |  |  |
| `extras` | `JSONField` | `text` | yes |  |  |  |
| `upload_fingerprint` | `CharField` | `varchar(255)` | yes |  |  |  |
| `default` | `BooleanField` | `bool` |  |  | `False` |  |

Relations:

- [PromptLog](#promptlog).`database_connection_id` → `id`, on delete set null
- [TableDescription](#tabledescription).`database_connection_id` → `id`, on delete cascade
- [Query](#query).`database_connection_id` → `id`, on delete set null
- [QueryLog](#querylog).`database_connection_id` → `id`, on delete set null

## PromptLog

Table `explorer_promptlog`, app `explorer`.

![PromptLog and the models it relates to](tables/PromptLog.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `prompt` | `TextField` | `text` |  |  |  |  |
| `user_request` | `TextField` | `text` |  |  |  |  |
| `response` | `TextField` | `text` |  |  |  |  |
| `run_by_user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `run_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `duration` | `FloatField` | `real` | yes |  |  |  |
| `model` | `CharField` | `varchar(128)` |  |  | `''` |  |
| `error` | `TextField` | `text` | yes |  |  |  |
| `database_connection_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `explorer_promptlog_run_by_user_id_ec301ea7` on `run_by_user_id`
- index `explorer_promptlog_database_connection_id_eb15d329` on `database_connection_id`

Relations:

- `run_by_user_id` → [User](#user).`id`, on delete cascade
- `database_connection_id` → [DatabaseConnection](#databaseconnection).`id`, on delete set null

## TableDescription

Table `explorer_tabledescription`, app `explorer`.

![TableDescription and the models it relates to](tables/TableDescription.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `database_connection_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `table_name` | `CharField` | `varchar(512)` |  |  |  |  |
| `description` | `TextField` | `text` |  |  |  |  |

Indexes:

- index `explorer_tabledescription_database_connection_id_16e24fab` on `database_connection_id`
- unique `explorer_tabledescription_database_connection_id_table_name_73ac361d_uniq` on `database_connection_id`, `table_name`

Relations:

- `database_connection_id` → [DatabaseConnection](#databaseconnection).`id`, on delete cascade

## Query

Table `explorer_query`, app `explorer`.

![Query and the models it relates to](tables/Query.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `title` | `CharField` | `varchar(255)` |  |  |  |  |
| `sql` | `TextField` | `text` |  |  |  |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `created_by_user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `last_run_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `snapshot` | `BooleanField` | `bool` |  |  | `False` | Include in snapshot task (if enabled) |
| `connection` | `CharField` | `varchar(128)` |  |  | `''` | Name of DB connection (as specified in settings) to use for this query.Will use EXPLORER_DEFAULT_CONNECTION if left blank |
| `database_connection_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `few_shot` | `BooleanField` | `bool` |  |  | `False` | Will be included as a good example of SQL in assistant queries that use relevant tables |

Indexes:

- index `explorer_query_created_by_user_id_182dd868` on `created_by_user_id`
- index `explorer_query_database_connection_id_64440483` on `database_connection_id`

Relations:

- `created_by_user_id` → [User](#user).`id`, on delete cascade
- `database_connection_id` → [DatabaseConnection](#databaseconnection).`id`, on delete set null
- [QueryLog](#querylog).`query_id` → `id`, on delete set null
- [QueryFavorite](#queryfavorite).`query_id` → `id`, on delete cascade

## QueryLog

Table `explorer_querylog`, app `explorer`.

![QueryLog and the models it relates to](tables/QueryLog.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `sql` | `TextField` | `text` |  |  |  |  |
| `query_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `run_by_user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `run_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `duration` | `FloatField` | `real` | yes |  |  |  |
| `connection` | `CharField` | `varchar(128)` |  |  | `''` |  |
| `database_connection_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `success` | `BooleanField` | `bool` |  |  | `True` |  |
| `error` | `TextField` | `text` | yes |  |  |  |

Indexes:

- index `explorer_querylog_query_id_1635a6b4` on `query_id`
- index `explorer_querylog_run_by_user_id_cf26a49f` on `run_by_user_id`
- index `explorer_querylog_database_connection_id_cf8ddd6a` on `database_connection_id`

Relations:

- `query_id` → [Query](#query).`id`, on delete set null
- `run_by_user_id` → [User](#user).`id`, on delete cascade
- `database_connection_id` → [DatabaseConnection](#databaseconnection).`id`, on delete set null

## QueryFavorite

Table `explorer_queryfavorite`, app `explorer`.

![QueryFavorite and the models it relates to](tables/QueryFavorite.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `query_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |

Indexes:

- index `explorer_queryfavorite_query_id_1f4e3cf7` on `query_id`
- index `explorer_queryfavorite_user_id_5ef2792c` on `user_id`
- unique `explorer_queryfavorite_query_id_user_id_602b82de_uniq` on `query_id`, `user_id`

Relations:

- `query_id` → [Query](#query).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## ExplorerValue

Table `explorer_explorervalue`, app `explorer`.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `key` | `ExplorerValueKey` | `varchar(5)` |  | unique |  |  |
| `value` | `TextField` | `text` | yes |  |  |  |

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## DatabaseConnectionEngine

| Name | Stored as |
| --- | --- |
| `SQLite3` | `django.db.backends.sqlite3` |
| `PostgreSQL` | `django.db.backends.postgresql` |
| `MySQL` | `django.db.backends.mysql` |
| `Oracle` | `django.db.backends.oracle` |
| `MariaDB` | `django.db.backends.mysql` |
| `CockroachDB` | `django_cockroachdb` |
| `SQL Server (mssql-django)` | `mssql` |
| `Snowflake` | `django_snowflake` |
| `Django Connection` | `django_connection` |

## ExplorerValueKey

| Name | Stored as |
| --- | --- |
| `Install Unique ID` | `UUID` |
| `Startup metric last send` | `SMLS` |
| `System prompt for SQL Assistant` | `ASP` |
