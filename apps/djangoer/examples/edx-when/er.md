```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    edx_when_datepolicy ||--o{ edx_when_contentdate : "(id) - (policy_id)"
    auth_user ||--o{ edx_when_userdate : "(id) - (user_id)"
    edx_when_contentdate ||--o{ edx_when_userdate : "(id) - (content_date_id)"
    auth_user |o--o{ edx_when_userdate : "(id) - (actor_id)"
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
    edx_when_datepolicy {
        BigAutoField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        DateTimeField abs_date
        DurationField rel_date
    }
    edx_when_contentdate {
        BigAutoField id PK
        CourseKeyField course_id
        ForeignKey policy_id FK
        UsageKeyField location
        CharField field
        BooleanField active
        CharField block_type
        CharField assignment_title
        CharField course_name
        CharField subsection_name
    }
    edx_when_userdate {
        BigAutoField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        ForeignKey user_id FK
        ForeignKey content_date_id FK
        DateTimeField abs_date
        DurationField rel_date
        TextField reason
        ForeignKey actor_id FK
        UsageKeyField first_component_block_id
        BooleanField is_content_gated
    }
```
