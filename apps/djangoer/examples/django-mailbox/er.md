```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    django_mailbox_mailbox ||--o{ django_mailbox_message : "(id) - (mailbox_id)"
    django_mailbox_message |o--o{ django_mailbox_message : "(id) - (in_reply_to_id)"
    django_mailbox_message |o--o{ django_mailbox_messageattachment : "(id) - (message_id)"
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
    django_mailbox_mailbox {
        AutoField id PK
        CharField name
        CharField uri "Example: imap+ssl://myusername:mypassword@someserver <br /><br />Internet transports include 'imap' and 'pop3'; common local file transports include 'maildir', 'mbox', and less commonly 'babyl', 'mh', and 'mmdf'. <br /><br />Be sure to urlencode your username and password should they contain illegal characters (like @, :, etc)."
        CharField from_email "Example: MailBot &lt;mailbot@yourdomain.com&gt;<br />'From' header to set for outgoing email.<br /><br />If you do not use this e-mail inbox for outgoing mail, this setting is unnecessary.<br />If you send e-mail without setting this, your 'From' header will'be set to match the setting `DEFAULT_FROM_EMAIL`."
        BooleanField active "Check this e-mail inbox for new e-mail messages during polling cycles.  This checkbox does not have an effect upon whether mail is collected here when this mailbox receives mail from a pipe, and does not affect whether e-mail messages can be dispatched from this mailbox."
        DateTimeField last_polling "The time of last successful polling for messages.It is blank for new mailboxes and is not set for mailboxes that only receive messages via a pipe."
    }
    django_mailbox_message {
        AutoField id PK
        ForeignKey mailbox_id FK
        CharField subject
        CharField message_id
        ForeignKey in_reply_to_id FK
        CharField from_header
        TextField to_header
        BooleanField outgoing
        TextField body
        BooleanField encoded "True if the e-mail body is Base64 encoded"
        DateTimeField processed
        DateTimeField read
        FileField eml "Original full content of message"
    }
    django_mailbox_messageattachment {
        AutoField id PK
        ForeignKey message_id FK
        TextField headers
        FileField document
    }
```
