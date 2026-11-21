```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user |o--o{ import_export_celery_exportjob : "(id) - (author_id)"
    auth_user |o--o{ import_export_celery_exportjob : "(id) - (updated_by_id)"
    auth_user |o--o{ import_export_celery_importjob : "(id) - (author_id)"
    auth_user |o--o{ import_export_celery_importjob : "(id) - (updated_by_id)"
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
    import_export_celery_exportjob {
        BigAutoField id PK
        ImportExportFileField file
        DateTimeField processing_initiated
        CharField job_status
        CharField format
        CharField app_label
        CharField model
        CharField resource
        TextField queryset
        BooleanField email_on_completion
        TextField site_of_origin
        ForeignKey author_id FK
        ForeignKey updated_by_id FK
    }
    import_export_celery_importjob {
        BigAutoField id PK
        ImportExportFileField file
        DateTimeField processing_initiated
        DateTimeField imported
        CharField format
        ImportExportFileField change_summary
        TextField errors
        CharField model
        CharField job_status
        ForeignKey author_id FK
        ForeignKey updated_by_id FK
    }
```
