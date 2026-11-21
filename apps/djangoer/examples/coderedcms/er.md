```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    coderedcms_carousel ||--o{ coderedcms_carouselslide : "(id) - (carousel_id)"
    wagtailimages_image |o--o{ coderedcms_carouselslide : "(id) - (image_id)"
    coderedcms_classifier ||--o{ coderedcms_classifierterm : "(id) - (classifier_id)"
    coderedcms_filmstrip ||--o{ coderedcms_filmpanel : "(id) - (film_strip_id)"
    wagtailimages_image |o--o{ coderedcms_filmpanel : "(id) - (background_image_id)"
    coderedcms_accordion ||--o{ coderedcms_accordionpanel : "(id) - (accordion_id)"
    wagtailcore_site ||--o| coderedcms_layoutsettings : "(id) - (site_id)"
    wagtailimages_image |o--o{ coderedcms_layoutsettings : "(id) - (logo_id)"
    wagtailimages_image |o--o{ coderedcms_layoutsettings : "(id) - (favicon_id)"
    coderedcms_layoutsettings ||--o{ coderedcms_navbarorderable : "(id) - (navbar_chooser_id)"
    coderedcms_navbar |o--o{ coderedcms_navbarorderable : "(id) - (navbar_id)"
    coderedcms_layoutsettings ||--o{ coderedcms_footerorderable : "(id) - (footer_chooser_id)"
    coderedcms_footer |o--o{ coderedcms_footerorderable : "(id) - (footer_id)"
    wagtailcore_site ||--o| coderedcms_analyticssettings : "(id) - (site_id)"
    taggit_tag ||--o{ coderedcms_coderedtag : "(id) - (tag_id)"
    coderedcms_coderedpage ||--o{ coderedcms_coderedtag : "(page_ptr_id) - (content_object_id)"
    wagtailcore_page ||--o| coderedcms_coderedpage : "is a, by page_ptr_id"
    wagtailimages_image |o--o{ coderedcms_coderedpage : "(id) - (og_image_id)"
    wagtailimages_image |o--o{ coderedcms_coderedpage : "(id) - (cover_image_id)"
    coderedcms_classifier |o--o{ coderedcms_coderedpage : "(id) - (index_order_by_classifier_id)"
    coderedcms_classifierterm |o--o{ coderedcms_coderedpage : "(id) - (related_classifier_term_id)"
    coderedcms_coderedpage }o--o{ coderedcms_classifier : "(index_classifiers) - (coderedpage_set)"
    coderedcms_coderedpage }o--o{ coderedcms_classifierterm : "(classifier_terms) - (coderedpage_set)"
    coderedcms_coderedpage ||..o{ wagtailcore_revision : "(page_ptr_id) - (object_id)"
    coderedcms_coderedpage ||..o{ wagtailcore_workflowstate : "(page_ptr_id) - (object_id)"
    coderedcms_coderedpage ||..o{ wagtailcore_workflowstate : "(page_ptr_id) - (object_id)"
    coderedcms_coderedpage ||..o{ wagtailcore_revision : "(page_ptr_id) - (object_id)"
    coderedcms_coderedpage ||..o{ wagtailcore_workflowstate : "(page_ptr_id) - (object_id)"
    coderedcms_coderedpage ||..o{ wagtailcore_workflowstate : "(page_ptr_id) - (object_id)"
    coderedcms_coderedpage ||..o{ wagtailsearch_indexentry : "(page_ptr_id) - (object_id)"
    django_content_type ||--o{ coderedcms_coderedsubmissionrevision : "(id) - (submission_ct_id)"
    wagtailcore_page ||--o{ coderedcms_coderedsessionformsubmission : "(id) - (page_id)"
    auth_user |o--o{ coderedcms_coderedsessionformsubmission : "(id) - (user_id)"
    taggit_tag ||--o{ taggit_taggeditem : "(id) - (tag_id)"
    django_content_type ||--o{ taggit_taggeditem : "(id) - (content_type_id)"
    wagtailcore_site ||--o| wagtailseo_seosettings : "(id) - (site_id)"
    wagtailimages_image |o--o{ wagtailseo_seosettings : "(id) - (struct_org_logo_id)"
    wagtailimages_image |o--o{ wagtailseo_seosettings : "(id) - (struct_org_image_id)"
    wagtailimages_image |o--o{ wagtailseo_seosettings : "(id) - (og_image_default_id)"
    wagtailcore_page ||--o{ wagtailforms_formsubmission : "(id) - (page_id)"
    wagtailcore_site |o--o{ wagtailredirects_redirect : "(id) - (site_id)"
    wagtailcore_page |o--o{ wagtailredirects_redirect : "(id) - (redirect_page_id)"
    wagtailsearchpromotions_query ||--o{ wagtailsearchpromotions_querydailyhits : "(id) - (query_id)"
    wagtailsearchpromotions_query ||--o{ wagtailsearchpromotions_searchpromotion : "(id) - (query_id)"
    wagtailcore_page |o--o{ wagtailsearchpromotions_searchpromotion : "(id) - (page_id)"
    auth_user ||--o| wagtailusers_userprofile : "(id) - (user_id)"
    wagtailcore_collection ||--o{ wagtaildocs_document : "(id) - (collection_id)"
    auth_user |o--o{ wagtaildocs_document : "(id) - (uploaded_by_user_id)"
    wagtaildocs_document ||..o{ taggit_taggeditem : "(id) - (object_id)"
    wagtaildocs_document ||..o{ wagtailsearch_indexentry : "(id) - (object_id)"
    wagtailcore_collection ||--o{ wagtailimages_image : "(id) - (collection_id)"
    auth_user |o--o{ wagtailimages_image : "(id) - (uploaded_by_user_id)"
    wagtailimages_image ||..o{ taggit_taggeditem : "(id) - (object_id)"
    wagtailimages_image ||..o{ wagtailsearch_indexentry : "(id) - (object_id)"
    wagtailimages_image ||--o{ wagtailimages_rendition : "(id) - (image_id)"
    django_content_type ||--o{ wagtailsearch_indexentry : "(id) - (content_type_id)"
    wagtailsearch_indexentry ||--o| wagtailsearch_indexentry_fts : "(id) - (rowid)"
    wagtailcore_page ||--o{ wagtailcore_site : "(id) - (root_page_id)"
    auth_group ||--o{ wagtailcore_groupsitepermission : "(id) - (group_id)"
    wagtailcore_site ||--o{ wagtailcore_groupsitepermission : "(id) - (site_id)"
    auth_permission ||--o{ wagtailcore_groupsitepermission : "(id) - (permission_id)"
    django_content_type |o--o{ wagtailcore_modellogentry : "(id) - (content_type_id)"
    auth_user |o..o{ wagtailcore_modellogentry : "(id) - (user_id), no constraint"
    wagtailcore_revision |o..o{ wagtailcore_modellogentry : "(id) - (revision_id), no constraint"
    django_content_type ||--o{ wagtailcore_revision : "(id) - (content_type_id)"
    django_content_type ||--o{ wagtailcore_revision : "(id) - (base_content_type_id)"
    auth_user |o--o{ wagtailcore_revision : "(id) - (user_id)"
    wagtailcore_collection ||--o{ wagtailcore_collectionviewrestriction : "(id) - (collection_id)"
    wagtailcore_collectionviewrestriction }o--o{ auth_group : "(groups) - (collectionviewrestriction_set)"
    auth_group ||--o{ wagtailcore_groupcollectionpermission : "(id) - (group_id)"
    wagtailcore_collection ||--o{ wagtailcore_groupcollectionpermission : "(id) - (collection_id)"
    auth_permission ||--o{ wagtailcore_groupcollectionpermission : "(id) - (permission_id)"
    django_content_type |o--o{ wagtailcore_uploadedfile : "(id) - (for_content_type_id)"
    auth_user |o--o{ wagtailcore_uploadedfile : "(id) - (uploaded_by_user_id)"
    django_content_type ||--o| wagtailcore_workflowcontenttype : "(id) - (content_type_id)"
    wagtailcore_workflow ||--o{ wagtailcore_workflowcontenttype : "(id) - (workflow_id)"
    django_content_type ||--o{ wagtailcore_workflowstate : "(id) - (content_type_id)"
    django_content_type ||--o{ wagtailcore_workflowstate : "(id) - (base_content_type_id)"
    wagtailcore_workflow ||--o{ wagtailcore_workflowstate : "(id) - (workflow_id)"
    auth_user |o--o{ wagtailcore_workflowstate : "(id) - (requested_by_id)"
    wagtailcore_taskstate |o--o| wagtailcore_workflowstate : "(id) - (current_task_state_id)"
    wagtailcore_workflow ||--o{ wagtailcore_workflowtask : "(id) - (workflow_id)"
    wagtailcore_task ||--o{ wagtailcore_workflowtask : "(id) - (task_id)"
    django_content_type ||--o{ wagtailcore_task : "(id) - (content_type_id)"
    wagtailcore_task ||--o| wagtailcore_groupapprovaltask : "is a, by task_ptr_id"
    wagtailcore_groupapprovaltask }o--o{ auth_group : "(groups) - (groupapprovaltask_set)"
    wagtailcore_workflowstate ||--o{ wagtailcore_taskstate : "(id) - (workflow_state_id)"
    wagtailcore_revision ||--o{ wagtailcore_taskstate : "(id) - (revision_id)"
    wagtailcore_task ||--o{ wagtailcore_taskstate : "(id) - (task_id)"
    auth_user |o--o{ wagtailcore_taskstate : "(id) - (finished_by_id)"
    django_content_type ||--o{ wagtailcore_taskstate : "(id) - (content_type_id)"
    wagtailcore_locale ||--o{ wagtailcore_page : "(id) - (locale_id)"
    wagtailcore_revision |o--o{ wagtailcore_page : "(id) - (latest_revision_id)"
    wagtailcore_revision |o--o{ wagtailcore_page : "(id) - (live_revision_id)"
    auth_user |o--o{ wagtailcore_page : "(id) - (locked_by_id)"
    django_content_type ||--o{ wagtailcore_page : "(id) - (content_type_id)"
    auth_user |o--o{ wagtailcore_page : "(id) - (owner_id)"
    wagtailcore_page |o--o{ wagtailcore_page : "(id) - (alias_of_id)"
    wagtailcore_page ||..o{ wagtailcore_revision : "(id) - (object_id)"
    wagtailcore_page ||..o{ wagtailcore_workflowstate : "(id) - (object_id)"
    wagtailcore_page ||..o{ wagtailcore_workflowstate : "(id) - (object_id)"
    wagtailcore_page ||..o{ wagtailsearch_indexentry : "(id) - (object_id)"
    auth_group ||--o{ wagtailcore_grouppagepermission : "(id) - (group_id)"
    wagtailcore_page ||--o{ wagtailcore_grouppagepermission : "(id) - (page_id)"
    auth_permission ||--o{ wagtailcore_grouppagepermission : "(id) - (permission_id)"
    wagtailcore_page ||--o{ wagtailcore_pageviewrestriction : "(id) - (page_id)"
    wagtailcore_pageviewrestriction }o--o{ auth_group : "(groups) - (pageviewrestriction_set)"
    wagtailcore_page ||--o| wagtailcore_workflowpage : "(id) - (page_id)"
    wagtailcore_workflow ||--o{ wagtailcore_workflowpage : "(id) - (workflow_id)"
    django_content_type |o--o{ wagtailcore_pagelogentry : "(id) - (content_type_id)"
    auth_user |o..o{ wagtailcore_pagelogentry : "(id) - (user_id), no constraint"
    wagtailcore_revision |o..o{ wagtailcore_pagelogentry : "(id) - (revision_id), no constraint"
    wagtailcore_page ||..o{ wagtailcore_pagelogentry : "(id) - (page_id), no constraint"
    wagtailcore_page ||--o{ wagtailcore_comment : "(id) - (page_id)"
    auth_user ||--o{ wagtailcore_comment : "(id) - (user_id)"
    wagtailcore_revision |o--o{ wagtailcore_comment : "(id) - (revision_created_id)"
    auth_user |o--o{ wagtailcore_comment : "(id) - (resolved_by_id)"
    wagtailcore_comment ||--o{ wagtailcore_commentreply : "(id) - (comment_id)"
    auth_user ||--o{ wagtailcore_commentreply : "(id) - (user_id)"
    auth_user ||--o{ wagtailcore_pagesubscription : "(id) - (user_id)"
    wagtailcore_page ||--o{ wagtailcore_pagesubscription : "(id) - (page_id)"
    django_content_type ||--o{ wagtailcore_referenceindex : "(id) - (content_type_id)"
    django_content_type ||--o{ wagtailcore_referenceindex : "(id) - (base_content_type_id)"
    django_content_type ||--o{ wagtailcore_referenceindex : "(id) - (to_content_type_id)"
    auth_user ||--o{ wagtailadmin_editingsession : "(id) - (user_id)"
    django_content_type ||--o{ wagtailadmin_editingsession : "(id) - (content_type_id)"
    auth_user ||--o{ wagtailadmin_formstate : "(id) - (user_id)"
    django_content_type ||--o{ wagtailadmin_formstate : "(id) - (content_type_id)"
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
    coderedcms_carousel {
        AutoField id PK
        CharField name
        BooleanField show_controls "Shows arrows on the left and right of the carousel to advance next or previous slides."
        BooleanField show_indicators "Shows small indicators at the bottom of the carousel based on the number of slides."
    }
    coderedcms_carouselslide {
        AutoField id PK
        IntegerField sort_order
        ParentalKey carousel_id FK
        ForeignKey image_id FK
        CharField background_color "Hexadecimal, rgba, or CSS color notation (e.g. #ff0011)"
        CharField custom_css_class
        CharField custom_id
        CoderedStreamField content
    }
    coderedcms_classifier {
        AutoField id PK
        SlugField slug UK
        CharField name
    }
    coderedcms_classifierterm {
        AutoField id PK
        IntegerField sort_order
        ParentalKey classifier_id FK
        SlugField slug UK
        CharField name
    }
    coderedcms_filmstrip {
        AutoField id PK
        CharField name
    }
    coderedcms_filmpanel {
        AutoField id PK
        IntegerField sort_order
        ParentalKey film_strip_id FK
        ForeignKey background_image_id FK
        CharField background_color "Hexadecimal, rgba, or CSS color notation (e.g. #ff0011)"
        CharField foreground_color "Hexadecimal, rgba, or CSS color notation (e.g. #ff0011)"
        CharField custom_css_class
        CharField custom_id
        CoderedStreamField content
    }
    coderedcms_navbar {
        AutoField id PK
        CharField name
        CharField custom_css_class
        CharField custom_id
        CoderedStreamField menu_items
    }
    coderedcms_footer {
        AutoField id PK
        CharField name
        CharField custom_css_class
        CharField custom_id
        CoderedStreamField content
    }
    coderedcms_reusablecontent {
        AutoField id PK
        CharField name
        CoderedStreamField content
    }
    coderedcms_accordion {
        AutoField id PK
        CharField name
    }
    coderedcms_accordionpanel {
        AutoField id PK
        IntegerField sort_order
        ParentalKey accordion_id FK
        CharField name
        CoderedStreamField content
        CharField custom_css_class
        CharField custom_id
    }
    coderedcms_contentwall {
        AutoField id PK
        CharField name
        CoderedStreamField content
        BooleanField is_dismissible
        BooleanField show_once "Do not show the content wall to the same user again after it has been closed."
    }
    coderedcms_layoutsettings {
        AutoField id PK
        OneToOneField site_id FK, UK
        ForeignKey logo_id FK "Brand logo used in the navbar and throughout the site"
        ForeignKey favicon_id FK
        CharField navbar_color_scheme "Optimizes text and other navbar elements for use with light or dark backgrounds."
        CharField navbar_class "Custom classes applied to navbar e.g. 'bg-light', 'bg-dark', 'bg-primary'."
        BooleanField navbar_fixed "Fixed navbar will remain at the top of the page when scrolling."
        BooleanField navbar_content_fluid "Content within the navbar will fill edge to edge."
        CharField navbar_collapse_mode "Control on what screen sizes to show and collapse the navbar menu links."
        CharField navbar_format
        BooleanField navbar_search "Show search box in navbar"
        CharField from_email_address "The default email address this site appears to send from. For example: 'sender@example.com' or 'Sender Name <sender@example.com>' (without quotes)"
        PositiveIntegerField search_num_results
        BooleanField external_new_tab
        LayoutSettingsSpamService spam_service "Choose a technique or 3rd party service to help block spam submissions."
        DecimalField recaptcha_threshold "reCAPTCHA v3 returns a score (0.0 is very likely a bot, 1.0 is very likely a good interaction). Reject submissions below this score (recommended 0.5)."
        CharField recaptcha_public_key "Create this key in the Google reCAPTCHA or Google Cloud dashboard."
        CharField recaptcha_secret_key "Create this key in the Google reCAPTCHA or Google Cloud dashboard."
        CharField google_maps_api_key "The API Key used for Google Maps."
        CharField mailchimp_api_key "The API Key used for Mailchimp."
    }
    coderedcms_navbarorderable {
        AutoField id PK
        IntegerField sort_order
        ParentalKey navbar_chooser_id FK
        ForeignKey navbar_id FK
    }
    coderedcms_footerorderable {
        AutoField id PK
        IntegerField sort_order
        ParentalKey footer_chooser_id FK
        ForeignKey footer_id FK
    }
    coderedcms_analyticssettings {
        AutoField id PK
        OneToOneField site_id FK, UK
        CharField ga_g_tracking_id "Your Google Analytics 4 tracking ID (begins with 'G-')"
        BooleanField ga_track_button_clicks "Track all button clicks using Google Analytics event tracking. Event tracking details can be specified in each button’s advanced settings options."
        CharField gtm_id "Begins with 'GTM-'"
        MonospaceField head_scripts "Add tracking scripts between the <head> tags."
        MonospaceField body_scripts "Add tracking scripts toward closing <body> tag."
    }
    coderedcms_coderedtag {
        AutoField id PK
        ForeignKey tag_id FK
        ParentalKey content_object_id FK
    }
    coderedcms_coderedpage {
        OneToOneField page_ptr_id PK
        URLField canonical_url "Leave blank to use the page's URL."
        ForeignKey og_image_id FK "Shown when linking to this page on social media. If blank, may show an image from the page, or the default from Settings > SEO."
        ForeignKey cover_image_id FK
        BooleanField index_show_subpages
        ForeignKey index_order_by_classifier_id FK "Child pages will first be sorted following the order of this classifier’s terms (from Snippets > Classifiers)."
        CoderedPageIndexOrderBy index_order_by "Child pages will then be sorted by this attribute."
        PositiveIntegerField index_num_per_page
        BooleanField related_show
        PositiveIntegerField related_num
        ForeignKey related_classifier_term_id FK "When getting related pages, pages with this term will be weighted over other classifier terms. By default, pages with the greatest number of classifiers in common are ranked highest."
        CharField custom_template
        CoderedStreamField content_walls
    }
    coderedcms_coderedsubmissionrevision {
        AutoField id PK
        CoderedSubmissionRevisionType type
        DateTimeField created_at
        ForeignKey submission_ct_id FK
        TextField submission_id
        TextField data
        TextField summary
    }
    coderedcms_coderedsessionformsubmission {
        AutoField id PK
        JSONField form_data
        ForeignKey page_id FK
        DateTimeField submit_time
        CharField session_key
        ForeignKey user_id FK
        DateTimeField last_modification
        CoderedSessionFormSubmissionStatus status
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
    wagtailseo_seosettings {
        AutoField id PK
        OneToOneField site_id FK, UK
        SeoSettingsStructOrgType struct_org_type "Organization business type."
        CharField struct_org_name "Leave blank to use the site name in Settings > Sites"
        ForeignKey struct_org_logo_id FK "Logo representative of the organisation. Must be 112x112px minimum; take note that it will be displayed on a white background"
        ForeignKey struct_org_image_id FK "A photo of the facility. This photo will be cropped to 1:1, 4:3, and 16:9 aspect ratios automatically."
        CharField struct_org_phone "Include country code for best results. For example: +1-216-555-8000"
        CharField struct_org_address_street "House number and street. For example, 55 Public Square Suite 1710"
        CharField struct_org_address_locality "City or locality. For example, Cleveland"
        CharField struct_org_address_region "State, province, county, or region. For example, OH"
        CharField struct_org_address_postal "Zip or postal code. For example, 44113"
        CharField struct_org_address_country "For example, USA. Two-letter ISO 3166-1 alpha-2 country code is also acceptable https://en.wikipedia.org/wiki/ISO_3166-1"
        DecimalField struct_org_geo_lat
        DecimalField struct_org_geo_lng
        StreamField struct_org_hours
        StreamField struct_org_actions
        TextField struct_org_extra_json "Additional JSON-LD inserted into the Organization dictionary. Must be properties of https://schema.org/Organization or the selected organization type."
        BooleanField og_meta "Show an optimized preview when linking to this site on social media. See https://ogp.me/"
        BooleanField twitter_meta "Shows content as a card when linking to this site on Twitter. See https://developer.twitter.com/en/docs/twitter-for-websites/cards"
        CharField twitter_site "The @username of the website owner’s Twitter handle."
        BooleanField struct_meta "Optimizes information about your organization for search engines. See https://schema.org/"
        ForeignKey og_image_default_id FK "Shown when linking to this page on social media. This can also be customized on each page."
    }
    wagtailforms_formsubmission {
        AutoField id PK
        JSONField form_data
        ForeignKey page_id FK
        DateTimeField submit_time
    }
    wagtailredirects_redirect {
        AutoField id PK
        CharField old_path
        ForeignKey site_id FK
        BooleanField is_permanent "Recommended. Permanent redirects ensure search engines forget the old page (the 'Redirect from') and index the new page instead."
        ForeignKey redirect_page_id FK
        CharField redirect_page_route_path "Optionally specify a route on the target page to redirect to. Leave blank to redirect to the default page route."
        URLField redirect_link
        BooleanField automatically_created
        DateTimeField created_at
    }
    wagtailsearchpromotions_query {
        AutoField id PK
        CharField query_string UK
    }
    wagtailsearchpromotions_querydailyhits {
        AutoField id PK
        ForeignKey query_id FK
        DateField date
        IntegerField hits
    }
    wagtailsearchpromotions_searchpromotion {
        AutoField id PK
        ForeignKey query_id FK
        ForeignKey page_id FK "Choose an internal page for this promotion"
        URLField external_link_url "Alternatively, use an external link for this promotion"
        CharField external_link_text
        TextField description "Applies to internal page or external link"
        IntegerField sort_order
    }
    wagtailembeds_embed {
        AutoField id PK
        TextField url
        SmallIntegerField max_width
        CharField hash UK
        EmbedType type
        TextField html
        TextField title
        TextField author_name
        TextField provider_name
        TextField thumbnail_url
        IntegerField width
        IntegerField height
        DateTimeField last_updated
        DateTimeField cache_until
    }
    wagtailusers_userprofile {
        AutoField id PK
        OneToOneField user_id FK, UK
        BooleanField submitted_notifications "Receive notification when a page is submitted for moderation"
        BooleanField approved_notifications "Receive notification when your page edit is approved"
        BooleanField rejected_notifications "Receive notification when your page edit is rejected"
        BooleanField updated_comments_notifications "Receive notification when comments have been created, resolved, or deleted on a page that you have subscribed to receive comment notifications on"
        CharField preferred_language "Select language for the admin"
        CharField current_time_zone "Select your current time zone"
        ImageField avatar
        JSONField dismissibles
        UserProfileAdminColorThemes theme
        UserProfileAdminContrastThemes contrast
        UserProfileAdminDensityThemes density
        BooleanField keyboard_shortcuts "Enable custom keyboard shortcuts specific to Wagtail."
    }
    wagtaildocs_document {
        AutoField id PK
        ForeignKey collection_id FK
        CharField title
        FileField file
        DateTimeField created_at
        ForeignKey uploaded_by_user_id FK
        PositiveBigIntegerField file_size
        CharField file_hash
    }
    wagtailimages_image {
        AutoField id PK
        ForeignKey collection_id FK
        CharField title
        WagtailImageField file
        CharField description
        IntegerField width
        IntegerField height
        DateTimeField created_at
        ForeignKey uploaded_by_user_id FK
        PositiveIntegerField focal_point_x
        PositiveIntegerField focal_point_y
        PositiveIntegerField focal_point_width
        PositiveIntegerField focal_point_height
        PositiveIntegerField file_size
        CharField file_hash
    }
    wagtailimages_rendition {
        AutoField id PK
        CharField filter_spec
        WagtailImageField file
        IntegerField width
        IntegerField height
        CharField focal_point_key
        ForeignKey image_id FK
    }
    wagtailsearch_indexentry {
        AutoField id PK
        ForeignKey content_type_id FK
        CharField object_id
        FloatField title_norm
        TextField autocomplete
        TextField title
        TextField body
    }
    wagtailsearch_indexentry_fts {
        TextField autocomplete
        TextField title
        TextField body
        OneToOneField rowid PK
    }
    wagtailcore_locale {
        AutoField id PK
        CharField language_code UK
    }
    wagtailcore_site {
        AutoField id PK
        CharField hostname
        IntegerField port "Set this to something other than 80 if you need a specific port number to appear in URLs (e.g. development on port 8000). Does not affect request handling (so port forwarding still works)."
        CharField site_name "Human-readable name for the site."
        ForeignKey root_page_id FK
        BooleanField is_default_site "If true, this site will handle requests for all other hostnames that do not have a site entry of their own"
    }
    wagtailcore_groupsitepermission {
        AutoField id PK
        ForeignKey group_id FK
        ForeignKey site_id FK
        ForeignKey permission_id FK
    }
    wagtailcore_modellogentry {
        AutoField id PK
        ForeignKey content_type_id FK
        TextField label
        CharField action
        JSONField data
        DateTimeField timestamp
        UUIDField uuid "Log entries that happened as part of the same user action are assigned the same UUID"
        ForeignKey user_id FK
        ForeignKey revision_id FK
        BooleanField content_changed
        BooleanField deleted
        CharField object_id
    }
    wagtailcore_revision {
        AutoField id PK
        ForeignKey content_type_id FK
        ForeignKey base_content_type_id FK
        CharField object_id
        DateTimeField created_at
        ForeignKey user_id FK
        TextField object_str
        JSONField content
        DateTimeField approved_go_live_at
    }
    wagtailcore_collectionviewrestriction {
        AutoField id PK
        CollectionViewRestrictionRestrictionType restriction_type
        CharField password "Shared passwords should not be used to protect sensitive content. Anyone who has this password will be able to view the content."
        ForeignKey collection_id FK
    }
    wagtailcore_collection {
        AutoField id PK
        CharField path UK
        PositiveIntegerField depth
        PositiveIntegerField numchild
        CharField name
    }
    wagtailcore_groupcollectionpermission {
        AutoField id PK
        ForeignKey group_id FK
        ForeignKey collection_id FK
        ForeignKey permission_id FK
    }
    wagtailcore_uploadedfile {
        AutoField id PK
        ForeignKey for_content_type_id FK
        FileField file
        ForeignKey uploaded_by_user_id FK
    }
    wagtailcore_workflowcontenttype {
        OneToOneField content_type_id PK
        ForeignKey workflow_id FK
    }
    wagtailcore_workflowstate {
        AutoField id PK
        ForeignKey content_type_id FK
        ForeignKey base_content_type_id FK
        CharField object_id
        ForeignKey workflow_id FK
        WorkflowStateStatus status
        DateTimeField created_at
        ForeignKey requested_by_id FK
        OneToOneField current_task_state_id FK, UK
    }
    wagtailcore_workflow {
        AutoField id PK
        CharField name
        BooleanField active "Active workflows can be added to pages/snippets. Deactivating a workflow does not remove it from existing pages/snippets."
    }
    wagtailcore_workflowtask {
        AutoField id PK
        IntegerField sort_order
        ParentalKey workflow_id FK
        ForeignKey task_id FK
    }
    wagtailcore_task {
        AutoField id PK
        CharField name
        ForeignKey content_type_id FK
        BooleanField active "Active tasks can be added to workflows. Deactivating a task does not remove it from existing workflows."
    }
    wagtailcore_groupapprovaltask {
        OneToOneField task_ptr_id PK
    }
    wagtailcore_taskstate {
        AutoField id PK
        ForeignKey workflow_state_id FK
        ForeignKey revision_id FK
        ForeignKey task_id FK
        TaskStateStatus status
        DateTimeField started_at
        DateTimeField finished_at
        ForeignKey finished_by_id FK
        TextField comment
        ForeignKey content_type_id FK
    }
    wagtailcore_page {
        AutoField id PK
        CharField path UK
        PositiveIntegerField depth
        PositiveIntegerField numchild
        UUIDField translation_key
        ForeignKey locale_id FK
        ForeignKey latest_revision_id FK
        BooleanField live
        BooleanField has_unpublished_changes
        DateTimeField first_published_at
        DateTimeField last_published_at
        ForeignKey live_revision_id FK
        DateTimeField go_live_at
        DateTimeField expire_at
        BooleanField expired
        BooleanField locked
        DateTimeField locked_at
        ForeignKey locked_by_id FK
        CharField title "The page title as you'd like it to be seen by the public"
        CharField draft_title
        SlugField slug "The name of the page as it will appear in URLs e.g http://domain.com/blog/[my-slug]/"
        ForeignKey content_type_id FK
        TextField url_path
        ForeignKey owner_id FK
        CharField seo_title "The name of the page displayed on search engine results as the clickable headline."
        BooleanField show_in_menus "Whether a link to this page will appear in automatically generated menus"
        TextField search_description "The descriptive text displayed underneath a headline in search engine results."
        DateTimeField latest_revision_created_at
        ForeignKey alias_of_id FK
    }
    wagtailcore_grouppagepermission {
        AutoField id PK
        ForeignKey group_id FK
        ForeignKey page_id FK
        ForeignKey permission_id FK
    }
    wagtailcore_pageviewrestriction {
        AutoField id PK
        PageViewRestrictionRestrictionType restriction_type
        CharField password "Shared passwords should not be used to protect sensitive content. Anyone who has this password will be able to view the content."
        ForeignKey page_id FK
    }
    wagtailcore_workflowpage {
        OneToOneField page_id PK
        ForeignKey workflow_id FK
    }
    wagtailcore_pagelogentry {
        AutoField id PK
        ForeignKey content_type_id FK
        TextField label
        CharField action
        JSONField data
        DateTimeField timestamp
        UUIDField uuid "Log entries that happened as part of the same user action are assigned the same UUID"
        ForeignKey user_id FK
        ForeignKey revision_id FK
        BooleanField content_changed
        BooleanField deleted
        ForeignKey page_id FK
    }
    wagtailcore_comment {
        AutoField id PK
        ParentalKey page_id FK
        ForeignKey user_id FK
        TextField text
        TextField contentpath
        TextField position
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey revision_created_id FK
        DateTimeField resolved_at
        ForeignKey resolved_by_id FK
    }
    wagtailcore_commentreply {
        AutoField id PK
        ParentalKey comment_id FK
        ForeignKey user_id FK
        TextField text
        DateTimeField created_at
        DateTimeField updated_at
    }
    wagtailcore_pagesubscription {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey page_id FK
        BooleanField comment_notifications
    }
    wagtailcore_referenceindex {
        AutoField id PK
        ForeignKey content_type_id FK
        ForeignKey base_content_type_id FK
        CharField object_id
        ForeignKey to_content_type_id FK
        CharField to_object_id
        TextField model_path
        TextField content_path
        UUIDField content_path_hash
    }
    wagtailadmin_admin {
        AutoField id PK
    }
    wagtailadmin_editingsession {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey content_type_id FK
        CharField object_id
        DateTimeField last_seen_at
        BooleanField is_editing
    }
    wagtailadmin_formstate {
        AutoField id PK
        TextField data
        ForeignKey user_id FK
        ForeignKey content_type_id FK
        CharField object_id
        CharField parent_object_id
        DateTimeField last_updated_at
    }
```
