# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [StaticDevice](#staticdevice) · [StaticToken](#statictoken) · [YubikeyDevice](#yubikeydevice) · [ValidationService](#validationservice) · [RemoteYubikeyDevice](#remoteyubikeydevice)

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

CREATE TABLE "otp_static_staticdevice" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(64) NOT NULL, "confirmed" bool NOT NULL, "throttling_failure_timestamp" datetime NULL, "throttling_failure_count" integer unsigned NOT NULL CHECK ("throttling_failure_count" >= 0), "created_at" datetime NULL, "last_used_at" datetime NULL);

CREATE TABLE "otp_static_statictoken" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "device_id" integer NOT NULL REFERENCES "otp_static_staticdevice" ("id") DEFERRABLE INITIALLY DEFERRED, "token" varchar(16) NOT NULL);

CREATE TABLE "otp_yubikey_yubikeydevice" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(64) NOT NULL, "confirmed" bool NOT NULL, "private_id" varchar(12) NOT NULL, "key" varchar(32) NOT NULL, "session" integer unsigned NOT NULL CHECK ("session" >= 0), "counter" integer unsigned NOT NULL CHECK ("counter" >= 0));

CREATE TABLE "otp_yubikey_validationservice" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(32) NOT NULL, "api_id" integer NOT NULL, "api_key" varchar(64) NOT NULL, "base_url" varchar(200) NOT NULL, "api_version" varchar(8) NOT NULL, "use_ssl" bool NOT NULL, "param_sl" varchar(16) NOT NULL, "param_timeout" varchar(16) NOT NULL);

CREATE TABLE "otp_yubikey_remoteyubikeydevice" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(64) NOT NULL, "confirmed" bool NOT NULL, "service_id" integer NOT NULL REFERENCES "otp_yubikey_validationservice" ("id") DEFERRABLE INITIALLY DEFERRED, "public_id" varchar(32) NOT NULL);

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

CREATE INDEX "otp_static_staticdevice_user_id_7f9cff2b" ON "otp_static_staticdevice" ("user_id");

CREATE INDEX "otp_static_statictoken_device_id_74b7c7d1" ON "otp_static_statictoken" ("device_id");

CREATE INDEX "otp_static_statictoken_token_d0a51866" ON "otp_static_statictoken" ("token");

CREATE INDEX "otp_yubikey_yubikeydevice_user_id_80a424a7" ON "otp_yubikey_yubikeydevice" ("user_id");

CREATE INDEX "otp_yubikey_remoteyubikeydevice_user_id_80de3c1b" ON "otp_yubikey_remoteyubikeydevice" ("user_id");

CREATE INDEX "otp_yubikey_remoteyubikeydevice_service_id_63162285" ON "otp_yubikey_remoteyubikeydevice" ("service_id");
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
- [StaticDevice](#staticdevice).`user_id` → `id`, on delete cascade
- [YubikeyDevice](#yubikeydevice).`user_id` → `id`, on delete cascade
- [RemoteYubikeyDevice](#remoteyubikeydevice).`user_id` → `id`, on delete cascade

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

## StaticDevice

Table `otp_static_staticdevice`, app `otp_static`.

Columns from `TimestampMixin` (abstract): `created_at`, `last_used_at`.

Columns from `ThrottlingMixin` (abstract): `throttling_failure_timestamp`, `throttling_failure_count`.

Columns from `Device` (abstract): `user_id`, `name`, `confirmed`.

A static :class:`~django_otp.models.Device` simply consists of random
tokens shared by the database and the user.

These are frequently used as emergency tokens in case a user's normal
device is lost or unavailable. They can be consumed in any order; each
token will be removed from the database as soon as it is used.

This model has no fields of its own, but serves as a container for
:class:`StaticToken` objects.

.. attribute:: token_set

    The RelatedManager for our tokens.

![StaticDevice and the models it relates to](tables/StaticDevice.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  | The user that this device belongs to. |
| `name` | `CharField` | `varchar(64)` |  |  |  | The human-readable name of this device. |
| `confirmed` | `BooleanField` | `bool` |  |  | `True` | Is this device ready for use? |
| `throttling_failure_timestamp` | `DateTimeField` | `datetime` | yes |  | `None` | A timestamp of the last failed verification attempt. Null if last attempt succeeded. |
| `throttling_failure_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | Number of successive failed attempts. |
| `created_at` | `DateTimeField` | `datetime` | yes |  | `now()` | The date and time when this device was initially created in the system. |
| `last_used_at` | `DateTimeField` | `datetime` | yes |  |  | The most recent date and time this device was used. |

Indexes:

- index `otp_static_staticdevice_user_id_7f9cff2b` on `user_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [StaticToken](#statictoken).`device_id` → `id`, on delete cascade

## StaticToken

Table `otp_static_statictoken`, app `otp_static`.

A single token belonging to a :class:`StaticDevice`.

.. attribute:: device

    *ForeignKey*: A foreign key to :class:`StaticDevice`.

.. attribute:: token

    *CharField*: A random string up to 16 characters.

![StaticToken and the models it relates to](tables/StaticToken.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `device_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `token` | `CharField` | `varchar(16)` |  |  |  |  |

Indexes:

- index `otp_static_statictoken_device_id_74b7c7d1` on `device_id`
- index `otp_static_statictoken_token_d0a51866` on `token`

Relations:

- `device_id` → [StaticDevice](#staticdevice).`id`, on delete cascade

## YubikeyDevice

Table `otp_yubikey_yubikeydevice`, app `otp_yubikey`.

Columns from `Device` (abstract): `user_id`, `name`, `confirmed`.

Represents a locally-verified YubiKey OTP
:class:`~django_otp.models.Device`.

.. attribute:: private_id

    *CharField*: The 6-byte private ID (hex-encoded).

.. attribute:: key

    *CharField*: The 16-byte AES key shared with this YubiKey
    (hex-encoded).

.. attribute:: session

    *PositiveIntegerField*: The non-volatile session counter most recently
    used by this device.

.. attribute:: counter

    *PositiveIntegerField*: The volatile session usage counter most
    recently used by this device.

![YubikeyDevice and the models it relates to](tables/YubikeyDevice.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  | The user that this device belongs to. |
| `name` | `CharField` | `varchar(64)` |  |  |  | The human-readable name of this device. |
| `confirmed` | `BooleanField` | `bool` |  |  | `True` | Is this device ready for use? |
| `private_id` | `CharField` | `varchar(12)` |  |  | `default_id()` | The 6-byte private ID (hex-encoded). |
| `key` | `CharField` | `varchar(32)` |  |  | `default_key()` | The 16-byte AES key shared with this YubiKey (hex-encoded). |
| `session` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | The non-volatile session counter most recently used by this device. |
| `counter` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | The volatile session usage counter most recently used by this device. |

Indexes:

- index `otp_yubikey_yubikeydevice_user_id_80a424a7` on `user_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

## ValidationService

Table `otp_yubikey_validationservice`, app `otp_yubikey`.

Represents a YubiKey validation web service. By default, this will point to
Yubico's official hosted service, which you can customize. You can also
create instances to point at any other service implementing the same
protocol.

.. attribute:: name

    *CharField*: The name of this validation service.

.. attribute:: api_id

    *IntegerField*: Your API ID. The server needs this to sign responsees.
    (Default: 1)

.. attribute:: api_key

    *CharField*: Your base64-encoded API key, used to sign requests. This
    is optional but strongly recommended. (Default: ``''``)

.. attribute:: base_url

    *URLField*: The base URL of the verification service. Defaults to
    Yubico's hosted API.

.. attribute:: api_version

    *CharField*: The version of the validation API to use: '1.0', '1.1', or
    '2.0'. (Default: '2.0')

.. attribute:: use_ssl

    *BooleanField*: If ``True``, we'll use the HTTPS versions of the
    default URLs. (Default: ``True``).

.. attribute:: param_sl

    *CharField*: The level of syncing required. See
    :class:`~yubiotp.client.YubiClient20`.

.. attribute:: param_timeout

    *CharField*: The time to allow for syncing. See
    :class:`~yubiotp.client.YubiClient20`.

![ValidationService and the models it relates to](tables/ValidationService.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(32)` |  |  |  | The name of this validation service. |
| `api_id` | `IntegerField` | `integer` |  |  | `1` | Your API ID. |
| `api_key` | `CharField` | `varchar(64)` |  |  | `''` | Your base64-encoded API key. |
| `base_url` | `URLField` | `varchar(200)` |  |  | `''` | The base URL of the verification service. Defaults to Yubico's hosted API. |
| `api_version` | `ValidationServiceApiVersion` | `varchar(8)` |  |  | `'2.0'` | The version of the validation api to use. |
| `use_ssl` | `BooleanField` | `bool` |  |  | `True` | Use HTTPS API URLs by default? |
| `param_sl` | `CharField` | `varchar(16)` |  |  | `None` | The level of syncing required. |
| `param_timeout` | `CharField` | `varchar(16)` |  |  | `None` | The time to allow for syncing. |

Relations:

- [RemoteYubikeyDevice](#remoteyubikeydevice).`service_id` → `id`, on delete cascade

## RemoteYubikeyDevice

Table `otp_yubikey_remoteyubikeydevice`, app `otp_yubikey`.

Columns from `Device` (abstract): `user_id`, `name`, `confirmed`.

Represents a YubiKey device that is to be verified with a remote validation
service. In order create these devices, you must have at least one
:class:`~otp_yubikey.models.ValidationService` in the database.

.. attribute:: service

    *ForeignKey*: The validation service to use for this device.

.. attribute:: public_id

    *CharField*: The public identity of the YubiKey (modhex-encoded).

![RemoteYubikeyDevice and the models it relates to](tables/RemoteYubikeyDevice.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  | The user that this device belongs to. |
| `name` | `CharField` | `varchar(64)` |  |  |  | The human-readable name of this device. |
| `confirmed` | `BooleanField` | `bool` |  |  | `True` | Is this device ready for use? |
| `service_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `public_id` | `CharField` | `varchar(32)` |  |  |  | The public identity of the YubiKey (modhex-encoded). |

Indexes:

- index `otp_yubikey_remoteyubikeydevice_user_id_80de3c1b` on `user_id`
- index `otp_yubikey_remoteyubikeydevice_service_id_63162285` on `service_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `service_id` → [ValidationService](#validationservice).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## ValidationServiceApiVersion

| Name | Stored as |
| --- | --- |
| `1.0` | `1.0` |
| `1.1` | `1.1` |
| `2.0` | `2.0` |
