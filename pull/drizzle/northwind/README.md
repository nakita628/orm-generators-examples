# drizzle/northwind

The tables drizzle-kit creates from the Drizzle schema of [apps/drizzler/examples/northwind](../../../apps/drizzler/examples/northwind), as SQL: [source.sql](source.sql). `prisma db pull` writes them up as [schema.prisma](schema.prisma), and from that write-up Prisma and every ORM hekireki generates create tables of their own; their SQL is in [sql/](sql), dumped the way source.sql is, and held to it here.

`=` identical · `≈` the same but for constraint and index names, and the order of columns and statements · `≠` tables written differently, or statements one side has and the other does not. The last column holds each ORM's SQL to what Prisma itself creates from the same write-up: a difference there is the ORM's or hekireki's, one only against source.sql is something the write-up does not say.

| From the write-up | Creates the tables with | SQL | Against source.sql | Against prisma.sql |
| --- | --- | --- | --- | --- |
| Prisma: schema.prisma | `prisma db push` | [prisma.sql](sql/prisma.sql) | = identical |  |
| hekireki-drizzle | `drizzle-kit export` | [drizzle.sql](sql/drizzle.sql) | ≠ 1 table only in source.sql ([diff](sql/drizzle.diff)) | ≠ 1 |
| hekireki-kysely | — | — | writes types only | — |
| hekireki-sqlalchemy | `metadata.create_all()` | [sqlalchemy.sql](sql/sqlalchemy.sql) | ≠ 1 table only in source.sql ([diff](sql/sqlalchemy.diff)) | ≠ 1 |
| hekireki-django | `makemigrations`, `migrate` | [django.sql](sql/django.sql) | ≠ 5 tables written differently, 1 table only in source.sql ([diff](sql/django.diff)) | ≠ 6 |
| hekireki-gorm | `AutoMigrate` | [gorm.sql](sql/gorm.sql) | ≠ 5 tables written differently, 1 table only in source.sql ([diff](sql/gorm.diff)) | ≠ 6 |
| hekireki-sea-orm | `Schema::create_table_from_entity` | [sea-orm.sql](sql/sea-orm.sql) | ≠ 1 table written differently, 1 table only in source.sql ([diff](sql/sea-orm.diff)) | ≠ 2 |
| hekireki-ecto | — | — | writes schemas only (no migrations) | — |
| hekireki-activerecord | — | — | writes models only (no migrations) | — |
| hekireki-eloquent | — | — | writes models only (no migrations) | — |
| hekireki-efcore | `EnsureCreated()` | [efcore.sql](sql/efcore.sql) | ≠ 1 table only in source.sql ([diff](sql/efcore.diff)) | ≠ 1 |
| hekireki-exposed | `SchemaUtils.create` | [exposed.sql](sql/exposed.sql) | ≠ 1 table only in source.sql ([diff](sql/exposed.diff)) | ≠ 1 |

6 models, postgresql. hekireki 0.52.2, Prisma 8.1.0-dev.9. The same code run against source.sql's tables: [results/northwind.md](../results/northwind.md).

<details><summary>hekireki-drizzle: 1 table only in source.sql</summary>

```diff
- public.order_details (the table is only in source.sql)
```

</details>

<details><summary>hekireki-sqlalchemy: 1 table only in source.sql</summary>

```diff
- public.order_details (the table is only in source.sql)
```

</details>

<details><summary>hekireki-django: 5 tables written differently, 1 table only in source.sql</summary>

```diff
@@ public.customers
- company_name character varying NOT NULL
+ company_name text NOT NULL
- contact_name character varying NOT NULL
+ contact_name text NOT NULL
- contact_title character varying NOT NULL
+ contact_title text NOT NULL
- address character varying NOT NULL
+ address text NOT NULL
- city character varying NOT NULL
+ city text NOT NULL
- postal_code character varying
+ postal_code text
- region character varying
+ region text
- country character varying NOT NULL
+ country text NOT NULL
- phone character varying NOT NULL
+ phone text NOT NULL
- fax character varying
+ fax text
+ INDEX USING btree (id varchar_pattern_ops)
@@ public.employees
- id character varying NOT NULL
+ id text NOT NULL
- last_name character varying NOT NULL
+ last_name text NOT NULL
- first_name character varying
+ first_name text
- title character varying NOT NULL
+ title text NOT NULL
- title_of_courtesy character varying NOT NULL
+ title_of_courtesy text NOT NULL
- address character varying NOT NULL
+ address text NOT NULL
- city character varying NOT NULL
+ city text NOT NULL
- postal_code character varying NOT NULL
+ postal_code text NOT NULL
- country character varying NOT NULL
+ country text NOT NULL
- home_phone character varying NOT NULL
+ home_phone text NOT NULL
- recipient_id character varying
+ recipient_id text
- FOREIGN KEY (recipient_id) REFERENCES public.employees(id)
+ FOREIGN KEY (recipient_id) REFERENCES public.employees(id) DEFERRABLE INITIALLY DEFERRED
+ INDEX USING btree (id text_pattern_ops)
- public.order_details (the table is only in source.sql)
@@ public.orders
- id character varying NOT NULL
+ id text NOT NULL
- ship_name character varying NOT NULL
+ ship_name text NOT NULL
- ship_city character varying NOT NULL
+ ship_city text NOT NULL
- ship_region character varying
+ ship_region text
- ship_postal_code character varying
+ ship_postal_code text
- ship_country character varying NOT NULL
+ ship_country text NOT NULL
- customer_id character varying NOT NULL
+ customer_id character varying(5) NOT NULL
- employee_id character varying NOT NULL
+ employee_id text NOT NULL
- FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON DELETE CASCADE
+ FOREIGN KEY (customer_id) REFERENCES public.customers(id) DEFERRABLE INITIALLY DEFERRED
- FOREIGN KEY (employee_id) REFERENCES public.employees(id) ON DELETE CASCADE
+ FOREIGN KEY (employee_id) REFERENCES public.employees(id) DEFERRABLE INITIALLY DEFERRED
+ INDEX USING btree (id text_pattern_ops)
@@ public.products
- id character varying NOT NULL
+ id text NOT NULL
- name character varying NOT NULL
+ name text NOT NULL
- qt_per_unit character varying NOT NULL
+ qt_per_unit text NOT NULL
- supplier_id character varying NOT NULL
+ supplier_id text NOT NULL
- FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) ON DELETE CASCADE
+ FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) DEFERRABLE INITIALLY DEFERRED
+ INDEX USING btree (id text_pattern_ops)
@@ public.suppliers
- id character varying NOT NULL
+ id text NOT NULL
- company_name character varying NOT NULL
+ company_name text NOT NULL
- contact_name character varying NOT NULL
+ contact_name text NOT NULL
- contact_title character varying NOT NULL
+ contact_title text NOT NULL
- address character varying NOT NULL
+ address text NOT NULL
- city character varying NOT NULL
+ city text NOT NULL
- region character varying
+ region text
- postal_code character varying NOT NULL
+ postal_code text NOT NULL
- country character varying NOT NULL
+ country text NOT NULL
- phone character varying NOT NULL
+ phone text NOT NULL
+ INDEX USING btree (id text_pattern_ops)
```

</details>

<details><summary>hekireki-gorm: 5 tables written differently, 1 table only in source.sql</summary>

```diff
@@ public.customers
- company_name character varying NOT NULL
+ company_name text NOT NULL
- contact_name character varying NOT NULL
+ contact_name text NOT NULL
- contact_title character varying NOT NULL
+ contact_title text NOT NULL
- address character varying NOT NULL
+ address text NOT NULL
- city character varying NOT NULL
+ city text NOT NULL
- postal_code character varying
+ postal_code text
- region character varying
+ region text
- country character varying NOT NULL
+ country text NOT NULL
- phone character varying NOT NULL
+ phone text NOT NULL
- fax character varying
+ fax text
@@ public.employees
- id character varying NOT NULL
+ id text NOT NULL
- last_name character varying NOT NULL
+ last_name text NOT NULL
- first_name character varying
+ first_name text
- title character varying NOT NULL
+ title text NOT NULL
- title_of_courtesy character varying NOT NULL
+ title_of_courtesy text NOT NULL
- address character varying NOT NULL
+ address text NOT NULL
- city character varying NOT NULL
+ city text NOT NULL
- postal_code character varying NOT NULL
+ postal_code text NOT NULL
- country character varying NOT NULL
+ country text NOT NULL
- home_phone character varying NOT NULL
+ home_phone text NOT NULL
- extension integer NOT NULL
+ extension bigint NOT NULL
- recipient_id character varying
+ recipient_id text
- public.order_details (the table is only in source.sql)
@@ public.orders
- id character varying NOT NULL
+ id text NOT NULL
- ship_via integer NOT NULL
+ ship_via bigint NOT NULL
- freight double precision NOT NULL
+ freight numeric NOT NULL
- ship_name character varying NOT NULL
+ ship_name text NOT NULL
- ship_city character varying NOT NULL
+ ship_city text NOT NULL
- ship_region character varying
+ ship_region text
- ship_postal_code character varying
+ ship_postal_code text
- ship_country character varying NOT NULL
+ ship_country text NOT NULL
- customer_id character varying NOT NULL
+ customer_id character varying(5) NOT NULL
- employee_id character varying NOT NULL
+ employee_id text NOT NULL
@@ public.products
- id character varying NOT NULL
+ id text NOT NULL
- name character varying NOT NULL
+ name text NOT NULL
- qt_per_unit character varying NOT NULL
+ qt_per_unit text NOT NULL
- unit_price double precision NOT NULL
+ unit_price numeric NOT NULL
- units_in_stock integer NOT NULL
+ units_in_stock bigint NOT NULL
- units_on_order integer NOT NULL
+ units_on_order bigint NOT NULL
- reorder_level integer NOT NULL
+ reorder_level bigint NOT NULL
- discontinued integer NOT NULL
+ discontinued bigint NOT NULL
- supplier_id character varying NOT NULL
+ supplier_id text NOT NULL
@@ public.suppliers
- id character varying NOT NULL
+ id text NOT NULL
- company_name character varying NOT NULL
+ company_name text NOT NULL
- contact_name character varying NOT NULL
+ contact_name text NOT NULL
- contact_title character varying NOT NULL
+ contact_title text NOT NULL
- address character varying NOT NULL
+ address text NOT NULL
- city character varying NOT NULL
+ city text NOT NULL
- region character varying
+ region text
- postal_code character varying NOT NULL
+ postal_code text NOT NULL
- country character varying NOT NULL
+ country text NOT NULL
- phone character varying NOT NULL
+ phone text NOT NULL
```

</details>

<details><summary>hekireki-sea-orm: 1 table written differently, 1 table only in source.sql</summary>

```diff
@@ public.employees
- notes text NOT NULL
+ notes character varying NOT NULL
- public.order_details (the table is only in source.sql)
```

</details>

<details><summary>hekireki-efcore: 1 table only in source.sql</summary>

```diff
- public.order_details (the table is only in source.sql)
```

</details>

<details><summary>hekireki-exposed: 1 table only in source.sql</summary>

```diff
- public.order_details (the table is only in source.sql)
```

</details>

Written by `node pull/pull.mjs`; do not edit.
