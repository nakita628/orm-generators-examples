# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [StudentItem](#studentitem) · [TeamSubmission](#teamsubmission) · [Submission](#submission) · [Score](#score) · [ScoreSummary](#scoresummary) · [ScoreAnnotation](#scoreannotation) · [ExternalGraderDetail](#externalgraderdetail) · [SubmissionFile](#submissionfile)

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

CREATE TABLE "submissions_studentitem" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "student_id" varchar(255) NOT NULL, "course_id" varchar(255) NOT NULL, "item_id" varchar(255) NOT NULL, "item_type" varchar(100) NOT NULL);

CREATE TABLE "submissions_teamsubmission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "created" datetime NOT NULL, "modified" datetime NOT NULL, "uuid" char(32) NOT NULL, "attempt_number" integer unsigned NOT NULL CHECK ("attempt_number" >= 0), "submitted_at" datetime NOT NULL, "course_id" varchar(255) NOT NULL, "item_id" varchar(255) NOT NULL, "team_id" varchar(255) NOT NULL, "submitted_by_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "status" varchar(1) NOT NULL);

CREATE TABLE "submissions_submission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "uuid" char(32) NOT NULL, "student_item_id" bigint NOT NULL REFERENCES "submissions_studentitem" ("id") DEFERRABLE INITIALLY DEFERRED, "attempt_number" integer unsigned NOT NULL CHECK ("attempt_number" >= 0), "submitted_at" datetime NOT NULL, "created_at" datetime NOT NULL, "raw_answer" text NOT NULL, "status" varchar(1) NOT NULL, "team_submission_id" bigint NULL REFERENCES "submissions_teamsubmission" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "submissions_score" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "student_item_id" bigint NOT NULL REFERENCES "submissions_studentitem" ("id") DEFERRABLE INITIALLY DEFERRED, "submission_id" bigint NULL REFERENCES "submissions_submission" ("id") DEFERRABLE INITIALLY DEFERRED, "points_earned" integer unsigned NOT NULL CHECK ("points_earned" >= 0), "points_possible" integer unsigned NOT NULL CHECK ("points_possible" >= 0), "created_at" datetime NOT NULL, "reset" bool NOT NULL);

CREATE TABLE "submissions_scoresummary" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "student_item_id" bigint NOT NULL UNIQUE REFERENCES "submissions_studentitem" ("id") DEFERRABLE INITIALLY DEFERRED, "highest_id" bigint NOT NULL REFERENCES "submissions_score" ("id") DEFERRABLE INITIALLY DEFERRED, "latest_id" bigint NOT NULL REFERENCES "submissions_score" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "submissions_scoreannotation" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "score_id" bigint NOT NULL REFERENCES "submissions_score" ("id") DEFERRABLE INITIALLY DEFERRED, "annotation_type" varchar(255) NOT NULL, "creator" varchar(255) NOT NULL, "reason" text NOT NULL);

CREATE TABLE "submissions_externalgraderdetail" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "submission_id" bigint NOT NULL UNIQUE REFERENCES "submissions_submission" ("id") DEFERRABLE INITIALLY DEFERRED, "queue_name" varchar(128) NOT NULL, "queue_key" varchar(128) NULL, "grader_file_name" varchar(128) NOT NULL, "points_possible" integer unsigned NOT NULL CHECK ("points_possible" >= 0), "status" varchar(20) NOT NULL, "pullkey" varchar(128) NULL, "grader_reply" text NULL, "status_time" datetime NOT NULL, "created_at" datetime NOT NULL, "num_failures" integer unsigned NOT NULL CHECK ("num_failures" >= 0));

CREATE TABLE "submissions_submissionfile" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "uuid" char(32) NOT NULL, "external_grader_id" bigint NULL REFERENCES "submissions_externalgraderdetail" ("id") DEFERRABLE INITIALLY DEFERRED, "file" varchar(512) NOT NULL, "original_filename" varchar(255) NOT NULL, "created_at" datetime NOT NULL);

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

CREATE UNIQUE INDEX "submissions_studentitem_course_id_student_id_item_id_5b02ecf8_uniq" ON "submissions_studentitem" ("course_id", "student_id", "item_id");

CREATE INDEX "submissions_studentitem_student_id_8e72bcd9" ON "submissions_studentitem" ("student_id");

CREATE INDEX "submissions_studentitem_course_id_05ee1efe" ON "submissions_studentitem" ("course_id");

CREATE INDEX "submissions_studentitem_item_id_6c409784" ON "submissions_studentitem" ("item_id");

CREATE INDEX "submissions_teamsubmission_uuid_4d1aef87" ON "submissions_teamsubmission" ("uuid");

CREATE INDEX "submissions_teamsubmission_submitted_at_74e28ed6" ON "submissions_teamsubmission" ("submitted_at");

CREATE INDEX "submissions_teamsubmission_course_id_68c6908d" ON "submissions_teamsubmission" ("course_id");

CREATE INDEX "submissions_teamsubmission_item_id_2bdcb26c" ON "submissions_teamsubmission" ("item_id");

CREATE INDEX "submissions_teamsubmission_team_id_5fda0e54" ON "submissions_teamsubmission" ("team_id");

CREATE INDEX "submissions_teamsubmission_submitted_by_id_5a27162a" ON "submissions_teamsubmission" ("submitted_by_id");

CREATE INDEX "submissions_submission_uuid_210428ab" ON "submissions_submission" ("uuid");

CREATE INDEX "submissions_submission_student_item_id_9d087470" ON "submissions_submission" ("student_item_id");

CREATE INDEX "submissions_submission_submitted_at_9653124d" ON "submissions_submission" ("submitted_at");

CREATE INDEX "submissions_submission_created_at_01c4bf22" ON "submissions_submission" ("created_at");

CREATE INDEX "submissions_submission_team_submission_id_40e6bc97" ON "submissions_submission" ("team_submission_id");

CREATE INDEX "submissions_score_student_item_id_de4f5954" ON "submissions_score" ("student_item_id");

CREATE INDEX "submissions_score_submission_id_ba095829" ON "submissions_score" ("submission_id");

CREATE INDEX "submissions_score_created_at_b65f0390" ON "submissions_score" ("created_at");

CREATE INDEX "submissions_scoresummary_highest_id_3efe897d" ON "submissions_scoresummary" ("highest_id");

CREATE INDEX "submissions_scoresummary_latest_id_dd8a17bb" ON "submissions_scoresummary" ("latest_id");

CREATE INDEX "submissions_scoreannotation_score_id_2dda6e02" ON "submissions_scoreannotation" ("score_id");

CREATE INDEX "submissions_scoreannotation_annotation_type_117a2607" ON "submissions_scoreannotation" ("annotation_type");

CREATE INDEX "submissions_scoreannotation_creator_5cc126cc" ON "submissions_scoreannotation" ("creator");

CREATE INDEX "submissions_externalgraderdetail_status_time_1f2a6499" ON "submissions_externalgraderdetail" ("status_time");

CREATE INDEX "submissions_externalgraderdetail_created_at_e1618de1" ON "submissions_externalgraderdetail" ("created_at");

CREATE INDEX "submissions_queue_n_fbabd8_idx" ON "submissions_externalgraderdetail" ("queue_name", "status", "status_time");

CREATE INDEX "submissions_submissionfile_external_grader_id_b21f9730" ON "submissions_submissionfile" ("external_grader_id");

CREATE INDEX "submissions_externa_ff8089_idx" ON "submissions_submissionfile" ("external_grader_id", "uuid");
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
- [TeamSubmission](#teamsubmission).`submitted_by_id` → `id`, on delete set null

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

## StudentItem

Table `submissions_studentitem`, app `submissions`.

Represents a single item for a single course for a single user.

This is typically an XBlock problem, but could be something more generic
like class participation.

.. no_pii:

![StudentItem and the models it relates to](tables/StudentItem.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `student_id` | `AnonymizedUserIDField` | `varchar(255)` |  |  |  |  |
| `course_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `item_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `item_type` | `CharField` | `varchar(100)` |  |  |  |  |

Indexes:

- index `submissions_studentitem_student_id_8e72bcd9` on `student_id`
- index `submissions_studentitem_course_id_05ee1efe` on `course_id`
- index `submissions_studentitem_item_id_6c409784` on `item_id`
- unique `submissions_studentitem_course_id_student_id_item_id_5b02ecf8_uniq` on `course_id`, `student_id`, `item_id`

Relations:

- [Submission](#submission).`student_item_id` → `id`, on delete cascade
- [Score](#score).`student_item_id` → `id`, on delete cascade
- [ScoreSummary](#scoresummary).`student_item_id` → `id`, on delete cascade

## TeamSubmission

Table `submissions_teamsubmission`, app `submissions`.

Columns from `TimeStampedModel` (abstract): `created`, `modified`.

A single response by a team for a given problem (ORA2) in a given course.
An abstraction layer over Submission used to for teams. Since we create a submission record for every team member,
there is a need to have a single point to connect the team to the workflows.
TeamSubmission is a 1 to many with Submission

.. no_pii:

![TeamSubmission and the models it relates to](tables/TeamSubmission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `uuid` | `UUIDField` | `char(32)` |  |  | `uuid4()` |  |
| `attempt_number` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `submitted_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `course_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `item_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `team_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `submitted_by_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `status` | `TeamSubmissionStatus` | `varchar(1)` |  |  | `'A'` |  |

Indexes:

- index `submissions_teamsubmission_uuid_4d1aef87` on `uuid`
- index `submissions_teamsubmission_submitted_at_74e28ed6` on `submitted_at`
- index `submissions_teamsubmission_course_id_68c6908d` on `course_id`
- index `submissions_teamsubmission_item_id_2bdcb26c` on `item_id`
- index `submissions_teamsubmission_team_id_5fda0e54` on `team_id`
- index `submissions_teamsubmission_submitted_by_id_5a27162a` on `submitted_by_id`

Relations:

- `submitted_by_id` → [User](#user).`id`, on delete set null
- [Submission](#submission).`team_submission_id` → `id`, on delete set null

## Submission

Table `submissions_submission`, app `submissions`.

A single response by a student for a given problem in a given course.

A student may have multiple submissions for the same problem. Submissions
should never be mutated. If the student makes another Submission, or if we
have to make corrective Submissions to fix bugs, those should be new
objects. We want to keep Submissions immutable both for audit purposes, and
because it makes caching trivial.

.. no_pii:

![Submission and the models it relates to](tables/Submission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `uuid` | `UUIDField` | `char(32)` |  |  | `uuid4()` |  |
| `student_item_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `attempt_number` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `submitted_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `raw_answer` | `JSONField` | `text` |  |  |  |  |
| `status` | `SubmissionStatus` | `varchar(1)` |  |  | `'A'` |  |
| `team_submission_id` | `ForeignKey` | `bigint` | yes | FK |  |  |

Indexes:

- index `submissions_submission_uuid_210428ab` on `uuid`
- index `submissions_submission_student_item_id_9d087470` on `student_item_id`
- index `submissions_submission_submitted_at_9653124d` on `submitted_at`
- index `submissions_submission_created_at_01c4bf22` on `created_at`
- index `submissions_submission_team_submission_id_40e6bc97` on `team_submission_id`

Relations:

- `student_item_id` → [StudentItem](#studentitem).`id`, on delete cascade
- `team_submission_id` → [TeamSubmission](#teamsubmission).`id`, on delete set null
- [Score](#score).`submission_id` → `id`, on delete cascade
- [ExternalGraderDetail](#externalgraderdetail).`submission_id` → `id`, on delete cascade

## Score

Table `submissions_score`, app `submissions`.

What the user scored for a given StudentItem at a given time.

Note that while a Score *can* be tied to a Submission, it doesn't *have* to.
Specifically, if we want to have scores for things that are not a part of
the courseware (like "class participation"), there would be no corresponding
Submission.

.. no_pii:

![Score and the models it relates to](tables/Score.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `student_item_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `submission_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `points_earned` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `points_possible` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `reset` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `submissions_score_student_item_id_de4f5954` on `student_item_id`
- index `submissions_score_submission_id_ba095829` on `submission_id`
- index `submissions_score_created_at_b65f0390` on `created_at`

Relations:

- `student_item_id` → [StudentItem](#studentitem).`id`, on delete cascade
- `submission_id` → [Submission](#submission).`id`, on delete cascade
- [ScoreSummary](#scoresummary).`highest_id` → `id`, on delete cascade
- [ScoreSummary](#scoresummary).`latest_id` → `id`, on delete cascade
- [ScoreAnnotation](#scoreannotation).`score_id` → `id`, on delete cascade

## ScoreSummary

Table `submissions_scoresummary`, app `submissions`.

Running store of the highest and most recent Scores for a StudentItem.

.. no_pii:

![ScoreSummary and the models it relates to](tables/ScoreSummary.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `student_item_id` | `OneToOneField` | `bigint` |  | FK, unique |  |  |
| `highest_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `latest_id` | `ForeignKey` | `bigint` |  | FK |  |  |

Indexes:

- index `submissions_scoresummary_highest_id_3efe897d` on `highest_id`
- index `submissions_scoresummary_latest_id_dd8a17bb` on `latest_id`

Relations:

- `student_item_id` → [StudentItem](#studentitem).`id`, on delete cascade
- `highest_id` → [Score](#score).`id`, on delete cascade
- `latest_id` → [Score](#score).`id`, on delete cascade

## ScoreAnnotation

Table `submissions_scoreannotation`, app `submissions`.

Annotate individual scores with extra information if necessary.

.. no_pii:

![ScoreAnnotation and the models it relates to](tables/ScoreAnnotation.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `score_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `annotation_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `creator` | `AnonymizedUserIDField` | `varchar(255)` |  |  |  |  |
| `reason` | `TextField` | `text` |  |  |  |  |

Indexes:

- index `submissions_scoreannotation_score_id_2dda6e02` on `score_id`
- index `submissions_scoreannotation_annotation_type_117a2607` on `annotation_type`
- index `submissions_scoreannotation_creator_5cc126cc` on `creator`

Relations:

- `score_id` → [Score](#score).`id`, on delete cascade

## ExternalGraderDetail

Table `submissions_externalgraderdetail`, app `submissions`.

Tracks queue processing information for a Submission.

![ExternalGraderDetail and the models it relates to](tables/ExternalGraderDetail.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `submission_id` | `OneToOneField` | `bigint` |  | FK, unique |  |  |
| `queue_name` | `CharField` | `varchar(128)` |  |  |  |  |
| `queue_key` | `CharField` | `varchar(128)` | yes |  |  |  |
| `grader_file_name` | `CharField` | `varchar(128)` |  |  | `''` |  |
| `points_possible` | `PositiveIntegerField` | `integer unsigned` |  |  | `1` |  |
| `status` | `ExternalGraderDetailStatus` | `varchar(20)` |  |  | `'pending'` |  |
| `pullkey` | `CharField` | `varchar(128)` | yes |  |  |  |
| `grader_reply` | `TextField` | `text` | yes |  |  |  |
| `status_time` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `num_failures` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |

Indexes:

- index `submissions_externalgraderdetail_status_time_1f2a6499` on `status_time`
- index `submissions_externalgraderdetail_created_at_e1618de1` on `created_at`
- index `submissions_queue_n_fbabd8_idx` on `queue_name`, `status`, `status_time`

Relations:

- `submission_id` → [Submission](#submission).`id`, on delete cascade
- [SubmissionFile](#submissionfile).`external_grader_id` → `id`, on delete set null

## SubmissionFile

Table `submissions_submissionfile`, app `submissions`.

Model to handle files associated with submissions

![SubmissionFile and the models it relates to](tables/SubmissionFile.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `uuid` | `UUIDField` | `char(32)` |  |  | `uuid4()` |  |
| `external_grader_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `file` | `FileField` | `varchar(512)` |  |  |  |  |
| `original_filename` | `CharField` | `varchar(255)` |  |  |  |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `submissions_submissionfile_external_grader_id_b21f9730` on `external_grader_id`
- index `submissions_externa_ff8089_idx` on `external_grader_id`, `uuid`

Relations:

- `external_grader_id` → [ExternalGraderDetail](#externalgraderdetail).`id`, on delete set null

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## TeamSubmissionStatus

| Name | Stored as |
| --- | --- |
| `Deleted` | `D` |
| `Active` | `A` |

## SubmissionStatus

| Name | Stored as |
| --- | --- |
| `Deleted` | `D` |
| `Active` | `A` |

## ExternalGraderDetailStatus

| Name | Stored as |
| --- | --- |
| `PENDING` | `pending` |
| `PULLED` | `pulled` |
| `RETIRED` | `retired` |
| `FAILED` | `failed` |
| `RETRY` | `retry` |
