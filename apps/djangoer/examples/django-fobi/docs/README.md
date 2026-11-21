# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [FormElement](#formelement) · [FormHandler](#formhandler) · [FormWizardHandler](#formwizardhandler) · [FormWizardEntry](#formwizardentry) · [FormEntry](#formentry) · [FormWizardFormEntry](#formwizardformentry) · [FormFieldsetEntry](#formfieldsetentry) · [FormElementEntry](#formelemententry) · [FormHandlerEntry](#formhandlerentry) · [FormWizardHandlerEntry](#formwizardhandlerentry) · [SavedFormDataEntry](#savedformdataentry) · [SavedFormWizardDataEntry](#savedformwizarddataentry)

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

CREATE TABLE "fobi_formelement" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "plugin_uid" varchar(255) NOT NULL UNIQUE);

CREATE TABLE "fobi_formelement_users" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "formelement_id" bigint NOT NULL REFERENCES "fobi_formelement" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "fobi_formelement_groups" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "formelement_id" bigint NOT NULL REFERENCES "fobi_formelement" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "fobi_formhandler" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "plugin_uid" varchar(255) NOT NULL UNIQUE);

CREATE TABLE "fobi_formhandler_users" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "formhandler_id" bigint NOT NULL REFERENCES "fobi_formhandler" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "fobi_formhandler_groups" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "formhandler_id" bigint NOT NULL REFERENCES "fobi_formhandler" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "fobi_formwizardhandler" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "plugin_uid" varchar(255) NOT NULL UNIQUE);

CREATE TABLE "fobi_formwizardhandler_users" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "formwizardhandler_id" bigint NOT NULL REFERENCES "fobi_formwizardhandler" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "fobi_formwizardhandler_groups" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "formwizardhandler_id" bigint NOT NULL REFERENCES "fobi_formwizardhandler" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "fobi_formwizardentry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(255) NOT NULL, "title" varchar(255) NULL, "slug" varchar(50) NOT NULL UNIQUE, "is_public" bool NOT NULL, "is_cloneable" bool NOT NULL, "success_page_title" varchar(255) NULL, "success_page_message" text NULL, "show_all_navigation_buttons" bool NOT NULL, "wizard_type" varchar(255) NOT NULL, "created" datetime NULL, "updated" datetime NULL);

CREATE TABLE "fobi_formentry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(255) NOT NULL, "title" varchar(255) NULL, "slug" varchar(50) NOT NULL UNIQUE, "is_public" bool NOT NULL, "active_date_from" datetime NULL, "active_date_to" datetime NULL, "inactive_page_title" varchar(255) NULL, "inactive_page_message" text NULL, "is_cloneable" bool NOT NULL, "success_page_title" varchar(255) NULL, "success_page_message" text NULL, "action" varchar(255) NULL, "created" datetime NULL, "updated" datetime NULL);

CREATE TABLE "fobi_formwizardformentry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "form_wizard_entry_id" bigint NOT NULL REFERENCES "fobi_formwizardentry" ("id") DEFERRABLE INITIALLY DEFERRED, "form_entry_id" bigint NOT NULL REFERENCES "fobi_formentry" ("id") DEFERRABLE INITIALLY DEFERRED, "position" integer unsigned NULL CHECK ("position" >= 0));

CREATE TABLE "fobi_formfieldsetentry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "form_entry_id" bigint NULL REFERENCES "fobi_formentry" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(255) NOT NULL, "is_repeatable" bool NOT NULL);

CREATE TABLE "fobi_formelemententry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "plugin_data" text NULL, "form_entry_id" bigint NOT NULL REFERENCES "fobi_formentry" ("id") DEFERRABLE INITIALLY DEFERRED, "plugin_uid" varchar(255) NOT NULL, "form_fieldset_entry_id" bigint NULL REFERENCES "fobi_formfieldsetentry" ("id") DEFERRABLE INITIALLY DEFERRED, "position" integer unsigned NULL CHECK ("position" >= 0));

CREATE TABLE "fobi_formhandlerentry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "plugin_data" text NULL, "form_entry_id" bigint NOT NULL REFERENCES "fobi_formentry" ("id") DEFERRABLE INITIALLY DEFERRED, "plugin_uid" varchar(255) NOT NULL);

CREATE TABLE "fobi_formwizardhandlerentry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "plugin_data" text NULL, "form_wizard_entry_id" bigint NOT NULL REFERENCES "fobi_formwizardentry" ("id") DEFERRABLE INITIALLY DEFERRED, "plugin_uid" varchar(255) NOT NULL);

CREATE TABLE "db_store_savedformdataentry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "form_data_headers" text NULL, "saved_data" text NULL, "created" datetime NOT NULL, "form_entry_id" bigint NULL REFERENCES "fobi_formentry" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "db_store_savedformwizarddataentry" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "form_data_headers" text NULL, "saved_data" text NULL, "created" datetime NOT NULL, "form_wizard_entry_id" bigint NULL REFERENCES "fobi_formwizardentry" ("id") DEFERRABLE INITIALLY DEFERRED);

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

CREATE UNIQUE INDEX "fobi_formelement_users_formelement_id_user_id_d703a088_uniq" ON "fobi_formelement_users" ("formelement_id", "user_id");

CREATE INDEX "fobi_formelement_users_formelement_id_87fde0e1" ON "fobi_formelement_users" ("formelement_id");

CREATE INDEX "fobi_formelement_users_user_id_d42ae322" ON "fobi_formelement_users" ("user_id");

CREATE UNIQUE INDEX "fobi_formelement_groups_formelement_id_group_id_b1f59b96_uniq" ON "fobi_formelement_groups" ("formelement_id", "group_id");

CREATE INDEX "fobi_formelement_groups_formelement_id_1a6d0aed" ON "fobi_formelement_groups" ("formelement_id");

CREATE INDEX "fobi_formelement_groups_group_id_e204d753" ON "fobi_formelement_groups" ("group_id");

CREATE UNIQUE INDEX "fobi_formhandler_users_formhandler_id_user_id_91993ff2_uniq" ON "fobi_formhandler_users" ("formhandler_id", "user_id");

CREATE INDEX "fobi_formhandler_users_formhandler_id_467d1fab" ON "fobi_formhandler_users" ("formhandler_id");

CREATE INDEX "fobi_formhandler_users_user_id_6901d013" ON "fobi_formhandler_users" ("user_id");

CREATE UNIQUE INDEX "fobi_formhandler_groups_formhandler_id_group_id_b774d9a5_uniq" ON "fobi_formhandler_groups" ("formhandler_id", "group_id");

CREATE INDEX "fobi_formhandler_groups_formhandler_id_d36cb512" ON "fobi_formhandler_groups" ("formhandler_id");

CREATE INDEX "fobi_formhandler_groups_group_id_4dd1f7b6" ON "fobi_formhandler_groups" ("group_id");

CREATE UNIQUE INDEX "fobi_formwizardhandler_users_formwizardhandler_id_user_id_03ec6d07_uniq" ON "fobi_formwizardhandler_users" ("formwizardhandler_id", "user_id");

CREATE INDEX "fobi_formwizardhandler_users_formwizardhandler_id_5ac41a2d" ON "fobi_formwizardhandler_users" ("formwizardhandler_id");

CREATE INDEX "fobi_formwizardhandler_users_user_id_101ba90e" ON "fobi_formwizardhandler_users" ("user_id");

CREATE UNIQUE INDEX "fobi_formwizardhandler_groups_formwizardhandler_id_group_id_dc81ca2d_uniq" ON "fobi_formwizardhandler_groups" ("formwizardhandler_id", "group_id");

CREATE INDEX "fobi_formwizardhandler_groups_formwizardhandler_id_f9e753a3" ON "fobi_formwizardhandler_groups" ("formwizardhandler_id");

CREATE INDEX "fobi_formwizardhandler_groups_group_id_071b1264" ON "fobi_formwizardhandler_groups" ("group_id");

CREATE UNIQUE INDEX "fobi_formwizardentry_user_id_slug_d03683d1_uniq" ON "fobi_formwizardentry" ("user_id", "slug");

CREATE UNIQUE INDEX "fobi_formwizardentry_user_id_name_e8dd542c_uniq" ON "fobi_formwizardentry" ("user_id", "name");

CREATE INDEX "fobi_formwizardentry_user_id_231dc66c" ON "fobi_formwizardentry" ("user_id");

CREATE UNIQUE INDEX "fobi_formentry_user_id_slug_de82881b_uniq" ON "fobi_formentry" ("user_id", "slug");

CREATE UNIQUE INDEX "fobi_formentry_user_id_name_1ed0a04a_uniq" ON "fobi_formentry" ("user_id", "name");

CREATE INDEX "fobi_formentry_user_id_dd71895a" ON "fobi_formentry" ("user_id");

CREATE UNIQUE INDEX "fobi_formwizardformentry_form_wizard_entry_id_form_entry_id_7ba3a3c9_uniq" ON "fobi_formwizardformentry" ("form_wizard_entry_id", "form_entry_id");

CREATE INDEX "fobi_formwizardformentry_form_wizard_entry_id_43f66eaf" ON "fobi_formwizardformentry" ("form_wizard_entry_id");

CREATE INDEX "fobi_formwizardformentry_form_entry_id_6b951198" ON "fobi_formwizardformentry" ("form_entry_id");

CREATE UNIQUE INDEX "fobi_formfieldsetentry_form_entry_id_name_dc693e98_uniq" ON "fobi_formfieldsetentry" ("form_entry_id", "name");

CREATE INDEX "fobi_formfieldsetentry_form_entry_id_6b3dcc22" ON "fobi_formfieldsetentry" ("form_entry_id");

CREATE INDEX "fobi_formelemententry_form_entry_id_fda4e235" ON "fobi_formelemententry" ("form_entry_id");

CREATE INDEX "fobi_formelemententry_form_fieldset_entry_id_2a8c5c7d" ON "fobi_formelemententry" ("form_fieldset_entry_id");

CREATE INDEX "fobi_formhandlerentry_form_entry_id_9a7dd382" ON "fobi_formhandlerentry" ("form_entry_id");

CREATE INDEX "fobi_formwizardhandlerentry_form_wizard_entry_id_ed113a50" ON "fobi_formwizardhandlerentry" ("form_wizard_entry_id");

CREATE INDEX "db_store_savedformdataentry_user_id_d6630c0b" ON "db_store_savedformdataentry" ("user_id");

CREATE INDEX "db_store_savedformdataentry_form_entry_id_e68ba302" ON "db_store_savedformdataentry" ("form_entry_id");

CREATE INDEX "db_store_savedformwizarddataentry_user_id_bcb2c0ec" ON "db_store_savedformwizarddataentry" ("user_id");

CREATE INDEX "db_store_savedformwizarddataentry_form_wizard_entry_id_1a42a01c" ON "db_store_savedformwizarddataentry" ("form_wizard_entry_id");
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
- `formelement_set` — many to many with [FormElement](#formelement)
- `formhandler_set` — many to many with [FormHandler](#formhandler)
- `formwizardhandler_set` — many to many with [FormWizardHandler](#formwizardhandler)

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
- `formelement_set` — many to many with [FormElement](#formelement)
- `formhandler_set` — many to many with [FormHandler](#formhandler)
- `formwizardhandler_set` — many to many with [FormWizardHandler](#formwizardhandler)
- [FormWizardEntry](#formwizardentry).`user_id` → `id`, on delete cascade
- [FormEntry](#formentry).`user_id` → `id`, on delete cascade
- [SavedFormDataEntry](#savedformdataentry).`user_id` → `id`, on delete cascade
- [SavedFormWizardDataEntry](#savedformwizarddataentry).`user_id` → `id`, on delete cascade

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

## FormElement

Table `fobi_formelement`, app `fobi`.

Form element.

Form field plugin. Used when ``fobi.settings.RESTRICT_PLUGIN_ACCESS``
is set to True.

:Properties:

    - `plugin_uid` (str): Plugin UID.
    - `users` (django.contrib.auth.models.User): White list of the users
      allowed to use the form element plugin.
    - `groups` (django.contrib.auth.models.Group): White list of the user
      groups allowed to use the form element plugin.

![FormElement and the models it relates to](tables/FormElement.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `plugin_uid` | `CharField` | `varchar(255)` |  | unique |  |  |

Relations:

- `users` — many to many with [User](#user)
- `groups` — many to many with [Group](#group)

## FormHandler

Table `fobi_formhandler`, app `fobi`.

Form handler plugin. Used when ``fobi.settings.RESTRICT_PLUGIN_ACCESS``
is set to True.

:Properties:

    - `plugin_uid` (str): Plugin UID.
    - `users` (django.contrib.auth.models.User): White list of the users
      allowed to use the form handler plugin.
    - `groups` (django.contrib.auth.models.Group): White list of the
      user groups allowed to use the form handler plugin.

![FormHandler and the models it relates to](tables/FormHandler.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `plugin_uid` | `CharField` | `varchar(255)` |  | unique |  |  |

Relations:

- `users` — many to many with [User](#user)
- `groups` — many to many with [Group](#group)

## FormWizardHandler

Table `fobi_formwizardhandler`, app `fobi`.

Form wizard handler plugin. Used when
``fobi.settings.RESTRICT_PLUGIN_ACCESS`` is set to True.

:Properties:

    - `plugin_uid` (str): Plugin UID.
    - `users` (django.contrib.auth.models.User): White list of the users
      allowed to use the form handler plugin.
    - `groups` (django.contrib.auth.models.Group): White list of the
      user groups allowed to use the form handler plugin.

![FormWizardHandler and the models it relates to](tables/FormWizardHandler.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `plugin_uid` | `CharField` | `varchar(255)` |  | unique |  |  |

Relations:

- `users` — many to many with [User](#user)
- `groups` — many to many with [Group](#group)

## FormWizardEntry

Table `fobi_formwizardentry`, app `fobi`.

Form wizard entry.

![FormWizardEntry and the models it relates to](tables/FormWizardEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `title` | `CharField` | `varchar(255)` | yes |  |  | Shown in templates if available. |
| `slug` | `AutoSlugField` | `varchar(50)` |  | unique |  |  |
| `is_public` | `BooleanField` | `bool` |  |  | `False` | Makes your form wizard visible to the public. |
| `is_cloneable` | `BooleanField` | `bool` |  |  | `False` | Makes your form wizard cloneable by other users. |
| `success_page_title` | `CharField` | `varchar(255)` | yes |  |  | Custom message title to display after valid form is submitted |
| `success_page_message` | `TextField` | `text` | yes |  |  | Custom message text to display after valid form is submitted |
| `show_all_navigation_buttons` | `BooleanField` | `bool` |  |  | `False` | Show all navigation buttons. |
| `wizard_type` | `FormWizardEntryWizardType` | `varchar(255)` |  |  | `'SessionWizardView'` | Type of the form wizard. |
| `created` | `DateTimeField` | `datetime` | yes |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` | yes |  | `now()` |  |

Indexes:

- index `fobi_formwizardentry_user_id_231dc66c` on `user_id`
- unique `fobi_formwizardentry_user_id_slug_d03683d1_uniq` on `user_id`, `slug`
- unique `fobi_formwizardentry_user_id_name_e8dd542c_uniq` on `user_id`, `name`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [FormWizardFormEntry](#formwizardformentry).`form_wizard_entry_id` → `id`, on delete cascade
- [FormWizardHandlerEntry](#formwizardhandlerentry).`form_wizard_entry_id` → `id`, on delete cascade
- [SavedFormWizardDataEntry](#savedformwizarddataentry).`form_wizard_entry_id` → `id`, on delete cascade

## FormEntry

Table `fobi_formentry`, app `fobi`.

Form entry.

![FormEntry and the models it relates to](tables/FormEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `title` | `CharField` | `varchar(255)` | yes |  |  | Shown in templates if available. |
| `slug` | `AutoSlugField` | `varchar(50)` |  | unique |  |  |
| `is_public` | `BooleanField` | `bool` |  |  | `False` | Makes your form visible to the public. |
| `active_date_from` | `DateTimeField` | `datetime` | yes |  |  | Date and time when the form becomes active in the format: 'YYYY-MM-DD HH:MM'. Leave it blank to activate immediately. |
| `active_date_to` | `DateTimeField` | `datetime` | yes |  |  | Date and time when the form becomes inactive in the format: 'YYYY-MM-DD HH:MM'. Leave it blank to keep active forever. |
| `inactive_page_title` | `CharField` | `varchar(255)` | yes |  |  | Custom message title to display if form is inactive. |
| `inactive_page_message` | `TextField` | `text` | yes |  |  | Custom message text to display if form is inactive. |
| `is_cloneable` | `BooleanField` | `bool` |  |  | `False` | Makes your form cloneable by other users. |
| `success_page_title` | `CharField` | `varchar(255)` | yes |  |  | Custom message title to display after valid form is submitted |
| `success_page_message` | `TextField` | `text` | yes |  |  | Custom message text to display after valid form is submitted |
| `action` | `CharField` | `varchar(255)` | yes |  |  | Custom form action; don't fill this field, unless really necessary. |
| `created` | `DateTimeField` | `datetime` | yes |  | `now()` |  |
| `updated` | `DateTimeField` | `datetime` | yes |  | `now()` |  |

Indexes:

- index `fobi_formentry_user_id_dd71895a` on `user_id`
- unique `fobi_formentry_user_id_slug_de82881b_uniq` on `user_id`, `slug`
- unique `fobi_formentry_user_id_name_1ed0a04a_uniq` on `user_id`, `name`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [FormWizardFormEntry](#formwizardformentry).`form_entry_id` → `id`, on delete cascade
- [FormFieldsetEntry](#formfieldsetentry).`form_entry_id` → `id`, on delete cascade
- [FormElementEntry](#formelemententry).`form_entry_id` → `id`, on delete cascade
- [FormHandlerEntry](#formhandlerentry).`form_entry_id` → `id`, on delete cascade
- [SavedFormDataEntry](#savedformdataentry).`form_entry_id` → `id`, on delete cascade

## FormWizardFormEntry

Table `fobi_formwizardformentry`, app `fobi`.

Form wizard form entry.

A coupling point between `FormWizardEntry` and `FormEntry`.

![FormWizardFormEntry and the models it relates to](tables/FormWizardFormEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `form_wizard_entry_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `form_entry_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `position` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |

Indexes:

- index `fobi_formwizardformentry_form_wizard_entry_id_43f66eaf` on `form_wizard_entry_id`
- index `fobi_formwizardformentry_form_entry_id_6b951198` on `form_entry_id`
- unique `fobi_formwizardformentry_form_wizard_entry_id_form_entry_id_7ba3a3c9_uniq` on `form_wizard_entry_id`, `form_entry_id`

Relations:

- `form_wizard_entry_id` → [FormWizardEntry](#formwizardentry).`id`, on delete cascade
- `form_entry_id` → [FormEntry](#formentry).`id`, on delete cascade

## FormFieldsetEntry

Table `fobi_formfieldsetentry`, app `fobi`.

Form fieldset entry.

![FormFieldsetEntry and the models it relates to](tables/FormFieldsetEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `form_entry_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `is_repeatable` | `BooleanField` | `bool` |  |  | `False` | Makes your form fieldset repeatable. |

Indexes:

- index `fobi_formfieldsetentry_form_entry_id_6b3dcc22` on `form_entry_id`
- unique `fobi_formfieldsetentry_form_entry_id_name_dc693e98_uniq` on `form_entry_id`, `name`

Relations:

- `form_entry_id` → [FormEntry](#formentry).`id`, on delete cascade
- [FormElementEntry](#formelemententry).`form_fieldset_entry_id` → `id`, on delete cascade

## FormElementEntry

Table `fobi_formelemententry`, app `fobi`.

Columns from `AbstractPluginEntry` (abstract): `form_entry_id`.

Columns from `BaseAbstractPluginEntry` (abstract): `plugin_data`.

Form field entry.

:Properties:

- `form` (fobi.models.FormEntry): Form to which the field plugin
  belongs to.
- `plugin_uid` (str): Plugin UID.
- `plugin_data` (str): JSON formatted string with plugin data.
- `form_fieldset_entry`: Fieldset.
- `position` (int): Entry position.

![FormElementEntry and the models it relates to](tables/FormElementEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `plugin_data` | `TextField` | `text` | yes |  |  |  |
| `form_entry_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `plugin_uid` | `CharField` | `varchar(255)` |  |  |  |  |
| `form_fieldset_entry_id` | `ForeignKey` | `bigint` | yes | FK |  |  |
| `position` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |

Indexes:

- index `fobi_formelemententry_form_entry_id_fda4e235` on `form_entry_id`
- index `fobi_formelemententry_form_fieldset_entry_id_2a8c5c7d` on `form_fieldset_entry_id`

Relations:

- `form_entry_id` → [FormEntry](#formentry).`id`, on delete cascade
- `form_fieldset_entry_id` → [FormFieldsetEntry](#formfieldsetentry).`id`, on delete cascade

## FormHandlerEntry

Table `fobi_formhandlerentry`, app `fobi`.

Columns from `AbstractPluginEntry` (abstract): `form_entry_id`.

Columns from `BaseAbstractPluginEntry` (abstract): `plugin_data`.

Form handler entry.

:Properties:

    - `form_entry` (fobi.models.FormEntry): Form to which the handler
      plugin belongs to.
    - `plugin_uid` (str): Plugin UID.
    - `plugin_data` (str): JSON formatted string with plugin data.

![FormHandlerEntry and the models it relates to](tables/FormHandlerEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `plugin_data` | `TextField` | `text` | yes |  |  |  |
| `form_entry_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `plugin_uid` | `CharField` | `varchar(255)` |  |  |  |  |

Indexes:

- index `fobi_formhandlerentry_form_entry_id_9a7dd382` on `form_entry_id`

Relations:

- `form_entry_id` → [FormEntry](#formentry).`id`, on delete cascade

## FormWizardHandlerEntry

Table `fobi_formwizardhandlerentry`, app `fobi`.

Columns from `AbstractFormWizardPluginEntry` (abstract): `form_wizard_entry_id`.

Columns from `BaseAbstractPluginEntry` (abstract): `plugin_data`.

Form wizard handler entry.

:Properties:

    - `form_wizard_entry` (fobi.models.FormWizardEntry): FormWizard to
      which the handler plugin belongs to.
    - `plugin_uid` (str): Plugin UID.
    - `plugin_data` (str): JSON formatted string with plugin data.

![FormWizardHandlerEntry and the models it relates to](tables/FormWizardHandlerEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `plugin_data` | `TextField` | `text` | yes |  |  |  |
| `form_wizard_entry_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `plugin_uid` | `CharField` | `varchar(255)` |  |  |  |  |

Indexes:

- index `fobi_formwizardhandlerentry_form_wizard_entry_id_ed113a50` on `form_wizard_entry_id`

Relations:

- `form_wizard_entry_id` → [FormWizardEntry](#formwizardentry).`id`, on delete cascade

## SavedFormDataEntry

Table `db_store_savedformdataentry`, app `fobi_contrib_plugins_form_handlers_db_store`.

Columns from `AbstractSavedFormDataEntry` (abstract): `user_id`, `form_data_headers`, `saved_data`, `created`.

Saved form data.

![SavedFormDataEntry and the models it relates to](tables/SavedFormDataEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `form_data_headers` | `TextField` | `text` | yes |  |  |  |
| `saved_data` | `TextField` | `text` | yes |  |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `form_entry_id` | `ForeignKey` | `bigint` | yes | FK |  |  |

Indexes:

- index `db_store_savedformdataentry_user_id_d6630c0b` on `user_id`
- index `db_store_savedformdataentry_form_entry_id_e68ba302` on `form_entry_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `form_entry_id` → [FormEntry](#formentry).`id`, on delete cascade

## SavedFormWizardDataEntry

Table `db_store_savedformwizarddataentry`, app `fobi_contrib_plugins_form_handlers_db_store`.

Columns from `AbstractSavedFormDataEntry` (abstract): `user_id`, `form_data_headers`, `saved_data`, `created`.

Saved form data.

![SavedFormWizardDataEntry and the models it relates to](tables/SavedFormWizardDataEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `form_data_headers` | `TextField` | `text` | yes |  |  |  |
| `saved_data` | `TextField` | `text` | yes |  |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `form_wizard_entry_id` | `ForeignKey` | `bigint` | yes | FK |  |  |

Indexes:

- index `db_store_savedformwizarddataentry_user_id_bcb2c0ec` on `user_id`
- index `db_store_savedformwizarddataentry_form_wizard_entry_id_1a42a01c` on `form_wizard_entry_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `form_wizard_entry_id` → [FormWizardEntry](#formwizardentry).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## FormWizardEntryWizardType

| Name | Stored as |
| --- | --- |
| `Session wizard` | `SessionWizardView` |
| `Cookie wizard` | `CookieWizardView` |
