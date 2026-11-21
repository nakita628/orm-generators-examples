```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    wagtailcore_page ||--o{ wagtailforms_formsubmission : "(id) - (page_id)"
    wagtailcore_site |o--o{ wagtailredirects_redirect : "(id) - (site_id)"
    wagtailcore_page |o--o{ wagtailredirects_redirect : "(id) - (redirect_page_id)"
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
    auth_user ||--o{ wagtailadmin_editingsession : "(id) - (user_id)"
    django_content_type ||--o{ wagtailadmin_editingsession : "(id) - (content_type_id)"
    auth_user ||--o{ wagtailadmin_formstate : "(id) - (user_id)"
    django_content_type ||--o{ wagtailadmin_formstate : "(id) - (content_type_id)"
    wagtailcore_page ||--o{ wagtailcore_site : "(id) - (root_page_id)"
    auth_group ||--o{ wagtailcore_groupsitepermission : "(id) - (group_id)"
    wagtailcore_site ||--o{ wagtailcore_groupsitepermission : "(id) - (site_id)"
    auth_permission ||--o{ wagtailcore_groupsitepermission : "(id) - (permission_id)"
    auth_user ||--o{ wagtailcore_apitoken : "(id) - (user_id)"
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
    taggit_tag ||--o{ taggit_taggeditem : "(id) - (tag_id)"
    django_content_type ||--o{ taggit_taggeditem : "(id) - (content_type_id)"
    wagtailstreamforms_formsubmission ||--o{ wagtailstreamforms_formsubmissionfile : "(id) - (submission_id)"
    wagtailstreamforms_form ||--o{ wagtailstreamforms_formsubmission : "(id) - (form_id)"
    wagtailcore_site |o--o{ wagtailstreamforms_form : "(id) - (site_id)"
    wagtailcore_page |o--o{ wagtailstreamforms_form : "(id) - (post_redirect_page_id)"
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
    wagtailcore_apitoken {
        AutoField id PK
        ForeignKey user_id FK
        CharField name "A human-readable label to identify this token."
        CharField key_hash UK
        CharField prefix
        DateTimeField created
        DateTimeField revoked_at
        DateTimeField last_used_at
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
        DateTimeField latest_revision_created_at
        ForeignKey alias_of_id FK
        BooleanField show_in_menus "Whether a link to this page will appear in automatically generated menus"
        CharField seo_title "The name of the page displayed on search engine results as the clickable headline."
        TextField search_description "The descriptive text displayed underneath a headline in search engine results."
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
    wagtailstreamforms_formsubmissionfile {
        BigAutoField id PK
        ForeignKey submission_id FK
        CharField field
        FileField file
    }
    wagtailstreamforms_formsubmission {
        BigAutoField id PK
        TextField form_data
        ForeignKey form_id FK
        DateTimeField submit_time
    }
    wagtailstreamforms_form {
        BigAutoField id PK
        ForeignKey site_id FK
        CharField title
        SlugField slug UK "Used to identify the form in template tags"
        FormTemplateName template_name
        FormFieldsStreamField fields
        CharField submit_button_text
        CharField success_message "An optional success message to show when the form has been successfully submitted"
        CharField error_message "An optional error message to show when the form has validation errors"
        ForeignKey post_redirect_page_id FK "The page to redirect to after a successful submission"
        HookSelectField process_form_submission_hooks
    }
```
