# Schema docs

![ER diagram](er.svg)

<details>
<summary>The whole schema as SQL</summary>

```sql
CREATE TABLE "account" (
	"userId" text NOT NULL,
	"type" text NOT NULL,
	"provider" text,
	"providerAccountId" text,
	"refresh_token" text,
	"access_token" text,
	"expires_at" integer,
	"token_type" text,
	"scope" text,
	"id_token" text,
	"session_state" text,
	CONSTRAINT "account_pkey" PRIMARY KEY("provider","providerAccountId")
);

CREATE TABLE "authenticator" (
	"credentialID" text UNIQUE,
	"userId" text,
	"providerAccountId" text NOT NULL,
	"credentialPublicKey" text NOT NULL,
	"counter" integer NOT NULL,
	"credentialDeviceType" text NOT NULL,
	"credentialBackedUp" boolean NOT NULL,
	"transports" text,
	CONSTRAINT "authenticator_pkey" PRIMARY KEY("userId","credentialID")
);

CREATE TABLE "session" (
	"sessionToken" text PRIMARY KEY,
	"userId" text NOT NULL,
	"expires" timestamp NOT NULL
);

CREATE TABLE "user" (
	"id" text PRIMARY KEY,
	"name" text,
	"email" text UNIQUE,
	"emailVerified" timestamp,
	"image" text
);

CREATE TABLE "verificationToken" (
	"identifier" text,
	"token" text,
	"expires" timestamp NOT NULL,
	CONSTRAINT "verificationToken_pkey" PRIMARY KEY("identifier","token")
);

ALTER TABLE "account" ADD CONSTRAINT "account_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;

ALTER TABLE "authenticator" ADD CONSTRAINT "authenticator_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;

ALTER TABLE "session" ADD CONSTRAINT "session_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;
```

</details>

## Contents

| Table                                                     | Name in the database | Description |
| --------------------------------------------------------- | -------------------- | ----------- |
| [usersTable](#table-usersTable)                           | `user`               |             |
| [accountsTable](#table-accountsTable)                     | `account`            |             |
| [sessionsTable](#table-sessionsTable)                     | `session`            |             |
| [verificationTokensTable](#table-verificationTokensTable) | `verificationToken`  |             |
| [authenticatorsTable](#table-authenticatorsTable)         | `authenticator`      |             |

## Tables

<a id="table-usersTable"></a>

### usersTable

Table `user`.

![usersTable and the tables it relates to](tables/usersTable.svg)

| Column          | Type        | Constraints | Default         | Description |
| --------------- | ----------- | ----------- | --------------- | ----------- |
| `id`            | `text`      | PK NOT NULL | `.$defaultFn()` |             |
| `name`          | `text`      | NULL        |                 |             |
| `email`         | `text`      | UNIQUE NULL |                 |             |
| `emailVerified` | `timestamp` | NULL        |                 |             |
| `image`         | `text`      | NULL        |                 |             |

#### Definition

```sql
CREATE TABLE "user" (
	"id" text PRIMARY KEY,
	"name" text,
	"email" text UNIQUE,
	"emailVerified" timestamp,
	"image" text
);
```

#### Referenced by

- `id` → [accountsTable.userId](#table-accountsTable) (one to many, on delete cascade)
- `id` → [sessionsTable.userId](#table-sessionsTable) (one to many, on delete cascade)
- `id` → [authenticatorsTable.userId](#table-authenticatorsTable) (one to many, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of usersTable.

```ts
db.select().from(usersTable);
```

```sql
select "id", "name", "email", "emailVerified", "image"
from "user"
```

**Count**: How many rows usersTable has.

```ts
db.select({ count: count() }).from(usersTable);
```

```sql
select count(*)
from "user"
```

**Find many**: The rows of usersTable, through the relational query API.

```ts
db.query.usersTable.findMany();
```

```sql
select "d0"."id" as "id", "d0"."name" as "name", "d0"."email" as "email", "d0"."emailVerified" as "emailVerified", "d0"."image" as "image"
from "user" as "d0"
```

**Find by primary key**: One row of usersTable, through the relational query API.

```ts
db.query.usersTable.findFirst({ where: { id: "id" } });
```

```sql
select "d0"."id" as "id", "d0"."name" as "name", "d0"."email" as "email", "d0"."emailVerified" as "emailVerified", "d0"."image" as "image"
from "user" as "d0"
where "d0"."id" = $1
limit $2
-- params: ["id",1]
```

**Insert**: Adds a row to usersTable, giving the columns that have no default.

```ts
db.insert(usersTable).values({});
```

```sql
insert into "user" ("id", "name", "email", "emailVerified", "image")
values ($1, default, default, default, default)
-- params: ["00000000-0000-4000-8000-000000000000"]
```

**Update**: Changes a row of usersTable, found by its primary key.

```ts
db.update(usersTable).set({ name: "name" }).where(eq(usersTable.id, "id"));
```

```sql
update "user"
set "name" = $1
where "user"."id" = $2
-- params: ["name","id"]
```

**Delete**: Deletes a row of usersTable, found by its primary key.

```ts
db.delete(usersTable).where(eq(usersTable.id, "id"));
```

```sql
delete from "user"
where "user"."id" = $1
-- params: ["id"]
```

</details>

<a id="table-accountsTable"></a>

### accountsTable

Table `account`.

![accountsTable and the tables it relates to](tables/accountsTable.svg)

| Column              | Type      | Constraints                                   | Default | Description |
| ------------------- | --------- | --------------------------------------------- | ------- | ----------- |
| `userId`            | `text`    | NOT NULL → [usersTable.id](#table-usersTable) |         |             |
| `type`              | `text`    | NOT NULL                                      |         |             |
| `provider`          | `text`    | PK NOT NULL                                   |         |             |
| `providerAccountId` | `text`    | PK NOT NULL                                   |         |             |
| `refresh_token`     | `text`    | NULL                                          |         |             |
| `access_token`      | `text`    | NULL                                          |         |             |
| `expires_at`        | `integer` | NULL                                          |         |             |
| `token_type`        | `text`    | NULL                                          |         |             |
| `scope`             | `text`    | NULL                                          |         |             |
| `id_token`          | `text`    | NULL                                          |         |             |
| `session_state`     | `text`    | NULL                                          |         |             |

#### Definition

```sql
CREATE TABLE "account" (
	"userId" text NOT NULL,
	"type" text NOT NULL,
	"provider" text,
	"providerAccountId" text,
	"refresh_token" text,
	"access_token" text,
	"expires_at" integer,
	"token_type" text,
	"scope" text,
	"id_token" text,
	"session_state" text,
	CONSTRAINT "account_pkey" PRIMARY KEY("provider","providerAccountId")
);

ALTER TABLE "account" ADD CONSTRAINT "account_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;
```

#### Indexes

- PRIMARY KEY `provider`, `providerAccountId`

#### References

- `userId` → [usersTable.id](#table-usersTable) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of accountsTable.

```ts
db.select().from(accountsTable);
```

```sql
select "userId", "type", "provider", "providerAccountId", "refresh_token", "access_token", "expires_at", "token_type", "scope", "id_token", "session_state"
from "account"
```

**Count**: How many rows accountsTable has.

```ts
db.select({ count: count() }).from(accountsTable);
```

```sql
select count(*)
from "account"
```

**Find many**: The rows of accountsTable, through the relational query API.

```ts
db.query.accountsTable.findMany();
```

```sql
select "d0"."userId" as "userId", "d0"."type" as "type", "d0"."provider" as "provider", "d0"."providerAccountId" as "providerAccountId", "d0"."refresh_token" as "refresh_token", "d0"."access_token" as "access_token", "d0"."expires_at" as "expires_at", "d0"."token_type" as "token_type", "d0"."scope" as "scope", "d0"."id_token" as "id_token", "d0"."session_state" as "session_state"
from "account" as "d0"
```

**Find by primary key**: One row of accountsTable, through the relational query API.

```ts
db.query.accountsTable.findFirst({
  where: { provider: "provider", providerAccountId: "providerAccountId" },
});
```

```sql
select "d0"."userId" as "userId", "d0"."type" as "type", "d0"."provider" as "provider", "d0"."providerAccountId" as "providerAccountId", "d0"."refresh_token" as "refresh_token", "d0"."access_token" as "access_token", "d0"."expires_at" as "expires_at", "d0"."token_type" as "token_type", "d0"."scope" as "scope", "d0"."id_token" as "id_token", "d0"."session_state" as "session_state"
from "account" as "d0"
where (("d0"."provider" = $1) and ("d0"."providerAccountId" = $2))
limit $3
-- params: ["provider","providerAccountId",1]
```

**Insert**: Adds a row to accountsTable, giving the columns that have no default.

```ts
db.insert(accountsTable).values({
  userId: "userId",
  type: "type",
  provider: "provider",
  providerAccountId: "providerAccountId",
});
```

```sql
insert into "account" ("userId", "type", "provider", "providerAccountId", "refresh_token", "access_token", "expires_at", "token_type", "scope", "id_token", "session_state")
values ($1, $2, $3, $4, default, default, default, default, default, default, default)
-- params: ["userId","type","provider","providerAccountId"]
```

**Update**: Changes a row of accountsTable, found by its primary key.

```ts
db.update(accountsTable)
  .set({ userId: "userId" })
  .where(
    and(
      eq(accountsTable.provider, "provider"),
      eq(accountsTable.providerAccountId, "providerAccountId"),
    ),
  );
```

```sql
update "account"
set "userId" = $1
where (("account"."provider" = $2) and ("account"."providerAccountId" = $3))
-- params: ["userId","provider","providerAccountId"]
```

**Delete**: Deletes a row of accountsTable, found by its primary key.

```ts
db.delete(accountsTable).where(
  and(
    eq(accountsTable.provider, "provider"),
    eq(accountsTable.providerAccountId, "providerAccountId"),
  ),
);
```

```sql
delete from "account"
where (("account"."provider" = $1) and ("account"."providerAccountId" = $2))
-- params: ["provider","providerAccountId"]
```

</details>

<a id="table-sessionsTable"></a>

### sessionsTable

Table `session`.

![sessionsTable and the tables it relates to](tables/sessionsTable.svg)

| Column         | Type        | Constraints                                   | Default | Description |
| -------------- | ----------- | --------------------------------------------- | ------- | ----------- |
| `sessionToken` | `text`      | PK NOT NULL                                   |         |             |
| `userId`       | `text`      | NOT NULL → [usersTable.id](#table-usersTable) |         |             |
| `expires`      | `timestamp` | NOT NULL                                      |         |             |

#### Definition

```sql
CREATE TABLE "session" (
	"sessionToken" text PRIMARY KEY,
	"userId" text NOT NULL,
	"expires" timestamp NOT NULL
);

ALTER TABLE "session" ADD CONSTRAINT "session_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;
```

#### References

- `userId` → [usersTable.id](#table-usersTable) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of sessionsTable.

```ts
db.select().from(sessionsTable);
```

```sql
select "sessionToken", "userId", "expires"
from "session"
```

**Count**: How many rows sessionsTable has.

```ts
db.select({ count: count() }).from(sessionsTable);
```

```sql
select count(*)
from "session"
```

**Find many**: The rows of sessionsTable, through the relational query API.

```ts
db.query.sessionsTable.findMany();
```

```sql
select "d0"."sessionToken" as "sessionToken", "d0"."userId" as "userId", "d0"."expires" as "expires"
from "session" as "d0"
```

**Find by primary key**: One row of sessionsTable, through the relational query API.

```ts
db.query.sessionsTable.findFirst({ where: { sessionToken: "sessionToken" } });
```

```sql
select "d0"."sessionToken" as "sessionToken", "d0"."userId" as "userId", "d0"."expires" as "expires"
from "session" as "d0"
where "d0"."sessionToken" = $1
limit $2
-- params: ["sessionToken",1]
```

**Insert**: Adds a row to sessionsTable, giving the columns that have no default.

```ts
db.insert(sessionsTable).values({
  sessionToken: "sessionToken",
  userId: "userId",
  expires: new Date("2026-01-01T00:00:00.000Z"),
});
```

```sql
insert into "session" ("sessionToken", "userId", "expires")
values ($1, $2, $3)
-- params: ["sessionToken","userId","2026-01-01T00:00:00.000Z"]
```

**Update**: Changes a row of sessionsTable, found by its primary key.

```ts
db.update(sessionsTable)
  .set({ userId: "userId" })
  .where(eq(sessionsTable.sessionToken, "sessionToken"));
```

```sql
update "session"
set "userId" = $1
where "session"."sessionToken" = $2
-- params: ["userId","sessionToken"]
```

**Delete**: Deletes a row of sessionsTable, found by its primary key.

```ts
db.delete(sessionsTable).where(eq(sessionsTable.sessionToken, "sessionToken"));
```

```sql
delete from "session"
where "session"."sessionToken" = $1
-- params: ["sessionToken"]
```

</details>

<a id="table-verificationTokensTable"></a>

### verificationTokensTable

Table `verificationToken`.

![verificationTokensTable](tables/verificationTokensTable.svg)

| Column       | Type        | Constraints | Default | Description |
| ------------ | ----------- | ----------- | ------- | ----------- |
| `identifier` | `text`      | PK NOT NULL |         |             |
| `token`      | `text`      | PK NOT NULL |         |             |
| `expires`    | `timestamp` | NOT NULL    |         |             |

#### Definition

```sql
CREATE TABLE "verificationToken" (
	"identifier" text,
	"token" text,
	"expires" timestamp NOT NULL,
	CONSTRAINT "verificationToken_pkey" PRIMARY KEY("identifier","token")
);
```

#### Indexes

- PRIMARY KEY `identifier`, `token`

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of verificationTokensTable.

```ts
db.select().from(verificationTokensTable);
```

```sql
select "identifier", "token", "expires"
from "verificationToken"
```

**Count**: How many rows verificationTokensTable has.

```ts
db.select({ count: count() }).from(verificationTokensTable);
```

```sql
select count(*)
from "verificationToken"
```

**Find many**: The rows of verificationTokensTable, through the relational query API.

```ts
db.query.verificationTokensTable.findMany();
```

```sql
select "d0"."identifier" as "identifier", "d0"."token" as "token", "d0"."expires" as "expires"
from "verificationToken" as "d0"
```

**Find by primary key**: One row of verificationTokensTable, through the relational query API.

```ts
db.query.verificationTokensTable.findFirst({ where: { identifier: "identifier", token: "token" } });
```

```sql
select "d0"."identifier" as "identifier", "d0"."token" as "token", "d0"."expires" as "expires"
from "verificationToken" as "d0"
where (("d0"."identifier" = $1) and ("d0"."token" = $2))
limit $3
-- params: ["identifier","token",1]
```

**Insert**: Adds a row to verificationTokensTable, giving the columns that have no default.

```ts
db.insert(verificationTokensTable).values({
  identifier: "identifier",
  token: "token",
  expires: new Date("2026-01-01T00:00:00.000Z"),
});
```

```sql
insert into "verificationToken" ("identifier", "token", "expires")
values ($1, $2, $3)
-- params: ["identifier","token","2026-01-01T00:00:00.000Z"]
```

**Update**: Changes a row of verificationTokensTable, found by its primary key.

```ts
db.update(verificationTokensTable)
  .set({ expires: new Date("2026-01-01T00:00:00.000Z") })
  .where(
    and(
      eq(verificationTokensTable.identifier, "identifier"),
      eq(verificationTokensTable.token, "token"),
    ),
  );
```

```sql
update "verificationToken"
set "expires" = $1
where (("verificationToken"."identifier" = $2) and ("verificationToken"."token" = $3))
-- params: ["2026-01-01T00:00:00.000Z","identifier","token"]
```

**Delete**: Deletes a row of verificationTokensTable, found by its primary key.

```ts
db.delete(verificationTokensTable).where(
  and(
    eq(verificationTokensTable.identifier, "identifier"),
    eq(verificationTokensTable.token, "token"),
  ),
);
```

```sql
delete from "verificationToken"
where (("verificationToken"."identifier" = $1) and ("verificationToken"."token" = $2))
-- params: ["identifier","token"]
```

</details>

<a id="table-authenticatorsTable"></a>

### authenticatorsTable

Table `authenticator`.

![authenticatorsTable and the tables it relates to](tables/authenticatorsTable.svg)

| Column                 | Type      | Constraints                                      | Default | Description |
| ---------------------- | --------- | ------------------------------------------------ | ------- | ----------- |
| `credentialID`         | `text`    | PK UNIQUE NOT NULL                               |         |             |
| `userId`               | `text`    | PK NOT NULL → [usersTable.id](#table-usersTable) |         |             |
| `providerAccountId`    | `text`    | NOT NULL                                         |         |             |
| `credentialPublicKey`  | `text`    | NOT NULL                                         |         |             |
| `counter`              | `integer` | NOT NULL                                         |         |             |
| `credentialDeviceType` | `text`    | NOT NULL                                         |         |             |
| `credentialBackedUp`   | `boolean` | NOT NULL                                         |         |             |
| `transports`           | `text`    | NULL                                             |         |             |

#### Definition

```sql
CREATE TABLE "authenticator" (
	"credentialID" text UNIQUE,
	"userId" text,
	"providerAccountId" text NOT NULL,
	"credentialPublicKey" text NOT NULL,
	"counter" integer NOT NULL,
	"credentialDeviceType" text NOT NULL,
	"credentialBackedUp" boolean NOT NULL,
	"transports" text,
	CONSTRAINT "authenticator_pkey" PRIMARY KEY("userId","credentialID")
);

ALTER TABLE "authenticator" ADD CONSTRAINT "authenticator_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;
```

#### Indexes

- PRIMARY KEY `userId`, `credentialID`

#### References

- `userId` → [usersTable.id](#table-usersTable) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of authenticatorsTable.

```ts
db.select().from(authenticatorsTable);
```

```sql
select "credentialID", "userId", "providerAccountId", "credentialPublicKey", "counter", "credentialDeviceType", "credentialBackedUp", "transports"
from "authenticator"
```

**Count**: How many rows authenticatorsTable has.

```ts
db.select({ count: count() }).from(authenticatorsTable);
```

```sql
select count(*)
from "authenticator"
```

**Find many**: The rows of authenticatorsTable, through the relational query API.

```ts
db.query.authenticatorsTable.findMany();
```

```sql
select "d0"."credentialID" as "credentialID", "d0"."userId" as "userId", "d0"."providerAccountId" as "providerAccountId", "d0"."credentialPublicKey" as "credentialPublicKey", "d0"."counter" as "counter", "d0"."credentialDeviceType" as "credentialDeviceType", "d0"."credentialBackedUp" as "credentialBackedUp", "d0"."transports" as "transports"
from "authenticator" as "d0"
```

**Find by primary key**: One row of authenticatorsTable, through the relational query API.

```ts
db.query.authenticatorsTable.findFirst({
  where: { credentialID: "credentialID", userId: "userId" },
});
```

```sql
select "d0"."credentialID" as "credentialID", "d0"."userId" as "userId", "d0"."providerAccountId" as "providerAccountId", "d0"."credentialPublicKey" as "credentialPublicKey", "d0"."counter" as "counter", "d0"."credentialDeviceType" as "credentialDeviceType", "d0"."credentialBackedUp" as "credentialBackedUp", "d0"."transports" as "transports"
from "authenticator" as "d0"
where (("d0"."credentialID" = $1) and ("d0"."userId" = $2))
limit $3
-- params: ["credentialID","userId",1]
```

**Insert**: Adds a row to authenticatorsTable, giving the columns that have no default.

```ts
db.insert(authenticatorsTable).values({
  credentialID: "credentialID",
  userId: "userId",
  providerAccountId: "providerAccountId",
  credentialPublicKey: "credentialPublicKey",
  counter: 1,
  credentialDeviceType: "credentialDeviceType",
  credentialBackedUp: true,
});
```

```sql
insert into "authenticator" ("credentialID", "userId", "providerAccountId", "credentialPublicKey", "counter", "credentialDeviceType", "credentialBackedUp", "transports")
values ($1, $2, $3, $4, $5, $6, $7, default)
-- params: ["credentialID","userId","providerAccountId","credentialPublicKey",1,"credentialDeviceType",true]
```

**Update**: Changes a row of authenticatorsTable, found by its primary key.

```ts
db.update(authenticatorsTable)
  .set({ providerAccountId: "providerAccountId" })
  .where(
    and(
      eq(authenticatorsTable.credentialID, "credentialID"),
      eq(authenticatorsTable.userId, "userId"),
    ),
  );
```

```sql
update "authenticator"
set "providerAccountId" = $1
where (("authenticator"."credentialID" = $2) and ("authenticator"."userId" = $3))
-- params: ["providerAccountId","credentialID","userId"]
```

**Delete**: Deletes a row of authenticatorsTable, found by its primary key.

```ts
db.delete(authenticatorsTable).where(
  and(
    eq(authenticatorsTable.credentialID, "credentialID"),
    eq(authenticatorsTable.userId, "userId"),
  ),
);
```

```sql
delete from "authenticator"
where (("authenticator"."credentialID" = $1) and ("authenticator"."userId" = $2))
-- params: ["credentialID","userId"]
```

</details>
