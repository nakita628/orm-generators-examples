# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Site](#site) · [KVStore](#kvstore) · [Newsletter](#newsletter) · [Subscription](#subscription) · [Article](#article) · [Attachment](#attachment) · [Message](#message) · [Submission](#submission)

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

CREATE TABLE "thumbnail_kvstore" ("key" varchar(200) NOT NULL PRIMARY KEY, "value" text NOT NULL);

CREATE TABLE "newsletter_newsletter" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "title" varchar(200) NOT NULL, "slug" varchar(50) NOT NULL UNIQUE, "email" varchar(254) NOT NULL, "sender" varchar(200) NOT NULL, "subscription_generator_class" varchar(200) NULL, "visible" bool NOT NULL, "send_html" bool NOT NULL, "enable_unsubscribe" bool NOT NULL);

CREATE TABLE "newsletter_newsletter_site" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "newsletter_id" bigint NOT NULL REFERENCES "newsletter_newsletter" ("id") DEFERRABLE INITIALLY DEFERRED, "site_id" integer NOT NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "newsletter_subscription" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(200) NULL, "email" varchar(254) NULL, "ip" char(39) NULL, "newsletter_id" bigint NOT NULL REFERENCES "newsletter_newsletter" ("id") DEFERRABLE INITIALLY DEFERRED, "create_date" datetime NOT NULL, "activation_code" varchar(40) NOT NULL, "subscribed" bool NOT NULL, "subscribe_date" datetime NULL, "unsubscribed" bool NOT NULL, "unsubscribe_date" datetime NULL);

CREATE TABLE "newsletter_article" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "sortorder" integer unsigned NOT NULL CHECK ("sortorder" >= 0), "title" varchar(200) NOT NULL, "text" text NOT NULL, "url" varchar(200) NULL, "image" varchar(100) NULL, "image_thumbnail_width" integer NULL, "image_use_original" bool NOT NULL, "image_below_text" bool NOT NULL, "post_id" bigint NOT NULL REFERENCES "newsletter_message" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "newsletter_attachment" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "file" varchar(100) NOT NULL, "message_id" bigint NOT NULL REFERENCES "newsletter_message" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "newsletter_message" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "title" varchar(200) NOT NULL, "slug" varchar(50) NOT NULL, "newsletter_id" bigint NOT NULL REFERENCES "newsletter_newsletter" ("id") DEFERRABLE INITIALLY DEFERRED, "date_create" datetime NOT NULL, "date_modify" datetime NOT NULL);

CREATE TABLE "newsletter_submission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "site_id" integer NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED, "newsletter_id" bigint NOT NULL REFERENCES "newsletter_newsletter" ("id") DEFERRABLE INITIALLY DEFERRED, "message_id" bigint NOT NULL REFERENCES "newsletter_message" ("id") DEFERRABLE INITIALLY DEFERRED, "publish_date" datetime NULL, "publish" bool NOT NULL, "prepared" bool NOT NULL, "sent" bool NOT NULL, "sending" bool NOT NULL);

CREATE TABLE "newsletter_submission_subscriptions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "submission_id" bigint NOT NULL REFERENCES "newsletter_submission" ("id") DEFERRABLE INITIALLY DEFERRED, "subscription_id" bigint NOT NULL REFERENCES "newsletter_subscription" ("id") DEFERRABLE INITIALLY DEFERRED);

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

CREATE INDEX "newsletter_newsletter_visible_db80e140" ON "newsletter_newsletter" ("visible");

CREATE UNIQUE INDEX "newsletter_newsletter_site_newsletter_id_site_id_eaadd4df_uniq" ON "newsletter_newsletter_site" ("newsletter_id", "site_id");

CREATE INDEX "newsletter_newsletter_site_newsletter_id_d5cf228e" ON "newsletter_newsletter_site" ("newsletter_id");

CREATE INDEX "newsletter_newsletter_site_site_id_2a546dc1" ON "newsletter_newsletter_site" ("site_id");

CREATE UNIQUE INDEX "newsletter_subscription_user_id_email_newsletter_id_0a9641e0_uniq" ON "newsletter_subscription" ("user_id", "email", "newsletter_id");

CREATE INDEX "newsletter_subscription_user_id_4185bd67" ON "newsletter_subscription" ("user_id");

CREATE INDEX "newsletter_subscription_email_3b8020aa" ON "newsletter_subscription" ("email");

CREATE INDEX "newsletter_subscription_newsletter_id_e9d08ab9" ON "newsletter_subscription" ("newsletter_id");

CREATE INDEX "newsletter_subscription_subscribed_350402fe" ON "newsletter_subscription" ("subscribed");

CREATE INDEX "newsletter_subscription_unsubscribed_8ae75c15" ON "newsletter_subscription" ("unsubscribed");

CREATE UNIQUE INDEX "newsletter_article_post_id_sortorder_cbe3c629_uniq" ON "newsletter_article" ("post_id", "sortorder");

CREATE INDEX "newsletter_article_post_id_c481e545" ON "newsletter_article" ("post_id");

CREATE INDEX "newsletter_attachment_message_id_52014d1f" ON "newsletter_attachment" ("message_id");

CREATE UNIQUE INDEX "newsletter_message_slug_newsletter_id_a2ab3653_uniq" ON "newsletter_message" ("slug", "newsletter_id");

CREATE INDEX "newsletter_message_slug_6dae36a9" ON "newsletter_message" ("slug");

CREATE INDEX "newsletter_message_newsletter_id_b77d9df1" ON "newsletter_message" ("newsletter_id");

CREATE INDEX "newsletter_submission_site_id_f4f389eb" ON "newsletter_submission" ("site_id");

CREATE INDEX "newsletter_submission_newsletter_id_8fa5ae76" ON "newsletter_submission" ("newsletter_id");

CREATE INDEX "newsletter_submission_message_id_35227c2e" ON "newsletter_submission" ("message_id");

CREATE INDEX "newsletter_submission_publish_date_6c3d8801" ON "newsletter_submission" ("publish_date");

CREATE INDEX "newsletter_submission_publish_2ceb081e" ON "newsletter_submission" ("publish");

CREATE INDEX "newsletter_submission_prepared_12eea099" ON "newsletter_submission" ("prepared");

CREATE INDEX "newsletter_submission_sent_7c1e3f56" ON "newsletter_submission" ("sent");

CREATE INDEX "newsletter_submission_sending_40b87529" ON "newsletter_submission" ("sending");

CREATE UNIQUE INDEX "newsletter_submission_subscriptions_submission_id_subscription_id_fa14b47b_uniq" ON "newsletter_submission_subscriptions" ("submission_id", "subscription_id");

CREATE INDEX "newsletter_submission_subscriptions_submission_id_0cda0c3f" ON "newsletter_submission_subscriptions" ("submission_id");

CREATE INDEX "newsletter_submission_subscriptions_subscription_id_be932b58" ON "newsletter_submission_subscriptions" ("subscription_id");
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
- [Subscription](#subscription).`user_id` → `id`, on delete cascade

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

![Site and the models it relates to](tables/Site.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `domain` | `CharField` | `varchar(100)` |  | unique |  |  |
| `name` | `CharField` | `varchar(50)` |  |  |  |  |

Relations:

- `newsletter_set` — many to many with [Newsletter](#newsletter)
- [Submission](#submission).`site_id` → `id`, on delete cascade

## KVStore

Table `thumbnail_kvstore`, app `thumbnail`.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `key` | `CharField` | `varchar(200)` |  | PK |  |  |
| `value` | `TextField` | `text` |  |  |  |  |

## Newsletter

Table `newsletter_newsletter`, app `newsletter`.

![Newsletter and the models it relates to](tables/Newsletter.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `title` | `CharField` | `varchar(200)` |  |  |  |  |
| `slug` | `SlugField` | `varchar(50)` |  | unique |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  | Sender e-mail |
| `sender` | `CharField` | `varchar(200)` |  |  |  | Sender name |
| `subscription_generator_class` | `CharField` | `varchar(200)` | yes |  |  | Class for generating subscriptions dynamically |
| `visible` | `BooleanField` | `bool` |  |  | `True` |  |
| `send_html` | `BooleanField` | `bool` |  |  | `True` | Whether or not to send HTML versions of e-mails. |
| `enable_unsubscribe` | `BooleanField` | `bool` |  |  | `True` | Enable unsubscribe links in e-mails. |

Indexes:

- index `newsletter_newsletter_visible_db80e140` on `visible`

Relations:

- `site` — many to many with [Site](#site)
- [Subscription](#subscription).`newsletter_id` → `id`, on delete cascade
- [Message](#message).`newsletter_id` → `id`, on delete cascade
- [Submission](#submission).`newsletter_id` → `id`, on delete cascade

## Subscription

Table `newsletter_subscription`, app `newsletter`.

![Subscription and the models it relates to](tables/Subscription.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `name` | `CharField` | `varchar(200)` | yes |  |  | optional |
| `email` | `EmailField` | `varchar(254)` | yes |  |  |  |
| `ip` | `GenericIPAddressField` | `char(39)` | yes |  |  |  |
| `newsletter_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `create_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `activation_code` | `CharField` | `varchar(40)` |  |  | `make_activation_code()` |  |
| `subscribed` | `BooleanField` | `bool` |  |  | `False` |  |
| `subscribe_date` | `DateTimeField` | `datetime` | yes |  |  |  |
| `unsubscribed` | `BooleanField` | `bool` |  |  | `False` |  |
| `unsubscribe_date` | `DateTimeField` | `datetime` | yes |  |  |  |

Indexes:

- index `newsletter_subscription_user_id_4185bd67` on `user_id`
- index `newsletter_subscription_email_3b8020aa` on `email_field`
- index `newsletter_subscription_newsletter_id_e9d08ab9` on `newsletter_id`
- index `newsletter_subscription_subscribed_350402fe` on `subscribed`
- index `newsletter_subscription_unsubscribed_8ae75c15` on `unsubscribed`
- unique `newsletter_subscription_user_id_email_newsletter_id_0a9641e0_uniq` on `user_id`, `email_field`, `newsletter_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `newsletter_id` → [Newsletter](#newsletter).`id`, on delete cascade
- `submission_set` — many to many with [Submission](#submission)

## Article

Table `newsletter_article`, app `newsletter`.

An Article within a Message which will be sent through a Submission.

![Article and the models it relates to](tables/Article.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `sortorder` | `PositiveIntegerField` | `integer unsigned` |  |  |  | Sort order determines the order in which articles are concatenated in a post. |
| `title` | `CharField` | `varchar(200)` |  |  |  |  |
| `text` | `TextField` | `text` |  |  |  |  |
| `url` | `URLField` | `varchar(200)` | yes |  |  |  |
| `image` | `DynamicImageField` | `varchar(100)` | yes |  |  |  |
| `image_thumbnail_width` | `IntegerField` | `integer` | yes |  |  |  |
| `image_use_original` | `BooleanField` | `bool` |  |  | `False` |  |
| `image_below_text` | `BooleanField` | `bool` |  |  | `False` |  |
| `post_id` | `ForeignKey` | `bigint` |  | FK |  |  |

Indexes:

- index `newsletter_article_post_id_c481e545` on `post_id`
- unique `newsletter_article_post_id_sortorder_cbe3c629_uniq` on `post_id`, `sortorder`

Relations:

- `post_id` → [Message](#message).`id`, on delete cascade

## Attachment

Table `newsletter_attachment`, app `newsletter`.

Attachment for a Message.

![Attachment and the models it relates to](tables/Attachment.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `file` | `FileField` | `varchar(100)` |  |  |  |  |
| `message_id` | `ForeignKey` | `bigint` |  | FK |  |  |

Indexes:

- index `newsletter_attachment_message_id_52014d1f` on `message_id`

Relations:

- `message_id` → [Message](#message).`id`, on delete cascade

## Message

Table `newsletter_message`, app `newsletter`.

Message as sent through a Submission.

![Message and the models it relates to](tables/Message.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `title` | `CharField` | `varchar(200)` |  |  |  |  |
| `slug` | `SlugField` | `varchar(50)` |  |  |  |  |
| `newsletter_id` | `ForeignKey` | `bigint` |  | FK | `get_default_newsletter()` |  |
| `date_create` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `date_modify` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `newsletter_message_slug_6dae36a9` on `slug`
- index `newsletter_message_newsletter_id_b77d9df1` on `newsletter_id`
- unique `newsletter_message_slug_newsletter_id_a2ab3653_uniq` on `slug`, `newsletter_id`

Relations:

- [Article](#article).`post_id` → `id`, on delete cascade
- [Attachment](#attachment).`message_id` → `id`, on delete cascade
- `newsletter_id` → [Newsletter](#newsletter).`id`, on delete cascade
- [Submission](#submission).`message_id` → `id`, on delete cascade

## Submission

Table `newsletter_submission`, app `newsletter`.

Submission represents a particular Message as it is being submitted
to a list of Subscribers. This is where actual queueing and submission
happen.

![Submission and the models it relates to](tables/Submission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `site_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `newsletter_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `message_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `publish_date` | `DateTimeField` | `datetime` | yes |  | `now()` |  |
| `publish` | `BooleanField` | `bool` |  |  | `True` | Publish in archive. |
| `prepared` | `BooleanField` | `bool` |  |  | `False` |  |
| `sent` | `BooleanField` | `bool` |  |  | `False` |  |
| `sending` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `newsletter_submission_site_id_f4f389eb` on `site_id`
- index `newsletter_submission_newsletter_id_8fa5ae76` on `newsletter_id`
- index `newsletter_submission_message_id_35227c2e` on `message_id`
- index `newsletter_submission_publish_date_6c3d8801` on `publish_date`
- index `newsletter_submission_publish_2ceb081e` on `publish`
- index `newsletter_submission_prepared_12eea099` on `prepared`
- index `newsletter_submission_sent_7c1e3f56` on `sent`
- index `newsletter_submission_sending_40b87529` on `sending`

Relations:

- `site_id` → [Site](#site).`id`, on delete cascade
- `newsletter_id` → [Newsletter](#newsletter).`id`, on delete cascade
- `message_id` → [Message](#message).`id`, on delete cascade
- `subscriptions` — many to many with [Subscription](#subscription)

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |
