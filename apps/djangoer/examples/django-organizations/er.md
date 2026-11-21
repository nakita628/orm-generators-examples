```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user ||--o{ organizations_organizationuser : "(id) - (user_id)"
    organizations_organization ||--o{ organizations_organizationuser : "(id) - (organization_id)"
    organizations_organizationuser ||--o| organizations_organizationowner : "(id) - (organization_user_id)"
    organizations_organization ||--o| organizations_organizationowner : "(id) - (organization_id)"
    auth_user ||--o{ organizations_organizationinvitation : "(id) - (invited_by_id)"
    auth_user |o--o{ organizations_organizationinvitation : "(id) - (invitee_id)"
    organizations_organization ||--o{ organizations_organizationinvitation : "(id) - (organization_id)"
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
    organizations_organization {
        AutoField id PK
        CharField name "The name of the organization"
        BooleanField is_active
        AutoCreatedField created
        AutoLastModifiedField modified
        SlugField slug UK "The name in all lowercase, suitable for URL identification"
    }
    organizations_organizationuser {
        AutoField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        BooleanField is_admin
        ForeignKey user_id FK
        ForeignKey organization_id FK
    }
    organizations_organizationowner {
        AutoField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        OneToOneField organization_user_id FK, UK
        OneToOneField organization_id FK, UK
    }
    organizations_organizationinvitation {
        AutoField id PK
        UUIDField guid
        CharField invitee_identifier "The contact identifier for the invitee, email, phone number, social media handle, etc."
        AutoCreatedField created
        AutoLastModifiedField modified
        ForeignKey invited_by_id FK
        ForeignKey invitee_id FK
        ForeignKey organization_id FK
    }
```
