```mermaid
erDiagram
    trans_component |o--o{ addons_addon : "(id) - (component_id)"
    trans_project |o--o{ addons_addon : "(id) - (project_id)"
    trans_category |o--o{ addons_addon : "(id) - (category_id)"
    addons_addon ||--o{ addons_event : "(id) - (addon_id)"
    addons_addon ||--o{ addons_addonactivitylog : "(id) - (addon_id)"
    trans_component |o--o{ addons_addonactivitylog : "(id) - (component_id)"
    kotlin_sdk_kotlinsdkbuild |o--o| kotlin_sdk_kotlinsdkcleanup : "(id) - (build_id)"
    addons_addon ||--o{ kotlin_sdk_kotlinsdkbuild : "(id) - (addon_id)"
    addons_addon ||--o{ kotlin_sdk_kotlinsdkartifact : "(id) - (addon_id)"
    kotlin_sdk_kotlinsdkartifact }o--o{ kotlin_sdk_kotlinsdkbuild : "(builds) - (artifacts)"
    weblate_auth_role }o--o{ weblate_auth_permission : "(permissions) - (role_set)"
    trans_project |o--o{ weblate_auth_group : "(id) - (defining_project_id)"
    workspaces_workspace |o--o{ weblate_auth_group : "(id) - (defining_workspace_id)"
    weblate_auth_group }o--o{ weblate_auth_role : "(roles) - (group_set)"
    weblate_auth_group }o--o{ trans_project : "(projects) - (group_set)"
    weblate_auth_group }o--o{ trans_component : "(components) - (group_set)"
    weblate_auth_group }o--o{ trans_componentlist : "(componentlists) - (group_set)"
    weblate_auth_group }o--o{ lang_language : "(languages) - (group_set)"
    weblate_auth_group }o--o{ weblate_auth_user : "(admins) - (administered_group_set)"
    weblate_auth_user ||--o{ weblate_auth_user_groups : "(id) - (user_id)"
    weblate_auth_group ||--o{ weblate_auth_user_groups : "(id) - (group_id)"
    weblate_auth_user_groups }o--o{ lang_language : "(limit_languages) - (teammembership_set)"
    weblate_auth_group ||--o{ weblate_auth_autogroup : "(id) - (group_id)"
    weblate_auth_user ||--o{ weblate_auth_userblock : "(id) - (user_id)"
    trans_project ||--o{ weblate_auth_userblock : "(id) - (project_id)"
    weblate_auth_user ||--o{ weblate_auth_invitation : "(id) - (author_id)"
    weblate_auth_user |o--o{ weblate_auth_invitation : "(id) - (user_id)"
    weblate_auth_group ||--o{ weblate_auth_invitation : "(id) - (group_id)"
    weblate_auth_invitation }o--o{ lang_language : "(limit_languages) - (invitation_set)"
    trans_unit ||--o{ checks_check : "(id) - (unit_id)"
    workspaces_workspace |o--o{ machinery_machineryinstallation : "(id) - (workspace_id)"
    trans_project |o--o{ machinery_machineryinstallation : "(id) - (project_id)"
    machinery_machineryinstallation ||--o{ machinery_machineryerror : "(id) - (installation_id)"
    workspaces_workspace |o--o{ machinery_machineryexclusion : "(id) - (workspace_id)"
    trans_project |o--o{ machinery_machineryexclusion : "(id) - (project_id)"
    machinery_machineryinstallation |o--o{ machinery_machineryexclusion : "(id) - (installation_id)"
    lang_language |o--o{ workspaces_workspace : "(id) - (secondary_language_id)"
    weblate_auth_user ||--o{ trans_contributoragreement : "(id) - (user_id)"
    trans_component |o--o{ trans_contributoragreement : "(id) - (component_id)"
    trans_project |o--o{ trans_contributoragreement : "(id) - (project_id)"
    trans_category |o--o{ trans_contributoragreement : "(id) - (category_id)"
    workspaces_workspace |o--o{ trans_contributoragreement : "(id) - (workspace_id)"
    trans_component ||--o{ trans_alert : "(id) - (component_id)"
    weblate_auth_user |o--o{ trans_alert : "(id) - (dismissed_by_id)"
    trans_project |o--o{ trans_announcement : "(id) - (project_id)"
    trans_category |o--o{ trans_announcement : "(id) - (category_id)"
    trans_component |o--o{ trans_announcement : "(id) - (component_id)"
    lang_language |o--o{ trans_announcement : "(id) - (language_id)"
    trans_project ||--o{ trans_backupauthorclaim : "(id) - (project_id)"
    weblate_auth_user ||--o{ trans_backupauthorclaim : "(id) - (user_id)"
    weblate_auth_invitation |o--o| trans_backupauthorclaim : "(uuid) - (invitation_id)"
    trans_backupauthorclaim ||--o{ trans_backupauthorvote : "(id) - (claim_id)"
    trans_suggestion ||--o{ trans_backupauthorvote : "(id) - (suggestion_id)"
    trans_vote |o--o{ trans_backupauthorvote : "(id) - (placeholder_id)"
    trans_project ||--o{ trans_category : "(id) - (project_id)"
    trans_category |o--o{ trans_category : "(id) - (category_id)"
    lang_language |o--o{ trans_category : "(id) - (secondary_language_id)"
    workspaces_workspace |o--o{ trans_project : "(id) - (workspace_id)"
    lang_language |o--o{ trans_project : "(id) - (secondary_language_id)"
    trans_backupauthorclaim |o--o{ trans_change : "(id) - (backup_author_claim_id)"
    trans_unit |o--o{ trans_change : "(id) - (unit_id)"
    lang_language |o--o{ trans_change : "(id) - (language_id)"
    trans_project |o--o{ trans_change : "(id) - (project_id)"
    workspaces_workspace |o--o{ trans_change : "(id) - (workspace_id)"
    trans_category |o--o{ trans_change : "(id) - (category_id)"
    trans_component |o--o{ trans_change : "(id) - (component_id)"
    trans_translation |o--o{ trans_change : "(id) - (translation_id)"
    trans_comment |o--o{ trans_change : "(id) - (comment_id)"
    trans_suggestion |o--o{ trans_change : "(id) - (suggestion_id)"
    trans_announcement |o--o{ trans_change : "(id) - (announcement_id)"
    screenshots_screenshot |o--o{ trans_change : "(id) - (screenshot_id)"
    trans_alert |o--o{ trans_change : "(id) - (alert_id)"
    addons_addon |o--o{ trans_change : "(id) - (addon_id)"
    weblate_auth_user |o--o{ trans_change : "(id) - (user_id)"
    weblate_auth_user |o--o{ trans_change : "(id) - (author_id)"
    trans_backupauthorclaim |o--o{ trans_pendingunitchange : "(id) - (backup_author_claim_id)"
    trans_unit ||--o{ trans_pendingunitchange : "(id) - (unit_id)"
    weblate_auth_user ||--o{ trans_pendingunitchange : "(id) - (author_id)"
    trans_backupauthorclaim |o--o{ trans_comment : "(id) - (backup_author_claim_id)"
    trans_unit ||--o{ trans_comment : "(id) - (unit_id)"
    weblate_auth_user |o--o{ trans_comment : "(id) - (user_id)"
    trans_backupauthorclaim |o--o{ trans_suggestion : "(id) - (backup_author_claim_id)"
    trans_unit ||--o{ trans_suggestion : "(id) - (unit_id)"
    weblate_auth_user |o--o{ trans_suggestion : "(id) - (user_id)"
    trans_suggestion ||--o{ trans_vote : "(id) - (suggestion_id)"
    weblate_auth_user ||--o{ trans_vote : "(id) - (user_id)"
    trans_component ||--o{ trans_variant : "(id) - (component_id)"
    trans_variant }o--o{ trans_unit : "(defining_units) - (defined_variants)"
    trans_translation ||--o{ trans_unit : "(id) - (translation_id)"
    trans_variant |o--o{ trans_unit : "(id) - (variant_id)"
    trans_unit |o--o{ trans_unit : "(id) - (source_unit_id)"
    trans_unit |o--o{ trans_unit : "(id) - (translation_parent_id)"
    trans_unit }o--o{ trans_label : "(labels) - (unit_set)"
    trans_component ||--o{ trans_translation : "(id) - (component_id)"
    lang_language ||--o{ trans_translation : "(id) - (language_id)"
    lang_plural ||--o{ trans_translation : "(id) - (plural_id)"
    trans_component ||--o{ trans_component_links : "(id) - (component_id)"
    trans_project ||--o{ trans_component_links : "(id) - (project_id)"
    trans_category |o--o{ trans_component_links : "(id) - (category_id)"
    trans_project ||--o{ trans_component : "(id) - (project_id)"
    trans_category |o--o{ trans_component : "(id) - (category_id)"
    trans_component |o--o{ trans_component : "(id) - (linked_component_id)"
    lang_language ||--o{ trans_component : "(id) - (source_language_id)"
    lang_language |o--o{ trans_component : "(id) - (secondary_language_id)"
    trans_componentlist }o--o{ trans_component : "(components) - (componentlist_set)"
    trans_componentlist ||--o{ trans_autocomponentlist : "(id) - (componentlist_id)"
    trans_project ||--o{ trans_label : "(id) - (project_id)"
    weblate_auth_user ||--o{ trans_report : "(id) - (creator_id)"
    workspaces_workspace |o--o{ trans_report : "(id) - (workspace_id)"
    trans_project |o--o{ trans_report : "(id) - (project_id)"
    trans_category |o--o{ trans_report : "(id) - (category_id)"
    trans_component |o--o{ trans_report : "(id) - (component_id)"
    trans_project |o--o{ trans_savedsearch : "(id) - (project_id)"
    weblate_auth_user |o--o{ trans_savedsearch : "(id) - (user_id)"
    trans_project |o--o{ trans_workflowsetting : "(id) - (project_id)"
    lang_language ||--o{ trans_workflowsetting : "(id) - (language_id)"
    lang_language |o--o{ trans_workflowsetting : "(id) - (source_language_id)"
    lang_language ||--o{ lang_plural : "(id) - (language_id)"
    lang_language ||--o{ memory_memory : "(id) - (source_language_id)"
    lang_language ||--o{ memory_memory : "(id) - (target_language_id)"
    weblate_auth_user |o--o{ memory_memory : "(id) - (legacy_user_id)"
    trans_project |o--o{ memory_memory : "(id) - (legacy_project_id)"
    memory_memory ||--o{ memory_memoryscope : "(id) - (memory_id)"
    trans_project |o--o{ memory_memoryscope : "(id) - (project_id)"
    workspaces_workspace |o--o{ memory_memoryscope : "(id) - (workspace_id)"
    trans_project |o--o{ memory_memoryscope : "(id) - (source_project_id)"
    trans_component |o--o{ memory_memoryscope : "(id) - (source_component_id)"
    weblate_auth_user |o--o{ memory_memoryscope : "(id) - (user_id)"
    trans_backupauthorclaim |o--o{ screenshots_screenshot : "(id) - (backup_author_claim_id)"
    trans_translation ||--o{ screenshots_screenshot : "(id) - (translation_id)"
    weblate_auth_user |o--o{ screenshots_screenshot : "(id) - (user_id)"
    screenshots_screenshot }o--o{ trans_unit : "(units) - (screenshots)"
    screenshots_screenshot ||--o| screenshots_screenshotocr : "(id) - (screenshot_id)"
    trans_project ||--o{ fonts_font : "(id) - (project_id)"
    weblate_auth_user |o--o{ fonts_font : "(id) - (user_id)"
    fonts_font ||--o{ fonts_fontgroup : "(id) - (font_id)"
    trans_project ||--o{ fonts_fontgroup : "(id) - (project_id)"
    fonts_fontgroup ||--o{ fonts_fontoverride : "(id) - (group_id)"
    fonts_font ||--o{ fonts_fontoverride : "(id) - (font_id)"
    lang_language ||--o{ fonts_fontoverride : "(id) - (language_id)"
    weblate_auth_user ||--o{ accounts_subscription : "(id) - (user_id)"
    trans_project |o--o{ accounts_subscription : "(id) - (project_id)"
    trans_component |o--o{ accounts_subscription : "(id) - (component_id)"
    weblate_auth_group |o--o{ accounts_teamauditlog : "(id) - (team_id)"
    trans_project |o--o{ accounts_teamauditlog : "(id) - (project_id)"
    workspaces_workspace |o--o{ accounts_teamauditlog : "(id) - (workspace_id)"
    weblate_auth_user |o--o{ accounts_teamauditlog : "(id) - (actor_id)"
    weblate_auth_user |o--o{ accounts_auditlog : "(id) - (user_id)"
    social_auth_usersocialauth ||--o{ accounts_verifiedemail : "(id) - (social_id)"
    weblate_auth_user ||--o| accounts_profile : "(id) - (user_id)"
    trans_componentlist |o--o{ accounts_profile : "(id) - (dashboard_component_list_id)"
    accounts_profile }o--o{ lang_language : "(languages) - (profile_set)"
    accounts_profile }o--o{ lang_language : "(secondary_languages) - (secondary_profile_set)"
    accounts_profile }o--o{ trans_project : "(watched) - (profile_set)"
    workspaces_workspace ||--o{ vcs_installation : "(id) - (workspace_id)"
    weblate_auth_user |o--o{ wladmin_securityscan : "(id) - (actor_id)"
    weblate_auth_user ||--o{ wladmin_securityfinding : "(id) - (user_id)"
    wladmin_securityfinding ||--o{ wladmin_securityfindingevent : "(id) - (finding_id)"
    wladmin_securityscan ||--o{ wladmin_securityfindingevent : "(id) - (scan_id)"
    weblate_auth_user |o--o{ wladmin_securityfindingevent : "(id) - (actor_id)"
    wladmin_backupservice ||--o{ wladmin_backuplog : "(id) - (service_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    weblate_auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    weblate_auth_user ||--o{ social_auth_usersocialauth : "(id) - (user_id)"
    weblate_auth_user ||--o| authtoken_token : "(id) - (user_id)"
    django_celery_beat_intervalschedule |o--o{ django_celery_beat_periodictask : "(id) - (interval_id)"
    django_celery_beat_crontabschedule |o--o{ django_celery_beat_periodictask : "(id) - (crontab_id)"
    django_celery_beat_solarschedule |o--o{ django_celery_beat_periodictask : "(id) - (solar_id)"
    django_celery_beat_clockedschedule |o--o{ django_celery_beat_periodictask : "(id) - (clocked_id)"
    weblate_auth_user ||--o{ otp_static_staticdevice : "(id) - (user_id)"
    otp_static_staticdevice ||--o{ otp_static_statictoken : "(id) - (device_id)"
    weblate_auth_user ||--o{ otp_totp_totpdevice : "(id) - (user_id)"
    weblate_auth_user ||--o{ django_otp_webauthn_webauthncredential : "(id) - (user_id)"
    django_otp_webauthn_webauthncredential ||--o| django_otp_webauthn_webauthnattestation : "(id) - (credential_id)"
    weblate_auth_user ||--o| django_otp_webauthn_webauthnuserhandle : "(id) - (user_id)"
    addons_addon {
        AutoField id PK
        ForeignKey component_id FK
        ForeignKey project_id FK
        ForeignKey category_id FK
        CharField name
        JSONField configuration
        JSONField state
        BooleanField repo_scope
    }
    addons_event {
        AutoField id PK
        ForeignKey addon_id FK
        AddonEvent event
    }
    addons_addonactivitylog {
        AutoField id PK
        ForeignKey addon_id FK
        ForeignKey component_id FK
        AddonEvent event
        DateTimeField created
        JSONField details
        AddonActivityLogStatus status
    }
    kotlin_sdk_kotlinsdkcleanup {
        AutoField id PK
        TextField path UK
        OneToOneField build_id FK, UK
    }
    kotlin_sdk_kotlinsdkbuild {
        AutoField id PK
        ForeignKey addon_id FK
        CharField package_name
        PositiveBigIntegerField version_code
        JSONField metadata
        JSONField pending_manifest
        CharField digest
        DateTimeField created
        DateTimeField retired
        CharField status
        TextField error
        DateTimeField published
    }
    kotlin_sdk_kotlinsdkartifact {
        AutoField id PK
        ForeignKey addon_id FK
        CharField digest
        DateTimeField unreferenced
    }
    weblate_auth_permission {
        AutoField id PK
        CharField codename UK
        CharField name
    }
    weblate_auth_role {
        AutoField id PK
        CharField name UK
    }
    weblate_auth_group {
        AutoField id PK
        CharField name
        ForeignKey defining_project_id FK
        ForeignKey defining_workspace_id FK
        WeblateAuthGroupProjectSelection project_selection
        WeblateAuthGroupLanguageSelection language_selection
        BooleanField internal
        BooleanField enforced_2fa "Requires users to have two-factor authentication configured."
    }
    weblate_auth_user_groups {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey group_id FK
    }
    weblate_auth_user {
        AutoField id PK
        CharField password
        DateTimeField last_login
        UsernameField username UK "Username may only contain letters, numbers or the following characters: @ . + - _"
        CharField full_name
        EmailField email UK
        BooleanField is_superuser "User has all possible permissions."
        BooleanField is_active "Mark user as inactive instead of removing."
        BooleanField is_bot
        DateTimeField date_expires "The account will be disabled after the expiry."
        DateTimeField date_joined
    }
    weblate_auth_autogroup {
        AutoField id PK
        RegexField match "Users with e-mail addresses found to match will be added to this team."
        ForeignKey group_id FK
    }
    weblate_auth_userblock {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey project_id FK
        DateTimeField expiry
        TextField note "Internal notes regarding blocking the user that are not visible to the user."
    }
    weblate_auth_invitation {
        UUIDField uuid PK
        BooleanField backup_authorship
        DateTimeField timestamp
        ForeignKey author_id FK
        ForeignKey user_id FK "Please type in an existing Weblate account name or e-mail address."
        UsernameField username "Suggest username for the user. It can be changed later."
        CharField full_name "Suggest full name for the user. It can be changed later."
        ForeignKey group_id FK "The user is granted all permissions included in membership of these teams."
        EmailField email
        BooleanField automatic
        BooleanField team_admin
        BooleanField is_superuser "User has all possible permissions."
    }
    checks_check {
        AutoField id PK
        ForeignKey unit_id FK
        CheckName name
        BooleanField dismissed
        JSONField metadata
    }
    machinery_machineryinstallation {
        AutoField id PK
        CharField engine
        SlugField slug
        JSONField configuration
        BooleanField enabled
        ForeignKey workspace_id FK
        ForeignKey project_id FK
    }
    machinery_machineryerror {
        AutoField id PK
        ForeignKey installation_id FK
        DateTimeField timestamp
        TextField error
    }
    machinery_machineryexclusion {
        AutoField id PK
        ForeignKey workspace_id FK
        ForeignKey project_id FK
        ForeignKey installation_id FK
        CharField engine
    }
    workspaces_workspace {
        UUIDField id PK
        IntegerField metric_id UK
        CharField name "Display name"
        WorkspacePropagationScope propagation_scope "Where translations of matching strings are propagated."
        BooleanField name_managed "Whether Weblate can update the name automatically."
        BooleanField use_workspace_tm "Uses the pool of shared translations between projects in this workspace."
        BooleanField contribute_workspace_tm "Contributes translations to the pool shared between projects in this workspace."
        WorkspaceLicense license
        TextField agreement "Contributor license agreement which needs to be approved before a user can translate components in this workspace."
        WorkspaceNewLang new_lang "How to handle requests for creating new translations."
        WorkspaceLanguageCodeStyle language_code_style "Customize language code used to generate the filename for translations created by Weblate."
        ForeignKey secondary_language_id FK "Additional language to show together with the source language while translating."
        TextField check_flags "Additional comma-separated flags to influence Weblate behavior."
        JSONField enforced_checks "List of checks which can not be dismissed."
        BooleanField inherit_enforced_checks "Use enforced checks from the workspace."
        TextField commit_message "You can use template language for various info, please consult the documentation for more details."
        TextField add_message "You can use template language for various info, please consult the documentation for more details."
        TextField delete_message "You can use template language for various info, please consult the documentation for more details."
        TextField merge_message "You can use template language for various info, please consult the documentation for more details."
        TextField addon_message "You can use template language for various info, please consult the documentation for more details."
        TextField pull_message "You can use template language for various info, please consult the documentation for more details."
    }
    trans_contributoragreement {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey component_id FK
        ForeignKey project_id FK
        ForeignKey category_id FK
        ForeignKey workspace_id FK
        DateTimeField timestamp
    }
    trans_alert {
        AutoField id PK
        ForeignKey component_id FK
        DateTimeField timestamp
        DateTimeField updated
        CharField name
        DateTimeField dismissed_at
        ForeignKey dismissed_by_id FK
        CharField dismissal_reason
        CharField dismissal_fingerprint
        AlertSeverity severity
        JSONField details
    }
    trans_announcement {
        AutoField id PK
        TextField message "You can use Markdown and mention users by @username."
        ForeignKey project_id FK
        ForeignKey category_id FK
        ForeignKey component_id FK
        ForeignKey language_id FK
        AnnouncementSeverity severity "Severity defines color used for the message."
        DateField expiry "The message will be not shown after this date. Use it to announce string freeze and translation deadline for next release."
        BooleanField notify "Send notification to subscribed users."
    }
    trans_backupauthorclaim {
        AutoField id PK
        ForeignKey project_id FK
        ForeignKey user_id FK
        OneToOneField invitation_id FK, UK
        DateTimeField accepted_at
    }
    trans_backupauthorvote {
        AutoField id PK
        ForeignKey claim_id FK
        ForeignKey suggestion_id FK
        ForeignKey placeholder_id FK
        SmallIntegerField value
    }
    trans_category {
        AutoField id PK
        CharField name "Display name"
        SlugField slug "Name used in URLs and filenames."
        ForeignKey project_id FK
        ForeignKey category_id FK
        TextField check_flags "Additional comma-separated flags to influence Weblate behavior."
        JSONField enforced_checks "List of checks which can not be dismissed."
        BooleanField inherit_enforced_checks "Use enforced checks from the parent category, project or workspace."
        CategoryLicense license
        BooleanField inherit_license "Use the translation license configured in the parent category or project."
        TextField agreement "Contributor license agreement which needs to be approved before a user can translate components in this category."
        BooleanField inherit_agreement "Use the contributor license agreement configured in the parent category or project."
        CategoryNewLang new_lang "How to handle requests for creating new translations."
        BooleanField inherit_new_lang "Use the adding new translations setting configured in the parent category or project."
        CategoryLanguageCodeStyle language_code_style "Customize language code used to generate the filename for translations created by Weblate."
        BooleanField inherit_language_code_style "Use the language code style configured in the parent category or project."
        ForeignKey secondary_language_id FK "Additional language to show together with the source language while translating."
        BooleanField inherit_secondary_language "Use the secondary language configured in the parent category or project."
        TextField commit_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_commit_message "Use the commit message when translating configured in the parent category or project."
        TextField add_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_add_message "Use the commit message when adding translation configured in the parent category or project."
        TextField delete_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_delete_message "Use the commit message when removing translation configured in the parent category or project."
        TextField merge_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_merge_message "Use the commit message when merging translation configured in the parent category or project."
        TextField addon_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_addon_message "Use the commit message when add-on makes a change configured in the parent category or project."
        TextField pull_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_pull_message "Use the merge request message configured in the parent category or project."
    }
    trans_project {
        AutoField id PK
        CharField name UK "Display name"
        SlugField slug UK "Name used in URLs and filenames."
        URLField web "Main website of translated project."
        TextField instructions "You can use Markdown and mention users by @username."
        ForeignKey workspace_id FK "Workspace this project belongs to. Standalone projects do not need one."
        BooleanField use_shared_tm "Uses the pool of shared translations between projects."
        BooleanField contribute_shared_tm "Publishes source strings, translations, and project and component origin to the global shared translation memory. Project access control does not restrict this shared data."
        BooleanField use_workspace_tm "Uses the pool of shared translations between projects in the workspace."
        BooleanField contribute_workspace_tm "Contributes translations to the pool shared between projects in the workspace."
        BooleanField autoclean_tm "Automatically removes outdated and obsolete entries from translation memory."
        ProjectAccessControl access_control "How to restrict access to this project is detailed in the documentation."
        BooleanField public_sharing "Allows anonymous access to the engage pages and status widgets for Private and Custom projects. Public and Protected projects are always publicly shared regardless of this setting."
        BooleanField enforced_2fa "Requires contributors to have two-factor authentication configured before being able to contribute."
        BooleanField translation_review "Requires dedicated reviewers to approve translations."
        BooleanField source_review "Requires dedicated reviewers to approve source strings."
        CommitPolicyChoices commit_policy "Select which translations should be included when committing changes. More restrictive options will skip translations with potential quality issues."
        BooleanField enable_hooks "Whether to allow updating this repository by remote hooks."
        TextField language_aliases "Comma-separated list of language code mappings, for example: en_GB:en,en_US:en"
        ForeignKey secondary_language_id FK "Additional language to show together with the source language while translating."
        TextField check_flags "Additional comma-separated flags to influence Weblate behavior."
        JSONField enforced_checks "List of checks which can not be dismissed."
        BooleanField inherit_enforced_checks "Use enforced checks from the workspace."
        ProjectLicense license
        BooleanField inherit_license "Use the translation license configured in the workspace."
        TextField agreement "Contributor license agreement which needs to be approved before a user can translate components in this project."
        BooleanField inherit_agreement "Use the contributor license agreement configured in the workspace."
        ProjectNewLang new_lang "How to handle requests for creating new translations."
        BooleanField inherit_new_lang "Use the adding new translations setting configured in the workspace."
        ProjectLanguageCodeStyle language_code_style "Customize language code used to generate the filename for translations created by Weblate."
        BooleanField inherit_language_code_style "Use the language code style configured in the workspace."
        BooleanField inherit_secondary_language "Use the secondary language configured in the workspace."
        TextField commit_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_commit_message "Use the commit message when translating configured in the workspace."
        TextField add_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_add_message "Use the commit message when adding translation configured in the workspace."
        TextField delete_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_delete_message "Use the commit message when removing translation configured in the workspace."
        TextField merge_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_merge_message "Use the commit message when merging translation configured in the workspace."
        TextField addon_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_addon_message "Use the commit message when add-on makes a change configured in the workspace."
        TextField pull_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_pull_message "Use the merge request message configured in the workspace."
    }
    trans_change {
        AutoField id PK
        ForeignKey backup_author_claim_id FK
        ForeignKey unit_id FK
        ForeignKey language_id FK
        ForeignKey project_id FK
        ForeignKey workspace_id FK
        ForeignKey category_id FK
        ForeignKey component_id FK
        ForeignKey translation_id FK
        ForeignKey comment_id FK
        ForeignKey suggestion_id FK
        ForeignKey announcement_id FK
        ForeignKey screenshot_id FK
        ForeignKey alert_id FK
        ForeignKey addon_id FK
        ForeignKey user_id FK
        ForeignKey author_id FK
        DateTimeField timestamp
        ActionEvents action
        TextField target
        TextField old
        JSONField details
    }
    trans_pendingunitchange {
        AutoField id PK
        ForeignKey backup_author_claim_id FK
        ForeignKey unit_id FK
        ForeignKey author_id FK
        TextField target
        TextField explanation
        TextField source_unit_explanation
        StringState state
        DateTimeField timestamp
        BooleanField add_unit
        BooleanField automatically_translated
        JSONField metadata
    }
    trans_comment {
        AutoField id PK
        ForeignKey backup_author_claim_id FK
        ForeignKey unit_id FK
        TextField comment
        ForeignKey user_id FK
        DateTimeField timestamp
        BooleanField resolved
        JSONField userdetails
    }
    trans_suggestion {
        AutoField id PK
        ForeignKey backup_author_claim_id FK
        ForeignKey unit_id FK
        TextField target
        ForeignKey user_id FK
        JSONField userdetails
        DateTimeField timestamp
    }
    trans_vote {
        AutoField id PK
        ForeignKey suggestion_id FK
        ForeignKey user_id FK
        SmallIntegerField value
    }
    trans_variant {
        AutoField id PK
        ForeignKey component_id FK
        RegexField variant_regex
        TextField key
    }
    trans_unit {
        AutoField id PK
        ForeignKey translation_id FK
        BigIntegerField id_hash
        TextField location
        TextField context
        TextField note
        TextField flags
        TextField source
        TextField previous_source
        TextField target
        StringState state
        StringState original_state
        JSONField details
        IntegerField position
        IntegerField num_words
        IntegerField priority
        DateTimeField timestamp
        DateTimeField last_updated
        TextField extra_flags "Additional comma-separated flags to influence Weblate behavior."
        TextField explanation "Additional explanation to clarify meaning or usage of the string."
        ForeignKey variant_id FK
        BooleanField automatically_translated "Indicates whether this string was translated automatically."
        ForeignKey source_unit_id FK
        ForeignKey translation_parent_id FK
    }
    trans_translation {
        AutoField id PK
        ForeignKey component_id FK
        ForeignKey language_id FK
        ForeignKey plural_id FK
        CharField revision
        CharField filename
        CharField language_code
        TextField check_flags
    }
    trans_component_links {
        AutoField id PK
        ForeignKey component_id FK
        ForeignKey project_id FK
        ForeignKey category_id FK
    }
    trans_component {
        AutoField id PK
        CharField name "Display name"
        SlugField slug "Name used in URLs and filenames."
        ForeignKey project_id FK
        ForeignKey category_id FK
        ComponentVcs vcs "Version control system to use to access your repository containing translations. You can also choose additional integration with third party providers to submit pull/merge requests."
        JSONField vcs_params
        CharField repo "URL of a repository, use weblate://project/component to share it with other component."
        ForeignKey linked_component_id FK
        CharField push "URL of a push repository, pushing is turned off if empty. When using pull/merge requests, keep empty to fork the repository."
        CharField repoweb "Link to repository browser, use {{branch}} for branch, {{filename}} and {{line}} as filename and line placeholders. You might want to strip leading directory by using {{filename|parentdir}}."
        CharField repoweb_translations "Link to repository browser for translation files, use {{branch}} for branch, {{filename}} and {{line}} as filename and line placeholders. If left empty, the Repository browser above will be used. You might want to strip leading directory by using {{filename|parentdir}}."
        CharField git_export "URL of repository where users can fetch changes from Weblate"
        EmailField report_source_bugs "E-mail address for reports on errors in source strings. Leave empty for no e-mails."
        CharField branch "Repository branch to translate"
        CharField push_branch "Branch for pushing changes. Leave empty to use repository branch when pushing directly to the repository. When using pull/merge requests, specify a branch name different from the repository branch."
        CharField filemask "Path of files to translate relative to repository root, use * instead of language code, for example: po/*.po or locale/*/LC_MESSAGES/django.po."
        CharField screenshot_filemask "Path of screenshots relative to repository root, for example: docs/screenshots/*.png."
        CharField template "Filename of translation base file, containing all strings and their source for monolingual translations."
        BooleanField edit_template "Whether users will be able to edit the base file for monolingual translations."
        CharField intermediate "Filename of intermediate translation file. In most cases this is a translation file provided by developers and is used when creating actual source strings."
        CharField new_base "Filename of file used for creating new translations. For gettext choose .pot file."
        ComponentFileFormat file_format
        JSONField file_format_params
        BooleanField locked "Locked component will not get any translation updates."
        BooleanField hide_glossary_matches "Hides the glossary panel in the translation editor."
        BooleanField allow_translation_propagation "Whether translation updates in other components will cause automatic translation in this one"
        BooleanField contribute_project_tm "Allow this component's translations to be added to the project-level translation memory."
        BooleanField enable_suggestions "Whether to allow translation suggestions at all."
        BooleanField suggestion_voting "Allows users to vote on suggestions."
        PositiveSmallIntegerField suggestion_autoaccept "Automatically accept suggestions with this number of votes, use 0 to turn it off."
        TextField check_flags "Additional comma-separated flags to influence Weblate behavior."
        JSONField enforced_checks "List of checks which can not be dismissed."
        BooleanField inherit_enforced_checks "Use enforced checks from the project, category or workspace."
        ComponentLicense license
        BooleanField inherit_license "Use the translation license configured in the category or project."
        TextField agreement "Contributor license agreement which needs to be approved before a user can translate this component."
        BooleanField inherit_agreement "Use the contributor license agreement configured in the category or project."
        ComponentNewLang new_lang "How to handle requests for creating new translations."
        BooleanField inherit_new_lang "Use the adding new translations setting configured in the category or project."
        ComponentLanguageCodeStyle language_code_style "Customize language code used to generate the filename for translations created by Weblate."
        BooleanField inherit_language_code_style "Use the language code style configured in the category or project."
        BooleanField manage_units "Enables adding, removing, and editing source strings and keys in Weblate. If your strings are extracted from the source code or managed externally you probably want to keep it disabled."
        ComponentMergeStyle merge_style "Define whether Weblate should merge the upstream repository or rebase changes onto it."
        TextField commit_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_commit_message "Use the commit message when translating configured in the category or project."
        TextField add_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_add_message "Use the commit message when adding translation configured in the category or project."
        TextField delete_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_delete_message "Use the commit message when removing translation configured in the category or project."
        TextField merge_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_merge_message "Use the commit message when merging translation configured in the category or project."
        TextField addon_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_addon_message "Use the commit message when add-on makes a change configured in the category or project."
        TextField pull_message "You can use template language for various info, please consult the documentation for more details."
        BooleanField inherit_pull_message "Use the merge request message configured in the category or project."
        BooleanField push_on_commit "Whether the repository should be pushed upstream on every commit."
        BooleanField push_on_update "Whether the repository should be pushed upstream after updating it, even when the update did not commit any translations. When turned off, commits made by the update are pushed with the next translation commit."
        SmallIntegerField commit_pending_age "Time in hours after which any pending changes will be committed to the VCS."
        BooleanField auto_lock_error "Whether the component should be locked on repository errors."
        ForeignKey source_language_id FK "Language used for source strings in all components"
        RegexField language_regex "Regular expression used to filter translation files when scanning for file mask."
        RegexField variant_regex "Regular expression used to determine variants of a string."
        ComponentPriority priority "Components with higher priority are offered first to translators."
        BooleanField restricted "Restrict access to the component to only those explicitly given permission."
        BooleanField is_glossary "Glossaries are different from regular translations but help keep track of and share consistent terminology."
        ColorChoices glossary_color
        CharField remote_revision
        CharField local_revision
        CharField processed_revision
        URLField pull_request_url
        RegexField key_filter "Regular expression used to filter keys. This is only available for monolingual formats."
        ForeignKey secondary_language_id FK "Additional language to show together with the source language while translating. This setting is inherited from the project if left empty."
        BooleanField inherit_secondary_language "Use the secondary language configured in the category or project."
    }
    trans_componentlist {
        AutoField id PK
        CharField name UK "Display name"
        SlugField slug UK "Name used in URLs and filenames."
        BooleanField show_dashboard "When enabled this component list will be shown as a tab on the dashboard"
    }
    trans_autocomponentlist {
        AutoField id PK
        RegexField project_match "Regular expression which is used to match project slug."
        RegexField component_match "Regular expression which is used to match component slug."
        ForeignKey componentlist_id FK
    }
    trans_label {
        AutoField id PK
        ForeignKey project_id FK
        CharField name
        ColorChoices color
        CharField description
    }
    trans_report {
        AutoField id PK
        ForeignKey creator_id FK
        DateTimeField created
        ReportKind kind
        PositiveSmallIntegerField version
        JSONField parameters
        JSONField data
        ForeignKey workspace_id FK
        ForeignKey project_id FK
        ForeignKey category_id FK
        ForeignKey component_id FK
    }
    trans_savedsearch {
        AutoField id PK
        ForeignKey project_id FK
        ForeignKey user_id FK
        CharField name
        TextField query
        CharField description
    }
    trans_workflowsetting {
        AutoField id PK
        ForeignKey project_id FK
        ForeignKey language_id FK
        ForeignKey source_language_id FK "Translate from this language. Leave empty to use the component source language."
        BooleanField translation_review "Requires dedicated reviewers to approve translations."
        BooleanField enable_suggestions "Whether to allow translation suggestions at all."
        BooleanField restrict_direct_editing "Only users with the “Edit string when suggestions are enforced” permission can make direct changes."
        BooleanField suggestion_voting "Allows users to vote on suggestions."
        PositiveSmallIntegerField suggestion_autoaccept "Automatically accept suggestions with this number of votes, use 0 to turn it off."
    }
    lang_language {
        AutoField id PK
        CharField code UK
        CharField name
        LanguageDirection direction
        BigIntegerField population "Number of people speaking this language."
    }
    lang_plural {
        AutoField id PK
        PluralSource source
        SmallIntegerField number
        TextField formula
        PluralType type
        ForeignKey language_id FK
    }
    memory_memory {
        AutoField id PK
        ForeignKey source_language_id FK
        ForeignKey target_language_id FK
        TextField source
        TextField target
        TextField origin
        TextField context
        ForeignKey legacy_user_id FK
        ForeignKey legacy_project_id FK
        BooleanField legacy_from_file
        BooleanField legacy_shared
        MemoryStatus status
    }
    memory_memoryscope {
        AutoField id PK
        ForeignKey memory_id FK
        MemoryScopeChoices scope
        ForeignKey project_id FK
        ForeignKey workspace_id FK
        ForeignKey source_project_id FK
        ForeignKey source_component_id FK
        ForeignKey user_id FK
    }
    memory_memoryscopemigrationstate {
        CharField name PK
        IntegerField last_memory_id
        BooleanField completed
        DateTimeField updated
    }
    screenshots_screenshot {
        AutoField id PK
        ForeignKey backup_author_claim_id FK
        CharField name
        CharField repository_filename "Scan for screenshot file change on repository update."
        ScreenshotField image "Upload image up to 2000x2000 pixels."
        ForeignKey translation_id FK
        DateTimeField timestamp
        ForeignKey user_id FK
    }
    screenshots_screenshotocr {
        OneToOneField screenshot_id PK
        JSONField signature
        UUIDField generation
        JSONField data
        UUIDField task_id
        DateTimeField requested_at
        CharField error
    }
    fonts_font {
        AutoField id PK
        CharField family
        CharField style
        FileField font "OpenType and TrueType fonts are supported."
        ForeignKey project_id FK
        DateTimeField timestamp
        ForeignKey user_id FK
    }
    fonts_fontgroup {
        AutoField id PK
        SlugField name "Identifier you will use in checks to select this font group. Avoid whitespaces and special characters."
        ForeignKey font_id FK "Default font is used unless per language override matches."
        ForeignKey project_id FK
    }
    fonts_fontoverride {
        AutoField id PK
        ForeignKey group_id FK
        ForeignKey font_id FK
        ForeignKey language_id FK
    }
    accounts_subscription {
        AutoField id PK
        ForeignKey user_id FK
        SubscriptionNotification notification
        NotificationScope scope
        NotificationFrequency frequency
        ForeignKey project_id FK
        ForeignKey component_id FK
        BooleanField onetime
    }
    accounts_teamauditlog {
        AutoField id PK
        ForeignKey team_id FK
        CharField team_name
        TeamAuditLogScope scope
        CharField scope_label
        ForeignKey project_id FK
        ForeignKey workspace_id FK
        ForeignKey actor_id FK
        TeamAuditLogEvent event
        JSONField details
        DateTimeField timestamp
    }
    accounts_auditlog {
        AutoField id PK
        ForeignKey user_id FK
        AuditLogActivity activity
        JSONField params
        GenericIPAddressField address
        CharField user_agent
        DateTimeField timestamp
    }
    accounts_verifiedemail {
        AutoField id PK
        BooleanField is_deliverable
        ForeignKey social_id FK
        EmailField email
    }
    accounts_profile {
        AutoField id PK
        OneToOneField user_id FK, UK
        ProfileLanguage language
        IntegerField suggested
        IntegerField translated
        IntegerField uploaded
        IntegerField commented
        ThemeChoices theme
        BooleanField hide_completed
        BooleanField secondary_in_zen
        BooleanField hide_source_secondary
        BooleanField wide_tables "Instead of hiding columns on narrow screens, keep all columns and scroll the table horizontally."
        JSONField listing_columns "Choose which statistics columns are shown in project, component, and language lists."
        CharField editor_link "Enter a custom URL to be used as link to the source code. You can use {{branch}} for branch, {{filename}} and {{line}} as filename and line placeholders."
        ProfileTranslateMode translate_mode
        ProfileZenMode zen_mode
        CharField special_chars "You can specify additional special visual keyboard characters to be shown while translating. It can be useful for characters you use frequently, but are hard to type on your keyboard."
        SmallIntegerField nearby_strings "Number of nearby strings to show in each direction in the full editor."
        BooleanField auto_watch "Projects are added to your watched projects when you contribute to them."
        BooleanField contribute_personal_tm "Allow your translations to be added to your personal translation memory."
        ProfileDashboardView dashboard_view
        ForeignKey dashboard_component_list_id FK
        URLField website
        URLField contact "Link to contact you online using services like Signal, SimpleX or Telegram."
        SlugField liberapay "Liberapay is a platform to donate money to teams, organizations and individuals."
        URLField fediverse "Link to your Fediverse profile for federated services like Mastodon or diaspora*."
        URLField codesite "Link to your code profile for services like Codeberg or GitLab."
        SlugField github
        SlugField twitter
        SlugField linkedin "Your LinkedIn profile name from linkedin.com/in/profilename"
        CharField location
        CharField company
        EmailField public_email
        EmailField commit_email
        ProfileCommitNameChoices commit_name
        ProfileLast2Fa last_2fa
    }
    configuration_setting {
        AutoField id PK
        SettingCategory category
        CharField name
        JSONField value
    }
    vcs_installation {
        AutoField id PK
        InstallationProvider provider
        CharField installation_id
        CharField target_type "Organization or User"
        CharField target_login "Hosting organization or user login"
        ForeignKey workspace_id FK
        CharField hostname
        BooleanField enabled
        DateTimeField created
        JSONField repositories
        DateTimeField repositories_updated
    }
    vcs_pendinginstallation {
        AutoField id PK
        InstallationProvider provider
        CharField hostname
        CharField installation_id
        JSONField payload
        DateTimeField created
        DateTimeField updated
    }
    vcs_githubappcredentials {
        AutoField id PK
        CharField hostname UK
        CharField app_id
        CharField app_slug
        TextField private_key
        CharField client_id
        CharField client_secret
        CharField webhook_secret
        UUIDField webhook_token UK
        URLField html_url
        DateTimeField created
        DateTimeField updated
    }
    wladmin_securityscan {
        UUIDField id PK
        ForeignKey actor_id FK
        CharField filename
        CharField operation
        CharField finding_type
        CharField state
        DateTimeField created
        DateTimeField completed
        PositiveIntegerField processed
        PositiveIntegerField total
        JSONField summary
        TextField error
    }
    wladmin_securityfinding {
        AutoField id PK
        ForeignKey user_id FK
        CharField finding_type
        CharField key
        JSONField evidence
        SecurityFindingStatus status
        DateTimeField first_seen
        DateTimeField last_seen
        DateTimeField handled_at
    }
    wladmin_securityfindingevent {
        AutoField id PK
        ForeignKey finding_id FK
        ForeignKey scan_id FK
        ForeignKey actor_id FK
        CharField action
        DateTimeField timestamp
        CharField reason
        JSONField evidence
    }
    wladmin_configurationerror {
        AutoField id PK
        CharField name UK
        TextField message
        DateTimeField timestamp
        BooleanField ignored
    }
    wladmin_supportstatus {
        AutoField id PK
        CharField name
        CharField secret
        DateTimeField expiry
        BooleanField in_limits
        BooleanField discoverable
        JSONField limits
        BooleanField has_subscription
        CharField backup_repository
        BooleanField enabled
    }
    wladmin_backupservice {
        AutoField id PK
        CharField repository "Use /path/to/repo for local backups or user@host:/path/to/repo or ssh://user@host:port/path/to/backups for remote SSH backups."
        BooleanField enabled
        DateTimeField timestamp
        CharField passphrase
        TextField paperkey
    }
    wladmin_backuplog {
        AutoField id PK
        ForeignKey service_id FK
        DateTimeField timestamp
        BackupLogEvent event
        BooleanField warning
        TextField log
    }
    metrics_metric {
        BigAutoField id PK
        DateField date
        SmallIntegerField scope
        IntegerField relation
        IntegerField secondary
        IntegerField changes
        JSONField data
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
    social_auth_usersocialauth {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField provider
        CharField id_key
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
    authtoken_token {
        CharField key PK
        OneToOneField user_id FK, UK
        DateTimeField created
    }
    django_celery_beat_solarschedule {
        AutoField id PK
        SolarScheduleEvent event "The type of solar event when the job should run"
        DecimalField latitude "Run the task when the event happens at this latitude"
        DecimalField longitude "Run the task when the event happens at this longitude"
    }
    django_celery_beat_intervalschedule {
        AutoField id PK
        IntegerField every "Number of interval periods to wait before running the task again"
        IntervalSchedulePeriod period "The type of period between task runs (Example: days)"
    }
    django_celery_beat_clockedschedule {
        AutoField id PK
        DateTimeField clocked_time "Run the task at clocked time"
    }
    django_celery_beat_crontabschedule {
        AutoField id PK
        CharField minute "Cron Minutes to Run. Use '*' for 'all'. (Example: '0,30')"
        CharField hour "Cron Hours to Run. Use '*' for 'all'. (Example: '8,20')"
        CharField day_of_month "Cron Days Of The Month to Run. Use '*' for 'all'. (Example: '1,15')"
        CharField month_of_year "Cron Months (1-12) Of The Year to Run. Use '*' for 'all'. (Example: '1,12')"
        CharField day_of_week "Cron Days Of The Week to Run. Use '*' for 'all', Sunday is 0 or 7, Monday is 1. (Example: '0,5')"
        CrontabScheduleTimezone timezone "Timezone to Run the Cron Schedule on. Default is UTC."
    }
    django_celery_beat_periodictasks {
        SmallIntegerField ident PK
        DateTimeField last_update
    }
    django_celery_beat_periodictask {
        AutoField id PK
        CharField name UK "Short Description For This Task"
        CharField task "The Name of the Celery Task that Should be Run.  (Example: 'proj.tasks.import_contacts')"
        ForeignKey interval_id FK "Interval Schedule to run the task on.  Set only one schedule type, leave the others null."
        ForeignKey crontab_id FK "Crontab Schedule to run the task on.  Set only one schedule type, leave the others null."
        ForeignKey solar_id FK "Solar Schedule to run the task on.  Set only one schedule type, leave the others null."
        ForeignKey clocked_id FK "Clocked Schedule to run the task on.  Set only one schedule type, leave the others null."
        TextField args "JSON encoded positional arguments (Example: ['arg1', 'arg2'])"
        TextField kwargs "JSON encoded keyword arguments (Example: {'argument': 'value'})"
        CharField queue "Queue defined in CELERY_TASK_QUEUES. Leave None for default queuing."
        CharField exchange "Override Exchange for low-level AMQP routing"
        CharField routing_key "Override Routing Key for low-level AMQP routing"
        TextField headers "JSON encoded message headers for the AMQP message."
        PositiveIntegerField priority "Priority Number between 0 and 255. Supported by: RabbitMQ, Redis (priority reversed, 0 is highest)."
        DateTimeField expires "Datetime after which the schedule will no longer trigger the task to run"
        PositiveIntegerField expire_seconds "Timedelta with seconds which the schedule will no longer trigger the task to run"
        BooleanField one_off "If True, the schedule will only run the task a single time"
        DateTimeField start_time "Datetime when the schedule should begin triggering the task to run"
        BooleanField enabled "Set to False to disable the schedule"
        DateTimeField last_run_at "Datetime that the schedule last triggered the task to run. Reset to None if enabled is set to False."
        PositiveIntegerField total_run_count "Running count of how many times the schedule has triggered the task"
        DateTimeField date_changed "Datetime that this PeriodicTask was last modified"
        TextField description "Detailed description about the details of this Periodic Task"
    }
    otp_static_staticdevice {
        AutoField id PK
        ForeignKey user_id FK "The user that this device belongs to."
        CharField name "The human-readable name of this device."
        BooleanField confirmed "Is this device ready for use?"
        DateTimeField throttling_failure_timestamp "A timestamp of the last failed verification attempt. Null if last attempt succeeded."
        PositiveIntegerField throttling_failure_count "Number of successive failed attempts."
        DateTimeField created_at "The date and time when this device was initially created in the system."
        DateTimeField last_used_at "The most recent date and time this device was used."
    }
    otp_static_statictoken {
        AutoField id PK
        ForeignKey device_id FK
        CharField token
    }
    otp_totp_totpdevice {
        AutoField id PK
        ForeignKey user_id FK "The user that this device belongs to."
        CharField name "The human-readable name of this device."
        BooleanField confirmed "Is this device ready for use?"
        DateTimeField throttling_failure_timestamp "A timestamp of the last failed verification attempt. Null if last attempt succeeded."
        PositiveIntegerField throttling_failure_count "Number of successive failed attempts."
        DateTimeField created_at "The date and time when this device was initially created in the system."
        DateTimeField last_used_at "The most recent date and time this device was used."
        CharField key "A hex-encoded secret key of up to 40 bytes."
        PositiveSmallIntegerField step "The time step in seconds."
        BigIntegerField t0 "The Unix time at which to begin counting steps."
        TOTPDeviceDigits digits "The number of digits to expect in a token."
        PositiveSmallIntegerField tolerance "The number of time steps in the past or future to allow."
        SmallIntegerField drift "The number of time steps the prover is known to deviate from our clock."
        BigIntegerField last_t "The t value of the latest verified token. The next token must be at a higher time step."
    }
    django_otp_webauthn_webauthncredential {
        BigAutoField id PK
        ForeignKey user_id FK "The user that this device belongs to."
        CharField name "The human-readable name of this device."
        BooleanField confirmed "Is this device ready for use?"
        DateTimeField created_at "The date and time when this device was initially created in the system."
        DateTimeField last_used_at "The most recent date and time this device was used."
        WebAuthnCredentialCredentialType credential_type
        BinaryField credential_id
        BinaryField public_key
        JSONField transports
        PositiveIntegerField sign_count
        BooleanField backup_eligible
        BooleanField backup_state
        CharField aaguid
        CharField credential_id_sha256 UK
        BooleanField discoverable
    }
    django_otp_webauthn_webauthnattestation {
        BigAutoField id PK
        WebAuthnAttestationFmt fmt
        BinaryField data
        BinaryField client_data_json
        OneToOneField credential_id FK, UK
    }
    django_otp_webauthn_webauthnuserhandle {
        CharField handle_hex UK
        OneToOneField user_id PK
    }
```
