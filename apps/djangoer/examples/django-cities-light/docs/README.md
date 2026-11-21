# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Country](#country) · [Region](#region) · [SubRegion](#subregion) · [City](#city)

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

CREATE TABLE "cities_light_country" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(200) NOT NULL, "name_ascii" varchar(200) NOT NULL, "slug" varchar(50) NOT NULL, "geoname_id" integer NULL UNIQUE, "alternate_names" text NULL, "translations" text NOT NULL CHECK ((JSON_VALID("translations") OR "translations" IS NULL)), "code2" varchar(2) NULL UNIQUE, "code3" varchar(3) NULL UNIQUE, "continent" varchar(2) NOT NULL, "tld" varchar(5) NOT NULL, "phone" varchar(20) NULL);

CREATE TABLE "cities_light_region" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(200) NOT NULL, "name_ascii" varchar(200) NOT NULL, "slug" varchar(50) NOT NULL, "geoname_id" integer NULL UNIQUE, "alternate_names" text NULL, "translations" text NOT NULL CHECK ((JSON_VALID("translations") OR "translations" IS NULL)), "display_name" varchar(200) NOT NULL, "geoname_code" varchar(50) NULL, "country_id" integer NOT NULL REFERENCES "cities_light_country" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "cities_light_subregion" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(200) NOT NULL, "name_ascii" varchar(200) NOT NULL, "slug" varchar(50) NOT NULL, "geoname_id" integer NULL UNIQUE, "alternate_names" text NULL, "translations" text NOT NULL CHECK ((JSON_VALID("translations") OR "translations" IS NULL)), "display_name" varchar(200) NOT NULL, "geoname_code" varchar(50) NULL, "country_id" integer NOT NULL REFERENCES "cities_light_country" ("id") DEFERRABLE INITIALLY DEFERRED, "region_id" integer NULL REFERENCES "cities_light_region" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "cities_light_city" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(200) NOT NULL, "name_ascii" varchar(200) NOT NULL, "slug" varchar(50) NOT NULL, "geoname_id" integer NULL UNIQUE, "alternate_names" text NULL, "translations" text NOT NULL CHECK ((JSON_VALID("translations") OR "translations" IS NULL)), "display_name" varchar(200) NOT NULL, "search_names" text NOT NULL, "latitude" decimal NULL, "longitude" decimal NULL, "subregion_id" integer NULL REFERENCES "cities_light_subregion" ("id") DEFERRABLE INITIALLY DEFERRED, "region_id" integer NULL REFERENCES "cities_light_region" ("id") DEFERRABLE INITIALLY DEFERRED, "country_id" integer NOT NULL REFERENCES "cities_light_country" ("id") DEFERRABLE INITIALLY DEFERRED, "population" bigint NULL, "feature_code" varchar(10) NULL, "timezone" varchar(40) NULL);

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

CREATE INDEX "cities_light_country_name_1d61d0d2" ON "cities_light_country" ("name");

CREATE INDEX "cities_light_country_name_ascii_648cc5e3" ON "cities_light_country" ("name_ascii");

CREATE INDEX "cities_light_country_slug_3707a22c" ON "cities_light_country" ("slug");

CREATE INDEX "cities_light_country_continent_e2c412a4" ON "cities_light_country" ("continent");

CREATE INDEX "cities_light_country_tld_1fb2faaa" ON "cities_light_country" ("tld");

CREATE UNIQUE INDEX "cities_light_region_country_id_name_6e5b3799_uniq" ON "cities_light_region" ("country_id", "name");

CREATE UNIQUE INDEX "cities_light_region_country_id_slug_3dd02103_uniq" ON "cities_light_region" ("country_id", "slug");

CREATE INDEX "cities_light_region_name_775f9496" ON "cities_light_region" ("name");

CREATE INDEX "cities_light_region_name_ascii_f085cb82" ON "cities_light_region" ("name_ascii");

CREATE INDEX "cities_light_region_slug_1653969f" ON "cities_light_region" ("slug");

CREATE INDEX "cities_light_region_geoname_code_1b0d4e58" ON "cities_light_region" ("geoname_code");

CREATE INDEX "cities_light_region_country_id_b2782d49" ON "cities_light_region" ("country_id");

CREATE INDEX "cities_light_subregion_name_2882337e" ON "cities_light_subregion" ("name");

CREATE INDEX "cities_light_subregion_name_ascii_ecf9a336" ON "cities_light_subregion" ("name_ascii");

CREATE INDEX "cities_light_subregion_slug_43494546" ON "cities_light_subregion" ("slug");

CREATE INDEX "cities_light_subregion_geoname_code_843acdc3" ON "cities_light_subregion" ("geoname_code");

CREATE INDEX "cities_light_subregion_country_id_9b32b484" ON "cities_light_subregion" ("country_id");

CREATE INDEX "cities_light_subregion_region_id_c6e0b71f" ON "cities_light_subregion" ("region_id");

CREATE UNIQUE INDEX "cities_light_city_region_id_subregion_id_name_cdfc77ea_uniq" ON "cities_light_city" ("region_id", "subregion_id", "name");

CREATE UNIQUE INDEX "cities_light_city_region_id_subregion_id_slug_efb2e768_uniq" ON "cities_light_city" ("region_id", "subregion_id", "slug");

CREATE INDEX "cities_light_city_name_4859e2a5" ON "cities_light_city" ("name");

CREATE INDEX "cities_light_city_name_ascii_1e56323b" ON "cities_light_city" ("name_ascii");

CREATE INDEX "cities_light_city_slug_cb2251f8" ON "cities_light_city" ("slug");

CREATE INDEX "cities_light_city_search_names_fb77fed2" ON "cities_light_city" ("search_names");

CREATE INDEX "cities_light_city_subregion_id_0926d2ad" ON "cities_light_city" ("subregion_id");

CREATE INDEX "cities_light_city_region_id_f7ab977b" ON "cities_light_city" ("region_id");

CREATE INDEX "cities_light_city_country_id_cf310fd2" ON "cities_light_city" ("country_id");

CREATE INDEX "cities_light_city_population_d597eeb6" ON "cities_light_city" ("population");

CREATE INDEX "cities_light_city_feature_code_d43c9217" ON "cities_light_city" ("feature_code");

CREATE INDEX "cities_light_city_timezone_0fb51b1e" ON "cities_light_city" ("timezone");
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

## Country

Table `cities_light_country`, app `cities_light`.

Columns from `AbstractCountry` (abstract): `code2`, `code3`, `continent`, `tld`, `phone`.

Columns from `Base` (abstract): `name`, `name_ascii`, `slug`, `geoname_id`, `alternate_names`, `translations`.

![Country and the models it relates to](tables/Country.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(200)` |  |  |  |  |
| `name_ascii` | `CharField` | `varchar(200)` |  |  |  |  |
| `slug` | `AutoSlugField` | `varchar(50)` |  |  |  |  |
| `geoname_id` | `IntegerField` | `integer` | yes | unique |  |  |
| `alternate_names` | `TextField` | `text` | yes |  | `''` |  |
| `translations` | `JSONField` | `text` |  |  | `dict()` |  |
| `code2` | `CharField` | `varchar(2)` | yes | unique |  |  |
| `code3` | `CharField` | `varchar(3)` | yes | unique |  |  |
| `continent` | `CountryContinent` | `varchar(2)` |  |  |  |  |
| `tld` | `CharField` | `varchar(5)` |  |  |  |  |
| `phone` | `CharField` | `varchar(20)` | yes |  |  |  |

Indexes:

- index `cities_light_country_name_1d61d0d2` on `name`
- index `cities_light_country_name_ascii_648cc5e3` on `name_ascii`
- index `cities_light_country_slug_3707a22c` on `slug`
- index `cities_light_country_continent_e2c412a4` on `continent`
- index `cities_light_country_tld_1fb2faaa` on `tld`

Relations:

- [Region](#region).`country_id` → `id`, on delete cascade
- [SubRegion](#subregion).`country_id` → `id`, on delete cascade
- [City](#city).`country_id` → `id`, on delete cascade

## Region

Table `cities_light_region`, app `cities_light`.

Columns from `AbstractRegion` (abstract): `display_name`, `geoname_code`, `country_id`.

Columns from `Base` (abstract): `name`, `name_ascii`, `slug`, `geoname_id`, `alternate_names`, `translations`.

![Region and the models it relates to](tables/Region.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(200)` |  |  |  |  |
| `name_ascii` | `CharField` | `varchar(200)` |  |  |  |  |
| `slug` | `AutoSlugField` | `varchar(50)` |  |  |  |  |
| `geoname_id` | `IntegerField` | `integer` | yes | unique |  |  |
| `alternate_names` | `TextField` | `text` | yes |  | `''` |  |
| `translations` | `JSONField` | `text` |  |  | `dict()` |  |
| `display_name` | `CharField` | `varchar(200)` |  |  |  |  |
| `geoname_code` | `CharField` | `varchar(50)` | yes |  |  |  |
| `country_id` | `ForeignKey` | `integer` |  | FK |  |  |

Indexes:

- index `cities_light_region_name_775f9496` on `name`
- index `cities_light_region_name_ascii_f085cb82` on `name_ascii`
- index `cities_light_region_slug_1653969f` on `slug`
- index `cities_light_region_geoname_code_1b0d4e58` on `geoname_code`
- index `cities_light_region_country_id_b2782d49` on `country_id`
- unique `cities_light_region_country_id_name_6e5b3799_uniq` on `country_id`, `name`
- unique `cities_light_region_country_id_slug_3dd02103_uniq` on `country_id`, `slug`

Relations:

- `country_id` → [Country](#country).`id`, on delete cascade
- [SubRegion](#subregion).`region_id` → `id`, on delete cascade
- [City](#city).`region_id` → `id`, on delete cascade

## SubRegion

Table `cities_light_subregion`, app `cities_light`.

Columns from `AbstractSubRegion` (abstract): `display_name`, `geoname_code`, `country_id`, `region_id`.

Columns from `Base` (abstract): `name`, `name_ascii`, `slug`, `geoname_id`, `alternate_names`, `translations`.

![SubRegion and the models it relates to](tables/SubRegion.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(200)` |  |  |  |  |
| `name_ascii` | `CharField` | `varchar(200)` |  |  |  |  |
| `slug` | `AutoSlugField` | `varchar(50)` |  |  |  |  |
| `geoname_id` | `IntegerField` | `integer` | yes | unique |  |  |
| `alternate_names` | `TextField` | `text` | yes |  | `''` |  |
| `translations` | `JSONField` | `text` |  |  | `dict()` |  |
| `display_name` | `CharField` | `varchar(200)` |  |  |  |  |
| `geoname_code` | `CharField` | `varchar(50)` | yes |  |  |  |
| `country_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `region_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `cities_light_subregion_name_2882337e` on `name`
- index `cities_light_subregion_name_ascii_ecf9a336` on `name_ascii`
- index `cities_light_subregion_slug_43494546` on `slug`
- index `cities_light_subregion_geoname_code_843acdc3` on `geoname_code`
- index `cities_light_subregion_country_id_9b32b484` on `country_id`
- index `cities_light_subregion_region_id_c6e0b71f` on `region_id`

Relations:

- `country_id` → [Country](#country).`id`, on delete cascade
- `region_id` → [Region](#region).`id`, on delete cascade
- [City](#city).`subregion_id` → `id`, on delete cascade

## City

Table `cities_light_city`, app `cities_light`.

Columns from `AbstractCity` (abstract): `display_name`, `search_names`, `latitude`, `longitude`, `subregion_id`, `region_id`, `country_id`, `population`, `feature_code`, `timezone`.

Columns from `Base` (abstract): `name`, `name_ascii`, `slug`, `geoname_id`, `alternate_names`, `translations`.

![City and the models it relates to](tables/City.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(200)` |  |  |  |  |
| `name_ascii` | `CharField` | `varchar(200)` |  |  |  |  |
| `slug` | `AutoSlugField` | `varchar(50)` |  |  |  |  |
| `geoname_id` | `IntegerField` | `integer` | yes | unique |  |  |
| `alternate_names` | `TextField` | `text` | yes |  | `''` |  |
| `translations` | `JSONField` | `text` |  |  | `dict()` |  |
| `display_name` | `CharField` | `varchar(200)` |  |  |  |  |
| `search_names` | `ToSearchTextField` | `text` |  |  | `''` |  |
| `latitude` | `DecimalField(8, 5)` | `decimal` | yes |  |  |  |
| `longitude` | `DecimalField(8, 5)` | `decimal` | yes |  |  |  |
| `subregion_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `region_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `country_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `population` | `BigIntegerField` | `bigint` | yes |  |  |  |
| `feature_code` | `CharField` | `varchar(10)` | yes |  |  |  |
| `timezone` | `CharField` | `varchar(40)` | yes |  |  |  |

Indexes:

- index `cities_light_city_name_4859e2a5` on `name`
- index `cities_light_city_name_ascii_1e56323b` on `name_ascii`
- index `cities_light_city_slug_cb2251f8` on `slug`
- index `cities_light_city_search_names_fb77fed2` on `search_names`
- index `cities_light_city_subregion_id_0926d2ad` on `subregion_id`
- index `cities_light_city_region_id_f7ab977b` on `region_id`
- index `cities_light_city_country_id_cf310fd2` on `country_id`
- index `cities_light_city_population_d597eeb6` on `population`
- index `cities_light_city_feature_code_d43c9217` on `feature_code`
- index `cities_light_city_timezone_0fb51b1e` on `timezone`
- unique `cities_light_city_region_id_subregion_id_name_cdfc77ea_uniq` on `region_id`, `subregion_id`, `name`
- unique `cities_light_city_region_id_subregion_id_slug_efb2e768_uniq` on `region_id`, `subregion_id`, `slug`

Relations:

- `subregion_id` → [SubRegion](#subregion).`id`, on delete cascade
- `region_id` → [Region](#region).`id`, on delete cascade
- `country_id` → [Country](#country).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## CountryContinent

| Name | Stored as |
| --- | --- |
| `Oceania` | `OC` |
| `Europe` | `EU` |
| `Africa` | `AF` |
| `North America` | `NA` |
| `Antarctica` | `AN` |
| `South America` | `SA` |
| `Asia` | `AS` |
