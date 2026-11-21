```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user ||--o| cms_usersettings : "(id) - (user_id)"
    cms_placeholder |o--o{ cms_usersettings : "(id) - (clipboard_id)"
    cms_page |o--o{ cms_page : "(id) - (parent_id)"
    django_site ||--o{ cms_page : "(id) - (site_id)"
    cms_page ||--o{ cms_pageurl : "(id) - (page_id)"
    django_site ||--o{ cms_pageurl : "(id) - (site_id)"
    auth_user |o--o{ cms_globalpagepermission : "(id) - (user_id)"
    auth_group |o--o{ cms_globalpagepermission : "(id) - (group_id)"
    cms_globalpagepermission }o--o{ django_site : "(sites) - (globalpagepermission_set)"
    auth_user |o--o{ cms_pagepermission : "(id) - (user_id)"
    auth_group |o--o{ cms_pagepermission : "(id) - (group_id)"
    cms_page |o--o{ cms_pagepermission : "(id) - (page_id)"
    auth_user ||--o| cms_pageuser : "is a, by user_ptr_id"
    auth_user |o--o{ cms_pageuser : "(id) - (created_by_id)"
    auth_group ||--o| cms_pageusergroup : "is a, by group_ptr_id"
    auth_user |o--o{ cms_pageusergroup : "(id) - (created_by_id)"
    django_content_type |o--o{ cms_placeholder : "(id) - (content_type_id)"
    cms_placeholder |o--o{ cms_cmsplugin : "(id) - (placeholder_id)"
    cms_cmsplugin |o--o{ cms_cmsplugin : "(id) - (parent_id)"
    cms_page ||--o{ cms_pagecontent : "(id) - (page_id)"
    cms_pagecontent ||..o{ cms_placeholder : "(id) - (object_id)"
    cms_cmsplugin ||--o| cms_placeholderreference : "is a, by cmsplugin_ptr_id"
    cms_placeholderreference ||..o{ cms_placeholder : "(cmsplugin_ptr_id) - (object_id)"
    cms_cmsplugin ||--o| cms_aliaspluginmodel : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin |o--o{ cms_aliaspluginmodel : "(id) - (plugin_id)"
    cms_placeholder |o--o{ cms_aliaspluginmodel : "(id) - (alias_placeholder_id)"
    easy_thumbnails_source ||--o{ easy_thumbnails_thumbnail : "(id) - (source_id)"
    easy_thumbnails_thumbnail ||--o| easy_thumbnails_thumbnaildimensions : "(id) - (thumbnail_id)"
    filer_folder |o--o{ filer_folder : "(id) - (parent_id)"
    auth_user |o--o{ filer_folder : "(id) - (owner_id)"
    filer_folder |o--o{ filer_folderpermission : "(id) - (folder_id)"
    auth_user |o--o{ filer_folderpermission : "(id) - (user_id)"
    auth_group |o--o{ filer_folderpermission : "(id) - (group_id)"
    django_content_type |o--o{ filer_file : "(id) - (polymorphic_ctype_id)"
    filer_folder |o--o{ filer_file : "(id) - (folder_id)"
    auth_user |o--o{ filer_file : "(id) - (owner_id)"
    auth_user ||--o{ filer_clipboard : "(id) - (user_id)"
    filer_file ||--o{ filer_clipboarditem : "(id) - (file_id)"
    filer_clipboard ||--o{ filer_clipboarditem : "(id) - (clipboard_id)"
    filer_file ||--o| filer_image : "is a, by file_ptr_id"
    cms_cmsplugin ||--o| djangocms_text_text : "is a, by cmsplugin_ptr_id"
    taggit_tag ||--o{ taggit_taggeditem : "(id) - (tag_id)"
    django_content_type ||--o{ taggit_taggeditem : "(id) - (content_type_id)"
    djangocms_stories_storiesconfig |o--o{ djangocms_stories_storiesconfig_translation : "(id) - (master_id)"
    filer_thumbnailoption |o--o{ djangocms_stories_storiesconfig : "(id) - (default_image_full_id)"
    filer_thumbnailoption |o--o{ djangocms_stories_storiesconfig : "(id) - (default_image_thumbnail_id)"
    djangocms_stories_postcategory |o--o{ djangocms_stories_postcategory_translation : "(id) - (master_id)"
    djangocms_stories_postcategory |o--o{ djangocms_stories_postcategory : "(id) - (parent_id)"
    djangocms_stories_storiesconfig |o--o{ djangocms_stories_postcategory : "(id) - (app_config_id)"
    filer_image |o--o{ djangocms_stories_postcategory : "(file_ptr_id) - (main_image_id)"
    filer_thumbnailoption |o--o{ djangocms_stories_postcategory : "(id) - (main_image_thumbnail_id)"
    filer_thumbnailoption |o--o{ djangocms_stories_postcategory : "(id) - (main_image_full_id)"
    auth_user |o--o{ djangocms_stories_post : "(id) - (author_id)"
    filer_image |o--o{ djangocms_stories_post : "(file_ptr_id) - (main_image_id)"
    filer_thumbnailoption |o--o{ djangocms_stories_post : "(id) - (main_image_thumbnail_id)"
    filer_thumbnailoption |o--o{ djangocms_stories_post : "(id) - (main_image_full_id)"
    djangocms_stories_storiesconfig |o--o{ djangocms_stories_post : "(id) - (app_config_id)"
    djangocms_stories_post }o--o{ djangocms_stories_postcategory : "(categories) - (posts)"
    djangocms_stories_post }o--o{ django_site : "(sites) - (post_set)"
    djangocms_stories_post }o--o{ djangocms_stories_post : "(related) - (post_set)"
    djangocms_stories_post ||..o{ taggit_taggeditem : "(id) - (object_id)"
    djangocms_stories_post ||--o{ djangocms_stories_postcontent : "(id) - (post_id)"
    djangocms_stories_postcontent ||..o{ cms_placeholder : "(id) - (object_id)"
    cms_cmsplugin ||--o| djangocms_stories_latestpostsplugin : "is a, by cmsplugin_ptr_id"
    djangocms_stories_storiesconfig |o--o{ djangocms_stories_latestpostsplugin : "(id) - (app_config_id)"
    djangocms_stories_latestpostsplugin }o--o{ djangocms_stories_postcategory : "(categories) - (latestpostsplugin_set)"
    djangocms_stories_latestpostsplugin ||..o{ taggit_taggeditem : "(cmsplugin_ptr_id) - (object_id)"
    cms_cmsplugin ||--o| djangocms_stories_authorentriesplugin : "is a, by cmsplugin_ptr_id"
    djangocms_stories_storiesconfig |o--o{ djangocms_stories_authorentriesplugin : "(id) - (app_config_id)"
    djangocms_stories_authorentriesplugin }o--o{ auth_user : "(authors) - (authorentriesplugin_set)"
    cms_cmsplugin ||--o| djangocms_stories_featuredpostsplugin : "is a, by cmsplugin_ptr_id"
    djangocms_stories_storiesconfig |o--o{ djangocms_stories_featuredpostsplugin : "(id) - (app_config_id)"
    djangocms_stories_featuredpostsplugin }o--o{ djangocms_stories_post : "(posts) - (featuredpostsplugin_set)"
    cms_cmsplugin ||--o| djangocms_stories_genericblogplugin : "is a, by cmsplugin_ptr_id"
    djangocms_stories_storiesconfig |o--o{ djangocms_stories_genericblogplugin : "(id) - (app_config_id)"
    django_admin_log {
        AutoField id PK
        DateTimeField action_time
        ForeignKey user_id FK
        ForeignKey content_type_id FK
        TextField object_id
        CharField object_repr
        LogEntryActionFlag action_flag
        TextField change_message
    }
    auth_permission {
        AutoField id PK
        CharField name
        ForeignKey content_type_id FK
        CharField codename
    }
    auth_group {
        AutoField id PK
        CharField name UK
    }
    auth_user {
        AutoField id PK
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        CharField username UK "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
        CharField first_name
        CharField last_name
        EmailField email
        BooleanField is_staff "Designates whether the user can log into this admin site."
        BooleanField is_active "Designates whether this user should be treated as active. Unselect this instead of deleting accounts."
        DateTimeField date_joined
    }
    django_content_type {
        AutoField id PK
        CharField app_label
        CharField model
    }
    django_session {
        CharField session_key PK
        TextField session_data
        DateTimeField expire_date
    }
    django_site {
        AutoField id PK
        CharField domain UK
        CharField name
    }
    cms_usersettings {
        AutoField id PK
        OneToOneField user_id FK, UK
        UserSettingsLanguage language "The language for the admin interface and toolbar"
        PlaceholderForeignKey clipboard_id FK
    }
    cms_page {
        AutoField id PK
        CharField path UK
        PositiveIntegerField depth
        PositiveIntegerField numchild
        ForeignKey parent_id FK
        ForeignKey site_id FK
        CharField created_by
        CharField changed_by
        DateTimeField creation_date
        DateTimeField changed_date
        CharField reverse_id "A unique identifier that is used with the page_url templatetag for linking to this page"
        CharField navigation_extenders
        BooleanField login_required
        BooleanField is_home
        CharField application_urls
        CharField application_namespace
        BooleanField is_page_type "Mark this page as a page type"
    }
    cms_pageurl {
        AutoField id PK
        SlugField slug
        CharField path
        CharField language
        ForeignKey page_id FK
        ForeignKey site_id FK
        BooleanField managed
    }
    cms_globalpagepermission {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey group_id FK
        BooleanField can_change
        BooleanField can_add
        BooleanField can_delete
        BooleanField can_publish
        BooleanField can_change_advanced_settings
        BooleanField can_change_permissions "Allows this user to grant and revoke page permissions for other users."
        BooleanField can_move_page
        BooleanField can_view "Grants frontend view access. Note: as soon as any user or group is given view access to a page, that page becomes restricted — only users/groups with view access can then see it on the frontend."
        BooleanField can_recover_page "can recover any deleted page"
    }
    cms_pagepermission {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey group_id FK
        BooleanField can_change
        BooleanField can_add
        BooleanField can_delete
        BooleanField can_publish
        BooleanField can_change_advanced_settings
        BooleanField can_change_permissions "Allows this user to grant and revoke page permissions for other users."
        BooleanField can_move_page
        BooleanField can_view "Grants frontend view access. Note: as soon as any user or group is given view access to a page, that page becomes restricted — only users/groups with view access can then see it on the frontend."
        PagePermissionGrantOn grant_on
        ForeignKey page_id FK
    }
    cms_pageuser {
        OneToOneField user_ptr_id PK
        ForeignKey created_by_id FK
    }
    cms_pageusergroup {
        OneToOneField group_ptr_id PK
        ForeignKey created_by_id FK
    }
    cms_placeholder {
        AutoField id PK
        CharField slot
        PositiveSmallIntegerField default_width
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
    }
    cms_cmsplugin {
        AutoField id PK
        ForeignKey placeholder_id FK
        ForeignKey parent_id FK
        SmallIntegerField position
        CharField language
        CharField plugin_type
        DateTimeField creation_date
        DateTimeField changed_date
    }
    cms_pagecontent {
        AutoField id PK
        CharField language
        CharField title "The default title"
        CharField page_title "Overwrites what is displayed at the top of your browser or in bookmarks"
        CharField menu_title "Overwrite what is displayed in the menu"
        TextField meta_description "A description of the page used by search engines."
        CharField redirect "Redirects to this URL."
        SlugField slug "The part of the title that is used in the URL"
        CharField overwrite_url "Keep this field empty if standard path should be used."
        ForeignKey page_id FK
        DateTimeField creation_date
        CharField created_by
        CharField changed_by
        DateTimeField changed_date
        BooleanField in_navigation
        BooleanField soft_root "All ancestors will not be displayed in the navigation"
        CharField template "The template used to render the content."
        PageContentLimitVisibilityInMenu limit_visibility_in_menu "limit when this page is visible in the menu"
        PageContentXframeOptions xframe_options
    }
    cms_placeholderreference {
        OneToOneField cmsplugin_ptr_id PK
        CharField name
    }
    cms_aliaspluginmodel {
        OneToOneField cmsplugin_ptr_id PK
        ForeignKey plugin_id FK
        ForeignKey alias_placeholder_id FK
    }
    cms_urlconfrevision {
        AutoField id PK
        CharField revision
    }
    menus_cachekey {
        AutoField id PK
        CharField language
        PositiveIntegerField site
        CharField key
    }
    easy_thumbnails_source {
        AutoField id PK
        CharField storage_hash
        CharField name
        DateTimeField modified
    }
    easy_thumbnails_thumbnail {
        AutoField id PK
        CharField storage_hash
        CharField name
        DateTimeField modified
        ForeignKey source_id FK
    }
    easy_thumbnails_thumbnaildimensions {
        AutoField id PK
        OneToOneField thumbnail_id FK, UK
        PositiveIntegerField width
        PositiveIntegerField height
    }
    filer_folder {
        AutoField id PK
        ForeignKey parent_id FK
        CharField name
        ForeignKey owner_id FK
        DateTimeField uploaded_at
        DateTimeField created_at
        DateTimeField modified_at
    }
    filer_folderpermission {
        AutoField id PK
        ForeignKey folder_id FK
        FolderPermissionType type
        ForeignKey user_id FK
        ForeignKey group_id FK
        BooleanField everybody
        FolderPermissionCanRead can_read
        FolderPermissionCanEdit can_edit
        FolderPermissionCanAddChildren can_add_children
    }
    filer_file {
        AutoField id PK
        ForeignKey polymorphic_ctype_id FK
        ForeignKey folder_id FK
        MultiStorageFileField file
        BigIntegerField _file_size
        CharField sha1
        BooleanField has_all_mandatory_data
        CharField original_filename
        CharField name
        TextField description
        ForeignKey owner_id FK
        DateTimeField uploaded_at
        DateTimeField modified_at
        BooleanField is_public "Disable any permission checking for this file. File will be publicly accessible to anyone."
        CharField mime_type "MIME type of uploaded content"
    }
    filer_clipboard {
        AutoField id PK
        ForeignKey user_id FK
    }
    filer_clipboarditem {
        AutoField id PK
        ForeignKey file_id FK
        ForeignKey clipboard_id FK
    }
    filer_image {
        FloatField _height
        FloatField _width
        BooleanField _transparent
        CharField default_alt_text
        CharField default_caption
        CharField subject_location
        OneToOneField file_ptr_id PK
        DateTimeField date_taken
        CharField author
        BooleanField must_always_publish_author_credit
        BooleanField must_always_publish_copyright
    }
    filer_thumbnailoption {
        AutoField id PK
        CharField name
        IntegerField width "width in pixel."
        IntegerField height "height in pixel."
        BooleanField crop
        BooleanField upscale
    }
    djangocms_text_text {
        OneToOneField cmsplugin_ptr_id PK
        TextField body
        JSONField json
        CharField rte "The rich text editor used to create this text. JSON formats vary between editors."
    }
    taggit_tag {
        AutoField id PK
        CharField name UK
        SlugField slug UK
    }
    taggit_taggeditem {
        AutoField id PK
        ForeignKey tag_id FK
        ForeignKey content_type_id FK
        IntegerField object_id
    }
    djangocms_stories_storiesconfig_translation {
        BigAutoField id PK
        StoriesConfigTranslationLanguageCode language_code
        CharField app_title
        CharField object_name
        TranslationsForeignKey master_id FK
    }
    djangocms_stories_storiesconfig {
        BigAutoField id PK
        CharField type
        CharField namespace UK
        ForeignKey default_image_full_id FK "If left empty the image size will have to be set for every newly created post."
        ForeignKey default_image_thumbnail_id FK "If left empty the thumbnail image size will have to be set for every newly created post."
        CharField url_patterns
        BooleanField use_placeholder
        BooleanField use_abstract
        StoriesConfigUseRelated use_related
        BooleanField set_author
        SmallIntegerField paginate_by "When paginating list views, how many articles per page?"
        CharField template_prefix "Alternative directory to load the stories templates from"
        CharField menu_structure "Structure of the django CMS menu"
        BooleanField menu_empty_categories "Show categories with no post attached in the menu"
        CharField sitemap_changefreq "Changefreq attribute for sitemap items"
        DecimalField sitemap_priority "Priority attribute for sitemap items"
        CharField object_type
        CharField og_type
        CharField og_app_id
        CharField og_profile_id
        CharField og_publisher
        CharField og_author_url
        CharField og_author
        CharField twitter_type
        CharField twitter_site
        CharField twitter_author
        CharField gplus_type
        CharField gplus_author
        BooleanField send_knock_create "Emits a desktop notification -if enabled- when publishing a new post"
        BooleanField send_knock_update "Emits a desktop notification -if enabled- when editing a published post"
    }
    djangocms_stories_postcategory_translation {
        BigAutoField id PK
        PostCategoryTranslationLanguageCode language_code
        CharField name
        SlugField slug
        TextField meta_description
        HTMLField abstract
        TranslationsForeignKey master_id FK
    }
    djangocms_stories_postcategory {
        BigAutoField id PK
        ForeignKey parent_id FK
        DateTimeField date_created
        DateTimeField date_modified
        ForeignKey app_config_id FK "When selecting a value, the form is reloaded to get the updated default"
        IntegerField priority
        FilerImageField main_image_id FK
        ForeignKey main_image_thumbnail_id FK
        ForeignKey main_image_full_id FK
    }
    djangocms_stories_post {
        BigAutoField id PK
        ForeignKey author_id FK
        DateTimeField date_created
        DateTimeField date_modified
        DateTimeField date_published
        DateTimeField date_published_end
        DateTimeField date_featured
        BooleanField include_in_rss
        FilerImageField main_image_id FK
        ForeignKey main_image_thumbnail_id FK
        ForeignKey main_image_full_id FK
        BooleanField enable_comments
        ForeignKey app_config_id FK "When selecting a value, the form is reloaded to get the updated default"
    }
    djangocms_stories_postcontent {
        BigAutoField id PK
        ForeignKey post_id FK
        CharField language
        CharField title
        SlugField slug
        CharField subtitle
        HTMLField abstract
        TextField meta_description
        TextField meta_keywords
        CharField meta_title "used in title tag and social sharing"
        HTMLField post_text
    }
    djangocms_stories_latestpostsplugin {
        OneToOneField cmsplugin_ptr_id PK
        ForeignKey app_config_id FK "When selecting a value, the form is reloaded to get the updated default"
        BooleanField current_site "Select items from the current site only"
        LatestPostsPluginTemplateFolder template_folder "Select plugin layout to load for this instance"
        IntegerField latest_posts "The number of latests entries to be displayed."
    }
    djangocms_stories_authorentriesplugin {
        OneToOneField cmsplugin_ptr_id PK
        ForeignKey app_config_id FK "When selecting a value, the form is reloaded to get the updated default"
        BooleanField current_site "Select items from the current site only"
        AuthorEntriesPluginTemplateFolder template_folder "Select plugin layout to load for this instance"
        IntegerField latest_posts "The number of author entries to be displayed."
    }
    djangocms_stories_featuredpostsplugin {
        OneToOneField cmsplugin_ptr_id PK
        ForeignKey app_config_id FK "When selecting a value, the form is reloaded to get the updated default"
        BooleanField current_site "Select items from the current site only"
        FeaturedPostsPluginTemplateFolder template_folder "Select plugin layout to load for this instance"
    }
    djangocms_stories_genericblogplugin {
        OneToOneField cmsplugin_ptr_id PK
        ForeignKey app_config_id FK "When selecting a value, the form is reloaded to get the updated default"
        BooleanField current_site "Select items from the current site only"
        GenericBlogPluginTemplateFolder template_folder "Select plugin layout to load for this instance"
    }
```
