# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Forum](#forum) · [Topic](#topic) · [Post](#post) · [Attachment](#attachment) · [TopicPoll](#topicpoll) · [TopicPollOption](#topicpolloption) · [TopicPollVote](#topicpollvote) · [ForumReadTrack](#forumreadtrack) · [TopicReadTrack](#topicreadtrack) · [ForumProfile](#forumprofile) · [ForumPermission](#forumpermission) · [GroupForumPermission](#groupforumpermission) · [UserForumPermission](#userforumpermission)

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

CREATE TABLE "forum_forum" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "created" datetime NOT NULL, "updated" datetime NOT NULL, "parent_id" bigint NULL REFERENCES "forum_forum" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(100) NOT NULL, "slug" varchar(255) NOT NULL, "description" text NULL, "image" varchar(100) NULL, "link" varchar(200) NULL, "link_redirects" bool NOT NULL, "type" smallint unsigned NOT NULL CHECK ("type" >= 0), "direct_posts_count" integer unsigned NOT NULL CHECK ("direct_posts_count" >= 0), "direct_topics_count" integer unsigned NOT NULL CHECK ("direct_topics_count" >= 0), "link_redirects_count" integer unsigned NOT NULL CHECK ("link_redirects_count" >= 0), "last_post_id" bigint NULL REFERENCES "forum_conversation_post" ("id") DEFERRABLE INITIALLY DEFERRED, "last_post_on" datetime NULL, "display_sub_forum_list" bool NOT NULL, "_description_rendered" text NULL, "lft" integer unsigned NOT NULL CHECK ("lft" >= 0), "rght" integer unsigned NOT NULL CHECK ("rght" >= 0), "tree_id" integer unsigned NOT NULL CHECK ("tree_id" >= 0), "level" integer unsigned NOT NULL CHECK ("level" >= 0));

CREATE TABLE "forum_conversation_topic" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "forum_id" bigint NOT NULL REFERENCES "forum_forum" ("id") DEFERRABLE INITIALLY DEFERRED, "poster_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "subject" varchar(255) NOT NULL, "slug" varchar(255) NOT NULL, "type" smallint unsigned NOT NULL CHECK ("type" >= 0), "status" integer unsigned NOT NULL CHECK ("status" >= 0), "approved" bool NOT NULL, "posts_count" integer unsigned NOT NULL CHECK ("posts_count" >= 0), "views_count" integer unsigned NOT NULL CHECK ("views_count" >= 0), "last_post_on" datetime NULL, "first_post_id" bigint NULL REFERENCES "forum_conversation_post" ("id") DEFERRABLE INITIALLY DEFERRED, "last_post_id" bigint NULL REFERENCES "forum_conversation_post" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL, "updated" datetime NOT NULL);

CREATE TABLE "forum_conversation_topic_subscribers" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "topic_id" bigint NOT NULL REFERENCES "forum_conversation_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "forum_conversation_post" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "created" datetime NOT NULL, "updated" datetime NOT NULL, "topic_id" bigint NOT NULL REFERENCES "forum_conversation_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "poster_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "anonymous_key" varchar(100) NULL, "subject" varchar(255) NOT NULL, "content" text NOT NULL, "username" varchar(155) NULL, "approved" bool NOT NULL, "enable_signature" bool NOT NULL, "update_reason" varchar(255) NULL, "updated_by_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "updates_count" integer unsigned NOT NULL CHECK ("updates_count" >= 0), "_content_rendered" text NULL);

CREATE TABLE "forum_attachments_attachment" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "post_id" bigint NOT NULL REFERENCES "forum_conversation_post" ("id") DEFERRABLE INITIALLY DEFERRED, "file" varchar(100) NOT NULL, "comment" varchar(255) NULL);

CREATE TABLE "forum_polls_topicpoll" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "created" datetime NOT NULL, "updated" datetime NOT NULL, "topic_id" bigint NOT NULL UNIQUE REFERENCES "forum_conversation_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "question" varchar(255) NOT NULL, "duration" integer unsigned NULL CHECK ("duration" >= 0), "max_options" smallint unsigned NOT NULL CHECK ("max_options" >= 0), "user_changes" bool NOT NULL, "hide_results" bool NOT NULL);

CREATE TABLE "forum_polls_topicpolloption" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "poll_id" bigint NOT NULL REFERENCES "forum_polls_topicpoll" ("id") DEFERRABLE INITIALLY DEFERRED, "text" varchar(255) NOT NULL);

CREATE TABLE "forum_polls_topicpollvote" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "voter_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "anonymous_key" varchar(100) NULL, "poll_option_id" bigint NOT NULL REFERENCES "forum_polls_topicpolloption" ("id") DEFERRABLE INITIALLY DEFERRED, "timestamp" datetime NOT NULL);

CREATE TABLE "forum_tracking_forumreadtrack" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "forum_id" bigint NOT NULL REFERENCES "forum_forum" ("id") DEFERRABLE INITIALLY DEFERRED, "mark_time" datetime NOT NULL);

CREATE TABLE "forum_tracking_topicreadtrack" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "topic_id" bigint NOT NULL REFERENCES "forum_conversation_topic" ("id") DEFERRABLE INITIALLY DEFERRED, "mark_time" datetime NOT NULL);

CREATE TABLE "forum_member_forumprofile" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL UNIQUE REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "avatar" varchar(100) NULL, "signature" text NULL, "posts_count" integer unsigned NOT NULL CHECK ("posts_count" >= 0), "_signature_rendered" text NULL);

CREATE TABLE "forum_permission_forumpermission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "codename" varchar(150) NOT NULL UNIQUE);

CREATE TABLE "forum_permission_groupforumpermission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "permission_id" bigint NOT NULL REFERENCES "forum_permission_forumpermission" ("id") DEFERRABLE INITIALLY DEFERRED, "has_perm" bool NOT NULL, "forum_id" bigint NULL REFERENCES "forum_forum" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "forum_permission_userforumpermission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "permission_id" bigint NOT NULL REFERENCES "forum_permission_forumpermission" ("id") DEFERRABLE INITIALLY DEFERRED, "has_perm" bool NOT NULL, "forum_id" bigint NULL REFERENCES "forum_forum" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "anonymous_user" bool NOT NULL, "authenticated_user" bool NOT NULL);

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

CREATE INDEX "forum_forum_parent_id_22edea05" ON "forum_forum" ("parent_id");

CREATE INDEX "forum_forum_slug_b9acc50d" ON "forum_forum" ("slug");

CREATE INDEX "forum_forum_type_30239563" ON "forum_forum" ("type");

CREATE INDEX "forum_forum_last_post_id_81b59e37" ON "forum_forum" ("last_post_id");

CREATE INDEX "forum_forum_tree_id_84a9227d" ON "forum_forum" ("tree_id");

CREATE INDEX "forum_forum_tree_id_lft_idx" ON "forum_forum" ("tree_id", "lft");

CREATE INDEX "forum_conversation_topic_forum_id_e9cfe592" ON "forum_conversation_topic" ("forum_id");

CREATE INDEX "forum_conversation_topic_poster_id_0dd4fa07" ON "forum_conversation_topic" ("poster_id");

CREATE INDEX "forum_conversation_topic_slug_c74ce2cc" ON "forum_conversation_topic" ("slug");

CREATE INDEX "forum_conversation_topic_type_92971eb5" ON "forum_conversation_topic" ("type");

CREATE INDEX "forum_conversation_topic_status_e78d0ae4" ON "forum_conversation_topic" ("status");

CREATE INDEX "forum_conversation_topic_approved_ad3fcbf9" ON "forum_conversation_topic" ("approved");

CREATE INDEX "forum_conversation_topic_last_post_on_a287686b" ON "forum_conversation_topic" ("last_post_on");

CREATE INDEX "forum_conversation_topic_first_post_id_ca473bd1" ON "forum_conversation_topic" ("first_post_id");

CREATE INDEX "forum_conversation_topic_last_post_id_e14396a2" ON "forum_conversation_topic" ("last_post_id");

CREATE INDEX "forum_conversation_topic_created_4e06996f" ON "forum_conversation_topic" ("created");

CREATE INDEX "forum_conversation_topic_updated_3a9e00a1" ON "forum_conversation_topic" ("updated");

CREATE INDEX "forum_conve_type_cc96d0_idx" ON "forum_conversation_topic" ("type", "last_post_on");

CREATE UNIQUE INDEX "forum_conversation_topic_subscribers_topic_id_user_id_b2c961d5_uniq" ON "forum_conversation_topic_subscribers" ("topic_id", "user_id");

CREATE INDEX "forum_conversation_topic_subscribers_topic_id_34ebca87" ON "forum_conversation_topic_subscribers" ("topic_id");

CREATE INDEX "forum_conversation_topic_subscribers_user_id_7e386a79" ON "forum_conversation_topic_subscribers" ("user_id");

CREATE INDEX "forum_conversation_post_topic_id_f6aaa418" ON "forum_conversation_post" ("topic_id");

CREATE INDEX "forum_conversation_post_poster_id_19c4e995" ON "forum_conversation_post" ("poster_id");

CREATE INDEX "forum_conversation_post_approved_a1090910" ON "forum_conversation_post" ("approved");

CREATE INDEX "forum_conversation_post_enable_signature_b1ce8b55" ON "forum_conversation_post" ("enable_signature");

CREATE INDEX "forum_conversation_post_updated_by_id_86093355" ON "forum_conversation_post" ("updated_by_id");

CREATE INDEX "forum_attachments_attachment_post_id_0476a843" ON "forum_attachments_attachment" ("post_id");

CREATE INDEX "forum_polls_topicpolloption_poll_id_a54cbd58" ON "forum_polls_topicpolloption" ("poll_id");

CREATE INDEX "forum_polls_topicpollvote_voter_id_0a287559" ON "forum_polls_topicpollvote" ("voter_id");

CREATE INDEX "forum_polls_topicpollvote_poll_option_id_a075b665" ON "forum_polls_topicpollvote" ("poll_option_id");

CREATE UNIQUE INDEX "forum_tracking_forumreadtrack_user_id_forum_id_3e64777a_uniq" ON "forum_tracking_forumreadtrack" ("user_id", "forum_id");

CREATE INDEX "forum_tracking_forumreadtrack_user_id_932d4605" ON "forum_tracking_forumreadtrack" ("user_id");

CREATE INDEX "forum_tracking_forumreadtrack_forum_id_bbd3fb47" ON "forum_tracking_forumreadtrack" ("forum_id");

CREATE INDEX "forum_tracking_forumreadtrack_mark_time_72eec39e" ON "forum_tracking_forumreadtrack" ("mark_time");

CREATE UNIQUE INDEX "forum_tracking_topicreadtrack_user_id_topic_id_6fe3e105_uniq" ON "forum_tracking_topicreadtrack" ("user_id", "topic_id");

CREATE INDEX "forum_tracking_topicreadtrack_user_id_2762562b" ON "forum_tracking_topicreadtrack" ("user_id");

CREATE INDEX "forum_tracking_topicreadtrack_topic_id_9a53bd45" ON "forum_tracking_topicreadtrack" ("topic_id");

CREATE INDEX "forum_tracking_topicreadtrack_mark_time_7dafc483" ON "forum_tracking_topicreadtrack" ("mark_time");

CREATE UNIQUE INDEX "forum_permission_groupforumpermission_permission_id_forum_id_group_id_a1e477c8_uniq" ON "forum_permission_groupforumpermission" ("permission_id", "forum_id", "group_id");

CREATE INDEX "forum_permission_groupforumpermission_permission_id_2475fe70" ON "forum_permission_groupforumpermission" ("permission_id");

CREATE INDEX "forum_permission_groupforumpermission_has_perm_48cae01d" ON "forum_permission_groupforumpermission" ("has_perm");

CREATE INDEX "forum_permission_groupforumpermission_forum_id_d59d5cac" ON "forum_permission_groupforumpermission" ("forum_id");

CREATE INDEX "forum_permission_groupforumpermission_group_id_b515635b" ON "forum_permission_groupforumpermission" ("group_id");

CREATE UNIQUE INDEX "forum_permission_userforumpermission_permission_id_forum_id_user_id_024a3693_uniq" ON "forum_permission_userforumpermission" ("permission_id", "forum_id", "user_id");

CREATE INDEX "forum_permission_userforumpermission_permission_id_9090e930" ON "forum_permission_userforumpermission" ("permission_id");

CREATE INDEX "forum_permission_userforumpermission_has_perm_1b5ee7ac" ON "forum_permission_userforumpermission" ("has_perm");

CREATE INDEX "forum_permission_userforumpermission_forum_id_f781f4d6" ON "forum_permission_userforumpermission" ("forum_id");

CREATE INDEX "forum_permission_userforumpermission_user_id_8106d02d" ON "forum_permission_userforumpermission" ("user_id");

CREATE INDEX "forum_permission_userforumpermission_anonymous_user_8aabbd9d" ON "forum_permission_userforumpermission" ("anonymous_user");

CREATE INDEX "forum_permission_userforumpermission_authenticated_user_b677afa4" ON "forum_permission_userforumpermission" ("authenticated_user");
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
- [GroupForumPermission](#groupforumpermission).`group_id` → `id`, on delete cascade

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
- [Topic](#topic).`poster_id` → `id`, on delete cascade
- `topic_subscriptions` — many to many with [Topic](#topic)
- [Post](#post).`poster_id` → `id`, on delete cascade
- [Post](#post).`updated_by_id` → `id`, on delete set null
- [TopicPollVote](#topicpollvote).`voter_id` → `id`, on delete cascade
- [ForumReadTrack](#forumreadtrack).`user_id` → `id`, on delete cascade
- [TopicReadTrack](#topicreadtrack).`user_id` → `id`, on delete cascade
- [ForumProfile](#forumprofile).`user_id` → `id`, on delete cascade
- [UserForumPermission](#userforumpermission).`user_id` → `id`, on delete cascade

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

## Forum

Table `forum_forum`, app `forum`.

Columns from `AbstractForum` (abstract): `parent_id`, `name`, `slug`, `description`, `image`, `link`, `link_redirects`, `type`, `direct_posts_count`, `direct_topics_count`, `link_redirects_count`, `last_post_id`, `last_post_on`, `display_sub_forum_list`.

Columns from `DatedModel` (abstract): `created`, `updated`.

![Forum and the models it relates to](tables/Forum.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `parent_id` | `TreeForeignKey` | `bigint` | yes | FK |  |  |
| `name` | `CharField` | `varchar(100)` |  |  |  |  |
| `slug` | `SlugField` | `varchar(255)` |  |  |  |  |
| `description` | `MarkupTextField` | `text` | yes |  |  |  |
| `image` | `ExtendedImageField` | `varchar(100)` | yes |  |  |  |
| `link` | `URLField` | `varchar(200)` | yes |  |  |  |
| `link_redirects` | `BooleanField` | `bool` |  |  | `False` | Records the number of times a forum link was clicked |
| `type` | `ForumType` | `smallint unsigned` |  |  |  |  |
| `direct_posts_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `direct_topics_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `link_redirects_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `last_post_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `last_post_on` | `DateTimeField` | `datetime` | yes |  |  |  |
| `display_sub_forum_list` | `BooleanField` | `bool` |  |  | `True` | Displays this forum on the legend of its parent-forum (sub forums list) |
| `_description_rendered` | `TextField` | `text` | yes |  |  |  |
| `lft` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `rght` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `tree_id` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `level` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |

Indexes:

- index `forum_forum_parent_id_22edea05` on `parent_id`
- index `forum_forum_slug_b9acc50d` on `slug`
- index `forum_forum_type_30239563` on `type`
- index `forum_forum_last_post_id_81b59e37` on `last_post_id`
- index `forum_forum_tree_id_84a9227d` on `tree_id`
- index `forum_forum_tree_id_lft_idx` on `tree_id`, `lft`

Relations:

- `parent_id` → [Forum](#forum).`id`, on delete cascade
- `last_post_id` → [Post](#post).`id`, on delete set null
- [Topic](#topic).`forum_id` → `id`, on delete cascade
- [ForumReadTrack](#forumreadtrack).`forum_id` → `id`, on delete cascade
- [GroupForumPermission](#groupforumpermission).`forum_id` → `id`, on delete cascade
- [UserForumPermission](#userforumpermission).`forum_id` → `id`, on delete cascade

## Topic

Table `forum_conversation_topic`, app `forum_conversation`.

Columns from `AbstractTopic` (abstract): `forum_id`, `poster_id`, `subject`, `slug`, `type`, `status`, `approved`, `posts_count`, `views_count`, `last_post_on`, `first_post_id`, `last_post_id`, `created`, `updated`.

![Topic and the models it relates to](tables/Topic.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `forum_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `poster_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `subject` | `CharField` | `varchar(255)` |  |  |  |  |
| `slug` | `SlugField` | `varchar(255)` |  |  |  |  |
| `type` | `TopicType` | `smallint unsigned` |  |  |  |  |
| `status` | `TopicStatus` | `integer unsigned` |  |  |  |  |
| `approved` | `BooleanField` | `bool` |  |  | `True` |  |
| `posts_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `views_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `last_post_on` | `DateTimeField` | `datetime` | yes |  |  |  |
| `first_post_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `last_post_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `forum_conversation_topic_forum_id_e9cfe592` on `forum_id`
- index `forum_conversation_topic_poster_id_0dd4fa07` on `poster_id`
- index `forum_conversation_topic_slug_c74ce2cc` on `slug`
- index `forum_conversation_topic_type_92971eb5` on `type`
- index `forum_conversation_topic_status_e78d0ae4` on `status`
- index `forum_conversation_topic_approved_ad3fcbf9` on `approved`
- index `forum_conversation_topic_last_post_on_a287686b` on `last_post_on`
- index `forum_conversation_topic_first_post_id_ca473bd1` on `first_post_id`
- index `forum_conversation_topic_last_post_id_e14396a2` on `last_post_id`
- index `forum_conversation_topic_created_4e06996f` on `created`
- index `forum_conversation_topic_updated_3a9e00a1` on `updated`
- index `forum_conve_type_cc96d0_idx` on `type`, `last_post_on`

Relations:

- `forum_id` → [Forum](#forum).`id`, on delete cascade
- `poster_id` → [User](#user).`id`, on delete cascade
- `first_post_id` → [Post](#post).`id`, on delete set null
- `last_post_id` → [Post](#post).`id`, on delete set null
- `subscribers` — many to many with [User](#user)
- [Post](#post).`topic_id` → `id`, on delete cascade
- [TopicPoll](#topicpoll).`topic_id` → `id`, on delete cascade
- [TopicReadTrack](#topicreadtrack).`topic_id` → `id`, on delete cascade

## Post

Table `forum_conversation_post`, app `forum_conversation`.

Columns from `AbstractPost` (abstract): `topic_id`, `poster_id`, `anonymous_key`, `subject`, `content`, `username`, `approved`, `enable_signature`, `update_reason`, `updated_by_id`, `updates_count`.

Columns from `DatedModel` (abstract): `created`, `updated`.

![Post and the models it relates to](tables/Post.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `topic_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `poster_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `anonymous_key` | `CharField` | `varchar(100)` | yes |  |  |  |
| `subject` | `CharField` | `varchar(255)` |  |  |  |  |
| `content` | `MarkupTextField` | `text` |  |  |  |  |
| `username` | `CharField` | `varchar(155)` | yes |  |  |  |
| `approved` | `BooleanField` | `bool` |  |  | `True` |  |
| `enable_signature` | `BooleanField` | `bool` |  |  | `True` |  |
| `update_reason` | `CharField` | `varchar(255)` | yes |  |  |  |
| `updated_by_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `updates_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `_content_rendered` | `TextField` | `text` | yes |  |  |  |

Indexes:

- index `forum_conversation_post_topic_id_f6aaa418` on `topic_id`
- index `forum_conversation_post_poster_id_19c4e995` on `poster_id`
- index `forum_conversation_post_approved_a1090910` on `approved`
- index `forum_conversation_post_enable_signature_b1ce8b55` on `enable_signature`
- index `forum_conversation_post_updated_by_id_86093355` on `updated_by_id`

Relations:

- [Forum](#forum).`last_post_id` → `id`, on delete set null
- [Topic](#topic).`first_post_id` → `id`, on delete set null
- [Topic](#topic).`last_post_id` → `id`, on delete set null
- `topic_id` → [Topic](#topic).`id`, on delete cascade
- `poster_id` → [User](#user).`id`, on delete cascade
- `updated_by_id` → [User](#user).`id`, on delete set null
- [Attachment](#attachment).`post_id` → `id`, on delete cascade

## Attachment

Table `forum_attachments_attachment`, app `forum_attachments`.

Columns from `AbstractAttachment` (abstract): `post_id`, `file`, `comment`.

![Attachment and the models it relates to](tables/Attachment.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `post_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `file` | `FileField` | `varchar(100)` |  |  |  |  |
| `comment` | `CharField` | `varchar(255)` | yes |  |  |  |

Indexes:

- index `forum_attachments_attachment_post_id_0476a843` on `post_id`

Relations:

- `post_id` → [Post](#post).`id`, on delete cascade

## TopicPoll

Table `forum_polls_topicpoll`, app `forum_polls`.

Columns from `AbstractTopicPoll` (abstract): `topic_id`, `question`, `duration`, `max_options`, `user_changes`, `hide_results`.

Columns from `DatedModel` (abstract): `created`, `updated`.

![TopicPoll and the models it relates to](tables/TopicPoll.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `topic_id` | `OneToOneField` | `bigint` |  | FK, unique |  |  |
| `question` | `CharField` | `varchar(255)` |  |  |  |  |
| `duration` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |
| `max_options` | `PositiveSmallIntegerField` | `smallint unsigned` |  |  | `1` |  |
| `user_changes` | `BooleanField` | `bool` |  |  | `False` |  |
| `hide_results` | `BooleanField` | `bool` |  |  | `False` |  |

Relations:

- `topic_id` → [Topic](#topic).`id`, on delete cascade
- [TopicPollOption](#topicpolloption).`poll_id` → `id`, on delete cascade

## TopicPollOption

Table `forum_polls_topicpolloption`, app `forum_polls`.

Columns from `AbstractTopicPollOption` (abstract): `poll_id`, `text`.

![TopicPollOption and the models it relates to](tables/TopicPollOption.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `poll_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `text` | `CharField` | `varchar(255)` |  |  |  |  |

Indexes:

- index `forum_polls_topicpolloption_poll_id_a54cbd58` on `poll_id`

Relations:

- `poll_id` → [TopicPoll](#topicpoll).`id`, on delete cascade
- [TopicPollVote](#topicpollvote).`poll_option_id` → `id`, on delete cascade

## TopicPollVote

Table `forum_polls_topicpollvote`, app `forum_polls`.

Columns from `AbstractTopicPollVote` (abstract): `voter_id`, `anonymous_key`, `poll_option_id`, `timestamp`.

![TopicPollVote and the models it relates to](tables/TopicPollVote.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `voter_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `anonymous_key` | `CharField` | `varchar(100)` | yes |  |  |  |
| `poll_option_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `timestamp` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `forum_polls_topicpollvote_voter_id_0a287559` on `voter_id`
- index `forum_polls_topicpollvote_poll_option_id_a075b665` on `poll_option_id`

Relations:

- `voter_id` → [User](#user).`id`, on delete cascade
- `poll_option_id` → [TopicPollOption](#topicpolloption).`id`, on delete cascade

## ForumReadTrack

Table `forum_tracking_forumreadtrack`, app `forum_tracking`.

Columns from `AbstractForumReadTrack` (abstract): `user_id`, `forum_id`, `mark_time`.

![ForumReadTrack and the models it relates to](tables/ForumReadTrack.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `forum_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `mark_time` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `forum_tracking_forumreadtrack_user_id_932d4605` on `user_id`
- index `forum_tracking_forumreadtrack_forum_id_bbd3fb47` on `forum_id`
- index `forum_tracking_forumreadtrack_mark_time_72eec39e` on `mark_time`
- unique `forum_tracking_forumreadtrack_user_id_forum_id_3e64777a_uniq` on `user_id`, `forum_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `forum_id` → [Forum](#forum).`id`, on delete cascade

## TopicReadTrack

Table `forum_tracking_topicreadtrack`, app `forum_tracking`.

Columns from `AbstractTopicReadTrack` (abstract): `user_id`, `topic_id`, `mark_time`.

![TopicReadTrack and the models it relates to](tables/TopicReadTrack.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `topic_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `mark_time` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `forum_tracking_topicreadtrack_user_id_2762562b` on `user_id`
- index `forum_tracking_topicreadtrack_topic_id_9a53bd45` on `topic_id`
- index `forum_tracking_topicreadtrack_mark_time_7dafc483` on `mark_time`
- unique `forum_tracking_topicreadtrack_user_id_topic_id_6fe3e105_uniq` on `user_id`, `topic_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `topic_id` → [Topic](#topic).`id`, on delete cascade

## ForumProfile

Table `forum_member_forumprofile`, app `forum_member`.

Columns from `AbstractForumProfile` (abstract): `user_id`, `avatar`, `signature`, `posts_count`.

![ForumProfile and the models it relates to](tables/ForumProfile.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `OneToOneField` | `integer` |  | FK, unique |  |  |
| `avatar` | `ExtendedImageField` | `varchar(100)` | yes |  |  |  |
| `signature` | `MarkupTextField` | `text` | yes |  |  |  |
| `posts_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `_signature_rendered` | `TextField` | `text` | yes |  |  |  |

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

## ForumPermission

Table `forum_permission_forumpermission`, app `forum_permission`.

Columns from `AbstractForumPermission` (abstract): `codename`.

![ForumPermission and the models it relates to](tables/ForumPermission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `codename` | `CharField` | `varchar(150)` |  | unique |  |  |

Relations:

- [GroupForumPermission](#groupforumpermission).`permission_id` → `id`, on delete cascade
- [UserForumPermission](#userforumpermission).`permission_id` → `id`, on delete cascade

## GroupForumPermission

Table `forum_permission_groupforumpermission`, app `forum_permission`.

Columns from `AbstractGroupForumPermission` (abstract): `group_id`.

Columns from `BaseAuthForumPermission` (abstract): `permission_id`, `has_perm`, `forum_id`.

![GroupForumPermission and the models it relates to](tables/GroupForumPermission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `permission_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `has_perm` | `BooleanField` | `bool` |  |  | `True` |  |
| `forum_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `group_id` | `ForeignKey` | `integer` |  | FK |  |  |

Indexes:

- index `forum_permission_groupforumpermission_permission_id_2475fe70` on `permission_id`
- index `forum_permission_groupforumpermission_has_perm_48cae01d` on `has_perm`
- index `forum_permission_groupforumpermission_forum_id_d59d5cac` on `forum_id`
- index `forum_permission_groupforumpermission_group_id_b515635b` on `group_id`
- unique `forum_permission_groupforumpermission_permission_id_forum_id_group_id_a1e477c8_uniq` on `permission_id`, `forum_id`, `group_id`

Relations:

- `permission_id` → [ForumPermission](#forumpermission).`id`, on delete cascade
- `forum_id` → [Forum](#forum).`id`, on delete cascade
- `group_id` → [Group](#group).`id`, on delete cascade

## UserForumPermission

Table `forum_permission_userforumpermission`, app `forum_permission`.

Columns from `AbstractUserForumPermission` (abstract): `user_id`, `anonymous_user`, `authenticated_user`.

Columns from `BaseAuthForumPermission` (abstract): `permission_id`, `has_perm`, `forum_id`.

![UserForumPermission and the models it relates to](tables/UserForumPermission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `permission_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `has_perm` | `BooleanField` | `bool` |  |  | `True` |  |
| `forum_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `anonymous_user` | `BooleanField` | `bool` |  |  | `False` |  |
| `authenticated_user` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `forum_permission_userforumpermission_permission_id_9090e930` on `permission_id`
- index `forum_permission_userforumpermission_has_perm_1b5ee7ac` on `has_perm`
- index `forum_permission_userforumpermission_forum_id_f781f4d6` on `forum_id`
- index `forum_permission_userforumpermission_user_id_8106d02d` on `user_id`
- index `forum_permission_userforumpermission_anonymous_user_8aabbd9d` on `anonymous_user`
- index `forum_permission_userforumpermission_authenticated_user_b677afa4` on `authenticated_user`
- unique `forum_permission_userforumpermission_permission_id_forum_id_user_id_024a3693_uniq` on `permission_id`, `forum_id`, `user_id`

Relations:

- `permission_id` → [ForumPermission](#forumpermission).`id`, on delete cascade
- `forum_id` → [Forum](#forum).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## ForumType

| Name | Stored as |
| --- | --- |
| `Default forum` | `0` |
| `Category forum` | `1` |
| `Link forum` | `2` |

## TopicType

| Name | Stored as |
| --- | --- |
| `Default topic` | `0` |
| `Sticky` | `1` |
| `Announce` | `2` |

## TopicStatus

| Name | Stored as |
| --- | --- |
| `Topic unlocked` | `0` |
| `Topic locked` | `1` |
| `Topic moved` | `2` |
