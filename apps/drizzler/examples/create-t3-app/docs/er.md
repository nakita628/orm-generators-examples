```mermaid
erDiagram
    project1_user ||--o{ project1_post : "(id) - (createdById)"
    project1_user ||--o{ project1_account : "(id) - (userId)"
    project1_user ||--o{ project1_session : "(id) - (userId)"
    project1_post {
        integer id PK
        varchar name
        varchar createdById FK
        timestamptz createdAt
        timestamptz updatedAt
    }
    project1_user {
        varchar id PK
        varchar name
        varchar email
        timestamptz emailVerified
        varchar image
    }
    project1_account {
        varchar userId FK
        varchar type
        varchar provider PK
        varchar providerAccountId PK
        text refresh_token
        text access_token
        integer expires_at
        varchar token_type
        varchar scope
        text id_token
        varchar session_state
    }
    project1_session {
        varchar sessionToken PK
        varchar userId FK
        timestamptz expires
    }
    project1_verification_token {
        varchar identifier PK
        varchar token PK
        timestamptz expires
    }
```
