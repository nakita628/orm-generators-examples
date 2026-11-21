# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Site](#site) · [JobDependency](#jobdependency) · [Job](#job) · [Log](#log)

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

CREATE TABLE "django_site" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "domain" varchar(100) NOT NULL UNIQUE, "name" varchar(50) NOT NULL);

CREATE TABLE "chroniker_jobdependency" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "dependent_id" integer NOT NULL REFERENCES "chroniker_job" ("id") DEFERRABLE INITIALLY DEFERRED, "dependee_id" integer NOT NULL REFERENCES "chroniker_job" ("id") DEFERRABLE INITIALLY DEFERRED, "wait_for_completion" bool NOT NULL, "wait_for_success" bool NOT NULL, "wait_for_next_run" bool NOT NULL);

CREATE TABLE "chroniker_job" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(200) NOT NULL, "frequency" varchar(10) NOT NULL, "params" text NULL, "command" varchar(200) NOT NULL, "args" varchar(200) NOT NULL, "raw_command" varchar(1000) NULL, "enabled" bool NOT NULL, "next_run" datetime NULL, "jitter_seconds" integer unsigned NOT NULL CHECK ("jitter_seconds" >= 0), "last_run_start_timestamp" datetime NULL, "last_run" datetime NULL, "last_heartbeat" datetime NULL, "is_running" bool NOT NULL, "last_run_successful" bool NULL, "email_errors_to_subscribers" bool NOT NULL, "email_success_to_subscribers" bool NOT NULL, "lock_file" varchar(255) NOT NULL, "force_run" bool NOT NULL, "force_stop" bool NOT NULL, "timeout_seconds" integer unsigned NOT NULL CHECK ("timeout_seconds" >= 0), "hostname" varchar(700) NULL, "current_hostname" varchar(700) NULL, "current_pid" varchar(50) NULL, "total_parts_complete" integer unsigned NOT NULL CHECK ("total_parts_complete" >= 0), "total_parts" integer unsigned NOT NULL CHECK ("total_parts" >= 0), "is_monitor" bool NOT NULL, "monitor_url" varchar(255) NULL, "monitor_error_template" text NULL, "monitor_description" text NULL, "monitor_records" integer NULL, "maximum_log_entries" integer unsigned NOT NULL CHECK ("maximum_log_entries" >= 0), "log_stdout" bool NOT NULL, "log_stderr" bool NOT NULL);

CREATE TABLE "chroniker_job_subscribers" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "job_id" integer NOT NULL REFERENCES "chroniker_job" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "chroniker_log" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "job_id" integer NOT NULL REFERENCES "chroniker_job" ("id") DEFERRABLE INITIALLY DEFERRED, "run_start_datetime" datetime NOT NULL, "run_end_datetime" datetime NULL, "duration_seconds" integer unsigned NULL CHECK ("duration_seconds" >= 0), "stdout" text NOT NULL, "stderr" text NOT NULL, "hostname" varchar(700) NULL, "success" bool NOT NULL, "on_time" bool NOT NULL);

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

CREATE UNIQUE INDEX "chroniker_jobdependency_dependent_id_dependee_id_a59a3257_uniq" ON "chroniker_jobdependency" ("dependent_id", "dependee_id");

CREATE INDEX "chroniker_jobdependency_dependent_id_48fb0cbb" ON "chroniker_jobdependency" ("dependent_id");

CREATE INDEX "chroniker_jobdependency_dependee_id_5107f29f" ON "chroniker_jobdependency" ("dependee_id");

CREATE INDEX "chroniker_job_current_pid_9996fbd0" ON "chroniker_job" ("current_pid");

CREATE UNIQUE INDEX "chroniker_job_subscribers_job_id_user_id_d297a4fc_uniq" ON "chroniker_job_subscribers" ("job_id", "user_id");

CREATE INDEX "chroniker_job_subscribers_job_id_cfb16efe" ON "chroniker_job_subscribers" ("job_id");

CREATE INDEX "chroniker_job_subscribers_user_id_51225b88" ON "chroniker_job_subscribers" ("user_id");

CREATE INDEX "chroniker_log_job_id_b73a0dc9" ON "chroniker_log" ("job_id");

CREATE INDEX "chroniker_log_run_start_datetime_ed601856" ON "chroniker_log" ("run_start_datetime");

CREATE INDEX "chroniker_log_run_end_datetime_ee9ac252" ON "chroniker_log" ("run_end_datetime");

CREATE INDEX "chroniker_log_duration_seconds_34305844" ON "chroniker_log" ("duration_seconds");

CREATE INDEX "chroniker_log_success_9d4f1251" ON "chroniker_log" ("success");

CREATE INDEX "chroniker_log_on_time_75c95897" ON "chroniker_log" ("on_time");
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
- `subscribed_jobs` — many to many with [Job](#job)

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

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `domain` | `CharField` | `varchar(100)` |  | unique |  |  |
| `name` | `CharField` | `varchar(50)` |  |  |  |  |

## JobDependency

Table `chroniker_jobdependency`, app `chroniker`.

Represents a scheduling dependency between two jobs.

![JobDependency and the models it relates to](tables/JobDependency.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `dependent_id` | `ForeignKey` | `integer` |  | FK |  | The thing that cannot run until another job completes. |
| `dependee_id` | `ForeignKey` | `integer` |  | FK |  | The thing that has other jobs waiting on it to complete. |
| `wait_for_completion` | `BooleanField` | `bool` |  |  | `True` | If checked, the dependent job will not run until the dependee job has completed. |
| `wait_for_success` | `BooleanField` | `bool` |  |  | `True` | If checked, the dependent job will not run until the dependee job has completed successfully. |
| `wait_for_next_run` | `BooleanField` | `bool` |  |  | `True` | If checked, the dependent job will not run until the dependee job has a next_run greater than its next_run. |

Indexes:

- index `chroniker_jobdependency_dependent_id_48fb0cbb` on `dependent_id`
- index `chroniker_jobdependency_dependee_id_5107f29f` on `dependee_id`
- unique `chroniker_jobdependency_dependent_id_dependee_id_a59a3257_uniq` on `dependent_id`, `dependee_id`

Relations:

- `dependent_id` → [Job](#job).`id`, on delete cascade
- `dependee_id` → [Job](#job).`id`, on delete cascade

## Job

Table `chroniker_job`, app `chroniker`.

Proxy models, on this table: `Monitor`.

A recurring ``django-admin`` command to be run.

![Job and the models it relates to](tables/Job.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(200)` |  |  |  |  |
| `frequency` | `JobFrequency` | `varchar(10)` |  |  |  |  |
| `params` | `TextField` | `text` | yes |  |  | Semicolon-separated list of <a href="http://labix.org/python-dateutil" target="_blank">rrule parameters</a>. e.g: interval:15;byhour:7,8,9 |
| `command` | `CharField` | `varchar(200)` |  |  |  | A valid django-admin command to run. |
| `args` | `CharField` | `varchar(200)` |  |  |  | Space separated list; e.g: arg1 option1=True |
| `raw_command` | `CharField` | `varchar(1000)` | yes |  |  | The raw shell command to run.             This is mutually exclusive with `command`. |
| `enabled` | `BooleanField` | `bool` |  |  | `True` | If checked, this job will be run automatically according             to the frequency options. |
| `next_run` | `DateTimeField` | `datetime` | yes |  |  | If you don't set this it will be determined automatically |
| `jitter_seconds` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | If given, a random offset of up to this many seconds is added to each scheduled run time.<br/>Useful for spreading load when several jobs or servers would otherwise fire at the same instant.<br/>A value of 0 disables jitter. |
| `last_run_start_timestamp` | `DateTimeField` | `datetime` | yes |  |  |  |
| `last_run` | `DateTimeField` | `datetime` | yes |  |  |  |
| `last_heartbeat` | `DateTimeField` | `datetime` | yes |  |  |  |
| `is_running` | `BooleanField` | `bool` |  |  | `False` |  |
| `last_run_successful` | `BooleanField` | `bool` | yes |  |  |  |
| `email_errors_to_subscribers` | `BooleanField` | `bool` |  |  | `True` | If checked, the stdout and stderr of a job will be emailed to the subscribers if an error occur. |
| `email_success_to_subscribers` | `BooleanField` | `bool` |  |  | `False` | If checked, the stdout of a job will be emailed to the subscribers if not errors occur. |
| `lock_file` | `CharField` | `varchar(255)` |  |  |  |  |
| `force_run` | `BooleanField` | `bool` |  |  | `False` | If checked, then this job will be run immediately. |
| `force_stop` | `BooleanField` | `bool` |  |  | `False` | If checked, and running then this job will be stopped. |
| `timeout_seconds` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | When non-zero, the job will be forcibly killed if             running for more than this amount of time. |
| `hostname` | `CharField` | `varchar(700)` | yes |  |  | <function hostname_help_text_setter at 0xffff827243b0> |
| `current_hostname` | `CharField` | `varchar(700)` | yes |  |  | The name of the host currently running the job. |
| `current_pid` | `CharField` | `varchar(50)` | yes |  |  | The ID of the process currently running the job. |
| `total_parts_complete` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | The total number of complete parts. |
| `total_parts` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | The total number of parts of the task. |
| `is_monitor` | `BooleanField` | `bool` |  |  | `False` | If checked, will appear in the monitors section. |
| `monitor_url` | `CharField` | `varchar(255)` | yes |  |  | URL provided to further explain the monitor. |
| `monitor_error_template` | `TextField` | `text` | yes |  | `'\nThe monitor "{{ job.name }}" has indicated a problem.\n\nPlease review this monitor at {{ url }}\n\n{{ job.monitor_description_safe }}\n\n{{ stderr }}\n'` | If this is a monitor, this is the template used to compose the error text email.<br/>Available variables: {{ job }} {{ stderr }} {{ url }} |
| `monitor_description` | `TextField` | `text` | yes |  |  | An explanation of the monitor's purpose. |
| `monitor_records` | `IntegerField` | `integer` | yes |  |  | The number of records that need attention. |
| `maximum_log_entries` | `PositiveIntegerField` | `integer unsigned` |  |  | `1000` | The maximum number of most recent log entries to keep.<br/>A value of 0 keeps all log entries. |
| `log_stdout` | `BooleanField` | `bool` |  |  | `True` | If checked, all characters printed to stdout will be             saved in a log record. |
| `log_stderr` | `BooleanField` | `bool` |  |  | `True` | If checked, all characters printed to stderr will be             saved in a log record. |

Indexes:

- index `chroniker_job_current_pid_9996fbd0` on `current_pid`

Relations:

- [JobDependency](#jobdependency).`dependent_id` → `id`, on delete cascade
- [JobDependency](#jobdependency).`dependee_id` → `id`, on delete cascade
- `subscribers` — many to many with [User](#user)
- [Log](#log).`job_id` → `id`, on delete cascade

## Log

Table `chroniker_log`, app `chroniker`.

A record of stdout and stderr of a ``Job``.

![Log and the models it relates to](tables/Log.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `job_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `run_start_datetime` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `run_end_datetime` | `DateTimeField` | `datetime` | yes |  |  |  |
| `duration_seconds` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |
| `stdout` | `TextField` | `text` |  |  |  |  |
| `stderr` | `TextField` | `text` |  |  |  |  |
| `hostname` | `CharField` | `varchar(700)` | yes |  |  | The hostname this job was executed on. |
| `success` | `BooleanField` | `bool` |  |  | `True` |  |
| `on_time` | `BooleanField` | `bool` |  |  | `True` | If true, indicates job completed of its own accord.             If false, the job exceeded a timeout threshold and was forcibly             killed. |

Indexes:

- index `chroniker_log_job_id_b73a0dc9` on `job_id`
- index `chroniker_log_run_start_datetime_ed601856` on `run_start_datetime`
- index `chroniker_log_run_end_datetime_ee9ac252` on `run_end_datetime`
- index `chroniker_log_duration_seconds_34305844` on `duration_seconds`
- index `chroniker_log_success_9d4f1251` on `success`
- index `chroniker_log_on_time_75c95897` on `on_time`

Relations:

- `job_id` → [Job](#job).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## JobFrequency

| Name | Stored as |
| --- | --- |
| `Yearly` | `YEARLY` |
| `Monthly` | `MONTHLY` |
| `Weekly` | `WEEKLY` |
| `Daily` | `DAILY` |
| `Hourly` | `HOURLY` |
| `Minutely` | `MINUTELY` |
| `Secondly` | `SECONDLY` |
