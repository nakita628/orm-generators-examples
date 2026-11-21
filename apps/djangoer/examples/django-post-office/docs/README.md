# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Email](#email) · [Log](#log) · [EmailTemplate](#emailtemplate) · [Attachment](#attachment)

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

CREATE TABLE "post_office_email" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "from_email" varchar(254) NOT NULL, "to" text NOT NULL, "cc" text NOT NULL, "bcc" text NOT NULL, "subject" varchar(989) NOT NULL, "message" text NOT NULL, "html_message" text NOT NULL, "status" smallint unsigned NULL CHECK ("status" >= 0), "recipient_delivery_status" smallint unsigned NULL CHECK ("recipient_delivery_status" >= 0), "priority" smallint unsigned NULL CHECK ("priority" >= 0), "created" datetime NOT NULL, "last_updated" datetime NOT NULL, "scheduled_time" datetime NULL, "expires_at" datetime NULL, "message_id" varchar(255) NULL, "number_of_retries" integer unsigned NULL CHECK ("number_of_retries" >= 0), "headers" text NULL CHECK ((JSON_VALID("headers") OR "headers" IS NULL)), "template_id" integer NULL REFERENCES "post_office_emailtemplate" ("id") DEFERRABLE INITIALLY DEFERRED, "context" text NULL CHECK ((JSON_VALID("context") OR "context" IS NULL)), "backend_alias" varchar(64) NOT NULL);

CREATE TABLE "post_office_log" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "email_id" integer NOT NULL REFERENCES "post_office_email" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL, "status" smallint unsigned NOT NULL CHECK ("status" >= 0), "exception_type" varchar(255) NOT NULL, "message" text NOT NULL);

CREATE TABLE "post_office_emailtemplate" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(255) NOT NULL, "description" text NOT NULL, "created" datetime NOT NULL, "last_updated" datetime NOT NULL, "subject" varchar(255) NOT NULL, "content" text NOT NULL, "html_content" text NOT NULL, "language" varchar(12) NOT NULL, "default_template_id" integer NULL REFERENCES "post_office_emailtemplate" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "post_office_attachment" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "file" varchar(100) NOT NULL, "name" varchar(255) NOT NULL, "mimetype" varchar(255) NOT NULL, "headers" text NULL CHECK ((JSON_VALID("headers") OR "headers" IS NULL)));

CREATE TABLE "post_office_attachment_emails" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "attachment_id" integer NOT NULL REFERENCES "post_office_attachment" ("id") DEFERRABLE INITIALLY DEFERRED, "email_id" integer NOT NULL REFERENCES "post_office_email" ("id") DEFERRABLE INITIALLY DEFERRED);

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

CREATE INDEX "post_office_email_status_013a896c" ON "post_office_email" ("status");

CREATE INDEX "post_office_email_created_1306952f" ON "post_office_email" ("created");

CREATE INDEX "post_office_email_last_updated_0ffcec35" ON "post_office_email" ("last_updated");

CREATE INDEX "post_office_email_scheduled_time_3869ebec" ON "post_office_email" ("scheduled_time");

CREATE INDEX "post_office_email_template_id_417da7da" ON "post_office_email" ("template_id");

CREATE INDEX "post_office_log_email_id_d42c8808" ON "post_office_log" ("email_id");

CREATE UNIQUE INDEX "post_office_emailtemplate_name_language_default_template_id_4023e3e4_uniq" ON "post_office_emailtemplate" ("name", "language", "default_template_id");

CREATE INDEX "post_office_emailtemplate_default_template_id_2ac2f889" ON "post_office_emailtemplate" ("default_template_id");

CREATE UNIQUE INDEX "post_office_attachment_emails_attachment_id_email_id_8e046917_uniq" ON "post_office_attachment_emails" ("attachment_id", "email_id");

CREATE INDEX "post_office_attachment_emails_attachment_id_6136fd9a" ON "post_office_attachment_emails" ("attachment_id");

CREATE INDEX "post_office_attachment_emails_email_id_96875fd9" ON "post_office_attachment_emails" ("email_id");
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

## Email

Table `post_office_email`, app `post_office`.

A model to hold email information.

![Email and the models it relates to](tables/Email.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `from_email` | `CharField` | `varchar(254)` |  |  |  |  |
| `to` | `CommaSeparatedEmailField` | `text` |  |  |  |  |
| `cc` | `CommaSeparatedEmailField` | `text` |  |  |  |  |
| `bcc` | `CommaSeparatedEmailField` | `text` |  |  |  |  |
| `subject` | `CharField` | `varchar(989)` |  |  |  |  |
| `message` | `TextField` | `text` |  |  |  |  |
| `html_message` | `TextField` | `text` |  |  |  |  |
| `status` | `EmailStatus` | `smallint unsigned` | yes |  |  |  |
| `recipient_delivery_status` | `RecipientDeliveryStatus` | `smallint unsigned` | yes |  |  |  |
| `priority` | `EmailPriority` | `smallint unsigned` | yes |  |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `last_updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `scheduled_time` | `DateTimeField` | `datetime` | yes |  |  | The scheduled sending time |
| `expires_at` | `DateTimeField` | `datetime` | yes |  |  | Email won't be sent after this timestamp |
| `message_id` | `CharField` | `varchar(255)` | yes |  |  |  |
| `number_of_retries` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |
| `headers` | `JSONField` | `text` | yes |  |  |  |
| `template_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `context` | `JSONField` | `text` | yes |  |  |  |
| `backend_alias` | `CharField` | `varchar(64)` |  |  | `''` |  |

Indexes:

- index `post_office_email_status_013a896c` on `status`
- index `post_office_email_created_1306952f` on `created`
- index `post_office_email_last_updated_0ffcec35` on `last_updated`
- index `post_office_email_scheduled_time_3869ebec` on `scheduled_time`
- index `post_office_email_template_id_417da7da` on `template_id`

Relations:

- `template_id` → [EmailTemplate](#emailtemplate).`id`, on delete cascade
- [Log](#log).`email_id` → `id`, on delete cascade
- `attachments` — many to many with [Attachment](#attachment)

## Log

Table `post_office_log`, app `post_office`.

A model to record sending email sending activities.

![Log and the models it relates to](tables/Log.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `email_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `status` | `LogStatus` | `smallint unsigned` |  |  |  |  |
| `exception_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `message` | `TextField` | `text` |  |  |  |  |

Indexes:

- index `post_office_log_email_id_d42c8808` on `email_id`

Relations:

- `email_id` → [Email](#email).`id`, on delete cascade

## EmailTemplate

Table `post_office_emailtemplate`, app `post_office`.

Model to hold template information from db

![EmailTemplate and the models it relates to](tables/EmailTemplate.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(255)` |  |  |  | e.g: 'welcome_email' |
| `description` | `TextField` | `text` |  |  |  | Description of this template. |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `last_updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `subject` | `CharField` | `varchar(255)` |  |  |  |  |
| `content` | `TextField` | `text` |  |  |  |  |
| `html_content` | `TextField` | `text` |  |  |  |  |
| `language` | `CharField` | `varchar(12)` |  |  | `''` | Render template in alternative language |
| `default_template_id` | `ForeignKey` | `integer` | yes | FK | `None` |  |

Indexes:

- index `post_office_emailtemplate_default_template_id_2ac2f889` on `default_template_id`
- unique `post_office_emailtemplate_name_language_default_template_id_4023e3e4_uniq` on `name`, `language`, `default_template_id`

Relations:

- [Email](#email).`template_id` → `id`, on delete cascade
- `default_template_id` → [EmailTemplate](#emailtemplate).`id`, on delete cascade

## Attachment

Table `post_office_attachment`, app `post_office`.

A model describing an email attachment.

![Attachment and the models it relates to](tables/Attachment.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `file` | `FileField` | `varchar(100)` |  |  |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  | The original filename |
| `mimetype` | `CharField` | `varchar(255)` |  |  | `''` |  |
| `headers` | `JSONField` | `text` | yes |  |  |  |

Relations:

- `emails` — many to many with [Email](#email)

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## EmailStatus

| Name | Stored as |
| --- | --- |
| `sent` | `0` |
| `failed` | `1` |
| `queued` | `2` |
| `requeued` | `3` |

## RecipientDeliveryStatus

| Name | Stored as |
| --- | --- |
| `ACCEPTED` | `10` |
| `DELIVERED` | `20` |
| `OPENED` | `30` |
| `CLICKED` | `40` |
| `DEFERRED` | `50` |
| `SOFT_BOUNCED` | `60` |
| `UNDETERMINED_BOUNCED` | `65` |
| `HARD_BOUNCED` | `70` |
| `SPAM_COMPLAINT` | `80` |
| `UNSUBSCRIBED` | `90` |

## EmailPriority

| Name | Stored as |
| --- | --- |
| `low` | `0` |
| `medium` | `1` |
| `high` | `2` |
| `now` | `3` |

## LogStatus

| Name | Stored as |
| --- | --- |
| `sent` | `0` |
| `failed` | `1` |
| `Accepted` | `10` |
| `Delivered` | `20` |
| `Opened` | `30` |
| `Clicked` | `40` |
| `Deferred` | `50` |
| `Soft Bounced` | `60` |
| `Undetermined Bounced` | `65` |
| `Hard Bounced` | `70` |
| `Spam Complaint` | `80` |
| `Unsubscribed` | `90` |
