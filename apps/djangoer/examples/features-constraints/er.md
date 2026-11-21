```mermaid
erDiagram
    catalog_products {
        BigAutoField id PK
        CharField sku
        CharField name
        DecimalField price
        DecimalField discounted
        IntegerField stock
        BooleanField active
        CharField category
    }
    legacy_codes {
        CharField code PK
        CharField label
    }
    v_totals {
        BigAutoField id PK
        IntegerField total
    }
    constraints_booking {
        CharField room PK
        DateField day PK
        CharField guest
    }
```
