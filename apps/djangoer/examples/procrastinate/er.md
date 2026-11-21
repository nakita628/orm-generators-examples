```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    procrastinate_workers |o--o{ procrastinate_jobs : "(id) - (worker_id)"
    procrastinate_jobs ||--o{ procrastinate_events : "(id) - (job_id)"
    procrastinate_jobs |o--o{ procrastinate_periodic_defers : "(id) - (job_id)"
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
    procrastinate_workers {
        BigAutoField id PK
        DateTimeField last_heartbeat
    }
    procrastinate_jobs {
        BigAutoField id PK
        CharField queue_name
        CharField task_name
        IntegerField priority
        TextField lock UK
        JSONField args
        ProcrastinateJobStatus status
        DateTimeField scheduled_at
        IntegerField attempts
        TextField queueing_lock UK
        BooleanField abort_requested
        ForeignKey worker_id FK
    }
    procrastinate_events {
        BigAutoField id PK
        ForeignKey job_id FK
        ProcrastinateEventType type
        DateTimeField at
    }
    procrastinate_periodic_defers {
        BigAutoField id PK
        CharField task_name
        BigIntegerField defer_timestamp
        ForeignKey job_id FK
        CharField periodic_id
    }
```
