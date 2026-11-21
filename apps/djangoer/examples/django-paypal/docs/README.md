# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [PayPalIPN](#paypalipn) · [PayPalPDT](#paypalpdt)

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

CREATE TABLE "paypal_ipn" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "business" varchar(127) NOT NULL, "charset" varchar(255) NOT NULL, "custom" varchar(256) NOT NULL, "notify_version" decimal NULL, "parent_txn_id" varchar(19) NOT NULL, "receiver_email" varchar(254) NOT NULL, "receiver_id" varchar(255) NOT NULL, "residence_country" varchar(2) NOT NULL, "test_ipn" bool NOT NULL, "txn_id" varchar(255) NOT NULL, "txn_type" varchar(255) NOT NULL, "verify_sign" varchar(255) NOT NULL, "address_country" varchar(64) NOT NULL, "address_city" varchar(40) NOT NULL, "address_country_code" varchar(64) NOT NULL, "address_name" varchar(128) NOT NULL, "address_state" varchar(40) NOT NULL, "address_status" varchar(255) NOT NULL, "address_street" varchar(200) NOT NULL, "address_zip" varchar(20) NOT NULL, "contact_phone" varchar(20) NOT NULL, "first_name" varchar(64) NOT NULL, "last_name" varchar(64) NOT NULL, "payer_business_name" varchar(127) NOT NULL, "payer_email" varchar(127) NOT NULL, "payer_id" varchar(13) NOT NULL, "auth_amount" decimal NULL, "auth_exp" varchar(28) NOT NULL, "auth_id" varchar(19) NOT NULL, "auth_status" varchar(255) NOT NULL, "exchange_rate" decimal NULL, "invoice" varchar(127) NOT NULL, "item_name" varchar(127) NOT NULL, "item_number" varchar(127) NOT NULL, "mc_currency" varchar(32) NOT NULL, "mc_fee" decimal NULL, "mc_gross" decimal NULL, "mc_handling" decimal NULL, "mc_shipping" decimal NULL, "memo" varchar(255) NOT NULL, "num_cart_items" integer NULL, "option_name1" varchar(64) NOT NULL, "option_name2" varchar(64) NOT NULL, "option_selection1" varchar(200) NOT NULL, "option_selection2" varchar(200) NOT NULL, "payer_status" varchar(255) NOT NULL, "payment_date" datetime NULL, "payment_gross" decimal NULL, "payment_status" varchar(255) NOT NULL, "payment_type" varchar(255) NOT NULL, "pending_reason" varchar(255) NOT NULL, "protection_eligibility" varchar(255) NOT NULL, "quantity" integer NULL, "reason_code" varchar(255) NOT NULL, "remaining_settle" decimal NULL, "settle_amount" decimal NULL, "settle_currency" varchar(32) NOT NULL, "shipping" decimal NULL, "shipping_method" varchar(255) NOT NULL, "tax" decimal NULL, "transaction_entity" varchar(255) NOT NULL, "auction_buyer_id" varchar(64) NOT NULL, "auction_closing_date" datetime NULL, "auction_multi_item" integer NULL, "for_auction" decimal NULL, "amount" decimal NULL, "amount_per_cycle" decimal NULL, "initial_payment_amount" decimal NULL, "next_payment_date" datetime NULL, "outstanding_balance" decimal NULL, "payment_cycle" varchar(255) NOT NULL, "period_type" varchar(255) NOT NULL, "product_name" varchar(255) NOT NULL, "product_type" varchar(255) NOT NULL, "profile_status" varchar(255) NOT NULL, "recurring_payment_id" varchar(255) NOT NULL, "rp_invoice_id" varchar(127) NOT NULL, "time_created" datetime NULL, "amount1" decimal NULL, "amount2" decimal NULL, "amount3" decimal NULL, "mc_amount1" decimal NULL, "mc_amount2" decimal NULL, "mc_amount3" decimal NULL, "password" varchar(24) NOT NULL, "period1" varchar(255) NOT NULL, "period2" varchar(255) NOT NULL, "period3" varchar(255) NOT NULL, "reattempt" varchar(1) NOT NULL, "recur_times" integer NULL, "recurring" varchar(1) NOT NULL, "retry_at" datetime NULL, "subscr_date" datetime NULL, "subscr_effective" datetime NULL, "subscr_id" varchar(19) NOT NULL, "username" varchar(64) NOT NULL, "mp_id" varchar(128) NULL, "case_creation_date" datetime NULL, "case_id" varchar(255) NOT NULL, "case_type" varchar(255) NOT NULL, "receipt_id" varchar(255) NOT NULL, "currency_code" varchar(32) NOT NULL, "handling_amount" decimal NULL, "transaction_subject" varchar(256) NOT NULL, "ipaddress" char(39) NULL, "flag" bool NOT NULL, "flag_code" varchar(16) NOT NULL, "flag_info" text NOT NULL, "query" text NOT NULL, "response" text NOT NULL, "created_at" datetime NOT NULL, "updated_at" datetime NOT NULL, "from_view" varchar(6) NULL);

CREATE TABLE "paypal_pdt" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "business" varchar(127) NOT NULL, "charset" varchar(255) NOT NULL, "custom" varchar(256) NOT NULL, "notify_version" decimal NULL, "parent_txn_id" varchar(19) NOT NULL, "receiver_email" varchar(254) NOT NULL, "receiver_id" varchar(255) NOT NULL, "residence_country" varchar(2) NOT NULL, "test_ipn" bool NOT NULL, "txn_id" varchar(255) NOT NULL, "txn_type" varchar(255) NOT NULL, "verify_sign" varchar(255) NOT NULL, "address_country" varchar(64) NOT NULL, "address_city" varchar(40) NOT NULL, "address_country_code" varchar(64) NOT NULL, "address_name" varchar(128) NOT NULL, "address_state" varchar(40) NOT NULL, "address_status" varchar(255) NOT NULL, "address_street" varchar(200) NOT NULL, "address_zip" varchar(20) NOT NULL, "contact_phone" varchar(20) NOT NULL, "first_name" varchar(64) NOT NULL, "last_name" varchar(64) NOT NULL, "payer_business_name" varchar(127) NOT NULL, "payer_email" varchar(127) NOT NULL, "payer_id" varchar(13) NOT NULL, "auth_amount" decimal NULL, "auth_exp" varchar(28) NOT NULL, "auth_id" varchar(19) NOT NULL, "auth_status" varchar(255) NOT NULL, "exchange_rate" decimal NULL, "invoice" varchar(127) NOT NULL, "item_name" varchar(127) NOT NULL, "item_number" varchar(127) NOT NULL, "mc_currency" varchar(32) NOT NULL, "mc_fee" decimal NULL, "mc_gross" decimal NULL, "mc_handling" decimal NULL, "mc_shipping" decimal NULL, "memo" varchar(255) NOT NULL, "num_cart_items" integer NULL, "option_name1" varchar(64) NOT NULL, "option_name2" varchar(64) NOT NULL, "option_selection1" varchar(200) NOT NULL, "option_selection2" varchar(200) NOT NULL, "payer_status" varchar(255) NOT NULL, "payment_date" datetime NULL, "payment_gross" decimal NULL, "payment_status" varchar(255) NOT NULL, "payment_type" varchar(255) NOT NULL, "pending_reason" varchar(255) NOT NULL, "protection_eligibility" varchar(255) NOT NULL, "quantity" integer NULL, "reason_code" varchar(255) NOT NULL, "remaining_settle" decimal NULL, "settle_amount" decimal NULL, "settle_currency" varchar(32) NOT NULL, "shipping" decimal NULL, "shipping_method" varchar(255) NOT NULL, "tax" decimal NULL, "transaction_entity" varchar(255) NOT NULL, "auction_buyer_id" varchar(64) NOT NULL, "auction_closing_date" datetime NULL, "auction_multi_item" integer NULL, "for_auction" decimal NULL, "amount" decimal NULL, "amount_per_cycle" decimal NULL, "initial_payment_amount" decimal NULL, "next_payment_date" datetime NULL, "outstanding_balance" decimal NULL, "payment_cycle" varchar(255) NOT NULL, "period_type" varchar(255) NOT NULL, "product_name" varchar(255) NOT NULL, "product_type" varchar(255) NOT NULL, "profile_status" varchar(255) NOT NULL, "recurring_payment_id" varchar(255) NOT NULL, "rp_invoice_id" varchar(127) NOT NULL, "time_created" datetime NULL, "amount1" decimal NULL, "amount2" decimal NULL, "amount3" decimal NULL, "mc_amount1" decimal NULL, "mc_amount2" decimal NULL, "mc_amount3" decimal NULL, "password" varchar(24) NOT NULL, "period1" varchar(255) NOT NULL, "period2" varchar(255) NOT NULL, "period3" varchar(255) NOT NULL, "reattempt" varchar(1) NOT NULL, "recur_times" integer NULL, "recurring" varchar(1) NOT NULL, "retry_at" datetime NULL, "subscr_date" datetime NULL, "subscr_effective" datetime NULL, "subscr_id" varchar(19) NOT NULL, "username" varchar(64) NOT NULL, "mp_id" varchar(128) NULL, "case_creation_date" datetime NULL, "case_id" varchar(255) NOT NULL, "case_type" varchar(255) NOT NULL, "receipt_id" varchar(255) NOT NULL, "currency_code" varchar(32) NOT NULL, "handling_amount" decimal NULL, "transaction_subject" varchar(256) NOT NULL, "ipaddress" char(39) NULL, "flag" bool NOT NULL, "flag_code" varchar(16) NOT NULL, "flag_info" text NOT NULL, "query" text NOT NULL, "response" text NOT NULL, "created_at" datetime NOT NULL, "updated_at" datetime NOT NULL, "from_view" varchar(6) NULL, "amt" decimal NULL, "cm" varchar(255) NOT NULL, "sig" varchar(255) NOT NULL, "tx" varchar(255) NOT NULL, "st" varchar(32) NOT NULL);

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

CREATE INDEX "paypal_ipn_txn_id_8fa22c44" ON "paypal_ipn" ("txn_id");

CREATE INDEX "paypal_pdt_txn_id_0b75b30b" ON "paypal_pdt" ("txn_id");
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

## PayPalIPN

Table `paypal_ipn`, app `ipn`.

Columns from `PayPalStandardBase` (abstract): `business`, `charset`, `custom`, `notify_version`, `parent_txn_id`, `receiver_email`, `receiver_id`, `residence_country`, `test_ipn`, `txn_id`, `txn_type`, `verify_sign`, `address_country`, `address_city`, `address_country_code`, `address_name`, `address_state`, `address_status`, `address_street`, `address_zip`, `contact_phone`, `first_name`, `last_name`, `payer_business_name`, `payer_email`, `payer_id`, `auth_amount`, `auth_exp`, `auth_id`, `auth_status`, `exchange_rate`, `invoice`, `item_name`, `item_number`, `mc_currency`, `mc_fee`, `mc_gross`, `mc_handling`, `mc_shipping`, `memo`, `num_cart_items`, `option_name1`, `option_name2`, `option_selection1`, `option_selection2`, `payer_status`, `payment_date`, `payment_gross`, `payment_status`, `payment_type`, `pending_reason`, `protection_eligibility`, `quantity`, `reason_code`, `remaining_settle`, `settle_amount`, `settle_currency`, `shipping`, `shipping_method`, `tax`, `transaction_entity`, `auction_buyer_id`, `auction_closing_date`, `auction_multi_item`, `for_auction`, `amount`, `amount_per_cycle`, `initial_payment_amount`, `next_payment_date`, `outstanding_balance`, `payment_cycle`, `period_type`, `product_name`, `product_type`, `profile_status`, `recurring_payment_id`, `rp_invoice_id`, `time_created`, `amount1`, `amount2`, `amount3`, `mc_amount1`, `mc_amount2`, `mc_amount3`, `password`, `period1`, `period2`, `period3`, `reattempt`, `recur_times`, `recurring`, `retry_at`, `subscr_date`, `subscr_effective`, `subscr_id`, `username`, `mp_id`, `case_creation_date`, `case_id`, `case_type`, `receipt_id`, `currency_code`, `handling_amount`, `transaction_subject`, `ipaddress`, `flag`, `flag_code`, `flag_info`, `query`, `response`, `created_at`, `updated_at`, `from_view`.

Logs PayPal IPN interactions.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `business` | `CharField` | `varchar(127)` |  |  |  | Email where the money was sent. |
| `charset` | `CharField` | `varchar(255)` |  |  |  |  |
| `custom` | `CharField` | `varchar(256)` |  |  |  |  |
| `notify_version` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `parent_txn_id` | `CharField` | `varchar(19)` |  |  |  |  |
| `receiver_email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `receiver_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `residence_country` | `CharField` | `varchar(2)` |  |  |  |  |
| `test_ipn` | `BooleanField` | `bool` |  |  | `False` |  |
| `txn_id` | `CharField` | `varchar(255)` |  |  |  | PayPal transaction ID. |
| `txn_type` | `CharField` | `varchar(255)` |  |  |  | PayPal transaction type. |
| `verify_sign` | `CharField` | `varchar(255)` |  |  |  |  |
| `address_country` | `CharField` | `varchar(64)` |  |  |  |  |
| `address_city` | `CharField` | `varchar(40)` |  |  |  |  |
| `address_country_code` | `CharField` | `varchar(64)` |  |  |  | ISO 3166 |
| `address_name` | `CharField` | `varchar(128)` |  |  |  |  |
| `address_state` | `CharField` | `varchar(40)` |  |  |  |  |
| `address_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `address_street` | `CharField` | `varchar(200)` |  |  |  |  |
| `address_zip` | `CharField` | `varchar(20)` |  |  |  |  |
| `contact_phone` | `CharField` | `varchar(20)` |  |  |  |  |
| `first_name` | `CharField` | `varchar(64)` |  |  |  |  |
| `last_name` | `CharField` | `varchar(64)` |  |  |  |  |
| `payer_business_name` | `CharField` | `varchar(127)` |  |  |  |  |
| `payer_email` | `CharField` | `varchar(127)` |  |  |  |  |
| `payer_id` | `CharField` | `varchar(13)` |  |  |  |  |
| `auth_amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `auth_exp` | `CharField` | `varchar(28)` |  |  |  |  |
| `auth_id` | `CharField` | `varchar(19)` |  |  |  |  |
| `auth_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `exchange_rate` | `DecimalField(64, 16)` | `decimal` | yes |  | `0` |  |
| `invoice` | `CharField` | `varchar(127)` |  |  |  |  |
| `item_name` | `CharField` | `varchar(127)` |  |  |  |  |
| `item_number` | `CharField` | `varchar(127)` |  |  |  |  |
| `mc_currency` | `CharField` | `varchar(32)` |  |  | `'USD'` |  |
| `mc_fee` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_gross` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_handling` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_shipping` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `memo` | `CharField` | `varchar(255)` |  |  |  |  |
| `num_cart_items` | `IntegerField` | `integer` | yes |  | `0` |  |
| `option_name1` | `CharField` | `varchar(64)` |  |  |  |  |
| `option_name2` | `CharField` | `varchar(64)` |  |  |  |  |
| `option_selection1` | `CharField` | `varchar(200)` |  |  |  |  |
| `option_selection2` | `CharField` | `varchar(200)` |  |  |  |  |
| `payer_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `payment_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `payment_gross` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `payment_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `payment_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `pending_reason` | `CharField` | `varchar(255)` |  |  |  |  |
| `protection_eligibility` | `CharField` | `varchar(255)` |  |  |  |  |
| `quantity` | `IntegerField` | `integer` | yes |  | `1` |  |
| `reason_code` | `CharField` | `varchar(255)` |  |  |  |  |
| `remaining_settle` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `settle_amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `settle_currency` | `CharField` | `varchar(32)` |  |  |  |  |
| `shipping` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `shipping_method` | `CharField` | `varchar(255)` |  |  |  |  |
| `tax` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `transaction_entity` | `CharField` | `varchar(255)` |  |  |  |  |
| `auction_buyer_id` | `CharField` | `varchar(64)` |  |  |  |  |
| `auction_closing_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `auction_multi_item` | `IntegerField` | `integer` | yes |  | `0` |  |
| `for_auction` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `amount_per_cycle` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `initial_payment_amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `next_payment_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `outstanding_balance` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `payment_cycle` | `CharField` | `varchar(255)` |  |  |  |  |
| `period_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `product_name` | `CharField` | `varchar(255)` |  |  |  |  |
| `product_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `profile_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `recurring_payment_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `rp_invoice_id` | `CharField` | `varchar(127)` |  |  |  |  |
| `time_created` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `amount1` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `amount2` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `amount3` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_amount1` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_amount2` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_amount3` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `password` | `CharField` | `varchar(24)` |  |  |  |  |
| `period1` | `CharField` | `varchar(255)` |  |  |  |  |
| `period2` | `CharField` | `varchar(255)` |  |  |  |  |
| `period3` | `CharField` | `varchar(255)` |  |  |  |  |
| `reattempt` | `CharField` | `varchar(1)` |  |  |  |  |
| `recur_times` | `IntegerField` | `integer` | yes |  | `0` |  |
| `recurring` | `CharField` | `varchar(1)` |  |  |  |  |
| `retry_at` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `subscr_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `subscr_effective` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `subscr_id` | `CharField` | `varchar(19)` |  |  |  |  |
| `username` | `CharField` | `varchar(64)` |  |  |  |  |
| `mp_id` | `CharField` | `varchar(128)` | yes |  |  |  |
| `case_creation_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `case_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `case_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `receipt_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `currency_code` | `CharField` | `varchar(32)` |  |  | `'USD'` |  |
| `handling_amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `transaction_subject` | `CharField` | `varchar(256)` |  |  |  |  |
| `ipaddress` | `GenericIPAddressField` | `char(39)` | yes |  |  |  |
| `flag` | `BooleanField` | `bool` |  |  | `False` |  |
| `flag_code` | `CharField` | `varchar(16)` |  |  |  |  |
| `flag_info` | `TextField` | `text` |  |  |  |  |
| `query` | `TextField` | `text` |  |  |  |  |
| `response` | `TextField` | `text` |  |  |  |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `from_view` | `CharField` | `varchar(6)` | yes |  |  |  |

Indexes:

- index `paypal_ipn_txn_id_8fa22c44` on `txn_id`

## PayPalPDT

Table `paypal_pdt`, app `pdt`.

Columns from `PayPalStandardBase` (abstract): `business`, `charset`, `custom`, `notify_version`, `parent_txn_id`, `receiver_email`, `receiver_id`, `residence_country`, `test_ipn`, `txn_id`, `txn_type`, `verify_sign`, `address_country`, `address_city`, `address_country_code`, `address_name`, `address_state`, `address_status`, `address_street`, `address_zip`, `contact_phone`, `first_name`, `last_name`, `payer_business_name`, `payer_email`, `payer_id`, `auth_amount`, `auth_exp`, `auth_id`, `auth_status`, `exchange_rate`, `invoice`, `item_name`, `item_number`, `mc_currency`, `mc_fee`, `mc_gross`, `mc_handling`, `mc_shipping`, `memo`, `num_cart_items`, `option_name1`, `option_name2`, `option_selection1`, `option_selection2`, `payer_status`, `payment_date`, `payment_gross`, `payment_status`, `payment_type`, `pending_reason`, `protection_eligibility`, `quantity`, `reason_code`, `remaining_settle`, `settle_amount`, `settle_currency`, `shipping`, `shipping_method`, `tax`, `transaction_entity`, `auction_buyer_id`, `auction_closing_date`, `auction_multi_item`, `for_auction`, `amount`, `amount_per_cycle`, `initial_payment_amount`, `next_payment_date`, `outstanding_balance`, `payment_cycle`, `period_type`, `product_name`, `product_type`, `profile_status`, `recurring_payment_id`, `rp_invoice_id`, `time_created`, `amount1`, `amount2`, `amount3`, `mc_amount1`, `mc_amount2`, `mc_amount3`, `password`, `period1`, `period2`, `period3`, `reattempt`, `recur_times`, `recurring`, `retry_at`, `subscr_date`, `subscr_effective`, `subscr_id`, `username`, `mp_id`, `case_creation_date`, `case_id`, `case_type`, `receipt_id`, `currency_code`, `handling_amount`, `transaction_subject`, `ipaddress`, `flag`, `flag_code`, `flag_info`, `query`, `response`, `created_at`, `updated_at`, `from_view`.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `business` | `CharField` | `varchar(127)` |  |  |  | Email where the money was sent. |
| `charset` | `CharField` | `varchar(255)` |  |  |  |  |
| `custom` | `CharField` | `varchar(256)` |  |  |  |  |
| `notify_version` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `parent_txn_id` | `CharField` | `varchar(19)` |  |  |  |  |
| `receiver_email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `receiver_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `residence_country` | `CharField` | `varchar(2)` |  |  |  |  |
| `test_ipn` | `BooleanField` | `bool` |  |  | `False` |  |
| `txn_id` | `CharField` | `varchar(255)` |  |  |  | PayPal transaction ID. |
| `txn_type` | `CharField` | `varchar(255)` |  |  |  | PayPal transaction type. |
| `verify_sign` | `CharField` | `varchar(255)` |  |  |  |  |
| `address_country` | `CharField` | `varchar(64)` |  |  |  |  |
| `address_city` | `CharField` | `varchar(40)` |  |  |  |  |
| `address_country_code` | `CharField` | `varchar(64)` |  |  |  | ISO 3166 |
| `address_name` | `CharField` | `varchar(128)` |  |  |  |  |
| `address_state` | `CharField` | `varchar(40)` |  |  |  |  |
| `address_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `address_street` | `CharField` | `varchar(200)` |  |  |  |  |
| `address_zip` | `CharField` | `varchar(20)` |  |  |  |  |
| `contact_phone` | `CharField` | `varchar(20)` |  |  |  |  |
| `first_name` | `CharField` | `varchar(64)` |  |  |  |  |
| `last_name` | `CharField` | `varchar(64)` |  |  |  |  |
| `payer_business_name` | `CharField` | `varchar(127)` |  |  |  |  |
| `payer_email` | `CharField` | `varchar(127)` |  |  |  |  |
| `payer_id` | `CharField` | `varchar(13)` |  |  |  |  |
| `auth_amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `auth_exp` | `CharField` | `varchar(28)` |  |  |  |  |
| `auth_id` | `CharField` | `varchar(19)` |  |  |  |  |
| `auth_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `exchange_rate` | `DecimalField(64, 16)` | `decimal` | yes |  | `0` |  |
| `invoice` | `CharField` | `varchar(127)` |  |  |  |  |
| `item_name` | `CharField` | `varchar(127)` |  |  |  |  |
| `item_number` | `CharField` | `varchar(127)` |  |  |  |  |
| `mc_currency` | `CharField` | `varchar(32)` |  |  | `'USD'` |  |
| `mc_fee` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_gross` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_handling` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_shipping` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `memo` | `CharField` | `varchar(255)` |  |  |  |  |
| `num_cart_items` | `IntegerField` | `integer` | yes |  | `0` |  |
| `option_name1` | `CharField` | `varchar(64)` |  |  |  |  |
| `option_name2` | `CharField` | `varchar(64)` |  |  |  |  |
| `option_selection1` | `CharField` | `varchar(200)` |  |  |  |  |
| `option_selection2` | `CharField` | `varchar(200)` |  |  |  |  |
| `payer_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `payment_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `payment_gross` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `payment_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `payment_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `pending_reason` | `CharField` | `varchar(255)` |  |  |  |  |
| `protection_eligibility` | `CharField` | `varchar(255)` |  |  |  |  |
| `quantity` | `IntegerField` | `integer` | yes |  | `1` |  |
| `reason_code` | `CharField` | `varchar(255)` |  |  |  |  |
| `remaining_settle` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `settle_amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `settle_currency` | `CharField` | `varchar(32)` |  |  |  |  |
| `shipping` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `shipping_method` | `CharField` | `varchar(255)` |  |  |  |  |
| `tax` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `transaction_entity` | `CharField` | `varchar(255)` |  |  |  |  |
| `auction_buyer_id` | `CharField` | `varchar(64)` |  |  |  |  |
| `auction_closing_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `auction_multi_item` | `IntegerField` | `integer` | yes |  | `0` |  |
| `for_auction` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `amount_per_cycle` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `initial_payment_amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `next_payment_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `outstanding_balance` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `payment_cycle` | `CharField` | `varchar(255)` |  |  |  |  |
| `period_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `product_name` | `CharField` | `varchar(255)` |  |  |  |  |
| `product_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `profile_status` | `CharField` | `varchar(255)` |  |  |  |  |
| `recurring_payment_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `rp_invoice_id` | `CharField` | `varchar(127)` |  |  |  |  |
| `time_created` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `amount1` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `amount2` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `amount3` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_amount1` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_amount2` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `mc_amount3` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `password` | `CharField` | `varchar(24)` |  |  |  |  |
| `period1` | `CharField` | `varchar(255)` |  |  |  |  |
| `period2` | `CharField` | `varchar(255)` |  |  |  |  |
| `period3` | `CharField` | `varchar(255)` |  |  |  |  |
| `reattempt` | `CharField` | `varchar(1)` |  |  |  |  |
| `recur_times` | `IntegerField` | `integer` | yes |  | `0` |  |
| `recurring` | `CharField` | `varchar(1)` |  |  |  |  |
| `retry_at` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `subscr_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `subscr_effective` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `subscr_id` | `CharField` | `varchar(19)` |  |  |  |  |
| `username` | `CharField` | `varchar(64)` |  |  |  |  |
| `mp_id` | `CharField` | `varchar(128)` | yes |  |  |  |
| `case_creation_date` | `DateTimeField` | `datetime` | yes |  |  | HH:MM:SS DD Mmm YY, YYYY PST |
| `case_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `case_type` | `CharField` | `varchar(255)` |  |  |  |  |
| `receipt_id` | `CharField` | `varchar(255)` |  |  |  |  |
| `currency_code` | `CharField` | `varchar(32)` |  |  | `'USD'` |  |
| `handling_amount` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `transaction_subject` | `CharField` | `varchar(256)` |  |  |  |  |
| `ipaddress` | `GenericIPAddressField` | `char(39)` | yes |  |  |  |
| `flag` | `BooleanField` | `bool` |  |  | `False` |  |
| `flag_code` | `CharField` | `varchar(16)` |  |  |  |  |
| `flag_info` | `TextField` | `text` |  |  |  |  |
| `query` | `TextField` | `text` |  |  |  |  |
| `response` | `TextField` | `text` |  |  |  |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `updated_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `from_view` | `CharField` | `varchar(6)` | yes |  |  |  |
| `amt` | `DecimalField(64, 2)` | `decimal` | yes |  | `0` |  |
| `cm` | `CharField` | `varchar(255)` |  |  |  |  |
| `sig` | `CharField` | `varchar(255)` |  |  |  |  |
| `tx` | `CharField` | `varchar(255)` |  |  |  |  |
| `st` | `CharField` | `varchar(32)` |  |  |  |  |

Indexes:

- index `paypal_pdt_txn_id_0b75b30b` on `txn_id`

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |
