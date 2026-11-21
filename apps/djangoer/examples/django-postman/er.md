```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user |o--o{ postman_message : "(id) - (sender_id)"
    auth_user |o--o{ postman_message : "(id) - (recipient_id)"
    postman_message |o--o{ postman_message : "(id) - (parent_id)"
    postman_message |o--o{ postman_message : "(id) - (thread_id)"
    auth_user |o--o{ postman_message : "(id) - (moderation_by_id)"
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
    postman_message {
        AutoField id PK
        CharField subject
        TextField body
        ForeignKey sender_id FK
        ForeignKey recipient_id FK
        EmailField email
        ForeignKey parent_id FK
        ForeignKey thread_id FK
        DateTimeField sent_at
        DateTimeField read_at
        DateTimeField replied_at
        BooleanField sender_archived
        BooleanField recipient_archived
        DateTimeField sender_deleted_at
        DateTimeField recipient_deleted_at
        MessageModerationStatus moderation_status
        ForeignKey moderation_by_id FK
        DateTimeField moderation_date
        CharField moderation_reason
    }
```
