# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Report](#report) · [Format](#format) · [DisplayField](#displayfield) · [FilterField](#filterfield)

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

CREATE TABLE "report_builder_report" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(255) NOT NULL, "slug" varchar(50) NOT NULL, "description" text NOT NULL, "root_model_id" integer NOT NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "created" date NOT NULL, "modified" date NOT NULL, "user_created_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "user_modified_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "distinct" bool NOT NULL, "row_limit" integer unsigned NULL CHECK ("row_limit" >= 0), "report_file" varchar(100) NOT NULL, "report_file_creation" datetime NULL, "report_file_task_id" varchar(64) NOT NULL);

CREATE TABLE "report_builder_report_starred" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "report_id" integer NOT NULL REFERENCES "report_builder_report" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "report_builder_format" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(50) NOT NULL, "string" varchar(300) NOT NULL);

CREATE TABLE "report_builder_displayfield" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "report_id" integer NOT NULL REFERENCES "report_builder_report" ("id") DEFERRABLE INITIALLY DEFERRED, "path" varchar(2000) NOT NULL, "path_verbose" varchar(2000) NOT NULL, "field" varchar(2000) NOT NULL, "field_verbose" varchar(2000) NOT NULL, "position" smallint unsigned NULL CHECK ("position" >= 0), "name" varchar(2000) NOT NULL, "sort" integer NULL, "sort_reverse" bool NOT NULL, "width" integer NOT NULL, "aggregate" varchar(5) NOT NULL, "total" bool NOT NULL, "group" bool NOT NULL, "display_format_id" integer NULL REFERENCES "report_builder_format" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "report_builder_filterfield" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "report_id" integer NOT NULL REFERENCES "report_builder_report" ("id") DEFERRABLE INITIALLY DEFERRED, "path" varchar(2000) NOT NULL, "path_verbose" varchar(2000) NOT NULL, "field" varchar(2000) NOT NULL, "field_verbose" varchar(2000) NOT NULL, "position" smallint unsigned NULL CHECK ("position" >= 0), "filter_type" varchar(20) NOT NULL, "filter_delta" bigint NULL, "filter_value" varchar(2000) NOT NULL, "filter_value2" varchar(2000) NOT NULL, "exclude" bool NOT NULL);

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

CREATE INDEX "report_builder_report_slug_559601ba" ON "report_builder_report" ("slug");

CREATE INDEX "report_builder_report_root_model_id_3bae34b3" ON "report_builder_report" ("root_model_id");

CREATE INDEX "report_builder_report_user_created_id_e8133ae8" ON "report_builder_report" ("user_created_id");

CREATE INDEX "report_builder_report_user_modified_id_04e38ffb" ON "report_builder_report" ("user_modified_id");

CREATE UNIQUE INDEX "report_builder_report_starred_report_id_user_id_8e38d9bd_uniq" ON "report_builder_report_starred" ("report_id", "user_id");

CREATE INDEX "report_builder_report_starred_report_id_0a54a683" ON "report_builder_report_starred" ("report_id");

CREATE INDEX "report_builder_report_starred_user_id_70388cb5" ON "report_builder_report_starred" ("user_id");

CREATE INDEX "report_builder_displayfield_report_id_642d3d44" ON "report_builder_displayfield" ("report_id");

CREATE INDEX "report_builder_displayfield_display_format_id_c649f0ea" ON "report_builder_displayfield" ("display_format_id");

CREATE INDEX "report_builder_filterfield_report_id_a5498074" ON "report_builder_filterfield" ("report_id");
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
- [Report](#report).`user_created_id` → `id`, on delete set null
- [Report](#report).`user_modified_id` → `id`, on delete set null
- `report_starred_set` — many to many with [Report](#report)

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
- [Report](#report).`root_model_id` → `id`, on delete cascade

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

## Report

Table `report_builder_report`, app `report_builder`.

A saved report with queryset and descriptive fields

![Report and the models it relates to](tables/Report.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `slug` | `SlugField` | `varchar(50)` |  |  |  |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `root_model_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created` | `DateField` | `date` |  |  | `now()` |  |
| `modified` | `DateField` | `date` |  |  | `now()` |  |
| `user_created_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `user_modified_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `distinct` | `BooleanField` | `bool` |  |  | `False` |  |
| `row_limit` | `PositiveIntegerField` | `integer unsigned` | yes |  |  | Maximum number of rows to include in exports. Leave blank for no limit. |
| `report_file` | `FileField` | `varchar(100)` |  |  |  |  |
| `report_file_creation` | `DateTimeField` | `datetime` | yes |  |  |  |
| `report_file_task_id` | `CharField` | `varchar(64)` |  |  | `''` |  |

Indexes:

- index `report_builder_report_slug_559601ba` on `slug`
- index `report_builder_report_root_model_id_3bae34b3` on `root_model_id`
- index `report_builder_report_user_created_id_e8133ae8` on `user_created_id`
- index `report_builder_report_user_modified_id_04e38ffb` on `user_modified_id`

Relations:

- `root_model_id` → [ContentType](#contenttype).`id`, on delete cascade
- `user_created_id` → [User](#user).`id`, on delete set null
- `user_modified_id` → [User](#user).`id`, on delete set null
- `starred` — many to many with [User](#user)
- [DisplayField](#displayfield).`report_id` → `id`, on delete cascade
- [FilterField](#filterfield).`report_id` → `id`, on delete cascade

## Format

Table `report_builder_format`, app `report_builder`.

A specifies a Python string format for e.g. `DisplayField`s.

![Format and the models it relates to](tables/Format.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(50)` |  |  | `''` |  |
| `string` | `CharField` | `varchar(300)` |  |  | `''` | Python string format. Ex ${} would place a $ in front of the result. |

Relations:

- [DisplayField](#displayfield).`display_format_id` → `id`, on delete set null

## DisplayField

Table `report_builder_displayfield`, app `report_builder`.

Columns from `AbstractField` (abstract): `report_id`, `path`, `path_verbose`, `field`, `field_verbose`, `position`.

A display field to show in a report. Always belongs to a Report

![DisplayField and the models it relates to](tables/DisplayField.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `report_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `path` | `CharField` | `varchar(2000)` |  |  |  |  |
| `path_verbose` | `CharField` | `varchar(2000)` |  |  |  |  |
| `field` | `CharField` | `varchar(2000)` |  |  |  |  |
| `field_verbose` | `CharField` | `varchar(2000)` |  |  |  |  |
| `position` | `PositiveSmallIntegerField` | `smallint unsigned` | yes |  |  |  |
| `name` | `CharField` | `varchar(2000)` |  |  |  |  |
| `sort` | `IntegerField` | `integer` | yes |  |  |  |
| `sort_reverse` | `BooleanField` | `bool` |  |  | `False` |  |
| `width` | `IntegerField` | `integer` |  |  | `15` |  |
| `aggregate` | `DisplayFieldAggregate` | `varchar(5)` |  |  |  |  |
| `total` | `BooleanField` | `bool` |  |  | `False` |  |
| `group` | `BooleanField` | `bool` |  |  | `False` |  |
| `display_format_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `report_builder_displayfield_report_id_642d3d44` on `report_id`
- index `report_builder_displayfield_display_format_id_c649f0ea` on `display_format_id`

Relations:

- `report_id` → [Report](#report).`id`, on delete cascade
- `display_format_id` → [Format](#format).`id`, on delete set null

## FilterField

Table `report_builder_filterfield`, app `report_builder`.

Columns from `AbstractField` (abstract): `report_id`, `path`, `path_verbose`, `field`, `field_verbose`, `position`.

A filter model used to filter DisplayFields using Django ORM
filter arguments or custom filter values.

![FilterField and the models it relates to](tables/FilterField.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `report_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `path` | `CharField` | `varchar(2000)` |  |  |  |  |
| `path_verbose` | `CharField` | `varchar(2000)` |  |  |  |  |
| `field` | `CharField` | `varchar(2000)` |  |  |  |  |
| `field_verbose` | `CharField` | `varchar(2000)` |  |  |  |  |
| `position` | `PositiveSmallIntegerField` | `smallint unsigned` | yes |  |  |  |
| `filter_type` | `FilterFieldFilterType` | `varchar(20)` |  |  | `'icontains'` |  |
| `filter_delta` | `BigIntegerField` | `bigint` | yes |  |  |  |
| `filter_value` | `CharField` | `varchar(2000)` |  |  |  |  |
| `filter_value2` | `CharField` | `varchar(2000)` |  |  |  |  |
| `exclude` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `report_builder_filterfield_report_id_a5498074` on `report_id`

Relations:

- `report_id` → [Report](#report).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## DisplayFieldAggregate

| Name | Stored as |
| --- | --- |
| `Sum` | `Sum` |
| `Count` | `Count` |
| `Avg` | `Avg` |
| `Max` | `Max` |
| `Min` | `Min` |

## FilterFieldFilterType

| Name | Stored as |
| --- | --- |
| `Equals` | `exact` |
| `Equals (case-insensitive)` | `iexact` |
| `Contains` | `contains` |
| `Contains (case-insensitive)` | `icontains` |
| `in (comma seperated 1,2,3)` | `in` |
| `Greater than` | `gt` |
| `Greater than equals` | `gte` |
| `Less than` | `lt` |
| `Less than equals` | `lte` |
| `Starts with` | `startswith` |
| `Starts with (case-insensitive)` | `istartswith` |
| `Ends with` | `endswith` |
| `Ends with  (case-insensitive)` | `iendswith` |
| `range` | `range` |
| `relative_range` | `relative_range` |
| `Week day` | `week_day` |
| `Is null` | `isnull` |
| `Regular Expression` | `regex` |
| `Reg. Exp. (case-insensitive)` | `iregex` |
| `Max (annotation-filter)` | `max` |
| `Min (annotation-filter)` | `min` |
