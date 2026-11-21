```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user |o..o{ edx_name_affirmation_historicalverifiedname : "(id) - (user_id), no constraint"
    auth_user |o--o{ edx_name_affirmation_historicalverifiedname : "(id) - (history_user_id)"
    auth_user ||--o{ nameaffirmation_verifiedname : "(id) - (user_id)"
    auth_user |o--o{ nameaffirmation_verifiednameconfig : "(id) - (changed_by_id)"
    auth_user ||--o{ nameaffirmation_verifiednameconfig : "(id) - (user_id)"
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
    edx_name_affirmation_historicalverifiedname {
        BigIntegerField id
        AutoCreatedField created
        AutoLastModifiedField modified
        CharField verified_name
        CharField profile_name
        PositiveIntegerField verification_attempt_id
        PositiveIntegerField proctored_exam_attempt_id
        PositiveIntegerField platform_verification_attempt_id
        HistoricalVerifiedNameStatus status
        ForeignKey user_id FK
        AutoField history_id PK
        DateTimeField history_date
        CharField history_change_reason
        HistoricalVerifiedNameHistoryType history_type
        ForeignKey history_user_id FK
    }
    nameaffirmation_verifiedname {
        BigAutoField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        ForeignKey user_id FK
        CharField verified_name
        CharField profile_name
        PositiveIntegerField verification_attempt_id
        PositiveIntegerField proctored_exam_attempt_id
        PositiveIntegerField platform_verification_attempt_id
        VerifiedNameStatus status
    }
    nameaffirmation_verifiednameconfig {
        BigAutoField id PK
        DateTimeField change_date
        ForeignKey changed_by_id FK
        BooleanField enabled
        ForeignKey user_id FK
        BooleanField use_verified_name_for_certs
    }
```
