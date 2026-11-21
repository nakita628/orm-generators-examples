```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    axes_accessattempt ||--o| axes_accessattemptexpiration : "(id) - (access_attempt_id)"
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
    axes_accessfailurelog {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        BooleanField locked_out
    }
    axes_accessattempt {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        TextField get_data
        TextField post_data
        PositiveIntegerField failures_since_start
    }
    axes_accessattemptexpiration {
        OneToOneField access_attempt_id PK
        DateTimeField expires_at "The time when access attempt expires and is no longer valid."
    }
    axes_accesslog {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        DateTimeField logout_time
        CharField session_hash
    }
```
