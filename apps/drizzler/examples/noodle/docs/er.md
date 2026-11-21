```mermaid
erDiagram
    early_access {
        uuid id PK
        text name
        text email
        text reason
        boolean approved
        timestamp created_at
        timestamp invitation_sent_at
    }
    modules {
        uuid id PK
        text user_id
        text name
        text description
        text code
        text icon
        text color
        boolean archived
        integer credits
        timestamp created_at
        timestamp modified_at
        timestamp last_visited
    }
```
