# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Request](#request) · [Response](#response) · [SQLQuery](#sqlquery) · [Profile](#profile)

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

CREATE TABLE "silk_request" ("id" varchar(36) NOT NULL PRIMARY KEY, "path" varchar(190) NOT NULL, "query_params" text NOT NULL, "raw_body" text NOT NULL, "body" text NOT NULL, "method" varchar(10) NOT NULL, "start_time" datetime NOT NULL, "view_name" varchar(190) NULL, "end_time" datetime NULL, "time_taken" real NULL, "encoded_headers" text NOT NULL, "meta_time" real NULL, "meta_num_queries" integer NULL, "meta_time_spent_queries" real NULL, "pyprofile" text NOT NULL, "prof_file" varchar(300) NOT NULL, "num_sql_queries" integer NOT NULL);

CREATE TABLE "silk_response" ("id" varchar(36) NOT NULL PRIMARY KEY, "request_id" varchar(36) NOT NULL UNIQUE REFERENCES "silk_request" ("id") DEFERRABLE INITIALLY DEFERRED, "status_code" integer NOT NULL, "raw_body" text NOT NULL, "body" text NOT NULL, "encoded_headers" text NOT NULL);

CREATE TABLE "silk_sqlquery" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "query" text NOT NULL, "start_time" datetime NULL, "end_time" datetime NULL, "time_taken" real NULL, "identifier" integer NOT NULL, "request_id" varchar(36) NULL REFERENCES "silk_request" ("id") DEFERRABLE INITIALLY DEFERRED, "traceback" text NOT NULL, "analysis" text NULL);

CREATE TABLE "silk_profile" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(300) NOT NULL, "start_time" datetime NOT NULL, "end_time" datetime NULL, "request_id" varchar(36) NULL REFERENCES "silk_request" ("id") DEFERRABLE INITIALLY DEFERRED, "time_taken" real NULL, "file_path" varchar(300) NOT NULL, "line_num" integer NULL, "end_line_num" integer NULL, "func_name" varchar(300) NOT NULL, "exception_raised" bool NOT NULL, "dynamic" bool NOT NULL);

CREATE TABLE "silk_profile_queries" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "profile_id" integer NOT NULL REFERENCES "silk_profile" ("id") DEFERRABLE INITIALLY DEFERRED, "sqlquery_id" integer NOT NULL REFERENCES "silk_sqlquery" ("id") DEFERRABLE INITIALLY DEFERRED);

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

CREATE INDEX "silk_request_path_9f3d798e" ON "silk_request" ("path");

CREATE INDEX "silk_request_start_time_1300bc58" ON "silk_request" ("start_time");

CREATE INDEX "silk_request_view_name_68559f7b" ON "silk_request" ("view_name");

CREATE INDEX "silk_sqlquery_request_id_6f8f0527" ON "silk_sqlquery" ("request_id");

CREATE INDEX "silk_profile_request_id_7b81bd69" ON "silk_profile" ("request_id");

CREATE UNIQUE INDEX "silk_profile_queries_profile_id_sqlquery_id_b2403d9b_uniq" ON "silk_profile_queries" ("profile_id", "sqlquery_id");

CREATE INDEX "silk_profile_queries_profile_id_a3d76db8" ON "silk_profile_queries" ("profile_id");

CREATE INDEX "silk_profile_queries_sqlquery_id_155df455" ON "silk_profile_queries" ("sqlquery_id");
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

## Request

Table `silk_request`, app `silk`.

![Request and the models it relates to](tables/Request.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `CharField` | `varchar(36)` |  | PK | `uuid4()` |  |
| `path` | `CharField` | `varchar(190)` |  |  |  |  |
| `query_params` | `TextField` | `text` |  |  | `''` |  |
| `raw_body` | `TextField` | `text` |  |  | `''` |  |
| `body` | `TextField` | `text` |  |  | `''` |  |
| `method` | `CharField` | `varchar(10)` |  |  |  |  |
| `start_time` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `view_name` | `CharField` | `varchar(190)` | yes |  | `''` |  |
| `end_time` | `DateTimeField` | `datetime` | yes |  |  |  |
| `time_taken` | `FloatField` | `real` | yes |  |  |  |
| `encoded_headers` | `TextField` | `text` |  |  | `''` |  |
| `meta_time` | `FloatField` | `real` | yes |  |  |  |
| `meta_num_queries` | `IntegerField` | `integer` | yes |  |  |  |
| `meta_time_spent_queries` | `FloatField` | `real` | yes |  |  |  |
| `pyprofile` | `TextField` | `text` |  |  | `''` |  |
| `prof_file` | `FileField` | `varchar(300)` |  |  |  |  |
| `num_sql_queries` | `IntegerField` | `integer` |  |  | `0` |  |

Indexes:

- index `silk_request_path_9f3d798e` on `path`
- index `silk_request_start_time_1300bc58` on `start_time`
- index `silk_request_view_name_68559f7b` on `view_name`

Relations:

- [Response](#response).`request_id` → `id`, on delete cascade
- [SQLQuery](#sqlquery).`request_id` → `id`, on delete cascade
- [Profile](#profile).`request_id` → `id`, on delete cascade

## Response

Table `silk_response`, app `silk`.

![Response and the models it relates to](tables/Response.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `CharField` | `varchar(36)` |  | PK | `uuid4()` |  |
| `request_id` | `OneToOneField` | `varchar(36)` |  | FK, unique |  |  |
| `status_code` | `IntegerField` | `integer` |  |  |  |  |
| `raw_body` | `TextField` | `text` |  |  | `''` |  |
| `body` | `TextField` | `text` |  |  | `''` |  |
| `encoded_headers` | `TextField` | `text` |  |  | `''` |  |

Relations:

- `request_id` → [Request](#request).`id`, on delete cascade

## SQLQuery

Table `silk_sqlquery`, app `silk`.

![SQLQuery and the models it relates to](tables/SQLQuery.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `query` | `TextField` | `text` |  |  |  |  |
| `start_time` | `DateTimeField` | `datetime` | yes |  | `now()` |  |
| `end_time` | `DateTimeField` | `datetime` | yes |  |  |  |
| `time_taken` | `FloatField` | `real` | yes |  |  |  |
| `identifier` | `IntegerField` | `integer` |  |  | `-1` |  |
| `request_id` | `ForeignKey` | `varchar(36)` | yes | FK |  |  |
| `traceback` | `TextField` | `text` |  |  |  |  |
| `analysis` | `TextField` | `text` | yes |  |  |  |

Indexes:

- index `silk_sqlquery_request_id_6f8f0527` on `request_id`

Relations:

- `request_id` → [Request](#request).`id`, on delete cascade
- `profiles` — many to many with [Profile](#profile)

## Profile

Table `silk_profile`, app `silk`.

Columns from `BaseProfile` (abstract): `name`, `start_time`, `end_time`, `request_id`, `time_taken`.

![Profile and the models it relates to](tables/Profile.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(300)` |  |  | `''` |  |
| `start_time` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `end_time` | `DateTimeField` | `datetime` | yes |  |  |  |
| `request_id` | `ForeignKey` | `varchar(36)` | yes | FK |  |  |
| `time_taken` | `FloatField` | `real` | yes |  |  |  |
| `file_path` | `CharField` | `varchar(300)` |  |  | `''` |  |
| `line_num` | `IntegerField` | `integer` | yes |  |  |  |
| `end_line_num` | `IntegerField` | `integer` | yes |  |  |  |
| `func_name` | `CharField` | `varchar(300)` |  |  | `''` |  |
| `exception_raised` | `BooleanField` | `bool` |  |  | `False` |  |
| `dynamic` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `silk_profile_request_id_7b81bd69` on `request_id`

Relations:

- `request_id` → [Request](#request).`id`, on delete cascade
- `queries` — many to many with [SQLQuery](#sqlquery)

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |
