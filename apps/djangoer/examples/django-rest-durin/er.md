```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user ||--o{ durin_authtoken : "(id) - (user_id)"
    durin_client ||--o{ durin_authtoken : "(id) - (client_id)"
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
    durin_client {
        BigAutoField id PK
        CharField name UK "A unique identification name for the client."
        DurationField token_ttl "Token Time To Live (TTL) in timedelta. Format: <code>DAYS HH:MM:SS</code>."
        CharField throttle_rate "Follows the same format as DRF's throttle rates.             Format: <em>'number_of_requests/period'</em>             where period should be one of: ('s', 'm', 'h', 'd').             Example: '100/h' implies 100 requests each hour."
    }
    durin_authtoken {
        BigAutoField id PK
        CharField token UK "Token is auto-generated on save."
        ForeignKey user_id FK
        ForeignKey client_id FK
        DateTimeField created
        DateTimeField expiry
    }
```
