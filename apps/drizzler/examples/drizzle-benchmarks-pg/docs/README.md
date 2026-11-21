# Schema docs

![ER diagram](er.svg)

<details>
<summary>The whole schema as SQL</summary>

```sql
CREATE TABLE "customers" (
	"id" serial PRIMARY KEY,
	"company_name" text NOT NULL,
	"contact_name" varchar NOT NULL,
	"contact_title" varchar NOT NULL,
	"address" varchar NOT NULL,
	"city" varchar NOT NULL,
	"postal_code" varchar,
	"region" varchar,
	"country" varchar NOT NULL,
	"phone" varchar NOT NULL,
	"fax" varchar
);

CREATE TABLE "order_details" (
	"unit_price" double precision NOT NULL,
	"quantity" integer NOT NULL,
	"discount" double precision NOT NULL,
	"order_id" integer NOT NULL,
	"product_id" integer NOT NULL
);

CREATE TABLE "employees" (
	"id" serial PRIMARY KEY,
	"last_name" varchar NOT NULL,
	"first_name" varchar,
	"title" varchar NOT NULL,
	"title_of_courtesy" varchar NOT NULL,
	"birth_date" date NOT NULL,
	"hire_date" date NOT NULL,
	"address" varchar NOT NULL,
	"city" varchar NOT NULL,
	"postal_code" varchar NOT NULL,
	"country" varchar NOT NULL,
	"home_phone" varchar NOT NULL,
	"extension" integer NOT NULL,
	"notes" text NOT NULL,
	"recipient_id" integer
);

CREATE TABLE "orders" (
	"id" serial PRIMARY KEY,
	"order_date" date NOT NULL,
	"required_date" date NOT NULL,
	"shipped_date" date,
	"ship_via" integer NOT NULL,
	"freight" double precision NOT NULL,
	"ship_name" varchar NOT NULL,
	"ship_city" varchar NOT NULL,
	"ship_region" varchar,
	"ship_postal_code" varchar,
	"ship_country" varchar NOT NULL,
	"customer_id" integer NOT NULL,
	"employee_id" integer NOT NULL
);

CREATE TABLE "products" (
	"id" serial PRIMARY KEY,
	"name" text NOT NULL,
	"quantity_per_unit" varchar NOT NULL,
	"unit_price" double precision NOT NULL,
	"units_in_stock" integer NOT NULL,
	"units_on_order" integer NOT NULL,
	"reorder_level" integer NOT NULL,
	"discontinued" integer NOT NULL,
	"supplier_id" serial
);

CREATE TABLE "suppliers" (
	"id" serial PRIMARY KEY,
	"company_name" varchar NOT NULL,
	"contact_name" varchar NOT NULL,
	"contact_title" varchar NOT NULL,
	"address" varchar NOT NULL,
	"city" varchar NOT NULL,
	"region" varchar,
	"postal_code" varchar NOT NULL,
	"country" varchar NOT NULL,
	"phone" varchar NOT NULL
);

CREATE INDEX "order_id_idx" ON "order_details" ("order_id");

CREATE INDEX "product_id_idx" ON "order_details" ("product_id");

CREATE INDEX "recepient_idx" ON "employees" ("recipient_id");

CREATE INDEX "supplier_idx" ON "products" ("supplier_id");

ALTER TABLE "order_details" ADD CONSTRAINT "order_details_order_id_orders_id_fkey" FOREIGN KEY ("order_id") REFERENCES "orders"("id") ON DELETE CASCADE;

ALTER TABLE "order_details" ADD CONSTRAINT "order_details_product_id_products_id_fkey" FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE;

ALTER TABLE "employees" ADD CONSTRAINT "employees_recipient_id_employees_id_fkey" FOREIGN KEY ("recipient_id") REFERENCES "employees"("id");

ALTER TABLE "orders" ADD CONSTRAINT "orders_customer_id_customers_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "customers"("id") ON DELETE CASCADE;

ALTER TABLE "orders" ADD CONSTRAINT "orders_employee_id_employees_id_fkey" FOREIGN KEY ("employee_id") REFERENCES "employees"("id") ON DELETE CASCADE;

ALTER TABLE "products" ADD CONSTRAINT "products_supplier_id_suppliers_id_fkey" FOREIGN KEY ("supplier_id") REFERENCES "suppliers"("id") ON DELETE CASCADE;
```

</details>

## Contents

| Table                         | Name in the database | Description |
| ----------------------------- | -------------------- | ----------- |
| [customers](#table-customers) | `customers`          |             |
| [employees](#table-employees) | `employees`          |             |
| [orders](#table-orders)       | `orders`             |             |
| [suppliers](#table-suppliers) | `suppliers`          |             |
| [products](#table-products)   | `products`           |             |
| [details](#table-details)     | `order_details`      |             |

## Tables

<a id="table-customers"></a>

### customers

![customers and the tables it relates to](tables/customers.svg)

| Column                           | Type      | Constraints | Default | Description |
| -------------------------------- | --------- | ----------- | ------- | ----------- |
| `id`                             | `serial`  | PK NOT NULL |         |             |
| `companyName` (`company_name`)   | `text`    | NOT NULL    |         |             |
| `contactName` (`contact_name`)   | `varchar` | NOT NULL    |         |             |
| `contactTitle` (`contact_title`) | `varchar` | NOT NULL    |         |             |
| `address`                        | `varchar` | NOT NULL    |         |             |
| `city`                           | `varchar` | NOT NULL    |         |             |
| `postalCode` (`postal_code`)     | `varchar` | NULL        |         |             |
| `region`                         | `varchar` | NULL        |         |             |
| `country`                        | `varchar` | NOT NULL    |         |             |
| `phone`                          | `varchar` | NOT NULL    |         |             |
| `fax`                            | `varchar` | NULL        |         |             |

#### Definition

```sql
CREATE TABLE "customers" (
	"id" serial PRIMARY KEY,
	"company_name" text NOT NULL,
	"contact_name" varchar NOT NULL,
	"contact_title" varchar NOT NULL,
	"address" varchar NOT NULL,
	"city" varchar NOT NULL,
	"postal_code" varchar,
	"region" varchar,
	"country" varchar NOT NULL,
	"phone" varchar NOT NULL,
	"fax" varchar
);
```

#### Referenced by

- `id` → [orders.customerId](#table-orders) (one to many, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of customers.

```ts
db.select().from(customers);
```

```sql
select "id", "company_name", "contact_name", "contact_title", "address", "city", "postal_code", "region", "country", "phone", "fax"
from "customers"
```

**Count**: How many rows customers has.

```ts
db.select({ count: count() }).from(customers);
```

```sql
select count(*)
from "customers"
```

**Find many**: The rows of customers, through the relational query API.

```ts
db.query.customers.findMany();
```

```sql
select "d0"."id" as "id", "d0"."company_name" as "companyName", "d0"."contact_name" as "contactName", "d0"."contact_title" as "contactTitle", "d0"."address" as "address", "d0"."city" as "city", "d0"."postal_code" as "postalCode", "d0"."region" as "region", "d0"."country" as "country", "d0"."phone" as "phone", "d0"."fax" as "fax"
from "customers" as "d0"
```

**Find by primary key**: One row of customers, through the relational query API.

```ts
db.query.customers.findFirst({ where: { id: 1 } });
```

```sql
select "d0"."id" as "id", "d0"."company_name" as "companyName", "d0"."contact_name" as "contactName", "d0"."contact_title" as "contactTitle", "d0"."address" as "address", "d0"."city" as "city", "d0"."postal_code" as "postalCode", "d0"."region" as "region", "d0"."country" as "country", "d0"."phone" as "phone", "d0"."fax" as "fax"
from "customers" as "d0"
where "d0"."id" = $1
limit $2
-- params: [1,1]
```

**Insert**: Adds a row to customers, giving the columns that have no default.

```ts
db.insert(customers).values({
  companyName: "companyName",
  contactName: "contactName",
  contactTitle: "contactTitle",
  address: "address",
  city: "city",
  country: "country",
  phone: "phone",
});
```

```sql
insert into "customers" ("id", "company_name", "contact_name", "contact_title", "address", "city", "postal_code", "region", "country", "phone", "fax")
values (default, $1, $2, $3, $4, $5, default, default, $6, $7, default)
-- params: ["companyName","contactName","contactTitle","address","city","country","phone"]
```

**Update**: Changes a row of customers, found by its primary key.

```ts
db.update(customers).set({ companyName: "companyName" }).where(eq(customers.id, 1));
```

```sql
update "customers"
set "company_name" = $1
where "customers"."id" = $2
-- params: ["companyName",1]
```

**Delete**: Deletes a row of customers, found by its primary key.

```ts
db.delete(customers).where(eq(customers.id, 1));
```

```sql
delete from "customers"
where "customers"."id" = $1
-- params: [1]
```

</details>

<a id="table-employees"></a>

### employees

![employees and the tables it relates to](tables/employees.svg)

| Column                                  | Type      | Constraints                             | Default | Description |
| --------------------------------------- | --------- | --------------------------------------- | ------- | ----------- |
| `id`                                    | `serial`  | PK NOT NULL                             |         |             |
| `lastName` (`last_name`)                | `varchar` | NOT NULL                                |         |             |
| `firstName` (`first_name`)              | `varchar` | NULL                                    |         |             |
| `title`                                 | `varchar` | NOT NULL                                |         |             |
| `titleOfCourtesy` (`title_of_courtesy`) | `varchar` | NOT NULL                                |         |             |
| `birthDate` (`birth_date`)              | `date`    | NOT NULL                                |         |             |
| `hireDate` (`hire_date`)                | `date`    | NOT NULL                                |         |             |
| `address`                               | `varchar` | NOT NULL                                |         |             |
| `city`                                  | `varchar` | NOT NULL                                |         |             |
| `postalCode` (`postal_code`)            | `varchar` | NOT NULL                                |         |             |
| `country`                               | `varchar` | NOT NULL                                |         |             |
| `homePhone` (`home_phone`)              | `varchar` | NOT NULL                                |         |             |
| `extension`                             | `integer` | NOT NULL                                |         |             |
| `notes`                                 | `text`    | NOT NULL                                |         |             |
| `recipientId` (`recipient_id`)          | `integer` | NULL → [employees.id](#table-employees) |         |             |

#### Definition

```sql
CREATE TABLE "employees" (
	"id" serial PRIMARY KEY,
	"last_name" varchar NOT NULL,
	"first_name" varchar,
	"title" varchar NOT NULL,
	"title_of_courtesy" varchar NOT NULL,
	"birth_date" date NOT NULL,
	"hire_date" date NOT NULL,
	"address" varchar NOT NULL,
	"city" varchar NOT NULL,
	"postal_code" varchar NOT NULL,
	"country" varchar NOT NULL,
	"home_phone" varchar NOT NULL,
	"extension" integer NOT NULL,
	"notes" text NOT NULL,
	"recipient_id" integer
);

CREATE INDEX "recepient_idx" ON "employees" ("recipient_id");

ALTER TABLE "employees" ADD CONSTRAINT "employees_recipient_id_employees_id_fkey" FOREIGN KEY ("recipient_id") REFERENCES "employees"("id");
```

#### Indexes

- INDEX `recipientId`

#### References

- `recipientId` → [employees.id](#table-employees) (many to one)

#### Referenced by

- `id` → [employees.recipientId](#table-employees) (one to many)
- `id` → [orders.employeeId](#table-orders) (one to many, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of employees.

```ts
db.select().from(employees);
```

```sql
select "id", "last_name", "first_name", "title", "title_of_courtesy", "birth_date", "hire_date", "address", "city", "postal_code", "country", "home_phone", "extension", "notes", "recipient_id"
from "employees"
```

**Count**: How many rows employees has.

```ts
db.select({ count: count() }).from(employees);
```

```sql
select count(*)
from "employees"
```

**Find many**: The rows of employees, through the relational query API.

```ts
db.query.employees.findMany();
```

```sql
select "d0"."id" as "id", "d0"."last_name" as "lastName", "d0"."first_name" as "firstName", "d0"."title" as "title", "d0"."title_of_courtesy" as "titleOfCourtesy", "d0"."birth_date" as "birthDate", "d0"."hire_date" as "hireDate", "d0"."address" as "address", "d0"."city" as "city", "d0"."postal_code" as "postalCode", "d0"."country" as "country", "d0"."home_phone" as "homePhone", "d0"."extension" as "extension", "d0"."notes" as "notes", "d0"."recipient_id" as "recipientId"
from "employees" as "d0"
```

**Find by primary key**: One row of employees, through the relational query API.

```ts
db.query.employees.findFirst({ where: { id: 1 } });
```

```sql
select "d0"."id" as "id", "d0"."last_name" as "lastName", "d0"."first_name" as "firstName", "d0"."title" as "title", "d0"."title_of_courtesy" as "titleOfCourtesy", "d0"."birth_date" as "birthDate", "d0"."hire_date" as "hireDate", "d0"."address" as "address", "d0"."city" as "city", "d0"."postal_code" as "postalCode", "d0"."country" as "country", "d0"."home_phone" as "homePhone", "d0"."extension" as "extension", "d0"."notes" as "notes", "d0"."recipient_id" as "recipientId"
from "employees" as "d0"
where "d0"."id" = $1
limit $2
-- params: [1,1]
```

**Insert**: Adds a row to employees, giving the columns that have no default.

```ts
db.insert(employees).values({
  lastName: "lastName",
  title: "title",
  titleOfCourtesy: "titleOfCourtesy",
  birthDate: new Date("2026-01-01T00:00:00.000Z"),
  hireDate: new Date("2026-01-01T00:00:00.000Z"),
  address: "address",
  city: "city",
  postalCode: "postalCode",
  country: "country",
  homePhone: "homePhone",
  extension: 1,
  notes: "notes",
});
```

```sql
insert into "employees" ("id", "last_name", "first_name", "title", "title_of_courtesy", "birth_date", "hire_date", "address", "city", "postal_code", "country", "home_phone", "extension", "notes", "recipient_id")
values (default, $1, default, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, default)
-- params: ["lastName","title","titleOfCourtesy","2026-01-01T00:00:00.000Z","2026-01-01T00:00:00.000Z","address","city","postalCode","country","homePhone",1,"notes"]
```

**Update**: Changes a row of employees, found by its primary key.

```ts
db.update(employees).set({ lastName: "lastName" }).where(eq(employees.id, 1));
```

```sql
update "employees"
set "last_name" = $1
where "employees"."id" = $2
-- params: ["lastName",1]
```

**Delete**: Deletes a row of employees, found by its primary key.

```ts
db.delete(employees).where(eq(employees.id, 1));
```

```sql
delete from "employees"
where "employees"."id" = $1
-- params: [1]
```

</details>

<a id="table-orders"></a>

### orders

![orders and the tables it relates to](tables/orders.svg)

| Column                                | Type              | Constraints                                 | Default | Description |
| ------------------------------------- | ----------------- | ------------------------------------------- | ------- | ----------- |
| `id`                                  | `serial`          | PK NOT NULL                                 |         |             |
| `orderDate` (`order_date`)            | `date`            | NOT NULL                                    |         |             |
| `requiredDate` (`required_date`)      | `date`            | NOT NULL                                    |         |             |
| `shippedDate` (`shipped_date`)        | `date`            | NULL                                        |         |             |
| `shipVia` (`ship_via`)                | `integer`         | NOT NULL                                    |         |             |
| `freight`                             | `doublePrecision` | NOT NULL                                    |         |             |
| `shipName` (`ship_name`)              | `varchar`         | NOT NULL                                    |         |             |
| `shipCity` (`ship_city`)              | `varchar`         | NOT NULL                                    |         |             |
| `shipRegion` (`ship_region`)          | `varchar`         | NULL                                        |         |             |
| `shipPostalCode` (`ship_postal_code`) | `varchar`         | NULL                                        |         |             |
| `shipCountry` (`ship_country`)        | `varchar`         | NOT NULL                                    |         |             |
| `customerId` (`customer_id`)          | `integer`         | NOT NULL → [customers.id](#table-customers) |         |             |
| `employeeId` (`employee_id`)          | `integer`         | NOT NULL → [employees.id](#table-employees) |         |             |

#### Definition

```sql
CREATE TABLE "orders" (
	"id" serial PRIMARY KEY,
	"order_date" date NOT NULL,
	"required_date" date NOT NULL,
	"shipped_date" date,
	"ship_via" integer NOT NULL,
	"freight" double precision NOT NULL,
	"ship_name" varchar NOT NULL,
	"ship_city" varchar NOT NULL,
	"ship_region" varchar,
	"ship_postal_code" varchar,
	"ship_country" varchar NOT NULL,
	"customer_id" integer NOT NULL,
	"employee_id" integer NOT NULL
);

ALTER TABLE "orders" ADD CONSTRAINT "orders_customer_id_customers_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "customers"("id") ON DELETE CASCADE;

ALTER TABLE "orders" ADD CONSTRAINT "orders_employee_id_employees_id_fkey" FOREIGN KEY ("employee_id") REFERENCES "employees"("id") ON DELETE CASCADE;
```

#### References

- `customerId` → [customers.id](#table-customers) (many to one, on delete cascade)
- `employeeId` → [employees.id](#table-employees) (many to one, on delete cascade)

#### Referenced by

- `id` → [details.orderId](#table-details) (one to many, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of orders.

```ts
db.select().from(orders);
```

```sql
select "id", "order_date", "required_date", "shipped_date", "ship_via", "freight", "ship_name", "ship_city", "ship_region", "ship_postal_code", "ship_country", "customer_id", "employee_id"
from "orders"
```

**Count**: How many rows orders has.

```ts
db.select({ count: count() }).from(orders);
```

```sql
select count(*)
from "orders"
```

**Find many**: The rows of orders, through the relational query API.

```ts
db.query.orders.findMany();
```

```sql
select "d0"."id" as "id", "d0"."order_date" as "orderDate", "d0"."required_date" as "requiredDate", "d0"."shipped_date" as "shippedDate", "d0"."ship_via" as "shipVia", "d0"."freight" as "freight", "d0"."ship_name" as "shipName", "d0"."ship_city" as "shipCity", "d0"."ship_region" as "shipRegion", "d0"."ship_postal_code" as "shipPostalCode", "d0"."ship_country" as "shipCountry", "d0"."customer_id" as "customerId", "d0"."employee_id" as "employeeId"
from "orders" as "d0"
```

**Find by primary key**: One row of orders, through the relational query API.

```ts
db.query.orders.findFirst({ where: { id: 1 } });
```

```sql
select "d0"."id" as "id", "d0"."order_date" as "orderDate", "d0"."required_date" as "requiredDate", "d0"."shipped_date" as "shippedDate", "d0"."ship_via" as "shipVia", "d0"."freight" as "freight", "d0"."ship_name" as "shipName", "d0"."ship_city" as "shipCity", "d0"."ship_region" as "shipRegion", "d0"."ship_postal_code" as "shipPostalCode", "d0"."ship_country" as "shipCountry", "d0"."customer_id" as "customerId", "d0"."employee_id" as "employeeId"
from "orders" as "d0"
where "d0"."id" = $1
limit $2
-- params: [1,1]
```

**Insert**: Adds a row to orders, giving the columns that have no default.

```ts
db.insert(orders).values({
  orderDate: new Date("2026-01-01T00:00:00.000Z"),
  requiredDate: new Date("2026-01-01T00:00:00.000Z"),
  shipVia: 1,
  freight: 1,
  shipName: "shipName",
  shipCity: "shipCity",
  shipCountry: "shipCountry",
  customerId: 1,
  employeeId: 1,
});
```

```sql
insert into "orders" ("id", "order_date", "required_date", "shipped_date", "ship_via", "freight", "ship_name", "ship_city", "ship_region", "ship_postal_code", "ship_country", "customer_id", "employee_id")
values (default, $1, $2, default, $3, $4, $5, $6, default, default, $7, $8, $9)
-- params: ["2026-01-01T00:00:00.000Z","2026-01-01T00:00:00.000Z",1,1,"shipName","shipCity","shipCountry",1,1]
```

**Update**: Changes a row of orders, found by its primary key.

```ts
db.update(orders)
  .set({ orderDate: new Date("2026-01-01T00:00:00.000Z") })
  .where(eq(orders.id, 1));
```

```sql
update "orders"
set "order_date" = $1
where "orders"."id" = $2
-- params: ["2026-01-01T00:00:00.000Z",1]
```

**Delete**: Deletes a row of orders, found by its primary key.

```ts
db.delete(orders).where(eq(orders.id, 1));
```

```sql
delete from "orders"
where "orders"."id" = $1
-- params: [1]
```

</details>

<a id="table-suppliers"></a>

### suppliers

![suppliers and the tables it relates to](tables/suppliers.svg)

| Column                           | Type      | Constraints | Default | Description |
| -------------------------------- | --------- | ----------- | ------- | ----------- |
| `id`                             | `serial`  | PK NOT NULL |         |             |
| `companyName` (`company_name`)   | `varchar` | NOT NULL    |         |             |
| `contactName` (`contact_name`)   | `varchar` | NOT NULL    |         |             |
| `contactTitle` (`contact_title`) | `varchar` | NOT NULL    |         |             |
| `address`                        | `varchar` | NOT NULL    |         |             |
| `city`                           | `varchar` | NOT NULL    |         |             |
| `region`                         | `varchar` | NULL        |         |             |
| `postalCode` (`postal_code`)     | `varchar` | NOT NULL    |         |             |
| `country`                        | `varchar` | NOT NULL    |         |             |
| `phone`                          | `varchar` | NOT NULL    |         |             |

#### Definition

```sql
CREATE TABLE "suppliers" (
	"id" serial PRIMARY KEY,
	"company_name" varchar NOT NULL,
	"contact_name" varchar NOT NULL,
	"contact_title" varchar NOT NULL,
	"address" varchar NOT NULL,
	"city" varchar NOT NULL,
	"region" varchar,
	"postal_code" varchar NOT NULL,
	"country" varchar NOT NULL,
	"phone" varchar NOT NULL
);
```

#### Referenced by

- `id` → [products.supplierId](#table-products) (one to many, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of suppliers.

```ts
db.select().from(suppliers);
```

```sql
select "id", "company_name", "contact_name", "contact_title", "address", "city", "region", "postal_code", "country", "phone"
from "suppliers"
```

**Count**: How many rows suppliers has.

```ts
db.select({ count: count() }).from(suppliers);
```

```sql
select count(*)
from "suppliers"
```

**Find many**: The rows of suppliers, through the relational query API.

```ts
db.query.suppliers.findMany();
```

```sql
select "d0"."id" as "id", "d0"."company_name" as "companyName", "d0"."contact_name" as "contactName", "d0"."contact_title" as "contactTitle", "d0"."address" as "address", "d0"."city" as "city", "d0"."region" as "region", "d0"."postal_code" as "postalCode", "d0"."country" as "country", "d0"."phone" as "phone"
from "suppliers" as "d0"
```

**Find by primary key**: One row of suppliers, through the relational query API.

```ts
db.query.suppliers.findFirst({ where: { id: 1 } });
```

```sql
select "d0"."id" as "id", "d0"."company_name" as "companyName", "d0"."contact_name" as "contactName", "d0"."contact_title" as "contactTitle", "d0"."address" as "address", "d0"."city" as "city", "d0"."region" as "region", "d0"."postal_code" as "postalCode", "d0"."country" as "country", "d0"."phone" as "phone"
from "suppliers" as "d0"
where "d0"."id" = $1
limit $2
-- params: [1,1]
```

**Insert**: Adds a row to suppliers, giving the columns that have no default.

```ts
db.insert(suppliers).values({
  companyName: "companyName",
  contactName: "contactName",
  contactTitle: "contactTitle",
  address: "address",
  city: "city",
  postalCode: "postalCode",
  country: "country",
  phone: "phone",
});
```

```sql
insert into "suppliers" ("id", "company_name", "contact_name", "contact_title", "address", "city", "region", "postal_code", "country", "phone")
values (default, $1, $2, $3, $4, $5, default, $6, $7, $8)
-- params: ["companyName","contactName","contactTitle","address","city","postalCode","country","phone"]
```

**Update**: Changes a row of suppliers, found by its primary key.

```ts
db.update(suppliers).set({ companyName: "companyName" }).where(eq(suppliers.id, 1));
```

```sql
update "suppliers"
set "company_name" = $1
where "suppliers"."id" = $2
-- params: ["companyName",1]
```

**Delete**: Deletes a row of suppliers, found by its primary key.

```ts
db.delete(suppliers).where(eq(suppliers.id, 1));
```

```sql
delete from "suppliers"
where "suppliers"."id" = $1
-- params: [1]
```

</details>

<a id="table-products"></a>

### products

![products and the tables it relates to](tables/products.svg)

| Column                                  | Type              | Constraints                                 | Default | Description |
| --------------------------------------- | ----------------- | ------------------------------------------- | ------- | ----------- |
| `id`                                    | `serial`          | PK NOT NULL                                 |         |             |
| `name`                                  | `text`            | NOT NULL                                    |         |             |
| `quantityPerUnit` (`quantity_per_unit`) | `varchar`         | NOT NULL                                    |         |             |
| `unitPrice` (`unit_price`)              | `doublePrecision` | NOT NULL                                    |         |             |
| `unitsInStock` (`units_in_stock`)       | `integer`         | NOT NULL                                    |         |             |
| `unitsOnOrder` (`units_on_order`)       | `integer`         | NOT NULL                                    |         |             |
| `reorderLevel` (`reorder_level`)        | `integer`         | NOT NULL                                    |         |             |
| `discontinued`                          | `integer`         | NOT NULL                                    |         |             |
| `supplierId` (`supplier_id`)            | `serial`          | NOT NULL → [suppliers.id](#table-suppliers) |         |             |

#### Definition

```sql
CREATE TABLE "products" (
	"id" serial PRIMARY KEY,
	"name" text NOT NULL,
	"quantity_per_unit" varchar NOT NULL,
	"unit_price" double precision NOT NULL,
	"units_in_stock" integer NOT NULL,
	"units_on_order" integer NOT NULL,
	"reorder_level" integer NOT NULL,
	"discontinued" integer NOT NULL,
	"supplier_id" serial
);

CREATE INDEX "supplier_idx" ON "products" ("supplier_id");

ALTER TABLE "products" ADD CONSTRAINT "products_supplier_id_suppliers_id_fkey" FOREIGN KEY ("supplier_id") REFERENCES "suppliers"("id") ON DELETE CASCADE;
```

#### Indexes

- INDEX `supplierId`

#### References

- `supplierId` → [suppliers.id](#table-suppliers) (many to one, on delete cascade)

#### Referenced by

- `id` → [details.productId](#table-details) (one to many, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of products.

```ts
db.select().from(products);
```

```sql
select "id", "name", "quantity_per_unit", "unit_price", "units_in_stock", "units_on_order", "reorder_level", "discontinued", "supplier_id"
from "products"
```

**Count**: How many rows products has.

```ts
db.select({ count: count() }).from(products);
```

```sql
select count(*)
from "products"
```

**Find many**: The rows of products, through the relational query API.

```ts
db.query.products.findMany();
```

```sql
select "d0"."id" as "id", "d0"."name" as "name", "d0"."quantity_per_unit" as "quantityPerUnit", "d0"."unit_price" as "unitPrice", "d0"."units_in_stock" as "unitsInStock", "d0"."units_on_order" as "unitsOnOrder", "d0"."reorder_level" as "reorderLevel", "d0"."discontinued" as "discontinued", "d0"."supplier_id" as "supplierId"
from "products" as "d0"
```

**Find by primary key**: One row of products, through the relational query API.

```ts
db.query.products.findFirst({ where: { id: 1 } });
```

```sql
select "d0"."id" as "id", "d0"."name" as "name", "d0"."quantity_per_unit" as "quantityPerUnit", "d0"."unit_price" as "unitPrice", "d0"."units_in_stock" as "unitsInStock", "d0"."units_on_order" as "unitsOnOrder", "d0"."reorder_level" as "reorderLevel", "d0"."discontinued" as "discontinued", "d0"."supplier_id" as "supplierId"
from "products" as "d0"
where "d0"."id" = $1
limit $2
-- params: [1,1]
```

**Insert**: Adds a row to products, giving the columns that have no default.

```ts
db.insert(products).values({
  name: "name",
  quantityPerUnit: "quantityPerUnit",
  unitPrice: 1,
  unitsInStock: 1,
  unitsOnOrder: 1,
  reorderLevel: 1,
  discontinued: 1,
});
```

```sql
insert into "products" ("id", "name", "quantity_per_unit", "unit_price", "units_in_stock", "units_on_order", "reorder_level", "discontinued", "supplier_id")
values (default, $1, $2, $3, $4, $5, $6, $7, default)
-- params: ["name","quantityPerUnit",1,1,1,1,1]
```

**Update**: Changes a row of products, found by its primary key.

```ts
db.update(products).set({ name: "name" }).where(eq(products.id, 1));
```

```sql
update "products"
set "name" = $1
where "products"."id" = $2
-- params: ["name",1]
```

**Delete**: Deletes a row of products, found by its primary key.

```ts
db.delete(products).where(eq(products.id, 1));
```

```sql
delete from "products"
where "products"."id" = $1
-- params: [1]
```

</details>

<a id="table-details"></a>

### details

Table `order_details`.

![details and the tables it relates to](tables/details.svg)

| Column                     | Type              | Constraints                               | Default | Description |
| -------------------------- | ----------------- | ----------------------------------------- | ------- | ----------- |
| `unitPrice` (`unit_price`) | `doublePrecision` | NOT NULL                                  |         |             |
| `quantity`                 | `integer`         | NOT NULL                                  |         |             |
| `discount`                 | `doublePrecision` | NOT NULL                                  |         |             |
| `orderId` (`order_id`)     | `integer`         | NOT NULL → [orders.id](#table-orders)     |         |             |
| `productId` (`product_id`) | `integer`         | NOT NULL → [products.id](#table-products) |         |             |

#### Definition

```sql
CREATE TABLE "order_details" (
	"unit_price" double precision NOT NULL,
	"quantity" integer NOT NULL,
	"discount" double precision NOT NULL,
	"order_id" integer NOT NULL,
	"product_id" integer NOT NULL
);

CREATE INDEX "order_id_idx" ON "order_details" ("order_id");

CREATE INDEX "product_id_idx" ON "order_details" ("product_id");

ALTER TABLE "order_details" ADD CONSTRAINT "order_details_order_id_orders_id_fkey" FOREIGN KEY ("order_id") REFERENCES "orders"("id") ON DELETE CASCADE;

ALTER TABLE "order_details" ADD CONSTRAINT "order_details_product_id_products_id_fkey" FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE;
```

#### Indexes

- INDEX `orderId`
- INDEX `productId`

#### References

- `orderId` → [orders.id](#table-orders) (many to one, on delete cascade)
- `productId` → [products.id](#table-products) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of details.

```ts
db.select().from(details);
```

```sql
select "unit_price", "quantity", "discount", "order_id", "product_id"
from "order_details"
```

**Count**: How many rows details has.

```ts
db.select({ count: count() }).from(details);
```

```sql
select count(*)
from "order_details"
```

**Find many**: The rows of details, through the relational query API.

```ts
db.query.details.findMany();
```

```sql
select "d0"."unit_price" as "unitPrice", "d0"."quantity" as "quantity", "d0"."discount" as "discount", "d0"."order_id" as "orderId", "d0"."product_id" as "productId"
from "order_details" as "d0"
```

**Insert**: Adds a row to details, giving the columns that have no default.

```ts
db.insert(details).values({ unitPrice: 1, quantity: 1, discount: 1, orderId: 1, productId: 1 });
```

```sql
insert into "order_details" ("unit_price", "quantity", "discount", "order_id", "product_id")
values ($1, $2, $3, $4, $5)
-- params: [1,1,1,1,1]
```

</details>
