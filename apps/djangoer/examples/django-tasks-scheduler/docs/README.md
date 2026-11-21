# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [TaskArg](#taskarg) · [TaskKwarg](#taskkwarg) · [Task](#task) · [Queue](#queue) · [Worker](#worker)

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

CREATE TABLE "scheduler_taskarg" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "arg_type" varchar(12) NOT NULL, "val" varchar(2048) NOT NULL, "content_type_id" integer NOT NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "object_id" integer unsigned NOT NULL CHECK ("object_id" >= 0));

CREATE TABLE "scheduler_taskkwarg" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "arg_type" varchar(12) NOT NULL, "val" varchar(2048) NOT NULL, "content_type_id" integer NOT NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "object_id" integer unsigned NOT NULL CHECK ("object_id" >= 0), "key" varchar(255) NOT NULL);

CREATE TABLE "scheduler_task" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "created_at" datetime NOT NULL, "updated_at" datetime NOT NULL, "name" varchar(128) NOT NULL UNIQUE, "task_type" varchar(32) NOT NULL, "callable" varchar(2048) NOT NULL, "enabled" bool NOT NULL, "queue" varchar(255) NOT NULL, "job_name" varchar(128) NULL, "at_front" bool NOT NULL, "timeout" integer NULL, "result_ttl" integer NULL, "failed_runs" integer unsigned NOT NULL CHECK ("failed_runs" >= 0), "successful_runs" integer unsigned NOT NULL CHECK ("successful_runs" >= 0), "last_successful_run" datetime NULL, "last_failed_run" datetime NULL, "interval" integer unsigned NULL CHECK ("interval" >= 0), "interval_unit" varchar(12) NULL, "repeat" integer unsigned NULL CHECK ("repeat" >= 0), "scheduled_time" datetime NULL, "cron_string" varchar(64) NULL);

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

CREATE INDEX "scheduler_taskarg_content_type_id_14c040cc" ON "scheduler_taskarg" ("content_type_id");

CREATE INDEX "scheduler_taskkwarg_content_type_id_d5f93ad6" ON "scheduler_taskkwarg" ("content_type_id");
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
- [TaskArg](#taskarg).`content_type_id` → `id`, on delete cascade
- [TaskKwarg](#taskkwarg).`content_type_id` → `id`, on delete cascade

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

## TaskArg

Table `scheduler_taskarg`, app `scheduler`.

Columns from `BaseTaskArg` (abstract): `arg_type`, `val`, `content_type_id`, `object_id`.

![TaskArg and the models it relates to](tables/TaskArg.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `arg_type` | `TaskArgArgType` | `varchar(12)` |  |  | `'str'` |  |
| `val` | `CharField` | `varchar(2048)` |  |  |  |  |
| `content_type_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `object_id` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |

Indexes:

- index `scheduler_taskarg_content_type_id_14c040cc` on `content_type_id`

Relations:

- `content_object` → any model, by `content_type_id` + `object_id` (generic foreign key)
- `content_type_id` → [ContentType](#contenttype).`id`, on delete cascade
- `object_id` → [Task](#task).`id` where `content_type_id` is Task (generic, `Task.callable_args`)

## TaskKwarg

Table `scheduler_taskkwarg`, app `scheduler`.

Columns from `BaseTaskArg` (abstract): `arg_type`, `val`, `content_type_id`, `object_id`.

![TaskKwarg and the models it relates to](tables/TaskKwarg.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `arg_type` | `TaskKwargArgType` | `varchar(12)` |  |  | `'str'` |  |
| `val` | `CharField` | `varchar(2048)` |  |  |  |  |
| `content_type_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `object_id` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `key` | `CharField` | `varchar(255)` |  |  |  |  |

Indexes:

- index `scheduler_taskkwarg_content_type_id_d5f93ad6` on `content_type_id`

Relations:

- `content_object` → any model, by `content_type_id` + `object_id` (generic foreign key)
- `content_type_id` → [ContentType](#contenttype).`id`, on delete cascade
- `object_id` → [Task](#task).`id` where `content_type_id` is Task (generic, `Task.callable_kwargs`)

## Task

Table `scheduler_task`, app `scheduler`.

![Task and the models it relates to](tables/Task.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `name` | `CharField` | `varchar(128)` |  | unique |  | Name of the job |
| `task_type` | `TaskType` | `varchar(32)` |  |  | `'OnceTaskType'` |  |
| `callable` | `CharField` | `varchar(2048)` |  |  |  |  |
| `enabled` | `BooleanField` | `bool` |  |  | `True` | Should job be scheduled? This field is useful to keep past jobs that should no longer be scheduled |
| `queue` | `TaskQueue` | `varchar(255)` |  |  |  | Queue name |
| `job_name` | `CharField` | `varchar(128)` | yes |  |  | Current job_name on queue |
| `at_front` | `BooleanField` | `bool` |  |  | `False` | When queuing the job, add it in the front of the queue |
| `timeout` | `IntegerField` | `integer` | yes |  |  | Timeout specifies the maximum runtime, in seconds, for the job before it'll be considered 'lost'. Blank uses the default timeout. |
| `result_ttl` | `IntegerField` | `integer` | yes |  |  | The TTL value (in seconds) of the job result.<br/>                -1: Result never expires, you should delete jobs manually. <br/>                 0: Result gets deleted immediately. <br/>                 >0: Result expires after n seconds. |
| `failed_runs` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | Number of times the task has failed |
| `successful_runs` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | Number of times the task has succeeded |
| `last_successful_run` | `DateTimeField` | `datetime` | yes |  |  | Last time the task has succeeded |
| `last_failed_run` | `DateTimeField` | `datetime` | yes |  |  | Last time the task has failed |
| `interval` | `PositiveIntegerField` | `integer unsigned` | yes |  |  | Interval for repeatable task |
| `interval_unit` | `TaskTimeUnits` | `varchar(12)` | yes |  | `'hours'` |  |
| `repeat` | `PositiveIntegerField` | `integer unsigned` | yes |  |  | Number of times to run the job. Leaving this blank means it will run forever. |
| `scheduled_time` | `DateTimeField` | `datetime` | yes |  |  |  |
| `cron_string` | `CharField` | `varchar(64)` | yes |  |  | Define the schedule in a crontab like syntax.             Times are in UTC. Use <a href="https://crontab.guru/">crontab.guru</a> to create a cron string. |

Relations:

- `callable_args` ← [TaskArg](#taskarg), by `content_type_id` + `object_id` (generic relation)
- `callable_kwargs` ← [TaskKwarg](#taskkwarg), by `content_type_id` + `object_id` (generic relation)

## Queue

Table `scheduler_queue`, app `scheduler`.

Not managed by Django (`managed = False`): it neither creates nor changes the table, whose indexes are the ones its own SQL gives it.

Placeholder model with no database table, but with django admin page and contenttype permission

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |

## Worker

Table `scheduler_worker`, app `scheduler`.

Not managed by Django (`managed = False`): it neither creates nor changes the table, whose indexes are the ones its own SQL gives it.

Placeholder model with no database table, but with django admin page and contenttype permission

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## TaskArgArgType

| Name | Stored as |
| --- | --- |
| `string` | `str` |
| `int` | `int` |
| `float` | `float` |
| `boolean` | `bool` |
| `datetime` | `datetime` |
| `JSON` | `json` |
| `callable` | `callable` |

## TaskKwargArgType

| Name | Stored as |
| --- | --- |
| `string` | `str` |
| `int` | `int` |
| `float` | `float` |
| `boolean` | `bool` |
| `datetime` | `datetime` |
| `JSON` | `json` |
| `callable` | `callable` |

## TaskType

| Name | Stored as |
| --- | --- |
| `CRON` | `CronTaskType` |
| `REPEATABLE` | `RepeatableTaskType` |
| `ONCE` | `OnceTaskType` |

## TaskQueue

| Name | Stored as |
| --- | --- |
| `default` | `default` |

## TaskTimeUnits

| Name | Stored as |
| --- | --- |
| `SECONDS` | `seconds` |
| `MINUTES` | `minutes` |
| `HOURS` | `hours` |
| `DAYS` | `days` |
| `WEEKS` | `weeks` |
