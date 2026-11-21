# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Key](#key) · [LtiRegistration](#ltiregistration) · [LtiPlatformInstance](#ltiplatforminstance) · [LtiDeployment](#ltideployment) · [LtiUser](#ltiuser) · [LtiContext](#lticontext) · [LtiMembership](#ltimembership) · [LtiResourceLink](#ltiresourcelink) · [LtiLineItem](#ltilineitem)

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

CREATE TABLE "lti_tool_key" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "public_key" text NOT NULL, "private_key" text NOT NULL, "is_active" bool NOT NULL, "datetime_created" datetime NOT NULL, "datetime_modified" datetime NOT NULL);

CREATE TABLE "lti_tool_ltiregistration" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(255) NOT NULL, "uuid" char(32) NOT NULL, "issuer" varchar(255) NOT NULL, "client_id" varchar(255) NOT NULL, "audience" varchar(255) NOT NULL, "auth_url" varchar(200) NOT NULL, "token_url" varchar(200) NOT NULL, "keyset_url" varchar(200) NOT NULL, "is_active" bool NOT NULL, "public_key" text NOT NULL, "private_key" text NOT NULL, "datetime_created" datetime NOT NULL, "datetime_modified" datetime NOT NULL, CONSTRAINT "unique_client_id_per_issuer" UNIQUE ("issuer", "client_id"));

CREATE TABLE "lti_tool_ltiplatforminstance" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "issuer" varchar(255) NOT NULL, "guid" varchar(255) NOT NULL, "lti1p1_id_on_platform" varchar(255) NOT NULL, "contact_email" varchar(254) NOT NULL, "description" text NOT NULL, "name" varchar(500) NOT NULL, "url" varchar(200) NOT NULL, "product_family_code" varchar(500) NOT NULL, "version" varchar(500) NOT NULL, CONSTRAINT "unique_guid_per_issuer" UNIQUE ("issuer", "guid"));

CREATE TABLE "lti_tool_ltideployment" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "registration_id" integer NOT NULL REFERENCES "lti_tool_ltiregistration" ("id") DEFERRABLE INITIALLY DEFERRED, "deployment_id" varchar(255) NOT NULL, "is_active" bool NOT NULL, "platform_instance_id" integer NULL REFERENCES "lti_tool_ltiplatforminstance" ("id") DEFERRABLE INITIALLY DEFERRED, "pns_url" varchar(200) NOT NULL, "supports_context_copy_notice" bool NOT NULL, "supports_asset_processor_submission_notice" bool NOT NULL, "datetime_created" datetime NOT NULL, "datetime_modified" datetime NOT NULL, CONSTRAINT "unique_deployment_id_per_registration" UNIQUE ("registration_id", "deployment_id"));

CREATE TABLE "lti_tool_ltiuser" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "registration_id" integer NOT NULL REFERENCES "lti_tool_ltiregistration" ("id") DEFERRABLE INITIALLY DEFERRED, "sub" varchar(255) NOT NULL, "lti1p1_id_on_platform" varchar(255) NOT NULL, "given_name" varchar(500) NOT NULL, "family_name" varchar(500) NOT NULL, "name" varchar(500) NOT NULL, "email" varchar(254) NOT NULL, "picture_url" varchar(200) NOT NULL, "auth_user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "datetime_created" datetime NOT NULL, "datetime_modified" datetime NOT NULL, CONSTRAINT "unique_sub_per_registration" UNIQUE ("registration_id", "sub"));

CREATE TABLE "lti_tool_lticontext" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "deployment_id" integer NOT NULL REFERENCES "lti_tool_ltideployment" ("id") DEFERRABLE INITIALLY DEFERRED, "id_on_platform" varchar(255) NOT NULL, "lti1p1_id_on_platform" varchar(255) NOT NULL, "label" varchar(255) NOT NULL, "title" varchar(500) NOT NULL, "is_course_template" bool NOT NULL, "is_course_offering" bool NOT NULL, "is_course_section" bool NOT NULL, "is_group" bool NOT NULL, "memberships_url" varchar(200) NOT NULL, "lineitems_url" varchar(200) NOT NULL, "can_query_lineitems" bool NOT NULL, "can_manage_lineitems" bool NOT NULL, "can_publish_scores" bool NOT NULL, "can_access_results" bool NOT NULL, "datetime_created" datetime NOT NULL, "datetime_modified" datetime NOT NULL, CONSTRAINT "unique_context_id_per_deployment" UNIQUE ("deployment_id", "id_on_platform"));

CREATE TABLE "lti_tool_ltimembership" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "lti_tool_ltiuser" ("id") DEFERRABLE INITIALLY DEFERRED, "context_id" integer NOT NULL REFERENCES "lti_tool_lticontext" ("id") DEFERRABLE INITIALLY DEFERRED, "is_administrator" bool NOT NULL, "is_content_developer" bool NOT NULL, "is_instructor" bool NOT NULL, "is_learner" bool NOT NULL, "is_mentor" bool NOT NULL, "is_active" bool NOT NULL, "datetime_created" datetime NOT NULL, "datetime_modified" datetime NOT NULL, CONSTRAINT "no_duplicate_memberships" UNIQUE ("user_id", "context_id"));

CREATE TABLE "lti_tool_ltiresourcelink" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "context_id" integer NOT NULL REFERENCES "lti_tool_lticontext" ("id") DEFERRABLE INITIALLY DEFERRED, "id_on_platform" varchar(255) NOT NULL, "lti1p1_id_on_platform" varchar(255) NOT NULL, "title" varchar(500) NOT NULL, "description" text NOT NULL, "datetime_created" datetime NOT NULL, "datetime_modified" datetime NOT NULL, CONSTRAINT "unique_resource_link_id_per_context" UNIQUE ("context_id", "id_on_platform"));

CREATE TABLE "lti_tool_ltilineitem" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "context_id" integer NOT NULL REFERENCES "lti_tool_lticontext" ("id") DEFERRABLE INITIALLY DEFERRED, "url" varchar(200) NOT NULL UNIQUE, "maximum_score" real NOT NULL, "label" varchar(255) NOT NULL, "tag" varchar(255) NOT NULL, "resource_id" varchar(255) NOT NULL, "resource_link_id" integer NULL REFERENCES "lti_tool_ltiresourcelink" ("id") DEFERRABLE INITIALLY DEFERRED, "start_datetime" datetime NULL, "end_datetime" datetime NULL);

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

CREATE INDEX "lti_tool_ltideployment_registration_id_cb081313" ON "lti_tool_ltideployment" ("registration_id");

CREATE INDEX "lti_tool_ltideployment_platform_instance_id_3af9e722" ON "lti_tool_ltideployment" ("platform_instance_id");

CREATE INDEX "lti_tool_ltiuser_registration_id_9eb044b9" ON "lti_tool_ltiuser" ("registration_id");

CREATE INDEX "lti_tool_ltiuser_auth_user_id_c6b9c91c" ON "lti_tool_ltiuser" ("auth_user_id");

CREATE INDEX "lti_tool_lticontext_deployment_id_2a02049c" ON "lti_tool_lticontext" ("deployment_id");

CREATE INDEX "lti_tool_ltimembership_user_id_f52055e9" ON "lti_tool_ltimembership" ("user_id");

CREATE INDEX "lti_tool_ltimembership_context_id_8eeef5a6" ON "lti_tool_ltimembership" ("context_id");

CREATE INDEX "lti_tool_ltiresourcelink_context_id_a17e2ea2" ON "lti_tool_ltiresourcelink" ("context_id");

CREATE INDEX "lti_tool_ltilineitem_context_id_812de9e7" ON "lti_tool_ltilineitem" ("context_id");

CREATE INDEX "lti_tool_ltilineitem_resource_link_id_28f7637a" ON "lti_tool_ltilineitem" ("resource_link_id");
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
- [LtiUser](#ltiuser).`auth_user_id` → `id`, on delete set null

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

## Key

Table `lti_tool_key`, app `lti_tool`.

A keypair for use in LTI integrations.

Attributes:
    public_key (str): Public key data.
    private_key (str): Private key data.
    is_active (bool): Indicates if the key is present in the JWKS.
    datetime_created (datetime): When the key was created.
    datetime_modified (datetime): When the key was last modified.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `public_key` | `TextField` | `text` |  |  |  |  |
| `private_key` | `TextField` | `text` |  |  |  |  |
| `is_active` | `BooleanField` | `bool` |  |  | `True` |  |
| `datetime_created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `datetime_modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |

## LtiRegistration

Table `lti_tool_ltiregistration`, app `lti_tool`.

An LTI platform registration.

Attributes:
    name (str): A display name for this registration
    uuid (uuid4): A unique identifier for registraitons for use in OIDC login init.
    issuer (str): The platform's issuer.
    client_id (str): The client ID assigned to the tool by the platform.
    audience (str): The platforms's OAuth2 audience.
    auth_url (str): The platform's auth login URL.
    token_url (str): The platform's access token retrieval URL.
    keyset_url (str): The platform's JWKS URL.
    is_active (bool): Indicates if the registration is active within the tool.
    public_key (str): Public key data specific to this registration. Only to be used
        when the platform cannot retrieve keys from a tool's JWKS URL.
    private_key (str): Private key data specific to this registration. Only to be
        used when the platform cannot retrieve keys from a tool's JWKS URL.
    datetime_created (datetime): When the registration was created.
    datetime_modified (datetime): When the registration was last modified.

![LtiRegistration and the models it relates to](tables/LtiRegistration.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `uuid` | `UUIDField` | `char(32)` |  |  | `uuid4()` |  |
| `issuer` | `CharField` | `varchar(255)` |  |  |  |  |
| `client_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `audience` | `CharField` | `varchar(255)` |  |  |  |  |
| `auth_url` | `URLField` | `varchar(200)` |  |  |  |  |
| `token_url` | `URLField` | `varchar(200)` |  |  |  |  |
| `keyset_url` | `URLField` | `varchar(200)` |  |  |  |  |
| `is_active` | `BooleanField` | `bool` |  |  | `True` |  |
| `public_key` | `TextField` | `text` |  |  |  |  |
| `private_key` | `TextField` | `text` |  |  |  |  |
| `datetime_created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `datetime_modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- unique `unique_client_id_per_issuer` on `issuer`, `client_id`

Relations:

- [LtiDeployment](#ltideployment).`registration_id` → `id`, on delete cascade
- [LtiUser](#ltiuser).`registration_id` → `id`, on delete cascade

## LtiPlatformInstance

Table `lti_tool_ltiplatforminstance`, app `lti_tool`.

An instance of a LTI learning platform.

The platform instance claim is described in the LTI spec at
https://www.imsglobal.org/spec/lti/v1p3/#platform-instance-claim

![LtiPlatformInstance and the models it relates to](tables/LtiPlatformInstance.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `issuer` | `CharField` | `varchar(255)` |  |  |  |  |
| `guid` | `CharField` | `varchar(255)` |  |  |  |  |
| `lti1p1_id_on_platform` | `CharField` | `varchar(255)` |  |  |  |  |
| `contact_email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `name` | `CharField` | `varchar(500)` |  |  |  |  |
| `url` | `URLField` | `varchar(200)` |  |  |  |  |
| `product_family_code` | `CharField` | `varchar(500)` |  |  |  |  |
| `version` | `CharField` | `varchar(500)` |  |  |  |  |

Indexes:

- unique `unique_guid_per_issuer` on `issuer`, `guid`

Relations:

- [LtiDeployment](#ltideployment).`platform_instance_id` → `id`, on delete set null

## LtiDeployment

Table `lti_tool_ltideployment`, app `lti_tool`.

A LTI deployment.

Deployments are described in the LTI 1.3 core specification at
https://www.imsglobal.org/spec/lti/v1p3/#tool-deployment.

Attributes:
    registration (LtiRegistration): The platform registration associated with
        this deployment.
    deployment_id (str): The deployment ID as provided by the platform.
    is_active (bool): Indicates if the deployment has been activated in the tool.
    platform_instance (LtiPlatformInstance): The platform instance associated with
        this deployment.
    pns_url (str): The URL for the Platform Notification Service.
    supports_context_copy_notice (bool): Indicates if the deployment supports
        the Context Copy Notice.
    supports_asset_processor_submission_notice (bool): Indicates if the deployment
        supports the Asset Processor Submission Notice.
    datetime_created (datetime): When the deployment was created.
    datetime_modified (datetime): When the deployment was last modified.

![LtiDeployment and the models it relates to](tables/LtiDeployment.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `registration_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `deployment_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `is_active` | `BooleanField` | `bool` |  |  | `False` |  |
| `platform_instance_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `pns_url` | `URLField` | `varchar(200)` |  |  |  |  |
| `supports_context_copy_notice` | `BooleanField` | `bool` |  |  | `False` |  |
| `supports_asset_processor_submission_notice` | `BooleanField` | `bool` |  |  | `False` |  |
| `datetime_created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `datetime_modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `lti_tool_ltideployment_registration_id_cb081313` on `registration_id`
- index `lti_tool_ltideployment_platform_instance_id_3af9e722` on `platform_instance_id`
- unique `unique_deployment_id_per_registration` on `registration_id`, `deployment_id`

Relations:

- `registration_id` → [LtiRegistration](#ltiregistration).`id`, on delete cascade
- `platform_instance_id` → [LtiPlatformInstance](#ltiplatforminstance).`id`, on delete set null
- [LtiContext](#lticontext).`deployment_id` → `id`, on delete cascade

## LtiUser

Table `lti_tool_ltiuser`, app `lti_tool`.

A platform user as described by an LTI launch.

The user identity claims are described in the LTI spec at
https://www.imsglobal.org/spec/lti/v1p3/#user-identity-claims

![LtiUser and the models it relates to](tables/LtiUser.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `registration_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `sub` | `CharField` | `varchar(255)` |  |  |  |  |
| `lti1p1_id_on_platform` | `CharField` | `varchar(255)` |  |  |  |  |
| `given_name` | `CharField` | `varchar(500)` |  |  |  |  |
| `family_name` | `CharField` | `varchar(500)` |  |  |  |  |
| `name` | `CharField` | `varchar(500)` |  |  |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `picture_url` | `URLField` | `varchar(200)` |  |  |  |  |
| `auth_user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `datetime_created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `datetime_modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `lti_tool_ltiuser_registration_id_9eb044b9` on `registration_id`
- index `lti_tool_ltiuser_auth_user_id_c6b9c91c` on `auth_user_id`
- unique `unique_sub_per_registration` on `registration_id`, `sub`

Relations:

- `contexts` — many to many with [LtiContext](#lticontext), through [LtiMembership](#ltimembership)
- `registration_id` → [LtiRegistration](#ltiregistration).`id`, on delete cascade
- `auth_user_id` → [User](#user).`id`, on delete set null
- [LtiMembership](#ltimembership).`user_id` → `id`, on delete cascade

## LtiContext

Table `lti_tool_lticontext`, app `lti_tool`.

Describes the context of a LTI launch.

The context claim are described in the LTI spec at
https://www.imsglobal.org/spec/lti/v1p3/#context-claim

![LtiContext and the models it relates to](tables/LtiContext.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `deployment_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `id_on_platform` | `CharField` | `varchar(255)` |  |  |  |  |
| `lti1p1_id_on_platform` | `CharField` | `varchar(255)` |  |  |  |  |
| `label` | `CharField` | `varchar(255)` |  |  |  |  |
| `title` | `CharField` | `varchar(500)` |  |  |  |  |
| `is_course_template` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_course_offering` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_course_section` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_group` | `BooleanField` | `bool` |  |  | `False` |  |
| `memberships_url` | `URLField` | `varchar(200)` |  |  |  |  |
| `lineitems_url` | `URLField` | `varchar(200)` |  |  |  |  |
| `can_query_lineitems` | `BooleanField` | `bool` |  |  | `False` |  |
| `can_manage_lineitems` | `BooleanField` | `bool` |  |  | `False` |  |
| `can_publish_scores` | `BooleanField` | `bool` |  |  | `False` |  |
| `can_access_results` | `BooleanField` | `bool` |  |  | `False` |  |
| `datetime_created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `datetime_modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `lti_tool_lticontext_deployment_id_2a02049c` on `deployment_id`
- unique `unique_context_id_per_deployment` on `deployment_id`, `id_on_platform`

Relations:

- `members` — many to many with [LtiUser](#ltiuser), through [LtiMembership](#ltimembership)
- `deployment_id` → [LtiDeployment](#ltideployment).`id`, on delete cascade
- [LtiMembership](#ltimembership).`context_id` → `id`, on delete cascade
- [LtiResourceLink](#ltiresourcelink).`context_id` → `id`, on delete cascade
- [LtiLineItem](#ltilineitem).`context_id` → `id`, on delete cascade

## LtiMembership

Table `lti_tool_ltimembership`, app `lti_tool`.

A user's role within a learning context.

Roles are described in the LTI spec at
https://www.imsglobal.org/spec/lti/v1p3/#roles-claim

![LtiMembership and the models it relates to](tables/LtiMembership.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `context_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `is_administrator` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_content_developer` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_instructor` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_learner` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_mentor` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_active` | `BooleanField` | `bool` |  |  | `True` |  |
| `datetime_created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `datetime_modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `lti_tool_ltimembership_user_id_f52055e9` on `user_id`
- index `lti_tool_ltimembership_context_id_8eeef5a6` on `context_id`
- unique `no_duplicate_memberships` on `user_id`, `context_id`

Relations:

- `user_id` → [LtiUser](#ltiuser).`id`, on delete cascade
- `context_id` → [LtiContext](#lticontext).`id`, on delete cascade

## LtiResourceLink

Table `lti_tool_ltiresourcelink`, app `lti_tool`.

Represents a resource link in an LTI launch.

The resource link claim is described in the LTI spec at
https://www.imsglobal.org/spec/lti/v1p3/#resource-link-claim

![LtiResourceLink and the models it relates to](tables/LtiResourceLink.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `context_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `id_on_platform` | `CharField` | `varchar(255)` |  |  |  |  |
| `lti1p1_id_on_platform` | `CharField` | `varchar(255)` |  |  |  |  |
| `title` | `CharField` | `varchar(500)` |  |  |  |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `datetime_created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `datetime_modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `lti_tool_ltiresourcelink_context_id_a17e2ea2` on `context_id`
- unique `unique_resource_link_id_per_context` on `context_id`, `id_on_platform`

Relations:

- `context_id` → [LtiContext](#lticontext).`id`, on delete cascade
- [LtiLineItem](#ltilineitem).`resource_link_id` → `id`, on delete set null

## LtiLineItem

Table `lti_tool_ltilineitem`, app `lti_tool`.

Represents an LTI LineItem for assignment scoring.

The line item scope and methods are described in the LTI Assignment and Grade
Services 2.0 spec at
https://www.imsglobal.org/spec/lti-ags/v2p0#line-item-service-scope-and-allowed-http-methods

![LtiLineItem and the models it relates to](tables/LtiLineItem.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `context_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `url` | `URLField` | `varchar(200)` |  | unique |  |  |
| `maximum_score` | `FloatField` | `real` |  |  |  |  |
| `label` | `CharField` | `varchar(255)` |  |  |  |  |
| `tag` | `CharField` | `varchar(255)` |  |  |  |  |
| `resource_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `resource_link_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `start_datetime` | `DateTimeField` | `datetime` | yes |  |  |  |
| `end_datetime` | `DateTimeField` | `datetime` | yes |  |  |  |

Indexes:

- index `lti_tool_ltilineitem_context_id_812de9e7` on `context_id`
- index `lti_tool_ltilineitem_resource_link_id_28f7637a` on `resource_link_id`

Relations:

- `context_id` → [LtiContext](#lticontext).`id`, on delete cascade
- `resource_link_id` → [LtiResourceLink](#ltiresourcelink).`id`, on delete set null

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |
