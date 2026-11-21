```mermaid
erDiagram
    misago_users_rank }o--o{ misago_acl_role : "(roles) - (rank_set)"
    misago_users_user ||--o| misago_users_online : "(id) - (user_id)"
    misago_users_rank |o--o{ misago_users_user : "(id) - (rank_id)"
    misago_users_user }o--o{ auth_group : "(groups) - (user_set)"
    misago_users_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    misago_users_user }o--o{ misago_acl_role : "(roles) - (user_set)"
    misago_users_user }o--o{ misago_users_user : "(follows) - (followed_by)"
    misago_users_user }o--o{ misago_users_user : "(blocks) - (blocked_by)"
    misago_users_user ||--o{ misago_users_usernamechange : "(id) - (user_id)"
    misago_users_user |o--o{ misago_users_usernamechange : "(id) - (changed_by_id)"
    misago_users_user ||--o{ misago_users_activityranking : "(id) - (user_id)"
    misago_users_user ||--o{ misago_users_avatar : "(id) - (user_id)"
    misago_users_user ||--o{ misago_users_audittrail : "(id) - (user_id)"
    django_content_type ||--o{ misago_users_audittrail : "(id) - (content_type_id)"
    misago_users_user ||--o| misago_users_bancache : "(id) - (user_id)"
    misago_users_ban |o--o{ misago_users_bancache : "(id) - (ban_id)"
    misago_users_user ||--o{ misago_users_datadownload : "(id) - (user_id)"
    misago_users_user |o--o{ misago_users_datadownload : "(id) - (requester_id)"
    misago_users_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    misago_users_user ||--o{ social_auth_usersocialauth : "(id) - (user_id)"
    misago_themes_theme |o--o{ misago_themes_theme : "(id) - (parent_id)"
    misago_themes_theme ||--o{ misago_themes_css : "(id) - (theme_id)"
    misago_themes_theme ||--o{ misago_themes_media : "(id) - (theme_id)"
    misago_users_user |o--o{ misago_legal_agreement : "(id) - (created_by_id)"
    misago_users_user |o--o{ misago_legal_agreement : "(id) - (last_modified_by_id)"
    misago_users_user ||--o{ misago_legal_useragreement : "(id) - (user_id)"
    misago_legal_agreement ||--o{ misago_legal_useragreement : "(id) - (agreement_id)"
    misago_users_user ||--o{ misago_notifications_notification : "(id) - (user_id)"
    misago_users_user |o--o{ misago_notifications_notification : "(id) - (actor_id)"
    misago_categories_category |o--o{ misago_notifications_notification : "(id) - (category_id)"
    misago_threads_thread |o--o{ misago_notifications_notification : "(id) - (thread_id)"
    misago_threads_post |o--o{ misago_notifications_notification : "(id) - (post_id)"
    misago_users_user ||--o{ misago_notifications_watchedthread : "(id) - (user_id)"
    misago_categories_category ||--o{ misago_notifications_watchedthread : "(id) - (category_id)"
    misago_threads_thread ||--o{ misago_notifications_watchedthread : "(id) - (thread_id)"
    misago_categories_category |o--o{ misago_categories_category : "(id) - (parent_id)"
    misago_threads_thread |o--o{ misago_categories_category : "(id) - (last_thread_id)"
    misago_users_user |o--o{ misago_categories_category : "(id) - (last_poster_id)"
    misago_categories_category |o--o{ misago_categories_category : "(id) - (archive_pruned_in_id)"
    misago_acl_role ||--o{ misago_categories_rolecategoryacl : "(id) - (role_id)"
    misago_categories_category ||--o{ misago_categories_rolecategoryacl : "(id) - (category_id)"
    misago_categories_categoryrole ||--o{ misago_categories_rolecategoryacl : "(id) - (category_role_id)"
    misago_categories_category ||--o{ misago_threads_post : "(id) - (category_id)"
    misago_threads_thread ||--o{ misago_threads_post : "(id) - (thread_id)"
    misago_users_user |o--o{ misago_threads_post : "(id) - (poster_id)"
    misago_users_user |o--o{ misago_threads_post : "(id) - (last_editor_id)"
    misago_users_user |o--o{ misago_threads_post : "(id) - (hidden_by_id)"
    misago_threads_post }o--o{ misago_users_user : "(mentions) - (mention_set)"
    misago_categories_category ||--o{ misago_threads_postedit : "(id) - (category_id)"
    misago_threads_thread ||--o{ misago_threads_postedit : "(id) - (thread_id)"
    misago_threads_post ||--o{ misago_threads_postedit : "(id) - (post_id)"
    misago_users_user |o--o{ misago_threads_postedit : "(id) - (editor_id)"
    misago_categories_category ||--o{ misago_threads_postlike : "(id) - (category_id)"
    misago_threads_thread ||--o{ misago_threads_postlike : "(id) - (thread_id)"
    misago_threads_post ||--o{ misago_threads_postlike : "(id) - (post_id)"
    misago_users_user |o--o{ misago_threads_postlike : "(id) - (liker_id)"
    misago_categories_category ||--o{ misago_threads_thread : "(id) - (category_id)"
    misago_threads_post |o--o{ misago_threads_thread : "(id) - (first_post_id)"
    misago_users_user |o--o{ misago_threads_thread : "(id) - (starter_id)"
    misago_threads_post |o--o{ misago_threads_thread : "(id) - (last_post_id)"
    misago_users_user |o--o{ misago_threads_thread : "(id) - (last_poster_id)"
    misago_threads_post |o--o{ misago_threads_thread : "(id) - (best_answer_id)"
    misago_users_user |o--o{ misago_threads_thread : "(id) - (best_answer_marked_by_id)"
    misago_threads_thread ||--o{ misago_threads_threadparticipant : "(id) - (thread_id)"
    misago_users_user ||--o{ misago_threads_threadparticipant : "(id) - (user_id)"
    misago_users_user ||--o{ misago_threads_subscription : "(id) - (user_id)"
    misago_threads_thread ||--o{ misago_threads_subscription : "(id) - (thread_id)"
    misago_categories_category ||--o{ misago_threads_subscription : "(id) - (category_id)"
    misago_threads_attachmenttype }o--o{ misago_acl_role : "(limit_uploads_to) - (_misago_threads_attachmenttype_limit_uploads_to_+)"
    misago_threads_attachmenttype }o--o{ misago_acl_role : "(limit_downloads_to) - (_misago_threads_attachmenttype_limit_downloads_to_+)"
    misago_threads_attachmenttype ||--o{ misago_threads_attachment : "(id) - (filetype_id)"
    misago_threads_post |o--o{ misago_threads_attachment : "(id) - (post_id)"
    misago_users_user |o--o{ misago_threads_attachment : "(id) - (uploader_id)"
    misago_categories_category ||--o{ misago_threads_poll : "(id) - (category_id)"
    misago_threads_thread ||--o| misago_threads_poll : "(id) - (thread_id)"
    misago_users_user |o--o{ misago_threads_poll : "(id) - (poster_id)"
    misago_categories_category ||--o{ misago_threads_pollvote : "(id) - (category_id)"
    misago_threads_thread ||--o{ misago_threads_pollvote : "(id) - (thread_id)"
    misago_threads_poll ||--o{ misago_threads_pollvote : "(id) - (poll_id)"
    misago_users_user |o--o{ misago_threads_pollvote : "(id) - (voter_id)"
    misago_users_user ||--o{ misago_readtracker_postread : "(id) - (user_id)"
    misago_categories_category ||--o{ misago_readtracker_postread : "(id) - (category_id)"
    misago_threads_thread ||--o{ misago_readtracker_postread : "(id) - (thread_id)"
    misago_threads_post ||--o{ misago_readtracker_postread : "(id) - (post_id)"
    misago_users_user ||--o{ misago_oauth2_subject : "(id) - (user_id)"
    misago_users_rank {
        AutoField id PK
        CharField name
        CharField slug UK
        TextField description
        CharField title
        CharField css_class
        BooleanField is_default
        BooleanField is_tab
        IntegerField order
    }
    misago_users_online {
        OneToOneField user_id PK
        DateTimeField last_click
    }
    misago_users_user {
        AutoField id PK
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        JSONField plugin_data
        CharField username
        CharField slug UK
        EmailField email
        CharField email_hash UK
        DateTimeField joined_on
        GenericIPAddressField joined_from_ip
        BooleanField is_hiding_presence
        ForeignKey rank_id FK
        CharField title
        PositiveIntegerField requires_activation
        BooleanField is_staff "Designates whether the user can log into admin sites."
        CharField acl_key
        BooleanField is_active "Designates whether this user should be treated as active. Unselect this instead of deleting accounts."
        TextField is_active_staff_message
        BooleanField is_deleting_account
        ImageField avatar_tmp
        ImageField avatar_src
        CharField avatar_crop
        JSONField avatars
        BooleanField is_avatar_locked
        TextField avatar_lock_user_message
        TextField avatar_lock_staff_message
        TextField signature
        TextField signature_parsed
        CharField signature_checksum
        BooleanField is_signature_locked
        TextField signature_lock_user_message
        TextField signature_lock_staff_message
        PositiveIntegerField followers
        PositiveIntegerField following
        UserLimitsPrivateThreadInvitesTo limits_private_thread_invites_to
        PositiveIntegerField unread_private_threads
        BooleanField sync_unread_private_threads
        UserSubscribeToStartedThreads subscribe_to_started_threads
        UserSubscribeToRepliedThreads subscribe_to_replied_threads
        PositiveIntegerField unread_notifications
        ThreadNotifications watch_started_threads
        ThreadNotifications watch_replied_threads
        ThreadNotifications watch_new_private_threads_by_followed
        ThreadNotifications watch_new_private_threads_by_other_users
        ThreadNotifications notify_new_private_threads_by_followed
        ThreadNotifications notify_new_private_threads_by_other_users
        PositiveIntegerField threads
        PositiveIntegerField posts
        DateTimeField last_posted_on
        HStoreField profile_fields
        ArrayField agreements
    }
    misago_users_usernamechange {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey changed_by_id FK
        CharField changed_by_username
        DateTimeField changed_on
        CharField new_username
        CharField old_username
    }
    misago_users_activityranking {
        AutoField id PK
        ForeignKey user_id FK
        PositiveIntegerField score
    }
    misago_users_avatar {
        AutoField id PK
        ForeignKey user_id FK
        PositiveIntegerField size
        ImageField image
    }
    misago_users_audittrail {
        AutoField id PK
        ForeignKey user_id FK
        DateTimeField created_on
        GenericIPAddressField ip_address
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
    }
    misago_users_avatargallery {
        AutoField id PK
        CharField gallery
        ImageField image
    }
    misago_users_ban {
        AutoField id PK
        BanCheckType check_type
        BooleanField registration_only
        CharField banned_value
        TextField user_message
        TextField staff_message
        DateTimeField expires_on
        BooleanField is_checked
    }
    misago_users_bancache {
        OneToOneField user_id PK
        ForeignKey ban_id FK
        CharField cache_version
        TextField user_message
        TextField staff_message
        DateTimeField expires_on
    }
    misago_users_datadownload {
        AutoField id PK
        ForeignKey user_id FK
        DataDownloadStatus status
        ForeignKey requester_id FK
        CharField requester_name
        DateTimeField requested_on
        DateTimeField expires_on
        FileField file
    }
    misago_users_deleteduser {
        AutoField id PK
        DateTimeField deleted_on
        DeletedUserDeletedBy deleted_by
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
    social_auth_usersocialauth {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField provider
        CharField uid
        JSONField extra_data
        DateTimeField created
        DateTimeField modified
    }
    social_auth_nonce {
        BigAutoField id PK
        CharField server_url
        IntegerField timestamp
        CharField salt
    }
    social_auth_association {
        BigAutoField id PK
        CharField server_url
        CharField handle
        CharField secret
        IntegerField issued
        IntegerField lifetime
        CharField assoc_type
    }
    social_auth_code {
        BigAutoField id PK
        EmailField email
        CharField code
        BooleanField verified
        DateTimeField timestamp
    }
    social_auth_partial {
        BigAutoField id PK
        CharField token
        PositiveSmallIntegerField next_step
        CharField backend
        JSONField data
        DateTimeField timestamp
    }
    misago_acl_role {
        AutoField id PK
        CharField name
        CharField special_role
        JSONField permissions
    }
    misago_cache_cacheversion {
        CharField cache PK
        CharField version
    }
    misago_conf_setting {
        AutoField id PK
        CharField setting UK
        CharField python_type
        TextField dry_value
        ImageField image
        PositiveIntegerField image_size
        PositiveIntegerField image_width
        PositiveIntegerField image_height
        BooleanField is_public
        BooleanField is_lazy
    }
    misago_icons_icon {
        AutoField id PK
        CharField type
        ImageField image
        PositiveIntegerField size
        PositiveIntegerField width
        PositiveIntegerField height
    }
    misago_themes_theme {
        AutoField id PK
        TreeForeignKey parent_id FK
        CharField name
        CharField dirname
        BooleanField is_default
        BooleanField is_active
        CharField version
        CharField author
        URLField url
        PositiveIntegerField lft
        PositiveIntegerField rght
        PositiveIntegerField tree_id
        PositiveIntegerField level
    }
    misago_themes_css {
        AutoField id PK
        ForeignKey theme_id FK
        CharField name
        URLField url
        FileField source_file
        CharField source_hash
        BooleanField source_needs_building
        FileField build_file
        CharField build_hash
        PositiveIntegerField size
        IntegerField order
        DateTimeField modified_on
    }
    misago_themes_media {
        AutoField id PK
        ForeignKey theme_id FK
        CharField name
        ImageField file
        CharField hash
        CharField type
        PositiveIntegerField width
        PositiveIntegerField height
        PositiveIntegerField size
        ImageField thumbnail
        DateTimeField modified_on
    }
    misago_legal_agreement {
        AutoField id PK
        AgreementType type
        CharField title
        URLField link
        TextField text
        BooleanField is_active
        DateTimeField created_on
        ForeignKey created_by_id FK
        CharField created_by_name
        DateTimeField last_modified_on
        ForeignKey last_modified_by_id FK
        CharField last_modified_by_name
    }
    misago_legal_useragreement {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey agreement_id FK
        DateTimeField accepted_on
    }
    misago_notifications_notification {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField verb
        BooleanField is_read
        ForeignKey actor_id FK
        CharField actor_name
        ForeignKey category_id FK
        ForeignKey thread_id FK
        CharField thread_title
        ForeignKey post_id FK
        DateTimeField created_at
    }
    misago_notifications_watchedthread {
        BigAutoField id PK
        ForeignKey user_id FK
        ForeignKey category_id FK
        ForeignKey thread_id FK
        BooleanField send_emails
        CharField secret
        DateTimeField created_at
        DateTimeField read_at
    }
    misago_categories_category {
        AutoField id PK
        JSONField plugin_data
        TreeForeignKey parent_id FK
        CharField special_role
        CharField name
        CharField slug
        CharField short_name
        CharField color
        TextField description
        BooleanField is_closed
        PositiveIntegerField threads
        PositiveIntegerField posts
        DateTimeField last_post_on
        ForeignKey last_thread_id FK
        CharField last_thread_title
        CharField last_thread_slug
        ForeignKey last_poster_id FK
        CharField last_poster_name
        CharField last_poster_slug
        BooleanField require_threads_approval
        BooleanField require_replies_approval
        BooleanField require_edits_approval
        PositiveIntegerField prune_started_after
        PositiveIntegerField prune_replied_after
        ForeignKey archive_pruned_in_id FK
        CharField css_class
        PositiveIntegerField lft
        PositiveIntegerField rght
        PositiveIntegerField tree_id
        PositiveIntegerField level
    }
    misago_categories_categoryrole {
        AutoField id PK
        CharField name
        CharField special_role
        JSONField permissions
    }
    misago_categories_rolecategoryacl {
        AutoField id PK
        ForeignKey role_id FK
        ForeignKey category_id FK
        ForeignKey category_role_id FK
    }
    misago_threads_post {
        AutoField id PK
        JSONField plugin_data
        ForeignKey category_id FK
        ForeignKey thread_id FK
        ForeignKey poster_id FK
        CharField poster_name
        TextField original
        TextField parsed
        CharField checksum
        JSONField attachments_cache
        DateTimeField posted_on
        DateTimeField updated_on
        DateTimeField hidden_on
        PositiveIntegerField edits
        ForeignKey last_editor_id FK
        CharField last_editor_name
        SlugField last_editor_slug
        ForeignKey hidden_by_id FK
        CharField hidden_by_name
        SlugField hidden_by_slug
        BooleanField has_reports
        BooleanField has_open_reports
        BooleanField is_unapproved
        BooleanField is_hidden
        BooleanField is_protected
        BooleanField is_event
        CharField event_type
        JSONField event_context
        PositiveIntegerField likes
        JSONField last_likes
        TextField search_document
        SearchVectorField search_vector
    }
    misago_threads_postedit {
        AutoField id PK
        ForeignKey category_id FK
        ForeignKey thread_id FK
        ForeignKey post_id FK
        DateTimeField edited_on
        ForeignKey editor_id FK
        CharField editor_name
        CharField editor_slug
        TextField edited_from
        TextField edited_to
    }
    misago_threads_postlike {
        AutoField id PK
        ForeignKey category_id FK
        ForeignKey thread_id FK
        ForeignKey post_id FK
        ForeignKey liker_id FK
        CharField liker_name
        CharField liker_slug
        DateTimeField liked_on
    }
    misago_threads_thread {
        AutoField id PK
        JSONField plugin_data
        ForeignKey category_id FK
        CharField title
        CharField slug
        PositiveIntegerField replies
        BooleanField has_events
        BooleanField has_poll
        BooleanField has_reported_posts
        BooleanField has_open_reports
        BooleanField has_unapproved_posts
        BooleanField has_hidden_posts
        DateTimeField started_on
        DateTimeField last_post_on
        ForeignKey first_post_id FK
        ForeignKey starter_id FK
        CharField starter_name
        CharField starter_slug
        ForeignKey last_post_id FK
        BooleanField last_post_is_event
        ForeignKey last_poster_id FK
        CharField last_poster_name
        CharField last_poster_slug
        PositiveIntegerField weight
        BooleanField is_unapproved
        BooleanField is_hidden
        BooleanField is_closed
        ForeignKey best_answer_id FK
        BooleanField best_answer_is_protected
        DateTimeField best_answer_marked_on
        ForeignKey best_answer_marked_by_id FK
        CharField best_answer_marked_by_name
        CharField best_answer_marked_by_slug
    }
    misago_threads_threadparticipant {
        AutoField id PK
        ForeignKey thread_id FK
        ForeignKey user_id FK
        BooleanField is_owner
    }
    misago_threads_subscription {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey thread_id FK
        ForeignKey category_id FK
        DateTimeField last_read_on
        BooleanField send_email
    }
    misago_threads_attachmenttype {
        AutoField id PK
        JSONField plugin_data
        CharField name
        CharField extensions
        CharField mimetypes
        PositiveIntegerField size_limit
        AttachmentTypeStatus status
    }
    misago_threads_attachment {
        AutoField id PK
        JSONField plugin_data
        CharField secret
        ForeignKey filetype_id FK
        ForeignKey post_id FK
        DateTimeField uploaded_on
        ForeignKey uploader_id FK
        CharField uploader_name
        CharField uploader_slug
        CharField filename
        PositiveIntegerField size
        ImageField thumbnail
        ImageField image
        FileField file
    }
    misago_threads_poll {
        AutoField id PK
        JSONField plugin_data
        ForeignKey category_id FK
        OneToOneField thread_id FK, UK
        ForeignKey poster_id FK
        CharField poster_name
        CharField poster_slug
        DateTimeField posted_on
        PositiveIntegerField length
        CharField question
        JSONField choices
        PositiveIntegerField allowed_choices
        BooleanField allow_revotes
        PositiveIntegerField votes
        BooleanField is_public
    }
    misago_threads_pollvote {
        AutoField id PK
        ForeignKey category_id FK
        ForeignKey thread_id FK
        ForeignKey poll_id FK
        ForeignKey voter_id FK
        CharField voter_name
        CharField voter_slug
        DateTimeField voted_on
        CharField choice_hash
    }
    misago_readtracker_postread {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey category_id FK
        ForeignKey thread_id FK
        ForeignKey post_id FK
        DateTimeField last_read_on
    }
    misago_oauth2_subject {
        CharField sub PK
        ForeignKey user_id FK
        DateTimeField created_on
        DateTimeField last_used_on
    }
    misago_socialauth_socialauthprovider {
        CharField provider PK
        CharField button_text
        CharField button_color
        JSONField settings
        BooleanField is_active
        IntegerField order
    }
    misago_menus_menuitem {
        AutoField id PK
        MenuItemMenu menu
        CharField title
        URLField url
        IntegerField order
        CharField css_class
        BooleanField target_blank
        CharField rel
    }
```
