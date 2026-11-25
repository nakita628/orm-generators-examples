```mermaid
erDiagram
    users ||--o{ sessions : "(id) - (user_id)"
    users ||--o{ accounts : "(id) - (user_id)"
    posts ||--o{ comments : "(slug) - (post_slug)"
    users |o--o{ comments : "(id) - (user_id)"
    comments |o--o{ comments : "(id) - (parent_id)"
    users ||--o{ comment_votes : "(id) - (user_id)"
    comments ||--o{ comment_votes : "(id) - (comment_id)"
    users ||--o| comment_notification_preferences : "(id) - (user_id)"
    users ||--o{ comment_reply_unsubscribes : "(id) - (user_id)"
    comments ||--o{ comment_reply_unsubscribes : "(id) - (comment_id)"
    users ||--o{ guestbook_messages : "(id) - (user_id)"
    posts ||--o{ post_likes : "(slug) - (post_slug)"
    users {
        text id PK
        text name
        text email
        boolean email_verified
        text image
        text role
        boolean banned
        text ban_reason
        timestamptz ban_expires
        text last_login_method
        timestamptz created_at
        timestamptz updated_at
    }
    sessions {
        text id PK
        timestamptz expires_at
        text token
        text ip_address
        text user_agent
        text user_id FK
        text impersonated_by
        timestamptz created_at
        timestamptz updated_at
    }
    accounts {
        text id PK
        text account_id
        text provider_id
        text user_id FK
        text access_token
        text refresh_token
        text id_token
        timestamptz access_token_expires_at
        timestamptz refresh_token_expires_at
        text scope
        text password
        timestamptz created_at
        timestamptz updated_at
    }
    verifications {
        text id PK
        text identifier
        text value
        timestamptz expires_at
        timestamptz created_at
        timestamptz updated_at
    }
    comments {
        text id PK
        text body
        text post_slug FK
        text user_id FK
        text parent_id FK
        boolean is_deleted
        timestamptz created_at
        timestamptz updated_at
    }
    comment_votes {
        text user_id PK, FK
        text comment_id PK, FK
        boolean is_like
    }
    comment_notification_preferences {
        text user_id PK, FK
        boolean comment_replies_enabled
    }
    comment_reply_unsubscribes {
        text user_id PK, FK
        text comment_id PK, FK
    }
    guestbook_messages {
        text id PK
        text body
        text user_id FK
        timestamptz created_at
        timestamptz updated_at
    }
    posts {
        text slug PK
        integer likes
        integer views
        timestamptz created_at
    }
    post_likes {
        text post_slug PK, FK
        text anon_key PK
        integer like_count
        timestamptz created_at
    }
```
