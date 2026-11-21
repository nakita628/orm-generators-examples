```mermaid
erDiagram
    auth_permission ||--o{ guardian_userobjectpermission : "(id) - (permission_id)"
    django_content_type ||--o{ guardian_userobjectpermission : "(id) - (content_type_id)"
    auth_user ||--o{ guardian_userobjectpermission : "(id) - (user_id)"
    auth_permission ||--o{ guardian_groupobjectpermission : "(id) - (permission_id)"
    django_content_type ||--o{ guardian_groupobjectpermission : "(id) - (content_type_id)"
    auth_group ||--o{ guardian_groupobjectpermission : "(id) - (group_id)"
    django_content_type ||--o{ django_comments : "(id) - (content_type_id)"
    django_site ||--o{ django_comments : "(id) - (site_id)"
    auth_user |o--o{ django_comments : "(id) - (user_id)"
    auth_user ||--o{ django_comment_flags : "(id) - (user_id)"
    django_comments ||--o{ django_comment_flags : "(id) - (comment_id)"
    django_content_type ||--o{ attachments_attachment : "(id) - (content_type_id)"
    auth_user ||--o{ attachments_attachment : "(id) - (creator_id)"
    testruns_testexecution ||--o{ linkreference_linkreference : "(id) - (execution_id)"
    management_classification ||--o{ management_product : "(id) - (classification_id)"
    management_product ||--o{ management_component : "(id) - (product_id)"
    auth_user |o--o{ management_component : "(id) - (initial_owner_id)"
    auth_user |o--o{ management_component : "(id) - (initial_qa_contact_id)"
    management_product ||--o{ management_version : "(id) - (product_id)"
    management_version ||--o{ management_build : "(id) - (version_id)"
    management_product ||--o{ testcases_category : "(id) - (product_id)"
    testcases_testcasestatus |o..o{ testcases_historicaltestcase : "(id) - (case_status_id), no constraint"
    testcases_category |o..o{ testcases_historicaltestcase : "(id) - (category_id), no constraint"
    management_priority |o..o{ testcases_historicaltestcase : "(id) - (priority_id), no constraint"
    auth_user |o..o{ testcases_historicaltestcase : "(id) - (author_id), no constraint"
    auth_user |o..o{ testcases_historicaltestcase : "(id) - (default_tester_id), no constraint"
    auth_user |o..o{ testcases_historicaltestcase : "(id) - (reviewer_id), no constraint"
    auth_user |o--o{ testcases_historicaltestcase : "(id) - (history_user_id)"
    testcases_testcasestatus ||--o{ testcases_testcase : "(id) - (case_status_id)"
    testcases_category ||--o{ testcases_testcase : "(id) - (category_id)"
    management_priority ||--o{ testcases_testcase : "(id) - (priority_id)"
    auth_user ||--o{ testcases_testcase : "(id) - (author_id)"
    auth_user |o--o{ testcases_testcase : "(id) - (default_tester_id)"
    auth_user |o--o{ testcases_testcase : "(id) - (reviewer_id)"
    testcases_testcase ||--o{ testcases_property : "(id) - (case_id)"
    testplans_testplan ||--o{ testcases_testcaseplan : "(id) - (plan_id)"
    testcases_testcase ||--o{ testcases_testcaseplan : "(id) - (case_id)"
    testcases_testcase ||--o{ testcases_testcasecomponent : "(id) - (case_id)"
    management_component ||--o{ testcases_testcasecomponent : "(id) - (component_id)"
    management_tag ||--o{ testcases_testcasetag : "(id) - (tag_id)"
    testcases_testcase ||--o{ testcases_testcasetag : "(id) - (case_id)"
    testcases_testcase ||--o| testcases_testcaseemailsettings : "(id) - (case_id)"
    auth_user |o--o{ testcases_historicaltemplate : "(id) - (history_user_id)"
    testplans_testplan |o..o{ testplans_historicaltestplan : "(id) - (parent_id), no constraint"
    management_version |o..o{ testplans_historicaltestplan : "(id) - (product_version_id), no constraint"
    auth_user |o..o{ testplans_historicaltestplan : "(id) - (author_id), no constraint"
    management_product |o..o{ testplans_historicaltestplan : "(id) - (product_id), no constraint"
    testplans_plantype |o..o{ testplans_historicaltestplan : "(id) - (type_id), no constraint"
    auth_user |o--o{ testplans_historicaltestplan : "(id) - (history_user_id)"
    testplans_testplan |o--o{ testplans_testplan : "(id) - (parent_id)"
    management_version ||--o{ testplans_testplan : "(id) - (product_version_id)"
    auth_user ||--o{ testplans_testplan : "(id) - (author_id)"
    management_product ||--o{ testplans_testplan : "(id) - (product_id)"
    testplans_plantype ||--o{ testplans_testplan : "(id) - (type_id)"
    management_tag ||--o{ testplans_testplantag : "(id) - (tag_id)"
    testplans_testplan ||--o{ testplans_testplantag : "(id) - (plan_id)"
    testplans_testplan ||--o| testplans_testplanemailsettings : "(id) - (plan_id)"
    testplans_testplan |o..o{ testruns_historicaltestrun : "(id) - (plan_id), no constraint"
    management_build |o..o{ testruns_historicaltestrun : "(id) - (build_id), no constraint"
    auth_user |o..o{ testruns_historicaltestrun : "(id) - (manager_id), no constraint"
    auth_user |o..o{ testruns_historicaltestrun : "(id) - (default_tester_id), no constraint"
    auth_user |o--o{ testruns_historicaltestrun : "(id) - (history_user_id)"
    testplans_testplan ||--o{ testruns_testrun : "(id) - (plan_id)"
    management_build ||--o{ testruns_testrun : "(id) - (build_id)"
    auth_user ||--o{ testruns_testrun : "(id) - (manager_id)"
    auth_user |o--o{ testruns_testrun : "(id) - (default_tester_id)"
    auth_user |o..o{ testruns_historicaltestexecution : "(id) - (assignee_id), no constraint"
    auth_user |o..o{ testruns_historicaltestexecution : "(id) - (tested_by_id), no constraint"
    testruns_testrun |o..o{ testruns_historicaltestexecution : "(id) - (run_id), no constraint"
    testcases_testcase |o..o{ testruns_historicaltestexecution : "(id) - (case_id), no constraint"
    testruns_testexecutionstatus |o..o{ testruns_historicaltestexecution : "(id) - (status_id), no constraint"
    management_build |o..o{ testruns_historicaltestexecution : "(id) - (build_id), no constraint"
    auth_user |o--o{ testruns_historicaltestexecution : "(id) - (history_user_id)"
    auth_user |o--o{ testruns_testexecution : "(id) - (assignee_id)"
    auth_user |o--o{ testruns_testexecution : "(id) - (tested_by_id)"
    testruns_testrun ||--o{ testruns_testexecution : "(id) - (run_id)"
    testcases_testcase ||--o{ testruns_testexecution : "(id) - (case_id)"
    testruns_testexecutionstatus ||--o{ testruns_testexecution : "(id) - (status_id)"
    management_build ||--o{ testruns_testexecution : "(id) - (build_id)"
    testruns_testexecution ||--o{ testruns_testexecutionproperty : "(id) - (execution_id)"
    management_tag ||--o{ testruns_testruntag : "(id) - (tag_id)"
    testruns_testrun ||--o{ testruns_testruntag : "(id) - (run_id)"
    testruns_testrun ||--o{ testruns_testruncc : "(id) - (run_id)"
    auth_user ||--o{ testruns_testruncc : "(id) - (user_id)"
    testruns_environment ||--o{ testruns_environmentproperty : "(id) - (environment_id)"
    testruns_testrun ||--o{ testruns_property : "(id) - (run_id)"
    auth_user ||--o{ bugs_bug : "(id) - (reporter_id)"
    auth_user |o--o{ bugs_bug : "(id) - (assignee_id)"
    management_product ||--o{ bugs_bug : "(id) - (product_id)"
    management_version ||--o{ bugs_bug : "(id) - (version_id)"
    management_build ||--o{ bugs_bug : "(id) - (build_id)"
    bugs_severity |o--o{ bugs_bug : "(id) - (severity_id)"
    bugs_bug }o--o{ testruns_testexecution : "(executions) - (bugs)"
    bugs_bug }o--o{ management_tag : "(tags) - (bugs)"
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user ||--o{ kiwi_auth_useractivationkey : "(id) - (user_id)"
    django_site {
        AutoField id PK
        CharField domain UK
        CharField name
    }
    guardian_userobjectpermission {
        AutoField id PK
        ForeignKey permission_id FK
        ForeignKey content_type_id FK
        CharField object_pk
        ForeignKey user_id FK
    }
    guardian_groupobjectpermission {
        AutoField id PK
        ForeignKey permission_id FK
        ForeignKey content_type_id FK
        CharField object_pk
        ForeignKey group_id FK
    }
    django_comments {
        AutoField id PK
        ForeignKey content_type_id FK
        CharField object_pk
        ForeignKey site_id FK
        ForeignKey user_id FK
        CharField user_name
        EmailField user_email
        URLField user_url
        TextField comment
        DateTimeField submit_date
        GenericIPAddressField ip_address
        BooleanField is_public "Uncheck this box to make the comment effectively disappear from the site."
        BooleanField is_removed "Check this box if the comment is inappropriate. A 'This comment has been removed' message will be displayed instead."
    }
    django_comment_flags {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey comment_id FK
        CharField flag
        DateTimeField flag_date
    }
    attachments_attachment {
        AutoField id PK
        ForeignKey content_type_id FK
        CharField object_id
        ForeignKey creator_id FK
        FileField attachment_file
        DateTimeField created
        DateTimeField modified
    }
    linkreference_linkreference {
        AutoField id PK
        ForeignKey execution_id FK
        CharField name
        URLField url
        DateTimeField created_on
        BooleanField is_defect
    }
    management_classification {
        AutoField id PK
        CharField name UK
    }
    management_product {
        AutoField id PK
        CharField name UK
        ForeignKey classification_id FK
        TextField description
    }
    management_priority {
        AutoField id PK
        CharField value UK
        BooleanField is_active
    }
    management_component {
        AutoField id PK
        CharField name
        ForeignKey product_id FK
        ForeignKey initial_owner_id FK
        ForeignKey initial_qa_contact_id FK
        TextField description
    }
    management_version {
        AutoField id PK
        CharField value
        ForeignKey product_id FK
    }
    management_build {
        AutoField id PK
        CharField name
        ForeignKey version_id FK
        BooleanField is_active
    }
    management_tag {
        AutoField id PK
        CharField name
    }
    testcases_testcasestatus {
        AutoField id PK
        CharField name
        TextField description
        BooleanField is_confirmed
    }
    testcases_category {
        AutoField id PK
        CharField name
        ForeignKey product_id FK
        TextField description
    }
    testcases_historicaltestcase {
        IntegerField id
        DateTimeField create_date
        BooleanField is_automated
        TextField script
        TextField arguments
        CharField extra_link
        CharField summary
        CharField requirement
        TextField notes
        TextField text
        DurationField setup_duration
        DurationField testing_duration
        ForeignKey case_status_id FK
        ForeignKey category_id FK
        ForeignKey priority_id FK
        ForeignKey author_id FK
        ForeignKey default_tester_id FK
        ForeignKey reviewer_id FK
        AutoField history_id PK
        DateTimeField history_date
        TextField history_change_reason
        HistoricalTestCaseHistoryType history_type
        ForeignKey history_user_id FK
    }
    testcases_testcase {
        AutoField id PK
        DateTimeField create_date
        BooleanField is_automated
        TextField script
        TextField arguments
        CharField extra_link
        CharField summary
        CharField requirement
        TextField notes
        TextField text
        DurationField setup_duration
        DurationField testing_duration
        ForeignKey case_status_id FK
        ForeignKey category_id FK
        ForeignKey priority_id FK
        ForeignKey author_id FK
        ForeignKey default_tester_id FK
        ForeignKey reviewer_id FK
    }
    testcases_property {
        AutoField id PK
        CharField name
        CharField value
        ForeignKey case_id FK
    }
    testcases_testcaseplan {
        AutoField id PK
        ForeignKey plan_id FK
        ForeignKey case_id FK
        IntegerField sortkey
    }
    testcases_testcasecomponent {
        AutoField id PK
        ForeignKey case_id FK
        ForeignKey component_id FK
    }
    testcases_testcasetag {
        AutoField id PK
        ForeignKey tag_id FK
        ForeignKey case_id FK
    }
    testcases_bugsystem {
        AutoField id PK
        CharField name UK
        CharField tracker_type "This determines how Kiwi TCMS integrates with the IT system"
        CharField base_url "Base URL, for example <strong>https://bugzilla.example.com</strong>! Leave empty to disable!"
        CharField api_url "This is the URL to which API requests will be sent. Leave empty to disable!"
        CharField api_username
        CharField api_password
    }
    testcases_testcaseemailsettings {
        AutoField id PK
        OneToOneField case_id FK, UK
        BooleanField notify_on_case_update
        BooleanField notify_on_case_delete
        BooleanField auto_to_case_author
        BooleanField auto_to_case_tester
        BooleanField auto_to_run_manager
        BooleanField auto_to_run_tester
        BooleanField auto_to_execution_assignee
        TextField cc_list
    }
    testcases_historicaltemplate {
        IntegerField id
        CharField name
        TextField text
        AutoField history_id PK
        DateTimeField history_date
        TextField history_change_reason
        HistoricalTemplateHistoryType history_type
        ForeignKey history_user_id FK
    }
    testcases_template {
        AutoField id PK
        CharField name
        TextField text
    }
    testplans_plantype {
        AutoField id PK
        CharField name UK
        TextField description
    }
    testplans_historicaltestplan {
        IntegerField id
        CharField name
        TextField text
        DateTimeField create_date
        BooleanField is_active
        CharField extra_link
        TreeNodeForeignKey parent_id FK
        ForeignKey product_version_id FK
        ForeignKey author_id FK
        ForeignKey product_id FK
        ForeignKey type_id FK
        AutoField history_id PK
        DateTimeField history_date
        TextField history_change_reason
        HistoricalTestPlanHistoryType history_type
        ForeignKey history_user_id FK
    }
    testplans_testplan {
        AutoField id PK
        TreeNodeForeignKey parent_id FK
        CharField name
        TextField text
        DateTimeField create_date
        BooleanField is_active
        CharField extra_link
        ForeignKey product_version_id FK
        ForeignKey author_id FK
        ForeignKey product_id FK
        ForeignKey type_id FK
    }
    testplans_testplantag {
        AutoField id PK
        ForeignKey tag_id FK
        ForeignKey plan_id FK
    }
    testplans_testplanemailsettings {
        AutoField id PK
        OneToOneField plan_id FK, UK
        BooleanField auto_to_plan_author
        BooleanField auto_to_case_owner
        BooleanField auto_to_case_default_tester
        BooleanField notify_on_plan_update
        BooleanField notify_on_case_update
    }
    testruns_historicaltestrun {
        IntegerField id
        DateTimeField start_date
        DateTimeField stop_date
        DateTimeField planned_start
        DateTimeField planned_stop
        TextField summary
        TextField notes
        ForeignKey plan_id FK
        ForeignKey build_id FK
        ForeignKey manager_id FK
        ForeignKey default_tester_id FK
        AutoField history_id PK
        DateTimeField history_date
        TextField history_change_reason
        HistoricalTestRunHistoryType history_type
        ForeignKey history_user_id FK
    }
    testruns_testrun {
        AutoField id PK
        DateTimeField start_date
        DateTimeField stop_date
        DateTimeField planned_start
        DateTimeField planned_stop
        TextField summary
        TextField notes
        ForeignKey plan_id FK
        ForeignKey build_id FK
        ForeignKey manager_id FK
        ForeignKey default_tester_id FK
    }
    testruns_testexecutionstatus {
        AutoField id PK
        CharField name UK
        IntegerField weight
        CharField icon
        ColorField color
    }
    testruns_historicaltestexecution {
        IntegerField id
        IntegerField case_text_version
        DateTimeField start_date
        DateTimeField stop_date
        IntegerField sortkey
        ForeignKey assignee_id FK
        ForeignKey tested_by_id FK
        ForeignKey run_id FK
        ForeignKey case_id FK
        ForeignKey status_id FK
        ForeignKey build_id FK
        AutoField history_id PK
        DateTimeField history_date
        TextField history_change_reason
        HistoricalTestExecutionHistoryType history_type
        ForeignKey history_user_id FK
    }
    testruns_testexecution {
        AutoField id PK
        ForeignKey assignee_id FK
        ForeignKey tested_by_id FK
        IntegerField case_text_version
        DateTimeField start_date
        DateTimeField stop_date
        IntegerField sortkey
        ForeignKey run_id FK
        ForeignKey case_id FK
        ForeignKey status_id FK
        ForeignKey build_id FK
    }
    testruns_testexecutionproperty {
        AutoField id PK
        CharField name
        CharField value
        ForeignKey execution_id FK
    }
    testruns_testruntag {
        AutoField id PK
        ForeignKey tag_id FK
        ForeignKey run_id FK
    }
    testruns_testruncc {
        AutoField id PK
        ForeignKey run_id FK
        ForeignKey user_id FK
    }
    testruns_environment {
        AutoField id PK
        CharField name UK
        TextField description
    }
    testruns_environmentproperty {
        AutoField id PK
        CharField name
        CharField value
        ForeignKey environment_id FK
    }
    testruns_property {
        AutoField id PK
        CharField name
        CharField value
        ForeignKey run_id FK
    }
    bugs_severity {
        AutoField id PK
        CharField name UK
        IntegerField weight
        CharField icon
        ColorField color
    }
    bugs_bug {
        AutoField id PK
        CharField summary
        DateTimeField created_at
        BooleanField status
        ForeignKey reporter_id FK
        ForeignKey assignee_id FK
        ForeignKey product_id FK
        ForeignKey version_id FK
        ForeignKey build_id FK
        ForeignKey severity_id FK
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
    captcha_captchastore {
        AutoField id PK
        CharField challenge
        CharField response
        CharField hashkey UK
        DateTimeField expiration
    }
    kiwi_auth_useractivationkey {
        AutoField id PK
        ForeignKey user_id FK
        CharField activation_key
        DateTimeField key_expires
    }
```
