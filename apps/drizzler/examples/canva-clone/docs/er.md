```mermaid
erDiagram
    user ||--o{ account : "(id) - (userId)"
    user ||--o{ session : "(id) - (userId)"
    user ||--o{ authenticator : "(id) - (userId)"
    user ||--o{ project : "(id) - (userId)"
    user ||--o{ subscription : "(id) - (userId)"
    user {
        text id PK
        text name
        text email
        timestamp emailVerified
        text image
        text password
    }
    account {
        text userId FK
        text type
        text provider PK
        text providerAccountId PK
        text refresh_token
        text access_token
        integer expires_at
        text token_type
        text scope
        text id_token
        text session_state
    }
    session {
        text sessionToken PK
        text userId FK
        timestamp expires
    }
    verificationToken {
        text identifier PK
        text token PK
        timestamp expires
    }
    authenticator {
        text credentialID PK
        text userId PK, FK
        text providerAccountId
        text credentialPublicKey
        integer counter
        text credentialDeviceType
        boolean credentialBackedUp
        text transports
    }
    project {
        text id PK
        text name
        text userId FK
        text json
        integer height
        integer width
        text thumbnailUrl
        boolean isTemplate
        boolean isPro
        timestamp createdAt
        timestamp updatedAt
    }
    subscription {
        text id PK
        text userId FK
        text subscriptionId
        text customerId
        text priceId
        text status
        timestamp currentPeriodEnd
        timestamp createdAt
        timestamp updatedAt
    }
```
