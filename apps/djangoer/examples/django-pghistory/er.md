```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
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
    pghistory_context {
        UUIDField id PK
        DateTimeField created_at
        DateTimeField updated_at
        JSONField metadata
    }
    pghistory_events {
        TextField pgh_slug PK "The unique identifier across all event tables."
        CharField pgh_model "The event model."
        BigIntegerField pgh_id "The primary key of the event."
        DateTimeField pgh_created_at "When the event was created."
        TextField pgh_label "The event label."
        JSONField pgh_data "The raw data of the event."
        JSONField pgh_diff "The diff between the previous event of the same label."
        UUIDField pgh_context_id "The context UUID."
        JSONField pgh_context "The context associated with the event."
        CharField pgh_obj_model "The object model."
        TextField pgh_obj_id "The primary key of the object."
    }
```
