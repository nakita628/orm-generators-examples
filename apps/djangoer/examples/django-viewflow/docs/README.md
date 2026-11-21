# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Process](#process) · [Task](#task)

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

CREATE TABLE "viewflow_process" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "flow_class" varchar(250) NOT NULL, "status" varchar(50) NOT NULL, "created" datetime NOT NULL, "finished" datetime NULL, "data" text NOT NULL CHECK ((JSON_VALID("data") OR "data" IS NULL)), "parent_task_id" bigint NULL REFERENCES "viewflow_task" ("id") DEFERRABLE INITIALLY DEFERRED, "seed_content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "seed_object_id" integer unsigned NULL CHECK ("seed_object_id" >= 0), "artifact_content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "artifact_object_id" integer unsigned NULL CHECK ("artifact_object_id" >= 0));

CREATE TABLE "viewflow_task" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "flow_task" varchar(255) NOT NULL, "flow_task_type" varchar(50) NOT NULL, "status" varchar(50) NOT NULL, "created" datetime NOT NULL, "assigned" datetime NULL, "started" datetime NULL, "finished" datetime NULL, "scheduled" datetime NULL, "token" varchar(150) NOT NULL, "external_task_id" varchar(50) NULL, "owner_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "owner_permission" varchar(255) NULL, "owner_permission_content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "owner_permission_obj_pk" varchar(255) NULL, "process_id" bigint NOT NULL REFERENCES "viewflow_process" ("id") DEFERRABLE INITIALLY DEFERRED, "data" text NOT NULL CHECK ((JSON_VALID("data") OR "data" IS NULL)), "seed_content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "seed_object_id" integer unsigned NULL CHECK ("seed_object_id" >= 0), "artifact_content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "artifact_object_id" integer unsigned NULL CHECK ("artifact_object_id" >= 0));

CREATE TABLE "viewflow_task_previous" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "from_task_id" bigint NOT NULL REFERENCES "viewflow_task" ("id") DEFERRABLE INITIALLY DEFERRED, "to_task_id" bigint NOT NULL REFERENCES "viewflow_task" ("id") DEFERRABLE INITIALLY DEFERRED);

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

CREATE INDEX "viewflow_process_parent_task_id_d3453359" ON "viewflow_process" ("parent_task_id");

CREATE INDEX "viewflow_process_seed_content_type_id_d648268f" ON "viewflow_process" ("seed_content_type_id");

CREATE INDEX "viewflow_process_artifact_content_type_id_8e40f778" ON "viewflow_process" ("artifact_content_type_id");

CREATE INDEX "viewflow_pr_artifac_64c8a3_idx" ON "viewflow_process" ("artifact_content_type_id", "artifact_object_id");

CREATE INDEX "viewflow_task_status_ffb4516b" ON "viewflow_task" ("status");

CREATE INDEX "viewflow_task_created_7518aa4b" ON "viewflow_task" ("created");

CREATE INDEX "viewflow_task_scheduled_bbd49899" ON "viewflow_task" ("scheduled");

CREATE INDEX "viewflow_task_external_task_id_2f93bc2f" ON "viewflow_task" ("external_task_id");

CREATE INDEX "viewflow_task_owner_id_01b087bd" ON "viewflow_task" ("owner_id");

CREATE INDEX "viewflow_task_owner_permission_content_type_id_5fc72a99" ON "viewflow_task" ("owner_permission_content_type_id");

CREATE INDEX "viewflow_task_process_id_619bc849" ON "viewflow_task" ("process_id");

CREATE INDEX "viewflow_task_seed_content_type_id_bed3f17b" ON "viewflow_task" ("seed_content_type_id");

CREATE INDEX "viewflow_task_artifact_content_type_id_50385e8d" ON "viewflow_task" ("artifact_content_type_id");

CREATE UNIQUE INDEX "viewflow_task_previous_from_task_id_to_task_id_a979d787_uniq" ON "viewflow_task_previous" ("from_task_id", "to_task_id");

CREATE INDEX "viewflow_task_previous_from_task_id_da5536f6" ON "viewflow_task_previous" ("from_task_id");

CREATE INDEX "viewflow_task_previous_to_task_id_7d3a2081" ON "viewflow_task_previous" ("to_task_id");
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
- [Task](#task).`owner_id` → `id`, on delete cascade

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
- [Process](#process).`seed_content_type_id` → `id`, on delete cascade
- [Process](#process).`artifact_content_type_id` → `id`, on delete cascade
- [Task](#task).`owner_permission_content_type_id` → `id`, on delete cascade
- [Task](#task).`seed_content_type_id` → `id`, on delete cascade
- [Task](#task).`artifact_content_type_id` → `id`, on delete cascade

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

## Process

Table `viewflow_process`, app `viewflow`.

Columns from `AbstractProcess` (abstract): `flow_class`, `status`, `created`, `finished`.

Default viewflow Process model.

![Process and the models it relates to](tables/Process.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `flow_class` | `FlowReferenceField` | `varchar(250)` |  |  |  |  |
| `status` | `ProcessStatus` | `varchar(50)` |  |  | `'NEW'` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `finished` | `DateTimeField` | `datetime` | yes |  |  |  |
| `data` | `JSONField` | `text` |  |  | `dict()` |  |
| `parent_task_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `seed_content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `seed_object_id` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |
| `artifact_content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `artifact_object_id` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |

Indexes:

- index `viewflow_process_parent_task_id_d3453359` on `parent_task_id`
- index `viewflow_process_seed_content_type_id_d648268f` on `seed_content_type_id`
- index `viewflow_process_artifact_content_type_id_8e40f778` on `artifact_content_type_id`
- index `viewflow_pr_artifac_64c8a3_idx` on `artifact_content_type_id`, `artifact_object_id`

Relations:

- `seed` → any model, by `seed_content_type_id` + `seed_object_id` (generic foreign key)
- `artifact` → any model, by `artifact_content_type_id` + `artifact_object_id` (generic foreign key)
- `parent_task_id` → [Task](#task).`id`, on delete cascade
- `seed_content_type_id` → [ContentType](#contenttype).`id`, on delete cascade
- `artifact_content_type_id` → [ContentType](#contenttype).`id`, on delete cascade
- [Task](#task).`process_id` → `id`, on delete cascade

## Task

Table `viewflow_task`, app `viewflow`.

Columns from `AbstractTask` (abstract): `flow_task`, `flow_task_type`, `status`, `created`, `assigned`, `started`, `finished`, `scheduled`, `token`, `external_task_id`, `owner_id`, `owner_permission`, `owner_permission_content_type_id`, `owner_permission_obj_pk`.

Default viewflow Task model.

![Task and the models it relates to](tables/Task.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `flow_task` | `TaskReferenceField` | `varchar(255)` |  |  |  |  |
| `flow_task_type` | `CharField` | `varchar(50)` |  |  |  |  |
| `status` | `TaskStatus` | `varchar(50)` |  |  | `'NEW'` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `assigned` | `DateTimeField` | `datetime` | yes |  |  |  |
| `started` | `DateTimeField` | `datetime` | yes |  |  |  |
| `finished` | `DateTimeField` | `datetime` | yes |  |  |  |
| `scheduled` | `DateTimeField` | `datetime` | yes |  |  |  |
| `token` | `TokenField` | `varchar(150)` |  |  | `start` |  |
| `external_task_id` | `CharField` | `varchar(50)` | yes |  |  |  |
| `owner_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `owner_permission` | `CharField` | `varchar(255)` | yes |  |  |  |
| `owner_permission_content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `owner_permission_obj_pk` | `CharField` | `varchar(255)` | yes |  |  |  |
| `process_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `data` | `JSONField` | `text` |  |  | `dict()` |  |
| `seed_content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `seed_object_id` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |
| `artifact_content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `artifact_object_id` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |

Indexes:

- index `viewflow_task_status_ffb4516b` on `status`
- index `viewflow_task_created_7518aa4b` on `created`
- index `viewflow_task_scheduled_bbd49899` on `scheduled`
- index `viewflow_task_external_task_id_2f93bc2f` on `external_task_id`
- index `viewflow_task_owner_id_01b087bd` on `owner_id`
- index `viewflow_task_owner_permission_content_type_id_5fc72a99` on `owner_permission_content_type_id`
- index `viewflow_task_process_id_619bc849` on `process_id`
- index `viewflow_task_seed_content_type_id_bed3f17b` on `seed_content_type_id`
- index `viewflow_task_artifact_content_type_id_50385e8d` on `artifact_content_type_id`

Relations:

- `seed` → any model, by `seed_content_type_id` + `seed_object_id` (generic foreign key)
- `artifact` → any model, by `artifact_content_type_id` + `artifact_object_id` (generic foreign key)
- `owner_permission_obj` → any model, by `owner_permission_content_type_id` + `owner_permission_obj_pk` (generic foreign key)
- [Process](#process).`parent_task_id` → `id`, on delete cascade
- `owner_id` → [User](#user).`id`, on delete cascade
- `owner_permission_content_type_id` → [ContentType](#contenttype).`id`, on delete cascade
- `process_id` → [Process](#process).`id`, on delete cascade
- `seed_content_type_id` → [ContentType](#contenttype).`id`, on delete cascade
- `artifact_content_type_id` → [ContentType](#contenttype).`id`, on delete cascade
- `previous` — many to many with [Task](#task)

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## ProcessStatus

| Name | Stored as |
| --- | --- |
| `Canceled` | `CANCELED` |
| `Done` | `DONE` |
| `New` | `NEW` |

## TaskStatus

| Name | Stored as |
| --- | --- |
| `Assigned` | `ASSIGNED` |
| `Canceled` | `CANCELED` |
| `Done` | `DONE` |
| `Error` | `ERROR` |
| `New` | `NEW` |
| `Revived` | `REVIVED` |
| `Scheduled` | `SCHEDULED` |
| `Started` | `STARTED` |
