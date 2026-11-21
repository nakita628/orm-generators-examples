# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Site](#site) · [Account](#account) · [SignupCode](#signupcode) · [SignupCodeResult](#signupcoderesult) · [EmailAddress](#emailaddress) · [EmailConfirmation](#emailconfirmation) · [AccountDeletion](#accountdeletion) · [PasswordHistory](#passwordhistory) · [PasswordExpiry](#passwordexpiry) · [ImageSet](#imageset) · [Image](#image) · [Announcement](#announcement) · [Dismissal](#dismissal) · [JoinInvitation](#joininvitation) · [InvitationStat](#invitationstat) · [SimpleTeam](#simpleteam) · [Team](#team) · [SimpleMembership](#simplemembership) · [Membership](#membership) · [Event](#event) · [Folder](#folder) · [Document](#document) · [FolderSharedUser](#foldershareduser) · [DocumentSharedUser](#documentshareduser) · [UserStorage](#userstorage)

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

CREATE TABLE "account_account" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL UNIQUE REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "timezone" varchar(100) NOT NULL, "language" varchar(10) NOT NULL);

CREATE TABLE "account_signupcode" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "code" varchar(64) NOT NULL UNIQUE, "max_uses" integer unsigned NOT NULL CHECK ("max_uses" >= 0), "expiry" datetime NULL, "inviter_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "email" varchar(254) NOT NULL, "notes" text NOT NULL, "sent" datetime NULL, "created" datetime NOT NULL, "use_count" integer unsigned NOT NULL CHECK ("use_count" >= 0));

CREATE TABLE "account_signupcoderesult" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "signup_code_id" bigint NOT NULL REFERENCES "account_signupcode" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "timestamp" datetime NOT NULL);

CREATE TABLE "account_emailaddress" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "email" varchar(254) NOT NULL UNIQUE, "verified" bool NOT NULL, "primary" bool NOT NULL);

CREATE TABLE "account_emailconfirmation" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "email_address_id" bigint NOT NULL REFERENCES "account_emailaddress" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL, "sent" datetime NULL, "key" varchar(64) NOT NULL UNIQUE);

CREATE TABLE "account_accountdeletion" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "email" varchar(254) NOT NULL, "date_requested" datetime NOT NULL, "date_expunged" datetime NULL);

CREATE TABLE "account_passwordhistory" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "password" varchar(255) NOT NULL, "timestamp" datetime NOT NULL);

CREATE TABLE "account_passwordexpiry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL UNIQUE REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "expiry" integer unsigned NOT NULL CHECK ("expiry" >= 0));

CREATE TABLE "pinax_images_imageset" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "primary_image_id" bigint NULL REFERENCES "pinax_images_image" ("id") DEFERRABLE INITIALLY DEFERRED, "created_by_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created_at" datetime NOT NULL);

CREATE TABLE "pinax_images_image" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "image_set_id" bigint NOT NULL REFERENCES "pinax_images_imageset" ("id") DEFERRABLE INITIALLY DEFERRED, "image" varchar(100) NOT NULL, "original_filename" varchar(500) NOT NULL, "created_by_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created_at" datetime NOT NULL);

CREATE TABLE "announcements_announcement" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "title" varchar(50) NOT NULL, "content" text NOT NULL, "creator_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "creation_date" datetime NOT NULL, "site_wide" bool NOT NULL, "members_only" bool NOT NULL, "dismissal_type" integer NOT NULL, "publish_start" datetime NOT NULL, "publish_end" datetime NULL);

CREATE TABLE "announcements_dismissal" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "announcement_id" bigint NOT NULL REFERENCES "announcements_announcement" ("id") DEFERRABLE INITIALLY DEFERRED, "dismissed_at" datetime NOT NULL);

CREATE TABLE "pinax_invitations_joininvitation" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "from_user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "to_user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "message" text NULL, "sent" datetime NOT NULL, "status" integer NOT NULL, "signup_code_id" bigint NOT NULL UNIQUE REFERENCES "account_signupcode" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "pinax_invitations_invitationstat" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL UNIQUE REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "invites_sent" integer NOT NULL, "invites_allocated" integer NOT NULL, "invites_accepted" integer NOT NULL);

CREATE TABLE "pinax_teams_simpleteam" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "member_access" varchar(20) NOT NULL, "manager_access" varchar(20) NOT NULL);

CREATE TABLE "pinax_teams_team" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "member_access" varchar(20) NOT NULL, "manager_access" varchar(20) NOT NULL, "slug" varchar(50) NOT NULL UNIQUE, "name" varchar(100) NOT NULL, "avatar" varchar(100) NOT NULL, "description" text NOT NULL, "creator_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL);

CREATE TABLE "pinax_teams_simplemembership" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "state" varchar(20) NOT NULL, "role" varchar(20) NOT NULL, "created" datetime NOT NULL, "team_id" bigint NOT NULL REFERENCES "pinax_teams_simpleteam" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "invite_id" bigint NULL REFERENCES "pinax_invitations_joininvitation" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "pinax_teams_membership" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "state" varchar(20) NOT NULL, "role" varchar(20) NOT NULL, "created" datetime NOT NULL, "team_id" bigint NOT NULL REFERENCES "pinax_teams_team" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "invite_id" bigint NULL REFERENCES "pinax_invitations_joininvitation" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "pinax_events_event" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "image" varchar(100) NOT NULL, "secondary_image" varchar(100) NOT NULL, "title" varchar(200) NOT NULL, "url" text NOT NULL, "where" varchar(200) NOT NULL, "what" text NOT NULL, "what_html" text NOT NULL, "start_date" date NOT NULL, "end_date" date NOT NULL, "published_at" datetime NOT NULL, "created_at" datetime NOT NULL);

CREATE TABLE "documents_folder" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(140) NOT NULL, "parent_id" bigint NULL REFERENCES "documents_folder" ("id") DEFERRABLE INITIALLY DEFERRED, "author_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL, "modified" datetime NOT NULL, "modified_by_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "documents_document" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(255) NOT NULL, "folder_id" bigint NULL REFERENCES "documents_folder" ("id") DEFERRABLE INITIALLY DEFERRED, "author_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL, "modified" datetime NOT NULL, "modified_by_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "file" varchar(100) NOT NULL, "original_filename" varchar(500) NOT NULL);

CREATE TABLE "documents_foldershareduser" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "folder_id" bigint NOT NULL REFERENCES "documents_folder" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "documents_documentshareduser" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "document_id" bigint NOT NULL REFERENCES "documents_document" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "documents_userstorage" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL UNIQUE REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "bytes_used" bigint NOT NULL, "bytes_total" bigint NOT NULL);

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

CREATE INDEX "account_signupcode_inviter_id_9706983e" ON "account_signupcode" ("inviter_id");

CREATE INDEX "account_signupcoderesult_signup_code_id_1adc486e" ON "account_signupcoderesult" ("signup_code_id");

CREATE INDEX "account_signupcoderesult_user_id_c10dc604" ON "account_signupcoderesult" ("user_id");

CREATE INDEX "account_emailaddress_user_id_2c513194" ON "account_emailaddress" ("user_id");

CREATE INDEX "account_emailconfirmation_email_address_id_5b7f8c58" ON "account_emailconfirmation" ("email_address_id");

CREATE INDEX "account_accountdeletion_user_id_c205475f" ON "account_accountdeletion" ("user_id");

CREATE INDEX "account_passwordhistory_user_id_dc325181" ON "account_passwordhistory" ("user_id");

CREATE INDEX "pinax_images_imageset_primary_image_id_5335cee9" ON "pinax_images_imageset" ("primary_image_id");

CREATE INDEX "pinax_images_imageset_created_by_id_52728f5f" ON "pinax_images_imageset" ("created_by_id");

CREATE INDEX "pinax_images_image_image_set_id_eb99d8d0" ON "pinax_images_image" ("image_set_id");

CREATE INDEX "pinax_images_image_created_by_id_31a81e5c" ON "pinax_images_image" ("created_by_id");

CREATE INDEX "announcements_announcement_creator_id_d9be5764" ON "announcements_announcement" ("creator_id");

CREATE INDEX "announcements_dismissal_user_id_927879ee" ON "announcements_dismissal" ("user_id");

CREATE INDEX "announcements_dismissal_announcement_id_719fe978" ON "announcements_dismissal" ("announcement_id");

CREATE INDEX "pinax_invitations_joininvitation_from_user_id_6298d80c" ON "pinax_invitations_joininvitation" ("from_user_id");

CREATE INDEX "pinax_invitations_joininvitation_to_user_id_d0fa4add" ON "pinax_invitations_joininvitation" ("to_user_id");

CREATE INDEX "pinax_teams_team_creator_id_464a2b68" ON "pinax_teams_team" ("creator_id");

CREATE UNIQUE INDEX "pinax_teams_simplemembership_team_id_user_id_invite_id_668da23b_uniq" ON "pinax_teams_simplemembership" ("team_id", "user_id", "invite_id");

CREATE INDEX "pinax_teams_simplemembership_team_id_37407f9d" ON "pinax_teams_simplemembership" ("team_id");

CREATE INDEX "pinax_teams_simplemembership_user_id_1152f95c" ON "pinax_teams_simplemembership" ("user_id");

CREATE INDEX "pinax_teams_simplemembership_invite_id_eedba0e4" ON "pinax_teams_simplemembership" ("invite_id");

CREATE UNIQUE INDEX "pinax_teams_membership_team_id_user_id_invite_id_8ead1786_uniq" ON "pinax_teams_membership" ("team_id", "user_id", "invite_id");

CREATE INDEX "pinax_teams_membership_team_id_72556e8a" ON "pinax_teams_membership" ("team_id");

CREATE INDEX "pinax_teams_membership_user_id_5237927b" ON "pinax_teams_membership" ("user_id");

CREATE INDEX "pinax_teams_membership_invite_id_8ab7ad37" ON "pinax_teams_membership" ("invite_id");

CREATE INDEX "documents_folder_parent_id_55715a39" ON "documents_folder" ("parent_id");

CREATE INDEX "documents_folder_author_id_332f41e9" ON "documents_folder" ("author_id");

CREATE INDEX "documents_folder_modified_by_id_6886f342" ON "documents_folder" ("modified_by_id");

CREATE INDEX "documents_document_folder_id_f2e8c7fd" ON "documents_document" ("folder_id");

CREATE INDEX "documents_document_author_id_6440570b" ON "documents_document" ("author_id");

CREATE INDEX "documents_document_modified_by_id_2f944095" ON "documents_document" ("modified_by_id");

CREATE UNIQUE INDEX "documents_foldershareduser_folder_id_user_id_addf94c2_uniq" ON "documents_foldershareduser" ("folder_id", "user_id");

CREATE INDEX "documents_foldershareduser_user_id_cd64431e" ON "documents_foldershareduser" ("user_id");

CREATE INDEX "documents_foldershareduser_folder_id_f5c77e47" ON "documents_foldershareduser" ("folder_id");

CREATE UNIQUE INDEX "documents_documentshareduser_document_id_user_id_fdf4687e_uniq" ON "documents_documentshareduser" ("document_id", "user_id");

CREATE INDEX "documents_documentshareduser_user_id_91a580b1" ON "documents_documentshareduser" ("user_id");

CREATE INDEX "documents_documentshareduser_document_id_37379472" ON "documents_documentshareduser" ("document_id");
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
- [Account](#account).`user_id` → `id`, on delete cascade
- [SignupCode](#signupcode).`inviter_id` → `id`, on delete cascade
- [SignupCodeResult](#signupcoderesult).`user_id` → `id`, on delete cascade
- [EmailAddress](#emailaddress).`user_id` → `id`, on delete cascade
- [AccountDeletion](#accountdeletion).`user_id` → `id`, on delete set null
- [PasswordHistory](#passwordhistory).`user_id` → `id`, on delete cascade
- [PasswordExpiry](#passwordexpiry).`user_id` → `id`, on delete cascade
- [ImageSet](#imageset).`created_by_id` → `id`, on delete cascade
- [Image](#image).`created_by_id` → `id`, on delete cascade
- [Announcement](#announcement).`creator_id` → `id`, on delete cascade
- [Dismissal](#dismissal).`user_id` → `id`, on delete cascade
- [JoinInvitation](#joininvitation).`from_user_id` → `id`, on delete cascade
- [JoinInvitation](#joininvitation).`to_user_id` → `id`, on delete cascade
- [InvitationStat](#invitationstat).`user_id` → `id`, on delete cascade
- [Team](#team).`creator_id` → `id`, on delete cascade
- [SimpleMembership](#simplemembership).`user_id` → `id`, on delete set null
- [Membership](#membership).`user_id` → `id`, on delete set null
- [Folder](#folder).`author_id` → `id`, on delete cascade
- [Folder](#folder).`modified_by_id` → `id`, on delete cascade
- [Document](#document).`author_id` → `id`, on delete cascade
- [Document](#document).`modified_by_id` → `id`, on delete cascade
- [FolderSharedUser](#foldershareduser).`user_id` → `id`, on delete cascade
- [DocumentSharedUser](#documentshareduser).`user_id` → `id`, on delete cascade
- [UserStorage](#userstorage).`user_id` → `id`, on delete cascade

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

## Account

Table `account_account`, app `account`.

![Account and the models it relates to](tables/Account.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `OneToOneField` | `integer` |  | FK, unique |  |  |
| `timezone` | `AccountTimezone` | `varchar(100)` |  |  | `''` |  |
| `language` | `AccountLanguage` | `varchar(10)` |  |  | `'en'` |  |

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

## SignupCode

Table `account_signupcode`, app `account`.

![SignupCode and the models it relates to](tables/SignupCode.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `code` | `CharField` | `varchar(64)` |  | unique |  |  |
| `max_uses` | `PositiveIntegerField` | `integer unsigned` |  |  | `1` |  |
| `expiry` | `DateTimeField` | `datetime` | yes |  |  |  |
| `inviter_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `notes` | `TextField` | `text` |  |  |  |  |
| `sent` | `DateTimeField` | `datetime` | yes |  |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `use_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |

Indexes:

- index `account_signupcode_inviter_id_9706983e` on `inviter_id`

Relations:

- `inviter_id` → [User](#user).`id`, on delete cascade
- [SignupCodeResult](#signupcoderesult).`signup_code_id` → `id`, on delete cascade
- [JoinInvitation](#joininvitation).`signup_code_id` → `id`, on delete cascade

## SignupCodeResult

Table `account_signupcoderesult`, app `account`.

![SignupCodeResult and the models it relates to](tables/SignupCodeResult.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `signup_code_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `timestamp` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `account_signupcoderesult_signup_code_id_1adc486e` on `signup_code_id`
- index `account_signupcoderesult_user_id_c10dc604` on `user_id`

Relations:

- `signup_code_id` → [SignupCode](#signupcode).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## EmailAddress

Table `account_emailaddress`, app `account`.

![EmailAddress and the models it relates to](tables/EmailAddress.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `email` | `EmailField` | `varchar(254)` |  | unique |  |  |
| `verified` | `BooleanField` | `bool` |  |  | `False` |  |
| `primary` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `account_emailaddress_user_id_2c513194` on `user_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [EmailConfirmation](#emailconfirmation).`email_address_id` → `id`, on delete cascade

## EmailConfirmation

Table `account_emailconfirmation`, app `account`.

![EmailConfirmation and the models it relates to](tables/EmailConfirmation.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `email_address_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `sent` | `DateTimeField` | `datetime` | yes |  |  |  |
| `key` | `CharField` | `varchar(64)` |  | unique |  |  |

Indexes:

- index `account_emailconfirmation_email_address_id_5b7f8c58` on `email_address_id`

Relations:

- `email_address_id` → [EmailAddress](#emailaddress).`id`, on delete cascade

## AccountDeletion

Table `account_accountdeletion`, app `account`.

![AccountDeletion and the models it relates to](tables/AccountDeletion.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `date_requested` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `date_expunged` | `DateTimeField` | `datetime` | yes |  |  |  |

Indexes:

- index `account_accountdeletion_user_id_c205475f` on `user_id`

Relations:

- `user_id` → [User](#user).`id`, on delete set null

## PasswordHistory

Table `account_passwordhistory`, app `account`.

Contains single password history for user.

![PasswordHistory and the models it relates to](tables/PasswordHistory.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `password` | `CharField` | `varchar(255)` |  |  |  |  |
| `timestamp` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `account_passwordhistory_user_id_dc325181` on `user_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

## PasswordExpiry

Table `account_passwordexpiry`, app `account`.

Holds the password expiration period for a single user.

![PasswordExpiry and the models it relates to](tables/PasswordExpiry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `OneToOneField` | `integer` |  | FK, unique |  |  |
| `expiry` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

## ImageSet

Table `pinax_images_imageset`, app `pinax_images`.

Container for a group of Images.

![ImageSet and the models it relates to](tables/ImageSet.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `primary_image_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `created_by_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `pinax_images_imageset_primary_image_id_5335cee9` on `primary_image_id`
- index `pinax_images_imageset_created_by_id_52728f5f` on `created_by_id`

Relations:

- `primary_image_id` → [Image](#image).`id`, on delete set null
- `created_by_id` → [User](#user).`id`, on delete cascade
- [Image](#image).`image_set_id` → `id`, on delete cascade

## Image

Table `pinax_images_image`, app `pinax_images`.

![Image and the models it relates to](tables/Image.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `image_set_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `image` | `ImageField` | `varchar(100)` |  |  |  |  |
| `original_filename` | `CharField` | `varchar(500)` |  |  |  |  |
| `created_by_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `pinax_images_image_image_set_id_eb99d8d0` on `image_set_id`
- index `pinax_images_image_created_by_id_31a81e5c` on `created_by_id`

Relations:

- [ImageSet](#imageset).`primary_image_id` → `id`, on delete set null
- `image_set_id` → [ImageSet](#imageset).`id`, on delete cascade
- `created_by_id` → [User](#user).`id`, on delete cascade

## Announcement

Table `announcements_announcement`, app `announcements`.

A single announcement.

![Announcement and the models it relates to](tables/Announcement.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `title` | `CharField` | `varchar(50)` |  |  |  |  |
| `content` | `TextField` | `text` |  |  |  |  |
| `creator_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `creation_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `site_wide` | `BooleanField` | `bool` |  |  | `False` |  |
| `members_only` | `BooleanField` | `bool` |  |  | `False` |  |
| `dismissal_type` | `AnnouncementDismissalType` | `integer` |  |  | `2` |  |
| `publish_start` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `publish_end` | `DateTimeField` | `datetime` | yes |  |  |  |

Indexes:

- index `announcements_announcement_creator_id_d9be5764` on `creator_id`

Relations:

- `creator_id` → [User](#user).`id`, on delete cascade
- [Dismissal](#dismissal).`announcement_id` → `id`, on delete cascade

## Dismissal

Table `announcements_dismissal`, app `announcements`.

![Dismissal and the models it relates to](tables/Dismissal.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `announcement_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `dismissed_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `announcements_dismissal_user_id_927879ee` on `user_id`
- index `announcements_dismissal_announcement_id_719fe978` on `announcement_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `announcement_id` → [Announcement](#announcement).`id`, on delete cascade

## JoinInvitation

Table `pinax_invitations_joininvitation`, app `pinax_invitations`.

![JoinInvitation and the models it relates to](tables/JoinInvitation.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `from_user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `to_user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `message` | `TextField` | `text` | yes |  |  |  |
| `sent` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `status` | `JoinInvitationStatus` | `integer` |  |  |  |  |
| `signup_code_id` | `OneToOneField` | `bigint` |  | FK, unique |  |  |

Indexes:

- index `pinax_invitations_joininvitation_from_user_id_6298d80c` on `from_user_id`
- index `pinax_invitations_joininvitation_to_user_id_d0fa4add` on `to_user_id`

Relations:

- `from_user_id` → [User](#user).`id`, on delete cascade
- `to_user_id` → [User](#user).`id`, on delete cascade
- `signup_code_id` → [SignupCode](#signupcode).`id`, on delete cascade
- [SimpleMembership](#simplemembership).`invite_id` → `id`, on delete set null
- [Membership](#membership).`invite_id` → `id`, on delete set null

## InvitationStat

Table `pinax_invitations_invitationstat`, app `pinax_invitations`.

![InvitationStat and the models it relates to](tables/InvitationStat.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `OneToOneField` | `integer` |  | FK, unique |  |  |
| `invites_sent` | `IntegerField` | `integer` |  |  | `0` |  |
| `invites_allocated` | `IntegerField` | `integer` |  |  | `0` |  |
| `invites_accepted` | `IntegerField` | `integer` |  |  | `0` |  |

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

## SimpleTeam

Table `pinax_teams_simpleteam`, app `pinax_teams`.

Columns from `BaseTeam` (abstract): `member_access`, `manager_access`.

![SimpleTeam and the models it relates to](tables/SimpleTeam.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `member_access` | `SimpleTeamMemberAccess` | `varchar(20)` |  |  |  |  |
| `manager_access` | `SimpleTeamManagerAccess` | `varchar(20)` |  |  |  |  |

Relations:

- [SimpleMembership](#simplemembership).`team_id` → `id`, on delete cascade

## Team

Table `pinax_teams_team`, app `pinax_teams`.

Columns from `BaseTeam` (abstract): `member_access`, `manager_access`.

![Team and the models it relates to](tables/Team.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `member_access` | `TeamMemberAccess` | `varchar(20)` |  |  |  |  |
| `manager_access` | `TeamManagerAccess` | `varchar(20)` |  |  |  |  |
| `slug` | `SlugField` | `varchar(50)` |  | unique |  |  |
| `name` | `CharField` | `varchar(100)` |  |  |  |  |
| `avatar` | `ImageField` | `varchar(100)` |  |  |  |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `creator_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `pinax_teams_team_creator_id_464a2b68` on `creator_id`

Relations:

- `creator_id` → [User](#user).`id`, on delete cascade
- [Membership](#membership).`team_id` → `id`, on delete cascade

## SimpleMembership

Table `pinax_teams_simplemembership`, app `pinax_teams`.

Columns from `BaseMembership` (abstract): `state`, `role`, `created`.

![SimpleMembership and the models it relates to](tables/SimpleMembership.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `state` | `SimpleMembershipState` | `varchar(20)` |  |  |  |  |
| `role` | `SimpleMembershipRole` | `varchar(20)` |  |  | `'member'` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `team_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `invite_id` | `ForeignKey` | `bigint` | yes | FK |  |  |

Indexes:

- index `pinax_teams_simplemembership_team_id_37407f9d` on `team_id`
- index `pinax_teams_simplemembership_user_id_1152f95c` on `user_id`
- index `pinax_teams_simplemembership_invite_id_eedba0e4` on `invite_id`
- unique `pinax_teams_simplemembership_team_id_user_id_invite_id_668da23b_uniq` on `team_id`, `user_id`, `invite_id`

Relations:

- `team_id` → [SimpleTeam](#simpleteam).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete set null
- `invite_id` → [JoinInvitation](#joininvitation).`id`, on delete set null

## Membership

Table `pinax_teams_membership`, app `pinax_teams`.

Columns from `BaseMembership` (abstract): `state`, `role`, `created`.

![Membership and the models it relates to](tables/Membership.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `state` | `MembershipState` | `varchar(20)` |  |  |  |  |
| `role` | `MembershipRole` | `varchar(20)` |  |  | `'member'` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `team_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `invite_id` | `ForeignKey` | `bigint` | yes | FK |  |  |

Indexes:

- index `pinax_teams_membership_team_id_72556e8a` on `team_id`
- index `pinax_teams_membership_user_id_5237927b` on `user_id`
- index `pinax_teams_membership_invite_id_8ab7ad37` on `invite_id`
- unique `pinax_teams_membership_team_id_user_id_invite_id_8ead1786_uniq` on `team_id`, `user_id`, `invite_id`

Relations:

- `team_id` → [Team](#team).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete set null
- `invite_id` → [JoinInvitation](#joininvitation).`id`, on delete set null

## Event

Table `pinax_events_event`, app `pinax_events`.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `image` | `ImageField` | `varchar(100)` |  |  |  |  |
| `secondary_image` | `ImageField` | `varchar(100)` |  |  |  |  |
| `title` | `CharField` | `varchar(200)` |  |  |  |  |
| `url` | `TextField` | `text` |  |  |  |  |
| `where` | `CharField` | `varchar(200)` |  |  |  |  |
| `what` | `TextField` | `text` |  |  |  |  |
| `what_html` | `TextField` | `text` |  |  |  |  |
| `start_date` | `DateField` | `date` |  |  |  |  |
| `end_date` | `DateField` | `date` |  |  |  |  |
| `published_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |

## Folder

Table `documents_folder`, app `documents`.

![Folder and the models it relates to](tables/Folder.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(140)` |  |  |  |  |
| `parent_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `author_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `modified_by_id` | `ForeignKey` | `integer` |  | FK |  |  |

Indexes:

- index `documents_folder_parent_id_55715a39` on `parent_id`
- index `documents_folder_author_id_332f41e9` on `author_id`
- index `documents_folder_modified_by_id_6886f342` on `modified_by_id`

Relations:

- `parent_id` → [Folder](#folder).`id`, on delete cascade
- `author_id` → [User](#user).`id`, on delete cascade
- `modified_by_id` → [User](#user).`id`, on delete cascade
- [Document](#document).`folder_id` → `id`, on delete cascade
- [FolderSharedUser](#foldershareduser).`folder_id` → `id`, on delete cascade

## Document

Table `documents_document`, app `documents`.

![Document and the models it relates to](tables/Document.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `folder_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `author_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `modified_by_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `file` | `FileField` | `varchar(100)` |  |  |  |  |
| `original_filename` | `CharField` | `varchar(500)` |  |  |  |  |

Indexes:

- index `documents_document_folder_id_f2e8c7fd` on `folder_id`
- index `documents_document_author_id_6440570b` on `author_id`
- index `documents_document_modified_by_id_2f944095` on `modified_by_id`

Relations:

- `folder_id` → [Folder](#folder).`id`, on delete cascade
- `author_id` → [User](#user).`id`, on delete cascade
- `modified_by_id` → [User](#user).`id`, on delete cascade
- [DocumentSharedUser](#documentshareduser).`document_id` → `id`, on delete cascade

## FolderSharedUser

Table `documents_foldershareduser`, app `documents`.

Columns from `MemberSharedUser` (abstract): `user_id`.

![FolderSharedUser and the models it relates to](tables/FolderSharedUser.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `folder_id` | `ForeignKey` | `bigint` |  | FK |  |  |

Indexes:

- index `documents_foldershareduser_user_id_cd64431e` on `user_id`
- index `documents_foldershareduser_folder_id_f5c77e47` on `folder_id`
- unique `documents_foldershareduser_folder_id_user_id_addf94c2_uniq` on `folder_id`, `user_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `folder_id` → [Folder](#folder).`id`, on delete cascade

## DocumentSharedUser

Table `documents_documentshareduser`, app `documents`.

Columns from `MemberSharedUser` (abstract): `user_id`.

![DocumentSharedUser and the models it relates to](tables/DocumentSharedUser.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `document_id` | `ForeignKey` | `bigint` |  | FK |  |  |

Indexes:

- index `documents_documentshareduser_user_id_91a580b1` on `user_id`
- index `documents_documentshareduser_document_id_37379472` on `document_id`
- unique `documents_documentshareduser_document_id_user_id_fdf4687e_uniq` on `document_id`, `user_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `document_id` → [Document](#document).`id`, on delete cascade

## UserStorage

Table `documents_userstorage`, app `documents`.

![UserStorage and the models it relates to](tables/UserStorage.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `OneToOneField` | `integer` |  | FK, unique |  |  |
| `bytes_used` | `BigIntegerField` | `bigint` |  |  | `0` |  |
| `bytes_total` | `BigIntegerField` | `bigint` |  |  | `0` |  |

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## AccountTimezone

| Name | Stored as |
| --- | --- |
| `Africa/Abidjan` | `Africa/Abidjan` |
| `Africa/Accra` | `Africa/Accra` |
| `Africa/Addis_Ababa` | `Africa/Addis_Ababa` |
| `Africa/Algiers` | `Africa/Algiers` |
| `Africa/Asmara` | `Africa/Asmara` |
| `Africa/Asmera` | `Africa/Asmera` |
| `Africa/Bamako` | `Africa/Bamako` |
| `Africa/Bangui` | `Africa/Bangui` |
| `Africa/Banjul` | `Africa/Banjul` |
| `Africa/Bissau` | `Africa/Bissau` |
| `Africa/Blantyre` | `Africa/Blantyre` |
| `Africa/Brazzaville` | `Africa/Brazzaville` |
| `Africa/Bujumbura` | `Africa/Bujumbura` |
| `Africa/Cairo` | `Africa/Cairo` |
| `Africa/Casablanca` | `Africa/Casablanca` |
| `Africa/Ceuta` | `Africa/Ceuta` |
| `Africa/Conakry` | `Africa/Conakry` |
| `Africa/Dakar` | `Africa/Dakar` |
| `Africa/Dar_es_Salaam` | `Africa/Dar_es_Salaam` |
| `Africa/Djibouti` | `Africa/Djibouti` |
| `Africa/Douala` | `Africa/Douala` |
| `Africa/El_Aaiun` | `Africa/El_Aaiun` |
| `Africa/Freetown` | `Africa/Freetown` |
| `Africa/Gaborone` | `Africa/Gaborone` |
| `Africa/Harare` | `Africa/Harare` |
| `Africa/Johannesburg` | `Africa/Johannesburg` |
| `Africa/Juba` | `Africa/Juba` |
| `Africa/Kampala` | `Africa/Kampala` |
| `Africa/Khartoum` | `Africa/Khartoum` |
| `Africa/Kigali` | `Africa/Kigali` |
| `Africa/Kinshasa` | `Africa/Kinshasa` |
| `Africa/Lagos` | `Africa/Lagos` |
| `Africa/Libreville` | `Africa/Libreville` |
| `Africa/Lome` | `Africa/Lome` |
| `Africa/Luanda` | `Africa/Luanda` |
| `Africa/Lubumbashi` | `Africa/Lubumbashi` |
| `Africa/Lusaka` | `Africa/Lusaka` |
| `Africa/Malabo` | `Africa/Malabo` |
| `Africa/Maputo` | `Africa/Maputo` |
| `Africa/Maseru` | `Africa/Maseru` |
| `Africa/Mbabane` | `Africa/Mbabane` |
| `Africa/Mogadishu` | `Africa/Mogadishu` |
| `Africa/Monrovia` | `Africa/Monrovia` |
| `Africa/Nairobi` | `Africa/Nairobi` |
| `Africa/Ndjamena` | `Africa/Ndjamena` |
| `Africa/Niamey` | `Africa/Niamey` |
| `Africa/Nouakchott` | `Africa/Nouakchott` |
| `Africa/Ouagadougou` | `Africa/Ouagadougou` |
| `Africa/Porto-Novo` | `Africa/Porto-Novo` |
| `Africa/Sao_Tome` | `Africa/Sao_Tome` |
| `Africa/Timbuktu` | `Africa/Timbuktu` |
| `Africa/Tripoli` | `Africa/Tripoli` |
| `Africa/Tunis` | `Africa/Tunis` |
| `Africa/Windhoek` | `Africa/Windhoek` |
| `America/Adak` | `America/Adak` |
| `America/Anchorage` | `America/Anchorage` |
| `America/Anguilla` | `America/Anguilla` |
| `America/Antigua` | `America/Antigua` |
| `America/Araguaina` | `America/Araguaina` |
| `America/Argentina/Buenos_Aires` | `America/Argentina/Buenos_Aires` |
| `America/Argentina/Catamarca` | `America/Argentina/Catamarca` |
| `America/Argentina/ComodRivadavia` | `America/Argentina/ComodRivadavia` |
| `America/Argentina/Cordoba` | `America/Argentina/Cordoba` |
| `America/Argentina/Jujuy` | `America/Argentina/Jujuy` |
| `America/Argentina/La_Rioja` | `America/Argentina/La_Rioja` |
| `America/Argentina/Mendoza` | `America/Argentina/Mendoza` |
| `America/Argentina/Rio_Gallegos` | `America/Argentina/Rio_Gallegos` |
| `America/Argentina/Salta` | `America/Argentina/Salta` |
| `America/Argentina/San_Juan` | `America/Argentina/San_Juan` |
| `America/Argentina/San_Luis` | `America/Argentina/San_Luis` |
| `America/Argentina/Tucuman` | `America/Argentina/Tucuman` |
| `America/Argentina/Ushuaia` | `America/Argentina/Ushuaia` |
| `America/Aruba` | `America/Aruba` |
| `America/Asuncion` | `America/Asuncion` |
| `America/Atikokan` | `America/Atikokan` |
| `America/Atka` | `America/Atka` |
| `America/Bahia` | `America/Bahia` |
| `America/Bahia_Banderas` | `America/Bahia_Banderas` |
| `America/Barbados` | `America/Barbados` |
| `America/Belem` | `America/Belem` |
| `America/Belize` | `America/Belize` |
| `America/Blanc-Sablon` | `America/Blanc-Sablon` |
| `America/Boa_Vista` | `America/Boa_Vista` |
| `America/Bogota` | `America/Bogota` |
| `America/Boise` | `America/Boise` |
| `America/Buenos_Aires` | `America/Buenos_Aires` |
| `America/Cambridge_Bay` | `America/Cambridge_Bay` |
| `America/Campo_Grande` | `America/Campo_Grande` |
| `America/Cancun` | `America/Cancun` |
| `America/Caracas` | `America/Caracas` |
| `America/Catamarca` | `America/Catamarca` |
| `America/Cayenne` | `America/Cayenne` |
| `America/Cayman` | `America/Cayman` |
| `America/Chicago` | `America/Chicago` |
| `America/Chihuahua` | `America/Chihuahua` |
| `America/Coral_Harbour` | `America/Coral_Harbour` |
| `America/Cordoba` | `America/Cordoba` |
| `America/Costa_Rica` | `America/Costa_Rica` |
| `America/Creston` | `America/Creston` |
| `America/Cuiaba` | `America/Cuiaba` |
| `America/Curacao` | `America/Curacao` |
| `America/Danmarkshavn` | `America/Danmarkshavn` |
| `America/Dawson` | `America/Dawson` |
| `America/Dawson_Creek` | `America/Dawson_Creek` |
| `America/Denver` | `America/Denver` |
| `America/Detroit` | `America/Detroit` |
| `America/Dominica` | `America/Dominica` |
| `America/Edmonton` | `America/Edmonton` |
| `America/Eirunepe` | `America/Eirunepe` |
| `America/El_Salvador` | `America/El_Salvador` |
| `America/Ensenada` | `America/Ensenada` |
| `America/Fort_Wayne` | `America/Fort_Wayne` |
| `America/Fortaleza` | `America/Fortaleza` |
| `America/Glace_Bay` | `America/Glace_Bay` |
| `America/Godthab` | `America/Godthab` |
| `America/Goose_Bay` | `America/Goose_Bay` |
| `America/Grand_Turk` | `America/Grand_Turk` |
| `America/Grenada` | `America/Grenada` |
| `America/Guadeloupe` | `America/Guadeloupe` |
| `America/Guatemala` | `America/Guatemala` |
| `America/Guayaquil` | `America/Guayaquil` |
| `America/Guyana` | `America/Guyana` |
| `America/Halifax` | `America/Halifax` |
| `America/Havana` | `America/Havana` |
| `America/Hermosillo` | `America/Hermosillo` |
| `America/Indiana/Indianapolis` | `America/Indiana/Indianapolis` |
| `America/Indiana/Knox` | `America/Indiana/Knox` |
| `America/Indiana/Marengo` | `America/Indiana/Marengo` |
| `America/Indiana/Petersburg` | `America/Indiana/Petersburg` |
| `America/Indiana/Tell_City` | `America/Indiana/Tell_City` |
| `America/Indiana/Vevay` | `America/Indiana/Vevay` |
| `America/Indiana/Vincennes` | `America/Indiana/Vincennes` |
| `America/Indiana/Winamac` | `America/Indiana/Winamac` |
| `America/Indianapolis` | `America/Indianapolis` |
| `America/Inuvik` | `America/Inuvik` |
| `America/Iqaluit` | `America/Iqaluit` |
| `America/Jamaica` | `America/Jamaica` |
| `America/Jujuy` | `America/Jujuy` |
| `America/Juneau` | `America/Juneau` |
| `America/Kentucky/Louisville` | `America/Kentucky/Louisville` |
| `America/Kentucky/Monticello` | `America/Kentucky/Monticello` |
| `America/Knox_IN` | `America/Knox_IN` |
| `America/Kralendijk` | `America/Kralendijk` |
| `America/La_Paz` | `America/La_Paz` |
| `America/Lima` | `America/Lima` |
| `America/Los_Angeles` | `America/Los_Angeles` |
| `America/Louisville` | `America/Louisville` |
| `America/Lower_Princes` | `America/Lower_Princes` |
| `America/Maceio` | `America/Maceio` |
| `America/Managua` | `America/Managua` |
| `America/Manaus` | `America/Manaus` |
| `America/Marigot` | `America/Marigot` |
| `America/Martinique` | `America/Martinique` |
| `America/Matamoros` | `America/Matamoros` |
| `America/Mazatlan` | `America/Mazatlan` |
| `America/Mendoza` | `America/Mendoza` |
| `America/Menominee` | `America/Menominee` |
| `America/Merida` | `America/Merida` |
| `America/Metlakatla` | `America/Metlakatla` |
| `America/Mexico_City` | `America/Mexico_City` |
| `America/Miquelon` | `America/Miquelon` |
| `America/Moncton` | `America/Moncton` |
| `America/Monterrey` | `America/Monterrey` |
| `America/Montevideo` | `America/Montevideo` |
| `America/Montreal` | `America/Montreal` |
| `America/Montserrat` | `America/Montserrat` |
| `America/Nassau` | `America/Nassau` |
| `America/New_York` | `America/New_York` |
| `America/Nipigon` | `America/Nipigon` |
| `America/Nome` | `America/Nome` |
| `America/Noronha` | `America/Noronha` |
| `America/North_Dakota/Beulah` | `America/North_Dakota/Beulah` |
| `America/North_Dakota/Center` | `America/North_Dakota/Center` |
| `America/North_Dakota/New_Salem` | `America/North_Dakota/New_Salem` |
| `America/Ojinaga` | `America/Ojinaga` |
| `America/Panama` | `America/Panama` |
| `America/Pangnirtung` | `America/Pangnirtung` |
| `America/Paramaribo` | `America/Paramaribo` |
| `America/Phoenix` | `America/Phoenix` |
| `America/Port-au-Prince` | `America/Port-au-Prince` |
| `America/Port_of_Spain` | `America/Port_of_Spain` |
| `America/Porto_Acre` | `America/Porto_Acre` |
| `America/Porto_Velho` | `America/Porto_Velho` |
| `America/Puerto_Rico` | `America/Puerto_Rico` |
| `America/Rainy_River` | `America/Rainy_River` |
| `America/Rankin_Inlet` | `America/Rankin_Inlet` |
| `America/Recife` | `America/Recife` |
| `America/Regina` | `America/Regina` |
| `America/Resolute` | `America/Resolute` |
| `America/Rio_Branco` | `America/Rio_Branco` |
| `America/Rosario` | `America/Rosario` |
| `America/Santa_Isabel` | `America/Santa_Isabel` |
| `America/Santarem` | `America/Santarem` |
| `America/Santiago` | `America/Santiago` |
| `America/Santo_Domingo` | `America/Santo_Domingo` |
| `America/Sao_Paulo` | `America/Sao_Paulo` |
| `America/Scoresbysund` | `America/Scoresbysund` |
| `America/Shiprock` | `America/Shiprock` |
| `America/Sitka` | `America/Sitka` |
| `America/St_Barthelemy` | `America/St_Barthelemy` |
| `America/St_Johns` | `America/St_Johns` |
| `America/St_Kitts` | `America/St_Kitts` |
| `America/St_Lucia` | `America/St_Lucia` |
| `America/St_Thomas` | `America/St_Thomas` |
| `America/St_Vincent` | `America/St_Vincent` |
| `America/Swift_Current` | `America/Swift_Current` |
| `America/Tegucigalpa` | `America/Tegucigalpa` |
| `America/Thule` | `America/Thule` |
| `America/Thunder_Bay` | `America/Thunder_Bay` |
| `America/Tijuana` | `America/Tijuana` |
| `America/Toronto` | `America/Toronto` |
| `America/Tortola` | `America/Tortola` |
| `America/Vancouver` | `America/Vancouver` |
| `America/Virgin` | `America/Virgin` |
| `America/Whitehorse` | `America/Whitehorse` |
| `America/Winnipeg` | `America/Winnipeg` |
| `America/Yakutat` | `America/Yakutat` |
| `America/Yellowknife` | `America/Yellowknife` |
| `Antarctica/Casey` | `Antarctica/Casey` |
| `Antarctica/Davis` | `Antarctica/Davis` |
| `Antarctica/DumontDUrville` | `Antarctica/DumontDUrville` |
| `Antarctica/Macquarie` | `Antarctica/Macquarie` |
| `Antarctica/Mawson` | `Antarctica/Mawson` |
| `Antarctica/McMurdo` | `Antarctica/McMurdo` |
| `Antarctica/Palmer` | `Antarctica/Palmer` |
| `Antarctica/Rothera` | `Antarctica/Rothera` |
| `Antarctica/South_Pole` | `Antarctica/South_Pole` |
| `Antarctica/Syowa` | `Antarctica/Syowa` |
| `Antarctica/Troll` | `Antarctica/Troll` |
| `Antarctica/Vostok` | `Antarctica/Vostok` |
| `Arctic/Longyearbyen` | `Arctic/Longyearbyen` |
| `Asia/Aden` | `Asia/Aden` |
| `Asia/Almaty` | `Asia/Almaty` |
| `Asia/Amman` | `Asia/Amman` |
| `Asia/Anadyr` | `Asia/Anadyr` |
| `Asia/Aqtau` | `Asia/Aqtau` |
| `Asia/Aqtobe` | `Asia/Aqtobe` |
| `Asia/Ashgabat` | `Asia/Ashgabat` |
| `Asia/Ashkhabad` | `Asia/Ashkhabad` |
| `Asia/Baghdad` | `Asia/Baghdad` |
| `Asia/Bahrain` | `Asia/Bahrain` |
| `Asia/Baku` | `Asia/Baku` |
| `Asia/Bangkok` | `Asia/Bangkok` |
| `Asia/Beirut` | `Asia/Beirut` |
| `Asia/Bishkek` | `Asia/Bishkek` |
| `Asia/Brunei` | `Asia/Brunei` |
| `Asia/Calcutta` | `Asia/Calcutta` |
| `Asia/Chita` | `Asia/Chita` |
| `Asia/Choibalsan` | `Asia/Choibalsan` |
| `Asia/Chongqing` | `Asia/Chongqing` |
| `Asia/Chungking` | `Asia/Chungking` |
| `Asia/Colombo` | `Asia/Colombo` |
| `Asia/Dacca` | `Asia/Dacca` |
| `Asia/Damascus` | `Asia/Damascus` |
| `Asia/Dhaka` | `Asia/Dhaka` |
| `Asia/Dili` | `Asia/Dili` |
| `Asia/Dubai` | `Asia/Dubai` |
| `Asia/Dushanbe` | `Asia/Dushanbe` |
| `Asia/Gaza` | `Asia/Gaza` |
| `Asia/Harbin` | `Asia/Harbin` |
| `Asia/Hebron` | `Asia/Hebron` |
| `Asia/Ho_Chi_Minh` | `Asia/Ho_Chi_Minh` |
| `Asia/Hong_Kong` | `Asia/Hong_Kong` |
| `Asia/Hovd` | `Asia/Hovd` |
| `Asia/Irkutsk` | `Asia/Irkutsk` |
| `Asia/Istanbul` | `Asia/Istanbul` |
| `Asia/Jakarta` | `Asia/Jakarta` |
| `Asia/Jayapura` | `Asia/Jayapura` |
| `Asia/Jerusalem` | `Asia/Jerusalem` |
| `Asia/Kabul` | `Asia/Kabul` |
| `Asia/Kamchatka` | `Asia/Kamchatka` |
| `Asia/Karachi` | `Asia/Karachi` |
| `Asia/Kashgar` | `Asia/Kashgar` |
| `Asia/Kathmandu` | `Asia/Kathmandu` |
| `Asia/Katmandu` | `Asia/Katmandu` |
| `Asia/Khandyga` | `Asia/Khandyga` |
| `Asia/Kolkata` | `Asia/Kolkata` |
| `Asia/Krasnoyarsk` | `Asia/Krasnoyarsk` |
| `Asia/Kuala_Lumpur` | `Asia/Kuala_Lumpur` |
| `Asia/Kuching` | `Asia/Kuching` |
| `Asia/Kuwait` | `Asia/Kuwait` |
| `Asia/Macao` | `Asia/Macao` |
| `Asia/Macau` | `Asia/Macau` |
| `Asia/Magadan` | `Asia/Magadan` |
| `Asia/Makassar` | `Asia/Makassar` |
| `Asia/Manila` | `Asia/Manila` |
| `Asia/Muscat` | `Asia/Muscat` |
| `Asia/Nicosia` | `Asia/Nicosia` |
| `Asia/Novokuznetsk` | `Asia/Novokuznetsk` |
| `Asia/Novosibirsk` | `Asia/Novosibirsk` |
| `Asia/Omsk` | `Asia/Omsk` |
| `Asia/Oral` | `Asia/Oral` |
| `Asia/Phnom_Penh` | `Asia/Phnom_Penh` |
| `Asia/Pontianak` | `Asia/Pontianak` |
| `Asia/Pyongyang` | `Asia/Pyongyang` |
| `Asia/Qatar` | `Asia/Qatar` |
| `Asia/Qyzylorda` | `Asia/Qyzylorda` |
| `Asia/Rangoon` | `Asia/Rangoon` |
| `Asia/Riyadh` | `Asia/Riyadh` |
| `Asia/Saigon` | `Asia/Saigon` |
| `Asia/Sakhalin` | `Asia/Sakhalin` |
| `Asia/Samarkand` | `Asia/Samarkand` |
| `Asia/Seoul` | `Asia/Seoul` |
| `Asia/Shanghai` | `Asia/Shanghai` |
| `Asia/Singapore` | `Asia/Singapore` |
| `Asia/Srednekolymsk` | `Asia/Srednekolymsk` |
| `Asia/Taipei` | `Asia/Taipei` |
| `Asia/Tashkent` | `Asia/Tashkent` |
| `Asia/Tbilisi` | `Asia/Tbilisi` |
| `Asia/Tehran` | `Asia/Tehran` |
| `Asia/Tel_Aviv` | `Asia/Tel_Aviv` |
| `Asia/Thimbu` | `Asia/Thimbu` |
| `Asia/Thimphu` | `Asia/Thimphu` |
| `Asia/Tokyo` | `Asia/Tokyo` |
| `Asia/Ujung_Pandang` | `Asia/Ujung_Pandang` |
| `Asia/Ulaanbaatar` | `Asia/Ulaanbaatar` |
| `Asia/Ulan_Bator` | `Asia/Ulan_Bator` |
| `Asia/Urumqi` | `Asia/Urumqi` |
| `Asia/Ust-Nera` | `Asia/Ust-Nera` |
| `Asia/Vientiane` | `Asia/Vientiane` |
| `Asia/Vladivostok` | `Asia/Vladivostok` |
| `Asia/Yakutsk` | `Asia/Yakutsk` |
| `Asia/Yekaterinburg` | `Asia/Yekaterinburg` |
| `Asia/Yerevan` | `Asia/Yerevan` |
| `Atlantic/Azores` | `Atlantic/Azores` |
| `Atlantic/Bermuda` | `Atlantic/Bermuda` |
| `Atlantic/Canary` | `Atlantic/Canary` |
| `Atlantic/Cape_Verde` | `Atlantic/Cape_Verde` |
| `Atlantic/Faeroe` | `Atlantic/Faeroe` |
| `Atlantic/Faroe` | `Atlantic/Faroe` |
| `Atlantic/Jan_Mayen` | `Atlantic/Jan_Mayen` |
| `Atlantic/Madeira` | `Atlantic/Madeira` |
| `Atlantic/Reykjavik` | `Atlantic/Reykjavik` |
| `Atlantic/South_Georgia` | `Atlantic/South_Georgia` |
| `Atlantic/St_Helena` | `Atlantic/St_Helena` |
| `Atlantic/Stanley` | `Atlantic/Stanley` |
| `Australia/ACT` | `Australia/ACT` |
| `Australia/Adelaide` | `Australia/Adelaide` |
| `Australia/Brisbane` | `Australia/Brisbane` |
| `Australia/Broken_Hill` | `Australia/Broken_Hill` |
| `Australia/Canberra` | `Australia/Canberra` |
| `Australia/Currie` | `Australia/Currie` |
| `Australia/Darwin` | `Australia/Darwin` |
| `Australia/Eucla` | `Australia/Eucla` |
| `Australia/Hobart` | `Australia/Hobart` |
| `Australia/LHI` | `Australia/LHI` |
| `Australia/Lindeman` | `Australia/Lindeman` |
| `Australia/Lord_Howe` | `Australia/Lord_Howe` |
| `Australia/Melbourne` | `Australia/Melbourne` |
| `Australia/NSW` | `Australia/NSW` |
| `Australia/North` | `Australia/North` |
| `Australia/Perth` | `Australia/Perth` |
| `Australia/Queensland` | `Australia/Queensland` |
| `Australia/South` | `Australia/South` |
| `Australia/Sydney` | `Australia/Sydney` |
| `Australia/Tasmania` | `Australia/Tasmania` |
| `Australia/Victoria` | `Australia/Victoria` |
| `Australia/West` | `Australia/West` |
| `Australia/Yancowinna` | `Australia/Yancowinna` |
| `Brazil/Acre` | `Brazil/Acre` |
| `Brazil/DeNoronha` | `Brazil/DeNoronha` |
| `Brazil/East` | `Brazil/East` |
| `Brazil/West` | `Brazil/West` |
| `CET` | `CET` |
| `CST6CDT` | `CST6CDT` |
| `Canada/Atlantic` | `Canada/Atlantic` |
| `Canada/Central` | `Canada/Central` |
| `Canada/East-Saskatchewan` | `Canada/East-Saskatchewan` |
| `Canada/Eastern` | `Canada/Eastern` |
| `Canada/Mountain` | `Canada/Mountain` |
| `Canada/Newfoundland` | `Canada/Newfoundland` |
| `Canada/Pacific` | `Canada/Pacific` |
| `Canada/Saskatchewan` | `Canada/Saskatchewan` |
| `Canada/Yukon` | `Canada/Yukon` |
| `Chile/Continental` | `Chile/Continental` |
| `Chile/EasterIsland` | `Chile/EasterIsland` |
| `Cuba` | `Cuba` |
| `EET` | `EET` |
| `EST` | `EST` |
| `EST5EDT` | `EST5EDT` |
| `Egypt` | `Egypt` |
| `Eire` | `Eire` |
| `Etc/GMT` | `Etc/GMT` |
| `Etc/GMT+0` | `Etc/GMT+0` |
| `Etc/GMT+1` | `Etc/GMT+1` |
| `Etc/GMT+10` | `Etc/GMT+10` |
| `Etc/GMT+11` | `Etc/GMT+11` |
| `Etc/GMT+12` | `Etc/GMT+12` |
| `Etc/GMT+2` | `Etc/GMT+2` |
| `Etc/GMT+3` | `Etc/GMT+3` |
| `Etc/GMT+4` | `Etc/GMT+4` |
| `Etc/GMT+5` | `Etc/GMT+5` |
| `Etc/GMT+6` | `Etc/GMT+6` |
| `Etc/GMT+7` | `Etc/GMT+7` |
| `Etc/GMT+8` | `Etc/GMT+8` |
| `Etc/GMT+9` | `Etc/GMT+9` |
| `Etc/GMT-0` | `Etc/GMT-0` |
| `Etc/GMT-1` | `Etc/GMT-1` |
| `Etc/GMT-10` | `Etc/GMT-10` |
| `Etc/GMT-11` | `Etc/GMT-11` |
| `Etc/GMT-12` | `Etc/GMT-12` |
| `Etc/GMT-13` | `Etc/GMT-13` |
| `Etc/GMT-14` | `Etc/GMT-14` |
| `Etc/GMT-2` | `Etc/GMT-2` |
| `Etc/GMT-3` | `Etc/GMT-3` |
| `Etc/GMT-4` | `Etc/GMT-4` |
| `Etc/GMT-5` | `Etc/GMT-5` |
| `Etc/GMT-6` | `Etc/GMT-6` |
| `Etc/GMT-7` | `Etc/GMT-7` |
| `Etc/GMT-8` | `Etc/GMT-8` |
| `Etc/GMT-9` | `Etc/GMT-9` |
| `Etc/GMT0` | `Etc/GMT0` |
| `Etc/Greenwich` | `Etc/Greenwich` |
| `Etc/UCT` | `Etc/UCT` |
| `Etc/UTC` | `Etc/UTC` |
| `Etc/Universal` | `Etc/Universal` |
| `Etc/Zulu` | `Etc/Zulu` |
| `Europe/Amsterdam` | `Europe/Amsterdam` |
| `Europe/Andorra` | `Europe/Andorra` |
| `Europe/Athens` | `Europe/Athens` |
| `Europe/Belfast` | `Europe/Belfast` |
| `Europe/Belgrade` | `Europe/Belgrade` |
| `Europe/Berlin` | `Europe/Berlin` |
| `Europe/Bratislava` | `Europe/Bratislava` |
| `Europe/Brussels` | `Europe/Brussels` |
| `Europe/Bucharest` | `Europe/Bucharest` |
| `Europe/Budapest` | `Europe/Budapest` |
| `Europe/Busingen` | `Europe/Busingen` |
| `Europe/Chisinau` | `Europe/Chisinau` |
| `Europe/Copenhagen` | `Europe/Copenhagen` |
| `Europe/Dublin` | `Europe/Dublin` |
| `Europe/Gibraltar` | `Europe/Gibraltar` |
| `Europe/Guernsey` | `Europe/Guernsey` |
| `Europe/Helsinki` | `Europe/Helsinki` |
| `Europe/Isle_of_Man` | `Europe/Isle_of_Man` |
| `Europe/Istanbul` | `Europe/Istanbul` |
| `Europe/Jersey` | `Europe/Jersey` |
| `Europe/Kaliningrad` | `Europe/Kaliningrad` |
| `Europe/Kiev` | `Europe/Kiev` |
| `Europe/Lisbon` | `Europe/Lisbon` |
| `Europe/Ljubljana` | `Europe/Ljubljana` |
| `Europe/London` | `Europe/London` |
| `Europe/Luxembourg` | `Europe/Luxembourg` |
| `Europe/Madrid` | `Europe/Madrid` |
| `Europe/Malta` | `Europe/Malta` |
| `Europe/Mariehamn` | `Europe/Mariehamn` |
| `Europe/Minsk` | `Europe/Minsk` |
| `Europe/Monaco` | `Europe/Monaco` |
| `Europe/Moscow` | `Europe/Moscow` |
| `Europe/Nicosia` | `Europe/Nicosia` |
| `Europe/Oslo` | `Europe/Oslo` |
| `Europe/Paris` | `Europe/Paris` |
| `Europe/Podgorica` | `Europe/Podgorica` |
| `Europe/Prague` | `Europe/Prague` |
| `Europe/Riga` | `Europe/Riga` |
| `Europe/Rome` | `Europe/Rome` |
| `Europe/Samara` | `Europe/Samara` |
| `Europe/San_Marino` | `Europe/San_Marino` |
| `Europe/Sarajevo` | `Europe/Sarajevo` |
| `Europe/Simferopol` | `Europe/Simferopol` |
| `Europe/Skopje` | `Europe/Skopje` |
| `Europe/Sofia` | `Europe/Sofia` |
| `Europe/Stockholm` | `Europe/Stockholm` |
| `Europe/Tallinn` | `Europe/Tallinn` |
| `Europe/Tirane` | `Europe/Tirane` |
| `Europe/Tiraspol` | `Europe/Tiraspol` |
| `Europe/Uzhgorod` | `Europe/Uzhgorod` |
| `Europe/Vaduz` | `Europe/Vaduz` |
| `Europe/Vatican` | `Europe/Vatican` |
| `Europe/Vienna` | `Europe/Vienna` |
| `Europe/Vilnius` | `Europe/Vilnius` |
| `Europe/Volgograd` | `Europe/Volgograd` |
| `Europe/Warsaw` | `Europe/Warsaw` |
| `Europe/Zagreb` | `Europe/Zagreb` |
| `Europe/Zaporozhye` | `Europe/Zaporozhye` |
| `Europe/Zurich` | `Europe/Zurich` |
| `GB` | `GB` |
| `GB-Eire` | `GB-Eire` |
| `GMT` | `GMT` |
| `GMT+0` | `GMT+0` |
| `GMT-0` | `GMT-0` |
| `GMT0` | `GMT0` |
| `Greenwich` | `Greenwich` |
| `HST` | `HST` |
| `Hongkong` | `Hongkong` |
| `Iceland` | `Iceland` |
| `Indian/Antananarivo` | `Indian/Antananarivo` |
| `Indian/Chagos` | `Indian/Chagos` |
| `Indian/Christmas` | `Indian/Christmas` |
| `Indian/Cocos` | `Indian/Cocos` |
| `Indian/Comoro` | `Indian/Comoro` |
| `Indian/Kerguelen` | `Indian/Kerguelen` |
| `Indian/Mahe` | `Indian/Mahe` |
| `Indian/Maldives` | `Indian/Maldives` |
| `Indian/Mauritius` | `Indian/Mauritius` |
| `Indian/Mayotte` | `Indian/Mayotte` |
| `Indian/Reunion` | `Indian/Reunion` |
| `Iran` | `Iran` |
| `Israel` | `Israel` |
| `Jamaica` | `Jamaica` |
| `Japan` | `Japan` |
| `Kwajalein` | `Kwajalein` |
| `Libya` | `Libya` |
| `MET` | `MET` |
| `MST` | `MST` |
| `MST7MDT` | `MST7MDT` |
| `Mexico/BajaNorte` | `Mexico/BajaNorte` |
| `Mexico/BajaSur` | `Mexico/BajaSur` |
| `Mexico/General` | `Mexico/General` |
| `NZ` | `NZ` |
| `NZ-CHAT` | `NZ-CHAT` |
| `Navajo` | `Navajo` |
| `PRC` | `PRC` |
| `PST8PDT` | `PST8PDT` |
| `Pacific/Apia` | `Pacific/Apia` |
| `Pacific/Auckland` | `Pacific/Auckland` |
| `Pacific/Bougainville` | `Pacific/Bougainville` |
| `Pacific/Chatham` | `Pacific/Chatham` |
| `Pacific/Chuuk` | `Pacific/Chuuk` |
| `Pacific/Easter` | `Pacific/Easter` |
| `Pacific/Efate` | `Pacific/Efate` |
| `Pacific/Enderbury` | `Pacific/Enderbury` |
| `Pacific/Fakaofo` | `Pacific/Fakaofo` |
| `Pacific/Fiji` | `Pacific/Fiji` |
| `Pacific/Funafuti` | `Pacific/Funafuti` |
| `Pacific/Galapagos` | `Pacific/Galapagos` |
| `Pacific/Gambier` | `Pacific/Gambier` |
| `Pacific/Guadalcanal` | `Pacific/Guadalcanal` |
| `Pacific/Guam` | `Pacific/Guam` |
| `Pacific/Honolulu` | `Pacific/Honolulu` |
| `Pacific/Johnston` | `Pacific/Johnston` |
| `Pacific/Kiritimati` | `Pacific/Kiritimati` |
| `Pacific/Kosrae` | `Pacific/Kosrae` |
| `Pacific/Kwajalein` | `Pacific/Kwajalein` |
| `Pacific/Majuro` | `Pacific/Majuro` |
| `Pacific/Marquesas` | `Pacific/Marquesas` |
| `Pacific/Midway` | `Pacific/Midway` |
| `Pacific/Nauru` | `Pacific/Nauru` |
| `Pacific/Niue` | `Pacific/Niue` |
| `Pacific/Norfolk` | `Pacific/Norfolk` |
| `Pacific/Noumea` | `Pacific/Noumea` |
| `Pacific/Pago_Pago` | `Pacific/Pago_Pago` |
| `Pacific/Palau` | `Pacific/Palau` |
| `Pacific/Pitcairn` | `Pacific/Pitcairn` |
| `Pacific/Pohnpei` | `Pacific/Pohnpei` |
| `Pacific/Ponape` | `Pacific/Ponape` |
| `Pacific/Port_Moresby` | `Pacific/Port_Moresby` |
| `Pacific/Rarotonga` | `Pacific/Rarotonga` |
| `Pacific/Saipan` | `Pacific/Saipan` |
| `Pacific/Samoa` | `Pacific/Samoa` |
| `Pacific/Tahiti` | `Pacific/Tahiti` |
| `Pacific/Tarawa` | `Pacific/Tarawa` |
| `Pacific/Tongatapu` | `Pacific/Tongatapu` |
| `Pacific/Truk` | `Pacific/Truk` |
| `Pacific/Wake` | `Pacific/Wake` |
| `Pacific/Wallis` | `Pacific/Wallis` |
| `Pacific/Yap` | `Pacific/Yap` |
| `Poland` | `Poland` |
| `Portugal` | `Portugal` |
| `ROC` | `ROC` |
| `ROK` | `ROK` |
| `Singapore` | `Singapore` |
| `Turkey` | `Turkey` |
| `UCT` | `UCT` |
| `US/Alaska` | `US/Alaska` |
| `US/Aleutian` | `US/Aleutian` |
| `US/Arizona` | `US/Arizona` |
| `US/Central` | `US/Central` |
| `US/East-Indiana` | `US/East-Indiana` |
| `US/Eastern` | `US/Eastern` |
| `US/Hawaii` | `US/Hawaii` |
| `US/Indiana-Starke` | `US/Indiana-Starke` |
| `US/Michigan` | `US/Michigan` |
| `US/Mountain` | `US/Mountain` |
| `US/Pacific` | `US/Pacific` |
| `US/Pacific-New` | `US/Pacific-New` |
| `US/Samoa` | `US/Samoa` |
| `UTC` | `UTC` |
| `Universal` | `Universal` |
| `W-SU` | `W-SU` |
| `WET` | `WET` |
| `Zulu` | `Zulu` |

## AccountLanguage

| Name | Stored as |
| --- | --- |
| `Afrikaans` | `af` |
| `العربيّة` | `ar` |
| `asturian` | `ast` |
| `Azərbaycanca` | `az` |
| `български` | `bg` |
| `беларуская` | `be` |
| `বাংলা` | `bn` |
| `brezhoneg` | `br` |
| `bosanski` | `bs` |
| `català` | `ca` |
| `česky` | `cs` |
| `Cymraeg` | `cy` |
| `dansk` | `da` |
| `Deutsch` | `de` |
| `Ελληνικά` | `el` |
| `English` | `en` |
| `Australian English` | `en-au` |
| `British English` | `en-gb` |
| `Esperanto` | `eo` |
| `español` | `es` |
| `español de Argentina` | `es-ar` |
| `español de Mexico` | `es-mx` |
| `español de Nicaragua` | `es-ni` |
| `español de Venezuela` | `es-ve` |
| `eesti` | `et` |
| `Basque` | `eu` |
| `فارسی` | `fa` |
| `suomi` | `fi` |
| `français` | `fr` |
| `frysk` | `fy` |
| `Gaeilge` | `ga` |
| `galego` | `gl` |
| `עברית` | `he` |
| `Hindi` | `hi` |
| `Hrvatski` | `hr` |
| `Magyar` | `hu` |
| `Interlingua` | `ia` |
| `Bahasa Indonesia` | `id` |
| `ido` | `io` |
| `Íslenska` | `is` |
| `italiano` | `it` |
| `日本語` | `ja` |
| `ქართული` | `ka` |
| `Қазақ` | `kk` |
| `Khmer` | `km` |
| `Kannada` | `kn` |
| `한국어` | `ko` |
| `Lëtzebuergesch` | `lb` |
| `Lietuviškai` | `lt` |
| `latvieš` | `lv` |
| `Македонски` | `mk` |
| `Malayalam` | `ml` |
| `Mongolian` | `mn` |
| `मराठी` | `mr` |
| `မြန်မာဘာသာ` | `my` |
| `norsk (bokmål)` | `nb` |
| `नेपाली` | `ne` |
| `Nederlands` | `nl` |
| `norsk (nynorsk)` | `nn` |
| `Ирон` | `os` |
| `Punjabi` | `pa` |
| `polski` | `pl` |
| `Português` | `pt` |
| `Português Brasileiro` | `pt-br` |
| `Română` | `ro` |
| `Русский` | `ru` |
| `slovenský` | `sk` |
| `Slovenščina` | `sl` |
| `shqip` | `sq` |
| `српски` | `sr` |
| `srpski (latinica)` | `sr-latn` |
| `svenska` | `sv` |
| `Kiswahili` | `sw` |
| `தமிழ்` | `ta` |
| `తెలుగు` | `te` |
| `ภาษาไทย` | `th` |
| `Türkçe` | `tr` |
| `Татарча` | `tt` |
| `Удмурт` | `udm` |
| `Українська` | `uk` |
| `اردو` | `ur` |
| `Tiếng Việt` | `vi` |
| `简体中文` | `zh-cn` |
| `简体中文` | `zh-hans` |
| `繁體中文` | `zh-hant` |
| `繁體中文` | `zh-tw` |

## AnnouncementDismissalType

| Name | Stored as |
| --- | --- |
| `No Dismissals Allowed` | `1` |
| `Session Only Dismissal` | `2` |
| `Permanent Dismissal Allowed` | `3` |

## JoinInvitationStatus

| Name | Stored as |
| --- | --- |
| `Sent` | `1` |
| `Accepted` | `2` |
| `Joined Independently` | `3` |

## SimpleTeamMemberAccess

| Name | Stored as |
| --- | --- |
| `open` | `open` |
| `by application` | `application` |
| `by invitation` | `invitation` |

## SimpleTeamManagerAccess

| Name | Stored as |
| --- | --- |
| `add someone` | `add someone` |
| `invite someone` | `invite someone` |

## TeamMemberAccess

| Name | Stored as |
| --- | --- |
| `open` | `open` |
| `by application` | `application` |
| `by invitation` | `invitation` |

## TeamManagerAccess

| Name | Stored as |
| --- | --- |
| `add someone` | `add someone` |
| `invite someone` | `invite someone` |

## SimpleMembershipState

| Name | Stored as |
| --- | --- |
| `applied` | `applied` |
| `invited` | `invited` |
| `declined` | `declined` |
| `rejected` | `rejected` |
| `accepted` | `accepted` |
| `waitlisted` | `waitlisted` |
| `auto joined` | `auto-joined` |

## SimpleMembershipRole

| Name | Stored as |
| --- | --- |
| `member` | `member` |
| `manager` | `manager` |
| `owner` | `owner` |

## MembershipState

| Name | Stored as |
| --- | --- |
| `applied` | `applied` |
| `invited` | `invited` |
| `declined` | `declined` |
| `rejected` | `rejected` |
| `accepted` | `accepted` |
| `waitlisted` | `waitlisted` |
| `auto joined` | `auto-joined` |

## MembershipRole

| Name | Stored as |
| --- | --- |
| `member` | `member` |
| `manager` | `manager` |
| `owner` | `owner` |
