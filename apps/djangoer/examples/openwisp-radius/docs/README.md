# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [ContentType](#contenttype) · [Session](#session) · [Site](#site) · [EmailAddress](#emailaddress) · [EmailConfirmation](#emailconfirmation) · [SocialApp](#socialapp) · [SocialAccount](#socialaccount) · [SocialToken](#socialtoken) · [User](#user) · [Organization](#organization) · [OrganizationUser](#organizationuser) · [OrganizationOwner](#organizationowner) · [OrganizationInvitation](#organizationinvitation) · [Token](#token) · [RadiusCheck](#radiuscheck) · [RadiusReply](#radiusreply) · [RadiusAccounting](#radiusaccounting) · [RadiusGroup](#radiusgroup) · [RadiusGroupCheck](#radiusgroupcheck) · [RadiusGroupReply](#radiusgroupreply) · [RadiusUserGroup](#radiususergroup) · [RadiusPostAuth](#radiuspostauth) · [Nas](#nas) · [RadiusBatch](#radiusbatch) · [RadiusToken](#radiustoken) · [OrganizationRadiusSettings](#organizationradiussettings) · [PhoneToken](#phonetoken) · [RegisteredUser](#registereduser)

<details>
<summary>The whole schema as SQL</summary>

```sql
CREATE TABLE "django_admin_log" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "action_time" datetime NOT NULL, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "object_id" text NULL, "object_repr" varchar(200) NOT NULL, "action_flag" smallint unsigned NOT NULL CHECK ("action_flag" >= 0), "change_message" text NOT NULL);

CREATE TABLE "auth_permission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(255) NOT NULL, "content_type_id" integer NOT NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "codename" varchar(100) NOT NULL);

CREATE TABLE "auth_group" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(150) NOT NULL UNIQUE);

CREATE TABLE "auth_group_permissions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED, "permission_id" integer NOT NULL REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "django_content_type" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "app_label" varchar(100) NOT NULL, "model" varchar(100) NOT NULL);

CREATE TABLE "django_session" ("session_key" varchar(40) NOT NULL PRIMARY KEY, "session_data" text NOT NULL, "expire_date" datetime NOT NULL);

CREATE TABLE "django_site" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "domain" varchar(100) NOT NULL UNIQUE, "name" varchar(50) NOT NULL);

CREATE TABLE "account_emailaddress" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "email" varchar(254) NOT NULL, "verified" bool NOT NULL, "primary" bool NOT NULL);

CREATE TABLE "account_emailconfirmation" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "email_address_id" integer NOT NULL REFERENCES "account_emailaddress" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL, "sent" datetime NULL, "key" varchar(64) NOT NULL UNIQUE);

CREATE TABLE "socialaccount_socialapp" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "provider" varchar(30) NOT NULL, "provider_id" varchar(200) NOT NULL, "name" varchar(40) NOT NULL, "client_id" varchar(191) NOT NULL, "secret" varchar(191) NOT NULL, "key" varchar(191) NOT NULL, "settings" text NOT NULL CHECK ((JSON_VALID("settings") OR "settings" IS NULL)));

CREATE TABLE "socialaccount_socialapp_sites" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "socialapp_id" integer NOT NULL REFERENCES "socialaccount_socialapp" ("id") DEFERRABLE INITIALLY DEFERRED, "site_id" integer NOT NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "socialaccount_socialaccount" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "provider" varchar(200) NOT NULL, "uid" varchar(191) NOT NULL, "last_login" datetime NOT NULL, "date_joined" datetime NOT NULL, "extra_data" text NOT NULL CHECK ((JSON_VALID("extra_data") OR "extra_data" IS NULL)));

CREATE TABLE "socialaccount_socialtoken" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "app_id" integer NULL REFERENCES "socialaccount_socialapp" ("id") DEFERRABLE INITIALLY DEFERRED, "account_id" integer NOT NULL REFERENCES "socialaccount_socialaccount" ("id") DEFERRABLE INITIALLY DEFERRED, "token" text NOT NULL, "token_secret" text NOT NULL, "expires_at" datetime NULL);

CREATE TABLE "openwisp_users_user" ("password" varchar(128) NOT NULL, "last_login" datetime NULL, "is_superuser" bool NOT NULL, "username" varchar(150) NOT NULL UNIQUE, "first_name" varchar(150) NOT NULL, "last_name" varchar(150) NOT NULL, "is_staff" bool NOT NULL, "is_active" bool NOT NULL, "date_joined" datetime NOT NULL, "id" char(32) NOT NULL PRIMARY KEY, "email" varchar(254) NULL UNIQUE, "bio" text NOT NULL, "url" varchar(200) NOT NULL, "company" varchar(30) NOT NULL, "location" varchar(256) NOT NULL, "phone_number" varchar(128) NULL UNIQUE, "birth_date" date NULL, "notes" text NOT NULL, "language" varchar(8) NOT NULL, "password_updated" date NULL, "password_based_token" bool NULL, "expiration_date" date NULL);

CREATE TABLE "openwisp_users_user_groups" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NOT NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "openwisp_users_user_user_permissions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "permission_id" integer NOT NULL REFERENCES "auth_permission" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "openwisp_users_organization" ("name" varchar(200) NOT NULL, "is_active" bool NOT NULL, "created" datetime NOT NULL, "modified" datetime NOT NULL, "slug" varchar(200) NOT NULL UNIQUE, "id" char(32) NOT NULL PRIMARY KEY, "description" text NOT NULL, "email" varchar(254) NOT NULL, "url" varchar(200) NOT NULL);

CREATE TABLE "openwisp_users_organizationuser" ("created" datetime NOT NULL, "modified" datetime NOT NULL, "id" char(32) NOT NULL PRIMARY KEY, "is_admin" bool NOT NULL, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "openwisp_users_organizationowner" ("created" datetime NOT NULL, "modified" datetime NOT NULL, "id" char(32) NOT NULL PRIMARY KEY, "organization_user_id" char(32) NOT NULL UNIQUE REFERENCES "openwisp_users_organizationuser" ("id") DEFERRABLE INITIALLY DEFERRED, "organization_id" char(32) NOT NULL UNIQUE REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "openwisp_users_organizationinvitation" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "guid" char(32) NOT NULL, "invitee_identifier" varchar(1000) NOT NULL, "created" datetime NOT NULL, "modified" datetime NOT NULL, "invited_by_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "invitee_id" char(32) NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "authtoken_token" ("key" varchar(40) NOT NULL PRIMARY KEY, "user_id" char(32) NOT NULL UNIQUE REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL);

CREATE TABLE "radcheck" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "username" varchar(64) NOT NULL, "value" varchar(253) NOT NULL, "op" varchar(2) NOT NULL, "attribute" varchar(64) NOT NULL, "user_id" char(32) NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "radreply" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "username" varchar(64) NOT NULL, "value" varchar(253) NOT NULL, "op" varchar(2) NOT NULL, "attribute" varchar(64) NOT NULL, "user_id" char(32) NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "radacct" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "acctsessionid" varchar(64) NOT NULL, "acctuniqueid" varchar(32) NOT NULL PRIMARY KEY, "username" varchar(64) NULL, "groupname" varchar(64) NULL, "realm" varchar(64) NULL, "nasipaddress" char(39) NOT NULL, "nasportid" varchar(15) NULL, "nasporttype" varchar(32) NULL, "acctstarttime" datetime NULL, "acctupdatetime" datetime NULL, "acctstoptime" datetime NULL, "acctinterval" integer NULL, "acctsessiontime" integer unsigned NULL CHECK ("acctsessiontime" >= 0), "acctauthentic" varchar(32) NULL, "connectinfo_start" varchar(50) NULL, "connectinfo_stop" varchar(50) NULL, "acctinputoctets" bigint NULL, "acctoutputoctets" bigint NULL, "calledstationid" varchar(253) NULL, "callingstationid" varchar(253) NULL, "acctterminatecause" varchar(32) NULL, "servicetype" varchar(32) NULL, "framedprotocol" varchar(32) NULL, "framedipaddress" char(39) NULL, "framedipv6address" char(39) NULL, "framedipv6prefix" varchar(44) NULL, "framedinterfaceid" varchar(19) NULL, "delegatedipv6prefix" varchar(44) NULL);

CREATE TABLE "openwisp_radius_radiusgroup" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "name" varchar(255) NOT NULL UNIQUE, "description" varchar(64) NULL, "default" bool NOT NULL);

CREATE TABLE "radgroupcheck" ("id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "groupname" varchar(64) NOT NULL, "attribute" varchar(64) NOT NULL, "op" varchar(2) NOT NULL, "value" varchar(253) NOT NULL, "group_id" char(32) NULL REFERENCES "openwisp_radius_radiusgroup" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "radgroupreply" ("id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "groupname" varchar(64) NOT NULL, "attribute" varchar(64) NOT NULL, "op" varchar(2) NOT NULL, "value" varchar(253) NOT NULL, "group_id" char(32) NULL REFERENCES "openwisp_radius_radiusgroup" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "radusergroup" ("id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "username" varchar(64) NOT NULL, "groupname" varchar(64) NOT NULL, "priority" integer NOT NULL, "user_id" char(32) NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" char(32) NULL REFERENCES "openwisp_radius_radiusgroup" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "radpostauth" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "id" char(32) NOT NULL PRIMARY KEY, "username" varchar(64) NOT NULL, "pass" varchar(64) NOT NULL, "reply" varchar(32) NOT NULL, "calledstationid" varchar(253) NULL, "callingstationid" varchar(253) NULL, "authdate" datetime NOT NULL);

CREATE TABLE "nas" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "nasname" varchar(128) NOT NULL, "shortname" varchar(32) NOT NULL, "type" varchar(30) NOT NULL, "ports" integer unsigned NULL CHECK ("ports" >= 0), "secret" varchar(60) NOT NULL, "server" varchar(64) NULL, "community" varchar(50) NULL, "description" varchar(200) NULL);

CREATE TABLE "radbatch" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "strategy" varchar(16) NOT NULL, "status" varchar(16) NOT NULL, "name" varchar(128) NOT NULL, "csvfile" varchar(100) NULL, "prefix" varchar(20) NULL, "user_credentials" text NULL CHECK ((JSON_VALID("user_credentials") OR "user_credentials" IS NULL)), "expiration_date" date NULL, "group_id" char(32) NULL REFERENCES "openwisp_radius_radiusgroup" ("id") DEFERRABLE INITIALLY DEFERRED, "notes" text NOT NULL);

CREATE TABLE "radbatch_users" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "radiusbatch_id" char(32) NOT NULL REFERENCES "radbatch" ("id") DEFERRABLE INITIALLY DEFERRED, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "radiustoken" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "created" datetime NOT NULL, "key" varchar(40) NOT NULL PRIMARY KEY, "user_id" char(32) NOT NULL UNIQUE REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "can_auth" bool NOT NULL, "password_based" bool NULL);

CREATE TABLE "openwisp_radius_organizationradiussettings" ("id" char(32) NOT NULL PRIMARY KEY, "organization_id" char(32) NOT NULL UNIQUE REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "token" varchar(32) NOT NULL, "sms_verification" bool NULL, "needs_identity_verification" bool NULL, "sms_sender" varchar(128) NULL, "sms_message" text NULL, "sms_cooldown" integer unsigned NULL CHECK ("sms_cooldown" >= 0), "sms_meta_data" text NULL CHECK ((JSON_VALID("sms_meta_data") OR "sms_meta_data" IS NULL)), "freeradius_allowed_hosts" text NULL, "coa_enabled" bool NULL, "allowed_mobile_prefixes" text NULL, "first_name" varchar(12) NULL, "last_name" varchar(12) NULL, "location" varchar(12) NULL, "birth_date" varchar(12) NULL, "registration_enabled" bool NULL, "saml_registration_enabled" bool NULL, "mac_addr_roaming_enabled" bool NULL, "social_registration_enabled" bool NULL, "login_url" varchar(200) NULL, "status_url" varchar(200) NULL, "password_reset_url" varchar(200) NULL);

CREATE TABLE "openwisp_radius_phonetoken" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "id" char(32) NOT NULL PRIMARY KEY, "created" datetime NOT NULL, "modified" datetime NOT NULL, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "valid_until" datetime NOT NULL, "attempts" integer unsigned NOT NULL CHECK ("attempts" >= 0), "verified" bool NOT NULL, "token" varchar(8) NOT NULL, "ip" char(39) NOT NULL, "phone_number" varchar(128) NOT NULL);

CREATE TABLE "openwisp_radius_registereduser" ("organization_id" char(32) NOT NULL REFERENCES "openwisp_users_organization" ("id") DEFERRABLE INITIALLY DEFERRED, "id" char(32) NOT NULL PRIMARY KEY, "user_id" char(32) NOT NULL REFERENCES "openwisp_users_user" ("id") DEFERRABLE INITIALLY DEFERRED, "method" varchar(64) NOT NULL, "is_verified" bool NOT NULL, "modified" datetime NOT NULL, CONSTRAINT "unique_registered_user_per_org" UNIQUE ("user_id", "organization_id"));

CREATE INDEX "django_admin_log_user_id_c564eba6" ON "django_admin_log" ("user_id");

CREATE INDEX "django_admin_log_content_type_id_c4bce8eb" ON "django_admin_log" ("content_type_id");

CREATE UNIQUE INDEX "auth_permission_content_type_id_codename_01ab375a_uniq" ON "auth_permission" ("content_type_id", "codename");

CREATE INDEX "auth_permission_content_type_id_2f476e4b" ON "auth_permission" ("content_type_id");

CREATE UNIQUE INDEX "auth_group_permissions_group_id_permission_id_0cd325b0_uniq" ON "auth_group_permissions" ("group_id", "permission_id");

CREATE INDEX "auth_group_permissions_group_id_b120cbf9" ON "auth_group_permissions" ("group_id");

CREATE INDEX "auth_group_permissions_permission_id_84c5c92e" ON "auth_group_permissions" ("permission_id");

CREATE UNIQUE INDEX "django_content_type_app_label_model_76bd3d3b_uniq" ON "django_content_type" ("app_label", "model");

CREATE INDEX "django_session_expire_date_a5c62663" ON "django_session" ("expire_date");

CREATE UNIQUE INDEX "account_emailaddress_user_id_email_987c8728_uniq" ON "account_emailaddress" ("user_id", "email");

CREATE UNIQUE INDEX "unique_primary_email" ON "account_emailaddress" ("user_id", "primary") WHERE "primary";

CREATE UNIQUE INDEX "unique_verified_email" ON "account_emailaddress" ("email") WHERE "verified";

CREATE INDEX "account_emailaddress_user_id_2c513194" ON "account_emailaddress" ("user_id");

CREATE INDEX "account_emailaddress_email_03be32b2" ON "account_emailaddress" ("email");

CREATE INDEX "account_emailconfirmation_email_address_id_5b7f8c58" ON "account_emailconfirmation" ("email_address_id");

CREATE UNIQUE INDEX "socialaccount_socialapp_sites_socialapp_id_site_id_71a9a768_uniq" ON "socialaccount_socialapp_sites" ("socialapp_id", "site_id");

CREATE INDEX "socialaccount_socialapp_sites_socialapp_id_97fb6e7d" ON "socialaccount_socialapp_sites" ("socialapp_id");

CREATE INDEX "socialaccount_socialapp_sites_site_id_2579dee5" ON "socialaccount_socialapp_sites" ("site_id");

CREATE UNIQUE INDEX "socialaccount_socialaccount_provider_uid_fc810c6e_uniq" ON "socialaccount_socialaccount" ("provider", "uid");

CREATE INDEX "socialaccount_socialaccount_user_id_8146e70c" ON "socialaccount_socialaccount" ("user_id");

CREATE UNIQUE INDEX "socialaccount_socialtoken_app_id_account_id_fca4e0ac_uniq" ON "socialaccount_socialtoken" ("app_id", "account_id");

CREATE INDEX "socialaccount_socialtoken_app_id_636a42d7" ON "socialaccount_socialtoken" ("app_id");

CREATE INDEX "socialaccount_socialtoken_account_id_951f210e" ON "socialaccount_socialtoken" ("account_id");

CREATE INDEX "user_id_email_idx" ON "openwisp_users_user" ("id", "email");

CREATE INDEX "user_active_expiry_idx" ON "openwisp_users_user" ("is_active", "expiration_date");

CREATE UNIQUE INDEX "openwisp_users_user_groups_user_id_group_id_7b92002e_uniq" ON "openwisp_users_user_groups" ("user_id", "group_id");

CREATE INDEX "openwisp_users_user_groups_user_id_2371794b" ON "openwisp_users_user_groups" ("user_id");

CREATE INDEX "openwisp_users_user_groups_group_id_e9ab4901" ON "openwisp_users_user_groups" ("group_id");

CREATE UNIQUE INDEX "openwisp_users_user_user_permissions_user_id_permission_id_24ba4ccf_uniq" ON "openwisp_users_user_user_permissions" ("user_id", "permission_id");

CREATE INDEX "openwisp_users_user_user_permissions_user_id_f3610345" ON "openwisp_users_user_user_permissions" ("user_id");

CREATE INDEX "openwisp_users_user_user_permissions_permission_id_f00dc354" ON "openwisp_users_user_user_permissions" ("permission_id");

CREATE UNIQUE INDEX "openwisp_users_organizationuser_user_id_organization_id_d40e96a2_uniq" ON "openwisp_users_organizationuser" ("user_id", "organization_id");

CREATE INDEX "openwisp_users_organizationuser_user_id_a5799aa4" ON "openwisp_users_organizationuser" ("user_id");

CREATE INDEX "openwisp_users_organizationuser_organization_id_a1c43bbd" ON "openwisp_users_organizationuser" ("organization_id");

CREATE INDEX "openwisp_users_organizationinvitation_invited_by_id_69a7098a" ON "openwisp_users_organizationinvitation" ("invited_by_id");

CREATE INDEX "openwisp_users_organizationinvitation_invitee_id_d3c1b419" ON "openwisp_users_organizationinvitation" ("invitee_id");

CREATE INDEX "openwisp_users_organizationinvitation_organization_id_17466f8f" ON "openwisp_users_organizationinvitation" ("organization_id");

CREATE INDEX "radcheck_organization_id_305f0018" ON "radcheck" ("organization_id");

CREATE INDEX "radcheck_username_761f015c" ON "radcheck" ("username");

CREATE INDEX "radcheck_user_id_ccbdb2b3" ON "radcheck" ("user_id");

CREATE INDEX "radreply_organization_id_848b9c0b" ON "radreply" ("organization_id");

CREATE INDEX "radreply_username_e78f923c" ON "radreply" ("username");

CREATE INDEX "radreply_user_id_7ec33bf1" ON "radreply" ("user_id");

CREATE INDEX "radacct_organization_id_53a2e5bc" ON "radacct" ("organization_id");

CREATE INDEX "radacct_acctsessionid_67584b28" ON "radacct" ("acctsessionid");

CREATE INDEX "radacct_username_2b99049d" ON "radacct" ("username");

CREATE INDEX "radacct_nasipaddress_130a761d" ON "radacct" ("nasipaddress");

CREATE INDEX "radacct_acctstarttime_3e3f73a6" ON "radacct" ("acctstarttime");

CREATE INDEX "radacct_acctstoptime_edf090cf" ON "radacct" ("acctstoptime");

CREATE INDEX "radacct_calledstationid_459c474f" ON "radacct" ("calledstationid");

CREATE INDEX "radacct_callingstationid_574c1364" ON "radacct" ("callingstationid");

CREATE INDEX "openwisp_radius_radiusgroup_organization_id_93a59144" ON "openwisp_radius_radiusgroup" ("organization_id");

CREATE INDEX "radgroupcheck_groupname_e379d2bb" ON "radgroupcheck" ("groupname");

CREATE INDEX "radgroupcheck_group_id_2ff27220" ON "radgroupcheck" ("group_id");

CREATE INDEX "radgroupreply_groupname_47789679" ON "radgroupreply" ("groupname");

CREATE INDEX "radgroupreply_group_id_fbded6b3" ON "radgroupreply" ("group_id");

CREATE UNIQUE INDEX "radusergroup_user_id_group_id_3d9a70aa_uniq" ON "radusergroup" ("user_id", "group_id");

CREATE INDEX "radusergroup_username_743b42e3" ON "radusergroup" ("username");

CREATE INDEX "radusergroup_user_id_5a7fd66b" ON "radusergroup" ("user_id");

CREATE INDEX "radusergroup_group_id_0dbc5c53" ON "radusergroup" ("group_id");

CREATE INDEX "radpostauth_organization_id_de44db05" ON "radpostauth" ("organization_id");

CREATE INDEX "nas_organization_id_a0932128" ON "nas" ("organization_id");

CREATE INDEX "nas_nasname_2a874c39" ON "nas" ("nasname");

CREATE UNIQUE INDEX "radbatch_name_organization_id_82270c3b_uniq" ON "radbatch" ("name", "organization_id");

CREATE INDEX "radbatch_organization_id_d184d391" ON "radbatch" ("organization_id");

CREATE INDEX "radbatch_strategy_1fe16152" ON "radbatch" ("strategy");

CREATE INDEX "radbatch_status_35f74861" ON "radbatch" ("status");

CREATE INDEX "radbatch_name_15eac8ea" ON "radbatch" ("name");

CREATE INDEX "radbatch_group_id_8bd9d0a5" ON "radbatch" ("group_id");

CREATE UNIQUE INDEX "radbatch_users_radiusbatch_id_user_id_8f828365_uniq" ON "radbatch_users" ("radiusbatch_id", "user_id");

CREATE INDEX "radbatch_users_radiusbatch_id_9f29169c" ON "radbatch_users" ("radiusbatch_id");

CREATE INDEX "radbatch_users_user_id_bcd57916" ON "radbatch_users" ("user_id");

CREATE INDEX "radiustoken_organization_id_8e1824bf" ON "radiustoken" ("organization_id");

CREATE INDEX "openwisp_radius_phonetoken_organization_id_5d939c4a" ON "openwisp_radius_phonetoken" ("organization_id");

CREATE INDEX "openwisp_radius_phonetoken_user_id_90d96493" ON "openwisp_radius_phonetoken" ("user_id");

CREATE INDEX "openwisp_ra_user_id_9fe207_idx" ON "openwisp_radius_phonetoken" ("user_id", "created");

CREATE INDEX "openwisp_ra_user_id_d4dd52_idx" ON "openwisp_radius_phonetoken" ("user_id", "created", "ip");

CREATE INDEX "openwisp_radius_registereduser_organization_id_b1772588" ON "openwisp_radius_registereduser" ("organization_id");

CREATE INDEX "openwisp_radius_registereduser_user_id_3af63f27" ON "openwisp_radius_registereduser" ("user_id");
```

</details>

## LogEntry

Table `django_admin_log`, app `admin`.

![LogEntry and the models it relates to](tables/LogEntry.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `action_time` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
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

Proxy models, on this table: `Group`.

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

![Site and the models it relates to](tables/Site.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `domain` | `CharField` | `varchar(100)` |  | unique |  |  |
| `name` | `CharField` | `varchar(50)` |  |  |  |  |

Relations:

- `socialapp_set` — many to many with [SocialApp](#socialapp)

## EmailAddress

Table `account_emailaddress`, app `account`.

![EmailAddress and the models it relates to](tables/EmailAddress.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `verified` | `BooleanField` | `bool` |  |  | `False` |  |
| `primary` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `account_emailaddress_user_id_2c513194` on `user_id`
- index `account_emailaddress_email_03be32b2` on `email`
- unique `account_emailaddress_user_id_email_987c8728_uniq` on `user_id`, `email`
- unique `unique_primary_email` on `user_id`, `primary` where `"primary"`
- unique `unique_verified_email` on `email` where `"verified"`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [EmailConfirmation](#emailconfirmation).`email_address_id` → `id`, on delete cascade

## EmailConfirmation

Table `account_emailconfirmation`, app `account`.

![EmailConfirmation and the models it relates to](tables/EmailConfirmation.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `email_address_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `sent` | `DateTimeField` | `datetime` | yes |  |  |  |
| `key` | `CharField` | `varchar(64)` |  | unique |  |  |

Indexes:

- index `account_emailconfirmation_email_address_id_5b7f8c58` on `email_address_id`

Relations:

- `email_address_id` → [EmailAddress](#emailaddress).`id`, on delete cascade

## SocialApp

Table `socialaccount_socialapp`, app `socialaccount`.

![SocialApp and the models it relates to](tables/SocialApp.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `provider` | `CharField` | `varchar(30)` |  |  |  |  |
| `provider_id` | `CharField` | `varchar(200)` |  |  |  |  |
| `name` | `CharField` | `varchar(40)` |  |  |  |  |
| `client_id` | `CharField` | `varchar(191)` |  |  |  | App ID, or consumer key |
| `secret` | `CharField` | `varchar(191)` |  |  |  | API secret, client secret, or consumer secret |
| `key` | `CharField` | `varchar(191)` |  |  |  | Key |
| `settings` | `JSONField` | `text` |  |  | `dict()` |  |

Relations:

- `sites` — many to many with [Site](#site)
- [SocialToken](#socialtoken).`app_id` → `id`, on delete set null

## SocialAccount

Table `socialaccount_socialaccount`, app `socialaccount`.

![SocialAccount and the models it relates to](tables/SocialAccount.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `provider` | `CharField` | `varchar(200)` |  |  |  |  |
| `uid` | `CharField` | `varchar(191)` |  |  |  |  |
| `last_login` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `date_joined` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `extra_data` | `JSONField` | `text` |  |  | `dict()` |  |

Indexes:

- index `socialaccount_socialaccount_user_id_8146e70c` on `user_id`
- unique `socialaccount_socialaccount_provider_uid_fc810c6e_uniq` on `provider`, `uid`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- [SocialToken](#socialtoken).`account_id` → `id`, on delete cascade

## SocialToken

Table `socialaccount_socialtoken`, app `socialaccount`.

![SocialToken and the models it relates to](tables/SocialToken.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `app_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `account_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `token` | `TextField` | `text` |  |  |  | "oauth_token" (OAuth1) or access token (OAuth2) |
| `token_secret` | `TextField` | `text` |  |  |  | "oauth_token_secret" (OAuth1) or refresh token (OAuth2) |
| `expires_at` | `DateTimeField` | `datetime` | yes |  |  |  |

Indexes:

- index `socialaccount_socialtoken_app_id_636a42d7` on `app_id`
- index `socialaccount_socialtoken_account_id_951f210e` on `account_id`
- unique `socialaccount_socialtoken_app_id_account_id_fca4e0ac_uniq` on `app_id`, `account_id`

Relations:

- `app_id` → [SocialApp](#socialapp).`id`, on delete set null
- `account_id` → [SocialAccount](#socialaccount).`id`, on delete cascade

## User

Table `openwisp_users_user`, app `openwisp_users`.

Columns from `AbstractUser` (abstract): `id`, `bio`, `url`, `company`, `location`, `phone_number`, `birth_date`, `notes`, `language`, `password_updated`, `password_based_token`, `expiration_date`.

Columns from `AbstractUser` (abstract): `username`, `first_name`, `last_name`, `email`, `is_staff`, `is_active`, `date_joined`.

Columns from `AbstractBaseUser` (abstract): `password`, `last_login`.

Columns from `PermissionsMixin` (abstract): `is_superuser`.

![User and the models it relates to](tables/User.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `password` | `CharField` | `varchar(128)` |  |  |  |  |
| `last_login` | `DateTimeField` | `datetime` | yes |  |  |  |
| `is_superuser` | `BooleanField` | `bool` |  |  | `False` | Designates that this user has all permissions without explicitly assigning them. |
| `username` | `CharField` | `varchar(150)` |  | unique |  | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. |
| `first_name` | `CharField` | `varchar(150)` |  |  |  |  |
| `last_name` | `CharField` | `varchar(150)` |  |  |  |  |
| `is_staff` | `BooleanField` | `bool` |  |  | `False` | Designates whether the user can log into this admin site. |
| `is_active` | `BooleanField` | `bool` |  |  | `True` | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. |
| `date_joined` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `email` | `EmailField` | `varchar(254)` | yes | unique |  |  |
| `bio` | `TextField` | `text` |  |  |  |  |
| `url` | `URLField` | `varchar(200)` |  |  |  |  |
| `company` | `CharField` | `varchar(30)` |  |  |  |  |
| `location` | `CharField` | `varchar(256)` |  |  |  |  |
| `phone_number` | `PhoneNumberField` | `varchar(128)` | yes | unique |  |  |
| `birth_date` | `DateField` | `date` | yes |  |  |  |
| `notes` | `TextField` | `text` |  |  |  | notes for internal usage |
| `language` | `UserLanguage` | `varchar(8)` |  |  | `'en-us'` |  |
| `password_updated` | `DateField` | `date` | yes |  |  |  |
| `password_based_token` | `BooleanField` | `bool` | yes |  | `None` | Indicates whether the last authentication token was obtained using the local password. When false, the token came from an external method (eg: SSO, SAML) and password expiration is not enforced for it. None means no token has been issued for this user since this feature was introduced. |
| `expiration_date` | `DateField` | `date` | yes |  |  | If set, the account will be deactivated after this date passes and the user will no longer be able to log in. |

Indexes:

- index `user_id_email_idx` on `id`, `email`
- index `user_active_expiry_idx` on `is_active`, `expiration_date`

Relations:

- `openwisp_users_organization` — many to many with [Organization](#organization), through [OrganizationUser](#organizationuser)
- [LogEntry](#logentry).`user_id` → `id`, on delete cascade
- [EmailAddress](#emailaddress).`user_id` → `id`, on delete cascade
- [SocialAccount](#socialaccount).`user_id` → `id`, on delete cascade
- `groups` — many to many with [Group](#group)
- `user_permissions` — many to many with [Permission](#permission)
- [OrganizationUser](#organizationuser).`user_id` → `id`, on delete cascade
- [OrganizationInvitation](#organizationinvitation).`invited_by_id` → `id`, on delete cascade
- [OrganizationInvitation](#organizationinvitation).`invitee_id` → `id`, on delete cascade
- [Token](#token).`user_id` → `id`, on delete cascade
- [RadiusCheck](#radiuscheck).`user_id` → `id`, on delete cascade
- [RadiusReply](#radiusreply).`user_id` → `id`, on delete cascade
- [RadiusUserGroup](#radiususergroup).`user_id` → `id`, on delete cascade
- `radius_batch` — many to many with [RadiusBatch](#radiusbatch)
- [RadiusToken](#radiustoken).`user_id` → `id`, on delete cascade
- [PhoneToken](#phonetoken).`user_id` → `id`, on delete cascade
- [RegisteredUser](#registereduser).`user_id` → `id`, on delete cascade

## Organization

Table `openwisp_users_organization`, app `openwisp_users`.

Columns from `BaseOrganization` (abstract): `id`, `description`, `email`, `url`.

Columns from `AbstractOrganization` (abstract): `slug`.

Columns from `SharedBaseModel` (abstract): `created`, `modified`.

Columns from `AbstractBaseOrganization` (abstract): `name`, `is_active`.

![Organization and the models it relates to](tables/Organization.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `name` | `CharField` | `varchar(200)` |  |  |  | The name of the organization |
| `is_active` | `BooleanField` | `bool` |  |  | `True` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `slug` | `SlugField` | `varchar(200)` |  | unique |  | The name in all lowercase, suitable for URL identification |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `email` | `EmailField` | `varchar(254)` |  |  |  |  |
| `url` | `URLField` | `varchar(200)` |  |  |  |  |

Relations:

- `users` — many to many with [User](#user), through [OrganizationUser](#organizationuser)
- [OrganizationUser](#organizationuser).`organization_id` → `id`, on delete cascade
- [OrganizationOwner](#organizationowner).`organization_id` → `id`, on delete cascade
- [OrganizationInvitation](#organizationinvitation).`organization_id` → `id`, on delete cascade
- [RadiusCheck](#radiuscheck).`organization_id` → `id`, on delete cascade
- [RadiusReply](#radiusreply).`organization_id` → `id`, on delete cascade
- [RadiusAccounting](#radiusaccounting).`organization_id` → `id`, on delete cascade
- [RadiusGroup](#radiusgroup).`organization_id` → `id`, on delete cascade
- [RadiusPostAuth](#radiuspostauth).`organization_id` → `id`, on delete cascade
- [Nas](#nas).`organization_id` → `id`, on delete cascade
- [RadiusBatch](#radiusbatch).`organization_id` → `id`, on delete cascade
- [RadiusToken](#radiustoken).`organization_id` → `id`, on delete cascade
- [OrganizationRadiusSettings](#organizationradiussettings).`organization_id` → `id`, on delete cascade
- [PhoneToken](#phonetoken).`organization_id` → `id`, on delete cascade
- [RegisteredUser](#registereduser).`organization_id` → `id`, on delete cascade

## OrganizationUser

Table `openwisp_users_organizationuser`, app `openwisp_users`.

Columns from `BaseOrganizationUser` (abstract): `id`.

Columns from `AbstractOrganizationUser` (abstract): `is_admin`.

Columns from `SharedBaseModel` (abstract): `created`, `modified`.

![OrganizationUser and the models it relates to](tables/OrganizationUser.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `is_admin` | `BooleanField` | `bool` |  |  | `False` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |

Indexes:

- index `openwisp_users_organizationuser_user_id_a5799aa4` on `user_id`
- index `openwisp_users_organizationuser_organization_id_a1c43bbd` on `organization_id`
- unique `openwisp_users_organizationuser_user_id_organization_id_d40e96a2_uniq` on `user_id`, `organization_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `organization_id` → [Organization](#organization).`id`, on delete cascade
- [OrganizationOwner](#organizationowner).`organization_user_id` → `id`, on delete cascade

## OrganizationOwner

Table `openwisp_users_organizationowner`, app `openwisp_users`.

Columns from `BaseOrganizationOwner` (abstract): `id`.

Columns from `SharedBaseModel` (abstract): `created`, `modified`.

![OrganizationOwner and the models it relates to](tables/OrganizationOwner.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `organization_user_id` | `OneToOneField` | `char(32)` |  | FK, unique |  |  |
| `organization_id` | `OneToOneField` | `char(32)` |  | FK, unique |  |  |

Relations:

- `organization_user_id` → [OrganizationUser](#organizationuser).`id`, on delete cascade
- `organization_id` → [Organization](#organization).`id`, on delete cascade

## OrganizationInvitation

Table `openwisp_users_organizationinvitation`, app `openwisp_users`.

Columns from `SharedBaseModel` (abstract): `created`, `modified`.

Columns from `AbstractBaseInvitation` (abstract): `guid`, `invitee_identifier`.

![OrganizationInvitation and the models it relates to](tables/OrganizationInvitation.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `guid` | `UUIDField` | `char(32)` |  |  |  |  |
| `invitee_identifier` | `CharField` | `varchar(1000)` |  |  |  | The contact identifier for the invitee, email, phone number, social media handle, etc. |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `invited_by_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `invitee_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |

Indexes:

- index `openwisp_users_organizationinvitation_invited_by_id_69a7098a` on `invited_by_id`
- index `openwisp_users_organizationinvitation_invitee_id_d3c1b419` on `invitee_id`
- index `openwisp_users_organizationinvitation_organization_id_17466f8f` on `organization_id`

Relations:

- `invited_by_id` → [User](#user).`id`, on delete cascade
- `invitee_id` → [User](#user).`id`, on delete cascade
- `organization_id` → [Organization](#organization).`id`, on delete cascade

## Token

Table `authtoken_token`, app `authtoken`.

Proxy models, on this table: `ApiKey`, `TokenProxy`.

The default authorization token model.

![Token and the models it relates to](tables/Token.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `key` | `CharField` | `varchar(40)` |  | PK |  |  |
| `user_id` | `OneToOneField` | `char(32)` |  | FK, unique |  |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Relations:

- `user_id` → [User](#user).`id`, on delete cascade

## RadiusCheck

Table `radcheck`, app `openwisp_radius`.

Columns from `AbstractRadiusCheck` (abstract): `username`, `value`, `op`, `attribute`, `user_id`.

Columns from `OrgMixin` (abstract): `organization_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![RadiusCheck and the models it relates to](tables/RadiusCheck.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `username` | `CharField` | `varchar(64)` |  |  |  |  |
| `value` | `CharField` | `varchar(253)` |  |  |  |  |
| `op` | `RadiusCheckOp` | `varchar(2)` |  |  | `':='` |  |
| `attribute` | `CharField` | `varchar(64)` |  |  |  |  |
| `user_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |

Indexes:

- index `radcheck_organization_id_305f0018` on `organization_id`
- index `radcheck_username_761f015c` on `username`
- index `radcheck_user_id_ccbdb2b3` on `user_id`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## RadiusReply

Table `radreply`, app `openwisp_radius`.

Columns from `AbstractRadiusReply` (abstract): `username`, `value`, `op`, `attribute`, `user_id`.

Columns from `OrgMixin` (abstract): `organization_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![RadiusReply and the models it relates to](tables/RadiusReply.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `username` | `CharField` | `varchar(64)` |  |  |  |  |
| `value` | `CharField` | `varchar(253)` |  |  |  |  |
| `op` | `RadiusReplyOp` | `varchar(2)` |  |  | `'='` |  |
| `attribute` | `CharField` | `varchar(64)` |  |  |  |  |
| `user_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |

Indexes:

- index `radreply_organization_id_848b9c0b` on `organization_id`
- index `radreply_username_e78f923c` on `username`
- index `radreply_user_id_7ec33bf1` on `user_id`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## RadiusAccounting

Table `radacct`, app `openwisp_radius`.

Columns from `AbstractRadiusAccounting` (abstract): `session_id`, `unique_id`, `username`, `groupname`, `realm`, `nas_ip_address`, `nas_port_id`, `nas_port_type`, `start_time`, `update_time`, `stop_time`, `interval`, `session_time`, `authentication`, `connection_info_start`, `connection_info_stop`, `input_octets`, `output_octets`, `called_station_id`, `calling_station_id`, `terminate_cause`, `service_type`, `framed_protocol`, `framed_ip_address`, `framed_ipv6_address`, `framed_ipv6_prefix`, `framed_interface_id`, `delegated_ipv6_prefix`.

Columns from `OrgMixin` (abstract): `organization_id`.

![RadiusAccounting and the models it relates to](tables/RadiusAccounting.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `acctsessionid` | `CharField` | `varchar(64)` |  |  |  |  |
| `acctuniqueid` | `CharField` | `varchar(32)` |  | PK |  |  |
| `username` | `CharField` | `varchar(64)` | yes |  |  |  |
| `groupname` | `CharField` | `varchar(64)` | yes |  |  |  |
| `realm` | `CharField` | `varchar(64)` | yes |  |  |  |
| `nasipaddress` | `GenericIPAddressField` | `char(39)` |  |  |  |  |
| `nasportid` | `CharField` | `varchar(15)` | yes |  |  |  |
| `nasporttype` | `CharField` | `varchar(32)` | yes |  |  |  |
| `acctstarttime` | `DateTimeField` | `datetime` | yes |  |  |  |
| `acctupdatetime` | `DateTimeField` | `datetime` | yes |  |  |  |
| `acctstoptime` | `DateTimeField` | `datetime` | yes |  |  |  |
| `acctinterval` | `IntegerField` | `integer` | yes |  |  |  |
| `acctsessiontime` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |
| `acctauthentic` | `CharField` | `varchar(32)` | yes |  |  |  |
| `connectinfo_start` | `CharField` | `varchar(50)` | yes |  |  |  |
| `connectinfo_stop` | `CharField` | `varchar(50)` | yes |  |  |  |
| `acctinputoctets` | `BigIntegerField` | `bigint` | yes |  |  |  |
| `acctoutputoctets` | `BigIntegerField` | `bigint` | yes |  |  |  |
| `calledstationid` | `CharField` | `varchar(253)` | yes |  |  |  |
| `callingstationid` | `CharField` | `varchar(253)` | yes |  |  |  |
| `acctterminatecause` | `CharField` | `varchar(32)` | yes |  |  |  |
| `servicetype` | `CharField` | `varchar(32)` | yes |  |  |  |
| `framedprotocol` | `CharField` | `varchar(32)` | yes |  |  |  |
| `framedipaddress` | `GenericIPAddressField` | `char(39)` | yes |  |  |  |
| `framedipv6address` | `GenericIPAddressField` | `char(39)` | yes |  |  |  |
| `framedipv6prefix` | `CharField` | `varchar(44)` | yes |  |  |  |
| `framedinterfaceid` | `CharField` | `varchar(19)` | yes |  |  |  |
| `delegatedipv6prefix` | `CharField` | `varchar(44)` | yes |  |  |  |

Indexes:

- index `radacct_organization_id_53a2e5bc` on `organization_id`
- index `radacct_acctsessionid_67584b28` on `session_id`
- index `radacct_username_2b99049d` on `username`
- index `radacct_nasipaddress_130a761d` on `nas_ip_address`
- index `radacct_acctstarttime_3e3f73a6` on `start_time`
- index `radacct_acctstoptime_edf090cf` on `stop_time`
- index `radacct_calledstationid_459c474f` on `called_station_id`
- index `radacct_callingstationid_574c1364` on `calling_station_id`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade

## RadiusGroup

Table `openwisp_radius_radiusgroup`, app `openwisp_radius`.

Columns from `AbstractRadiusGroup` (abstract): `name`, `description`, `default`.

Columns from `OrgMixin` (abstract): `organization_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![RadiusGroup and the models it relates to](tables/RadiusGroup.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `name` | `CharField` | `varchar(255)` |  | unique |  |  |
| `description` | `CharField` | `varchar(64)` | yes |  |  |  |
| `default` | `BooleanField` | `bool` |  |  | `False` | The default group is automatically assigned to new users; changing the default group has only effect on new users (existing users will keep being members of their current group) |

Indexes:

- index `openwisp_radius_radiusgroup_organization_id_93a59144` on `organization_id`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade
- [RadiusGroupCheck](#radiusgroupcheck).`group_id` → `id`, on delete cascade
- [RadiusGroupReply](#radiusgroupreply).`group_id` → `id`, on delete cascade
- [RadiusUserGroup](#radiususergroup).`group_id` → `id`, on delete cascade
- [RadiusBatch](#radiusbatch).`group_id` → `id`, on delete set null

## RadiusGroupCheck

Table `radgroupcheck`, app `openwisp_radius`.

Columns from `AbstractRadiusGroupCheck` (abstract): `groupname`, `attribute`, `op`, `value`, `group_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![RadiusGroupCheck and the models it relates to](tables/RadiusGroupCheck.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `groupname` | `CharField` | `varchar(64)` |  |  |  |  |
| `attribute` | `CharField` | `varchar(64)` |  |  |  |  |
| `op` | `RadiusGroupCheckOp` | `varchar(2)` |  |  | `':='` |  |
| `value` | `CharField` | `varchar(253)` |  |  |  |  |
| `group_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |

Indexes:

- index `radgroupcheck_groupname_e379d2bb` on `groupname`
- index `radgroupcheck_group_id_2ff27220` on `group_id`

Relations:

- `group_id` → [RadiusGroup](#radiusgroup).`id`, on delete cascade

## RadiusGroupReply

Table `radgroupreply`, app `openwisp_radius`.

Columns from `AbstractRadiusGroupReply` (abstract): `groupname`, `attribute`, `op`, `value`, `group_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![RadiusGroupReply and the models it relates to](tables/RadiusGroupReply.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `groupname` | `CharField` | `varchar(64)` |  |  |  |  |
| `attribute` | `CharField` | `varchar(64)` |  |  |  |  |
| `op` | `RadiusGroupReplyOp` | `varchar(2)` |  |  | `'='` |  |
| `value` | `CharField` | `varchar(253)` |  |  |  |  |
| `group_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |

Indexes:

- index `radgroupreply_groupname_47789679` on `groupname`
- index `radgroupreply_group_id_fbded6b3` on `group_id`

Relations:

- `group_id` → [RadiusGroup](#radiusgroup).`id`, on delete cascade

## RadiusUserGroup

Table `radusergroup`, app `openwisp_radius`.

Columns from `AbstractRadiusUserGroup` (abstract): `username`, `groupname`, `priority`, `user_id`, `group_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![RadiusUserGroup and the models it relates to](tables/RadiusUserGroup.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `username` | `CharField` | `varchar(64)` |  |  |  |  |
| `groupname` | `CharField` | `varchar(64)` |  |  |  |  |
| `priority` | `IntegerField` | `integer` |  |  | `1` |  |
| `user_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |
| `group_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |

Indexes:

- index `radusergroup_username_743b42e3` on `username`
- index `radusergroup_user_id_5a7fd66b` on `user_id`
- index `radusergroup_group_id_0dbc5c53` on `group_id`
- unique `radusergroup_user_id_group_id_3d9a70aa_uniq` on `user_id`, `group_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `group_id` → [RadiusGroup](#radiusgroup).`id`, on delete cascade

## RadiusPostAuth

Table `radpostauth`, app `openwisp_radius`.

Columns from `AbstractRadiusPostAuth` (abstract): `username`, `password`, `reply`, `called_station_id`, `calling_station_id`, `date`.

Columns from `OrgMixin` (abstract): `organization_id`.

Columns from `UUIDModel` (abstract): `id`.

![RadiusPostAuth and the models it relates to](tables/RadiusPostAuth.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `username` | `CharField` | `varchar(64)` |  |  |  |  |
| `pass` | `CharField` | `varchar(64)` |  |  |  |  |
| `reply` | `CharField` | `varchar(32)` |  |  |  |  |
| `calledstationid` | `CharField` | `varchar(253)` | yes |  |  |  |
| `callingstationid` | `CharField` | `varchar(253)` | yes |  |  |  |
| `authdate` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `radpostauth_organization_id_de44db05` on `organization_id`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade

## Nas

Table `nas`, app `openwisp_radius`.

Columns from `AbstractNas` (abstract): `name`, `short_name`, `type`, `ports`, `secret`, `server`, `community`, `description`.

Columns from `OrgMixin` (abstract): `organization_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![Nas and the models it relates to](tables/Nas.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `nasname` | `CharField` | `varchar(128)` |  |  |  | NAS Name (or IP address) |
| `shortname` | `CharField` | `varchar(32)` |  |  |  |  |
| `type` | `NasType` | `varchar(30)` |  |  | `'other'` |  |
| `ports` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |
| `secret` | `CharField` | `varchar(60)` |  |  |  | Shared Secret |
| `server` | `CharField` | `varchar(64)` | yes |  |  |  |
| `community` | `CharField` | `varchar(50)` | yes |  |  |  |
| `description` | `CharField` | `varchar(200)` | yes |  |  |  |

Indexes:

- index `nas_organization_id_a0932128` on `organization_id`
- index `nas_nasname_2a874c39` on `name`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade

## RadiusBatch

Table `radbatch`, app `openwisp_radius`.

Columns from `AbstractRadiusBatch` (abstract): `strategy`, `status`, `name`, `csvfile`, `prefix`, `user_credentials`, `expiration_date`, `group_id`, `notes`.

Columns from `OrgMixin` (abstract): `organization_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![RadiusBatch and the models it relates to](tables/RadiusBatch.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `strategy` | `RadiusBatchStrategy` | `varchar(16)` |  |  |  | Import users from a CSV or generate using a prefix |
| `status` | `RadiusBatchStatus` | `varchar(16)` |  |  | `'pending'` |  |
| `name` | `CharField` | `varchar(128)` |  |  |  | A unique batch name |
| `csvfile` | `PrivateFileField` | `varchar(100)` | yes |  |  | The csv file containing the user details to be uploaded |
| `prefix` | `CharField` | `varchar(20)` | yes |  |  | Usernames generated will be of the format [prefix][number] |
| `user_credentials` | `JSONField` | `text` | yes |  |  |  |
| `expiration_date` | `DateField` | `date` | yes |  |  | If left blank users will never expire |
| `group_id` | `ForeignKey` | `char(32)` | yes | FK |  |  |
| `notes` | `TextField` | `text` |  |  |  | internal notes |

Indexes:

- index `radbatch_organization_id_d184d391` on `organization_id`
- index `radbatch_strategy_1fe16152` on `strategy`
- index `radbatch_status_35f74861` on `status`
- index `radbatch_name_15eac8ea` on `name`
- index `radbatch_group_id_8bd9d0a5` on `group_id`
- unique `radbatch_name_organization_id_82270c3b_uniq` on `name`, `organization_id`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade
- `group_id` → [RadiusGroup](#radiusgroup).`id`, on delete set null
- `users` — many to many with [User](#user)

## RadiusToken

Table `radiustoken`, app `openwisp_radius`.

Columns from `AbstractRadiusToken` (abstract): `key`, `user_id`, `can_auth`, `password_based`.

Columns from `OrgMixin` (abstract): `organization_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`.

![RadiusToken and the models it relates to](tables/RadiusToken.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `key` | `CharField` | `varchar(40)` |  | PK |  |  |
| `user_id` | `OneToOneField` | `char(32)` |  | FK, unique |  |  |
| `can_auth` | `BooleanField` | `bool` |  |  | `False` | Enable the radius token to be used for freeradius authorization request |
| `password_based` | `BooleanField` | `bool` | yes |  | `None` | Indicates whether the user's most recent authentication used the local password. When false, the radius token is treated as externally authenticated (eg: SSO, SAML), so the local password expiration policy does not apply to it. |

Indexes:

- index `radiustoken_organization_id_8e1824bf` on `organization_id`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## OrganizationRadiusSettings

Table `openwisp_radius_organizationradiussettings`, app `openwisp_radius`.

Columns from `AbstractOrganizationRadiusSettings` (abstract): `organization_id`, `token`, `sms_verification`, `needs_identity_verification`, `sms_sender`, `sms_message`, `sms_cooldown`, `sms_meta_data`, `freeradius_allowed_hosts`, `coa_enabled`, `allowed_mobile_prefixes`, `first_name`, `last_name`, `location`, `birth_date`, `registration_enabled`, `saml_registration_enabled`, `mac_addr_roaming_enabled`, `social_registration_enabled`, `login_url`, `status_url`, `password_reset_url`.

Columns from `UUIDModel` (abstract): `id`.

![OrganizationRadiusSettings and the models it relates to](tables/OrganizationRadiusSettings.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `organization_id` | `OneToOneField` | `char(32)` |  | FK, unique |  |  |
| `token` | `KeyField` | `varchar(32)` |  |  | `get_random_key()` |  |
| `sms_verification` | `FallbackBooleanChoiceField` | `bool` | yes |  | `None` | Whether users who sign up should be required to verify their mobile phone number via SMS |
| `needs_identity_verification` | `FallbackBooleanChoiceField` | `bool` | yes |  | `None` | Whether identity verification is required at the time of user registration |
| `sms_sender` | `CharField` | `varchar(128)` | yes |  |  | alpha numeric identifier used as sender for SMS sent by this organization |
| `sms_message` | `FallbackTextField` | `text` | yes |  | `None` | SMS message template used for sending verification code. Must contain "{code}" placeholder for OTP value. |
| `sms_cooldown` | `FallbackPositiveIntegerField` | `integer unsigned` | yes |  | `None` | Time period a user will have to wait before requesting another SMS token (in seconds). |
| `sms_meta_data` | `JSONField` | `text` | yes |  |  | Additional configuration for SMS backend in JSON format (optional, leave blank if unsure) |
| `freeradius_allowed_hosts` | `FallbackTextField` | `text` | yes |  | `None` | Comma separated list of IP addresses allowed to access freeradius API |
| `coa_enabled` | `FallbackBooleanChoiceField` | `bool` | yes |  | `None` | Whether RADIUS Change Of Authoization (CoA) is enabled |
| `allowed_mobile_prefixes` | `FallbackTextField` | `text` | yes |  | `None` | Comma separated list of international mobile prefixes allowed to register via the user registration API. |
| `first_name` | `OrganizationRadiusSettingsFirstName` | `varchar(12)` | yes |  | `None` | Whether this field should be disabled, allowed or mandatory in the user registration API. |
| `last_name` | `OrganizationRadiusSettingsLastName` | `varchar(12)` | yes |  | `None` | Whether this field should be disabled, allowed or mandatory in the user registration API. |
| `location` | `OrganizationRadiusSettingsLocation` | `varchar(12)` | yes |  | `None` | Whether this field should be disabled, allowed or mandatory in the user registration API. |
| `birth_date` | `OrganizationRadiusSettingsBirthDate` | `varchar(12)` | yes |  | `None` | Whether this field should be disabled, allowed or mandatory in the user registration API. |
| `registration_enabled` | `FallbackBooleanChoiceField` | `bool` | yes |  | `None` | Whether the registration API endpoint should be enabled or not |
| `saml_registration_enabled` | `FallbackBooleanChoiceField` | `bool` | yes |  | `None` | Whether the registration using SAML should be enabled or not |
| `mac_addr_roaming_enabled` | `FallbackBooleanChoiceField` | `bool` | yes |  | `None` | Whether the MAC address roaming should be enabled or not. |
| `social_registration_enabled` | `FallbackBooleanChoiceField` | `bool` | yes |  | `None` | Whether the registration using social applications should be enabled or not |
| `login_url` | `URLField` | `varchar(200)` | yes |  |  | Enter the URL where users can log in to the wifi service |
| `status_url` | `URLField` | `varchar(200)` | yes |  |  | Enter the URL where users can log out from the wifi service |
| `password_reset_url` | `FallbackCharField` | `varchar(200)` | yes |  | `None` | Enter the URL where users can reset their password |

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade

## PhoneToken

Table `openwisp_radius_phonetoken`, app `openwisp_radius`.

Columns from `AbstractPhoneToken` (abstract): `user_id`, `valid_until`, `attempts`, `verified`, `token`, `ip`, `phone_number`.

Columns from `OrgMixin` (abstract): `organization_id`.

Columns from `TimeStampedEditableModel` (abstract): `created`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

![PhoneToken and the models it relates to](tables/PhoneToken.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `created` | `AutoCreatedField` | `datetime` |  |  | `now()` |  |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `valid_until` | `DateTimeField` | `datetime` |  |  | `get_sms_default_valid_until()` |  |
| `attempts` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `verified` | `BooleanField` | `bool` |  |  | `False` |  |
| `token` | `CharField` | `varchar(8)` |  |  | `generate_sms_token()` |  |
| `ip` | `GenericIPAddressField` | `char(39)` |  |  |  |  |
| `phone_number` | `PhoneNumberField` | `varchar(128)` |  |  |  |  |

Indexes:

- index `openwisp_radius_phonetoken_organization_id_5d939c4a` on `organization_id`
- index `openwisp_radius_phonetoken_user_id_90d96493` on `user_id`
- index `openwisp_ra_user_id_9fe207_idx` on `user_id`, `created`
- index `openwisp_ra_user_id_d4dd52_idx` on `user_id`, `created`, `ip`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## RegisteredUser

Table `openwisp_radius_registereduser`, app `openwisp_radius`.

Columns from `AbstractRegisteredUser` (abstract): `user_id`, `method`, `is_verified`, `modified`.

Columns from `UUIDModel` (abstract): `id`.

Columns from `OrgMixin` (abstract): `organization_id`.

![RegisteredUser and the models it relates to](tables/RegisteredUser.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `organization_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `id` | `UUIDField` | `char(32)` |  | PK | `uuid4()` |  |
| `user_id` | `ForeignKey` | `char(32)` |  | FK |  |  |
| `method` | `RegisteredUserMethod` | `varchar(64)` |  |  | `''` | users can sign up in different ways, some methods are valid as indirect identity verification (eg: mobile phone SIM card in most countries) |
| `is_verified` | `BooleanField` | `bool` |  |  | `False` | whether the user has completed any identity verification process sucessfully |
| `modified` | `AutoLastModifiedField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `openwisp_radius_registereduser_organization_id_b1772588` on `organization_id`
- index `openwisp_radius_registereduser_user_id_3af63f27` on `user_id`
- unique `unique_registered_user_per_org` on `user_id`, `organization_id`

Relations:

- `organization_id` → [Organization](#organization).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## UserLanguage

| Name | Stored as |
| --- | --- |
| `Afrikaans` | `af` |
| `Arabic` | `ar` |
| `Algerian Arabic` | `ar-dz` |
| `Asturian` | `ast` |
| `Azerbaijani` | `az` |
| `Bulgarian` | `bg` |
| `Belarusian` | `be` |
| `Bengali` | `bn` |
| `Breton` | `br` |
| `Bosnian` | `bs` |
| `Catalan` | `ca` |
| `Central Kurdish (Sorani)` | `ckb` |
| `Czech` | `cs` |
| `Welsh` | `cy` |
| `Danish` | `da` |
| `German` | `de` |
| `Lower Sorbian` | `dsb` |
| `Greek` | `el` |
| `English` | `en` |
| `Australian English` | `en-au` |
| `British English` | `en-gb` |
| `Esperanto` | `eo` |
| `Spanish` | `es` |
| `Argentinian Spanish` | `es-ar` |
| `Colombian Spanish` | `es-co` |
| `Mexican Spanish` | `es-mx` |
| `Nicaraguan Spanish` | `es-ni` |
| `Venezuelan Spanish` | `es-ve` |
| `Estonian` | `et` |
| `Basque` | `eu` |
| `Persian` | `fa` |
| `Finnish` | `fi` |
| `French` | `fr` |
| `Frisian` | `fy` |
| `Irish` | `ga` |
| `Scottish Gaelic` | `gd` |
| `Galician` | `gl` |
| `Hebrew` | `he` |
| `Hindi` | `hi` |
| `Croatian` | `hr` |
| `Upper Sorbian` | `hsb` |
| `Hungarian` | `hu` |
| `Armenian` | `hy` |
| `Interlingua` | `ia` |
| `Indonesian` | `id` |
| `Igbo` | `ig` |
| `Ido` | `io` |
| `Icelandic` | `is` |
| `Italian` | `it` |
| `Japanese` | `ja` |
| `Georgian` | `ka` |
| `Kabyle` | `kab` |
| `Kazakh` | `kk` |
| `Khmer` | `km` |
| `Kannada` | `kn` |
| `Korean` | `ko` |
| `Kyrgyz` | `ky` |
| `Luxembourgish` | `lb` |
| `Lithuanian` | `lt` |
| `Latvian` | `lv` |
| `Macedonian` | `mk` |
| `Malayalam` | `ml` |
| `Mongolian` | `mn` |
| `Marathi` | `mr` |
| `Malay` | `ms` |
| `Burmese` | `my` |
| `Norwegian Bokmål` | `nb` |
| `Nepali` | `ne` |
| `Dutch` | `nl` |
| `Norwegian Nynorsk` | `nn` |
| `Ossetic` | `os` |
| `Punjabi` | `pa` |
| `Polish` | `pl` |
| `Portuguese` | `pt` |
| `Brazilian Portuguese` | `pt-br` |
| `Romanian` | `ro` |
| `Russian` | `ru` |
| `Slovak` | `sk` |
| `Slovenian` | `sl` |
| `Albanian` | `sq` |
| `Serbian` | `sr` |
| `Serbian Latin` | `sr-latn` |
| `Swedish` | `sv` |
| `Swahili` | `sw` |
| `Tamil` | `ta` |
| `Telugu` | `te` |
| `Tajik` | `tg` |
| `Thai` | `th` |
| `Turkmen` | `tk` |
| `Turkish` | `tr` |
| `Tatar` | `tt` |
| `Udmurt` | `udm` |
| `Uyghur` | `ug` |
| `Ukrainian` | `uk` |
| `Urdu` | `ur` |
| `Uzbek` | `uz` |
| `Vietnamese` | `vi` |
| `Simplified Chinese` | `zh-hans` |
| `Traditional Chinese` | `zh-hant` |

## RadiusCheckOp

| Name | Stored as |
| --- | --- |
| `=` | `=` |
| `:=` | `:=` |
| `==` | `==` |
| `+=` | `+=` |
| `!=` | `!=` |
| `>` | `>` |
| `>=` | `>=` |
| `<` | `<` |
| `<=` | `<=` |
| `=~` | `=~` |
| `!~` | `!~` |
| `=*` | `=*` |
| `!*` | `!*` |

## RadiusReplyOp

| Name | Stored as |
| --- | --- |
| `=` | `=` |
| `:=` | `:=` |
| `+=` | `+=` |

## RadiusGroupCheckOp

| Name | Stored as |
| --- | --- |
| `=` | `=` |
| `:=` | `:=` |
| `==` | `==` |
| `+=` | `+=` |
| `!=` | `!=` |
| `>` | `>` |
| `>=` | `>=` |
| `<` | `<` |
| `<=` | `<=` |
| `=~` | `=~` |
| `!~` | `!~` |
| `=*` | `=*` |
| `!*` | `!*` |

## RadiusGroupReplyOp

| Name | Stored as |
| --- | --- |
| `=` | `=` |
| `:=` | `:=` |
| `+=` | `+=` |

## NasType

| Name | Stored as |
| --- | --- |
| `Async` | `Async` |
| `Sync` | `Sync` |
| `ISDN Sync` | `ISDN Sync` |
| `ISDN Async V.120` | `ISDN Async V.120` |
| `ISDN Async V.110` | `ISDN Async V.110` |
| `Virtual` | `Virtual` |
| `PIAFS` | `PIAFS` |
| `HDLC Clear` | `HDLC Clear` |
| `Channel` | `Channel` |
| `X.25` | `X.25` |
| `X.75` | `X.75` |
| `G.3 Fax` | `G.3 Fax` |
| `SDSL - Symmetric DSL` | `SDSL` |
| `ADSL-CAP` | `ADSL-CAP` |
| `ADSL-DMT` | `ADSL-DMT` |
| `IDSL` | `IDSL` |
| `Ethernet` | `Ethernet` |
| `xDSL` | `xDSL` |
| `Cable` | `Cable` |
| `Wireless - Other` | `Wireless - Other` |
| `Wireless - IEEE 802.11` | `IEEE 802.11` |
| `Token-Ring` | `Token-Ring` |
| `FDDI` | `FDDI` |
| `Wireless - CDMA2000` | `Wireless - CDMA2000` |
| `Wireless - UMTS` | `Wireless - UMTS` |
| `Wireless - 1X-EV` | `Wireless - 1X-EV` |
| `IAPP` | `IAPP` |
| `FTTP` | `FTTP` |
| `Wireless - IEEE 802.16` | `IEEE 802.16` |
| `Wireless - IEEE 802.20` | `IEEE 802.20` |
| `Wireless - IEEE 802.22` | `IEEE 802.22` |
| `PPPoA - PPP over ATM` | `PPPoA` |
| `PPPoEoA - PPP over Ethernet over ATM` | `PPPoEoA` |
| `PPPoEoE - PPP over Ethernet over Ethernet` | `PPPoEoE` |
| `PPPoEoVLAN - PPP over Ethernet over VLAN` | `PPPoEoVLAN` |
| `PPPoEoQinQ - PPP over Ethernet over IEEE 802.1QinQ` | `PPPoEoQinQ` |
| `xPON - Passive Optical Network` | `xPON` |
| `Wireless - XGP` | `Wireless - XGP` |
| ` WiMAX Pre-Release 8 IWK Function` | `WiMAX` |
| `WIMAX-WIFI-IWK: WiMAX WIFI Interworking` | `WIMAX-WIFI-IWK` |
| `WIMAX-SFF: Signaling Forwarding Function for LTE/3GPP2` | `WIMAX-SFF` |
| `WIMAX-HA-LMA: WiMAX HA and or LMA function` | `WIMAX-HA-LMA` |
| `WIMAX-DHCP: WIMAX DCHP service` | `WIMAX-DHCP` |
| `WIMAX-LBS: WiMAX location based service` | `WIMAX-LBS` |
| `WIMAX-WVS: WiMAX voice service` | `WIMAX-WVS` |
| `Other` | `Other` |

## RadiusBatchStrategy

| Name | Stored as |
| --- | --- |
| `Generate from prefix` | `prefix` |
| `Import from CSV` | `csv` |

## RadiusBatchStatus

| Name | Stored as |
| --- | --- |
| `Pending` | `pending` |
| `Processing` | `processing` |
| `Completed` | `completed` |
| `Failed` | `failed` |

## OrganizationRadiusSettingsFirstName

| Name | Stored as |
| --- | --- |
| `Disabled` | `disabled` |
| `Allowed` | `allowed` |
| `Mandatory` | `mandatory` |

## OrganizationRadiusSettingsLastName

| Name | Stored as |
| --- | --- |
| `Disabled` | `disabled` |
| `Allowed` | `allowed` |
| `Mandatory` | `mandatory` |

## OrganizationRadiusSettingsLocation

| Name | Stored as |
| --- | --- |
| `Disabled` | `disabled` |
| `Allowed` | `allowed` |
| `Mandatory` | `mandatory` |

## OrganizationRadiusSettingsBirthDate

| Name | Stored as |
| --- | --- |
| `Disabled` | `disabled` |
| `Allowed` | `allowed` |
| `Mandatory` | `mandatory` |

## RegisteredUserMethod

| Name | Stored as |
| --- | --- |
| `Unspecified` |  |
| `Manually created` | `manual` |
| `Email` | `email` |
| `Mobile phone` | `mobile_phone` |
| `Pending Verification` | `pending_verification` |
| `Social login` | `social_login` |
