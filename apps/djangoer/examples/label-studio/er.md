```mermaid
erDiagram
    htx_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    htx_user ||--o| authtoken_token : "(id) - (user_id)"
    htx_user |o--o{ token_blacklist_outstandingtoken : "(id) - (user_id)"
    token_blacklist_outstandingtoken ||--o| token_blacklist_blacklistedtoken : "(id) - (token_id)"
    htx_user |o--o{ fsm_taskstate : "(id) - (triggered_by_id)"
    task ||--o{ fsm_taskstate : "(id) - (task_id)"
    htx_user |o--o{ fsm_annotationstate : "(id) - (triggered_by_id)"
    task_completion ||--o{ fsm_annotationstate : "(id) - (annotation_id)"
    htx_user |o--o{ fsm_projectstate : "(id) - (triggered_by_id)"
    project ||--o{ fsm_projectstate : "(id) - (project_id)"
    project |o--o{ core_asyncmigrationstatus : "(id) - (project_id)"
    organization |o--o{ htx_user : "(id) - (active_organization_id)"
    htx_user }o--o{ auth_group : "(groups) - (user_set)"
    htx_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    htx_user ||--o{ organizations_organizationmember : "(id) - (user_id)"
    organization ||--o{ organizations_organizationmember : "(id) - (organization_id)"
    htx_user |o--o| organization : "(id) - (created_by_id)"
    htx_user ||--o{ data_import_fileupload : "(id) - (user_id)"
    project ||--o{ data_import_fileupload : "(id) - (project_id)"
    project ||--o{ data_export_export : "(id) - (project_id)"
    htx_user |o--o{ data_export_export : "(id) - (created_by_id)"
    project |o--o{ data_export_convertedformat : "(id) - (project_id)"
    organization |o--o{ data_export_convertedformat : "(id) - (organization_id)"
    data_export_export ||--o{ data_export_convertedformat : "(id) - (export_id)"
    htx_user |o--o{ data_export_convertedformat : "(id) - (created_by_id)"
    organization |o--o{ project : "(id) - (organization_id)"
    htx_user |o--o{ project : "(id) - (created_by_id)"
    htx_user |o--o{ project : "(id) - (deleted_by_id)"
    projects_projectonboardingsteps ||--o{ projects_projectonboarding : "(id) - (step_id)"
    project ||--o{ projects_projectonboarding : "(id) - (project_id)"
    htx_user ||--o{ projects_labelstreamhistory : "(id) - (user_id)"
    project ||--o{ projects_labelstreamhistory : "(id) - (project_id)"
    htx_user ||--o{ projects_projectmember : "(id) - (user_id)"
    project ||--o{ projects_projectmember : "(id) - (project_id)"
    project ||--o| projects_projectsummary : "(id) - (project_id)"
    project |o--o{ projects_projectimport : "(id) - (project_id)"
    project |o--o{ projects_projectreimport : "(id) - (project_id)"
    project |o--o{ task : "(id) - (project_id)"
    htx_user |o--o{ task : "(id) - (updated_by_id)"
    data_import_fileupload |o--o{ task : "(id) - (file_upload_id)"
    task }o--o{ htx_user : "(comment_authors) - (tasks_with_comments)"
    task |o--o{ task_completion : "(id) - (task_id)"
    project |o--o{ task_completion : "(id) - (project_id)"
    htx_user |o--o{ task_completion : "(id) - (completed_by_id)"
    htx_user |o--o{ task_completion : "(id) - (updated_by_id)"
    prediction |o--o{ task_completion : "(id) - (parent_prediction_id)"
    task_completion |o--o{ task_completion : "(id) - (parent_annotation_id)"
    htx_user |o--o{ task_completion : "(id) - (last_created_by_id)"
    task ||--o{ tasks_tasklock : "(id) - (task_id)"
    htx_user ||--o{ tasks_tasklock : "(id) - (user_id)"
    task |o--o{ tasks_annotationdraft : "(id) - (task_id)"
    task_completion |o--o{ tasks_annotationdraft : "(id) - (annotation_id)"
    htx_user ||--o{ tasks_annotationdraft : "(id) - (user_id)"
    ml_mlbackend |o--o{ prediction : "(id) - (model_id)"
    ml_models_modelrun |o--o{ prediction : "(id) - (model_run_id)"
    task ||--o{ prediction : "(id) - (task_id)"
    project |o--o{ prediction : "(id) - (project_id)"
    ml_mlbackend |o--o{ tasks_failedprediction : "(id) - (ml_backend_model_id)"
    ml_models_modelrun |o--o{ tasks_failedprediction : "(id) - (model_run_id)"
    project |o--o{ tasks_failedprediction : "(id) - (project_id)"
    task ||--o{ tasks_failedprediction : "(id) - (task_id)"
    prediction |o--o| prediction_meta : "(id) - (prediction_id)"
    tasks_failedprediction |o--o| prediction_meta : "(id) - (failed_prediction_id)"
    data_manager_filtergroup |o--o{ data_manager_view : "(id) - (filter_group_id)"
    htx_user |o--o{ data_manager_view : "(id) - (user_id)"
    project ||--o{ data_manager_view : "(id) - (project_id)"
    data_manager_filtergroup }o--o{ data_manager_filter : "(filters) - (filter_groups)"
    data_manager_filter |o--o{ data_manager_filter : "(id) - (parent_id)"
    io_storages_azureblobstoragemixin ||--o| io_storages_azureblobimportstorage : "is a, by azureblobstoragemixin_ptr_id"
    project ||--o{ io_storages_azureblobimportstorage : "(id) - (project_id)"
    io_storages_azureblobstoragemixin ||--o| io_storages_azureblobexportstorage : "is a, by azureblobstoragemixin_ptr_id"
    project ||--o{ io_storages_azureblobexportstorage : "(id) - (project_id)"
    task ||--o| io_storages_azureblobimportstoragelink : "(id) - (task_id)"
    io_storages_azureblobimportstorage ||--o{ io_storages_azureblobimportstoragelink : "(azureblobstoragemixin_ptr_id) - (storage_id)"
    task_completion ||--o{ io_storages_azureblobexportstoragelink : "(id) - (annotation_id)"
    io_storages_azureblobexportstorage ||--o{ io_storages_azureblobexportstoragelink : "(azureblobstoragemixin_ptr_id) - (storage_id)"
    project ||--o{ io_storages_s3importstorage : "(id) - (project_id)"
    project ||--o{ io_storages_s3exportstorage : "(id) - (project_id)"
    task ||--o| io_storages_s3importstoragelink : "(id) - (task_id)"
    io_storages_s3importstorage ||--o{ io_storages_s3importstoragelink : "(id) - (storage_id)"
    task_completion ||--o{ io_storages_s3exportstoragelink : "(id) - (annotation_id)"
    io_storages_s3exportstorage ||--o{ io_storages_s3exportstoragelink : "(id) - (storage_id)"
    io_storages_gcsstoragemixin ||--o| io_storages_gcsimportstorage : "is a, by gcsstoragemixin_ptr_id"
    project ||--o{ io_storages_gcsimportstorage : "(id) - (project_id)"
    io_storages_gcsstoragemixin ||--o| io_storages_gcsexportstorage : "is a, by gcsstoragemixin_ptr_id"
    project ||--o{ io_storages_gcsexportstorage : "(id) - (project_id)"
    task ||--o| io_storages_gcsimportstoragelink : "(id) - (task_id)"
    io_storages_gcsimportstorage ||--o{ io_storages_gcsimportstoragelink : "(gcsstoragemixin_ptr_id) - (storage_id)"
    task_completion ||--o{ io_storages_gcsexportstoragelink : "(id) - (annotation_id)"
    io_storages_gcsexportstorage ||--o{ io_storages_gcsexportstoragelink : "(gcsstoragemixin_ptr_id) - (storage_id)"
    io_storages_redisstoragemixin ||--o| io_storages_redisimportstorage : "is a, by redisstoragemixin_ptr_id"
    project ||--o{ io_storages_redisimportstorage : "(id) - (project_id)"
    io_storages_redisstoragemixin ||--o| io_storages_redisexportstorage : "is a, by redisstoragemixin_ptr_id"
    project ||--o{ io_storages_redisexportstorage : "(id) - (project_id)"
    task ||--o| io_storages_redisimportstoragelink : "(id) - (task_id)"
    io_storages_redisimportstorage ||--o{ io_storages_redisimportstoragelink : "(redisstoragemixin_ptr_id) - (storage_id)"
    task_completion ||--o{ io_storages_redisexportstoragelink : "(id) - (annotation_id)"
    io_storages_redisexportstorage ||--o{ io_storages_redisexportstoragelink : "(redisstoragemixin_ptr_id) - (storage_id)"
    project ||--o{ ml_mlbackend : "(id) - (project_id)"
    ml_mlbackend ||--o{ ml_mlbackendpredictionjob : "(id) - (ml_backend_id)"
    ml_mlbackend ||--o{ ml_mlbackendtrainjob : "(id) - (ml_backend_id)"
    organization ||--o{ webhook : "(id) - (organization_id)"
    project |o--o{ webhook : "(id) - (project_id)"
    webhook ||--o{ webhook_action : "(id) - (webhook_id)"
    htx_user ||--o{ labels_manager_label : "(id) - (created_by_id)"
    htx_user |o--o{ labels_manager_label : "(id) - (approved_by_id)"
    organization ||--o{ labels_manager_label : "(id) - (organization_id)"
    project ||--o{ labels_manager_labellink : "(id) - (project_id)"
    labels_manager_label ||--o{ labels_manager_labellink : "(id) - (label_id)"
    htx_user |o--o{ ml_models_modelinterface : "(id) - (created_by_id)"
    organization |o--o{ ml_models_modelinterface : "(id) - (organization_id)"
    ml_models_modelinterface }o--o{ project : "(associated_projects) - (modelinterface_set)"
    ml_models_modelinterface ||--o{ ml_models_thirdpartymodelversion : "(id) - (parent_model_id)"
    ml_model_providers_modelproviderconnection |o--o{ ml_models_thirdpartymodelversion : "(id) - (model_provider_connection_id)"
    htx_user |o--o{ ml_models_thirdpartymodelversion : "(id) - (created_by_id)"
    organization |o--o{ ml_models_thirdpartymodelversion : "(id) - (organization_id)"
    organization |o--o{ ml_models_modelrun : "(id) - (organization_id)"
    project ||--o{ ml_models_modelrun : "(id) - (project_id)"
    ml_models_thirdpartymodelversion ||--o{ ml_models_modelrun : "(id) - (model_version_id)"
    htx_user |o--o{ ml_models_modelrun : "(id) - (created_by_id)"
    organization |o--o{ ml_model_providers_modelproviderconnection : "(id) - (organization_id)"
    htx_user |o--o{ ml_model_providers_modelproviderconnection : "(id) - (created_by_id)"
    organization ||--o| jwt_auth_jwtsettings : "(id) - (organization_id)"
    organization ||--o| session_policy_sessiontimeoutpolicy : "(id) - (organization_id)"
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
    django_rq_queue {
        AutoField id PK
    }
    authtoken_token {
        CharField key PK
        OneToOneField user_id FK, UK
        DateTimeField created
    }
    token_blacklist_outstandingtoken {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField jti UK
        TextField token
        DateTimeField created_at
        DateTimeField expires_at
    }
    token_blacklist_blacklistedtoken {
        BigAutoField id PK
        OneToOneField token_id FK, UK
        DateTimeField blacklisted_at
    }
    fsm_taskstate {
        UUIDField id PK "UUID7 provides natural time ordering and global uniqueness"
        PositiveIntegerField organization_id "Organization ID that owns this state record (for multi-tenant applications)"
        CharField previous_state "Previous state before this transition"
        CharField transition_name "Name of the transition method that triggered this state change"
        ForeignKey triggered_by_id FK "User who triggered this state transition"
        JSONField context_data "Additional context data for this transition (e.g., validation results, external IDs)"
        TextField reason "Human-readable reason for this state transition"
        DateTimeField created_at "Human-readable timestamp for debugging (UUID7 id contains precise timestamp)"
        ForeignKey task_id FK
        ProjectStateChoices state
        PositiveIntegerField project_id "From task.project_id - denormalized for performance"
    }
    fsm_annotationstate {
        UUIDField id PK "UUID7 provides natural time ordering and global uniqueness"
        PositiveIntegerField organization_id "Organization ID that owns this state record (for multi-tenant applications)"
        CharField previous_state "Previous state before this transition"
        CharField transition_name "Name of the transition method that triggered this state change"
        ForeignKey triggered_by_id FK "User who triggered this state transition"
        JSONField context_data "Additional context data for this transition (e.g., validation results, external IDs)"
        TextField reason "Human-readable reason for this state transition"
        DateTimeField created_at "Human-readable timestamp for debugging (UUID7 id contains precise timestamp)"
        ForeignKey annotation_id FK
        AnnotationStateChoices state
        PositiveIntegerField task_id "From annotation.task_id - denormalized for performance"
        PositiveIntegerField project_id "From annotation.task.project_id - denormalized for performance"
        PositiveIntegerField completed_by_id "From annotation.completed_by_id - denormalized for performance"
    }
    fsm_projectstate {
        UUIDField id PK "UUID7 provides natural time ordering and global uniqueness"
        PositiveIntegerField organization_id "Organization ID that owns this state record (for multi-tenant applications)"
        CharField previous_state "Previous state before this transition"
        CharField transition_name "Name of the transition method that triggered this state change"
        ForeignKey triggered_by_id FK "User who triggered this state transition"
        JSONField context_data "Additional context data for this transition (e.g., validation results, external IDs)"
        TextField reason "Human-readable reason for this state transition"
        DateTimeField created_at "Human-readable timestamp for debugging (UUID7 id contains precise timestamp)"
        ForeignKey project_id FK
        ProjectStateChoices state
        PositiveIntegerField created_by_id "From project.created_by_id - denormalized for performance"
    }
    core_asyncmigrationstatus {
        AutoField id PK
        JSONField meta "Meta is for any params for migrations, e.g.: project, filter or error message."
        ForeignKey project_id FK "Project ID for this migration"
        TextField name "Migration name"
        AsyncMigrationStatusStatus status
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Last updated time"
    }
    core_deletedrow {
        AutoField id PK
        CharField model
        IntegerField row_id
        JSONField data
        TextField reason
        DateTimeField created_at
        DateTimeField updated_at
        IntegerField organization_id
        IntegerField project_id
        IntegerField user_id
    }
    htx_user {
        AutoField id PK
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        DateTimeField last_activity
        CharField username
        EmailField email UK
        CharField first_name
        CharField last_name
        CharField phone
        ImageField avatar
        JSONField custom_hotkeys "Custom keyboard shortcuts configuration for the user interface"
        BooleanField is_staff "Designates whether the user can log into this admin site."
        BooleanField is_active "Designates whether to treat this user as active. Unselect this instead of deleting accounts."
        DateTimeField date_joined
        DateTimeField activity_at
        ForeignKey active_organization_id FK
        BooleanField allow_newsletters "Allow sending newsletters to user"
    }
    organizations_organizationmember {
        AutoField id PK
        ForeignKey user_id FK "User ID"
        ForeignKey organization_id FK "Organization ID"
        DateTimeField created_at
        DateTimeField updated_at
        DateTimeField deleted_at "Timestamp indicating when the organization member was marked as deleted.  If NULL, the member is not considered deleted."
    }
    organization {
        AutoField id PK
        CharField title
        CharField token UK
        OneToOneField created_by_id FK, UK
        DateTimeField created_at
        DateTimeField updated_at
        EmailField contact_info
    }
    data_import_fileupload {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey project_id FK
        FileField file
    }
    data_export_export {
        AutoField id PK
        CharField title
        DateTimeField created_at "Creation time"
        FileField file
        CharField md5
        DateTimeField finished_at "Complete or fail time"
        ExportStatus status
        JSONField counters
        ForeignKey project_id FK
        ForeignKey created_by_id FK
    }
    data_export_convertedformat {
        AutoField id PK
        ForeignKey project_id FK
        ForeignKey organization_id FK
        ForeignKey export_id FK "Export snapshot for this converted file"
        FileField file
        ConvertedFormatStatus status
        TextField traceback "Traceback report in case of errors"
        CharField export_type
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Updated time"
        DateTimeField finished_at "Complete or fail time"
        ForeignKey created_by_id FK
    }
    project {
        AutoField id PK
        CharField title "Project name. Must be between 3 and 50 characters long."
        TextField description "Project description"
        ForeignKey organization_id FK
        TextField label_config "Label config in XML format. See more about it in documentation"
        JSONField parsed_label_config "Parsed label config in JSON format. See more about it in documentation"
        BigIntegerField label_config_hash
        TextField expert_instruction "Labeling instructions in HTML format"
        BooleanField show_instruction "Show instructions to the annotator before they start"
        BooleanField show_skip_button "Show a skip button in interface and allow annotators to skip the task"
        BooleanField enable_empty_annotation "Allow annotators to submit empty annotations"
        BooleanField reveal_preannotations_interactively "Reveal pre-annotations interactively"
        BooleanField show_annotation_history "Show annotation history to annotator"
        BooleanField show_collab_predictions "If set, the annotator can view model predictions"
        BooleanField evaluate_predictions_automatically "Retrieve and display predictions when loading a task"
        CharField token
        IntegerField result_count "Total results inside of annotations counter"
        CharField color
        ForeignKey created_by_id FK
        IntegerField maximum_annotations "Maximum number of annotations for one task. If the number of annotations per task is equal or greater to this value, the task is completed (is_labeled=True)"
        IntegerField min_annotations_to_start_training "Minimum number of completed tasks after which model training is started"
        JSONField control_weights "Dict of weights for each control tag in metric calculation. Each control tag (e.g. label or choice) will have it's own key in control weight dict with weight for each label and overall weight.For example, if bounding box annotation with control tag named my_bbox should be included with 0.33 weight in agreement calculation, and the first label Car should be twice more important than Airplaine, then you have to need the specify: {'my_bbox': {'type': 'RectangleLabels', 'labels': {'Car': 1.0, 'Airplaine': 0.5}, 'overall': 0.33}"
        TextField model_version "Machine learning model version"
        JSONField data_types
        BooleanField is_draft "Whether or not the project is in the middle of being created"
        BooleanField is_published "Whether or not the project is published to annotators"
        DateTimeField created_at
        DateTimeField updated_at
        ProjectSampling sampling
        ProjectSkipQueue skip_queue
        BooleanField show_ground_truth_first "Onboarding mode (true): show ground truth tasks first in the labeling stream"
        BooleanField annotator_evaluation_enabled "Enable annotator evaluation for the project"
        BooleanField show_overlap_first
        IntegerField overlap_cohort_percentage
        CharField task_data_login "Task data credentials: login"
        CharField task_data_password "Task data credentials: password"
        DateTimeField pinned_at "Pinned date and time"
        IntegerField custom_task_lock_ttl "Custom task lock TTL in seconds. If not set, the default value is used"
        DateTimeField deleted_at
        ForeignKey deleted_by_id FK
        DateTimeField purge_at
        TextField search_vector
    }
    projects_projectonboardingsteps {
        AutoField id PK
        ProjectOnboardingStepsCode code
        CharField title
        TextField description
        IntegerField order
        DateTimeField created_at
        DateTimeField updated_at
    }
    projects_projectonboarding {
        AutoField id PK
        ForeignKey step_id FK
        ForeignKey project_id FK
        BooleanField finished
        DateTimeField created_at
        DateTimeField updated_at
    }
    projects_labelstreamhistory {
        AutoField id PK
        ForeignKey user_id FK "User ID"
        ForeignKey project_id FK "Project ID"
        JSONField data
    }
    projects_projectmember {
        AutoField id PK
        ForeignKey user_id FK "User ID"
        ForeignKey project_id FK "Project ID"
        BooleanField enabled "Project member is enabled"
        DateTimeField created_at
        DateTimeField updated_at
    }
    projects_projectsummary {
        AutoOneToOneField project_id PK
        DateTimeField created_at "Creation time"
        JSONField all_data_columns "All data columns found in imported tasks"
        JSONField common_data_columns "Common data columns found across imported tasks"
        JSONField created_annotations "Unique annotation types identified by tuple (from_name, to_name, type)"
        JSONField created_labels "Unique labels"
        JSONField created_labels_drafts "Unique drafts labels"
    }
    projects_projectimport {
        AutoField id PK
        ForeignKey project_id FK
        JSONField preannotated_from_fields
        BooleanField commit_to_project
        BooleanField return_task_ids
        ProjectImportStatus status
        CharField url
        TextField traceback
        TextField error
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Updated time"
        DateTimeField finished_at "Complete or fail time"
        IntegerField task_count
        IntegerField annotation_count
        IntegerField prediction_count
        IntegerField duration
        JSONField file_upload_ids
        BooleanField could_be_tasks_list
        JSONField found_formats
        JSONField data_columns
        JSONField tasks
        JSONField task_ids
    }
    projects_projectreimport {
        AutoField id PK
        ForeignKey project_id FK
        ProjectReimportStatus status
        TextField error
        IntegerField task_count
        IntegerField annotation_count
        IntegerField prediction_count
        IntegerField duration
        JSONField file_upload_ids
        BooleanField files_as_tasks_list
        JSONField found_formats
        JSONField data_columns
        TextField traceback
    }
    task {
        AutoField id PK
        JSONField data "User imported or uploaded data for a task. Data is formatted according to the project label config. You can find examples of data for your project on the Import page in the Label Studio Data Manager UI."
        JSONField meta "Meta is user imported (uploaded) data and can be useful as input for an ML Backend for embeddings, advanced vectors, and other info. It is passed to ML during training/predicting steps."
        ForeignKey project_id FK "Project ID for this task"
        DateTimeField created_at "Time a task was created"
        DateTimeField updated_at "Last time a task was updated"
        ForeignKey updated_by_id FK "Last annotator or reviewer who updated this task"
        BooleanField is_labeled "True if the number of annotations for this task is greater than or equal to the number of maximum_completions for the project"
        BooleanField allow_skip "Whether this task can be skipped. Set to False to make task unskippable."
        IntegerField overlap "Number of distinct annotators that processed the current task"
        ForeignKey file_upload_id FK "Uploaded file used as data source for this task"
        BigIntegerField inner_id "Internal task ID in the project, starts with 1"
        IntegerField total_annotations "Number of total annotations for the current task except cancelled annotations"
        IntegerField cancelled_annotations "Number of total cancelled annotations for the current task"
        IntegerField total_predictions "Number of total predictions for the current task"
        FloatField precomputed_agreement "Average agreement score for the task"
        IntegerField comment_count "Number of comments in the task including all annotations"
        IntegerField unresolved_comment_count "Number of unresolved comments in the task including all annotations"
        DateTimeField last_comment_updated_at "When the last comment was updated"
    }
    task_completion {
        AutoField id PK
        JSONField result "The main value of annotator work - labeling result in JSON format"
        ForeignKey task_id FK "Corresponding task for this annotation"
        ForeignKey project_id FK "Project ID for this annotation"
        ForeignKey completed_by_id FK "User ID of the person who created this annotation"
        ForeignKey updated_by_id FK "Last user who updated this annotation"
        BooleanField was_cancelled "User skipped the task"
        BooleanField ground_truth "This annotation is a Ground Truth (ground_truth)"
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Last updated time"
        DateTimeField draft_created_at "Draft creation time"
        FloatField lead_time "How much time it took to annotate the task"
        JSONField prediction "Prediction viewed at the time of annotation"
        IntegerField result_count "Results inside of annotation counter"
        ForeignKey parent_prediction_id FK "Points to the prediction from which this annotation was created"
        ForeignKey parent_annotation_id FK "Points to the parent annotation from which this annotation was created"
        UUIDField unique_id UK
        BigIntegerField import_id "Original annotation ID that was at the import step or NULL if this annotation wasn't imported"
        ActionType last_action "Action which was performed in the last annotation history item"
        ForeignKey last_created_by_id FK "User who created the last annotation history item"
        BooleanField bulk_created "Annotation was created in bulk mode"
    }
    tasks_tasklock {
        AutoField id PK
        ForeignKey task_id FK "Locked task"
        DateTimeField expire_at
        UUIDField unique_id UK
        ForeignKey user_id FK "User who locked this task"
        DateTimeField created_at "Creation time"
    }
    tasks_annotationdraft {
        AutoField id PK
        JSONField result "Draft result in JSON format"
        FloatField lead_time "How much time it took to annotate the task"
        ForeignKey task_id FK "Corresponding task for this draft"
        ForeignKey annotation_id FK "Corresponding annotation for this draft"
        ForeignKey user_id FK "User who created this draft"
        BooleanField was_postponed "User postponed this draft (clicked a forward button) in the label stream"
        BigIntegerField import_id "Original draft ID that was at the import step or NULL if this draft wasn't imported"
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Last update time"
    }
    prediction {
        AutoField id PK
        JSONField result "Prediction result"
        FloatField score "Prediction score"
        TextField model_version "A string value that for model version that produced the prediction. Used in both live models and when uploading offline predictions."
        ForeignKey model_id FK "An ML Backend instance that created the prediction."
        ForeignKey model_run_id FK "A run of a ModelVersion that created the prediction."
        IntegerField cluster "Cluster for the current prediction"
        JSONField neighbors "Array of task IDs of the closest neighbors"
        FloatField mislabeling "Related task mislabeling score"
        ForeignKey task_id FK
        ForeignKey project_id FK
        DateTimeField created_at
        DateTimeField updated_at
    }
    tasks_failedprediction {
        AutoField id PK
        TextField message "The message explaining why generating this prediction failed"
        CharField error_type "The type of error that caused prediction to fail"
        ForeignKey ml_backend_model_id FK "An ML Backend instance that created the failed prediction."
        TextField model_version "A string value that for model version that produced the failed prediction. Used in both live models and when uploading offline predictions."
        ForeignKey model_run_id FK "A run of a ModelVersion that created the failed prediction."
        ForeignKey project_id FK
        ForeignKey task_id FK
        DateTimeField created_at
    }
    prediction_meta {
        AutoField id PK
        OneToOneField prediction_id FK, UK "Reference to the associated prediction"
        OneToOneField failed_prediction_id FK, UK "Reference to the associated failed prediction"
        FloatField inference_time "Time taken for inference in seconds"
        IntegerField prompt_tokens_count "Number of tokens in the prompt"
        IntegerField completion_tokens_count "Number of tokens in the completion"
        IntegerField total_tokens_count "Total number of tokens"
        FloatField prompt_cost "Cost of the prompt"
        FloatField completion_cost "Cost of the completion"
        FloatField total_cost "Total cost"
        JSONField extra "Additional metadata in JSON format"
    }
    data_manager_view {
        AutoField id PK
        JSONField data "Custom view data"
        JSONField ordering "Ordering parameters"
        IntegerField order "Position of the tab, starting at the left in data manager and increasing as the tabs go left to right"
        JSONField selected_items "Selected items"
        ForeignKey filter_group_id FK "Groups of filters"
        ForeignKey user_id FK "User who made this view"
        ForeignKey project_id FK "Project ID"
    }
    data_manager_filtergroup {
        AutoField id PK
        CharField conjunction "Type of conjunction"
    }
    data_manager_filter {
        AutoField id PK
        ForeignKey parent_id FK "Optional parent filter to create one-level hierarchy (child filters are AND-merged with parent)"
        IntegerField index "Display order among root filters only"
        CharField column "Field name"
        CharField type "Field type"
        CharField operator "Filter operator"
        JSONField value "Filter value"
    }
    io_storages_azureblobstoragemixin {
        AutoField id PK
        TextField container "Azure blob container"
        TextField prefix "Azure blob prefix name"
        TextField regex_filter "Cloud storage regex for filtering objects"
        BooleanField use_blob_urls "Interpret objects as BLOBs and generate URLs"
        TextField account_name "Azure Blob account name"
        TextField account_key "Azure Blob account key"
    }
    io_storages_azureblobimportstorage {
        OneToOneField azureblobstoragemixin_ptr_id PK
        DateTimeField last_sync "Last sync finished time"
        PositiveIntegerField last_sync_count "Count of tasks synced last time"
        CharField last_sync_job "Last sync job ID"
        AzureBlobImportStorageStatus status
        TextField traceback "Traceback report for the last failed sync"
        JSONField meta "Meta and debug information about storage processes"
        CharField title "Cloud storage title"
        TextField description "Cloud storage description"
        DateTimeField created_at "Creation time"
        BooleanField synchronizable "If storage can be synced"
        ForeignKey project_id FK "A unique integer value identifying this project."
        BooleanField presign "Generate presigned URLs"
        PositiveSmallIntegerField presign_ttl "Presigned URLs TTL (in minutes)"
        BooleanField recursive_scan "Perform recursive scan over the container content"
    }
    io_storages_azureblobexportstorage {
        OneToOneField azureblobstoragemixin_ptr_id PK
        DateTimeField last_sync "Last sync finished time"
        PositiveIntegerField last_sync_count "Count of tasks synced last time"
        CharField last_sync_job "Last sync job ID"
        AzureBlobExportStorageStatus status
        TextField traceback "Traceback report for the last failed sync"
        JSONField meta "Meta and debug information about storage processes"
        CharField title "Cloud storage title"
        TextField description "Cloud storage description"
        DateTimeField created_at "Creation time"
        BooleanField synchronizable "If storage can be synced"
        ForeignKey project_id FK "A unique integer value identifying this project."
        BooleanField can_delete_objects "Deletion from storage enabled"
    }
    io_storages_azureblobimportstoragelink {
        AutoField id PK
        OneToOneField task_id FK, UK
        TextField key "External link key"
        BooleanField object_exists "Whether object under external link still exists"
        DateTimeField created_at "Creation time"
        IntegerField row_group "Parquet row group"
        IntegerField row_index "Parquet row index, or JSON[L] object index"
        ForeignKey storage_id FK
    }
    io_storages_azureblobexportstoragelink {
        AutoField id PK
        ForeignKey annotation_id FK
        BooleanField object_exists "Whether object under external link still exists"
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Update time"
        ForeignKey storage_id FK
    }
    io_storages_s3importstorage {
        AutoField id PK
        DateTimeField last_sync "Last sync finished time"
        PositiveIntegerField last_sync_count "Count of tasks synced last time"
        CharField last_sync_job "Last sync job ID"
        S3ImportStorageStatus status
        TextField traceback "Traceback report for the last failed sync"
        JSONField meta "Meta and debug information about storage processes"
        CharField title "Cloud storage title"
        TextField description "Cloud storage description"
        DateTimeField created_at "Creation time"
        BooleanField synchronizable "If storage can be synced"
        ForeignKey project_id FK "A unique integer value identifying this project."
        TextField bucket "S3 bucket name"
        TextField prefix "S3 bucket prefix"
        TextField regex_filter "Cloud storage regex for filtering objects"
        BooleanField use_blob_urls "Interpret objects as BLOBs and generate URLs"
        TextField aws_access_key_id "AWS_ACCESS_KEY_ID"
        TextField aws_secret_access_key "AWS_SECRET_ACCESS_KEY"
        TextField aws_session_token "AWS_SESSION_TOKEN"
        TextField aws_sse_kms_key_id "AWS SSE KMS Key ID"
        TextField region_name "AWS Region"
        TextField s3_endpoint "S3 Endpoint"
        BooleanField presign "Generate presigned URLs"
        PositiveSmallIntegerField presign_ttl "Presigned URLs TTL (in minutes)"
        BooleanField recursive_scan "Perform recursive scan over the bucket content"
    }
    io_storages_s3exportstorage {
        AutoField id PK
        DateTimeField last_sync "Last sync finished time"
        PositiveIntegerField last_sync_count "Count of tasks synced last time"
        CharField last_sync_job "Last sync job ID"
        S3ExportStorageStatus status
        TextField traceback "Traceback report for the last failed sync"
        JSONField meta "Meta and debug information about storage processes"
        CharField title "Cloud storage title"
        TextField description "Cloud storage description"
        DateTimeField created_at "Creation time"
        BooleanField synchronizable "If storage can be synced"
        ForeignKey project_id FK "A unique integer value identifying this project."
        BooleanField can_delete_objects "Deletion from storage enabled"
        TextField bucket "S3 bucket name"
        TextField prefix "S3 bucket prefix"
        TextField regex_filter "Cloud storage regex for filtering objects"
        BooleanField use_blob_urls "Interpret objects as BLOBs and generate URLs"
        TextField aws_access_key_id "AWS_ACCESS_KEY_ID"
        TextField aws_secret_access_key "AWS_SECRET_ACCESS_KEY"
        TextField aws_session_token "AWS_SESSION_TOKEN"
        TextField aws_sse_kms_key_id "AWS SSE KMS Key ID"
        TextField region_name "AWS Region"
        TextField s3_endpoint "S3 Endpoint"
    }
    io_storages_s3importstoragelink {
        AutoField id PK
        OneToOneField task_id FK, UK
        TextField key "External link key"
        BooleanField object_exists "Whether object under external link still exists"
        DateTimeField created_at "Creation time"
        IntegerField row_group "Parquet row group"
        IntegerField row_index "Parquet row index, or JSON[L] object index"
        ForeignKey storage_id FK
    }
    io_storages_s3exportstoragelink {
        AutoField id PK
        ForeignKey annotation_id FK
        BooleanField object_exists "Whether object under external link still exists"
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Update time"
        ForeignKey storage_id FK
    }
    io_storages_gcsstoragemixin {
        AutoField id PK
        TextField bucket "GCS bucket name"
        TextField prefix "GCS bucket prefix"
        TextField regex_filter "Cloud storage regex for filtering objects"
        BooleanField use_blob_urls "Interpret objects as BLOBs and generate URLs"
        TextField google_application_credentials "The content of GOOGLE_APPLICATION_CREDENTIALS json file"
        TextField google_project_id "Google project ID"
    }
    io_storages_gcsimportstorage {
        OneToOneField gcsstoragemixin_ptr_id PK
        DateTimeField last_sync "Last sync finished time"
        PositiveIntegerField last_sync_count "Count of tasks synced last time"
        CharField last_sync_job "Last sync job ID"
        GCSImportStorageStatus status
        TextField traceback "Traceback report for the last failed sync"
        JSONField meta "Meta and debug information about storage processes"
        CharField title "Cloud storage title"
        TextField description "Cloud storage description"
        DateTimeField created_at "Creation time"
        BooleanField synchronizable "If storage can be synced"
        ForeignKey project_id FK "A unique integer value identifying this project."
        BooleanField presign "Generate presigned URLs"
        PositiveSmallIntegerField presign_ttl "Presigned URLs TTL (in minutes)"
        BooleanField recursive_scan "Perform recursive scan over the bucket content"
    }
    io_storages_gcsexportstorage {
        OneToOneField gcsstoragemixin_ptr_id PK
        DateTimeField last_sync "Last sync finished time"
        PositiveIntegerField last_sync_count "Count of tasks synced last time"
        CharField last_sync_job "Last sync job ID"
        GCSExportStorageStatus status
        TextField traceback "Traceback report for the last failed sync"
        JSONField meta "Meta and debug information about storage processes"
        CharField title "Cloud storage title"
        TextField description "Cloud storage description"
        DateTimeField created_at "Creation time"
        BooleanField synchronizable "If storage can be synced"
        ForeignKey project_id FK "A unique integer value identifying this project."
        BooleanField can_delete_objects "Deletion from storage enabled"
    }
    io_storages_gcsimportstoragelink {
        AutoField id PK
        OneToOneField task_id FK, UK
        TextField key "External link key"
        BooleanField object_exists "Whether object under external link still exists"
        DateTimeField created_at "Creation time"
        IntegerField row_group "Parquet row group"
        IntegerField row_index "Parquet row index, or JSON[L] object index"
        ForeignKey storage_id FK
    }
    io_storages_gcsexportstoragelink {
        AutoField id PK
        ForeignKey annotation_id FK
        BooleanField object_exists "Whether object under external link still exists"
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Update time"
        ForeignKey storage_id FK
    }
    io_storages_redisstoragemixin {
        AutoField id PK
        TextField path "Storage prefix (optional)"
        TextField host "Server Host IP (optional)"
        TextField port "Server Port (optional)"
        TextField password "Server Password (optional)"
        TextField regex_filter "Cloud storage regex for filtering objects"
        BooleanField use_blob_urls "Interpret objects as BLOBs and generate URLs"
    }
    io_storages_redisimportstorage {
        OneToOneField redisstoragemixin_ptr_id PK
        DateTimeField last_sync "Last sync finished time"
        PositiveIntegerField last_sync_count "Count of tasks synced last time"
        CharField last_sync_job "Last sync job ID"
        RedisImportStorageStatus status
        TextField traceback "Traceback report for the last failed sync"
        JSONField meta "Meta and debug information about storage processes"
        CharField title "Cloud storage title"
        TextField description "Cloud storage description"
        DateTimeField created_at "Creation time"
        BooleanField synchronizable "If storage can be synced"
        ForeignKey project_id FK "A unique integer value identifying this project."
        PositiveSmallIntegerField db "Server Database"
    }
    io_storages_redisexportstorage {
        OneToOneField redisstoragemixin_ptr_id PK
        DateTimeField last_sync "Last sync finished time"
        PositiveIntegerField last_sync_count "Count of tasks synced last time"
        CharField last_sync_job "Last sync job ID"
        RedisExportStorageStatus status
        TextField traceback "Traceback report for the last failed sync"
        JSONField meta "Meta and debug information about storage processes"
        CharField title "Cloud storage title"
        TextField description "Cloud storage description"
        DateTimeField created_at "Creation time"
        BooleanField synchronizable "If storage can be synced"
        ForeignKey project_id FK "A unique integer value identifying this project."
        BooleanField can_delete_objects "Deletion from storage enabled"
        PositiveSmallIntegerField db "Server Database"
    }
    io_storages_redisimportstoragelink {
        AutoField id PK
        OneToOneField task_id FK, UK
        TextField key "External link key"
        BooleanField object_exists "Whether object under external link still exists"
        DateTimeField created_at "Creation time"
        IntegerField row_group "Parquet row group"
        IntegerField row_index "Parquet row index, or JSON[L] object index"
        ForeignKey storage_id FK
    }
    io_storages_redisexportstoragelink {
        AutoField id PK
        ForeignKey annotation_id FK
        BooleanField object_exists "Whether object under external link still exists"
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Update time"
        ForeignKey storage_id FK
    }
    ml_mlbackend {
        AutoField id PK
        MLBackendState state
        BooleanField is_interactive "Used to interactively annotate tasks. If true, model returns one list with results"
        TextField url "URL for the machine learning model server"
        TextField error_message "Error message in error state"
        TextField title "Name of the machine learning backend"
        MLBackendAuth auth_method
        TextField basic_auth_user "HTTP Basic Auth user"
        TextField basic_auth_pass "HTTP Basic Auth password"
        TextField description "Description for the machine learning backend"
        JSONField extra_params "Any extra parameters passed to the ML Backend during the setup"
        TextField model_version "Current model version associated with this machine learning backend"
        FloatField timeout "Response model timeout"
        ForeignKey project_id FK
        DateTimeField created_at
        DateTimeField updated_at
        BooleanField auto_update "If false, model version is set by the user, if true - getting latest version from backend."
    }
    ml_mlbackendpredictionjob {
        AutoField id PK
        CharField job_id
        ForeignKey ml_backend_id FK
        TextField model_version "Model version this job is associated with"
        PositiveSmallIntegerField batch_size "Number of tasks processed per batch"
        DateTimeField created_at
        DateTimeField updated_at
    }
    ml_mlbackendtrainjob {
        AutoField id PK
        CharField job_id
        ForeignKey ml_backend_id FK
        TextField model_version "Model version this job is associated with"
        DateTimeField created_at
        DateTimeField updated_at
    }
    webhook {
        AutoField id PK
        ForeignKey organization_id FK
        ForeignKey project_id FK
        URLField url "URL of webhook"
        BooleanField send_payload "If value is False send only action"
        BooleanField send_for_all_actions "If value is False - used only for actions from WebhookAction"
        JSONField headers "Key Value Json of headers"
        BooleanField is_active "If value is False the webhook is disabled"
        DateTimeField created_at "Creation time"
        DateTimeField updated_at "Last update time"
    }
    webhook_action {
        AutoField id PK
        ForeignKey webhook_id FK
        WebhookActionAction action "Action value"
    }
    labels_manager_label {
        AutoField id PK
        DateTimeField created_at "Time of label creation"
        DateTimeField updated_at "Time of label modification"
        ForeignKey created_by_id FK "User who made this label"
        JSONField value "Label value"
        CharField title "Label title"
        TextField description "Label description"
        ForeignKey approved_by_id FK "User who approved this label"
        BooleanField approved "Status of label"
        ForeignKey organization_id FK
    }
    labels_manager_labellink {
        AutoField id PK
        ForeignKey project_id FK
        ForeignKey label_id FK
        CharField from_name "Tag name"
    }
    ml_models_modelinterface {
        AutoField id PK
        CharField title "Model name"
        TextField description "Model description"
        ForeignKey created_by_id FK
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey organization_id FK
        SkillNames skill_name
        JSONField input_fields
        JSONField output_classes
    }
    ml_models_thirdpartymodelversion {
        AutoField id PK
        CharField title "Model name"
        ForeignKey parent_model_id FK
        TextField prompt "Prompt to execute"
        ForeignKey model_provider_connection_id FK
        ModelProviders provider "The model provider to use e.g. OpenAI"
        CharField provider_model_id "The model ID to use within the given provider, e.g. gpt-3.5"
        ForeignKey created_by_id FK
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey organization_id FK
    }
    ml_models_modelrun {
        AutoField id PK
        ForeignKey organization_id FK
        ForeignKey project_id FK
        ForeignKey model_version_id FK
        ForeignKey created_by_id FK
        ModelRunProjectSubset project_subset
        ModelRunModelRunStatus status
        CharField job_id "Job ID for inference job for a ModelRun e.g. Adala job ID"
        IntegerField total_predictions
        IntegerField total_correct_predictions
        IntegerField total_tasks
        DateTimeField created_at
        DateTimeField triggered_at
        DateTimeField predictions_updated_at
        DateTimeField completed_at
    }
    ml_model_providers_modelproviderconnection {
        AutoField id PK
        ModelProviders provider
        TextField api_key "Model provider API key"
        TextField auth_token "Model provider Auth token"
        CharField deployment_name "Azure OpenAI deployment name"
        CharField endpoint "Azure OpenAI endpoint"
        TextField google_application_credentials "The content of GOOGLE_APPLICATION_CREDENTIALS json file"
        CharField google_project_id "Google project ID"
        CharField google_location "Google project location"
        CharField cached_available_models "List of available models from the provider"
        ModelProviderConnectionScopes scope
        ForeignKey organization_id FK
        ForeignKey created_by_id FK
        DateTimeField created_at
        DateTimeField updated_at
        BooleanField is_internal "Whether the model provider connection is internal, not visible to the user"
        FloatField budget_limit "Budget limit for the model provider connection (null if unlimited)"
        DateTimeField budget_last_reset_date "Date and time the budget was last reset"
        ModelProviderConnectionBudgetResetPeriod budget_reset_period "Budget reset period for the model provider connection (null if not reset)"
        FloatField budget_total_spent "Tracked total budget spent for the given provider connection within the current budget period"
        FloatField budget_alert_threshold "Budget alert threshold for the given provider connection"
    }
    jwt_auth_jwtsettings {
        AutoOneToOneField organization_id PK
        BooleanField api_tokens_enabled "Enable JWT API token authentication for this organization"
        IntegerField api_token_ttl_days "Number of days before JWT API tokens expire"
        BooleanField legacy_api_tokens_enabled "Enable legacy API token authentication for this organization"
        DateTimeField created_at
        DateTimeField updated_at
    }
    session_policy_sessiontimeoutpolicy {
        AutoOneToOneField organization_id PK
        IntegerField max_session_age "Number of minutes that a session can be active before needing to re-login"
        IntegerField max_time_between_activity "Number of minutes that a session stays active without any activity"
        DateTimeField created_at
        DateTimeField updated_at
    }
```
