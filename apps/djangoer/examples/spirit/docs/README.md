# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [UserProfile](#userprofile) · [Category](#category) · [Topic](#topic) · [TopicFavorite](#topicfavorite) · [TopicNotification](#topicnotification) · [TopicPrivate](#topicprivate) · [TopicUnread](#topicunread) · [Comment](#comment) · [CommentBookmark](#commentbookmark) · [CommentFlag](#commentflag) · [Flag](#flag) · [CommentHistory](#commenthistory) · [CommentLike](#commentlike) · [CommentPoll](#commentpoll) · [CommentPollChoice](#commentpollchoice) · [CommentPollVote](#commentpollvote) · [Config](#config)

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

CREATE TABLE "spirit_user_userprofile" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL UNIQUE REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "slug" varchar(50) NOT NULL, "nickname" varchar(255) NOT NULL, "location" varchar(75) NOT NULL, "last_seen" datetime NOT NULL, "last_ip" char(39) NULL, "timezone" varchar(32) NOT NULL, "avatar" varchar(255) NOT NULL, "notify" integer unsigned NOT NULL CHECK ("notify" >= 0), "is_administrator" bool NOT NULL, "is_moderator" bool NOT NULL, "is_verified" bool NOT NULL, "topic_count" integer unsigned NOT NULL CHECK ("topic_count" >= 0), "comment_count" integer unsigned NOT NULL CHECK ("comment_count" >= 0), "last_post_hash" varchar(32) NOT NULL, "last_post_on" datetime NULL);

CREATE TABLE "spirit_category_category" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "parent_id" bigint NULL REFERENCES "spirit_category_category" ("id") DEFERRABLE INITIALLY DEFERRED, "title" varchar(75) NOT NULL, "slug" varchar(50) NOT NULL, "description" varchar(255) NOT NULL, "color" varchar(7) NOT NULL, "sort" integer unsigned NOT NULL CHECK ("sort" >= 0), "reindex_at" datetime NOT NULL, "is_global" bool NOT NULL, "is_closed" bool NOT NULL, "is_removed" bool NOT NULL, "is_private" bool NOT NULL);

CREATE TABLE "spirit_topic_topic" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "category_id" bigint NOT NULL REFERENCES "spirit_category_category" ("id") DEFERRABLE INITIALLY DEFERRED, "title" varchar(255) NOT NULL, "slug" varchar(50) NOT NULL, "date" datetime NOT NULL, "last_active" datetime NOT NULL, "reindex_at" datetime NOT NULL, "is_pinned" bool NOT NULL, "is_globally_pinned" bool NOT NULL, "is_closed" bool NOT NULL, "is_removed" bool NOT NULL, "view_count" integer unsigned NOT NULL CHECK ("view_count" >= 0), "comment_count" integer unsigned NOT NULL CHECK ("comment_count" >= 0));

CREATE TABLE "spirit_topic_favorite_topicfavorite" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "topic_id" bigint NOT NULL REFERENCES "spirit_topic_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL);

CREATE TABLE "spirit_topic_notification_topicnotification" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "topic_id" bigint NOT NULL REFERENCES "spirit_topic_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "comment_id" bigint NOT NULL REFERENCES "spirit_comment_comment" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL, "action" integer NOT NULL, "is_read" bool NOT NULL, "is_active" bool NOT NULL);

CREATE TABLE "spirit_topic_private_topicprivate" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "topic_id" bigint NOT NULL REFERENCES "spirit_topic_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL);

CREATE TABLE "spirit_topic_unread_topicunread" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "topic_id" bigint NOT NULL REFERENCES "spirit_topic_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL, "is_read" bool NOT NULL);

CREATE TABLE "spirit_comment_comment" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "topic_id" bigint NOT NULL REFERENCES "spirit_topic_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "comment" text NOT NULL, "comment_html" text NOT NULL, "action" integer NOT NULL, "date" datetime NOT NULL, "is_removed" bool NOT NULL, "is_modified" bool NOT NULL, "ip_address" char(39) NULL, "modified_count" integer unsigned NOT NULL CHECK ("modified_count" >= 0), "likes_count" integer unsigned NOT NULL CHECK ("likes_count" >= 0));

CREATE TABLE "spirit_comment_bookmark_commentbookmark" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "topic_id" bigint NOT NULL REFERENCES "spirit_topic_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "comment_number" integer unsigned NOT NULL CHECK ("comment_number" >= 0));

CREATE TABLE "spirit_comment_flag_commentflag" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "moderator_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "comment_id" bigint NOT NULL UNIQUE REFERENCES "spirit_comment_comment" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL, "is_closed" bool NOT NULL);

CREATE TABLE "spirit_comment_flag_flag" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "comment_id" bigint NOT NULL REFERENCES "spirit_comment_comment" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL, "reason" integer NOT NULL, "body" text NOT NULL);

CREATE TABLE "spirit_comment_history_commenthistory" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "comment_fk_id" bigint NOT NULL REFERENCES "spirit_comment_comment" ("id") DEFERRABLE INITIALLY DEFERRED, "comment_html" text NOT NULL, "date" datetime NOT NULL);

CREATE TABLE "spirit_comment_like_commentlike" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "comment_id" bigint NOT NULL REFERENCES "spirit_comment_comment" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL);

CREATE TABLE "spirit_comment_poll_commentpoll" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "comment_id" bigint NOT NULL REFERENCES "spirit_comment_comment" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(255) NOT NULL, "title" varchar(255) NOT NULL, "choice_min" integer unsigned NOT NULL CHECK ("choice_min" >= 0), "choice_max" integer unsigned NOT NULL CHECK ("choice_max" >= 0), "mode" integer NOT NULL, "close_at" datetime NULL, "is_removed" bool NOT NULL, "created_at" datetime NOT NULL);

CREATE TABLE "spirit_comment_poll_commentpollchoice" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "poll_id" bigint NOT NULL REFERENCES "spirit_comment_poll_commentpoll" ("id") DEFERRABLE INITIALLY DEFERRED, "number" integer unsigned NOT NULL CHECK ("number" >= 0), "description" varchar(255) NOT NULL, "vote_count" integer unsigned NOT NULL CHECK ("vote_count" >= 0), "is_removed" bool NOT NULL);

CREATE TABLE "spirit_comment_poll_commentpollvote" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "voter_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "choice_id" bigint NOT NULL REFERENCES "spirit_comment_poll_commentpollchoice" ("id") DEFERRABLE INITIALLY DEFERRED, "is_removed" bool NOT NULL, "created_at" datetime NOT NULL);

CREATE TABLE "djconfig_config" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "key" varchar(75) NOT NULL UNIQUE, "value" text NULL);

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

CREATE INDEX "spirit_category_category_parent_id_e59f3015" ON "spirit_category_category" ("parent_id");

CREATE INDEX "spirit_topic_topic_user_id_5d0151a9" ON "spirit_topic_topic" ("user_id");

CREATE INDEX "spirit_topic_topic_category_id_824dc7e8" ON "spirit_topic_topic" ("category_id");

CREATE UNIQUE INDEX "spirit_topic_favorite_topicfavorite_user_id_topic_id_167416f4_uniq" ON "spirit_topic_favorite_topicfavorite" ("user_id", "topic_id");

CREATE INDEX "spirit_topic_favorite_topicfavorite_user_id_e406cc12" ON "spirit_topic_favorite_topicfavorite" ("user_id");

CREATE INDEX "spirit_topic_favorite_topicfavorite_topic_id_7a363899" ON "spirit_topic_favorite_topicfavorite" ("topic_id");

CREATE UNIQUE INDEX "spirit_topic_notification_topicnotification_user_id_topic_id_1f927541_uniq" ON "spirit_topic_notification_topicnotification" ("user_id", "topic_id");

CREATE INDEX "spirit_topic_notification_topicnotification_user_id_a9d2003a" ON "spirit_topic_notification_topicnotification" ("user_id");

CREATE INDEX "spirit_topic_notification_topicnotification_topic_id_449cdcfe" ON "spirit_topic_notification_topicnotification" ("topic_id");

CREATE INDEX "spirit_topic_notification_topicnotification_comment_id_703f40ab" ON "spirit_topic_notification_topicnotification" ("comment_id");

CREATE UNIQUE INDEX "spirit_topic_private_topicprivate_user_id_topic_id_8c297113_uniq" ON "spirit_topic_private_topicprivate" ("user_id", "topic_id");

CREATE INDEX "spirit_topic_private_topicprivate_user_id_682c6ef5" ON "spirit_topic_private_topicprivate" ("user_id");

CREATE INDEX "spirit_topic_private_topicprivate_topic_id_47c6f3fd" ON "spirit_topic_private_topicprivate" ("topic_id");

CREATE UNIQUE INDEX "spirit_topic_unread_topicunread_user_id_topic_id_8f3c8a54_uniq" ON "spirit_topic_unread_topicunread" ("user_id", "topic_id");

CREATE INDEX "spirit_topic_unread_topicunread_user_id_f5a459f5" ON "spirit_topic_unread_topicunread" ("user_id");

CREATE INDEX "spirit_topic_unread_topicunread_topic_id_04fc7131" ON "spirit_topic_unread_topicunread" ("topic_id");

CREATE INDEX "spirit_comment_comment_user_id_0f9dbd41" ON "spirit_comment_comment" ("user_id");

CREATE INDEX "spirit_comment_comment_topic_id_4f5636aa" ON "spirit_comment_comment" ("topic_id");

CREATE UNIQUE INDEX "spirit_comment_bookmark_commentbookmark_user_id_topic_id_5ad0da74_uniq" ON "spirit_comment_bookmark_commentbookmark" ("user_id", "topic_id");

CREATE INDEX "spirit_comment_bookmark_commentbookmark_user_id_b31a0d88" ON "spirit_comment_bookmark_commentbookmark" ("user_id");

CREATE INDEX "spirit_comment_bookmark_commentbookmark_topic_id_cd606f44" ON "spirit_comment_bookmark_commentbookmark" ("topic_id");

CREATE INDEX "spirit_comment_flag_commentflag_moderator_id_1eb69535" ON "spirit_comment_flag_commentflag" ("moderator_id");

CREATE UNIQUE INDEX "spirit_comment_flag_flag_user_id_comment_id_900c59f8_uniq" ON "spirit_comment_flag_flag" ("user_id", "comment_id");

CREATE INDEX "spirit_comment_flag_flag_user_id_ffbf846e" ON "spirit_comment_flag_flag" ("user_id");

CREATE INDEX "spirit_comment_flag_flag_comment_id_1223e22e" ON "spirit_comment_flag_flag" ("comment_id");

CREATE INDEX "spirit_comment_history_commenthistory_comment_fk_id_48f6b502" ON "spirit_comment_history_commenthistory" ("comment_fk_id");

CREATE UNIQUE INDEX "spirit_comment_like_commentlike_user_id_comment_id_32b1f85f_uniq" ON "spirit_comment_like_commentlike" ("user_id", "comment_id");

CREATE INDEX "spirit_comment_like_commentlike_user_id_85e21014" ON "spirit_comment_like_commentlike" ("user_id");

CREATE INDEX "spirit_comment_like_commentlike_comment_id_a81f0f4f" ON "spirit_comment_like_commentlike" ("comment_id");

CREATE UNIQUE INDEX "spirit_comment_poll_commentpoll_comment_id_name_b54ee48a_uniq" ON "spirit_comment_poll_commentpoll" ("comment_id", "name");

CREATE INDEX "spirit_comment_poll_commentpoll_comment_id_453c04b5" ON "spirit_comment_poll_commentpoll" ("comment_id");

CREATE UNIQUE INDEX "spirit_comment_poll_commentpollchoice_poll_id_number_36ccebb7_uniq" ON "spirit_comment_poll_commentpollchoice" ("poll_id", "number");

CREATE INDEX "spirit_comment_poll_commentpollchoice_poll_id_01dd8cb4" ON "spirit_comment_poll_commentpollchoice" ("poll_id");

CREATE UNIQUE INDEX "spirit_comment_poll_commentpollvote_voter_id_choice_id_b2b15b1f_uniq" ON "spirit_comment_poll_commentpollvote" ("voter_id", "choice_id");

CREATE INDEX "spirit_comment_poll_commentpollvote_voter_id_7fd9a8eb" ON "spirit_comment_poll_commentpollvote" ("voter_id");

CREATE INDEX "spirit_comment_poll_commentpollvote_choice_id_1ecb02ca" ON "spirit_comment_poll_commentpollvote" ("choice_id");
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
- [UserProfile](#userprofile).`user_id` → `id`, on delete cascade
- [Topic](#topic).`user_id` → `id`, on delete cascade
- [TopicFavorite](#topicfavorite).`user_id` → `id`, on delete cascade
- [TopicNotification](#topicnotification).`user_id` → `id`, on delete cascade
- [TopicPrivate](#topicprivate).`user_id` → `id`, on delete cascade
- [TopicUnread](#topicunread).`user_id` → `id`, on delete cascade
- [Comment](#comment).`user_id` → `id`, on delete cascade
- [CommentBookmark](#commentbookmark).`user_id` → `id`, on delete cascade
- [CommentFlag](#commentflag).`moderator_id` → `id`, on delete cascade
- [Flag](#flag).`user_id` → `id`, on delete cascade
- [CommentLike](#commentlike).`user_id` → `id`, on delete cascade
- [CommentPollVote](#commentpollvote).`voter_id` → `id`, on delete cascade

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

## UserProfile

Table `spirit_user_userprofile`, app `spirit_user`.

![UserProfile and the models it relates to](tables/UserProfile.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `OneToOneField` | `integer` |  | FK, unique |  |  |
| `slug` | `AutoSlugField` | `varchar(50)` |  |  |  |  |
| `nickname` | `CharField` | `varchar(255)` |  |  |  |  |
| `location` | `CharField` | `varchar(75)` |  |  |  |  |
| `last_seen` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `last_ip` | `GenericIPAddressField` | `char(39)` | yes |  |  |  |
| `timezone` | `CharField` | `varchar(32)` |  |  | `'UTC'` |  |
| `avatar` | `ImageField` | `varchar(255)` |  |  |  |  |
| `notify` | `PositiveIntegerField` | `integer unsigned` |  |  | `25` |  |
| `is_administrator` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_moderator` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_verified` | `BooleanField` | `bool` |  |  | `False` | Designates whether the user has verified his account by email or by other means. Un-select this to let the user activate his account. |
| `topic_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `comment_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `last_post_hash` | `CharField` | `varchar(32)` |  |  |  |  |
| `last_post_on` | `DateTimeField` | `datetime` | yes |  |  |  |

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

## Category

Table `spirit_category_category`, app `spirit_category`.

Category model

:ivar reindex_at: Last time this model was marked    for reindex. It makes the search re-index the topic,    it must be set explicitly
:vartype reindex_at: `:py:class:models.DateTimeField`

![Category and the models it relates to](tables/Category.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `parent_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `title` | `CharField` | `varchar(75)` |  |  |  |  |
| `slug` | `AutoSlugField` | `varchar(50)` |  |  |  |  |
| `description` | `CharField` | `varchar(255)` |  |  |  |  |
| `color` | `CharField` | `varchar(7)` |  |  |  | Title color in hex format (i.e: #1aafd0). |
| `sort` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `reindex_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `is_global` | `BooleanField` | `bool` |  |  | `True` | Designates whether the topics will bedisplayed in the all-categories list. |
| `is_closed` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_removed` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_private` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `spirit_category_category_parent_id_e59f3015` on `parent_id`

Relations:

- `parent_id` → [Category](#category).`id`, on delete cascade
- [Topic](#topic).`category_id` → `id`, on delete cascade

## Topic

Table `spirit_topic_topic`, app `spirit_topic`.

Topic model

:ivar last_active: Last time a comment was added/removed,    it makes the search re-index the topic
:vartype last_active: `:py:class:models.DateTimeField`
:ivar reindex_at: Last time this model was marked    for reindex. It makes the search re-index the topic,    it must be set explicitly
:vartype reindex_at: `:py:class:models.DateTimeField`

![Topic and the models it relates to](tables/Topic.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `category_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `title` | `CharField` | `varchar(255)` |  |  |  |  |
| `slug` | `AutoSlugField` | `varchar(50)` |  |  |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `last_active` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `reindex_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `is_pinned` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_globally_pinned` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_closed` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_removed` | `BooleanField` | `bool` |  |  | `False` |  |
| `view_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `comment_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |

Indexes:

- index `spirit_topic_topic_user_id_5d0151a9` on `user_id`
- index `spirit_topic_topic_category_id_824dc7e8` on `category_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `category_id` → [Category](#category).`id`, on delete cascade
- [TopicFavorite](#topicfavorite).`topic_id` → `id`, on delete cascade
- [TopicNotification](#topicnotification).`topic_id` → `id`, on delete cascade
- [TopicPrivate](#topicprivate).`topic_id` → `id`, on delete cascade
- [TopicUnread](#topicunread).`topic_id` → `id`, on delete cascade
- [Comment](#comment).`topic_id` → `id`, on delete cascade
- [CommentBookmark](#commentbookmark).`topic_id` → `id`, on delete cascade

## TopicFavorite

Table `spirit_topic_favorite_topicfavorite`, app `spirit_topic_favorite`.

![TopicFavorite and the models it relates to](tables/TopicFavorite.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `topic_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `spirit_topic_favorite_topicfavorite_user_id_e406cc12` on `user_id`
- index `spirit_topic_favorite_topicfavorite_topic_id_7a363899` on `topic_id`
- unique `spirit_topic_favorite_topicfavorite_user_id_topic_id_167416f4_uniq` on `user_id`, `topic_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `topic_id` → [Topic](#topic).`id`, on delete cascade

## TopicNotification

Table `spirit_topic_notification_topicnotification`, app `spirit_topic_notification`.

![TopicNotification and the models it relates to](tables/TopicNotification.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `topic_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `comment_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `action` | `TopicNotificationAction` | `integer` |  |  | `0` |  |
| `is_read` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_active` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `spirit_topic_notification_topicnotification_user_id_a9d2003a` on `user_id`
- index `spirit_topic_notification_topicnotification_topic_id_449cdcfe` on `topic_id`
- index `spirit_topic_notification_topicnotification_comment_id_703f40ab` on `comment_id`
- unique `spirit_topic_notification_topicnotification_user_id_topic_id_1f927541_uniq` on `user_id`, `topic_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `topic_id` → [Topic](#topic).`id`, on delete cascade
- `comment_id` → [Comment](#comment).`id`, on delete cascade

## TopicPrivate

Table `spirit_topic_private_topicprivate`, app `spirit_topic_private`.

![TopicPrivate and the models it relates to](tables/TopicPrivate.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `topic_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `spirit_topic_private_topicprivate_user_id_682c6ef5` on `user_id`
- index `spirit_topic_private_topicprivate_topic_id_47c6f3fd` on `topic_id`
- unique `spirit_topic_private_topicprivate_user_id_topic_id_8c297113_uniq` on `user_id`, `topic_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `topic_id` → [Topic](#topic).`id`, on delete cascade

## TopicUnread

Table `spirit_topic_unread_topicunread`, app `spirit_topic_unread`.

![TopicUnread and the models it relates to](tables/TopicUnread.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `topic_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `is_read` | `BooleanField` | `bool` |  |  | `True` |  |

Indexes:

- index `spirit_topic_unread_topicunread_user_id_f5a459f5` on `user_id`
- index `spirit_topic_unread_topicunread_topic_id_04fc7131` on `topic_id`
- unique `spirit_topic_unread_topicunread_user_id_topic_id_8f3c8a54_uniq` on `user_id`, `topic_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `topic_id` → [Topic](#topic).`id`, on delete cascade

## Comment

Table `spirit_comment_comment`, app `spirit_comment`.

![Comment and the models it relates to](tables/Comment.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `topic_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `comment` | `TextField` | `text` |  |  |  |  |
| `comment_html` | `TextField` | `text` |  |  |  |  |
| `action` | `CommentAction` | `integer` |  |  | `0` |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `is_removed` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_modified` | `BooleanField` | `bool` |  |  | `False` |  |
| `ip_address` | `GenericIPAddressField` | `char(39)` | yes |  |  |  |
| `modified_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `likes_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |

Indexes:

- index `spirit_comment_comment_user_id_0f9dbd41` on `user_id`
- index `spirit_comment_comment_topic_id_4f5636aa` on `topic_id`

Relations:

- [TopicNotification](#topicnotification).`comment_id` → `id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade
- `topic_id` → [Topic](#topic).`id`, on delete cascade
- [CommentFlag](#commentflag).`comment_id` → `id`, on delete cascade
- [Flag](#flag).`comment_id` → `id`, on delete cascade
- [CommentHistory](#commenthistory).`comment_fk_id` → `id`, on delete cascade
- [CommentLike](#commentlike).`comment_id` → `id`, on delete cascade
- [CommentPoll](#commentpoll).`comment_id` → `id`, on delete cascade

## CommentBookmark

Table `spirit_comment_bookmark_commentbookmark`, app `spirit_comment_bookmark`.

![CommentBookmark and the models it relates to](tables/CommentBookmark.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `topic_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `comment_number` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |

Indexes:

- index `spirit_comment_bookmark_commentbookmark_user_id_b31a0d88` on `user_id`
- index `spirit_comment_bookmark_commentbookmark_topic_id_cd606f44` on `topic_id`
- unique `spirit_comment_bookmark_commentbookmark_user_id_topic_id_5ad0da74_uniq` on `user_id`, `topic_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `topic_id` → [Topic](#topic).`id`, on delete cascade

## CommentFlag

Table `spirit_comment_flag_commentflag`, app `spirit_comment_flag`.

![CommentFlag and the models it relates to](tables/CommentFlag.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `moderator_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `comment_id` | `OneToOneField` | `bigint` |  | FK, unique |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `is_closed` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `spirit_comment_flag_commentflag_moderator_id_1eb69535` on `moderator_id`

Relations:

- `moderator_id` → [User](#user).`id`, on delete cascade
- `comment_id` → [Comment](#comment).`id`, on delete cascade

## Flag

Table `spirit_comment_flag_flag`, app `spirit_comment_flag`.

![Flag and the models it relates to](tables/Flag.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `comment_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `reason` | `FlagReason` | `integer` |  |  |  |  |
| `body` | `TextField` | `text` |  |  |  |  |

Indexes:

- index `spirit_comment_flag_flag_user_id_ffbf846e` on `user_id`
- index `spirit_comment_flag_flag_comment_id_1223e22e` on `comment_id`
- unique `spirit_comment_flag_flag_user_id_comment_id_900c59f8_uniq` on `user_id`, `comment_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `comment_id` → [Comment](#comment).`id`, on delete cascade

## CommentHistory

Table `spirit_comment_history_commenthistory`, app `spirit_comment_history`.

![CommentHistory and the models it relates to](tables/CommentHistory.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `comment_fk_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `comment_html` | `TextField` | `text` |  |  |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `spirit_comment_history_commenthistory_comment_fk_id_48f6b502` on `comment_fk_id`

Relations:

- `comment_fk_id` → [Comment](#comment).`id`, on delete cascade

## CommentLike

Table `spirit_comment_like_commentlike`, app `spirit_comment_like`.

![CommentLike and the models it relates to](tables/CommentLike.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `comment_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `spirit_comment_like_commentlike_user_id_85e21014` on `user_id`
- index `spirit_comment_like_commentlike_comment_id_a81f0f4f` on `comment_id`
- unique `spirit_comment_like_commentlike_user_id_comment_id_32b1f85f_uniq` on `user_id`, `comment_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `comment_id` → [Comment](#comment).`id`, on delete cascade

## CommentPoll

Table `spirit_comment_poll_commentpoll`, app `spirit_comment_poll`.

![CommentPoll and the models it relates to](tables/CommentPoll.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `comment_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `title` | `CharField` | `varchar(255)` |  |  |  |  |
| `choice_min` | `PositiveIntegerField` | `integer unsigned` |  |  | `1` |  |
| `choice_max` | `PositiveIntegerField` | `integer unsigned` |  |  | `1` |  |
| `mode` | `CommentPollMode` | `integer` |  |  | `0` |  |
| `close_at` | `DateTimeField` | `datetime` | yes |  |  |  |
| `is_removed` | `BooleanField` | `bool` |  |  | `False` |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `spirit_comment_poll_commentpoll_comment_id_453c04b5` on `comment_id`
- unique `spirit_comment_poll_commentpoll_comment_id_name_b54ee48a_uniq` on `comment_id`, `name`

Relations:

- `comment_id` → [Comment](#comment).`id`, on delete cascade
- [CommentPollChoice](#commentpollchoice).`poll_id` → `id`, on delete cascade

## CommentPollChoice

Table `spirit_comment_poll_commentpollchoice`, app `spirit_comment_poll`.

![CommentPollChoice and the models it relates to](tables/CommentPollChoice.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `poll_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `number` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `description` | `CharField` | `varchar(255)` |  |  |  |  |
| `vote_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `is_removed` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `spirit_comment_poll_commentpollchoice_poll_id_01dd8cb4` on `poll_id`
- unique `spirit_comment_poll_commentpollchoice_poll_id_number_36ccebb7_uniq` on `poll_id`, `number`

Relations:

- `poll_id` → [CommentPoll](#commentpoll).`id`, on delete cascade
- [CommentPollVote](#commentpollvote).`choice_id` → `id`, on delete cascade

## CommentPollVote

Table `spirit_comment_poll_commentpollvote`, app `spirit_comment_poll`.

![CommentPollVote and the models it relates to](tables/CommentPollVote.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `voter_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `choice_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `is_removed` | `BooleanField` | `bool` |  |  | `False` |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `spirit_comment_poll_commentpollvote_voter_id_7fd9a8eb` on `voter_id`
- index `spirit_comment_poll_commentpollvote_choice_id_1ecb02ca` on `choice_id`
- unique `spirit_comment_poll_commentpollvote_voter_id_choice_id_b2b15b1f_uniq` on `voter_id`, `choice_id`

Relations:

- `voter_id` → [User](#user).`id`, on delete cascade
- `choice_id` → [CommentPollChoice](#commentpollchoice).`id`, on delete cascade

## Config

Table `djconfig_config`, app `djconfig`.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `key` | `CharField` | `varchar(75)` |  | unique |  |  |
| `value` | `TextField` | `text` | yes |  |  |  |

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## TopicNotificationAction

| Name | Stored as |
| --- | --- |
| `Undefined` | `0` |
| `Mention` | `1` |
| `Comment` | `2` |

## CommentAction

| Name | Stored as |
| --- | --- |
| `comment` | `0` |
| `topic moved` | `1` |
| `topic closed` | `2` |
| `topic unclosed` | `3` |
| `topic pinned` | `4` |
| `topic unpinned` | `5` |

## FlagReason

| Name | Stored as |
| --- | --- |
| `Spam` | `0` |
| `Other` | `1` |

## CommentPollMode

| Name | Stored as |
| --- | --- |
| `default` | `0` |
| `secret` | `1` |
