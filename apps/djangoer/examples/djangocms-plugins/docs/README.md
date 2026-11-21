# Schema docs

![ER diagram](er.svg)

[LogEntry](#logentry) · [Permission](#permission) · [Group](#group) · [User](#user) · [ContentType](#contenttype) · [Session](#session) · [Site](#site) · [UserSettings](#usersettings) · [Page](#page) · [PageUrl](#pageurl) · [GlobalPagePermission](#globalpagepermission) · [PagePermission](#pagepermission) · [PageUser](#pageuser) · [PageUserGroup](#pageusergroup) · [Placeholder](#placeholder) · [CMSPlugin](#cmsplugin) · [PageContent](#pagecontent) · [PlaceholderReference](#placeholderreference) · [AliasPluginModel](#aliaspluginmodel) · [UrlconfRevision](#urlconfrevision) · [CacheKey](#cachekey) · [Source](#source) · [Thumbnail](#thumbnail) · [ThumbnailDimensions](#thumbnaildimensions) · [filer.Folder](#filerfolder) · [FolderPermission](#folderpermission) · [filer.File](#filerfile) · [Clipboard](#clipboard) · [ClipboardItem](#clipboarditem) · [Image](#image) · [ThumbnailOption](#thumbnailoption) · [Text](#text) · [Version](#version) · [StateTracking](#statetracking) · [CategoryTranslation](#categorytranslation) · [Category](#category) · [Alias](#alias) · [AliasContent](#aliascontent) · [AliasPlugin](#aliasplugin) · [SnippetGrouper](#snippetgrouper) · [Snippet](#snippet) · [SnippetPtr](#snippetptr) · [Picture](#picture) · [djangocms_file.File](#djangocms_filefile) · [djangocms_file.Folder](#djangocms_filefolder) · [VideoPlayer](#videoplayer) · [VideoSource](#videosource) · [HlsStreamSource](#hlsstreamsource) · [VideoTrack](#videotrack)

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

CREATE TABLE "cms_usersettings" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL UNIQUE REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "language" varchar(10) NOT NULL, "clipboard_id" integer NULL REFERENCES "cms_placeholder" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "cms_page" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "path" varchar(255) NOT NULL UNIQUE, "depth" integer unsigned NOT NULL CHECK ("depth" >= 0), "numchild" integer unsigned NOT NULL CHECK ("numchild" >= 0), "parent_id" integer NULL REFERENCES "cms_page" ("id") DEFERRABLE INITIALLY DEFERRED, "site_id" integer NOT NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED, "created_by" varchar(255) NOT NULL, "changed_by" varchar(255) NOT NULL, "creation_date" datetime NOT NULL, "changed_date" datetime NOT NULL, "reverse_id" varchar(40) NULL, "navigation_extenders" varchar(80) NULL, "login_required" bool NOT NULL, "is_home" bool NOT NULL, "application_urls" varchar(200) NULL, "application_namespace" varchar(200) NULL, "is_page_type" bool NOT NULL);

CREATE TABLE "cms_pageurl" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "slug" varchar(255) NOT NULL, "path" varchar(255) NULL, "language" varchar(15) NOT NULL, "page_id" integer NOT NULL REFERENCES "cms_page" ("id") DEFERRABLE INITIALLY DEFERRED, "site_id" integer NOT NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED, "managed" bool NOT NULL, CONSTRAINT "unique_together_page_language" UNIQUE ("page_id", "language"), CONSTRAINT "unique_site_language_path" UNIQUE ("site_id", "language", "path"));

CREATE TABLE "cms_globalpagepermission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED, "can_change" bool NOT NULL, "can_add" bool NOT NULL, "can_delete" bool NOT NULL, "can_publish" bool NOT NULL, "can_change_advanced_settings" bool NOT NULL, "can_change_permissions" bool NOT NULL, "can_move_page" bool NOT NULL, "can_view" bool NOT NULL, "can_recover_page" bool NOT NULL);

CREATE TABLE "cms_globalpagepermission_sites" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "globalpagepermission_id" integer NOT NULL REFERENCES "cms_globalpagepermission" ("id") DEFERRABLE INITIALLY DEFERRED, "site_id" integer NOT NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "cms_pagepermission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED, "can_change" bool NOT NULL, "can_add" bool NOT NULL, "can_delete" bool NOT NULL, "can_publish" bool NOT NULL, "can_change_advanced_settings" bool NOT NULL, "can_change_permissions" bool NOT NULL, "can_move_page" bool NOT NULL, "can_view" bool NOT NULL, "grant_on" integer NOT NULL, "page_id" integer NULL REFERENCES "cms_page" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "cms_pageuser" ("user_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "created_by_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "cms_pageusergroup" ("group_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED, "created_by_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "cms_placeholder" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "slot" varchar(255) NOT NULL, "default_width" smallint unsigned NULL CHECK ("default_width" >= 0), "content_type_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "object_id" integer unsigned NULL CHECK ("object_id" >= 0));

CREATE TABLE "cms_cmsplugin" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "placeholder_id" integer NULL REFERENCES "cms_placeholder" ("id") DEFERRABLE INITIALLY DEFERRED, "parent_id" integer NULL REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "position" smallint NOT NULL, "language" varchar(15) NOT NULL, "plugin_type" varchar(50) NOT NULL, "creation_date" datetime NOT NULL, "changed_date" datetime NOT NULL);

CREATE TABLE "cms_pagecontent" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "language" varchar(15) NOT NULL, "title" varchar(255) NOT NULL, "page_title" varchar(255) NULL, "menu_title" varchar(255) NULL, "meta_description" text NULL, "redirect" varchar(2048) NULL, "slug" varchar(255) NOT NULL, "overwrite_url" varchar(255) NULL, "page_id" integer NOT NULL REFERENCES "cms_page" ("id") DEFERRABLE INITIALLY DEFERRED, "creation_date" datetime NOT NULL, "created_by" varchar(255) NOT NULL, "changed_by" varchar(255) NOT NULL, "changed_date" datetime NOT NULL, "in_navigation" bool NOT NULL, "soft_root" bool NOT NULL, "template" varchar(100) NOT NULL, "limit_visibility_in_menu" smallint NULL, "xframe_options" integer NOT NULL);

CREATE TABLE "cms_placeholderreference" ("cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(255) NOT NULL);

CREATE TABLE "cms_aliaspluginmodel" ("cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "plugin_id" integer NULL REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "alias_placeholder_id" integer NULL REFERENCES "cms_placeholder" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "cms_urlconfrevision" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "revision" varchar(255) NOT NULL);

CREATE TABLE "menus_cachekey" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "language" varchar(255) NOT NULL, "site" integer unsigned NOT NULL CHECK ("site" >= 0), "key" varchar(255) NOT NULL);

CREATE TABLE "easy_thumbnails_source" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "storage_hash" varchar(40) NOT NULL, "name" varchar(255) NOT NULL, "modified" datetime NOT NULL);

CREATE TABLE "easy_thumbnails_thumbnail" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "storage_hash" varchar(40) NOT NULL, "name" varchar(255) NOT NULL, "modified" datetime NOT NULL, "source_id" integer NOT NULL REFERENCES "easy_thumbnails_source" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "easy_thumbnails_thumbnaildimensions" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "thumbnail_id" integer NOT NULL UNIQUE REFERENCES "easy_thumbnails_thumbnail" ("id") DEFERRABLE INITIALLY DEFERRED, "width" integer unsigned NULL CHECK ("width" >= 0), "height" integer unsigned NULL CHECK ("height" >= 0));

CREATE TABLE "filer_folder" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "parent_id" integer NULL REFERENCES "filer_folder" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(255) NOT NULL, "owner_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "uploaded_at" datetime NOT NULL, "created_at" datetime NOT NULL, "modified_at" datetime NOT NULL);

CREATE TABLE "filer_folderpermission" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "folder_id" integer NULL REFERENCES "filer_folder" ("id") DEFERRABLE INITIALLY DEFERRED, "type" smallint NOT NULL, "user_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "group_id" integer NULL REFERENCES "auth_group" ("id") DEFERRABLE INITIALLY DEFERRED, "everybody" bool NOT NULL, "can_read" smallint NULL, "can_edit" smallint NULL, "can_add_children" smallint NULL);

CREATE TABLE "filer_file" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "polymorphic_ctype_id" integer NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "folder_id" integer NULL REFERENCES "filer_folder" ("id") DEFERRABLE INITIALLY DEFERRED, "file" varchar(255) NULL, "_file_size" bigint NULL, "sha1" varchar(40) NOT NULL, "has_all_mandatory_data" bool NOT NULL, "original_filename" varchar(255) NULL, "name" varchar(255) NOT NULL, "description" text NULL, "owner_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "uploaded_at" datetime NOT NULL, "modified_at" datetime NOT NULL, "is_public" bool NOT NULL, "mime_type" varchar(255) NOT NULL);

CREATE TABLE "filer_clipboard" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "filer_clipboarditem" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "file_id" integer NOT NULL REFERENCES "filer_file" ("id") DEFERRABLE INITIALLY DEFERRED, "clipboard_id" integer NOT NULL REFERENCES "filer_clipboard" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "filer_image" ("_height" real NULL, "_width" real NULL, "_transparent" bool NOT NULL, "default_alt_text" varchar(255) NULL, "default_caption" varchar(255) NULL, "subject_location" varchar(64) NOT NULL, "file_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "filer_file" ("id") DEFERRABLE INITIALLY DEFERRED, "date_taken" datetime NULL, "author" varchar(255) NULL, "must_always_publish_author_credit" bool NOT NULL, "must_always_publish_copyright" bool NOT NULL);

CREATE TABLE "filer_thumbnailoption" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(100) NOT NULL, "width" integer NOT NULL, "height" integer NOT NULL, "crop" bool NOT NULL, "upscale" bool NOT NULL);

CREATE TABLE "djangocms_text_text" ("cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "body" text NOT NULL, "json" text NULL CHECK ((JSON_VALID("json") OR "json" IS NULL)), "rte" varchar(16) NOT NULL);

CREATE TABLE "djangocms_versioning_version" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "created" datetime NOT NULL, "modified" datetime NOT NULL, "created_by_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "number" varchar(11) NOT NULL, "content_type_id" integer NOT NULL REFERENCES "django_content_type" ("id") DEFERRABLE INITIALLY DEFERRED, "object_id" integer unsigned NOT NULL CHECK ("object_id" >= 0), "state" varchar(50) NOT NULL, "locked_by_id" integer NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED, "source_id" integer NULL REFERENCES "djangocms_versioning_version" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_versioning_statetracking" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "version_id" integer NOT NULL REFERENCES "djangocms_versioning_version" ("id") DEFERRABLE INITIALLY DEFERRED, "date" datetime NOT NULL, "old_state" varchar(100) NOT NULL, "new_state" varchar(100) NOT NULL, "user_id" integer NOT NULL REFERENCES "auth_user" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_alias_category_translation" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "language_code" varchar(15) NOT NULL, "name" varchar(120) NOT NULL, "master_id" bigint NULL REFERENCES "djangocms_alias_category" ("id") DEFERRABLE INITIALLY DEFERRED, CONSTRAINT "djangocms_alias_category_translation_uniq_lang" UNIQUE ("language_code", "master_id"));

CREATE TABLE "djangocms_alias_category" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT);

CREATE TABLE "djangocms_alias_alias" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "creation_method" varchar(20) NOT NULL, "category_id" bigint NOT NULL REFERENCES "djangocms_alias_category" ("id") DEFERRABLE INITIALLY DEFERRED, "position" integer unsigned NOT NULL CHECK ("position" >= 0), "static_code" varchar(255) NULL, "site_id" integer NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_alias_aliascontent" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "alias_id" bigint NOT NULL REFERENCES "djangocms_alias_alias" ("id") DEFERRABLE INITIALLY DEFERRED, "name" varchar(120) NOT NULL, "language" varchar(10) NOT NULL);

CREATE TABLE "djangocms_alias_aliasplugin" ("cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "alias_id" bigint NOT NULL REFERENCES "djangocms_alias_alias" ("id") DEFERRABLE INITIALLY DEFERRED, "template" varchar(255) NOT NULL);

CREATE TABLE "djangocms_snippet_snippetgrouper" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT);

CREATE TABLE "djangocms_snippet_snippet" ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "name" varchar(255) NOT NULL, "snippet_grouper_id" bigint NOT NULL REFERENCES "djangocms_snippet_snippetgrouper" ("id") DEFERRABLE INITIALLY DEFERRED, "html" text NOT NULL, "template" varchar(255) NOT NULL, "slug" varchar(255) NOT NULL, "site_id" integer NULL REFERENCES "django_site" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_snippet_snippetptr" ("cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "snippet_grouper_id" bigint NOT NULL REFERENCES "djangocms_snippet_snippetgrouper" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_picture_picture" ("template" varchar(255) NOT NULL, "picture_id" integer NULL REFERENCES "filer_image" ("file_ptr_id") DEFERRABLE INITIALLY DEFERRED, "external_picture" varchar(255) NULL, "width" integer unsigned NULL CHECK ("width" >= 0), "height" integer unsigned NULL CHECK ("height" >= 0), "alignment" varchar(255) NOT NULL, "caption_text" text NULL, "attributes" text NOT NULL, "link_url" varchar(2040) NULL, "link_page_id" integer NULL REFERENCES "cms_page" ("id") DEFERRABLE INITIALLY DEFERRED, "link_target" varchar(255) NOT NULL, "link_attributes" text NOT NULL, "use_automatic_scaling" bool NOT NULL, "use_no_cropping" bool NOT NULL, "use_crop" bool NOT NULL, "use_upscale" bool NOT NULL, "use_responsive_image" varchar(7) NOT NULL, "thumbnail_options_id" integer NULL REFERENCES "filer_thumbnailoption" ("id") DEFERRABLE INITIALLY DEFERRED, "cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_file_file" ("template" varchar(255) NOT NULL, "file_src_id" integer NULL REFERENCES "filer_file" ("id") DEFERRABLE INITIALLY DEFERRED, "file_name" varchar(255) NOT NULL, "link_target" varchar(255) NOT NULL, "link_title" varchar(255) NOT NULL, "show_file_size" bool NOT NULL, "attributes" text NOT NULL, "cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_file_folder" ("template" varchar(255) NOT NULL, "folder_src_id" integer NULL REFERENCES "filer_folder" ("id") DEFERRABLE INITIALLY DEFERRED, "link_target" varchar(255) NOT NULL, "show_file_size" bool NOT NULL, "attributes" text NOT NULL, "cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_video_videoplayer" ("template" varchar(255) NOT NULL, "label" varchar(255) NOT NULL, "embed_link" text NOT NULL, "parameters" text NOT NULL, "poster_id" integer NULL REFERENCES "filer_image" ("file_ptr_id") DEFERRABLE INITIALLY DEFERRED, "attributes" text NOT NULL, "show_controls" bool NOT NULL, "autoplay" bool NOT NULL, "cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED);

CREATE TABLE "djangocms_video_videosource" ("cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "source_file_id" integer NULL REFERENCES "filer_file" ("id") DEFERRABLE INITIALLY DEFERRED, "text_title" varchar(255) NOT NULL, "text_description" text NOT NULL, "attributes" text NOT NULL);

CREATE TABLE "djangocms_video_hlsstreamsource" ("cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "hls_source_url" varchar(200) NOT NULL);

CREATE TABLE "djangocms_video_videotrack" ("cmsplugin_ptr_id" integer NOT NULL PRIMARY KEY REFERENCES "cms_cmsplugin" ("id") DEFERRABLE INITIALLY DEFERRED, "kind" varchar(255) NOT NULL, "src_id" integer NULL REFERENCES "filer_file" ("id") DEFERRABLE INITIALLY DEFERRED, "srclang" varchar(255) NOT NULL, "label" varchar(255) NOT NULL, "attributes" text NOT NULL);

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

CREATE INDEX "cms_usersettings_clipboard_id_3ae17c19" ON "cms_usersettings" ("clipboard_id");

CREATE INDEX "cms_page_parent_id_f89b72e4" ON "cms_page" ("parent_id");

CREATE INDEX "cms_page_site_id_4323d3ff" ON "cms_page" ("site_id");

CREATE INDEX "cms_page_reverse_id_ffc9ede2" ON "cms_page" ("reverse_id");

CREATE INDEX "cms_page_navigation_extenders_c24af8dd" ON "cms_page" ("navigation_extenders");

CREATE INDEX "cms_page_is_home_edadca07" ON "cms_page" ("is_home");

CREATE INDEX "cms_page_application_urls_9ef47497" ON "cms_page" ("application_urls");

CREATE INDEX "cms_pageurl_slug_f8d42287" ON "cms_pageurl" ("slug");

CREATE INDEX "cms_pageurl_path_9f26e673" ON "cms_pageurl" ("path");

CREATE INDEX "cms_pageurl_language_a10302dc" ON "cms_pageurl" ("language");

CREATE INDEX "cms_pageurl_page_id_dbf4793a" ON "cms_pageurl" ("page_id");

CREATE INDEX "cms_pageurl_site_id_7d6fd976" ON "cms_pageurl" ("site_id");

CREATE INDEX "cms_globalpagepermission_user_id_a227cee1" ON "cms_globalpagepermission" ("user_id");

CREATE INDEX "cms_globalpagepermission_group_id_991b4733" ON "cms_globalpagepermission" ("group_id");

CREATE UNIQUE INDEX "cms_globalpagepermission_sites_globalpagepermission_id_site_id_db684f41_uniq" ON "cms_globalpagepermission_sites" ("globalpagepermission_id", "site_id");

CREATE INDEX "cms_globalpagepermission_sites_globalpagepermission_id_46bd2681" ON "cms_globalpagepermission_sites" ("globalpagepermission_id");

CREATE INDEX "cms_globalpagepermission_sites_site_id_00460b53" ON "cms_globalpagepermission_sites" ("site_id");

CREATE INDEX "cms_pagepermission_user_id_0c7d2b3c" ON "cms_pagepermission" ("user_id");

CREATE INDEX "cms_pagepermission_group_id_af5af193" ON "cms_pagepermission" ("group_id");

CREATE INDEX "cms_pagepermission_page_id_0ae9a163" ON "cms_pagepermission" ("page_id");

CREATE INDEX "cms_pageuser_created_by_id_8e9fbf83" ON "cms_pageuser" ("created_by_id");

CREATE INDEX "cms_pageusergroup_created_by_id_7d57fa39" ON "cms_pageusergroup" ("created_by_id");

CREATE INDEX "cms_placeholder_slot_0bc04d21" ON "cms_placeholder" ("slot");

CREATE INDEX "cms_placeholder_content_type_id_a7659d9c" ON "cms_placeholder" ("content_type_id");

CREATE UNIQUE INDEX "cms_cmsplugin_placeholder_id_language_position_bad60ebe_uniq" ON "cms_cmsplugin" ("placeholder_id", "language", "position");

CREATE INDEX "cms_cmsplugin_placeholder_id_0bfa3b26" ON "cms_cmsplugin" ("placeholder_id");

CREATE INDEX "cms_cmsplugin_parent_id_fd3bd9dd" ON "cms_cmsplugin" ("parent_id");

CREATE INDEX "cms_cmsplugin_language_bbea8a48" ON "cms_cmsplugin" ("language");

CREATE INDEX "cms_cmsplugin_plugin_type_94e96ebf" ON "cms_cmsplugin" ("plugin_type");

CREATE INDEX "cms_cmsplug_placeho_3f2b0f_idx" ON "cms_cmsplugin" ("placeholder_id", "language", "position");

CREATE INDEX "cms_pagecontent_language_0b8ccf12" ON "cms_pagecontent" ("language");

CREATE INDEX "cms_pagecontent_slug_23c7b868" ON "cms_pagecontent" ("slug");

CREATE INDEX "cms_pagecontent_page_id_b958451e" ON "cms_pagecontent" ("page_id");

CREATE INDEX "cms_pagecontent_in_navigation_c35c26ba" ON "cms_pagecontent" ("in_navigation");

CREATE INDEX "cms_pagecontent_soft_root_ea2486e9" ON "cms_pagecontent" ("soft_root");

CREATE INDEX "cms_pagecontent_limit_visibility_in_menu_e1d68dcf" ON "cms_pagecontent" ("limit_visibility_in_menu");

CREATE INDEX "cms_aliaspluginmodel_plugin_id_9867676e" ON "cms_aliaspluginmodel" ("plugin_id");

CREATE INDEX "cms_aliaspluginmodel_alias_placeholder_id_6d6e0c8a" ON "cms_aliaspluginmodel" ("alias_placeholder_id");

CREATE UNIQUE INDEX "easy_thumbnails_source_storage_hash_name_481ce32d_uniq" ON "easy_thumbnails_source" ("storage_hash", "name");

CREATE INDEX "easy_thumbnails_source_storage_hash_946cbcc9" ON "easy_thumbnails_source" ("storage_hash");

CREATE INDEX "easy_thumbnails_source_name_5fe0edc6" ON "easy_thumbnails_source" ("name");

CREATE UNIQUE INDEX "easy_thumbnails_thumbnail_storage_hash_name_source_id_fb375270_uniq" ON "easy_thumbnails_thumbnail" ("storage_hash", "name", "source_id");

CREATE INDEX "easy_thumbnails_thumbnail_storage_hash_f1435f49" ON "easy_thumbnails_thumbnail" ("storage_hash");

CREATE INDEX "easy_thumbnails_thumbnail_name_b5882c31" ON "easy_thumbnails_thumbnail" ("name");

CREATE INDEX "easy_thumbnails_thumbnail_source_id_5b57bc77" ON "easy_thumbnails_thumbnail" ("source_id");

CREATE UNIQUE INDEX "filer_folder_parent_id_name_bc773258_uniq" ON "filer_folder" ("parent_id", "name");

CREATE INDEX "filer_folder_parent_id_308aecda" ON "filer_folder" ("parent_id");

CREATE INDEX "filer_folder_owner_id_be530fb4" ON "filer_folder" ("owner_id");

CREATE INDEX "filer_folderpermission_folder_id_5d02f1da" ON "filer_folderpermission" ("folder_id");

CREATE INDEX "filer_folderpermission_user_id_7673d4b6" ON "filer_folderpermission" ("user_id");

CREATE INDEX "filer_folderpermission_group_id_8901bafa" ON "filer_folderpermission" ("group_id");

CREATE INDEX "filer_file_polymorphic_ctype_id_f44903c1" ON "filer_file" ("polymorphic_ctype_id");

CREATE INDEX "filer_file_folder_id_af803bbb" ON "filer_file" ("folder_id");

CREATE INDEX "filer_file_owner_id_b9e32671" ON "filer_file" ("owner_id");

CREATE INDEX "filer_clipboard_user_id_b52ff0bc" ON "filer_clipboard" ("user_id");

CREATE INDEX "filer_clipboarditem_file_id_06196f80" ON "filer_clipboarditem" ("file_id");

CREATE INDEX "filer_clipboarditem_clipboard_id_7a76518b" ON "filer_clipboarditem" ("clipboard_id");

CREATE UNIQUE INDEX "djangocms_versioning_version_content_type_id_object_id_fac3dca5_uniq" ON "djangocms_versioning_version" ("content_type_id", "object_id");

CREATE INDEX "djangocms_versioning_version_created_by_id_0a7b7c9d" ON "djangocms_versioning_version" ("created_by_id");

CREATE INDEX "djangocms_versioning_version_content_type_id_04f95e7d" ON "djangocms_versioning_version" ("content_type_id");

CREATE INDEX "djangocms_versioning_version_locked_by_id_87332f7c" ON "djangocms_versioning_version" ("locked_by_id");

CREATE INDEX "djangocms_versioning_version_source_id_d0664b18" ON "djangocms_versioning_version" ("source_id");

CREATE INDEX "djangocms_versioning_statetracking_version_id_bda1e5ca" ON "djangocms_versioning_statetracking" ("version_id");

CREATE INDEX "djangocms_versioning_statetracking_user_id_93cc4ccb" ON "djangocms_versioning_statetracking" ("user_id");

CREATE UNIQUE INDEX "djangocms_alias_category_translation_name_language_code_45c565e3_uniq" ON "djangocms_alias_category_translation" ("name", "language_code");

CREATE INDEX "djangocms_alias_category_translation_language_code_95644158" ON "djangocms_alias_category_translation" ("language_code");

CREATE INDEX "djangocms_alias_category_translation_master_id_65c3ac44" ON "djangocms_alias_category_translation" ("master_id");

CREATE UNIQUE INDEX "djangocms_alias_alias_static_code_site_id_e5afbc06_uniq" ON "djangocms_alias_alias" ("static_code", "site_id");

CREATE INDEX "djangocms_alias_alias_category_id_96da6da6" ON "djangocms_alias_alias" ("category_id");

CREATE INDEX "djangocms_alias_alias_site_id_d4e42897" ON "djangocms_alias_alias" ("site_id");

CREATE INDEX "djangocms_alias_aliascontent_alias_id_e36dc314" ON "djangocms_alias_aliascontent" ("alias_id");

CREATE INDEX "djangocms_alias_aliasplugin_alias_id_d425895e" ON "djangocms_alias_aliasplugin" ("alias_id");

CREATE INDEX "djangocms_snippet_snippet_snippet_grouper_id_6ab0e516" ON "djangocms_snippet_snippet" ("snippet_grouper_id");

CREATE INDEX "djangocms_snippet_snippet_slug_bd43cd96" ON "djangocms_snippet_snippet" ("slug");

CREATE INDEX "djangocms_snippet_snippet_site_id_aec16190" ON "djangocms_snippet_snippet" ("site_id");

CREATE INDEX "djangocms_snippet_snippetptr_snippet_grouper_id_a7aa3c91" ON "djangocms_snippet_snippetptr" ("snippet_grouper_id");

CREATE INDEX "djangocms_picture_picture_picture_id_f7d6711b" ON "djangocms_picture_picture" ("picture_id");

CREATE INDEX "djangocms_picture_picture_link_page_id_dbec9beb" ON "djangocms_picture_picture" ("link_page_id");

CREATE INDEX "djangocms_picture_picture_thumbnail_options_id_59cf80d1" ON "djangocms_picture_picture" ("thumbnail_options_id");

CREATE INDEX "djangocms_file_file_file_src_id_74ac14a5" ON "djangocms_file_file" ("file_src_id");

CREATE INDEX "djangocms_file_folder_folder_src_id_9558406a" ON "djangocms_file_folder" ("folder_src_id");

CREATE INDEX "djangocms_video_videoplayer_poster_id_07790e24" ON "djangocms_video_videoplayer" ("poster_id");

CREATE INDEX "djangocms_video_videosource_source_file_id_16f11167" ON "djangocms_video_videosource" ("source_file_id");

CREATE INDEX "djangocms_video_videotrack_src_id_e5a015d8" ON "djangocms_video_videotrack" ("src_id");
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

Subclasses (multi-table): [PageUserGroup](#pageusergroup).

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
- [GlobalPagePermission](#globalpagepermission).`group_id` → `id`, on delete cascade
- [PagePermission](#pagepermission).`group_id` → `id`, on delete cascade
- [PageUserGroup](#pageusergroup).`group_ptr_id` → `id`, parent link, on delete cascade
- [FolderPermission](#folderpermission).`group_id` → `id`, on delete cascade

## User

Table `auth_user`, app `auth`.

Subclasses (multi-table): [PageUser](#pageuser).

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
- [UserSettings](#usersettings).`user_id` → `id`, on delete cascade
- [GlobalPagePermission](#globalpagepermission).`user_id` → `id`, on delete cascade
- [PagePermission](#pagepermission).`user_id` → `id`, on delete cascade
- [PageUser](#pageuser).`user_ptr_id` → `id`, parent link, on delete cascade
- [PageUser](#pageuser).`created_by_id` → `id`, on delete set null
- [PageUserGroup](#pageusergroup).`created_by_id` → `id`, on delete set null
- [filer.Folder](#filerfolder).`owner_id` → `id`, on delete set null
- [FolderPermission](#folderpermission).`user_id` → `id`, on delete set null
- [filer.File](#filerfile).`owner_id` → `id`
- [Clipboard](#clipboard).`user_id` → `id`, on delete cascade
- [Version](#version).`created_by_id` → `id`, on delete protect
- [Version](#version).`locked_by_id` → `id`, on delete set null
- [StateTracking](#statetracking).`user_id` → `id`, on delete cascade

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
- [Placeholder](#placeholder).`content_type_id` → `id`, on delete set null
- [filer.File](#filerfile).`polymorphic_ctype_id` → `id`
- [Version](#version).`content_type_id` → `id`, on delete protect

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

- [Page](#page).`site_id` → `id`, on delete cascade
- [PageUrl](#pageurl).`site_id` → `id`, on delete cascade
- `globalpagepermission_set` — many to many with [GlobalPagePermission](#globalpagepermission)
- [Alias](#alias).`site_id` → `id`, on delete cascade
- [Snippet](#snippet).`site_id` → `id`, on delete cascade

## UserSettings

Table `cms_usersettings`, app `cms`.

![UserSettings and the models it relates to](tables/UserSettings.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `OneToOneField` | `integer` |  | FK, unique |  |  |
| `language` | `UserSettingsLanguage` | `varchar(10)` |  |  |  | The language for the admin interface and toolbar |
| `clipboard_id` | `PlaceholderForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `cms_usersettings_clipboard_id_3ae17c19` on `clipboard_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `clipboard_id` → [Placeholder](#placeholder).`id`, on delete cascade

## Page

Table `cms_page`, app `cms`.

Proxy models, on this table: `PageType`.

Columns from `MP_Node` (abstract): `path`, `depth`, `numchild`.

A ``Page`` is the basic unit of site structure in django CMS. The CMS uses a hierarchical page model: each page
stands in relation to other pages as parent, child or sibling. This hierarchy is managed by the `django-treebeard
<http://django-treebeard.readthedocs.io/en/latest/>`_ library.

A ``Page`` is an abstract entity. It does not have any content associated with it, nor does is provide any slugs to
build a URL. This is part of the model :class:`~cms.models.contentmodels.PageContent` and
:class:`~cms.models.pagemodel.PageUrl` respectively.

![Page and the models it relates to](tables/Page.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `path` | `CharField` | `varchar(255)` |  | unique |  |  |
| `depth` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `numchild` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `parent_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `site_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `created_by` | `CharField` | `varchar(255)` |  |  |  |  |
| `changed_by` | `CharField` | `varchar(255)` |  |  |  |  |
| `creation_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `changed_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `reverse_id` | `CharField` | `varchar(40)` | yes |  |  | A unique identifier that is used with the page_url templatetag for linking to this page |
| `navigation_extenders` | `CharField` | `varchar(80)` | yes |  |  |  |
| `login_required` | `BooleanField` | `bool` |  |  | `False` |  |
| `is_home` | `BooleanField` | `bool` |  |  | `False` |  |
| `application_urls` | `CharField` | `varchar(200)` | yes |  |  |  |
| `application_namespace` | `CharField` | `varchar(200)` | yes |  |  |  |
| `is_page_type` | `BooleanField` | `bool` |  |  | `False` | Mark this page as a page type |

Indexes:

- index `cms_page_parent_id_f89b72e4` on `parent_id`
- index `cms_page_site_id_4323d3ff` on `site_id`
- index `cms_page_reverse_id_ffc9ede2` on `reverse_id`
- index `cms_page_navigation_extenders_c24af8dd` on `navigation_extenders`
- index `cms_page_is_home_edadca07` on `is_home`
- index `cms_page_application_urls_9ef47497` on `application_urls`

Relations:

- `parent_id` → [Page](#page).`id`, on delete cascade
- `site_id` → [Site](#site).`id`, on delete cascade
- [PageUrl](#pageurl).`page_id` → `id`, on delete cascade
- [PagePermission](#pagepermission).`page_id` → `id`, on delete cascade
- [PageContent](#pagecontent).`page_id` → `id`, on delete cascade
- [Picture](#picture).`link_page_id` → `id`, on delete cascade

## PageUrl

Table `cms_pageurl`, app `cms`.

![PageUrl and the models it relates to](tables/PageUrl.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `slug` | `SlugField` | `varchar(255)` |  |  |  |  |
| `path` | `CharField` | `varchar(255)` | yes |  |  |  |
| `language` | `CharField` | `varchar(15)` |  |  |  |  |
| `page_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `site_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `managed` | `BooleanField` | `bool` |  |  | `False` |  |

Indexes:

- index `cms_pageurl_slug_f8d42287` on `slug`
- index `cms_pageurl_path_9f26e673` on `path`
- index `cms_pageurl_language_a10302dc` on `language`
- index `cms_pageurl_page_id_dbf4793a` on `page_id`
- index `cms_pageurl_site_id_7d6fd976` on `site_id`
- unique `unique_together_page_language` on `page_id`, `language`
- unique `unique_site_language_path` on `site_id`, `language`, `path`

Relations:

- `page_id` → [Page](#page).`id`, on delete cascade
- `site_id` → [Site](#site).`id`, on delete cascade

## GlobalPagePermission

Table `cms_globalpagepermission`, app `cms`.

Columns from `AbstractPagePermission` (abstract): `user_id`, `group_id`, `can_change`, `can_add`, `can_delete`, `can_publish`, `can_change_advanced_settings`, `can_change_permissions`, `can_move_page`, `can_view`.

Permissions for all pages (global).

![GlobalPagePermission and the models it relates to](tables/GlobalPagePermission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `group_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `can_change` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_add` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_delete` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_publish` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_change_advanced_settings` | `BooleanField` | `bool` |  |  | `False` |  |
| `can_change_permissions` | `BooleanField` | `bool` |  |  | `False` | Allows this user to grant and revoke page permissions for other users. |
| `can_move_page` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_view` | `BooleanField` | `bool` |  |  | `False` | Grants frontend view access. Note: as soon as any user or group is given view access to a page, that page becomes restricted — only users/groups with view access can then see it on the frontend. |
| `can_recover_page` | `BooleanField` | `bool` |  |  | `True` | can recover any deleted page |

Indexes:

- index `cms_globalpagepermission_user_id_a227cee1` on `user_id`
- index `cms_globalpagepermission_group_id_991b4733` on `group_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `group_id` → [Group](#group).`id`, on delete cascade
- `sites` — many to many with [Site](#site)

## PagePermission

Table `cms_pagepermission`, app `cms`.

Columns from `AbstractPagePermission` (abstract): `user_id`, `group_id`, `can_change`, `can_add`, `can_delete`, `can_publish`, `can_change_advanced_settings`, `can_change_permissions`, `can_move_page`, `can_view`.

Page permissions for a single page

![PagePermission and the models it relates to](tables/PagePermission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `group_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `can_change` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_add` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_delete` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_publish` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_change_advanced_settings` | `BooleanField` | `bool` |  |  | `False` |  |
| `can_change_permissions` | `BooleanField` | `bool` |  |  | `False` | Allows this user to grant and revoke page permissions for other users. |
| `can_move_page` | `BooleanField` | `bool` |  |  | `True` |  |
| `can_view` | `BooleanField` | `bool` |  |  | `False` | Grants frontend view access. Note: as soon as any user or group is given view access to a page, that page becomes restricted — only users/groups with view access can then see it on the frontend. |
| `grant_on` | `PagePermissionGrantOn` | `integer` |  |  | `5` |  |
| `page_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `cms_pagepermission_user_id_0c7d2b3c` on `user_id`
- index `cms_pagepermission_group_id_af5af193` on `group_id`
- index `cms_pagepermission_page_id_0ae9a163` on `page_id`

Relations:

- `user_id` → [User](#user).`id`, on delete cascade
- `group_id` → [Group](#group).`id`, on delete cascade
- `page_id` → [Page](#page).`id`, on delete cascade

## PageUser

Table `cms_pageuser`, app `cms`.

Inherits from [User](#user) (multi-table): its rows are joined on `user_ptr_id`.

Cms specific user data, required for permission system

![PageUser and the models it relates to](tables/PageUser.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `user_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `created_by_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `cms_pageuser_created_by_id_8e9fbf83` on `created_by_id`

Relations:

- `user_ptr_id` → [User](#user).`id`, parent link, on delete cascade
- `created_by_id` → [User](#user).`id`, on delete set null

## PageUserGroup

Table `cms_pageusergroup`, app `cms`.

Inherits from [Group](#group) (multi-table): its rows are joined on `group_ptr_id`.

Cms specific group data, required for permission system

![PageUserGroup and the models it relates to](tables/PageUserGroup.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `group_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `created_by_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `cms_pageusergroup_created_by_id_7d57fa39` on `created_by_id`

Relations:

- `group_ptr_id` → [Group](#group).`id`, parent link, on delete cascade
- `created_by_id` → [User](#user).`id`, on delete set null

## Placeholder

Table `cms_placeholder`, app `cms`.

``Placeholders`` can be filled with plugins, which store or generate content.

![Placeholder and the models it relates to](tables/Placeholder.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `slot` | `CharField` | `varchar(255)` |  |  |  |  |
| `default_width` | `PositiveSmallIntegerField` | `smallint unsigned` | yes |  |  |  |
| `content_type_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `object_id` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |

Indexes:

- index `cms_placeholder_slot_0bc04d21` on `slot`
- index `cms_placeholder_content_type_id_a7659d9c` on `content_type_id`

Relations:

- `source` → any model, by `content_type_id` + `object_id` (generic foreign key)
- [UserSettings](#usersettings).`clipboard_id` → `id`, on delete cascade
- `content_type_id` → [ContentType](#contenttype).`id`, on delete set null
- [CMSPlugin](#cmsplugin).`placeholder_id` → `id`, on delete cascade
- `object_id` → [PageContent](#pagecontent).`id` where `content_type_id` is PageContent (generic, `PageContent.placeholders`)
- `object_id` → [PlaceholderReference](#placeholderreference).`cmsplugin_ptr_id` where `content_type_id` is PlaceholderReference (generic, `PlaceholderReference.placeholders`)
- [AliasPluginModel](#aliaspluginmodel).`alias_placeholder_id` → `id`, on delete cascade
- `object_id` → [AliasContent](#aliascontent).`id` where `content_type_id` is AliasContent (generic, `AliasContent.placeholders`)

## CMSPlugin

Table `cms_cmsplugin`, app `cms`.

Subclasses (multi-table): [PlaceholderReference](#placeholderreference), [AliasPluginModel](#aliaspluginmodel), [Text](#text), [AliasPlugin](#aliasplugin), [SnippetPtr](#snippetptr), [Picture](#picture), [djangocms_file.File](#djangocms_filefile), [djangocms_file.Folder](#djangocms_filefolder), [VideoPlayer](#videoplayer), [VideoSource](#videosource), [HlsStreamSource](#hlsstreamsource), [VideoTrack](#videotrack).

The base class for a CMS plugin model. When defining a new custom plugin, you should
store plugin-instance specific information on a subclass of this class. (An example for this
would be to store the number of pictures to display in a gallery.)

Two restrictions apply when subclassing this to use in your own models:

1. Subclasses of CMSPlugin **cannot be further subclassed**
2. Subclasses of CMSPlugin cannot define a "text" field.

![CMSPlugin and the models it relates to](tables/CMSPlugin.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `placeholder_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `parent_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `position` | `SmallIntegerField` | `smallint` |  |  | `1` |  |
| `language` | `CharField` | `varchar(15)` |  |  |  |  |
| `plugin_type` | `CharField` | `varchar(50)` |  |  |  |  |
| `creation_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `changed_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `cms_cmsplugin_placeholder_id_0bfa3b26` on `placeholder_id`
- index `cms_cmsplugin_parent_id_fd3bd9dd` on `parent_id`
- index `cms_cmsplugin_language_bbea8a48` on `language`
- index `cms_cmsplugin_plugin_type_94e96ebf` on `plugin_type`
- unique `cms_cmsplugin_placeholder_id_language_position_bad60ebe_uniq` on `placeholder_id`, `language`, `position`
- index `cms_cmsplug_placeho_3f2b0f_idx` on `placeholder_id`, `language`, `position`

Relations:

- `placeholder_id` → [Placeholder](#placeholder).`id`, on delete cascade
- `parent_id` → [CMSPlugin](#cmsplugin).`id`, on delete cascade
- [PlaceholderReference](#placeholderreference).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [AliasPluginModel](#aliaspluginmodel).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [AliasPluginModel](#aliaspluginmodel).`plugin_id` → `id`, on delete cascade
- [Text](#text).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [AliasPlugin](#aliasplugin).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [SnippetPtr](#snippetptr).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [Picture](#picture).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [djangocms_file.File](#djangocms_filefile).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [djangocms_file.Folder](#djangocms_filefolder).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [VideoPlayer](#videoplayer).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [VideoSource](#videosource).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [HlsStreamSource](#hlsstreamsource).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade
- [VideoTrack](#videotrack).`cmsplugin_ptr_id` → `id`, parent link, on delete cascade

## PageContent

Table `cms_pagecontent`, app `cms`.

![PageContent and the models it relates to](tables/PageContent.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `language` | `CharField` | `varchar(15)` |  |  |  |  |
| `title` | `CharField` | `varchar(255)` |  |  |  | The default title |
| `page_title` | `CharField` | `varchar(255)` | yes |  |  | Overwrites what is displayed at the top of your browser or in bookmarks |
| `menu_title` | `CharField` | `varchar(255)` | yes |  |  | Overwrite what is displayed in the menu |
| `meta_description` | `TextField` | `text` | yes |  |  | A description of the page used by search engines. |
| `redirect` | `CharField` | `varchar(2048)` | yes |  |  | Redirects to this URL. |
| `slug` | `SlugField` | `varchar(255)` |  |  |  | The part of the title that is used in the URL |
| `overwrite_url` | `CharField` | `varchar(255)` | yes |  |  | Keep this field empty if standard path should be used. |
| `page_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `creation_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `created_by` | `CharField` | `varchar(255)` |  |  |  |  |
| `changed_by` | `CharField` | `varchar(255)` |  |  |  |  |
| `changed_date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `in_navigation` | `BooleanField` | `bool` |  |  | `True` |  |
| `soft_root` | `BooleanField` | `bool` |  |  | `False` | All ancestors will not be displayed in the navigation |
| `template` | `CharField` | `varchar(100)` |  |  | `'INHERIT'` | The template used to render the content. |
| `limit_visibility_in_menu` | `PageContentLimitVisibilityInMenu` | `smallint` | yes |  | `None` | limit when this page is visible in the menu |
| `xframe_options` | `PageContentXframeOptions` | `integer` |  |  | `0` |  |

Indexes:

- index `cms_pagecontent_language_0b8ccf12` on `language`
- index `cms_pagecontent_slug_23c7b868` on `slug`
- index `cms_pagecontent_page_id_b958451e` on `page_id`
- index `cms_pagecontent_in_navigation_c35c26ba` on `in_navigation`
- index `cms_pagecontent_soft_root_ea2486e9` on `soft_root`
- index `cms_pagecontent_limit_visibility_in_menu_e1d68dcf` on `limit_visibility_in_menu`

Relations:

- `page_id` → [Page](#page).`id`, on delete cascade
- `placeholders` ← [Placeholder](#placeholder), by `content_type_id` + `object_id` (generic relation)
- `versions` ← [Version](#version), by `content_type_id` + `object_id` (generic relation)

## PlaceholderReference

Table `cms_placeholderreference`, app `cms`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

![PlaceholderReference and the models it relates to](tables/PlaceholderReference.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |

Relations:

- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade
- `placeholders` ← [Placeholder](#placeholder), by `content_type_id` + `object_id` (generic relation)

## AliasPluginModel

Table `cms_aliaspluginmodel`, app `cms`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

AliasPlugin is deprecated,
and it will be removed;
please use the package djangocms-alias instead

![AliasPluginModel and the models it relates to](tables/AliasPluginModel.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `plugin_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `alias_placeholder_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `cms_aliaspluginmodel_plugin_id_9867676e` on `plugin_id`
- index `cms_aliaspluginmodel_alias_placeholder_id_6d6e0c8a` on `alias_placeholder_id`

Relations:

- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade
- `plugin_id` → [CMSPlugin](#cmsplugin).`id`, on delete cascade
- `alias_placeholder_id` → [Placeholder](#placeholder).`id`, on delete cascade

## UrlconfRevision

Table `cms_urlconfrevision`, app `cms`.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `revision` | `CharField` | `varchar(255)` |  |  |  |  |

## CacheKey

Table `menus_cachekey`, app `menus`.

This model stores a set of cache keys accessible by multiple processes/machines.
Multiple Django instances will then share the keys, allowing selective invalidation
of menu trees (per site, per language) in the cache.

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `language` | `CharField` | `varchar(255)` |  |  |  |  |
| `site` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `key` | `CharField` | `varchar(255)` |  |  |  |  |

## Source

Table `easy_thumbnails_source`, app `easy_thumbnails`.

Columns from `File` (abstract): `storage_hash`, `name`, `modified`.

![Source and the models it relates to](tables/Source.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `storage_hash` | `CharField` | `varchar(40)` |  |  |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `easy_thumbnails_source_storage_hash_946cbcc9` on `storage_hash`
- index `easy_thumbnails_source_name_5fe0edc6` on `name`
- unique `easy_thumbnails_source_storage_hash_name_481ce32d_uniq` on `storage_hash`, `name`

Relations:

- [Thumbnail](#thumbnail).`source_id` → `id`, on delete cascade

## Thumbnail

Table `easy_thumbnails_thumbnail`, app `easy_thumbnails`.

Columns from `File` (abstract): `storage_hash`, `name`, `modified`.

![Thumbnail and the models it relates to](tables/Thumbnail.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `storage_hash` | `CharField` | `varchar(40)` |  |  |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `source_id` | `ForeignKey` | `integer` |  | FK |  |  |

Indexes:

- index `easy_thumbnails_thumbnail_storage_hash_f1435f49` on `storage_hash`
- index `easy_thumbnails_thumbnail_name_b5882c31` on `name`
- index `easy_thumbnails_thumbnail_source_id_5b57bc77` on `source_id`
- unique `easy_thumbnails_thumbnail_storage_hash_name_source_id_fb375270_uniq` on `storage_hash`, `name`, `source_id`

Relations:

- `source_id` → [Source](#source).`id`, on delete cascade
- [ThumbnailDimensions](#thumbnaildimensions).`thumbnail_id` → `id`, on delete cascade

## ThumbnailDimensions

Table `easy_thumbnails_thumbnaildimensions`, app `easy_thumbnails`.

![ThumbnailDimensions and the models it relates to](tables/ThumbnailDimensions.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `thumbnail_id` | `OneToOneField` | `integer` |  | FK, unique |  |  |
| `width` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |
| `height` | `PositiveIntegerField` | `integer unsigned` | yes |  |  |  |

Relations:

- `thumbnail_id` → [Thumbnail](#thumbnail).`id`, on delete cascade

## filer.Folder

Table `filer_folder`, app `filer`.

Represents a Folder that things (files) can be put into. Folders are *NOT*
mirrored in the Filesystem and can have any unicode chars as their name.
Other models may attach to a folder with a ForeignKey. If the related name
ends with "_files" they will automatically be listed in the
folder.files list along with all the other models that link to the folder
in this way. Make sure the linked models obey the AbstractFile interface
(Duck Type).

![filer.Folder and the models it relates to](tables/filer.Folder.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `parent_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `owner_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `uploaded_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `created_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `modified_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |

Indexes:

- index `filer_folder_parent_id_308aecda` on `parent_id`
- index `filer_folder_owner_id_be530fb4` on `owner_id`
- unique `filer_folder_parent_id_name_bc773258_uniq` on `parent_id`, `name`

Relations:

- `parent_id` → [filer.Folder](#filerfolder).`id`, on delete cascade
- `owner_id` → [User](#user).`id`, on delete set null
- [FolderPermission](#folderpermission).`folder_id` → `id`, on delete cascade
- [filer.File](#filerfile).`folder_id` → `id`
- [djangocms_file.Folder](#djangocms_filefolder).`folder_src_id` → `id`, on delete set null

## FolderPermission

Table `filer_folderpermission`, app `filer`.

![FolderPermission and the models it relates to](tables/FolderPermission.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `folder_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `type` | `FolderPermissionType` | `smallint` |  |  | `0` |  |
| `user_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `group_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `everybody` | `BooleanField` | `bool` |  |  | `False` |  |
| `can_read` | `FolderPermissionCanRead` | `smallint` | yes |  | `None` |  |
| `can_edit` | `FolderPermissionCanEdit` | `smallint` | yes |  | `None` |  |
| `can_add_children` | `FolderPermissionCanAddChildren` | `smallint` | yes |  | `None` |  |

Indexes:

- index `filer_folderpermission_folder_id_5d02f1da` on `folder_id`
- index `filer_folderpermission_user_id_7673d4b6` on `user_id`
- index `filer_folderpermission_group_id_8901bafa` on `group_id`

Relations:

- `folder_id` → [filer.Folder](#filerfolder).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete set null
- `group_id` → [Group](#group).`id`, on delete cascade

## filer.File

Table `filer_file`, app `filer`.

Subclasses (multi-table): [Image](#image).

Columns from `PolymorphicModel` (abstract): `polymorphic_ctype_id`.

![filer.File and the models it relates to](tables/filer.File.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `polymorphic_ctype_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `folder_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `file` | `MultiStorageFileField` | `varchar(255)` | yes |  |  |  |
| `_file_size` | `BigIntegerField` | `bigint` | yes |  |  |  |
| `sha1` | `CharField` | `varchar(40)` |  |  | `''` |  |
| `has_all_mandatory_data` | `BooleanField` | `bool` |  |  | `False` |  |
| `original_filename` | `CharField` | `varchar(255)` | yes |  |  |  |
| `name` | `CharField` | `varchar(255)` |  |  | `''` |  |
| `description` | `TextField` | `text` | yes |  |  |  |
| `owner_id` | `ForeignKey` | `integer` | yes | FK |  |  |
| `uploaded_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `modified_at` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `is_public` | `BooleanField` | `bool` |  |  | `is_public_default()` | Disable any permission checking for this file. File will be publicly accessible to anyone. |
| `mime_type` | `CharField` | `varchar(255)` |  |  | `'application/octet-stream'` | MIME type of uploaded content |

Indexes:

- index `filer_file_polymorphic_ctype_id_f44903c1` on `polymorphic_ctype_id`
- index `filer_file_folder_id_af803bbb` on `folder_id`
- index `filer_file_owner_id_b9e32671` on `owner_id`

Relations:

- `in_clipboards` — many to many with [Clipboard](#clipboard), through [ClipboardItem](#clipboarditem)
- `polymorphic_ctype_id` → [ContentType](#contenttype).`id`
- `folder_id` → [filer.Folder](#filerfolder).`id`
- `owner_id` → [User](#user).`id`
- [ClipboardItem](#clipboarditem).`file_id` → `id`, on delete cascade
- [Image](#image).`file_ptr_id` → `id`, parent link
- [djangocms_file.File](#djangocms_filefile).`file_src_id` → `id`, on delete set null
- [VideoSource](#videosource).`source_file_id` → `id`, on delete set null
- [VideoTrack](#videotrack).`src_id` → `id`, on delete set null

## Clipboard

Table `filer_clipboard`, app `filer`.

![Clipboard and the models it relates to](tables/Clipboard.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |

Indexes:

- index `filer_clipboard_user_id_b52ff0bc` on `user_id`

Relations:

- `files` — many to many with [filer.File](#filerfile), through [ClipboardItem](#clipboarditem)
- `user_id` → [User](#user).`id`, on delete cascade
- [ClipboardItem](#clipboarditem).`clipboard_id` → `id`, on delete cascade

## ClipboardItem

Table `filer_clipboarditem`, app `filer`.

![ClipboardItem and the models it relates to](tables/ClipboardItem.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `file_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `clipboard_id` | `ForeignKey` | `integer` |  | FK |  |  |

Indexes:

- index `filer_clipboarditem_file_id_06196f80` on `file_id`
- index `filer_clipboarditem_clipboard_id_7a76518b` on `clipboard_id`

Relations:

- `file_id` → [filer.File](#filerfile).`id`, on delete cascade
- `clipboard_id` → [Clipboard](#clipboard).`id`, on delete cascade

## Image

Table `filer_image`, app `filer`.

Inherits from [filer.File](#filerfile) (multi-table): its rows are joined on `file_ptr_id`.

Columns from `BaseImage` (abstract): `_height`, `_width`, `_transparent`, `default_alt_text`, `default_caption`, `subject_location`, `file_ptr_id`.

![Image and the models it relates to](tables/Image.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `_height` | `FloatField` | `real` | yes |  |  |  |
| `_width` | `FloatField` | `real` | yes |  |  |  |
| `_transparent` | `BooleanField` | `bool` |  |  | `False` |  |
| `default_alt_text` | `CharField` | `varchar(255)` | yes |  |  |  |
| `default_caption` | `CharField` | `varchar(255)` | yes |  |  |  |
| `subject_location` | `CharField` | `varchar(64)` |  |  | `''` |  |
| `file_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `date_taken` | `DateTimeField` | `datetime` | yes |  |  |  |
| `author` | `CharField` | `varchar(255)` | yes |  |  |  |
| `must_always_publish_author_credit` | `BooleanField` | `bool` |  |  | `False` |  |
| `must_always_publish_copyright` | `BooleanField` | `bool` |  |  | `False` |  |

Relations:

- `file_ptr_id` → [filer.File](#filerfile).`id`, parent link
- [Picture](#picture).`picture_id` → `file_ptr_id`, on delete set null
- [VideoPlayer](#videoplayer).`poster_id` → `file_ptr_id`, on delete set null

## ThumbnailOption

Table `filer_thumbnailoption`, app `filer`.

This class defines the option use to create the thumbnail.

![ThumbnailOption and the models it relates to](tables/ThumbnailOption.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(100)` |  |  |  |  |
| `width` | `IntegerField` | `integer` |  |  |  | width in pixel. |
| `height` | `IntegerField` | `integer` |  |  |  | height in pixel. |
| `crop` | `BooleanField` | `bool` |  |  | `True` |  |
| `upscale` | `BooleanField` | `bool` |  |  | `True` |  |

Relations:

- [Picture](#picture).`thumbnail_options_id` → `id`, on delete cascade

## Text

Table `djangocms_text_text`, app `djangocms_text`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

Columns from `AbstractText` (abstract): `cmsplugin_ptr_id`, `body`, `json`, `rte`.

![Text and the models it relates to](tables/Text.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `body` | `TextField` | `text` |  |  |  |  |
| `json` | `JSONField` | `text` | yes |  |  |  |
| `rte` | `CharField` | `varchar(16)` |  |  | `''` | The rich text editor used to create this text. JSON formats vary between editors. |

Relations:

- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade

## Version

Table `djangocms_versioning_version`, app `djangocms_versioning`.

Proxy models, on this table: `PageContentVersion`, `AliasContentVersion`, `SnippetVersion`.

![Version and the models it relates to](tables/Version.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `created` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `modified` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `created_by_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `number` | `CharField` | `varchar(11)` |  |  |  |  |
| `content_type_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `object_id` | `PositiveIntegerField` | `integer unsigned` |  |  |  |  |
| `state` | `VersionState` | `varchar(50)` |  |  | `'draft'` |  |
| `locked_by_id` | `ForeignKey` | `integer` | yes | FK | `None` |  |
| `source_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `djangocms_versioning_version_created_by_id_0a7b7c9d` on `created_by_id`
- index `djangocms_versioning_version_content_type_id_04f95e7d` on `content_type_id`
- index `djangocms_versioning_version_locked_by_id_87332f7c` on `locked_by_id`
- index `djangocms_versioning_version_source_id_d0664b18` on `source_id`
- unique `djangocms_versioning_version_content_type_id_object_id_fac3dca5_uniq` on `content_type_id`, `object_id`

Relations:

- `content` → any model, by `content_type_id` + `object_id` (generic foreign key)
- `object_id` → [PageContent](#pagecontent).`id` where `content_type_id` is PageContent (generic, `PageContent.versions`)
- `created_by_id` → [User](#user).`id`, on delete protect
- `content_type_id` → [ContentType](#contenttype).`id`, on delete protect
- `locked_by_id` → [User](#user).`id`, on delete set null
- `source_id` → [Version](#version).`id`, on delete allow deleting versions
- [StateTracking](#statetracking).`version_id` → `id`, on delete cascade
- `object_id` → [AliasContent](#aliascontent).`id` where `content_type_id` is AliasContent (generic, `AliasContent.versions`)
- `object_id` → [Snippet](#snippet).`id` where `content_type_id` is Snippet (generic, `Snippet.versions`)

## StateTracking

Table `djangocms_versioning_statetracking`, app `djangocms_versioning`.

![StateTracking and the models it relates to](tables/StateTracking.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `AutoField` | `integer` |  | PK | `autoincrement()` |  |
| `version_id` | `ForeignKey` | `integer` |  | FK |  |  |
| `date` | `DateTimeField` | `datetime` |  |  | `now()` |  |
| `old_state` | `StateTrackingOldState` | `varchar(100)` |  |  |  |  |
| `new_state` | `StateTrackingNewState` | `varchar(100)` |  |  |  |  |
| `user_id` | `ForeignKey` | `integer` |  | FK |  |  |

Indexes:

- index `djangocms_versioning_statetracking_version_id_bda1e5ca` on `version_id`
- index `djangocms_versioning_statetracking_user_id_93cc4ccb` on `user_id`

Relations:

- `version_id` → [Version](#version).`id`, on delete cascade
- `user_id` → [User](#user).`id`, on delete cascade

## CategoryTranslation

Table `djangocms_alias_category_translation`, app `djangocms_alias`.

Columns from `TranslatedFieldsModel` (abstract): `language_code`.

![CategoryTranslation and the models it relates to](tables/CategoryTranslation.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `language_code` | `CategoryTranslationLanguageCode` | `varchar(15)` |  |  |  |  |
| `name` | `CharField` | `varchar(120)` |  |  |  |  |
| `master_id` | `TranslationsForeignKey` | `bigint` | yes | FK |  |  |

Indexes:

- index `djangocms_alias_category_translation_language_code_95644158` on `language_code`
- index `djangocms_alias_category_translation_master_id_65c3ac44` on `master_id`
- unique `djangocms_alias_category_translation_name_language_code_45c565e3_uniq` on `name`, `language_code`
- unique `djangocms_alias_category_translation_uniq_lang` on `language_code`, `master_id`

Relations:

- `master_id` → [Category](#category).`id`, on delete cascade

## Category

Table `djangocms_alias_category`, app `djangocms_alias`.

![Category and the models it relates to](tables/Category.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |

Relations:

- [CategoryTranslation](#categorytranslation).`master_id` → `id`, on delete cascade
- [Alias](#alias).`category_id` → `id`, on delete protect

## Alias

Table `djangocms_alias_alias`, app `djangocms_alias`.

![Alias and the models it relates to](tables/Alias.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `creation_method` | `AliasCreationMethod` | `varchar(20)` |  |  | `'code'` |  |
| `category_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `position` | `PositiveIntegerField` | `integer unsigned` |  |  | `0` |  |
| `static_code` | `CharField` | `varchar(255)` | yes |  |  | To render the alias in templates. |
| `site_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `djangocms_alias_alias_category_id_96da6da6` on `category_id`
- index `djangocms_alias_alias_site_id_d4e42897` on `site_id`
- unique `djangocms_alias_alias_static_code_site_id_e5afbc06_uniq` on `static_code`, `site_id`

Relations:

- `category_id` → [Category](#category).`id`, on delete protect
- `site_id` → [Site](#site).`id`, on delete cascade
- [AliasContent](#aliascontent).`alias_id` → `id`, on delete cascade
- [AliasPlugin](#aliasplugin).`alias_id` → `id`, on delete protect

## AliasContent

Table `djangocms_alias_aliascontent`, app `djangocms_alias`.

![AliasContent and the models it relates to](tables/AliasContent.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `alias_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `name` | `CharField` | `varchar(120)` |  |  |  |  |
| `language` | `CharField` | `varchar(10)` |  |  | `get_language()` |  |

Indexes:

- index `djangocms_alias_aliascontent_alias_id_e36dc314` on `alias_id`

Relations:

- `alias_id` → [Alias](#alias).`id`, on delete cascade
- `placeholders` ← [Placeholder](#placeholder), by `content_type_id` + `object_id` (generic relation)
- `versions` ← [Version](#version), by `content_type_id` + `object_id` (generic relation)

## AliasPlugin

Table `djangocms_alias_aliasplugin`, app `djangocms_alias`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

![AliasPlugin and the models it relates to](tables/AliasPlugin.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `alias_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `template` | `AliasPluginTemplate` | `varchar(255)` |  |  | `'default'` |  |

Indexes:

- index `djangocms_alias_aliasplugin_alias_id_d425895e` on `alias_id`

Relations:

- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade
- `alias_id` → [Alias](#alias).`id`, on delete protect

## SnippetGrouper

Table `djangocms_snippet_snippetgrouper`, app `djangocms_snippet`.

The Grouper model for snippet, this is required for versioning

![SnippetGrouper and the models it relates to](tables/SnippetGrouper.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |

Relations:

- [Snippet](#snippet).`snippet_grouper_id` → `id`, on delete protect
- [SnippetPtr](#snippetptr).`snippet_grouper_id` → `id`, on delete cascade

## Snippet

Table `djangocms_snippet_snippet`, app `djangocms_snippet`.

A snippet of HTML or a Django template

![Snippet and the models it relates to](tables/Snippet.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `id` | `BigAutoField` | `integer` |  | PK | `autoincrement()` |  |
| `name` | `CharField` | `varchar(255)` |  |  |  |  |
| `snippet_grouper_id` | `ForeignKey` | `bigint` |  | FK |  |  |
| `html` | `TextField` | `text` |  |  |  |  |
| `template` | `CharField` | `varchar(255)` |  |  |  | Enter a template (e.g. "snippets/plugin_xy.html") to be rendered. If "template" is given, the contents of field "HTML" will be passed as template variable {{ html }} to the template. Otherwise, the content of "HTML" is rendered. |
| `slug` | `SlugField` | `varchar(255)` |  |  | `''` |  |
| `site_id` | `ForeignKey` | `integer` | yes | FK |  |  |

Indexes:

- index `djangocms_snippet_snippet_snippet_grouper_id_6ab0e516` on `snippet_grouper_id`
- index `djangocms_snippet_snippet_slug_bd43cd96` on `slug`
- index `djangocms_snippet_snippet_site_id_aec16190` on `site_id`

Relations:

- `snippet_grouper_id` → [SnippetGrouper](#snippetgrouper).`id`, on delete protect
- `site_id` → [Site](#site).`id`, on delete cascade
- `versions` ← [Version](#version), by `content_type_id` + `object_id` (generic relation)

## SnippetPtr

Table `djangocms_snippet_snippetptr`, app `djangocms_snippet`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

![SnippetPtr and the models it relates to](tables/SnippetPtr.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `snippet_grouper_id` | `ForeignKey` | `bigint` |  | FK |  |  |

Indexes:

- index `djangocms_snippet_snippetptr_snippet_grouper_id_a7aa3c91` on `snippet_grouper_id`

Relations:

- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade
- `snippet_grouper_id` → [SnippetGrouper](#snippetgrouper).`id`, on delete cascade

## Picture

Table `djangocms_picture_picture`, app `djangocms_picture`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

Columns from `AbstractPicture` (abstract): `template`, `picture_id`, `external_picture`, `width`, `height`, `alignment`, `caption_text`, `attributes`, `link_url`, `link_page_id`, `link_target`, `link_attributes`, `use_automatic_scaling`, `use_no_cropping`, `use_crop`, `use_upscale`, `use_responsive_image`, `thumbnail_options_id`, `cmsplugin_ptr_id`.

![Picture and the models it relates to](tables/Picture.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `template` | `PictureTemplate` | `varchar(255)` |  |  | `'default'` |  |
| `picture_id` | `FilerImageField` | `integer` | yes | FK |  |  |
| `external_picture` | `URLField` | `varchar(255)` | yes |  |  | If provided, overrides the embedded image. Certain options such as cropping are not applicable to external images. |
| `width` | `PositiveIntegerField` | `integer unsigned` | yes |  |  | The image width as number in pixels. Example: "720" and not "720px". |
| `height` | `PositiveIntegerField` | `integer unsigned` | yes |  |  | The image height as number in pixels. Example: "720" and not "720px". |
| `alignment` | `PictureAlignment` | `varchar(255)` |  |  |  | Aligns the image according to the selected option. |
| `caption_text` | `TextField` | `text` | yes |  |  | Provide a description, attribution, copyright or other information. |
| `attributes` | `AttributesField` | `text` |  |  | `dict()` |  |
| `link_url` | `URLField` | `varchar(2040)` | yes |  |  | Wraps the image in a link to an external URL. |
| `link_page_id` | `PageField` | `integer` | yes | FK |  | Wraps the image in a link to an internal (page) URL. |
| `link_target` | `PictureLinkTarget` | `varchar(255)` |  |  |  |  |
| `link_attributes` | `AttributesField` | `text` |  |  | `dict()` |  |
| `use_automatic_scaling` | `BooleanField` | `bool` |  |  | `True` | Uses the placeholder dimensions to automatically calculate the size. |
| `use_no_cropping` | `BooleanField` | `bool` |  |  | `False` | Outputs the raw image without cropping. |
| `use_crop` | `BooleanField` | `bool` |  |  | `False` | Crops the image according to the thumbnail settings provided in the template. |
| `use_upscale` | `BooleanField` | `bool` |  |  | `False` | Upscales the image to the size of the thumbnail settings in the template. |
| `use_responsive_image` | `PictureUseResponsiveImage` | `varchar(7)` |  |  | `'inherit'` | Uses responsive image technique to choose better image to display based upon screen viewport. This configuration only applies to uploaded images (external pictures will not be affected). |
| `thumbnail_options_id` | `ForeignKey` | `integer` | yes | FK |  | Overrides width, height, and crop; scales up to the provided preset dimensions. |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |

Indexes:

- index `djangocms_picture_picture_picture_id_f7d6711b` on `picture_id`
- index `djangocms_picture_picture_link_page_id_dbec9beb` on `link_page_id`
- index `djangocms_picture_picture_thumbnail_options_id_59cf80d1` on `thumbnail_options_id`

Relations:

- `picture_id` → [Image](#image).`file_ptr_id`, on delete set null
- `link_page_id` → [Page](#page).`id`, on delete cascade
- `thumbnail_options_id` → [ThumbnailOption](#thumbnailoption).`id`, on delete cascade
- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade

## djangocms_file.File

Table `djangocms_file_file`, app `djangocms_file`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

Columns from `AbstractFile` (abstract): `template`, `file_src_id`, `file_name`, `link_target`, `link_title`, `show_file_size`, `attributes`, `cmsplugin_ptr_id`.

![djangocms_file.File and the models it relates to](tables/djangocms_file.File.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `template` | `DjangocmsFileFileTemplate` | `varchar(255)` |  |  | `'default'` |  |
| `file_src_id` | `FilerFileField` | `integer` | yes | FK |  |  |
| `file_name` | `CharField` | `varchar(255)` |  |  |  | Overrides the default file name with the given value. |
| `link_target` | `DjangocmsFileFileLinkTarget` | `varchar(255)` |  |  | `''` |  |
| `link_title` | `CharField` | `varchar(255)` |  |  |  |  |
| `show_file_size` | `BooleanField` | `bool` |  |  | `False` | Appends the file size at the end of the name. |
| `attributes` | `AttributesField` | `text` |  |  | `dict()` |  |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |

Indexes:

- index `djangocms_file_file_file_src_id_74ac14a5` on `file_src_id`

Relations:

- `file_src_id` → [filer.File](#filerfile).`id`, on delete set null
- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade

## djangocms_file.Folder

Table `djangocms_file_folder`, app `djangocms_file`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

Columns from `AbstractFolder` (abstract): `template`, `folder_src_id`, `link_target`, `show_file_size`, `attributes`, `cmsplugin_ptr_id`.

![djangocms_file.Folder and the models it relates to](tables/djangocms_file.Folder.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `template` | `DjangocmsFileFolderTemplate` | `varchar(255)` |  |  | `'default'` |  |
| `folder_src_id` | `FilerFolderField` | `integer` | yes | FK |  |  |
| `link_target` | `DjangocmsFileFolderLinkTarget` | `varchar(255)` |  |  | `''` |  |
| `show_file_size` | `BooleanField` | `bool` |  |  | `False` | Appends the file size at the end of the name. |
| `attributes` | `AttributesField` | `text` |  |  | `dict()` |  |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |

Indexes:

- index `djangocms_file_folder_folder_src_id_9558406a` on `folder_src_id`

Relations:

- `folder_src_id` → [filer.Folder](#filerfolder).`id`, on delete set null
- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade

## VideoPlayer

Table `djangocms_video_videoplayer`, app `djangocms_video`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

Renders either an Iframe when ``link`` is provided or the HTML5 <video> tag

![VideoPlayer and the models it relates to](tables/VideoPlayer.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `template` | `VideoPlayerTemplate` | `varchar(255)` |  |  | `'default'` |  |
| `label` | `CharField` | `varchar(255)` |  |  |  |  |
| `embed_link` | `TextField` | `text` |  |  |  | Use this field to embed videos from external services such as YouTube, Vimeo or others. Leave it blank to upload video files by adding nested "Source" plugins. |
| `parameters` | `AttributesField` | `text` |  |  | `dict()` | Parameters are appended to the video link if provided. |
| `poster_id` | `FilerImageField` | `integer` | yes | FK |  |  |
| `attributes` | `AttributesField` | `text` |  |  | `dict()` |  |
| `show_controls` | `BooleanField` | `bool` |  |  | `True` | If enabled, the video will be shown with Play, Pause and Seek elements that allow the user to control playback. |
| `autoplay` | `BooleanField` | `bool` |  |  | `False` | If enabled, the video will automatically play once the page is loaded. This might not work depending on how the user has configured their browser. |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |

Indexes:

- index `djangocms_video_videoplayer_poster_id_07790e24` on `poster_id`

Relations:

- `poster_id` → [Image](#image).`file_ptr_id`, on delete set null
- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade

## VideoSource

Table `djangocms_video_videosource`, app `djangocms_video`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

Renders the HTML <source> element inside of <video>.

![VideoSource and the models it relates to](tables/VideoSource.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `source_file_id` | `FilerFileField` | `integer` | yes | FK |  |  |
| `text_title` | `CharField` | `varchar(255)` |  |  |  |  |
| `text_description` | `TextField` | `text` |  |  |  |  |
| `attributes` | `AttributesField` | `text` |  |  | `dict()` |  |

Indexes:

- index `djangocms_video_videosource_source_file_id_16f11167` on `source_file_id`

Relations:

- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade
- `source_file_id` → [filer.File](#filerfile).`id`, on delete set null

## HlsStreamSource

Table `djangocms_video_hlsstreamsource`, app `djangocms_video`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

Renders the HTML <source> element inside of <video> for an HLS stream defined by a .m3u8 URL.

![HlsStreamSource and the models it relates to](tables/HlsStreamSource.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `hls_source_url` | `URLField` | `varchar(200)` |  |  |  |  |

Relations:

- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade

## VideoTrack

Table `djangocms_video_videotrack`, app `djangocms_video`.

Inherits from [CMSPlugin](#cmsplugin) (multi-table): its rows are joined on `cmsplugin_ptr_id`.

Renders the HTML <track> element inside <video>.

![VideoTrack and the models it relates to](tables/VideoTrack.svg)

| Column | Field | Type | Null | Key | Default | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `cmsplugin_ptr_id` | `OneToOneField` | `integer` |  | PK, FK |  |  |
| `kind` | `VideoTrackKind` | `varchar(255)` |  |  |  |  |
| `src_id` | `FilerFileField` | `integer` | yes | FK |  |  |
| `srclang` | `CharField` | `varchar(255)` |  |  |  | Examples: "en" or "de" etc. |
| `label` | `CharField` | `varchar(255)` |  |  |  |  |
| `attributes` | `AttributesField` | `text` |  |  | `dict()` |  |

Indexes:

- index `djangocms_video_videotrack_src_id_e5a015d8` on `src_id`

Relations:

- `cmsplugin_ptr_id` → [CMSPlugin](#cmsplugin).`id`, parent link, on delete cascade
- `src_id` → [filer.File](#filerfile).`id`, on delete set null

# Enums

## LogEntryActionFlag

| Name | Stored as |
| --- | --- |
| `Addition` | `1` |
| `Change` | `2` |
| `Deletion` | `3` |

## UserSettingsLanguage

| Name | Stored as |
| --- | --- |
| `English` | `en` |

## PagePermissionGrantOn

| Name | Stored as |
| --- | --- |
| `Current page` | `1` |
| `Page children (immediate)` | `2` |
| `Page and children (immediate)` | `3` |
| `Page descendants` | `4` |
| `Page and descendants` | `5` |

## PageContentLimitVisibilityInMenu

| Name | Stored as |
| --- | --- |
| `no limit set` | `None` |
| `for logged in users only` | `1` |
| `for anonymous users only` | `2` |

## PageContentXframeOptions

| Name | Stored as |
| --- | --- |
| `Inherit from parent page` | `0` |
| `Deny` | `1` |
| `Only this website` | `2` |
| `Allow` | `3` |

## FolderPermissionType

| Name | Stored as |
| --- | --- |
| `all items` | `0` |
| `this item only` | `1` |
| `this item and all children` | `2` |

## FolderPermissionCanRead

| Name | Stored as |
| --- | --- |
| `inherit` | `None` |
| `allow` | `1` |
| `deny` | `0` |

## FolderPermissionCanEdit

| Name | Stored as |
| --- | --- |
| `inherit` | `None` |
| `allow` | `1` |
| `deny` | `0` |

## FolderPermissionCanAddChildren

| Name | Stored as |
| --- | --- |
| `inherit` | `None` |
| `allow` | `1` |
| `deny` | `0` |

## VersionState

| Name | Stored as |
| --- | --- |
| `Draft` | `draft` |
| `Published` | `published` |
| `Unpublished` | `unpublished` |
| `Archived` | `archived` |

## StateTrackingOldState

| Name | Stored as |
| --- | --- |
| `Draft` | `draft` |
| `Published` | `published` |
| `Unpublished` | `unpublished` |
| `Archived` | `archived` |

## StateTrackingNewState

| Name | Stored as |
| --- | --- |
| `Draft` | `draft` |
| `Published` | `published` |
| `Unpublished` | `unpublished` |
| `Archived` | `archived` |

## CategoryTranslationLanguageCode

| Name | Stored as |
| --- | --- |
| `English` | `en` |

## AliasCreationMethod

| Name | Stored as |
| --- | --- |
| `by template` | `template` |
| `by code` | `code` |

## AliasPluginTemplate

| Name | Stored as |
| --- | --- |
| `Default` | `default` |

## PictureTemplate

| Name | Stored as |
| --- | --- |
| `Default` | `default` |

## PictureAlignment

| Name | Stored as |
| --- | --- |
| `Align left` | `left` |
| `Align right` | `right` |
| `Align center` | `center` |

## PictureLinkTarget

| Name | Stored as |
| --- | --- |
| `Open in new window` | `_blank` |
| `Open in same window` | `_self` |
| `Delegate to parent` | `_parent` |
| `Delegate to top` | `_top` |

## PictureUseResponsiveImage

| Name | Stored as |
| --- | --- |
| `Let settings.DJANGOCMS_PICTURE_RESPONSIVE_IMAGES decide` | `inherit` |
| `Yes` | `yes` |
| `No` | `no` |

## DjangocmsFileFileTemplate

| Name | Stored as |
| --- | --- |
| `Default` | `default` |

## DjangocmsFileFileLinkTarget

| Name | Stored as |
| --- | --- |
| `Open in same window` | `_self` |
| `Open in new window` | `_blank` |
| `Delegate to parent` | `_parent` |
| `Delegate to top` | `_top` |

## DjangocmsFileFolderTemplate

| Name | Stored as |
| --- | --- |
| `Default` | `default` |

## DjangocmsFileFolderLinkTarget

| Name | Stored as |
| --- | --- |
| `Open in same window` | `_self` |
| `Open in new window` | `_blank` |
| `Delegate to parent` | `_parent` |
| `Delegate to top` | `_top` |

## VideoPlayerTemplate

| Name | Stored as |
| --- | --- |
| `Default` | `default` |

## VideoTrackKind

| Name | Stored as |
| --- | --- |
| `Subtitles` | `subtitles` |
| `Captions` | `captions` |
| `Descriptions` | `descriptions` |
| `Chapters` | `chapters` |
