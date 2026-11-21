```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user ||--o{ data_wizard_run : "(id) - (user_id)"
    django_content_type |o--o{ data_wizard_run : "(id) - (content_type_id)"
    data_wizard_run ||--o{ data_wizard_runlog : "(id) - (run_id)"
    data_wizard_run ||--o{ data_wizard_range : "(id) - (run_id)"
    data_wizard_identifier ||--o{ data_wizard_range : "(id) - (identifier_id)"
    data_wizard_run ||--o{ data_wizard_record : "(id) - (run_id)"
    django_content_type |o--o{ data_wizard_record : "(id) - (content_type_id)"
    auth_user |o--o{ sources_filesource : "(id) - (user_id)"
    auth_user |o--o{ sources_urlsource : "(id) - (user_id)"
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
    data_wizard_run {
        AutoField id PK
        ForeignKey user_id FK
        IntegerField record_count
        CharField loader
        CharField serializer
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
    }
    data_wizard_runlog {
        AutoField id PK
        ForeignKey run_id FK
        CharField event
        DateTimeField date
    }
    data_wizard_identifier {
        AutoField id PK
        CharField serializer
        CharField name
        CharField value
        CharField field
        CharField attr_field
        PositiveIntegerField attr_id
        BooleanField resolved
    }
    data_wizard_range {
        AutoField id PK
        ForeignKey run_id FK
        ForeignKey identifier_id FK
        RangeType type
        IntegerField header_col
        IntegerField start_col
        IntegerField end_col
        IntegerField header_row
        IntegerField start_row
        IntegerField end_row
        IntegerField count
    }
    data_wizard_record {
        AutoField id PK
        ForeignKey run_id FK
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
        PositiveIntegerField row
        BooleanField success
        TextField fail_reason
    }
    sources_filesource {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField name
        FileField file
        DateTimeField date
    }
    sources_urlsource {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField name
        URLField url
        DateTimeField date
    }
```
