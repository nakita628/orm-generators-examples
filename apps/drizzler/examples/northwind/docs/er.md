```mermaid
erDiagram
    employees |o--o{ employees : "(id) - (recipient_id)"
    customers ||--o{ orders : "(id) - (customer_id)"
    employees ||--o{ orders : "(id) - (employee_id)"
    suppliers ||--o{ products : "(id) - (supplier_id)"
    orders ||--o{ order_details : "(id) - (order_id)"
    products ||--o{ order_details : "(id) - (product_id)"
    customers {
        varchar id PK
        varchar company_name
        varchar contact_name
        varchar contact_title
        varchar address
        varchar city
        varchar postal_code
        varchar region
        varchar country
        varchar phone
        varchar fax
    }
    employees {
        varchar id PK
        varchar last_name
        varchar first_name
        varchar title
        varchar title_of_courtesy
        date birth_date
        date hire_date
        varchar address
        varchar city
        varchar postal_code
        varchar country
        varchar home_phone
        integer extension
        text notes
        varchar recipient_id FK
    }
    orders {
        varchar id PK
        date order_date
        date required_date
        date shipped_date
        integer ship_via
        double_precision freight
        varchar ship_name
        varchar ship_city
        varchar ship_region
        varchar ship_postal_code
        varchar ship_country
        varchar customer_id FK
        varchar employee_id FK
    }
    suppliers {
        varchar id PK
        varchar company_name
        varchar contact_name
        varchar contact_title
        varchar address
        varchar city
        varchar region
        varchar postal_code
        varchar country
        varchar phone
    }
    products {
        varchar id PK
        varchar name
        varchar qt_per_unit
        double_precision unit_price
        integer units_in_stock
        integer units_on_order
        integer reorder_level
        integer discontinued
        varchar supplier_id FK
    }
    order_details {
        double_precision unit_price
        integer quantity
        double_precision discount
        varchar order_id FK
        varchar product_id FK
    }
```
