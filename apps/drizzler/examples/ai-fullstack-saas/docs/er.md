```mermaid
erDiagram
    user ||--o{ session : "(id) - (user_id)"
    user |o--o{ session : "(id) - (impersonated_by)"
    user ||--o{ account : "(id) - (user_id)"
    user |o--o{ message : "(id) - (sender_id)"
    user {
        uuid id PK
        text name
        integer age
        text image
        text email
        boolean email_verified
        boolean banned
        timestamp created_at
        timestamp updated_at
        roles role
    }
    session {
        uuid id PK
        uuid user_id FK
        text token
        timestamp expires_at
        text ip_address
        text user_agent
        uuid impersonated_by FK
        timestamp created_at
        timestamp updated_at
    }
    account {
        uuid id PK
        uuid user_id FK
        text account_id
        text provider_id
        text access_token
        text refresh_token
        timestamp access_token_expires_at
        timestamp refresh_token_expires_at
        text scope
        text id_token
        text password
        timestamp created_at
        timestamp updated_at
    }
    verification {
        uuid id PK
        text identifier
        text value
        timestamp expires_at
        timestamp created_at
        timestamp updated_at
    }
    example {
        uuid id PK
        text name
        text description
        timestamp created_at
    }
    message {
        uuid id PK
        text message
        uuid sender_id FK
        timestamp created_at
    }
```
