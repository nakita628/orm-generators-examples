# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Survey](#survey) · [Category](#category) · [Question](#question) · [Response](#response) · [Answer](#answer)

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

CREATE TABLE "survey_survey" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(400) NOT NULL, "description" text NOT NULL, "is_published" bool NOT NULL, "need_logged_user" bool NOT NULL, "editable_answers" bool NOT NULL, "display_method" smallint NOT NULL, "template" varchar(255) NULL, "publish_date" date NOT NULL, "expire_date" date NOT NULL, "redirect_url" varchar(200) NOT NULL);

CREATE TABLE "survey_category" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(400) NOT NULL, "survey_id" bigint NOT NULL REFERENCES "survey_survey" ("id") DEFERRABLE INITIALLY DEFERRED, "order" integer NULL, "description" varchar(2000) NULL);

CREATE TABLE "survey_question" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "text" text NOT NULL, "order" integer NOT NULL, "required" bool NOT NULL, "category_id" bigint NULL REFERENCES "survey_category" ("id") DEFERRABLE INITIALLY DEFERRED, "survey_id" bigint NOT NULL REFERENCES "survey_survey" ("id") DEFERRABLE INITIALLY DEFERRED, "type" varchar(200) NOT NULL, "choices" text NULL, "maximum_choices" integer unsigned NULL CHECK ("maximum_choices" >= 0));

CREATE TABLE "survey_response" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "created" datetime NOT NULL, "updated" datetime NOT NULL, "survey_id" bigint NOT NULL REFERENCES "survey_survey" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "interview_uuid" varchar(36) NOT NULL);

CREATE TABLE "survey_answer" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "question_id" bigint NOT NULL REFERENCES "survey_question" ("id") DEFERRABLE INITIALLY DEFERRED, "response_id" bigint NOT NULL REFERENCES "survey_response" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL, "updated" datetime NOT NULL, "body" text NULL);

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

CREATE INDEX "survey_category_survey_id_bb87ed31" ON "survey_category" ("survey_id");

CREATE INDEX "survey_question_category_id_30d30593" ON "survey_question" ("category_id");

CREATE INDEX "survey_question_survey_id_aa755a68" ON "survey_question" ("survey_id");

CREATE INDEX "survey_response_survey_id_be749458" ON "survey_response" ("survey_id");

CREATE INDEX "survey_response_user_id_3e2a2e81" ON "survey_response" ("user_id");

CREATE INDEX "survey_answer_question_id_1e622dbb" ON "survey_answer" ("question_id");

CREATE INDEX "survey_answer_response_id_df20db0c" ON "survey_answer" ("response_id");
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
- [Response](#response).`user_id` → `id`, on delete set null

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

## Survey

Table `survey_survey`, app `survey`.

![Survey and the models it relates to](tables/Survey.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(400)` |  |  |  |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `is_published` | `BooleanField` | `bool` |  |  | `True` |  |
| `need_logged_user` | `BooleanField` | `bool` |  |  |  |  |
| `editable_answers` | `BooleanField` | `bool` |  |  | `True` |  |
| `display_method` | `SurveyDisplayMethod` | `smallint` |  |  | `0` |  |
| `template` | `CharField` | `varchar(255)` | yes |  |  |  |
| `publish_date` | `DateField` | `date` |  |  | `now()` |  |
| `expire_date` | `DateField` | `date` |  |  | `in_duration_day()` |  |
| `redirect_url` | `URLField` | `varchar(200)` |  |  |  |  |

Relations:

- [Category](#category).`survey_id` → `id`, on delete cascade
- [Question](#question).`survey_id` → `id`, on delete cascade
- [Response](#response).`survey_id` → `id`, on delete cascade

## Category

Table `survey_category`, app `survey`.

![Category and the models it relates to](tables/Category.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(400)` |  |  |  |  |
| `survey_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `order` | `IntegerField` | `integer` | yes |  |  |  |
| `description` | `CharField` | `varchar(2000)` | yes |  |  |  |

Indexes:

- index `survey_category_survey_id_bb87ed31` on `survey_id`

Relations:

- `survey_id` → [Survey](#survey).`id`, on delete cascade
- [Question](#question).`category_id` → `id`, on delete set null

## Question

Table `survey_question`, app `survey`.

![Question and the models it relates to](tables/Question.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `text` | `TextField` | `text` |  |  |  |  |
| `order` | `IntegerField` | `integer` |  |  |  |  |
| `required` | `BooleanField` | `bool` |  |  |  |  |
| `category_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `survey_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `type` | `QuestionType` | `varchar(200)` |  |  | `'text'` |  |
| `choices` | `TextField` | `text` | yes |  |  | The choices field is only used if the question type if the question type is 'radio', 'select', or 'select multiple' provide a comma-separated list of options for this question . |
| `maximum_choices` | `PositiveIntegerField` | `integer unsigned` | yes |  |  | Only used for 'select multiple' questions. Leave empty for no limit. |

Indexes:

- index `survey_question_category_id_30d30593` on `category_id`
- index `survey_question_survey_id_aa755a68` on `survey_id`

Relations:

- `category_id` → [Category](#category).`id`, on delete set null
- `survey_id` → [Survey](#survey).`id`, on delete cascade
- [Answer](#answer).`question_id` → `id`, on delete cascade

## Response

Table `survey_response`, app `survey`.

A Response object is a collection of questions and answers with a
unique interview uuid.

![Response and the models it relates to](tables/Response.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `survey_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `interview_uuid` | `CharField` | `varchar(36)` |  |  |  |  |

Indexes:

- index `survey_response_survey_id_be749458` on `survey_id`
- index `survey_response_user_id_3e2a2e81` on `user_id`

Relations:

- `survey_id` → [Survey](#survey).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete set null
- [Answer](#answer).`response_id` → `id`, on delete cascade

## Answer

Table `survey_answer`, app `survey`.

![Answer and the models it relates to](tables/Answer.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `question_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `response_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `body` | `TextField` | `text` | yes |  |  |  |

Indexes:

- index `survey_answer_question_id_1e622dbb` on `question_id`
- index `survey_answer_response_id_df20db0c` on `response_id`

Relations:

- `question_id` → [Question](#question).`id`, on delete cascade
- `response_id` → [Response](#response).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## SurveyDisplayMethod

| Name | Stored as |
| --- | --- |
| `By question` | `1` |
| `By category` | `2` |
| `All in one page` | `0` |

## QuestionType

| Name | Stored as |
| --- | --- |
| `text (multiple line)` | `text` |
| `short text (one line)` | `short-text` |
| `radio` | `radio` |
| `select` | `select` |
| `Select Multiple` | `select-multiple` |
| `Select Image` | `select_image` |
| `integer` | `integer` |
| `float` | `float` |
| `date` | `date` |
