# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Theme](#theme)

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

CREATE TABLE "admin_interface_theme" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(50) NOT NULL UNIQUE, "active" bool NOT NULL, "title" varchar(50) NOT NULL, "title_color" varchar(10) NOT NULL, "title_visible" bool NOT NULL, "logo" varchar(100) NOT NULL, "logo_color" varchar(10) NOT NULL, "logo_max_width" smallint unsigned NOT NULL CHECK ("logo_max_width" >= 0), "logo_max_height" smallint unsigned NOT NULL CHECK ("logo_max_height" >= 0), "logo_visible" bool NOT NULL, "favicon" varchar(100) NOT NULL, "env_name" varchar(50) NOT NULL, "env_color" varchar(10) NOT NULL, "env_visible_in_header" bool NOT NULL, "env_visible_in_favicon" bool NOT NULL, "language_chooser_active" bool NOT NULL, "language_chooser_control" varchar(20) NOT NULL, "language_chooser_display" varchar(10) NOT NULL, "css_header_background_color" varchar(10) NOT NULL, "css_header_text_color" varchar(10) NOT NULL, "css_header_link_color" varchar(10) NOT NULL, "css_header_link_hover_color" varchar(10) NOT NULL, "css_module_background_color" varchar(10) NOT NULL, "css_module_background_selected_color" varchar(10) NOT NULL, "css_module_text_color" varchar(10) NOT NULL, "css_module_link_color" varchar(10) NOT NULL, "css_module_link_selected_color" varchar(10) NOT NULL, "css_module_link_hover_color" varchar(10) NOT NULL, "css_module_rounded_corners" bool NOT NULL, "css_generic_link_color" varchar(10) NOT NULL, "css_generic_link_hover_color" varchar(10) NOT NULL, "css_generic_link_active_color" varchar(10) NOT NULL, "css_save_button_background_color" varchar(10) NOT NULL, "css_save_button_background_hover_color" varchar(10) NOT NULL, "css_save_button_text_color" varchar(10) NOT NULL, "css_delete_button_background_color" varchar(10) NOT NULL, "css_delete_button_background_hover_color" varchar(10) NOT NULL, "css_delete_button_text_color" varchar(10) NOT NULL, "related_modal_active" bool NOT NULL, "related_modal_background_color" varchar(10) NOT NULL, "related_modal_background_opacity" varchar(5) NOT NULL, "related_modal_rounded_corners" bool NOT NULL, "related_modal_close_button_visible" bool NOT NULL, "list_filter_highlight" bool NOT NULL, "list_filter_dropdown" bool NOT NULL, "list_filter_sticky" bool NOT NULL, "list_filter_removal_links" bool NOT NULL, "foldable_apps" bool NOT NULL, "show_fieldsets_as_tabs" bool NOT NULL, "show_inlines_as_tabs" bool NOT NULL, "collapsible_stacked_inlines" bool NOT NULL, "collapsible_stacked_inlines_collapsed" bool NOT NULL, "collapsible_tabular_inlines" bool NOT NULL, "collapsible_tabular_inlines_collapsed" bool NOT NULL, "recent_actions_visible" bool NOT NULL, "form_actions_sticky" bool NOT NULL, "form_submit_sticky" bool NOT NULL, "form_pagination_sticky" bool NOT NULL);

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

## Theme

Table `admin_interface_theme`, app `admin_interface`.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(50)` |  | unique | `'Django'` |  |
| `active` | `BooleanField` | `bool` |  |  | `True` |  |
| `title` | `CharField` | `varchar(50)` |  |  | `'Django administration'` |  |
| `title_color` | `ColorField` | `varchar(10)` |  |  | `'#F5DD5D'` | #F5DD5D |
| `title_visible` | `BooleanField` | `bool` |  |  | `True` |  |
| `logo` | `FileField` | `varchar(100)` |  |  |  | Leave blank to use the default Django logo |
| `logo_color` | `ColorField` | `varchar(10)` |  |  | `'#FFFFFF'` | #FFFFFF |
| `logo_max_width` | `PositiveSmallIntegerField` | `smallint unsigned` |  |  | `400` |  |
| `logo_max_height` | `PositiveSmallIntegerField` | `smallint unsigned` |  |  | `100` |  |
| `logo_visible` | `BooleanField` | `bool` |  |  | `True` |  |
| `favicon` | `FileField` | `varchar(100)` |  |  |  | (.ico\|.png\|.gif - 16x16\|32x32 px) |
| `env_name` | `CharField` | `varchar(50)` |  |  |  |  |
| `env_color` | `ColorField` | `varchar(10)` |  |  | `'#E74C3C'` | (red: #E74C3C, orange: #E67E22, yellow: #F1C40F, green: #2ECC71, blue: #3498DB) |
| `env_visible_in_header` | `BooleanField` | `bool` |  |  | `True` |  |
| `env_visible_in_favicon` | `BooleanField` | `bool` |  |  | `True` |  |
| `language_chooser_active` | `BooleanField` | `bool` |  |  | `True` |  |
| `language_chooser_control` | `ThemeLanguageChooserControl` | `varchar(20)` |  |  | `'default-select'` |  |
| `language_chooser_display` | `ThemeLanguageChooserDisplay` | `varchar(10)` |  |  | `'code'` |  |
| `css_header_background_color` | `ColorField` | `varchar(10)` |  |  | `'#0C4B33'` | #0C4B33 |
| `css_header_text_color` | `ColorField` | `varchar(10)` |  |  | `'#44B78B'` | #44B78B |
| `css_header_link_color` | `ColorField` | `varchar(10)` |  |  | `'#FFFFFF'` | #FFFFFF |
| `css_header_link_hover_color` | `ColorField` | `varchar(10)` |  |  | `'#C9F0DD'` | #C9F0DD |
| `css_module_background_color` | `ColorField` | `varchar(10)` |  |  | `'#44B78B'` | #44B78B |
| `css_module_background_selected_color` | `ColorField` | `varchar(10)` |  |  | `'#FFFFCC'` | #FFFFCC |
| `css_module_text_color` | `ColorField` | `varchar(10)` |  |  | `'#FFFFFF'` | #FFFFFF |
| `css_module_link_color` | `ColorField` | `varchar(10)` |  |  | `'#FFFFFF'` | #FFFFFF |
| `css_module_link_selected_color` | `ColorField` | `varchar(10)` |  |  | `'#FFFFFF'` | #FFFFFF |
| `css_module_link_hover_color` | `ColorField` | `varchar(10)` |  |  | `'#C9F0DD'` | #C9F0DD |
| `css_module_rounded_corners` | `BooleanField` | `bool` |  |  | `True` |  |
| `css_generic_link_color` | `ColorField` | `varchar(10)` |  |  | `'#0C3C26'` | #0C3C26 |
| `css_generic_link_hover_color` | `ColorField` | `varchar(10)` |  |  | `'#156641'` | #156641 |
| `css_generic_link_active_color` | `ColorField` | `varchar(10)` |  |  | `'#29B864'` | #29B864 |
| `css_save_button_background_color` | `ColorField` | `varchar(10)` |  |  | `'#0C4B33'` | #0C4B33 |
| `css_save_button_background_hover_color` | `ColorField` | `varchar(10)` |  |  | `'#0C3C26'` | #0C3C26 |
| `css_save_button_text_color` | `ColorField` | `varchar(10)` |  |  | `'#FFFFFF'` | #FFFFFF |
| `css_delete_button_background_color` | `ColorField` | `varchar(10)` |  |  | `'#BA2121'` | #BA2121 |
| `css_delete_button_background_hover_color` | `ColorField` | `varchar(10)` |  |  | `'#A41515'` | #A41515 |
| `css_delete_button_text_color` | `ColorField` | `varchar(10)` |  |  | `'#FFFFFF'` | #FFFFFF |
| `related_modal_active` | `BooleanField` | `bool` |  |  | `True` |  |
| `related_modal_background_color` | `ColorField` | `varchar(10)` |  |  | `'#000000'` | #000000 |
| `related_modal_background_opacity` | `ThemeRelatedModalBackgroundOpacity` | `varchar(5)` |  |  | `'0.3'` | 20% |
| `related_modal_rounded_corners` | `BooleanField` | `bool` |  |  | `True` |  |
| `related_modal_close_button_visible` | `BooleanField` | `bool` |  |  | `True` |  |
| `list_filter_highlight` | `BooleanField` | `bool` |  |  | `True` |  |
| `list_filter_dropdown` | `BooleanField` | `bool` |  |  | `True` |  |
| `list_filter_sticky` | `BooleanField` | `bool` |  |  | `True` |  |
| `list_filter_removal_links` | `BooleanField` | `bool` |  |  | `True` |  |
| `foldable_apps` | `BooleanField` | `bool` |  |  | `True` |  |
| `show_fieldsets_as_tabs` | `BooleanField` | `bool` |  |  | `False` |  |
| `show_inlines_as_tabs` | `BooleanField` | `bool` |  |  | `False` |  |
| `collapsible_stacked_inlines` | `BooleanField` | `bool` |  |  | `False` |  |
| `collapsible_stacked_inlines_collapsed` | `BooleanField` | `bool` |  |  | `True` |  |
| `collapsible_tabular_inlines` | `BooleanField` | `bool` |  |  | `False` |  |
| `collapsible_tabular_inlines_collapsed` | `BooleanField` | `bool` |  |  | `True` |  |
| `recent_actions_visible` | `BooleanField` | `bool` |  |  | `True` |  |
| `form_actions_sticky` | `BooleanField` | `bool` |  |  | `True` |  |
| `form_submit_sticky` | `BooleanField` | `bool` |  |  | `True` |  |
| `form_pagination_sticky` | `BooleanField` | `bool` |  |  | `True` |  |

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## ThemeLanguageChooserControl

| Name | Stored as |
| --- | --- |
| `Default Select` | `default-select` |
| `Minimal Select` | `minimal-select` |

## ThemeLanguageChooserDisplay

| Name | Stored as |
| --- | --- |
| `code` | `code` |
| `name` | `name` |

## ThemeRelatedModalBackgroundOpacity

| Name | Stored as |
| --- | --- |
| `10%` | `0.1` |
| `20%` | `0.2` |
| `30%` | `0.3` |
| `40%` | `0.4` |
| `50%` | `0.5` |
| `60%` | `0.6` |
| `70%` | `0.7` |
| `80%` | `0.8` |
| `90%` | `0.9` |
