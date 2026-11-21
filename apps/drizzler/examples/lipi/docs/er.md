```mermaid
erDiagram
    lipi_workspaces |o--o{ lipi_folders : "(id) - (workspace_id)"
    lipi_workspaces |o--o{ lipi_files : "(id) - (workspace_id)"
    lipi_folders |o--o{ lipi_files : "(id) - (folder_id)"
    user ||--o{ lipi_accounts : "(id) - (user_id)"
    lipi_products |o--o{ lipi_prices : "(id) - (product_id)"
    lipi_prices |o--o{ lipi_subscriptions : "(id) - (price_id)"
    lipi_workspaces ||--o{ lipi_collaborators : "(id) - (workspace_id)"
    user ||--o{ lipi_collaborators : "(id) - (user_id)"
    user ||--o{ account : "(id) - (userId)"
    lipi_workspaces {
        uuid id PK
        text title
        text icon_id
        text data
        text logo
        text banner_url
        uuid workspace_owner_id
        boolean in_trash
        timestamptz created_at
    }
    lipi_folders {
        uuid id PK
        text title
        text icon_id
        text data
        text banner_url
        uuid workspace_id FK
        boolean in_trash
        timestamptz created_at
    }
    lipi_files {
        uuid id PK
        text title
        text icon_id
        text data
        text banner_url
        uuid workspace_id FK
        uuid folder_id FK
        boolean in_trash
        timestamptz created_at
    }
    lipi_accounts {
        uuid user_id FK
        jsonb billing_address
        timestamptz updated_at
        jsonb payment_method
    }
    lipi_customers {
        uuid id PK
        text stripe_customer_id
    }
    lipi_prices {
        text id PK
        text product_id FK
        boolean active
        text description
        bigint unit_amount "You can use { mode: 'bigint' } if numbers are exceeding js number limitations"
        text currency
        pricing_type type
        pricing_plan_interval interval
        integer interval_count
        integer trial_period_days
        jsonb metadata
    }
    lipi_products {
        text id PK
        boolean active
        text name
        text description
        text image
        jsonb metadata
    }
    lipi_subscriptions {
        text id PK
        uuid user_id
        subscription_status status
        jsonb metadata
        text price_id FK
        integer quantity
        boolean cancel_at_period_end
        timestamptz created
        timestamptz current_period_start
        timestamptz current_period_end
        timestamptz ended_at
        timestamptz cancel_at
        timestamptz canceled_at
        timestamptz trial_start
        timestamptz trial_end
    }
    lipi_collaborators {
        uuid id PK
        uuid workspace_id FK
        timestamptz created_at
        uuid user_id FK
    }
    user {
        uuid id PK
        text name
        text email
        text username
        text password
        timestamp emailVerified
        text image
    }
    account {
        uuid userId FK
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
    verificationToken {
        text identifier PK
        text token PK
        timestamp expires
    }
```
