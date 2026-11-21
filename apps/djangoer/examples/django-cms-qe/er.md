```mermaid
erDiagram
    cms_qe_auth_user }o--o{ auth_group : "(groups) - (user_set)"
    cms_qe_auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    cms_cmsplugin ||--o| cms_qe_menu_menupluginmodel : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| cms_qe_newsletter_newsletterpluginmodel : "is a, by cmsplugin_ptr_id"
    cms_qe_newsletter_newsletterpluginmodel }o--o{ cms_qe_newsletter_mailinglist : "(mailing_lists) - (newsletterpluginmodel_set)"
    cms_qe_newsletter_mailinglist ||--o{ cms_qe_newsletter_subscriber : "(id) - (mailing_list_id)"
    cms_qe_newsletter_mailinglist |o--o{ cms_qe_newsletter_subscribetask : "(id) - (mailing_list_id)"
    cms_cmsplugin ||--o| cms_qe_table_tablepluginmodel : "is a, by cmsplugin_ptr_id"
    filer_image |o--o{ cms_qe_video_sourcefilevideoplayer : "(file_ptr_id) - (poster_id)"
    cms_cmsplugin ||--o| cms_qe_video_sourcefilevideoplayer : "is a, by cmsplugin_ptr_id"
    filer_file |o--o{ cms_qe_video_sourcefilevideoplayer : "(id) - (source_file_id)"
    filer_image |o--o{ cms_qe_video_hostingvideoplayer : "(file_ptr_id) - (poster_id)"
    cms_cmsplugin ||--o| cms_qe_video_hostingvideoplayer : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| cms_qe_video_videotrack : "is a, by cmsplugin_ptr_id"
    filer_file |o--o{ cms_qe_video_videotrack : "(id) - (src_id)"
    cms_cmsplugin ||--o| cms_qe_plugins_script : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| cms_qe_plugins_link : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| cms_qe_plugins_iframe : "is a, by cmsplugin_ptr_id"
    cms_qe_auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    cms_qe_auth_user ||--o| cms_usersettings : "(id) - (user_id)"
    cms_placeholder |o--o{ cms_usersettings : "(id) - (clipboard_id)"
    cms_page |o--o{ cms_page : "(id) - (parent_id)"
    django_site ||--o{ cms_page : "(id) - (site_id)"
    cms_page ||--o{ cms_pageurl : "(id) - (page_id)"
    django_site ||--o{ cms_pageurl : "(id) - (site_id)"
    cms_qe_auth_user |o--o{ cms_globalpagepermission : "(id) - (user_id)"
    auth_group |o--o{ cms_globalpagepermission : "(id) - (group_id)"
    cms_globalpagepermission }o--o{ django_site : "(sites) - (globalpagepermission_set)"
    cms_qe_auth_user |o--o{ cms_pagepermission : "(id) - (user_id)"
    auth_group |o--o{ cms_pagepermission : "(id) - (group_id)"
    cms_page |o--o{ cms_pagepermission : "(id) - (page_id)"
    cms_qe_auth_user ||--o| cms_pageuser : "is a, by user_ptr_id"
    cms_qe_auth_user |o--o{ cms_pageuser : "(id) - (created_by_id)"
    auth_group ||--o| cms_pageusergroup : "is a, by group_ptr_id"
    cms_qe_auth_user |o--o{ cms_pageusergroup : "(id) - (created_by_id)"
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
    cms_cmsplugin ||--o| djangocms_text_text : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| djangocms_link_link : "is a, by cmsplugin_ptr_id"
    djangocms_alias_category |o--o{ djangocms_alias_category_translation : "(id) - (master_id)"
    djangocms_alias_category ||--o{ djangocms_alias_alias : "(id) - (category_id)"
    django_site |o--o{ djangocms_alias_alias : "(id) - (site_id)"
    djangocms_alias_alias ||--o{ djangocms_alias_aliascontent : "(id) - (alias_id)"
    djangocms_alias_aliascontent ||..o{ cms_placeholder : "(id) - (object_id)"
    djangocms_alias_aliascontent ||..o{ djangocms_versioning_version : "(id) - (object_id)"
    cms_cmsplugin ||--o| djangocms_alias_aliasplugin : "is a, by cmsplugin_ptr_id"
    djangocms_alias_alias ||--o{ djangocms_alias_aliasplugin : "(id) - (alias_id)"
    cms_qe_auth_user ||--o{ djangocms_versioning_version : "(id) - (created_by_id)"
    django_content_type ||--o{ djangocms_versioning_version : "(id) - (content_type_id)"
    cms_qe_auth_user |o--o{ djangocms_versioning_version : "(id) - (locked_by_id)"
    djangocms_versioning_version |o--o{ djangocms_versioning_version : "(id) - (source_id)"
    djangocms_versioning_version ||--o{ djangocms_versioning_statetracking : "(id) - (version_id)"
    cms_qe_auth_user ||--o{ djangocms_versioning_statetracking : "(id) - (user_id)"
    cms_cmsplugin ||--o| djangocms_googlemap_googlemap : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| djangocms_googlemap_googlemapmarker : "is a, by cmsplugin_ptr_id"
    filer_image |o--o{ djangocms_googlemap_googlemapmarker : "(file_ptr_id) - (icon_id)"
    cms_cmsplugin ||--o| djangocms_googlemap_googlemaproute : "is a, by cmsplugin_ptr_id"
    filer_folder |o--o{ filer_folder : "(id) - (parent_id)"
    cms_qe_auth_user |o--o{ filer_folder : "(id) - (owner_id)"
    filer_folder |o--o{ filer_folderpermission : "(id) - (folder_id)"
    cms_qe_auth_user |o--o{ filer_folderpermission : "(id) - (user_id)"
    auth_group |o--o{ filer_folderpermission : "(id) - (group_id)"
    django_content_type |o--o{ filer_file : "(id) - (polymorphic_ctype_id)"
    filer_folder |o--o{ filer_file : "(id) - (folder_id)"
    cms_qe_auth_user |o--o{ filer_file : "(id) - (owner_id)"
    cms_qe_auth_user ||--o{ filer_clipboard : "(id) - (user_id)"
    filer_file ||--o{ filer_clipboarditem : "(id) - (file_id)"
    filer_clipboard ||--o{ filer_clipboarditem : "(id) - (clipboard_id)"
    filer_file ||--o| filer_image : "is a, by file_ptr_id"
    easy_thumbnails_source ||--o{ easy_thumbnails_thumbnail : "(id) - (source_id)"
    easy_thumbnails_thumbnail ||--o| easy_thumbnails_thumbnaildimensions : "(id) - (thumbnail_id)"
    filer_file |o--o{ djangocms_file_file : "(id) - (file_src_id)"
    cms_cmsplugin ||--o| djangocms_file_file : "is a, by cmsplugin_ptr_id"
    filer_folder |o--o{ djangocms_file_folder : "(id) - (folder_src_id)"
    cms_cmsplugin ||--o| djangocms_file_folder : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| djangocms_icon_icon : "is a, by cmsplugin_ptr_id"
    filer_image |o--o{ djangocms_picture_picture : "(file_ptr_id) - (picture_id)"
    cms_page |o--o{ djangocms_picture_picture : "(id) - (link_page_id)"
    filer_thumbnailoption |o--o{ djangocms_picture_picture : "(id) - (thumbnail_options_id)"
    cms_cmsplugin ||--o| djangocms_picture_picture : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| djangocms_frontend_frontenduiitem : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| djangocms_render_context_rendercontext : "is a, by cmsplugin_ptr_id"
    filer_file |o--o{ djangocms_render_context_rendercontext : "(id) - (file_id)"
    cms_cmsplugin ||--o| djangocms_render_context_default : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| djangocms_render_context_stack : "is a, by cmsplugin_ptr_id"
    axes_accessattempt ||--o| axes_accessattemptexpiration : "(id) - (access_attempt_id)"
    cms_cmsplugin ||--o| aldryn_forms_formplugin : "is a, by cmsplugin_ptr_id"
    aldryn_forms_formplugin }o--o{ cms_qe_auth_user : "(recipients) - (formplugin_set)"
    aldryn_forms_formplugin }o--o{ aldryn_forms_webhook : "(webhooks) - (formplugin_set)"
    cms_cmsplugin ||--o| aldryn_forms_fieldsetplugin : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| aldryn_forms_fieldplugin : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| aldryn_forms_textareafieldplugin : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| aldryn_forms_emailfieldplugin : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| aldryn_forms_fileuploadfieldplugin : "is a, by cmsplugin_ptr_id"
    filer_folder ||--o{ aldryn_forms_fileuploadfieldplugin : "(id) - (upload_to_id)"
    cms_cmsplugin ||--o| aldryn_forms_multiplefilesuploadfieldplugin : "is a, by cmsplugin_ptr_id"
    filer_folder ||--o{ aldryn_forms_multiplefilesuploadfieldplugin : "(id) - (upload_to_id)"
    cms_cmsplugin ||--o| aldryn_forms_imageuploadfieldplugin : "is a, by cmsplugin_ptr_id"
    filer_folder ||--o{ aldryn_forms_imageuploadfieldplugin : "(id) - (upload_to_id)"
    cms_cmsplugin ||--o| aldryn_forms_datefieldplugin : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| aldryn_forms_datetimelocalfieldplugin : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| aldryn_forms_timefieldplugin : "is a, by cmsplugin_ptr_id"
    cms_cmsplugin ||--o| aldryn_forms_urlfieldplugin : "is a, by cmsplugin_ptr_id"
    aldryn_forms_fieldplugin ||--o{ aldryn_forms_option : "(cmsplugin_ptr_id) - (field_id)"
    cms_cmsplugin ||--o| aldryn_forms_formbuttonplugin : "is a, by cmsplugin_ptr_id"
    aldryn_forms_formsubmission }o--o{ aldryn_forms_webhook : "(webhooks) - (formsubmission_set)"
    aldryn_forms_submittedtobesent }o--o{ aldryn_forms_webhook : "(webhooks) - (submittedtobesent_set)"
    cms_qe_auth_user |o--o{ email_notifications_emailnotification : "(id) - (to_user_id)"
    cms_qe_auth_user ||--o{ knox_authtoken : "(id) - (user_id)"
    cms_qe_auth_user {
        BigAutoField id PK
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
    cms_qe_menu_menupluginmodel {
        OneToOneField cmsplugin_ptr_id PK
        MenuPluginModelOrientation orientation
        CharField start_page "ID of page in advanced settings."
        IntegerField end_level "Specify at which level the navigation should stop.              Keep in mind that Bootstrap by default does not support dropdown submenu.             Any deep dropdown menu is not very well accessible from mobile devices."
        BooleanField submenu_in_inactive_item
    }
    cms_qe_newsletter_newsletterpluginmodel {
        OneToOneField cmsplugin_ptr_id PK
        CharField title "The title of subscribes that will be shown to users"
        BooleanField fullname_show
        BooleanField fullname_require
    }
    cms_qe_newsletter_mailinglist {
        BigAutoField id PK
        CharField name
        MailingListExternalService external_service
        CharField external_id
    }
    cms_qe_newsletter_subscriber {
        BigAutoField id PK
        ForeignKey mailing_list_id FK
        EmailField email
        CharField first_name
        CharField last_name
        CharField external_id
    }
    cms_qe_newsletter_subscribetask {
        BigAutoField id PK
        ForeignKey mailing_list_id FK
        EmailField email
        CharField first_name
        CharField last_name
        CharField external_id
        SubscribeTaskType type
        DateTimeField created
        PositiveSmallIntegerField attempts
        TextField last_error
    }
    cms_qe_table_tablepluginmodel {
        OneToOneField cmsplugin_ptr_id PK
        CharField table
        JSONField columns
        JSONField filter
        BooleanField paging_show
        IntegerField paging_per_page
    }
    cms_qe_video_sourcefilevideoplayer {
        CharField label
        FilerImageField poster_id FK
        IntegerField width "Leave it blank to make a video player of the default width of video source"
        IntegerField height "Leave it blank to make a video player of the default width of video source"
        BooleanField controls
        BooleanField autoplay
        BooleanField loop
        AttributesField other_attributes
        OneToOneField cmsplugin_ptr_id PK
        FilerFileField source_file_id FK
        CharField text_title
        TextField text_description
        AttributesField attributes
        BooleanField muted
    }
    cms_qe_video_hostingvideoplayer {
        CharField label
        FilerImageField poster_id FK
        IntegerField width "Leave it blank to make a video player of the default width of video source"
        IntegerField height "Leave it blank to make a video player of the default width of video source"
        BooleanField controls
        BooleanField autoplay
        BooleanField loop
        AttributesField other_attributes
        OneToOneField cmsplugin_ptr_id PK
        HostingVideoPlayerVideoHostingService video_hosting_service
        URLField video_url "Use this field to embed videos from external services such as YouTube, Vimeo or others."
    }
    cms_qe_video_videotrack {
        OneToOneField cmsplugin_ptr_id PK
        VideoTrackKind kind
        FilerFileField src_id FK
        CharField srclang "Examples: 'en' or 'de' etc."
        CharField label
        AttributesField attributes
    }
    cms_qe_plugins_script {
        OneToOneField cmsplugin_ptr_id PK
        CharField src "Specifies the location of the linked script."
        BooleanField asyncf "Specifies that the script is downloaded in parallel to parsing the page, and executed as soon as it is available (before parsing completes) (only for external scripts)"
        ScriptCrossOrigin cross_origin "Sets the mode of the request to an HTTP CORS Request."
        BooleanField defer "Specifies that the script is downloaded in parallel to parsing the page, and executed after the page has finished parsing (only for external scripts)"
        CharField integrity "Allows a browser to check the fetched script to ensure that the code is never loaded if the source has been manipulated"
        BooleanField nomodule "Specifies that the script should not be executed in browsers supporting ES2015 modules."
        ScriptReferrerPolicy referrer_policy "Specifies which referrer to use when fetching the resource."
        CharField type "Specifies the media type of the script."
        JSONField attributes "More attributes as JSON data. E.g. {'data-name': 'value', 'id': 42}"
    }
    cms_qe_plugins_link {
        OneToOneField cmsplugin_ptr_id PK
        CharField href "Specifies the location of the linked document."
        CmsQePluginsLinkRel rel "Required. Specifies the relationship between the current document and the linked document."
        CmsQePluginsLinkCrossOrigin cross_origin "Specifies how the element handles cross-origin requests."
        CharField hreflang "Specifies the language of the text in the linked document."
        CharField media "Specifies on what device the linked document will be displayed."
        CmsQePluginsLinkReferrerPolicy referrer_policy "Specifies which referrer to use when fetching the resource."
        CharField sizes "Specifies the size of the linked resource. Only for rel='icon'."
        CharField title "Defines a preferred or an alternate stylesheet."
        CharField type "Specifies the media type of the linked document. E.g. 'text/css'."
        JSONField attributes "More attributes as JSON data. E.g. {'data-name': 'value', 'id': 42}"
    }
    cms_qe_plugins_iframe {
        OneToOneField cmsplugin_ptr_id PK
        CharField src "Specifies the address of the document to embed in the &lt;iframe&gt;."
        CharField allow "Specifies a feature policy for the &lt;iframe&gt;. Form more see <a href='https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe#allow' target='_blank'>IFrame attributes.</a>."
        BooleanField allow_full_screen "Set to true if the &lt;iframe&gt; can activate fullscreen mode by calling the requestFullscreen() method."
        BooleanField allow_payment_request "Set to true if a cross-origin &lt;iframe&gt; should be allowed to invoke the Payment Request API."
        IntegerField height "Specifies the height of an &lt;iframe&gt;. Default height is 150 pixels."
        IntegerField width "Specifies the width of an &lt;iframe&gt;. Default width is 300 pixels."
        IframeLoading loading "Specifies whether a browser should load an iframe immediately or to defer loading of iframes until some conditions are met."
        CharField name "Specifies the name of an &lt;iframe&gt;."
        IframeReferrerPolicy referrer_policy "Specifies which referrer information to send when fetching the iframe."
        IframeSandbox sandbox "Enables an extra set of restrictions for the content in an &lt;iframe&gt;."
        TextField src_doc "Specifies the HTML content of the page to show in the &lt;iframe&gt;."
        JSONField attributes "More attributes as JSON data. E.g. {'data-name': 'value', 'id': 42}"
    }
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
        PageContentTemplate template "The template used to render the content."
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
    djangocms_text_text {
        OneToOneField cmsplugin_ptr_id PK
        TextField body
        JSONField json
        CharField rte "The rich text editor used to create this text. JSON formats vary between editors."
    }
    djangocms_link_link {
        DjangocmsLinkLinkTemplate template
        CharField name
        LinkField link "-"
        DjangocmsLinkLinkLinkTarget link_target
        AttributesField attributes
        OneToOneField cmsplugin_ptr_id PK
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
    djangocms_googlemap_googlemap {
        GoogleMapTemplate template
        CharField title
        CharField width "Width of the map, including the CSS length units (e.g. '100%', '400px' or '400rem')."
        CharField height "Height of the map, including the CSS length units (e.g. '400px' or '400rem')."
        TextField style "Provide a valid (escaped) JSON configuration. See http://developers.google.com/maps/documentation/javascript/styling"
        CharField map_id "Map ID of Cloud-based maps styling. See https://developers.google.com/maps/documentation/cloud-customization/overview"
        FloatField lat "Default Geographical latitude in degrees (e.g. '46.947973')."
        FloatField lng "Default Geographical longitude in degrees (e.g. '7.447446')."
        GoogleMapZoom zoom
        GoogleMapMapTypeControl map_type_control
        BooleanField zoom_control "Enabling zooming allows the user to view the map at different scales."
        BooleanField street_view_control "Street View allows the user to see street-level imagery."
        BooleanField rotate_control "Allows the user to change the orientation of the map."
        BooleanField scale_control "Displays a linear distance scale on the map."
        BooleanField fullscreen_control "Allows the user to select a fullscreen view of the map."
        BooleanField pan_control "Allows the user to move the map using the cursor keys to see new areas."
        BooleanField double_click_zoom "Allows the user to zoom and centre the map with a double-click or double-tap."
        BooleanField draggable "Allows the user to drag the map to see new areas."
        BooleanField keyboard_shortcuts "Allows the user to control the map using the keyboard."
        BooleanField scrollwheel "Allows the user to control the zoom level with a scroll-wheel."
        OneToOneField cmsplugin_ptr_id PK
    }
    djangocms_googlemap_googlemapmarker {
        OneToOneField cmsplugin_ptr_id PK
        CharField title
        CharField address "Note: Latitude and longitude can be used to fine-tune the location."
        FloatField lat "Geographical latitude in degrees (e.g. '46.947973')."
        FloatField lng "Geographical longitude in degrees (e.g. '7.447446')."
        FilerImageField icon_id FK "A marker icon identifies a location on a map. By default, it uses a standard image from Google."
        BooleanField show_content "Display the info window when the map loads."
        TextField info_content "Will be displayed in the info window attached to the marker."
    }
    djangocms_googlemap_googlemaproute {
        OneToOneField cmsplugin_ptr_id PK
        CharField title
        CharField origin "Will be determined by user's location (if possible) if left blank."
        CharField destination
        GoogleMapRouteTravelMode travel_mode
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
    djangocms_icon_icon {
        OneToOneField cmsplugin_ptr_id PK
        Icon icon
        IconTemplate template
        CharField label
        AttributesField attributes
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
    djangocms_frontend_frontenduiitem {
        OneToOneField cmsplugin_ptr_id PK
        CharField ui_item
        TagTypeField tag_type "Select the HTML tag to be used."
        JSONField config
    }
    djangocms_render_context_rendercontext {
        OneToOneField cmsplugin_ptr_id PK
        CharField mimetype "The format must be consistent with the data in the context."
        TextField context "Context data in selected format. They take precedence over the source file and source URL."
        JSONField config "Additional configuration for spreadsheet with source data. The configuration is automatically added or updated when the spreadsheet is saved. To edit the configuration 'manually', you must have the Data for the context set to Source."
        FilerFileField file_id FK "Context data file. If Data is specified for the context, the source file is not used even though it is specified."
        URLField source "The URL of the source from which the data will be downloaded. If Data for Context or Source File is specified, the source URL is not used even though it is specified."
        PositiveSmallIntegerField cached "The time in minutes during which data will not be retrieved from the Source URL, but will remain in the cache. If set to zero, data from the source URL is not cached, but is loaded every time."
        TextField template "Django template for the context. It takes precedence over the template selected from the list. The context is available through the 'data' value."
        CharField template_list "List of templates. If Template is specified, the template set in the list will not be used."
        CharField path "The path to the data in the form 'name.name.name'. The 'data' context in the template starts from this key."
        CharField processor "List of functions that process the context."
        CharField detail_extends "The detail uses a template. If not specified, the default template is used."
        TextField detail_template "Django template for the context. It takes precedence over the template selected from the list. The context is available through the 'detail' value. The entire context is still available via the 'data' key."
        CharField detail_template_list "List of templates. If Template is specified, the template set in the list will not be used."
        CharField detail_processor "List of functions that process the context."
    }
    djangocms_render_context_default {
        OneToOneField cmsplugin_ptr_id PK
        CharField code "The value name for the default type."
    }
    djangocms_render_context_stack {
        OneToOneField cmsplugin_ptr_id PK
        CharField selector "Keys or indexes separated by a dot."
    }
    axes_accessfailurelog {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        BooleanField locked_out
    }
    axes_accessattempt {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        TextField get_data
        TextField post_data
        PositiveIntegerField failures_since_start
    }
    axes_accessattemptexpiration {
        OneToOneField access_attempt_id PK
        DateTimeField expires_at "The time when access attempt expires and is no longer valid."
    }
    axes_accesslog {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        DateTimeField logout_time
        CharField session_hash
    }
    constance_constance {
        AutoField id PK
        CharField key UK
        TextField value
    }
    aldryn_forms_webhook {
        BigAutoField id PK
        CharField name UK
        URLField url UK
        WebhookMethod method
        JSONField transform
    }
    aldryn_forms_formplugin {
        CharField name "Used to filter out form submissions."
        TextField error_message "An error message that will be displayed if the form doesn't validate."
        TextField success_message "An success message that will be displayed."
        BooleanField message_on_form
        BooleanField hide_form_after_redirection
        AldrynFormsLinkField redirect_to "-"
        CharField custom_classes
        FormPluginFormTemplate form_template
        CharField action_backend
        AttributesField form_attributes
        BooleanField is_enable_autofill_from_url_params "Eg if you open the form with a url that contains parameters as 'https://example.com/sub-page/?email=hello@example.com&name=Alex', then the fields 'email' and 'name' are going to to be filled in automatically."
        BooleanField use_form_action "Use the URL or CMS Page in the form attribute 'action'."
        OneToOneField cmsplugin_ptr_id PK
    }
    aldryn_forms_fieldsetplugin {
        CharField legend
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
    }
    aldryn_forms_fieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        CharField pattern "A regular expression to match for the value. It must be a valid JavaScript regular expression."
    }
    aldryn_forms_textareafieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        PositiveIntegerField text_area_columns
        PositiveIntegerField text_area_rows
    }
    aldryn_forms_emailfieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        BooleanField email_send_notification "When checked, the value of this field will be used to send an email notification."
        CharField email_subject "Used as the email subject when email_send_notification is checked."
        TextField email_body "Additional body text used when email notifications are active."
        CharField pattern "A regular expression to match for the value. It must be a valid JavaScript regular expression."
    }
    aldryn_forms_fileuploadfieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        FilerFolderField upload_to_id FK "Select a folder to which all files submitted through this field will be uploaded to."
        FileSizeField max_size "The maximum size of all files to upload, in bytes. You can use common size suffixes (kB, MB, GB, ...)."
        BooleanField enable_js "Enable javascript to view files for upload."
        CharField accepted_types "The list of accepted types. E.g. '.pdf .jpg .png text/plain application/msword image/*'."
        BooleanField is_private "The file is private."
    }
    aldryn_forms_multiplefilesuploadfieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        FilerFolderField upload_to_id FK "Select a folder to which all files submitted through this field will be uploaded to."
        FileSizeField max_size "The maximum size of all files to upload, in bytes. You can use common size suffixes (kB, MB, GB, ...)."
        BooleanField enable_js "Enable javascript to view files for upload."
        CharField accepted_types "The list of accepted types. E.g. '.pdf .jpg .png text/plain application/msword image/*'."
        PositiveSmallIntegerField max_files "Maximum number of uploaded files."
        BooleanField is_private "The files are private."
    }
    aldryn_forms_imageuploadfieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        FilerFolderField upload_to_id FK "Select a folder to which all files submitted through this field will be uploaded to."
        FileSizeField max_size "The maximum size of all files to upload, in bytes. You can use common size suffixes (kB, MB, GB, ...)."
        BooleanField enable_js "Enable javascript to view files for upload."
        PositiveIntegerField max_width "The maximum width of the uploaded image, in pixels."
        PositiveIntegerField max_height "The maximum height of the uploaded image, in pixels."
        BooleanField is_private "The file is private."
    }
    aldryn_forms_datefieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        DateField earliest_date "The earliest date to accept."
        DateField latest_date "The latest date to accept."
        PositiveIntegerField input_step "The granularity numnber."
    }
    aldryn_forms_datetimelocalfieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        DateTimeField earliest_datetime "The earliest datetime to accept."
        DateTimeField latest_datetime "The latest datetime to accept."
        PositiveIntegerField input_step "The granularity numnber."
    }
    aldryn_forms_timefieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        TimeField earliest_time "The earliest time to accept."
        TimeField latest_time "The latest time to accept."
        PositiveIntegerField input_step "The granularity numnber."
        SlugField data_list "Datalist id of html element datalist with options."
        BooleanField readonly "The field cannot be edited by the user."
    }
    aldryn_forms_urlfieldplugin {
        CharField name "Used to set the field name"
        CharField label
        BooleanField required
        TextField required_message "Error message displayed if the required field is left empty. Default: 'This field is required'."
        CharField placeholder_text "Default text in a form. Disappears when user starts typing. Example: 'email@example.com'"
        TextField help_text "Explanatory text displayed next to input field. Just like this one."
        AttributesField attributes
        PositiveIntegerField min_value
        PositiveIntegerField max_value
        CharField initial_value "Default value of field."
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
        TextField list "A list of predefined url values. Each url is on a separate line. The first word is the url. The next (optional) words are a description for the url."
        CharField pattern "A regular expression to match for the value. It must be a valid JavaScript regular expression."
        BooleanField readonly "The field cannot be edited by the user."
        SmallIntegerField size "How many characters wide the input field should be. The default value is 20."
        BooleanField spellcheck "To enable spell-checking for an element."
    }
    aldryn_forms_option {
        BigAutoField id PK
        ForeignKey field_id FK
        CharField value
        CharField widget_value
        BooleanField default_value
        PositiveIntegerField position
    }
    aldryn_forms_formbuttonplugin {
        CharField label
        CharField custom_classes
        OneToOneField cmsplugin_ptr_id PK
    }
    aldryn_forms_formsubmission {
        BigAutoField id PK
        CharField name
        TextField data
        TextField recipients "People who got a notification when form was submitted."
        CharField language
        CharField form_url
        DateTimeField sent_at
        CharField post_ident
        BooleanField honeypot_filled
    }
    aldryn_forms_submittedtobesent {
        BigAutoField id PK
        CharField name
        TextField data
        TextField recipients "People who got a notification when form was submitted."
        CharField language
        CharField form_url
        DateTimeField sent_at
        CharField post_ident
        BooleanField honeypot_filled
    }
    email_notifications_emailnotification {
        BigAutoField id PK
        EmailNotificationTheme theme "Provides the base theme for the email."
        CharField to_name
        CharField to_email
        ForeignKey to_user_id FK
        CharField from_name
        CharField from_email
        CharField reply_to_email
        CharField subject
        TextField body_text "used when rendering the email in text only mode."
        HTMLField body_html "used when rendering the email in html."
        ForeignKey form_id FK
    }
    captcha_captchastore {
        AutoField id PK
        CharField challenge
        CharField response
        CharField hashkey UK
        DateTimeField expiration
    }
    knox_authtoken {
        CharField digest PK
        CharField token_key
        ForeignKey user_id FK
        DateTimeField created
        DateTimeField expiry
    }
```
