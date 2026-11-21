# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [StripeUser](#stripeuser) · [Feature](#feature) · [Product](#product) · [ProductFeature](#productfeature) · [Price](#price) · [Subscription](#subscription) · [SubscriptionItem](#subscriptionitem)

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

CREATE TABLE "drf_stripe_stripeuser" ("user_id" integer NOT NULL PRIMARY KEY REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "customer_id" varchar(128) NULL);

CREATE TABLE "drf_stripe_feature" ("feature_id" varchar(64) NOT NULL PRIMARY KEY, "description" varchar(256) NULL);

CREATE TABLE "drf_stripe_product" ("product_id" varchar(256) NOT NULL PRIMARY KEY, "active" bool NOT NULL, "description" varchar(1024) NULL, "name" varchar(256) NULL);

CREATE TABLE "drf_stripe_productfeature" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "product_id" varchar(256) NOT NULL REFERENCES "drf_stripe_product" ("product_id") DEFERRABLE INITIALLY DEFERRED, "feature_id" varchar(64) NOT NULL REFERENCES "drf_stripe_feature" ("feature_id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "drf_stripe_price" ("price_id" varchar(256) NOT NULL PRIMARY KEY, "product_id" varchar(256) NOT NULL REFERENCES "drf_stripe_product" ("product_id") DEFERRABLE INITIALLY DEFERRED, "nickname" varchar(256) NULL, "price" integer unsigned NOT NULL CHECK ("price" >= 0), "freq" varchar(64) NULL, "active" bool NOT NULL, "currency" varchar(3) NOT NULL);

CREATE TABLE "drf_stripe_subscription" ("subscription_id" varchar(256) NOT NULL PRIMARY KEY, "stripe_user_id" integer NOT NULL REFERENCES "drf_stripe_stripeuser" ("user_id") DEFERRABLE INITIALLY DEFERRED, "period_start" datetime NULL, "period_end" datetime NULL, "cancel_at" datetime NULL, "cancel_at_period_end" bool NOT NULL, "ended_at" datetime NULL, "status" varchar(64) NOT NULL, "trial_end" datetime NULL, "trial_start" datetime NULL);

CREATE TABLE "drf_stripe_subscriptionitem" ("sub_item_id" varchar(256) NOT NULL PRIMARY KEY, "subscription_id" varchar(256) NOT NULL REFERENCES "drf_stripe_subscription" ("subscription_id") DEFERRABLE INITIALLY DEFERRED, "price_id" varchar(256) NOT NULL REFERENCES "drf_stripe_price" ("price_id") DEFERRABLE INITIALLY DEFERRED, "quantity" integer unsigned NOT NULL CHECK ("quantity" >= 0));

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

CREATE INDEX "drf_stripe__user_id_6bbc0d_idx" ON "drf_stripe_stripeuser" ("user_id", "customer_id");

CREATE INDEX "drf_stripe_productfeature_product_id_ee68f65a" ON "drf_stripe_productfeature" ("product_id");

CREATE INDEX "drf_stripe_productfeature_feature_id_59595fe3" ON "drf_stripe_productfeature" ("feature_id");

CREATE INDEX "drf_stripe_price_product_id_bc402469" ON "drf_stripe_price" ("product_id");

CREATE INDEX "drf_stripe__active_3854b4_idx" ON "drf_stripe_price" ("active", "freq");

CREATE INDEX "drf_stripe_subscription_stripe_user_id_c0df8936" ON "drf_stripe_subscription" ("stripe_user_id");

CREATE INDEX "drf_stripe__stripe__7d71e7_idx" ON "drf_stripe_subscription" ("stripe_user_id", "status");

CREATE INDEX "drf_stripe_subscriptionitem_subscription_id_cdb63cde" ON "drf_stripe_subscriptionitem" ("subscription_id");

CREATE INDEX "drf_stripe_subscriptionitem_price_id_56df13e9" ON "drf_stripe_subscriptionitem" ("price_id");
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
- [StripeUser](#stripeuser).`user_id` → `id`, on delete cascade

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

## StripeUser

Table `drf_stripe_stripeuser`, app `drf_stripe`.

A model linking Django user model with a Stripe User

![StripeUser and the models it relates to](tables/StripeUser.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `user_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `customer_id` | `CharField` | `varchar(128)` | yes |  |  |  |

Indexes:

- index `drf_stripe__user_id_6bbc0d_idx` on `user_id`, `customer_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [Subscription](#subscription).`stripe_user_id` → `user_id`, on delete cascade

## Feature

Table `drf_stripe_feature`, app `drf_stripe`.

A model used to keep track of features provided by your application.
This does not correspond to a Stripe object, but the feature ids should be listed as
 a space delimited strings in Stripe.product.metadata.features

![Feature and the models it relates to](tables/Feature.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `feature_id` | `CharField` | `varchar(64)` |  | PK |  |  |
| `description` | `CharField` | `varchar(256)` | yes |  |  |  |

Relations:

- [ProductFeature](#productfeature).`feature_id` → `feature_id`, on delete cascade

## Product

Table `drf_stripe_product`, app `drf_stripe`.

A model representing a Stripe Product

![Product and the models it relates to](tables/Product.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `product_id` | `CharField` | `varchar(256)` |  | PK |  |  |
| `active` | `BooleanField` | `bool` |  |  |  |  |
| `description` | `CharField` | `varchar(1024)` | yes |  |  |  |
| `name` | `CharField` | `varchar(256)` | yes |  |  |  |

Relations:

- [ProductFeature](#productfeature).`product_id` → `product_id`, on delete cascade
- [Price](#price).`product_id` → `product_id`, on delete cascade

## ProductFeature

Table `drf_stripe_productfeature`, app `drf_stripe`.

A model representing association of Product and Feature instances. They have many-to-many relationship.

![ProductFeature and the models it relates to](tables/ProductFeature.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `product_id` | `ForeignKey` | `varchar(256)` |  | FK |  |  |
| `feature_id` | `ForeignKey` | `varchar(64)` |  | FK |  |  |

Indexes:

- index `drf_stripe_productfeature_product_id_ee68f65a` on `product_id`
- index `drf_stripe_productfeature_feature_id_59595fe3` on `feature_id`

Relations:

- `product_id` → [Product](#product).`product_id`, on delete cascade
- `feature_id` → [Feature](#feature).`feature_id`, on delete cascade

## Price

Table `drf_stripe_price`, app `drf_stripe`.

A model representing to a Stripe Price object, with enhanced attributes.

![Price and the models it relates to](tables/Price.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `price_id` | `CharField` | `varchar(256)` |  | PK |  |  |
| `product_id` | `ForeignKey` | `varchar(256)` |  | FK |  |  |
| `nickname` | `CharField` | `varchar(256)` | yes |  |  |  |
| `price` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `freq` | `CharField` | `varchar(64)` | yes |  |  |  |
| `active` | `BooleanField` | `bool` |  |  |  |  |
| `currency` | `CharField` | `varchar(3)` |  |  |  |  |

Indexes:

- index `drf_stripe_price_product_id_bc402469` on `product_id`
- index `drf_stripe__active_3854b4_idx` on `active`, `freq`

Relations:

- `product_id` → [Product](#product).`product_id`, on delete cascade
- [SubscriptionItem](#subscriptionitem).`price_id` → `price_id`, on delete cascade

## Subscription

Table `drf_stripe_subscription`, app `drf_stripe`.

A model representing Subscription, corresponding to a Stripe Subscription object.

![Subscription and the models it relates to](tables/Subscription.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `subscription_id` | `CharField` | `varchar(256)` |  | PK |  |  |
| `stripe_user_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `period_start` | `DateTimeField` | `datetime` | yes |  |  |  |
| `period_end` | `DateTimeField` | `datetime` | yes |  |  |  |
| `cancel_at` | `DateTimeField` | `datetime` | yes |  |  |  |
| `cancel_at_period_end` | `BooleanField` | `bool` |  |  |  |  |
| `ended_at` | `DateTimeField` | `datetime` | yes |  |  |  |
| `status` | `CharField` | `varchar(64)` |  |  |  |  |
| `trial_end` | `DateTimeField` | `datetime` | yes |  |  |  |
| `trial_start` | `DateTimeField` | `datetime` | yes |  |  |  |

Indexes:

- index `drf_stripe_subscription_stripe_user_id_c0df8936` on `stripe_user_id`
- index `drf_stripe__stripe__7d71e7_idx` on `stripe_user_id`, `status`

Relations:

- `stripe_user_id` → [StripeUser](#stripeuser).`user_id`, on delete cascade
- [SubscriptionItem](#subscriptionitem).`subscription_id` → `subscription_id`, on delete cascade

## SubscriptionItem

Table `drf_stripe_subscriptionitem`, app `drf_stripe`.

A model representing relation of Price and Subscription, corresponding to Stripe Subscription line item.

![SubscriptionItem and the models it relates to](tables/SubscriptionItem.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `sub_item_id` | `CharField` | `varchar(256)` |  | PK |  |  |
| `subscription_id` | `ForeignKey` | `varchar(256)` |  | FK |  |  |
| `price_id` | `ForeignKey` | `varchar(256)` |  | FK |  |  |
| `quantity` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |

Indexes:

- index `drf_stripe_subscriptionitem_subscription_id_cdb63cde` on `subscription_id`
- index `drf_stripe_subscriptionitem_price_id_56df13e9` on `price_id`

Relations:

- `subscription_id` → [Subscription](#subscription).`subscription_id`, on delete cascade
- `price_id` → [Price](#price).`price_id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |
