# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Site](#site) · [Gallery](#gallery) · [Photo](#photo) · [PhotoEffect](#photoeffect) · [Watermark](#watermark) · [PhotoSize](#photosize)

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

CREATE TABLE "photologue_gallery" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "date_added" datetime NOT NULL, "title" varchar(250) NOT NULL UNIQUE, "slug" varchar(250) NOT NULL UNIQUE, "description" text NOT NULL, "is_public" bool NOT NULL);

CREATE TABLE "photologue_gallery_photos" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "sort_value" integer NOT NULL, "gallery_id" integer NOT NULL REFERENCES "photologue_gallery" ("id") DEFERRABLE INITIALLY DEFERRED, "photo_id" integer NOT NULL REFERENCES "photologue_photo" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "photologue_gallery_sites" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "gallery_id" integer NOT NULL REFERENCES "photologue_gallery" ("id") DEFERRABLE INITIALLY DEFERRED, "site_id" integer NOT NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "photologue_photo" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "image" varchar(100) NOT NULL, "date_taken" datetime NULL, "view_count" integer unsigned NOT NULL CHECK ("view_count" >= 0), "crop_from" varchar(10) NOT NULL, "effect_id" integer NULL REFERENCES "photologue_photoeffect" ("id") DEFERRABLE INITIALLY DEFERRED, "title" varchar(250) NOT NULL UNIQUE, "slug" varchar(250) NOT NULL UNIQUE, "caption" text NOT NULL, "date_added" datetime NOT NULL, "is_public" bool NOT NULL);

CREATE TABLE "photologue_photo_sites" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "photo_id" integer NOT NULL REFERENCES "photologue_photo" ("id") DEFERRABLE INITIALLY DEFERRED, "site_id" integer NOT NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "photologue_photoeffect" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(30) NOT NULL UNIQUE, "description" text NOT NULL, "transpose_method" varchar(15) NOT NULL, "color" real NOT NULL, "brightness" real NOT NULL, "contrast" real NOT NULL, "sharpness" real NOT NULL, "filters" varchar(200) NOT NULL, "reflection_size" real NOT NULL, "reflection_strength" real NOT NULL, "background_color" varchar(7) NOT NULL);

CREATE TABLE "photologue_watermark" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(30) NOT NULL UNIQUE, "description" text NOT NULL, "image" varchar(100) NOT NULL, "style" varchar(5) NOT NULL, "opacity" real NOT NULL);

CREATE TABLE "photologue_photosize" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(40) NOT NULL UNIQUE, "width" integer unsigned NOT NULL CHECK ("width" >= 0), "height" integer unsigned NOT NULL CHECK ("height" >= 0), "quality" integer unsigned NOT NULL CHECK ("quality" >= 0), "upscale" bool NOT NULL, "crop" bool NOT NULL, "pre_cache" bool NOT NULL, "increment_count" bool NOT NULL, "effect_id" integer NULL REFERENCES "photologue_photoeffect" ("id") DEFERRABLE INITIALLY DEFERRED, "watermark_id" integer NULL REFERENCES "photologue_watermark" ("id") DEFERRABLE INITIALLY DEFERRED);

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

CREATE UNIQUE INDEX "photologue_gallery_photos_gallery_id_photo_id_ac3d7969_uniq" ON "photologue_gallery_photos" ("gallery_id", "photo_id");

CREATE INDEX "photologue_gallery_photos_gallery_id_770eac69" ON "photologue_gallery_photos" ("gallery_id");

CREATE INDEX "photologue_gallery_photos_photo_id_38de5ebc" ON "photologue_gallery_photos" ("photo_id");

CREATE UNIQUE INDEX "photologue_gallery_sites_gallery_id_site_id_4158504d_uniq" ON "photologue_gallery_sites" ("gallery_id", "site_id");

CREATE INDEX "photologue_gallery_sites_gallery_id_263bc875" ON "photologue_gallery_sites" ("gallery_id");

CREATE INDEX "photologue_gallery_sites_site_id_ff8bed9d" ON "photologue_gallery_sites" ("site_id");

CREATE INDEX "photologue_photo_effect_id_5f27aba5" ON "photologue_photo" ("effect_id");

CREATE UNIQUE INDEX "photologue_photo_sites_photo_id_site_id_a8528863_uniq" ON "photologue_photo_sites" ("photo_id", "site_id");

CREATE INDEX "photologue_photo_sites_photo_id_1e4e4a98" ON "photologue_photo_sites" ("photo_id");

CREATE INDEX "photologue_photo_sites_site_id_a26b734c" ON "photologue_photo_sites" ("site_id");

CREATE INDEX "photologue_photosize_effect_id_1d9b1b30" ON "photologue_photosize" ("effect_id");

CREATE INDEX "photologue_photosize_watermark_id_f9419f99" ON "photologue_photosize" ("watermark_id");
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

## Site

Table `django_site`, app `sites`.

![Site and the models it relates to](tables/Site.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `domain` | `CharField` | `varchar(100)` |  | unique |  |  |
| `name` | `CharField` | `varchar(50)` |  |  |  |  |

Relations:

- `gallery_set` — many to many with [Gallery](#gallery)
- `photo_set` — many to many with [Photo](#photo)

## Gallery

Table `photologue_gallery`, app `photologue`.

![Gallery and the models it relates to](tables/Gallery.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `date_added` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `title` | `CharField` | `varchar(250)` |  | unique |  |  |
| `slug` | `SlugField` | `varchar(250)` |  | unique |  | A "slug" is a unique URL-friendly title for an object. |
| `description` | `TextField` | `text` |  |  |  |  |
| `is_public` | `BooleanField` | `bool` |  |  | `True` | Public galleries will be displayed in the default views. |

Relations:

- `photos` — many to many with [Photo](#photo)
- `sites` — many to many with [Site](#site)

## Photo

Table `photologue_photo`, app `photologue`.

Columns from `ImageModel` (abstract): `image`, `date_taken`, `view_count`, `crop_from`, `effect_id`.

![Photo and the models it relates to](tables/Photo.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `image` | `ImageField` | `varchar(100)` |  |  |  |  |
| `date_taken` | `DateTimeField` | `datetime` | yes |  |  | Date image was taken; is obtained from the image EXIF data. |
| `view_count` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `crop_from` | `PhotoCropFrom` | `varchar(10)` |  |  | `'center'` |  |
| `effect_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `title` | `CharField` | `varchar(250)` |  | unique |  |  |
| `slug` | `SlugField` | `varchar(250)` |  | unique |  | A "slug" is a unique URL-friendly title for an object. |
| `caption` | `TextField` | `text` |  |  |  |  |
| `date_added` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `is_public` | `BooleanField` | `bool` |  |  | `True` | Public photographs will be displayed in the default views. |

Indexes:

- index `photologue_photo_effect_id_5f27aba5` on `effect_id`

Relations:

- `galleries` — many to many with [Gallery](#gallery)
- `effect_id` → [PhotoEffect](#photoeffect).`id`, on delete cascade
- `sites` — many to many with [Site](#site)

## PhotoEffect

Table `photologue_photoeffect`, app `photologue`.

Columns from `BaseEffect` (abstract): `name`, `description`.

A pre-defined effect to apply to photos

![PhotoEffect and the models it relates to](tables/PhotoEffect.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(30)` |  | unique |  |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `transpose_method` | `PhotoEffectTransposeMethod` | `varchar(15)` |  |  |  |  |
| `color` | `FloatField` | `real` |  |  | `1.0` | A factor of 0.0 gives a black and white image, a factor of 1.0 gives the original image. |
| `brightness` | `FloatField` | `real` |  |  | `1.0` | A factor of 0.0 gives a black image, a factor of 1.0 gives the original image. |
| `contrast` | `FloatField` | `real` |  |  | `1.0` | A factor of 0.0 gives a solid grey image, a factor of 1.0 gives the original image. |
| `sharpness` | `FloatField` | `real` |  |  | `1.0` | A factor of 0.0 gives a blurred image, a factor of 1.0 gives the original image. |
| `filters` | `CharField` | `varchar(200)` |  |  |  | Chain multiple filters using the following pattern "FILTER_ONE->FILTER_TWO->FILTER_THREE". Image filters will be applied in order. The following filters are available: BLUR, CONTOUR, DETAIL, EDGE_ENHANCE, EDGE_ENHANCE_MORE, EMBOSS, FIND_EDGES, Kernel, SHARPEN, SMOOTH, SMOOTH_MORE. |
| `reflection_size` | `FloatField` | `real` |  |  | `0` | The height of the reflection as a percentage of the orignal image. A factor of 0.0 adds no reflection, a factor of 1.0 adds a reflection equal to the height of the orignal image. |
| `reflection_strength` | `FloatField` | `real` |  |  | `0.6` | The initial opacity of the reflection gradient. |
| `background_color` | `CharField` | `varchar(7)` |  |  | `'#FFFFFF'` | The background color of the reflection gradient. Set this to match the background color of your page. |

Relations:

- [Photo](#photo).`effect_id` → `id`, on delete cascade
- [PhotoSize](#photosize).`effect_id` → `id`, on delete cascade

## Watermark

Table `photologue_watermark`, app `photologue`.

Columns from `BaseEffect` (abstract): `name`, `description`.

![Watermark and the models it relates to](tables/Watermark.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(30)` |  | unique |  |  |
| `description` | `TextField` | `text` |  |  |  |  |
| `image` | `ImageField` | `varchar(100)` |  |  |  |  |
| `style` | `WatermarkStyle` | `varchar(5)` |  |  | `'scale'` |  |
| `opacity` | `FloatField` | `real` |  |  | `1` | The opacity of the overlay. |

Relations:

- [PhotoSize](#photosize).`watermark_id` → `id`, on delete cascade

## PhotoSize

Table `photologue_photosize`, app `photologue`.

About the Photosize name: it's used to create get_PHOTOSIZE_url() methods,
so the name has to follow the same restrictions as any Python method name,
e.g. no spaces or non-ascii characters.

![PhotoSize and the models it relates to](tables/PhotoSize.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(40)` |  | unique |  | Photo size name should contain only letters, numbers and underscores. Examples: "thumbnail", "display", "small", "main_page_widget". |
| `width` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | If width is set to "0" the image will be scaled to the supplied height. |
| `height` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` | If height is set to "0" the image will be scaled to the supplied width |
| `quality` | `PhotoSizeQuality` | `integer unsigned` |  |  | `70` | JPEG image quality. |
| `upscale` | `BooleanField` | `bool` |  |  | `False` | If selected the image will be scaled up if necessary to fit the supplied dimensions. Cropped sizes will be upscaled regardless of this setting. |
| `crop` | `BooleanField` | `bool` |  |  | `False` | If selected the image will be scaled and cropped to fit the supplied dimensions. |
| `pre_cache` | `BooleanField` | `bool` |  |  | `False` | If selected this photo size will be pre-cached as photos are added. |
| `increment_count` | `BooleanField` | `bool` |  |  | `False` | If selected the image's "view_count" will be incremented when this photo size is displayed. |
| `effect_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `watermark_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `photologue_photosize_effect_id_1d9b1b30` on `effect_id`
- index `photologue_photosize_watermark_id_f9419f99` on `watermark_id`

Relations:

- `effect_id` → [PhotoEffect](#photoeffect).`id`, on delete cascade
- `watermark_id` → [Watermark](#watermark).`id`, on delete cascade

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## PhotoCropFrom

| Name | Stored as |
| --- | --- |
| `Top` | `top` |
| `Right` | `right` |
| `Bottom` | `bottom` |
| `Left` | `left` |
| `Center (Default)` | `center` |

## PhotoEffectTransposeMethod

| Name | Stored as |
| --- | --- |
| `Flip left to right` | `FLIP_LEFT_RIGHT` |
| `Flip top to bottom` | `FLIP_TOP_BOTTOM` |
| `Rotate 90 degrees counter-clockwise` | `ROTATE_90` |
| `Rotate 90 degrees clockwise` | `ROTATE_270` |
| `Rotate 180 degrees` | `ROTATE_180` |

## WatermarkStyle

| Name | Stored as |
| --- | --- |
| `Tile` | `tile` |
| `Scale` | `scale` |

## PhotoSizeQuality

| Name | Stored as |
| --- | --- |
| `Very Low` | `30` |
| `Low` | `40` |
| `Medium-Low` | `50` |
| `Medium` | `60` |
| `Medium-High` | `70` |
| `High` | `80` |
| `Very High` | `90` |
