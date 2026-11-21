```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    django_celery_beat_intervalschedule |o--o{ django_celery_beat_periodictask : "(id) - (interval_id)"
    django_celery_beat_crontabschedule |o--o{ django_celery_beat_periodictask : "(id) - (crontab_id)"
    django_celery_beat_solarschedule |o--o{ django_celery_beat_periodictask : "(id) - (solar_id)"
    django_celery_beat_clockedschedule |o--o{ django_celery_beat_periodictask : "(id) - (clocked_id)"
    auth_user |o--o{ esi_token : "(id) - (user_id)"
    esi_token }o--o{ esi_scope : "(scopes) - (token_set)"
    esi_token |o--o{ esi_callbackredirect : "(id) - (token_id)"
    authentication_state }o--o{ eveonline_evecharacter : "(member_characters) - (state_set)"
    authentication_state }o--o{ eveonline_evecorporationinfo : "(member_corporations) - (state_set)"
    authentication_state }o--o{ eveonline_eveallianceinfo : "(member_alliances) - (state_set)"
    authentication_state }o--o{ eveonline_evefactioninfo : "(member_factions) - (state_set)"
    eveonline_evecharacter |o--o| authentication_userprofile : "(id) - (main_character_id)"
    authentication_state ||--o{ authentication_userprofile : "(id) - (state_id)"
    eveonline_evecharacter ||--o| authentication_characterownership : "(id) - (character_id)"
    eveonline_evecharacter ||--o{ authentication_ownershiprecord : "(id) - (character_id)"
    services_nameformatconfig }o--o{ authentication_state : "(states) - (nameformatconfig_set)"
    eveonline_evefactioninfo |o--o{ eveonline_eveallianceinfo : "(id) - (faction_id)"
    eveonline_eveallianceinfo |o--o{ eveonline_evecorporationinfo : "(id) - (alliance_id)"
    eveonline_evefactioninfo |o--o{ eveonline_evecorporationinfo : "(id) - (faction_id)"
    groupmanagement_authgroup }o--o{ authentication_state : "(states) - (valid_states)"
    menu_menuitem |o--o{ menu_menuitem : "(id) - (parent_id)"
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
    esi_scope {
        AutoField id PK
        CharField name UK "The official EVE name for the scope."
        TextField help_text "The official EVE description of the scope."
    }
    esi_token {
        AutoField id PK
        DateTimeField created
        TextField access_token "The access token granted by SSO."
        TextField refresh_token "A re-usable token to generate new access tokens upon expiry."
        ForeignKey user_id FK "The user to whom this token belongs."
        IntegerField character_id "The ID of the EVE character who authenticated by SSO."
        CharField character_name "The name of the EVE character who authenticated by SSO."
        TokenTokenType token_type "The applicable range of the token."
        CharField character_owner_hash "The unique string identifying this character and its owning EVE account. Changes if the owning account changes."
        IntegerField sso_version "EVE SSO Version."
    }
    esi_callbackredirect {
        AutoField id PK
        CharField session_key UK "Session key identifying the session this redirect was created for."
        TextField url "The internal URL to redirect this callback towards."
        CharField state "OAuth2 state string representing this session."
        DateTimeField created
        ForeignKey token_id FK "Token generated by a completed code exchange from callback processing."
    }
    admin_status_applicationannouncement {
        AutoField id PK
        CharField application_name "Name of the application that issued the announcement"
        IntegerField announcement_number "Issue number on the notification source"
        TextField announcement_text "Issue title text displayed on the dashboard"
        TextField announcement_url
        CharField announcement_hash UK "hash of an announcement."
        BooleanField hide_announcement "Set to true if the announcement should not be displayed on the dashboard"
    }
    admin_status_softwareversion {
        AutoField id PK
        CharField latest_stable_version "Latest stable version of Alliance Auth"
        CharField latest_development_version "Latest development version of Alliance Auth"
    }
    authentication_state {
        AutoField id PK
        CharField name UK
        IntegerField priority UK "Users get assigned the state with the highest priority available to them."
        BooleanField public "Make this state available to any character."
    }
    authentication_userprofile {
        AutoField id PK
        OneToOneField user_id FK, UK
        OneToOneField main_character_id FK, UK
        ForeignKey state_id FK
        UserProfileLanguage language
        BooleanField night_mode
        CharField theme "Bootstrap 5 Themes from https://bootswatch.com/ or Community Apps"
        BooleanField minimize_sidebar "Keep the sidebar menu minimized"
    }
    authentication_characterownership {
        AutoField id PK
        OneToOneField character_id FK, UK
        CharField owner_hash UK
        ForeignKey user_id FK
    }
    authentication_ownershiprecord {
        AutoField id PK
        ForeignKey character_id FK
        CharField owner_hash
        ForeignKey user_id FK
        DateTimeField created
    }
    services_nameformatconfig {
        AutoField id PK
        CharField service_name
        BooleanField default_to_username "If a user has no main_character, default to using their Auth username instead."
        CharField format "For information on constructing name formats please see the official documentation, topic 'Services Name Formats'."
    }
    eveonline_evefactioninfo {
        AutoField id PK
        PositiveIntegerField faction_id UK
        PositiveIntegerField corporation_id UK
        TextField description
        PositiveIntegerField militia_corporation_id UK
        CharField faction_name UK
        PositiveSmallIntegerField size_factor
        PositiveIntegerField solar_system_id
        PositiveIntegerField station_count
        PositiveIntegerField station_system_count
        DateTimeField last_updated "Last time the faction's details were updated, (GetUniverseFactions), 24 hr Cache"
    }
    eveonline_eveallianceinfo {
        AutoField id PK
        PositiveIntegerField alliance_id UK
        PositiveIntegerField creator_corporation_id "Alliance's creator corporation ID"
        PositiveIntegerField creator_id "Alliance's creator ID"
        DateField date_founded "Alliance's founding date"
        PositiveIntegerField executor_corp_id "Alliance's executor corporation ID"
        ForeignKey faction_id FK "Alliance's faction ID"
        CharField alliance_name "Alliance's name"
        CharField alliance_ticker "Alliance's ticker"
        DateTimeField last_updated "Last time the alliance's details were updated, (GetAlliancesAllianceId), 1 hr Cache"
    }
    eveonline_evecorporationinfo {
        AutoField id PK
        PositiveIntegerField corporation_id UK
        ForeignKey alliance_id FK
        PositiveIntegerField ceo_id "Corporation's CEO ID"
        PositiveIntegerField creator_id "Corporation's creator ID"
        DateField date_founded "Corporation's founding date"
        TextField description "Corporation's description"
        ForeignKey faction_id FK
        PositiveIntegerField home_station_id "Corporation's home station ID"
        PositiveIntegerField member_count "Corporation's member count"
        CharField corporation_name "Corporation's name"
        PositiveBigIntegerField shares "Corporation's shares"
        FloatField tax_rate "Corporation's tax rate"
        CharField corporation_ticker "Corporation's short name"
        URLField url "Corporation's URL"
        BooleanField war_eligible "Corporation's war eligibility"
        DateTimeField last_updated "Last time the corporation's details were updated, (GetCorporationsCorporationId), 1 hr Cache"
    }
    eveonline_evecharacter {
        AutoField id PK
        PositiveIntegerField character_id UK
        PositiveIntegerField alliance_id "Character's alliance ID"
        DateField birthday "Character's creation date"
        PositiveSmallIntegerField bloodline_id "Character's bloodline ID"
        PositiveIntegerField corporation_id "Character's corporation ID"
        TextField description "Character's description (biography)"
        PositiveIntegerField faction_id "Character's faction ID"
        EveCharacterGender gender "Character's gender"
        CharField character_name "Character's name"
        PositiveSmallIntegerField race_id "Character's race ID"
        FloatField security_status "Character's security status"
        CharField title "Character's title"
        CharField corporation_name
        CharField corporation_ticker
        CharField alliance_name
        CharField alliance_ticker
        CharField faction_name
        DateTimeField last_updated_affiliations "Last time the character's affiliations were updated (PostCharactersAffiliation, 1 hr Cache)"
        DateTimeField last_updated_other "Last time the character's details were updated, (GetCharactersCharacterId), 24 hr Cache"
    }
    groupmanagement_grouprequest {
        AutoField id PK
        BooleanField leave_request
        ForeignKey user_id FK
        ForeignKey group_id FK
    }
    groupmanagement_requestlog {
        AutoField id PK
        RequestLogType request_type
        ForeignKey group_id FK
        RequestLogAction action
        ForeignKey request_actor_id FK
        ForeignKey requestor_id FK
        DateTimeField date
    }
    groupmanagement_authgroup {
        OneToOneField group_id PK
        BooleanField internal "Internal group, users cannot see, join or request to join this group.<br>Used for groups such as Members, Corp_*, Alliance_* etc.<br><b>Overrides Hidden and Open options when selected.</b>"
        BooleanField hidden "Group is hidden from users but can still join with the correct link."
        BooleanField open "Group is open and users will be automatically added upon request.<br>If the group is not open users will need their request manually approved."
        BooleanField public "Group is public. Any registered user is able to join this group, with visibility based on the other options set for this group.<br>Auth will not remove users from this group automatically when they are no longer authenticated."
        BooleanField restricted "Group is restricted. This means that adding or removing users for this group requires a superuser admin."
        TextField description "Short description <i>(max. 512 characters)</i> of the group shown to users."
    }
    groupmanagement_reservedgroupname {
        AutoField id PK
        CharField name UK "Name that can not be used for groups."
        TextField reason "Reason why this name is reserved."
        CharField created_by "Name of the user who created this entry."
        DateTimeField created_at "Date when this entry was created"
    }
    notifications_notification {
        AutoField id PK
        ForeignKey user_id FK
        NotificationLevel level
        CharField title
        TextField message
        DateTimeField timestamp
        BooleanField viewed
    }
    analytics_analyticsidentifier {
        AutoField id PK
        UUIDField identifier
    }
    analytics_analyticstokens {
        AutoField id PK
        CharField name
        AnalyticsTokensAnalyticsType type
        CharField token
        CharField secret
        BooleanField send_stats
    }
    menu_menuitem {
        AutoField id PK
        CharField text "Text to show on menu"
        IntegerField order "Order of the menu. Lowest First"
        ForeignKey parent_id FK "Folder this item is in (optional)"
        BooleanField is_hidden "Hide this menu item.If this item is a folder all items under it will be hidden too"
        CharField hook_hash UK "hash of a menu item hook. Must be nullable for unique comparison."
        CharField classes "Font Awesome classes to show as icon on menu, e.g. <code>fa-solid fa-house</code>"
        TextField url "External URL this menu items will link to"
        MenuItemPermissionMode permission_mode "Whether a user needs any or all of the assigned permissions to see this menu item."
    }
    custom_css_customcss {
        AutoField id PK
        TextField css "This CSS will be added to the site after the default CSS."
        TextField css_compressed "This field is automatically generated and should not be edited."
        DateTimeField timestamp
    }
    crontab_cronoffset {
        AutoField id PK
        FloatField minute
        FloatField hour
        FloatField day_of_month
        FloatField month_of_year
        FloatField day_of_week
    }
```
