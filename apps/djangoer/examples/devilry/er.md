```mermaid
erDiagram
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    devilry_account_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    devilry_account_user |o--o{ ievv_batchframework_batchoperation : "(id) - (started_by_id)"
    django_content_type |o--o{ ievv_batchframework_batchoperation : "(id) - (context_content_type_id)"
    devilry_account_user ||--o{ cradmin_temporaryfileuploadstore_temporaryfilecollection : "(id) - (user_id)"
    cradmin_temporaryfileuploadstore_temporaryfilecollection ||--o{ cradmin_temporaryfileuploadstore_temporaryfile : "(id) - (collection_id)"
    django_content_type ||--o{ cradmin_generic_token_with_metadata_generictokenwithmetadata : "(id) - (content_type_id)"
    core_subject }o--o{ devilry_account_user : "(admins) - (subject_set)"
    core_subject ||--o{ core_period : "(id) - (parentnode_id)"
    core_period }o--o{ devilry_account_user : "(admins) - (period_set)"
    core_period ||--o{ core_periodapplicationkeyvalue : "(id) - (period_id)"
    core_period ||--o{ core_relatedexaminer : "(id) - (period_id)"
    devilry_account_user ||--o{ core_relatedexaminer : "(id) - (user_id)"
    core_period ||--o{ core_relatedstudent : "(id) - (period_id)"
    devilry_account_user ||--o{ core_relatedstudent : "(id) - (user_id)"
    core_relatedstudent ||--o{ core_relatedstudentkeyvalue : "(id) - (relatedstudent_id)"
    core_period ||--o{ core_periodtag : "(id) - (period_id)"
    core_periodtag }o--o{ core_relatedexaminer : "(relatedexaminers) - (periodtag_set)"
    core_periodtag }o--o{ core_relatedstudent : "(relatedstudents) - (periodtag_set)"
    core_period ||--o{ core_assignment : "(id) - (parentnode_id)"
    core_assignment }o--o{ devilry_account_user : "(admins) - (assignment_set)"
    core_assignment ||--o| core_pointtogrademap : "(id) - (assignment_id)"
    core_pointtogrademap ||--o{ core_pointrangetograde : "(id) - (point_to_grade_map_id)"
    core_assignment ||--o{ core_assignmentgroup : "(id) - (parentnode_id)"
    ievv_batchframework_batchoperation |o--o{ core_assignmentgroup : "(id) - (batchoperation_id)"
    core_assignmentgroup |o--o{ core_assignmentgroup : "(id) - (copied_from_id)"
    core_assignmentgroup ||--o{ core_assignmentgrouptag : "(id) - (assignment_group_id)"
    core_assignmentgroup ||--o| core_assignmentgrouphistory : "(id) - (assignment_group_id)"
    devilry_account_user |o--o{ core_candidate : "(id) - (old_reference_not_in_use_student_id)"
    core_relatedstudent ||--o{ core_candidate : "(id) - (relatedstudent_id)"
    core_assignmentgroup ||--o{ core_candidate : "(id) - (assignment_group_id)"
    devilry_account_user ||--o| core_devilryuserprofile : "(id) - (user_id)"
    devilry_account_user |o--o{ core_assignmentgroup_examiners : "(id) - (old_reference_not_in_use_user_id)"
    core_assignmentgroup ||--o{ core_assignmentgroup_examiners : "(id) - (assignmentgroup_id)"
    core_relatedexaminer ||--o{ core_assignmentgroup_examiners : "(id) - (relatedexaminer_id)"
    core_assignmentgroup ||--o{ core_groupinvite : "(id) - (group_id)"
    devilry_account_user ||--o{ core_groupinvite : "(id) - (sent_by_id)"
    devilry_account_user ||--o{ core_groupinvite : "(id) - (sent_to_id)"
    core_assignmentgroup ||--o{ core_examinerassignmentgrouphistory : "(id) - (assignment_group_id)"
    devilry_account_user |o--o{ core_examinerassignmentgrouphistory : "(id) - (user_id)"
    core_assignmentgroup ||--o{ core_candidateassignmentgrouphistory : "(id) - (assignment_group_id)"
    devilry_account_user |o--o{ core_candidateassignmentgrouphistory : "(id) - (user_id)"
    devilry_account_user ||--o{ devilry_account_mergeduser : "(id) - (source_user_id)"
    devilry_account_user ||--o{ devilry_account_mergeduser : "(id) - (target_user_id)"
    devilry_account_user ||--o{ devilry_account_useremail : "(id) - (user_id)"
    devilry_account_user ||--o{ devilry_account_username : "(id) - (user_id)"
    devilry_account_permissiongroup ||--o{ devilry_account_permissiongroupuser : "(id) - (permissiongroup_id)"
    devilry_account_user ||--o{ devilry_account_permissiongroupuser : "(id) - (user_id)"
    devilry_account_permissiongroup ||--o{ devilry_account_periodpermissiongroup : "(id) - (permissiongroup_id)"
    core_period ||--o{ devilry_account_periodpermissiongroup : "(id) - (period_id)"
    devilry_account_permissiongroup ||--o{ devilry_account_subjectpermissiongroup : "(id) - (permissiongroup_id)"
    core_subject ||--o{ devilry_account_subjectpermissiongroup : "(id) - (subject_id)"
    devilry_account_user ||--o{ devilry_account_perioduserguidelineacceptance : "(id) - (user_id)"
    core_period ||--o{ devilry_account_perioduserguidelineacceptance : "(id) - (period_id)"
    django_content_type ||--o{ devilry_compressionutil_compressedarchivemeta : "(id) - (content_type_id)"
    devilry_account_user |o--o{ devilry_compressionutil_compressedarchivemeta : "(id) - (created_by_id)"
    core_assignmentgroup ||--o{ devilry_group_feedbackset : "(id) - (group_id)"
    devilry_account_user |o--o{ devilry_group_feedbackset : "(id) - (created_by_id)"
    devilry_account_user |o--o{ devilry_group_feedbackset : "(id) - (last_updated_by_id)"
    devilry_account_user |o--o{ devilry_group_feedbackset : "(id) - (grading_published_by_id)"
    devilry_group_feedbackset |o--o| devilry_group_feedbacksetpassedpreviousperiod : "(id) - (feedbackset_id)"
    devilry_account_user |o--o{ devilry_group_feedbacksetpassedpreviousperiod : "(id) - (grading_published_by_id)"
    devilry_account_user |o--o{ devilry_group_feedbacksetpassedpreviousperiod : "(id) - (created_by_id)"
    devilry_group_feedbackset ||--o{ devilry_group_feedbacksetgradingupdatehistory : "(id) - (feedback_set_id)"
    devilry_account_user |o--o{ devilry_group_feedbacksetgradingupdatehistory : "(id) - (updated_by_id)"
    devilry_account_user |o--o{ devilry_group_feedbacksetgradingupdatehistory : "(id) - (old_grading_published_by_id)"
    devilry_group_feedbackset ||--o{ devilry_group_feedbacksetdeadlinehistory : "(id) - (feedback_set_id)"
    devilry_account_user |o--o{ devilry_group_feedbacksetdeadlinehistory : "(id) - (changed_by_id)"
    devilry_comment_comment ||--o| devilry_group_groupcomment : "is a, by comment_ptr_id"
    devilry_group_feedbackset ||--o{ devilry_group_groupcomment : "(id) - (feedback_set_id)"
    devilry_comment_commentedithistory ||--o| devilry_group_groupcommentedithistory : "is a, by commentedithistory_ptr_id"
    devilry_group_groupcomment ||--o{ devilry_group_groupcommentedithistory : "(comment_ptr_id) - (group_comment_id)"
    devilry_comment_comment ||--o| devilry_group_imageannotationcomment : "is a, by comment_ptr_id"
    devilry_group_feedbackset ||--o{ devilry_group_imageannotationcomment : "(id) - (feedback_set_id)"
    devilry_comment_commentfileimage ||--o{ devilry_group_imageannotationcomment : "(id) - (image_id)"
    devilry_account_user |o--o{ devilry_comment_comment : "(id) - (user_id)"
    devilry_comment_comment |o--o{ devilry_comment_comment : "(id) - (parent_id)"
    devilry_comment_comment ||--o{ devilry_comment_commentedithistory : "(id) - (comment_id)"
    devilry_account_user |o--o{ devilry_comment_commentedithistory : "(id) - (edited_by_id)"
    devilry_comment_comment ||--o{ devilry_comment_commentfile : "(id) - (comment_id)"
    devilry_comment_commentfile ||--o{ devilry_comment_commentfileimage : "(id) - (comment_file_id)"
    core_period ||--o| devilry_qualifiesforexam_periodtag : "(id) - (period_id)"
    devilry_qualifiesforexam_deadlinetag ||--o{ devilry_qualifiesforexam_periodtag : "(id) - (deadlinetag_id)"
    core_period ||--o{ devilry_qualifiesforexam_status : "(id) - (period_id)"
    devilry_account_user ||--o{ devilry_qualifiesforexam_status : "(id) - (user_id)"
    core_relatedstudent ||--o{ devilry_qualifiesforexam_qualifiesforfinalexam : "(id) - (relatedstudent_id)"
    devilry_qualifiesforexam_status ||--o{ devilry_qualifiesforexam_qualifiesforfinalexam : "(id) - (status_id)"
    core_relatedstudent ||--o{ devilry_qualifiesforexam_deletedqualifiesforfinalexam : "(id) - (relatedstudent_id)"
    devilry_qualifiesforexam_status ||--o{ devilry_qualifiesforexam_deletedqualifiesforfinalexam : "(id) - (status_id)"
    devilry_account_user ||--o{ devilry_qualifiesforexam_draftstatus : "(id) - (created_by_id)"
    core_period ||--o{ devilry_qualifiesforexam_draftstatus : "(id) - (period_id)"
    core_relatedstudent ||--o{ devilry_qualifiesforexam_draftqualifiesforfinalexam : "(id) - (relatedstudent_id)"
    devilry_qualifiesforexam_draftstatus ||--o{ devilry_qualifiesforexam_draftqualifiesforfinalexam : "(id) - (draft_status_id)"
    core_assignmentgroup ||--o| devilry_dbcache_assignmentgroupcacheddata : "(id) - (group_id)"
    devilry_group_feedbackset |o--o{ devilry_dbcache_assignmentgroupcacheddata : "(id) - (first_feedbackset_id)"
    devilry_group_feedbackset |o--o{ devilry_dbcache_assignmentgroupcacheddata : "(id) - (last_feedbackset_id)"
    devilry_group_feedbackset |o--o{ devilry_dbcache_assignmentgroupcacheddata : "(id) - (last_published_feedbackset_id)"
    devilry_account_user |o--o{ devilry_message_message : "(id) - (created_by_id)"
    devilry_message_message ||--o{ devilry_message_messagereceiver : "(id) - (message_id)"
    devilry_account_user ||--o{ devilry_message_messagereceiver : "(id) - (user_id)"
    devilry_account_user ||--o{ devilry_report_devilryreport : "(id) - (generated_by_user_id)"
    devilry_account_user ||--o{ account_emailaddress : "(id) - (user_id)"
    account_emailaddress ||--o{ account_emailconfirmation : "(id) - (email_address_id)"
    socialaccount_socialapp }o--o{ django_site : "(sites) - (socialapp_set)"
    devilry_account_user ||--o{ socialaccount_socialaccount : "(id) - (user_id)"
    socialaccount_socialapp |o--o{ socialaccount_socialtoken : "(id) - (app_id)"
    socialaccount_socialaccount ||--o{ socialaccount_socialtoken : "(id) - (account_id)"
    django_session {
        CharField session_key PK
        TextField session_data
        DateTimeField expire_date
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
    django_site {
        AutoField id PK
        CharField domain UK
        CharField name
    }
    ievv_batchframework_batchoperation {
        AutoField id PK
        ForeignKey started_by_id FK
        DateTimeField created_datetime
        DateTimeField started_running_datetime
        DateTimeField finished_datetime
        ForeignKey context_content_type_id FK
        PositiveIntegerField context_object_id
        CharField operationtype
        BatchOperationStatus status
        BatchOperationResult result
        TextField input_data_json
        TextField output_data_json
    }
    ievv_developemail_developemail {
        AutoField id PK
        DateTimeField created_datetime
        TextField subject
        TextField from_email
        TextField to_emails
        TextField raw_message
    }
    django_rq_dashboard {
        AutoField id PK
    }
    cradmin_temporaryfileuploadstore_temporaryfilecollection {
        AutoField id PK
        ForeignKey user_id FK "The user that owns this temporary file. Users should notbe allowed access to other users temporary files."
        DateTimeField created_datetime
        PositiveIntegerField minutes_to_live "The number of minutes the app requests that this file collection should be kept before automatic removal. This is used by automatic cleanup jobs to determine what to delete. You should not rely on the file beeing automatically deleted after this number of minutes, and you should always delete files explicitly as part of the file upload process."
        TextField accept "An html input field accept attribute formatted string. This is validated by the API on upload."
        IntegerField max_filename_length "If specified, we shorten filenames to maximum the specified length. This is validated by the API on upload."
        BooleanField unique_filenames "If this is True, we add random data when we detect duplicate filenames. The duplicate prevention algorithm handles max_filename.This is validated by the API on upload."
        PositiveIntegerField max_filesize_bytes
        BooleanField singlemode "If this is True, only a single file can be added to the collection. This means that the last file added to the collection will be the only file in the collection."
    }
    cradmin_temporaryfileuploadstore_temporaryfile {
        AutoField id PK
        ForeignKey collection_id FK
        TextField filename
        FileField file
        TextField mimetype
    }
    cradmin_generic_token_with_metadata_generictokenwithmetadata {
        AutoField id PK
        CharField app
        CharField token UK
        DateTimeField created_datetime
        DateTimeField expiration_datetime
        BooleanField single_use
        TextField metadata_json
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
    }
    core_subject {
        AutoField id PK
        ShortNameField short_name UK "Up to 20 letters of lowercase english letters (a-z), numbers, underscore ('_') and hyphen ('-'). Used when the name takes too much space."
        LongNameField long_name
        DateTimeField etag
    }
    core_period {
        AutoField id PK
        ShortNameField short_name "Up to 20 letters of lowercase english letters (a-z), numbers, underscore ('_') and hyphen ('-'). Used when the name takes too much space."
        LongNameField long_name
        ForeignKey parentnode_id FK
        DateTimeField start_time "Start time and end time defines when the period is active."
        DateTimeField end_time "Start time and end time defines when the period is active."
        DateTimeField etag
    }
    core_periodapplicationkeyvalue {
        AutoField id PK
        CharField application "Application identifier. Max 300 chars."
        CharField key "Key. Max 300 chars."
        TextField value "Value."
        ForeignKey period_id FK "The period where this metadata belongs."
    }
    core_relatedexaminer {
        AutoField id PK
        ForeignKey period_id FK "The period."
        ForeignKey user_id FK "The related user."
        TextField tags "Comma-separated list of tags. Each tag is a word with the following letters allowed: a-z, 0-9, ``_`` and ``-``. Each word is separated by a comma, and no whitespace."
        CharField automatic_anonymous_id
        BooleanField active
    }
    core_relatedstudent {
        AutoField id PK
        ForeignKey period_id FK "The period."
        ForeignKey user_id FK "The related user."
        TextField tags "Comma-separated list of tags. Each tag is a word with the following letters allowed: a-z, 0-9, ``_`` and ``-``. Each word is separated by a comma, and no whitespace."
        CharField automatic_anonymous_id
        BooleanField active
        CharField candidate_id
    }
    core_relatedstudentkeyvalue {
        AutoField id PK
        CharField application "Application identifier. Max 300 chars."
        CharField key "Key. Max 300 chars."
        TextField value "Value."
        ForeignKey relatedstudent_id FK
        BooleanField student_can_read "Specifies if a student can read the value or not."
    }
    core_periodtag {
        AutoField id PK
        ForeignKey period_id FK
        CharField prefix
        TextField tag
        BooleanField is_hidden
        DateTimeField created_datetime
        DateTimeField modified_datetime
    }
    core_assignment {
        AutoField id PK
        ShortNameField short_name "Up to 20 letters of lowercase english letters (a-z), numbers, underscore ('_') and hyphen ('-'). Used when the name takes too much space."
        LongNameField long_name
        TextField gradeform_setup_json
        ForeignKey parentnode_id FK
        DateTimeField publishing_time "The time when the assignment is to be published (visible to students and examiners)."
        BooleanField deprecated_field_anonymous "Deprecated anonymous field. Will be removed in 3.1."
        BooleanField uses_custom_candidate_ids "If this is enabled, the assignment does not inherit candidate IDs from the semester, and instead have their own set of candidate IDs only for this assignment."
        AssignmentAnonymizationmode anonymizationmode
        BooleanField students_can_see_points
        AssignmentDeliveryTypes delivery_types "This option controls what types of deliveries this assignment accepts. See the Delivery documentation for more info."
        AssignmentDeadlineHandling deadline_handling "With HARD deadlines, students will be unable to make deliveries when a deadline has expired. With SOFT deadlines students will be able to make deliveries after the deadline has expired. All deliveries after their deadline are clearly highligted. NOTE: Devilry is designed from the bottom up to gracefully handle SOFT deadlines. Students have to perform an extra confirm-step when adding deliveries after their active deadline, and assignments where the deadline has expired is clearly marked for both students and examiners."
        PositiveIntegerField scale_points_percent "Percent to scale points on this assignment by for period overviews. The default is 100, which means no change to the points."
        DateTimeField first_deadline
        PositiveIntegerField max_points
        PositiveIntegerField passing_grade_min_points
        AssignmentPointsToGradeMapper points_to_grade_mapper
        AssignmentGradingSystemPluginId grading_system_plugin_id
        BooleanField students_can_create_groups "Select this if students should be allowed to join/leave groups. Even if this is not selected, you can still organize your students in groups manually."
        DateTimeField students_can_not_create_groups_after "Students can not create project groups after this time. Ignored if 'Students can create project groups' is not selected."
        AssignmentFeedbackWorkflow feedback_workflow
        BooleanField examiners_can_self_assign "Select this if you want examiners to be able to self-assign to the assignment."
        PositiveIntegerField examiner_self_assign_limit "The max number of examiners that are allowed to self-assign to each project group within the assignment."
    }
    core_pointtogrademap {
        AutoField id PK
        OneToOneField assignment_id FK, UK
        BooleanField invalid
    }
    core_pointrangetograde {
        AutoField id PK
        ForeignKey point_to_grade_map_id FK
        PositiveIntegerField minimum_points
        PositiveIntegerField maximum_points
        CharField grade
    }
    core_assignmentgroup {
        AutoField id PK
        ForeignKey parentnode_id FK
        CharField name "An optional name for the group. Typically used a project name on project assignments."
        BooleanField is_open "If this is checked, the group can add deliveries."
        DateTimeField etag
        ForeignKey batchoperation_id FK
        ForeignKey copied_from_id FK
        DateTimeField created_datetime
        BooleanField internal_is_being_deleted
    }
    core_assignmentgrouptag {
        AutoField id PK
        ForeignKey assignment_group_id FK
        SlugField tag "A tag can contain a-z, A-Z, 0-9 and '_'."
    }
    core_assignmentgrouphistory {
        AutoField id PK
        OneToOneField assignment_group_id FK, UK
        TextField merge_history_json
    }
    core_candidate {
        AutoField id PK
        ForeignKey old_reference_not_in_use_student_id FK
        ForeignKey relatedstudent_id FK
        ForeignKey assignment_group_id FK
        CharField candidate_id "An optional candidate id. This can be anything as long as it is less than 30 characters. Used to show the user on anonymous assignmens."
    }
    core_devilryuserprofile {
        AutoField id PK
        OneToOneField user_id FK, UK
        CharField full_name
        CharField languagecode
    }
    core_assignmentgroup_examiners {
        AutoField id PK
        ForeignKey old_reference_not_in_use_user_id FK
        ForeignKey assignmentgroup_id FK
        ForeignKey relatedexaminer_id FK
    }
    core_groupinvite {
        AutoField id PK
        ForeignKey group_id FK
        DateTimeField sent_datetime
        ForeignKey sent_by_id FK
        ForeignKey sent_to_id FK
        BooleanField accepted
        DateTimeField responded_datetime
    }
    core_examinerassignmentgrouphistory {
        AutoField id PK
        ForeignKey assignment_group_id FK
        ForeignKey user_id FK
        DateTimeField created_datetime
        BooleanField is_add
    }
    core_candidateassignmentgrouphistory {
        AutoField id PK
        ForeignKey assignment_group_id FK
        ForeignKey user_id FK
        DateTimeField created_datetime
        BooleanField is_add
    }
    devilry_account_user {
        AutoField id PK
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        CharField shortname UK "The short name for the user. This is set automatically to the email or username depending on the method used for authentication."
        TextField fullname
        TextField lastname
        DateTimeField datetime_joined
        DateTimeField suspended_datetime "Time when the account was suspended"
        TextField suspended_reason
        CharField languagecode
    }
    devilry_account_mergeduser {
        AutoField id PK
        ForeignKey source_user_id FK
        ForeignKey target_user_id FK
        JSONField summary_json
    }
    devilry_account_useremail {
        AutoField id PK
        ForeignKey user_id FK
        DateTimeField created_datetime
        DateTimeField last_updated_datetime
        EmailField email UK
        BooleanField use_for_notifications
        UserEmailIsPrimary is_primary "Your primary email is the email address used when we need to display a single email address."
    }
    devilry_account_username {
        AutoField id PK
        ForeignKey user_id FK
        DateTimeField created_datetime
        DateTimeField last_updated_datetime
        CharField username UK
        UserNameIsPrimary is_primary "Your primary username is shown alongside your full name to identify you to teachers, examiners and other students."
    }
    devilry_account_permissiongroupuser {
        AutoField id PK
        ForeignKey permissiongroup_id FK
        ForeignKey user_id FK
    }
    devilry_account_permissiongroup {
        AutoField id PK
        CharField name UK "A unique name for this group."
        DateTimeField created_datetime "The time when this group was created."
        DateTimeField updated_datetime "The time when this group last updated."
        DateTimeField syncsystem_update_datetime "The time when this group last updated from a third party sync system."
        PermissionGroupGrouptype grouptype "Course and semester administrator groups can only be assigned to a single course or semester. Department administrator groups can be assigned to multiple courses. You can not change this for existing permission groups."
        BooleanField is_custom_manageable "Is this group mageable by non-superusers. Can not be enabled for department administrator groups."
    }
    devilry_account_periodpermissiongroup {
        AutoField id PK
        ForeignKey permissiongroup_id FK
        ForeignKey period_id FK
    }
    devilry_account_subjectpermissiongroup {
        AutoField id PK
        ForeignKey permissiongroup_id FK
        ForeignKey subject_id FK
    }
    devilry_account_perioduserguidelineacceptance {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey period_id FK
        CharField devilryrole
        DateTimeField accepted_datetime
        CharField guidelines_version
        TextField matched_regex
    }
    devilry_compressionutil_compressedarchivemeta {
        AutoField id PK
        ForeignKey content_type_id FK
        PositiveIntegerField content_object_id
        ForeignKey created_by_id FK
        CompressedArchiveMetaCreatedByRole created_by_role
        DateTimeField created_datetime
        CharField archive_name
        CharField archive_path
        BigIntegerField archive_size
        CharField backend_id
        DateTimeField deleted_datetime
    }
    devilry_group_feedbackset {
        AutoField id PK
        ForeignKey group_id FK
        FeedbackSetFeedbacksetType feedbackset_type
        BooleanField ignored
        TextField ignored_reason
        DateTimeField ignored_datetime
        ForeignKey created_by_id FK
        DateTimeField created_datetime
        ForeignKey last_updated_by_id FK
        DateTimeField deadline_datetime
        DateTimeField grading_published_datetime
        ForeignKey grading_published_by_id FK
        PositiveIntegerField grading_points
        TextField gradeform_data_json
    }
    devilry_group_feedbacksetpassedpreviousperiod {
        AutoField id PK
        OneToOneField feedbackset_id FK, UK
        FeedbacksetPassedPreviousPeriodPassedPreviousPeriodType passed_previous_period_type
        ShortNameField assignment_short_name "Up to 20 letters of lowercase english letters (a-z), numbers, underscore ('_') and hyphen ('-'). Used when the name takes too much space."
        LongNameField assignment_long_name
        PositiveIntegerField assignment_max_points
        PositiveIntegerField assignment_passing_grade_min_points
        ShortNameField period_short_name "Up to 20 letters of lowercase english letters (a-z), numbers, underscore ('_') and hyphen ('-'). Used when the name takes too much space."
        LongNameField period_long_name
        DateTimeField period_start_time
        DateTimeField period_end_time
        PositiveIntegerField grading_points
        ForeignKey grading_published_by_id FK
        DateTimeField grading_published_datetime
        DateTimeField created_datetime
        ForeignKey created_by_id FK
    }
    devilry_group_feedbacksetgradingupdatehistory {
        AutoField id PK
        ForeignKey feedback_set_id FK
        ForeignKey updated_by_id FK
        DateTimeField updated_datetime
        PositiveIntegerField old_grading_points
        ForeignKey old_grading_published_by_id FK
        DateTimeField old_grading_published_datetime
    }
    devilry_group_feedbacksetdeadlinehistory {
        AutoField id PK
        ForeignKey feedback_set_id FK
        ForeignKey changed_by_id FK
        DateTimeField changed_datetime
        DateTimeField deadline_old
        DateTimeField deadline_new
    }
    devilry_group_groupcomment {
        OneToOneField comment_ptr_id PK
        ForeignKey feedback_set_id FK
        BooleanField part_of_grading
        GroupCommentVisibility visibility
        CharField v2_id
    }
    devilry_group_groupcommentedithistory {
        OneToOneField commentedithistory_ptr_id PK
        ForeignKey group_comment_id FK
        CharField visibility
    }
    devilry_group_imageannotationcomment {
        OneToOneField comment_ptr_id PK
        ForeignKey feedback_set_id FK
        BooleanField part_of_grading
        ImageAnnotationCommentVisibility visibility
        ForeignKey image_id FK
        PositiveIntegerField x_coordinate
        PositiveIntegerField y_coordinate
    }
    devilry_comment_comment {
        AutoField id PK
        TextField text
        TextField draft_text
        ForeignKey user_id FK
        ForeignKey parent_id FK
        DateTimeField created_datetime
        DateTimeField published_datetime
        CommentUserRole user_role
        CommentCommentType comment_type
    }
    devilry_comment_commentedithistory {
        AutoField id PK
        ForeignKey comment_id FK
        ForeignKey edited_by_id FK
        DateTimeField edited_datetime
        TextField post_edit_text
        TextField pre_edit_text
    }
    devilry_comment_commentfile {
        AutoField id PK
        CharField mimetype
        FileField file
        CharField filename
        PositiveIntegerField filesize
        ForeignKey comment_id FK
        DateTimeField processing_started_datetime
        DateTimeField processing_completed_datetime
        BooleanField processing_successful
        DateTimeField created_datetime
        CharField v2_id
    }
    devilry_comment_commentfileimage {
        AutoField id PK
        ForeignKey comment_file_id FK
        FileField image
        PositiveIntegerField image_width
        PositiveIntegerField image_height
        FileField thumbnail
        PositiveIntegerField thumbnail_width
        PositiveIntegerField thumbnail_height
    }
    devilry_qualifiesforexam_deadlinetag {
        AutoField id PK
        DateTimeField timestamp
        CharField tag
    }
    devilry_qualifiesforexam_periodtag {
        OneToOneField period_id PK
        ForeignKey deadlinetag_id FK
    }
    devilry_qualifiesforexam_status {
        AutoField id PK
        StatusStatus status
        ForeignKey period_id FK
        DateTimeField createtime
        TextField message
        ForeignKey user_id FK
        CharField plugin
        JSONField plugin_data
        DateTimeField exported_timestamp
    }
    devilry_qualifiesforexam_qualifiesforfinalexam {
        AutoField id PK
        ForeignKey relatedstudent_id FK
        ForeignKey status_id FK
        BooleanField qualifies
    }
    devilry_qualifiesforexam_deletedqualifiesforfinalexam {
        AutoField id PK
        ForeignKey relatedstudent_id FK
        ForeignKey status_id FK
        BooleanField qualifies
    }
    devilry_qualifiesforexam_draftstatus {
        AutoField id PK
        DateTimeField created_datetime
        ForeignKey created_by_id FK
        ForeignKey period_id FK
        CharField plugin
        JSONField plugin_data
        DateTimeField processing_started_datetime
        DateTimeField processing_completed_datetime
        DraftStatusProcessingStatusChoices processing_status
        JSONField processing_status_data
    }
    devilry_qualifiesforexam_draftqualifiesforfinalexam {
        AutoField id PK
        ForeignKey relatedstudent_id FK
        ForeignKey draft_status_id FK
        BooleanField qualifies
    }
    devilry_dbcache_assignmentgroupcacheddata {
        AutoField id PK
        OneToOneField group_id FK, UK
        ForeignKey first_feedbackset_id FK
        ForeignKey last_feedbackset_id FK
        ForeignKey last_published_feedbackset_id FK
        PositiveIntegerField new_attempt_count
        PositiveIntegerField public_total_comment_count
        PositiveIntegerField public_student_comment_count
        PositiveIntegerField public_examiner_comment_count
        PositiveIntegerField public_admin_comment_count
        PositiveIntegerField public_student_file_upload_count
        PositiveIntegerField public_student_attempts_with_delivered_files
        DateTimeField last_public_comment_by_student_datetime
        DateTimeField last_public_comment_by_examiner_datetime
        PositiveIntegerField examiner_count
        PositiveIntegerField candidate_count
    }
    devilry_message_message {
        AutoField id PK
        DateTimeField created_datetime
        ForeignKey created_by_id FK
        MessageStatus status
        JSONField status_data
        JSONField metadata
        MessageContextType context_type
        ArrayField message_type
        JSONField virtual_message_receivers
    }
    devilry_message_messagereceiver {
        AutoField id PK
        DateTimeField created_datetime
        MessageReceiverStatus status
        JSONField status_data
        JSONField metadata
        CharField subject
        TextField message_content_plain
        TextField message_content_html
        ForeignKey message_id FK
        CharField message_type
        CharField send_to
        ForeignKey user_id FK
        DateTimeField sent_datetime
        IntegerField sending_failed_count
        IntegerField sending_success_count
    }
    devilry_report_devilryreport {
        AutoField id PK
        ForeignKey generated_by_user_id FK
        DateTimeField created_datetime
        DateTimeField started_datetime
        DateTimeField finished_datetime
        CharField generator_type
        JSONField generator_options
        DevilryReportStatus status
        JSONField status_data
        CharField output_filename
        CharField content_type
        BinaryField result
    }
    account_emailaddress {
        AutoField id PK
        ForeignKey user_id FK
        EmailField email
        BooleanField verified
        BooleanField primary
    }
    account_emailconfirmation {
        AutoField id PK
        ForeignKey email_address_id FK
        DateTimeField created
        DateTimeField sent
        CharField key UK
    }
    socialaccount_socialapp {
        AutoField id PK
        CharField provider
        CharField provider_id
        CharField name
        CharField client_id "App ID, or consumer key"
        CharField secret "API secret, client secret, or consumer secret"
        CharField key "Key"
        JSONField settings
    }
    socialaccount_socialaccount {
        AutoField id PK
        ForeignKey user_id FK
        CharField provider
        CharField uid
        DateTimeField last_login
        DateTimeField date_joined
        JSONField extra_data
    }
    socialaccount_socialtoken {
        AutoField id PK
        ForeignKey app_id FK
        ForeignKey account_id FK
        TextField token "'oauth_token' (OAuth1) or access token (OAuth2)"
        TextField token_secret "'oauth_token_secret' (OAuth1) or refresh token (OAuth2)"
        DateTimeField expires_at
    }
```
