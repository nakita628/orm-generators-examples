```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user |o--o| djblets_privacy_storedconsentdata : "(id) - (user_id)"
    django_site ||--o{ siteconfig_siteconfiguration : "(id) - (site_id)"
    auth_user |o--o{ oauth2_provider_application : "(id) - (user_id)"
    auth_user ||--o{ oauth2_provider_grant : "(id) - (user_id)"
    reviewboard_oauth_application ||--o{ oauth2_provider_grant : "(id) - (application_id)"
    auth_user |o--o{ oauth2_provider_accesstoken : "(id) - (user_id)"
    oauth2_provider_refreshtoken |o--o| oauth2_provider_accesstoken : "(id) - (source_refresh_token_id)"
    oauth2_provider_idtoken |o--o| oauth2_provider_accesstoken : "(id) - (id_token_id)"
    reviewboard_oauth_application |o--o{ oauth2_provider_accesstoken : "(id) - (application_id)"
    auth_user ||--o{ oauth2_provider_refreshtoken : "(id) - (user_id)"
    reviewboard_oauth_application ||--o{ oauth2_provider_refreshtoken : "(id) - (application_id)"
    oauth2_provider_accesstoken |o--o| oauth2_provider_refreshtoken : "(id) - (access_token_id)"
    auth_user |o--o{ oauth2_provider_idtoken : "(id) - (user_id)"
    reviewboard_oauth_application |o--o{ oauth2_provider_idtoken : "(id) - (application_id)"
    auth_user ||--o{ accounts_reviewrequestvisit : "(id) - (user_id)"
    reviews_reviewrequest ||--o{ accounts_reviewrequestvisit : "(id) - (review_request_id)"
    auth_user ||--o{ accounts_linkedaccount : "(id) - (user_id)"
    auth_user ||--o{ accounts_profile : "(id) - (user_id)"
    accounts_profile }o--o{ reviews_reviewrequest : "(starred_review_requests) - (starred_by)"
    accounts_profile }o--o{ reviews_group : "(starred_groups) - (starred_by)"
    auth_user ||--o{ accounts_localsiteprofile : "(id) - (user_id)"
    accounts_profile ||--o{ accounts_localsiteprofile : "(id) - (profile_id)"
    site_localsite |o--o{ accounts_localsiteprofile : "(id) - (local_site_id)"
    reviews_reviewrequest ||--o{ accounts_trophy : "(id) - (review_request_id)"
    site_localsite |o--o{ accounts_trophy : "(id) - (local_site_id)"
    auth_user ||--o{ accounts_trophy : "(id) - (user_id)"
    auth_user |o--o{ attachments_fileattachment : "(id) - (user_id)"
    site_localsite |o--o{ attachments_fileattachment : "(id) - (local_site_id)"
    scmtools_repository |o--o{ attachments_fileattachment : "(id) - (repository_id)"
    diffviewer_filediff |o--o{ attachments_fileattachment : "(id) - (added_in_filediff_id)"
    attachments_fileattachmenthistory |o--o{ attachments_fileattachment : "(id) - (attachment_history_id)"
    auth_user |o--o{ changedescs_changedescription : "(id) - (user_id)"
    diffviewer_diffsethistory |o--o{ diffviewer_diffset : "(id) - (history_id)"
    scmtools_repository ||--o{ diffviewer_diffset : "(id) - (repository_id)"
    diffviewer_diffset ||--o{ diffviewer_diffcommit : "(id) - (diffset_id)"
    diffviewer_diffset ||--o{ diffviewer_filediff : "(id) - (diffset_id)"
    diffviewer_diffcommit |o--o{ diffviewer_filediff : "(id) - (commit_id)"
    diffviewer_filediffdata |o--o{ diffviewer_filediff : "(binary_hash) - (diff_hash_id)"
    diffviewer_rawfilediffdata |o--o{ diffviewer_filediff : "(id) - (raw_diff_hash_id)"
    diffviewer_filediffdata |o--o{ diffviewer_filediff : "(binary_hash) - (parent_diff_hash_id)"
    diffviewer_rawfilediffdata |o--o{ diffviewer_filediff : "(id) - (raw_parent_diff_hash_id)"
    site_localsite |o--o{ hostingsvcs_hostingserviceaccount : "(id) - (local_site_id)"
    site_localsite |o--o{ integrations_integrationconfig : "(id) - (local_site_id)"
    site_localsite |o--o{ notifications_webhooktarget : "(id) - (local_site_id)"
    notifications_webhooktarget }o--o{ scmtools_repository : "(repositories) - (webhooks)"
    auth_user |o--o{ reviewboard_oauth_application : "(id) - (user_id)"
    auth_user |o--o{ reviewboard_oauth_application : "(id) - (original_user_id)"
    site_localsite |o--o{ reviewboard_oauth_application : "(id) - (local_site_id)"
    site_localsite |o--o{ reviews_group : "(id) - (local_site_id)"
    reviews_group }o--o{ auth_user : "(users) - (review_groups)"
    site_localsite |o--o{ reviews_defaultreviewer : "(id) - (local_site_id)"
    reviews_defaultreviewer }o--o{ scmtools_repository : "(repository) - (defaultreviewer_set)"
    reviews_defaultreviewer }o--o{ reviews_group : "(groups) - (defaultreviewer_set)"
    reviews_defaultreviewer }o--o{ auth_user : "(people) - (default_review_paths)"
    reviews_comment |o--o{ reviews_comment : "(id) - (reply_to_id)"
    diffviewer_filediff ||--o{ reviews_comment : "(id) - (filediff_id)"
    diffviewer_filediff |o--o{ reviews_comment : "(id) - (interfilediff_id)"
    reviews_fileattachmentcomment |o--o{ reviews_fileattachmentcomment : "(id) - (reply_to_id)"
    attachments_fileattachment ||--o{ reviews_fileattachmentcomment : "(id) - (file_attachment_id)"
    attachments_fileattachment |o--o{ reviews_fileattachmentcomment : "(id) - (diff_against_file_attachment_id)"
    reviews_generalcomment |o--o{ reviews_generalcomment : "(id) - (reply_to_id)"
    auth_user ||--o{ reviews_reviewrequest : "(id) - (submitter_id)"
    scmtools_repository |o--o{ reviews_reviewrequest : "(id) - (repository_id)"
    diffviewer_diffsethistory |o--o{ reviews_reviewrequest : "(id) - (diffset_history_id)"
    site_localsite |o--o{ reviews_reviewrequest : "(id) - (local_site_id)"
    reviews_reviewrequest }o--o{ reviews_group : "(target_groups) - (review_requests)"
    reviews_reviewrequest }o--o{ auth_user : "(target_people) - (directed_review_requests)"
    reviews_reviewrequest }o--o{ reviews_screenshot : "(screenshots) - (review_request)"
    reviews_reviewrequest }o--o{ reviews_screenshot : "(inactive_screenshots) - (inactive_review_request)"
    reviews_reviewrequest }o--o{ attachments_fileattachment : "(file_attachments) - (review_request)"
    reviews_reviewrequest }o--o{ attachments_fileattachment : "(inactive_file_attachments) - (inactive_review_request)"
    reviews_reviewrequest }o--o{ attachments_fileattachmenthistory : "(file_attachment_histories) - (review_request)"
    reviews_reviewrequest }o--o{ changedescs_changedescription : "(changedescs) - (review_request)"
    reviews_reviewrequest }o--o{ reviews_reviewrequest : "(depends_on) - (blocks)"
    reviews_screenshotcomment |o--o{ reviews_screenshotcomment : "(id) - (reply_to_id)"
    reviews_screenshot ||--o{ reviews_screenshotcomment : "(id) - (screenshot_id)"
    reviews_reviewrequest ||--o{ reviews_review : "(id) - (review_request_id)"
    auth_user ||--o{ reviews_review : "(id) - (user_id)"
    reviews_review |o--o{ reviews_review : "(id) - (base_reply_to_id)"
    reviews_review |o--o{ reviews_review : "(id) - (body_top_reply_to_id)"
    reviews_review |o--o{ reviews_review : "(id) - (body_bottom_reply_to_id)"
    diffviewer_diffset |o--o{ reviews_review : "(id) - (reviewed_diffset_id)"
    reviews_review }o--o{ reviews_comment : "(comments) - (review)"
    reviews_review }o--o{ reviews_screenshotcomment : "(screenshot_comments) - (review)"
    reviews_review }o--o{ reviews_fileattachmentcomment : "(file_attachment_comments) - (review)"
    reviews_review }o--o{ reviews_generalcomment : "(general_comments) - (review)"
    auth_user |o--o{ reviews_reviewrequestdraft : "(id) - (owner_id)"
    reviews_reviewrequest ||--o{ reviews_reviewrequestdraft : "(id) - (review_request_id)"
    diffviewer_diffset |o--o{ reviews_reviewrequestdraft : "(id) - (diffset_id)"
    changedescs_changedescription |o--o{ reviews_reviewrequestdraft : "(id) - (changedesc_id)"
    reviews_reviewrequestdraft }o--o{ reviews_group : "(target_groups) - (drafts)"
    reviews_reviewrequestdraft }o--o{ auth_user : "(target_people) - (directed_drafts)"
    reviews_reviewrequestdraft }o--o{ reviews_screenshot : "(screenshots) - (drafts)"
    reviews_reviewrequestdraft }o--o{ reviews_screenshot : "(inactive_screenshots) - (inactive_drafts)"
    reviews_reviewrequestdraft }o--o{ attachments_fileattachment : "(file_attachments) - (drafts)"
    reviews_reviewrequestdraft }o--o{ attachments_fileattachment : "(inactive_file_attachments) - (inactive_drafts)"
    reviews_reviewrequestdraft }o--o{ reviews_reviewrequest : "(depends_on) - (draft_blocks)"
    auth_user |o--o{ reviews_statusupdate : "(id) - (user_id)"
    reviews_reviewrequest ||--o{ reviews_statusupdate : "(id) - (review_request_id)"
    changedescs_changedescription |o--o{ reviews_statusupdate : "(id) - (change_description_id)"
    reviews_review |o--o| reviews_statusupdate : "(id) - (review_id)"
    scmtools_tool ||--o{ scmtools_repository : "(id) - (tool_id)"
    hostingsvcs_hostingserviceaccount |o--o{ scmtools_repository : "(id) - (hosting_account_id)"
    site_localsite |o--o{ scmtools_repository : "(id) - (local_site_id)"
    scmtools_repository }o--o{ auth_user : "(users) - (repositories)"
    scmtools_repository }o--o{ reviews_group : "(review_groups) - (repositories)"
    site_localsite }o--o{ auth_user : "(users) - (local_site)"
    site_localsite }o--o{ auth_user : "(admins) - (local_site_admins)"
    auth_user ||--o{ webapi_webapitoken : "(id) - (user_id)"
    site_localsite |o--o{ webapi_webapitoken : "(id) - (local_site_id)"
    django_project_version ||--o{ django_evolution : "(id) - (version_id)"
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
    django_site {
        AutoField id PK
        CharField domain UK
        CharField name
    }
    django_session {
        CharField session_key PK
        TextField session_data
        DateTimeField expire_date
    }
    extensions_registeredextension {
        AutoField id PK
        CharField class_name UK
        CharField name
        BooleanField enabled
        BooleanField installed
        JSONField settings
    }
    djblets_privacy_storedconsentdata {
        AutoField id PK
        OneToOneField user_id FK, UK "The user whose consent decisions are being stored."
        CharField audit_identifier "An identifier for locating the audit record as needed. The value must be able to be compared against information from the original user, but may be erased or encrypted at any point."
        DateTimeField time_added "The date and time when the stored consent data was first added to the database."
        DateTimeField last_updated "The date and time when the stored consent data was last updated."
        JSONField consent_grants "The consents granted or denied by the user."
        JSONField audit_trail "A record of all consent changes made by the user."
    }
    siteconfig_siteconfiguration {
        AutoField id PK
        ForeignKey site_id FK
        CharField version
        JSONField settings
    }
    oauth2_provider_application {
        BigAutoField id PK
        CharField client_id UK
        ForeignKey user_id FK
        TextField redirect_uris "Allowed URIs list, space separated"
        Oauth2ProviderApplicationClientType client_type
        Oauth2ProviderApplicationAuthorizationGrantType authorization_grant_type
        CharField client_secret
        CharField name
        BooleanField skip_authorization
        DateTimeField created
        DateTimeField updated
        Oauth2ProviderApplicationAlgorithm algorithm
    }
    oauth2_provider_grant {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField code UK
        ForeignKey application_id FK
        DateTimeField expires
        TextField redirect_uri
        TextField scope
        DateTimeField created
        DateTimeField updated
        CharField code_challenge
        GrantCodeChallengeMethod code_challenge_method
        CharField nonce
        TextField claims
    }
    oauth2_provider_accesstoken {
        BigAutoField id PK
        ForeignKey user_id FK
        OneToOneField source_refresh_token_id FK, UK
        CharField token UK
        OneToOneField id_token_id FK, UK
        ForeignKey application_id FK
        DateTimeField expires
        TextField scope
        DateTimeField created
        DateTimeField updated
    }
    oauth2_provider_refreshtoken {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField token
        ForeignKey application_id FK
        OneToOneField access_token_id FK, UK
        DateTimeField created
        DateTimeField updated
        DateTimeField revoked
    }
    oauth2_provider_idtoken {
        BigAutoField id PK
        ForeignKey user_id FK
        UUIDField jti UK
        ForeignKey application_id FK
        DateTimeField expires
        TextField scope
        DateTimeField created
        DateTimeField updated
    }
    accounts_reviewrequestvisit {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey review_request_id FK
        DateTimeField timestamp
        ReviewRequestVisitVisibility visibility
    }
    accounts_linkedaccount {
        AutoField id PK
        ForeignKey user_id FK
        CharField service_id
        CharField service_user_id
        JSONField service_data "Used to store private data such as session keys."
        JSONField extra_data
    }
    accounts_profile {
        AutoField id PK
        ForeignKey user_id FK, UK
        BooleanField first_time_setup_done "Indicates whether the user has already gone through the first time setup process by saving their user preferences."
        BooleanField should_send_email "Indicates whether the user wishes to receive emails."
        BooleanField should_send_own_updates "Indicates whether the user wishes to receive emails about their own activity."
        BooleanField collapsed_diffs "Indicates whether diffs should be shown in their collapsed state by default."
        BooleanField wordwrapped_diffs "This field is unused and will be removed in a future version."
        BooleanField syntax_highlighting "Indicates whether the user wishes to see syntax highlighting in the diffs."
        BooleanField is_private "Indicates whether the user wishes to keep his/her profile private."
        BooleanField open_an_issue "Indicates whether the user wishes to default to opening an issue or not."
        BooleanField default_use_rich_text "Indicates whether new posts or comments should default to being in Markdown format."
        BooleanField show_closed
        CharField sort_review_request_columns
        CharField sort_dashboard_columns
        CharField sort_submitter_columns
        CharField sort_group_columns
        CharField review_request_columns
        CharField dashboard_columns
        CharField submitter_columns
        CharField group_columns
        ProfileTimezone timezone
        JSONField settings
        JSONField extra_data
    }
    accounts_localsiteprofile {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey profile_id FK
        ForeignKey local_site_id FK
        JSONField permissions
        CounterField direct_incoming_request_count
        CounterField total_incoming_request_count
        CounterField pending_outgoing_request_count
        CounterField total_outgoing_request_count
        CounterField starred_public_request_count
    }
    accounts_trophy {
        AutoField id PK
        CharField category
        DateTimeField received_date
        ForeignKey review_request_id FK
        ForeignKey local_site_id FK
        ForeignKey user_id FK
    }
    attachments_fileattachmenthistory {
        AutoField id PK
        IntegerField display_position
        RelationCounterField latest_revision
    }
    attachments_fileattachment {
        AutoField id PK
        CharField caption
        CharField draft_caption
        CharField orig_filename
        ForeignKey user_id FK
        ForeignKey local_site_id FK
        CharField uuid
        FileField file
        CharField mimetype
        JSONField extra_data
        CharField repo_path
        CharField repo_revision
        ForeignKey repository_id FK
        ForeignKey added_in_filediff_id FK
        ForeignKey attachment_history_id FK
        IntegerField attachment_revision
    }
    changedescs_changedescription {
        AutoField id PK
        ForeignKey user_id FK
        DateTimeField timestamp
        BooleanField public
        TextField text
        BooleanField rich_text
        JSONField fields_changed
    }
    diffviewer_diffset {
        AutoField id PK
        CharField name
        IntegerField revision
        DateTimeField timestamp
        CharField basedir
        ForeignKey history_id FK
        ForeignKey repository_id FK
        IntegerField diffcompat "The diff generator compatibility version to use. This can and should be ignored."
        CharField base_commit_id "The ID/revision this change is built upon."
        RelationCounterField commit_count
        JSONField extra_data
    }
    diffviewer_diffcommit {
        AutoField id PK
        CharField filename "The original file name of the diff."
        ForeignKey diffset_id FK
        CharField author_name "The name of the commit author."
        CharField author_email "The e-mail address of the commit author."
        DateTimeField author_date_utc "The date the commit was authored in UTC."
        IntegerField author_date_offset "The author's UTC offset."
        CharField committer_name "The name of the committer (if applicable)."
        CharField committer_email "The e-mail address of the committer (if applicable)."
        DateTimeField committer_date_utc "The date the commit was committed in UTC (if applicable)."
        IntegerField committer_date_offset "The committer's UTC offset (if applicable)."
        TextField commit_message "The commit message."
        CharField commit_id "The unique identifier of the commit."
        CharField parent_id "The unique identifier of the parent commit."
        DateTimeField last_modified
        JSONField extra_data
    }
    diffviewer_diffsethistory {
        AutoField id PK
        CharField name
        DateTimeField timestamp
        DateTimeField last_diff_updated
        JSONField extra_data
    }
    diffviewer_filediffdata {
        CharField binary_hash PK
        Base64Field binary
        JSONField extra_data
    }
    diffviewer_rawfilediffdata {
        AutoField id PK
        CharField binary_hash UK
        BinaryField binary
        RawFileDiffDataCompression compression
        JSONField extra_data
    }
    diffviewer_filediff {
        AutoField id PK
        ForeignKey diffset_id FK
        ForeignKey commit_id FK
        CharField source_file
        CharField dest_file
        CharField source_revision
        CharField dest_detail
        BooleanField binary
        FileDiffStatus status
        Base64Field diff_base64
        ForeignKey diff_hash_id FK
        ForeignKey raw_diff_hash_id FK
        Base64Field parent_diff_base64
        ForeignKey parent_diff_hash_id FK
        ForeignKey raw_parent_diff_hash_id FK
        JSONField extra_data
    }
    hostingsvcs_hostingserviceaccount {
        AutoField id PK
        CharField service_name "The ID of the hosting service that is associated with this account."
        CharField hosting_url "The hosting URL."
        CharField username "The username of the account on the hosting service."
        JSONField data "Account data specific to the hosting service. This should generally not be changed."
        BooleanField visible "Whether this account shows up as an option when configuring a repository."
        ForeignKey local_site_id FK "The LocalSite to associate with this account."
    }
    integrations_integrationconfig {
        AutoField id PK
        CharField integration_id
        DateTimeField time_added
        DateTimeField last_updated
        CharField name
        BooleanField enabled
        JSONField settings
        JSONField extra_data
        ForeignKey local_site_id FK
    }
    notifications_webhooktarget {
        AutoField id PK
        BooleanField enabled
        CommaSeparatedValuesField events "Select any or all events that should trigger this Webhook."
        URLField url "When the event is triggered, HTTP requests will be made against this URL."
        WebHookTargetEncoding encoding "Payload contents will be encoded in this format."
        BooleanField use_custom_content
        TextField custom_content "You can override what is sent to the URL above. If left blank, the default payload will be sent."
        CharField secret "If specified, the HMAC digest for the Webhook payload will be signed with the given secret."
        WebHookTargetApplyTo apply_to
        ForeignKey local_site_id FK "If set, this Webhook will be limited to this site."
        JSONField extra_data "Extra JSON data that can be tied to this Webhook registration. It will not be sent with the Webhook request."
    }
    reviewboard_oauth_application {
        BigAutoField id PK
        CharField client_id UK
        ForeignKey user_id FK
        TextField redirect_uris "Allowed URIs list, space separated"
        OauthApplicationClientType client_type
        OauthApplicationAuthorizationGrantType authorization_grant_type
        CharField client_secret
        CharField name
        BooleanField skip_authorization
        DateTimeField created
        DateTimeField updated
        OauthApplicationAlgorithm algorithm
        BooleanField enabled "Whether or not this application can be used to authenticate with Review Board."
        ForeignKey original_user_id FK "The original owner of this application."
        ForeignKey local_site_id FK "An optional Local Site to limit this application to.<br>If specified, only users with access to the Local Site will be able to use the application."
        JSONField extra_data
    }
    reviews_group {
        AutoField id PK
        SlugField name
        CharField display_name
        CharField mailing_list "The mailing list review requests and discussions are sent to."
        BooleanField email_list_only "If a mailing list is specified and this option is checked, group members will not be individually included on e-mails, and only the mailing list will be used. This is highly recommended for large groups."
        ForeignKey local_site_id FK
        BooleanField is_default_group "If a local site is set, this will automatically add users to this group when those users are added to the local site. If there is no local site, users will be automatically added to this group when they are registered."
        CounterField incoming_request_count
        BooleanField invite_only "If checked, only the users listed below will be able to view review requests sent to this group."
        BooleanField visible
        JSONField extra_data
    }
    reviews_defaultreviewer {
        AutoField id PK
        CharField name
        CharField file_regex "File paths are matched against this regular expression to determine if these reviewers should be added."
        ForeignKey local_site_id FK
    }
    reviews_comment {
        AutoField id PK
        BooleanField issue_opened
        CommentIssueStatus issue_status
        ForeignKey reply_to_id FK
        DateTimeField timestamp
        TextField text
        BooleanField rich_text
        JSONField extra_data
        ForeignKey filediff_id FK
        ForeignKey interfilediff_id FK
        PositiveIntegerField first_line
        PositiveIntegerField num_lines
    }
    reviews_fileattachmentcomment {
        AutoField id PK
        BooleanField issue_opened
        FileAttachmentCommentIssueStatus issue_status
        ForeignKey reply_to_id FK
        DateTimeField timestamp
        TextField text
        BooleanField rich_text
        JSONField extra_data
        ForeignKey file_attachment_id FK
        ForeignKey diff_against_file_attachment_id FK
    }
    reviews_generalcomment {
        AutoField id PK
        BooleanField issue_opened
        GeneralCommentIssueStatus issue_status
        ForeignKey reply_to_id FK
        DateTimeField timestamp
        TextField text
        BooleanField rich_text
        JSONField extra_data
    }
    reviews_screenshot {
        AutoField id PK
        CharField caption
        CharField draft_caption
        ImageField image
    }
    reviews_reviewrequest {
        AutoField id PK
        TextField description
        BooleanField description_rich_text
        TextField testing_done
        BooleanField testing_done_rich_text
        CharField bugs_closed
        CharField branch
        CharField commit_id
        JSONField extra_data
        BooleanField rich_text
        CharField summary
        ForeignKey submitter_id FK
        DateTimeField time_added
        ModificationTimestampField last_updated
        ReviewRequestStatus status
        BooleanField public
        PositiveIntegerField changenum
        ForeignKey repository_id FK
        CharField email_message_id
        DateTimeField time_emailed
        ForeignKey diffset_history_id FK
        DateTimeField last_review_timestamp
        CounterField shipit_count
        CounterField issue_open_count
        CounterField issue_resolved_count
        CounterField issue_dropped_count
        CounterField issue_verifying_count
        RelationCounterField screenshots_count
        RelationCounterField inactive_screenshots_count
        RelationCounterField file_attachments_count
        RelationCounterField inactive_file_attachments_count
        ForeignKey local_site_id FK
        IntegerField local_id
    }
    reviews_screenshotcomment {
        AutoField id PK
        BooleanField issue_opened
        ScreenshotCommentIssueStatus issue_status
        ForeignKey reply_to_id FK
        DateTimeField timestamp
        TextField text
        BooleanField rich_text
        JSONField extra_data
        ForeignKey screenshot_id FK
        PositiveSmallIntegerField x
        PositiveSmallIntegerField y
        PositiveSmallIntegerField w
        PositiveSmallIntegerField h
    }
    reviews_review {
        AutoField id PK
        ForeignKey review_request_id FK
        ForeignKey user_id FK
        DateTimeField timestamp
        BooleanField public
        BooleanField ship_it "Indicates whether the reviewer thinks this code is ready to ship."
        ForeignKey base_reply_to_id FK "The top-most review in the discussion thread for this review reply."
        CharField email_message_id
        DateTimeField time_emailed
        TextField body_top "The review text shown above the diff and screenshot comments."
        BooleanField body_top_rich_text
        TextField body_bottom "The review text shown below the diff and screenshot comments."
        BooleanField body_bottom_rich_text
        ForeignKey body_top_reply_to_id FK "The review that the body (top) field is in reply to."
        ForeignKey body_bottom_reply_to_id FK "The review that the body (bottom) field is in reply to."
        JSONField extra_data
        BooleanField rich_text
        ForeignKey reviewed_diffset_id FK "This field is unused and will be removed in a future version."
    }
    reviews_reviewrequestdraft {
        AutoField id PK
        TextField description
        BooleanField description_rich_text
        TextField testing_done
        BooleanField testing_done_rich_text
        CharField bugs_closed
        CharField branch
        CharField commit_id
        JSONField extra_data
        BooleanField rich_text
        CharField summary
        ForeignKey owner_id FK
        ForeignKey review_request_id FK, UK
        ModificationTimestampField last_updated
        ForeignKey diffset_id FK
        ForeignKey changedesc_id FK
        RelationCounterField screenshots_count
        RelationCounterField inactive_screenshots_count
        RelationCounterField file_attachments_count
        RelationCounterField inactive_file_attachments_count
    }
    reviews_statusupdate {
        AutoField id PK
        CharField service_id
        ForeignKey user_id FK
        DateTimeField timestamp
        CharField summary
        CharField description
        URLField url
        CharField url_text
        StatusUpdateState state
        ForeignKey review_request_id FK
        ForeignKey change_description_id FK
        OneToOneField review_id FK, UK
        JSONField extra_data
        IntegerField timeout
    }
    scmtools_tool {
        AutoField id PK
        CharField name UK
        CharField class_name UK
    }
    scmtools_repository {
        AutoField id PK
        CharField name
        CharField path
        CharField mirror_path
        CharField raw_file_url
        CharField username
        CharField password
        JSONField extra_data
        ForeignKey tool_id FK
        CharField scmtool_id
        ForeignKey hosting_account_id FK
        CharField bug_tracker "This should be the full path to a bug in the bug tracker for this repository, using '%s' in place of the bug ID."
        CharField encoding "The encoding used for files in this repository. This is an advanced setting and should only be used if you're sure you need it."
        BooleanField visible "Use this to control whether or not a repository is shown when creating new review requests. Existing review requests are unaffected."
        BooleanField archived "Archived repositories do not show up in lists of repositories, and aren't open to new review requests."
        DateTimeField archived_timestamp
        ForeignKey local_site_id FK
        BooleanField public "Review requests and files on public repositories are visible to anyone. Private repositories must explicitly list the users and groups that can access them."
        CharField hooks_uuid "Unique identifier used for validating incoming webhooks."
    }
    site_localsite {
        AutoField id PK
        SlugField name UK
        BooleanField public "Allow people outside the team to access and post review requests and reviews."
        JSONField extra_data
    }
    webapi_webapitoken {
        AutoField id PK
        ForeignKey user_id FK "The user that owns the token."
        CharField token UK "The access token."
        CharField token_generator_id "The token generator that generated the token."
        DateTimeField time_added "The date and time when the token was first added to the database."
        ModificationTimestampField last_updated "The date and time when the token was last updated."
        DateTimeField last_used "The date and time when the token was last used for authentication."
        DateTimeField expires "An optional field for the date and time that the token will expire. The token will be invalid and unusable for authentication after this point."
        BooleanField valid "Whether the token is currently valid."
        DateTimeField invalid_date "The date and time at which the token became invalid."
        TextField invalid_reason "A message indicating why the token is no longer valid."
        TextField note "A message describing the token."
        JSONField policy "The policy document describing what this token can access in the API. If empty, this provides full access."
        JSONField extra_data
        ForeignKey local_site_id FK
    }
    django_project_version {
        AutoField id PK
        SignatureField signature
        DateTimeField when
    }
    django_evolution {
        AutoField id PK
        ForeignKey version_id FK
        CharField app_label
        CharField label
    }
```
