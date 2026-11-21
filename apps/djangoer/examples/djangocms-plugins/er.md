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
    cms_pagecontent ||..o{ djangocms_versioning_version : "(id) - (object_id)"
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
    auth_user ||--o{ djangocms_versioning_version : "(id) - (created_by_id)"
    django_content_type ||--o{ djangocms_versioning_version : "(id) - (content_type_id)"
    auth_user |o--o{ djangocms_versioning_version : "(id) - (locked_by_id)"
    djangocms_versioning_version |o--o{ djangocms_versioning_version : "(id) - (source_id)"
    djangocms_versioning_version ||--o{ djangocms_versioning_statetracking : "(id) - (version_id)"
    auth_user ||--o{ djangocms_versioning_statetracking : "(id) - (user_id)"
    djangocms_alias_category |o--o{ djangocms_alias_category_translation : "(id) - (master_id)"
    djangocms_alias_category ||--o{ djangocms_alias_alias : "(id) - (category_id)"
    django_site |o--o{ djangocms_alias_alias : "(id) - (site_id)"
    djangocms_alias_alias ||--o{ djangocms_alias_aliascontent : "(id) - (alias_id)"
    djangocms_alias_aliascontent ||..o{ cms_placeholder : "(id) - (object_id)"
    djangocms_alias_aliascontent ||..o{ djangocms_versioning_version : "(id) - (object_id)"
    cms_cmsplugin ||--o| djangocms_alias_aliasplugin : "is a, by cmsplugin_ptr_id"
    djangocms_alias_alias ||--o{ djangocms_alias_aliasplugin : "(id) - (alias_id)"
    djangocms_snippet_snippetgrouper ||--o{ djangocms_snippet_snippet : "(id) - (snippet_grouper_id)"
    django_site |o--o{ djangocms_snippet_snippet : "(id) - (site_id)"
    djangocms_snippet_snippet ||..o{ djangocms_versioning_version : "(id) - (object_id)"
    cms_cmsplugin ||--o| djangocms_snippet_snippetptr : "is a, by cmsplugin_ptr_id"
    djangocms_snippet_snippetgrouper ||--o{ djangocms_snippet_snippetptr : "(id) - (snippet_grouper_id)"
    filer_image |o--o{ djangocms_picture_picture : "(file_ptr_id) - (picture_id)"
    cms_page |o--o{ djangocms_picture_picture : "(id) - (link_page_id)"
    filer_thumbnailoption |o--o{ djangocms_picture_picture : "(id) - (thumbnail_options_id)"
    cms_cmsplugin ||--o| djangocms_picture_picture : "is a, by cmsplugin_ptr_id"
    filer_file |o--o{ djangocms_file_file : "(id) - (file_src_id)"
    cms_cmsplugin ||--o| djangocms_file_file : "is a, by cmsplugin_ptr_id"
    filer_folder |o--o{ djangocms_file_folder : "(id) - (folder_src_id)"
    cms_cmsplugin ||--o| djangocms_file_folder : "is a, by cmsplugin_ptr_id"
    filer_image |o--o{ djangocms_video_videoplayer : "(file_ptr_id) - (poster_id)"
    cms_cmsplugin ||--o| djangocms_video_videoplayer : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| djangocms_video_videosource : "is a, by cmsplugin_ptr_id"
    filer_file |o--o{ djangocms_video_videosource : "(id) - (source_file_id)"
    cms_cmsplugin ||--o| djangocms_video_hlsstreamsource : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| djangocms_video_videotrack : "is a, by cmsplugin_ptr_id"
    filer_file |o--o{ djangocms_video_videotrack : "(id) - (src_id)"
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
    djangocms_versioning_version {
        AutoField id PK
        DateTimeField created
        DateTimeField modified
        ForeignKey created_by_id FK
        CharField number
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
        VersionState state
        ForeignKey locked_by_id FK
        ForeignKey source_id FK
    }
    djangocms_versioning_statetracking {
        AutoField id PK
        ForeignKey version_id FK
        DateTimeField date
        StateTrackingOldState old_state
        StateTrackingNewState new_state
        ForeignKey user_id FK
    }
    djangocms_alias_category_translation {
        BigAutoField id PK
        CategoryTranslationLanguageCode language_code
        CharField name
        TranslationsForeignKey master_id FK
    }
    djangocms_alias_category {
        BigAutoField id PK
    }
    djangocms_alias_alias {
        BigAutoField id PK
        AliasCreationMethod creation_method
        ForeignKey category_id FK
        PositiveIntegerField position
        CharField static_code "To render the alias in templates."
        ForeignKey site_id FK
    }
    djangocms_alias_aliascontent {
        BigAutoField id PK
        ForeignKey alias_id FK
        CharField name
        CharField language
    }
    djangocms_alias_aliasplugin {
        OneToOneField cmsplugin_ptr_id PK
        ForeignKey alias_id FK
        AliasPluginTemplate template
    }
    djangocms_snippet_snippetgrouper {
        BigAutoField id PK
    }
    djangocms_snippet_snippet {
        BigAutoField id PK
        CharField name
        ForeignKey snippet_grouper_id FK
        TextField html
        CharField template "Enter a template (e.g. 'snippets/plugin_xy.html') to be rendered. If 'template' is given, the contents of field 'HTML' will be passed as template variable {{ html }} to the template. Otherwise, the content of 'HTML' is rendered."
        SlugField slug
        ForeignKey site_id FK
    }
    djangocms_snippet_snippetptr {
        OneToOneField cmsplugin_ptr_id PK
        ForeignKey snippet_grouper_id FK
    }
    djangocms_picture_picture {
        PictureTemplate template
        FilerImageField picture_id FK
        URLField external_picture "If provided, overrides the embedded image. Certain options such as cropping are not applicable to external images."
        PositiveIntegerField width "The image width as number in pixels. Example: '720' and not '720px'."
        PositiveIntegerField height "The image height as number in pixels. Example: '720' and not '720px'."
        PictureAlignment alignment "Aligns the image according to the selected option."
        TextField caption_text "Provide a description, attribution, copyright or other information."
        AttributesField attributes
        URLField link_url "Wraps the image in a link to an external URL."
        PageField link_page_id FK "Wraps the image in a link to an internal (page) URL."
        PictureLinkTarget link_target
        AttributesField link_attributes
        BooleanField use_automatic_scaling "Uses the placeholder dimensions to automatically calculate the size."
        BooleanField use_no_cropping "Outputs the raw image without cropping."
        BooleanField use_crop "Crops the image according to the thumbnail settings provided in the template."
        BooleanField use_upscale "Upscales the image to the size of the thumbnail settings in the template."
        PictureUseResponsiveImage use_responsive_image "Uses responsive image technique to choose better image to display based upon screen viewport. This configuration only applies to uploaded images (external pictures will not be affected)."
        ForeignKey thumbnail_options_id FK "Overrides width, height, and crop; scales up to the provided preset dimensions."
        OneToOneField cmsplugin_ptr_id PK
    }
    djangocms_file_file {
        DjangocmsFileFileTemplate template
        FilerFileField file_src_id FK
        CharField file_name "Overrides the default file name with the given value."
        DjangocmsFileFileLinkTarget link_target
        CharField link_title
        BooleanField show_file_size "Appends the file size at the end of the name."
        AttributesField attributes
        OneToOneField cmsplugin_ptr_id PK
    }
    djangocms_file_folder {
        DjangocmsFileFolderTemplate template
        FilerFolderField folder_src_id FK
        DjangocmsFileFolderLinkTarget link_target
        BooleanField show_file_size "Appends the file size at the end of the name."
        AttributesField attributes
        OneToOneField cmsplugin_ptr_id PK
    }
    djangocms_video_videoplayer {
        VideoPlayerTemplate template
        CharField label
        TextField embed_link "Use this field to embed videos from external services such as YouTube, Vimeo or others. Leave it blank to upload video files by adding nested 'Source' plugins."
        AttributesField parameters "Parameters are appended to the video link if provided."
        FilerImageField poster_id FK
        AttributesField attributes
        BooleanField show_controls "If enabled, the video will be shown with Play, Pause and Seek elements that allow the user to control playback."
        BooleanField autoplay "If enabled, the video will automatically play once the page is loaded. This might not work depending on how the user has configured their browser."
        OneToOneField cmsplugin_ptr_id PK
    }
    djangocms_video_videosource {
        OneToOneField cmsplugin_ptr_id PK
        FilerFileField source_file_id FK
        CharField text_title
        TextField text_description
        AttributesField attributes
    }
    djangocms_video_hlsstreamsource {
        OneToOneField cmsplugin_ptr_id PK
        URLField hls_source_url
    }
    djangocms_video_videotrack {
        OneToOneField cmsplugin_ptr_id PK
        VideoTrackKind kind
        FilerFileField src_id FK
        CharField srclang "Examples: 'en' or 'de' etc."
        CharField label
        AttributesField attributes
    }
```
