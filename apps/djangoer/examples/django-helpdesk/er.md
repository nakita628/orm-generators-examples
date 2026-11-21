```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user ||--o{ pinax_invitations_joininvitation : "(id) - (from_user_id)"
    auth_user |o--o{ pinax_invitations_joininvitation : "(id) - (to_user_id)"
    account_signupcode ||--o| pinax_invitations_joininvitation : "(id) - (signup_code_id)"
    auth_user ||--o| pinax_invitations_invitationstat : "(id) - (user_id)"
    auth_user ||--o{ pinax_teams_team : "(id) - (creator_id)"
    pinax_teams_simpleteam ||--o{ pinax_teams_simplemembership : "(id) - (team_id)"
    auth_user |o--o{ pinax_teams_simplemembership : "(id) - (user_id)"
    pinax_invitations_joininvitation |o--o{ pinax_teams_simplemembership : "(id) - (invite_id)"
    pinax_teams_team ||--o{ pinax_teams_membership : "(id) - (team_id)"
    auth_user |o--o{ pinax_teams_membership : "(id) - (user_id)"
    pinax_invitations_joininvitation |o--o{ pinax_teams_membership : "(id) - (invite_id)"
    auth_user ||--o| account_account : "(id) - (user_id)"
    auth_user |o--o{ account_signupcode : "(id) - (inviter_id)"
    account_signupcode ||--o{ account_signupcoderesult : "(id) - (signup_code_id)"
    auth_user ||--o{ account_signupcoderesult : "(id) - (user_id)"
    auth_user ||--o{ account_emailaddress : "(id) - (user_id)"
    account_emailaddress ||--o{ account_emailconfirmation : "(id) - (email_address_id)"
    auth_user |o--o{ account_accountdeletion : "(id) - (user_id)"
    auth_user ||--o{ account_passwordhistory : "(id) - (user_id)"
    auth_user ||--o| account_passwordexpiry : "(id) - (user_id)"
    auth_user |o--o{ helpdesk_queue : "(id) - (default_owner_id)"
    helpdesk_queue ||--o{ helpdesk_ticket : "(id) - (queue_id)"
    auth_user |o--o{ helpdesk_ticket : "(id) - (assigned_to_id)"
    helpdesk_kbitem |o--o{ helpdesk_ticket : "(id) - (kbitem_id)"
    helpdesk_ticket |o--o{ helpdesk_ticket : "(id) - (merged_to_id)"
    helpdesk_ticket ||--o{ helpdesk_followup : "(id) - (ticket_id)"
    auth_user |o--o{ helpdesk_followup : "(id) - (user_id)"
    helpdesk_followup ||--o{ helpdesk_ticketchange : "(id) - (followup_id)"
    helpdesk_followup ||--o{ helpdesk_followupattachment : "(id) - (followup_id)"
    helpdesk_kbitem ||--o{ helpdesk_kbiattachment : "(id) - (kbitem_id)"
    helpdesk_presetreply }o--o{ helpdesk_queue : "(queues) - (presetreply_set)"
    helpdesk_escalationexclusion }o--o{ helpdesk_queue : "(queues) - (escalationexclusion_set)"
    helpdesk_queue |o--o{ helpdesk_kbcategory : "(id) - (queue_id)"
    helpdesk_kbcategory ||--o{ helpdesk_kbitem : "(id) - (category_id)"
    pinax_teams_team |o--o{ helpdesk_kbitem : "(id) - (team_id)"
    helpdesk_kbitem }o--o{ auth_user : "(voted_by) - (votes)"
    helpdesk_kbitem }o--o{ auth_user : "(downvoted_by) - (downvotes)"
    auth_user ||--o{ helpdesk_savedsearch : "(id) - (user_id)"
    auth_user ||--o| helpdesk_usersettings : "(id) - (user_id)"
    helpdesk_ignoreemail }o--o{ helpdesk_queue : "(queues) - (ignoreemail_set)"
    helpdesk_ticket ||--o{ helpdesk_ticketcc : "(id) - (ticket_id)"
    auth_user |o--o{ helpdesk_ticketcc : "(id) - (user_id)"
    helpdesk_ticket ||--o{ helpdesk_ticketcustomfieldvalue : "(id) - (ticket_id)"
    helpdesk_customfield ||--o{ helpdesk_ticketcustomfieldvalue : "(id) - (field_id)"
    helpdesk_ticket ||--o{ helpdesk_ticketdependency : "(id) - (ticket_id)"
    helpdesk_ticket ||--o{ helpdesk_ticketdependency : "(id) - (depends_on_id)"
    helpdesk_ticket ||--o{ helpdesk_checklist : "(id) - (ticket_id)"
    helpdesk_checklist ||--o{ helpdesk_checklisttask : "(id) - (checklist_id)"
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
    django_site {
        AutoField id PK
        CharField domain UK
        CharField name
    }
    pinax_invitations_joininvitation {
        BigAutoField id PK
        ForeignKey from_user_id FK
        ForeignKey to_user_id FK
        TextField message
        DateTimeField sent
        JoinInvitationStatus status
        OneToOneField signup_code_id FK, UK
    }
    pinax_invitations_invitationstat {
        BigAutoField id PK
        OneToOneField user_id FK, UK
        IntegerField invites_sent
        IntegerField invites_allocated
        IntegerField invites_accepted
    }
    pinax_teams_simpleteam {
        BigAutoField id PK
        SimpleTeamMemberAccess member_access
        SimpleTeamManagerAccess manager_access
    }
    pinax_teams_team {
        BigAutoField id PK
        TeamMemberAccess member_access
        TeamManagerAccess manager_access
        SlugField slug UK
        CharField name
        ImageField avatar
        TextField description
        ForeignKey creator_id FK
        DateTimeField created
    }
    pinax_teams_simplemembership {
        BigAutoField id PK
        SimpleMembershipState state
        SimpleMembershipRole role
        DateTimeField created
        ForeignKey team_id FK
        ForeignKey user_id FK
        ForeignKey invite_id FK
    }
    pinax_teams_membership {
        BigAutoField id PK
        MembershipState state
        MembershipRole role
        DateTimeField created
        ForeignKey team_id FK
        ForeignKey user_id FK
        ForeignKey invite_id FK
    }
    account_account {
        BigAutoField id PK
        OneToOneField user_id FK, UK
        AccountTimezone timezone
        AccountLanguage language
    }
    account_signupcode {
        BigAutoField id PK
        CharField code UK
        PositiveIntegerField max_uses
        DateTimeField expiry
        ForeignKey inviter_id FK
        EmailField email
        TextField notes
        DateTimeField sent
        DateTimeField created
        PositiveIntegerField use_count
    }
    account_signupcoderesult {
        BigAutoField id PK
        ForeignKey signup_code_id FK
        ForeignKey user_id FK
        DateTimeField timestamp
    }
    account_emailaddress {
        BigAutoField id PK
        ForeignKey user_id FK
        EmailField email UK
        BooleanField verified
        BooleanField primary
    }
    account_emailconfirmation {
        BigAutoField id PK
        ForeignKey email_address_id FK
        DateTimeField created
        DateTimeField sent
        CharField key UK
    }
    account_accountdeletion {
        BigAutoField id PK
        ForeignKey user_id FK
        EmailField email
        DateTimeField date_requested
        DateTimeField date_expunged
    }
    account_passwordhistory {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField password
        DateTimeField timestamp
    }
    account_passwordexpiry {
        BigAutoField id PK
        OneToOneField user_id FK, UK
        PositiveIntegerField expiry
    }
    helpdesk_queue {
        AutoField id PK
        CharField title
        SlugField slug UK "This slug is used when building ticket ID's. Once set, try not to change it or e-mailing may get messy."
        EmailField email_address "All outgoing e-mails for this queue will use this e-mail address. If you use IMAP or POP3, this should be the e-mail address for that mailbox."
        CharField locale "Locale of this queue. All correspondence in this queue will be in this language."
        BooleanField allow_public_submission "Should this queue be listed on the public submission form?"
        BooleanField allow_email_submission "Do you want to poll the e-mail box below for new tickets?"
        IntegerField escalate_days "For tickets which are not held, how often do you wish to increase their priority? Set to 0 for no escalation."
        CharField new_ticket_cc "If an e-mail address is entered here, then it will receive notification of all new tickets created for this queue. Enter a comma between multiple e-mail addresses."
        CharField updated_ticket_cc "If an e-mail address is entered here, then it will receive notification of all activity (new tickets, closed tickets, updates, reassignments, etc) for this queue. Separate multiple addresses with a comma."
        BooleanField enable_notifications_on_email_events "When an email arrives to either create a ticket or to interact with an existing discussion. Should email notifications be sent ? Note: the new_ticket_cc and updated_ticket_cc work independently of this feature"
        QueueEmailBoxType email_box_type "E-Mail server type for creating tickets automatically from a mailbox - both POP3 and IMAP are supported, as well as reading from a local directory."
        CharField email_box_host "Your e-mail server address - either the domain name or IP address. May be 'localhost'."
        IntegerField email_box_port "Port number to use for accessing e-mail. Default for POP3 is '110', and for IMAP is '143'. This may differ on some servers. Leave it blank to use the defaults."
        BooleanField email_box_ssl "Whether to use SSL for IMAP or POP3 - the default ports when using SSL are 993 for IMAP and 995 for POP3."
        CharField email_box_user "Username for accessing this mailbox."
        CharField email_box_pass "Password for the above username"
        CharField email_box_imap_folder "If using IMAP, what folder do you wish to fetch messages from? This allows you to use one IMAP account for multiple queues, by filtering messages on your IMAP server into separate folders. Default: INBOX."
        CharField email_box_local_dir "If using a local directory, what directory path do you wish to poll for new email? Example: /var/lib/mail/helpdesk/"
        CharField permission_name "Name used in the django.contrib.auth permission system"
        IntegerField email_box_interval "How often do you wish to check this mailbox? (in Minutes)"
        DateTimeField email_box_last_check
        QueueSocksProxyType socks_proxy_type "SOCKS4 or SOCKS5 allows you to proxy your connections through a SOCKS server."
        GenericIPAddressField socks_proxy_host "Socks proxy IP address. Default: 127.0.0.1"
        IntegerField socks_proxy_port "Socks proxy port number. Default: 9150 (default TOR port)"
        QueueLoggingType logging_type "Set the default logging level. All messages at that level or above will be logged to the directory set below. If no level is set, logging will be disabled."
        CharField logging_dir "If logging is enabled, what directory should we use to store log files for this queue? The standard logging mechanisms are used if no directory is set"
        ForeignKey default_owner_id FK
        DurationField dedicated_time "Time to be spent on this Queue in total"
    }
    helpdesk_ticket {
        AutoField id PK
        CharField title
        ForeignKey queue_id FK
        DateTimeField created "Date this ticket was first created"
        DateTimeField modified "Date this ticket was most recently changed."
        EmailField submitter_email "The submitter will receive an email for all public follow-ups left for this task."
        ForeignKey assigned_to_id FK
        TicketStatus status
        BooleanField on_hold "If a ticket is on hold, it will not automatically be escalated."
        TextField description "The content of the customers query."
        TextField resolution "The resolution provided to the customer by our staff."
        TicketPriority priority "1 = Highest Priority, 5 = Low Priority"
        DateTimeField due_date
        DateTimeField last_escalation "The date this ticket was last escalated - updated automatically by management/commands/escalate_tickets.py."
        CharField secret_key
        ForeignKey kbitem_id FK
        ForeignKey merged_to_id FK
    }
    helpdesk_followup {
        AutoField id PK
        ForeignKey ticket_id FK
        DateTimeField date
        CharField title
        TextField comment
        BooleanField public "Public tickets are viewable by the submitter and all staff, but non-public tickets can only be seen by staff."
        ForeignKey user_id FK
        FollowUpNewStatus new_status "If the status was changed, what was it changed to?"
        CharField message_id "The Message ID of the submitter's email."
        DurationField time_spent "Time spent on this follow up"
        JSONField email_recipients "Addresses successfully notified by e-mail about this follow-up. Null means recipients were not recorded."
    }
    helpdesk_ticketchange {
        AutoField id PK
        ForeignKey followup_id FK
        CharField field
        TextField old_value
        TextField new_value
    }
    helpdesk_followupattachment {
        AutoField id PK
        FileField file
        CharField filename
        CharField mime_type
        IntegerField size "Size of this file in bytes"
        ForeignKey followup_id FK
    }
    helpdesk_kbiattachment {
        AutoField id PK
        FileField file
        CharField filename
        CharField mime_type
        IntegerField size "Size of this file in bytes"
        ForeignKey kbitem_id FK
    }
    helpdesk_presetreply {
        AutoField id PK
        CharField name "Only used to assist users with selecting a reply - not shown to the user."
        TextField body "Context available: {{ ticket }} - ticket object (eg {{ ticket.title }}); {{ queue }} - The queue; and {{ user }} - the current user."
    }
    helpdesk_escalationexclusion {
        AutoField id PK
        CharField name
        DateField date "Date on which escalation should not happen"
    }
    helpdesk_emailtemplate {
        AutoField id PK
        CharField template_name
        CharField subject "This will be prefixed with '[ticket.ticket] ticket.title'. We recommend something simple such as '(Updated') or '(Closed)' - the same context is available as in plain_text, below."
        CharField heading "In HTML e-mails, this will be the heading at the top of the email - the same context is available as in plain_text, below."
        TextField plain_text "The context available to you includes {{ ticket }}, {{ queue }}, and depending on the time of the call: {{ resolution }} or {{ comment }}."
        TextField html "The same context is available here as in plain_text, above."
        CharField locale "Locale of this template."
    }
    helpdesk_kbcategory {
        AutoField id PK
        CharField name
        CharField title
        SlugField slug
        TextField description
        ForeignKey queue_id FK
        BooleanField public
    }
    helpdesk_kbitem {
        AutoField id PK
        ForeignKey category_id FK
        CharField title
        TextField question
        TextField answer
        IntegerField votes "Total number of votes cast for this item"
        IntegerField recommendations "Number of votes for this item which were POSITIVE."
        DateTimeField last_updated "The date on which this question was most recently changed."
        ForeignKey team_id FK
        PositiveIntegerField order
        BooleanField enabled
    }
    helpdesk_savedsearch {
        AutoField id PK
        ForeignKey user_id FK
        CharField title "User-provided name for this query"
        BooleanField shared "Should other users see this query?"
        TextField query "Pickled query object. Be wary changing this."
    }
    helpdesk_usersettings {
        AutoField id PK
        OneToOneField user_id FK, UK
        TextField settings_pickled "DEPRECATED! This is a base64-encoded representation of a pickled Python dictionary. Do not change this field via the admin."
        BooleanField login_view_ticketlist "Display the ticket list upon login? Otherwise, the dashboard is shown."
        BooleanField email_on_ticket_change "If you're the ticket owner and the ticket is changed via the web by somebody else,do you want to receive an e-mail?"
        BooleanField email_on_ticket_assign "If you are assigned a ticket via the web, do you want to receive an e-mail?"
        UserSettingsTicketsPerPage tickets_per_page "How many tickets do you want to see on the Ticket List page?"
        BooleanField use_email_as_submitter "When you submit a ticket, do you want to automatically use your e-mail address as the submitter address? You can type a different e-mail address when entering the ticket if needed, this option only changes the default."
        UserSettingsTicketRespondLayout ticket_respond_layout "Where should the respond form appear when viewing a ticket?"
    }
    helpdesk_ignoreemail {
        AutoField id PK
        CharField name
        DateField date "Date on which this e-mail address was added"
        CharField email_address "Enter a full e-mail address, or portions with wildcards, eg *@domain.com or postmaster@*."
        BooleanField keep_in_mailbox "Do you want to save emails from this address in the mailbox? If this is unticked, emails from this address will be deleted."
    }
    helpdesk_ticketcc {
        AutoField id PK
        ForeignKey ticket_id FK
        ForeignKey user_id FK "User who wishes to receive updates for this ticket."
        EmailField email "For non-user followers, enter their e-mail address"
        BooleanField can_view "Can this CC login to view the ticket details?"
        BooleanField can_update "Can this CC login and update the ticket?"
    }
    helpdesk_customfield {
        AutoField id PK
        SlugField name UK "As used in the database and behind the scenes. Must be unique and consist of only lowercase letters with no punctuation."
        CharField label "The display label for this field"
        TextField help_text "Shown to the user when editing the ticket"
        CustomFieldDataType data_type "Allows you to restrict the data entered into this field"
        IntegerField max_length
        IntegerField decimal_places "Only used for decimal fields"
        BooleanField empty_selection_list "Only for List: adds an empty first entry to the choices list, which enforces that the user makes an active choice."
        TextField list_values "For list fields only. Enter one option per line."
        IntegerField ordering "Lower numbers are displayed first; higher numbers are listed later"
        BooleanField required "Does the user have to enter a value for this field?"
        BooleanField staff_only "If this is ticked, then the public submission form will NOT show this field"
    }
    helpdesk_ticketcustomfieldvalue {
        AutoField id PK
        ForeignKey ticket_id FK
        ForeignKey field_id FK
        TextField value
    }
    helpdesk_ticketdependency {
        AutoField id PK
        ForeignKey ticket_id FK
        ForeignKey depends_on_id FK
    }
    helpdesk_checklisttemplate {
        AutoField id PK
        CharField name
        JSONField task_list
    }
    helpdesk_checklist {
        AutoField id PK
        ForeignKey ticket_id FK
        CharField name
    }
    helpdesk_checklisttask {
        AutoField id PK
        ForeignKey checklist_id FK
        CharField description
        DateTimeField completion_date
        PositiveSmallIntegerField position
    }
```
