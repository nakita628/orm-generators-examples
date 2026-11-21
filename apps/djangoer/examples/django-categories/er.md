```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    categories_category |o--o{ categories_category : "(id) - (parent_id)"
    categories_category ||--o{ categories_categoryrelation : "(id) - (category_id)"
    django_content_type ||--o{ categories_categoryrelation : "(id) - (content_type_id)"
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
    categories_category {
        BigAutoField id PK
        TreeForeignKey parent_id FK
        CharField name
        SlugField slug UK
        BooleanField active
        FileField thumbnail
        IntegerField thumbnail_width
        IntegerField thumbnail_height
        IntegerField order
        CharField alternate_title "An alternative title to use on pages with this category."
        CharField alternate_url "An alternative URL to use instead of the one derived from the category hierarchy."
        TextField description
        CharField meta_keywords "Comma-separated keywords for search engines."
        TextField meta_extra "(Advanced) Any additional HTML to be placed verbatim in the &lt;head&gt;"
        PositiveIntegerField lft
        PositiveIntegerField rght
        PositiveIntegerField tree_id
        PositiveIntegerField level
    }
    categories_categoryrelation {
        BigAutoField id PK
        ForeignKey category_id FK
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
        CharField relation_type "A generic text field to tag a relation, like 'leadphoto'."
    }
```
