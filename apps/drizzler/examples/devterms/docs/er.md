```mermaid
erDiagram
    users ||..o{ sessions : "(id) - (user_id)"
    users ||..o{ definitions : "(id) - (user_id)"
    users ||..o{ reports : "(id) - (user_id)"
    definitions ||..o{ reports : "(id) - (definition_id)"
    definitions ||..o{ wotds : "(id) - (definition_id)"
    users {
        text id PK
        text name
        text role
        text email
        text avatar
        integer github_id
        integer created_at
    }
    sessions {
        text id PK
        text user_id FK
        integer expires_at
    }
    definitions {
        text id PK
        text user_id FK
        text status
        text term
        text definition
        text example
        integer upvotes
        integer downvotes
        integer created_at
    }
    reports {
        text id PK
        text user_id FK
        text definition_id FK
        integer read
        text reason
        integer created_at
    }
    wotds {
        text id PK
        text definition_id FK
        integer created_at
    }
```
