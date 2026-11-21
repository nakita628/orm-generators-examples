```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    django_content_type ||..o{ easyaudit_crudevent : "(id) - (content_type_id), no constraint"
    auth_user |o..o{ easyaudit_crudevent : "(id) - (user_id), no constraint"
    auth_user |o..o{ easyaudit_loginevent : "(id) - (user_id), no constraint"
    auth_user |o..o{ easyaudit_requestevent : "(id) - (user_id), no constraint"
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
    easyaudit_crudevent {
        AutoField id PK
        CRUDEventEventType event_type
        CharField object_id
        ForeignKey content_type_id FK
        TextField object_repr
        TextField object_json_repr
        TextField changed_fields
        ForeignKey user_id FK
        CharField user_pk_as_string "String version of the user pk"
        DateTimeField datetime
    }
    easyaudit_loginevent {
        AutoField id PK
        LoginEventLoginType login_type
        CharField username
        ForeignKey user_id FK
        CharField remote_ip
        DateTimeField datetime
    }
    easyaudit_requestevent {
        AutoField id PK
        CharField url
        CharField method
        TextField query_string
        ForeignKey user_id FK
        CharField remote_ip
        DateTimeField datetime
    }
```
