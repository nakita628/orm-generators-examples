```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    django_content_type ||--o{ django_comments : "(id) - (content_type_id)"
    django_site ||--o{ django_comments : "(id) - (site_id)"
    auth_user |o--o{ django_comments : "(id) - (user_id)"
    auth_user ||--o{ django_comment_flags : "(id) - (user_id)"
    django_comments ||--o{ django_comment_flags : "(id) - (comment_id)"
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
    django_site {
        AutoField id PK
        CharField domain UK
        CharField name
    }
    django_comments {
        AutoField id PK
        ForeignKey content_type_id FK
        CharField object_pk
        ForeignKey site_id FK
        ForeignKey user_id FK
        CharField user_name
        EmailField user_email
        URLField user_url
        TextField comment
        DateTimeField submit_date
        GenericIPAddressField ip_address
        BooleanField is_public "Uncheck this box to make the comment effectively disappear from the site."
        BooleanField is_removed "Check this box if the comment is inappropriate. A 'This comment has been removed' message will be displayed instead."
    }
    django_comment_flags {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey comment_id FK
        CharField flag
        DateTimeField flag_date
    }
```
