# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Task](#task) · [CompletedTask](#completedtask)

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

CREATE TABLE "background_task" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "task_name" varchar(190) NOT NULL, "task_params" text NOT NULL, "task_hash" varchar(40) NOT NULL, "verbose_name" varchar(255) NULL, "priority" integer NOT NULL, "run_at" datetime NOT NULL, "repeat" bigint NOT NULL, "repeat_until" datetime NULL, "queue" varchar(190) NULL, "attempts" integer NOT NULL, "failed_at" datetime NULL, "last_error" text NOT NULL, "locked_by" varchar(64) NULL, "locked_at" datetime NULL, "creator_content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "creator_object_id" integer unsigned NULL CHECK ("creator_object_id" >= 0));

CREATE TABLE "background_task_completedtask" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "task_name" varchar(190) NOT NULL, "task_params" text NOT NULL, "task_hash" varchar(40) NOT NULL, "verbose_name" varchar(255) NULL, "priority" integer NOT NULL, "run_at" datetime NOT NULL, "repeat" bigint NOT NULL, "repeat_until" datetime NULL, "queue" varchar(190) NULL, "attempts" integer NOT NULL, "failed_at" datetime NULL, "last_error" text NOT NULL, "locked_by" varchar(64) NULL, "locked_at" datetime NULL, "creator_content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "creator_object_id" integer unsigned NULL CHECK ("creator_object_id" >= 0));

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

CREATE INDEX "background_task_task_name_4562d56a" ON "background_task" ("task_name");

CREATE INDEX "background_task_task_hash_d8f233bd" ON "background_task" ("task_hash");

CREATE INDEX "background_task_priority_88bdbce9" ON "background_task" ("priority");

CREATE INDEX "background_task_run_at_7baca3aa" ON "background_task" ("run_at");

CREATE INDEX "background_task_queue_1d5f3a40" ON "background_task" ("queue");

CREATE INDEX "background_task_attempts_a9ade23d" ON "background_task" ("attempts");

CREATE INDEX "background_task_failed_at_b81bba14" ON "background_task" ("failed_at");

CREATE INDEX "background_task_locked_by_db7779e3" ON "background_task" ("locked_by");

CREATE INDEX "background_task_locked_at_0fb0f225" ON "background_task" ("locked_at");

CREATE INDEX "background_task_creator_content_type_id_61cc9af3" ON "background_task" ("creator_content_type_id");

CREATE INDEX "background_task_completedtask_task_name_388dabc2" ON "background_task_completedtask" ("task_name");

CREATE INDEX "background_task_completedtask_task_hash_91187576" ON "background_task_completedtask" ("task_hash");

CREATE INDEX "background_task_completedtask_priority_9080692e" ON "background_task_completedtask" ("priority");

CREATE INDEX "background_task_completedtask_run_at_77c80f34" ON "background_task_completedtask" ("run_at");

CREATE INDEX "background_task_completedtask_queue_61fb0415" ON "background_task_completedtask" ("queue");

CREATE INDEX "background_task_completedtask_attempts_772a6783" ON "background_task_completedtask" ("attempts");

CREATE INDEX "background_task_completedtask_failed_at_3de56618" ON "background_task_completedtask" ("failed_at");

CREATE INDEX "background_task_completedtask_locked_by_edc8a213" ON "background_task_completedtask" ("locked_by");

CREATE INDEX "background_task_completedtask_locked_at_29c62708" ON "background_task_completedtask" ("locked_at");

CREATE INDEX "background_task_completedtask_creator_content_type_id_21d6a741" ON "background_task_completedtask" ("creator_content_type_id");
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
- [Task](#task).`creator_content_type_id` → `id`, on delete cascade
- [CompletedTask](#completedtask).`creator_content_type_id` → `id`, on delete cascade

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

## Task

Table `background_task`, app `background_task`.

![Task and the models it relates to](tables/Task.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `task_name` | `CharField` | `varchar(190)` |  |  |  |  |
| `task_params` | `TextField` | `text` |  |  |  |  |
| `task_hash` | `CharField` | `varchar(40)` |  |  |  |  |
| `verbose_name` | `CharField` | `varchar(255)` | yes |  |  |  |
| `priority` | `IntegerField` | `integer` |  |  | `0` |  |
| `run_at` | `DateTimeField` | `datetime` |  |  |  |  |
| `repeat` | `TaskRepeat` | `bigint` |  |  | `0` |  |
| `repeat_until` | `DateTimeField` | `datetime` | yes |  |  |  |
| `queue` | `CharField` | `varchar(190)` | yes |  |  |  |
| `attempts` | `IntegerField` | `integer` |  |  | `0` |  |
| `failed_at` | `DateTimeField` | `datetime` | yes |  |  |  |
| `last_error` | `TextField` | `text` |  |  |  |  |
| `locked_by` | `CharField` | `varchar(64)` | yes |  |  |  |
| `locked_at` | `DateTimeField` | `datetime` | yes |  |  |  |
| `creator_content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `creator_object_id` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |

Indexes:

- index `background_task_task_name_4562d56a` on `task_name`
- index `background_task_task_hash_d8f233bd` on `task_hash`
- index `background_task_priority_88bdbce9` on `priority`
- index `background_task_run_at_7baca3aa` on `run_at`
- index `background_task_queue_1d5f3a40` on `queue`
- index `background_task_attempts_a9ade23d` on `attempts`
- index `background_task_failed_at_b81bba14` on `failed_at`
- index `background_task_locked_by_db7779e3` on `locked_by`
- index `background_task_locked_at_0fb0f225` on `locked_at`
- index `background_task_creator_content_type_id_61cc9af3` on `creator_content_type_id`

Relations:

- `creator` → any model, by `creator_content_type_id` + `creator_object_id` (generic foreign key)
- `creator_content_type_id` → [ContentType](#contenttype).`id`, on delete cascade

## CompletedTask

Table `background_task_completedtask`, app `background_task`.

![CompletedTask and the models it relates to](tables/CompletedTask.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `task_name` | `CharField` | `varchar(190)` |  |  |  |  |
| `task_params` | `TextField` | `text` |  |  |  |  |
| `task_hash` | `CharField` | `varchar(40)` |  |  |  |  |
| `verbose_name` | `CharField` | `varchar(255)` | yes |  |  |  |
| `priority` | `IntegerField` | `integer` |  |  | `0` |  |
| `run_at` | `DateTimeField` | `datetime` |  |  |  |  |
| `repeat` | `CompletedTaskRepeat` | `bigint` |  |  | `0` |  |
| `repeat_until` | `DateTimeField` | `datetime` | yes |  |  |  |
| `queue` | `CharField` | `varchar(190)` | yes |  |  |  |
| `attempts` | `IntegerField` | `integer` |  |  | `0` |  |
| `failed_at` | `DateTimeField` | `datetime` | yes |  |  |  |
| `last_error` | `TextField` | `text` |  |  |  |  |
| `locked_by` | `CharField` | `varchar(64)` | yes |  |  |  |
| `locked_at` | `DateTimeField` | `datetime` | yes |  |  |  |
| `creator_content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `creator_object_id` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |

Indexes:

- index `background_task_completedtask_task_name_388dabc2` on `task_name`
- index `background_task_completedtask_task_hash_91187576` on `task_hash`
- index `background_task_completedtask_priority_9080692e` on `priority`
- index `background_task_completedtask_run_at_77c80f34` on `run_at`
- index `background_task_completedtask_queue_61fb0415` on `queue`
- index `background_task_completedtask_attempts_772a6783` on `attempts`
- index `background_task_completedtask_failed_at_3de56618` on `failed_at`
- index `background_task_completedtask_locked_by_edc8a213` on `locked_by`
- index `background_task_completedtask_locked_at_29c62708` on `locked_at`
- index `background_task_completedtask_creator_content_type_id_21d6a741` on `creator_content_type_id`

Relations:

- `creator` → any model, by `creator_content_type_id` + `creator_object_id` (generic foreign key)
- `creator_content_type_id` → [ContentType](#contenttype).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## TaskRepeat

| Name | Stored as |
| --- | --- |
| `hourly` | `3600` |
| `daily` | `86400` |
| `weekly` | `604800` |
| `every 2 weeks` | `1209600` |
| `every 4 weeks` | `2419200` |
| `never` | `0` |

## CompletedTaskRepeat

| Name | Stored as |
| --- | --- |
| `hourly` | `3600` |
| `daily` | `86400` |
| `weekly` | `604800` |
| `every 2 weeks` | `1209600` |
| `every 4 weeks` | `2419200` |
| `never` | `0` |
