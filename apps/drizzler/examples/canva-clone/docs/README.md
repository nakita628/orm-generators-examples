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

CREATE TABLE "project" (
	"id" text PRIMARY KEY,
	"name" text NOT NULL,
	"userId" text NOT NULL,
	"json" text NOT NULL,
	"height" integer NOT NULL,
	"width" integer NOT NULL,
	"thumbnailUrl" text,
	"isTemplate" boolean,
	"isPro" boolean,
	"createdAt" timestamp NOT NULL,
	"updatedAt" timestamp NOT NULL
);

CREATE TABLE "session" (
	"sessionToken" text PRIMARY KEY,
	"userId" text NOT NULL,
	"expires" timestamp NOT NULL
);

CREATE TABLE "subscription" (
	"id" text PRIMARY KEY,
	"userId" text NOT NULL,
	"subscriptionId" text NOT NULL,
	"customerId" text NOT NULL,
	"priceId" text NOT NULL,
	"status" text NOT NULL,
	"currentPeriodEnd" timestamp,
	"createdAt" timestamp NOT NULL,
	"updatedAt" timestamp NOT NULL
);

CREATE TABLE "user" (
	"id" text PRIMARY KEY,
	"name" text,
	"email" text NOT NULL,
	"emailVerified" timestamp,
	"image" text,
	"password" text
);

CREATE TABLE "verificationToken" (
	"identifier" text,
	"token" text,
	"expires" timestamp NOT NULL,
	CONSTRAINT "verificationToken_pkey" PRIMARY KEY("identifier","token")
);

ALTER TABLE "account" ADD CONSTRAINT "account_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;

ALTER TABLE "authenticator" ADD CONSTRAINT "authenticator_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;

ALTER TABLE "project" ADD CONSTRAINT "project_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;

ALTER TABLE "session" ADD CONSTRAINT "session_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;

ALTER TABLE "subscription" ADD CONSTRAINT "subscription_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;
```

</details>

## Contents

| Table                                           | Name in the database | Description |
| ----------------------------------------------- | -------------------- | ----------- |
| [users](#table-users)                           | `user`               |             |
| [accounts](#table-accounts)                     | `account`            |             |
| [sessions](#table-sessions)                     | `session`            |             |
| [verificationTokens](#table-verificationTokens) | `verificationToken`  |             |
| [authenticators](#table-authenticators)         | `authenticator`      |             |
| [projects](#table-projects)                     | `project`            |             |
| [subscriptions](#table-subscriptions)           | `subscription`       |             |

## Tables

<a id="table-users"></a>

### users

Table `user`.

![users and the tables it relates to](tables/users.svg)

| Column          | Type        | Constraints | Default         | Description |
| --------------- | ----------- | ----------- | --------------- | ----------- |
| `id`            | `text`      | PK NOT NULL | `.$defaultFn()` |             |
| `name`          | `text`      | NULL        |                 |             |
| `email`         | `text`      | NOT NULL    |                 |             |
| `emailVerified` | `timestamp` | NULL        |                 |             |
| `image`         | `text`      | NULL        |                 |             |
| `password`      | `text`      | NULL        |                 |             |

#### Definition

```sql
CREATE TABLE "user" (
	"id" text PRIMARY KEY,
	"name" text,
	"email" text NOT NULL,
	"emailVerified" timestamp,
	"image" text,
	"password" text
);
```

#### Referenced by

- `id` → [accounts.userId](#table-accounts) (one to many, on delete cascade)
- `id` → [sessions.userId](#table-sessions) (one to many, on delete cascade)
- `id` → [authenticators.userId](#table-authenticators) (one to many, on delete cascade)
- `id` → [projects.userId](#table-projects) (one to many, on delete cascade)
- `id` → [subscriptions.userId](#table-subscriptions) (one to many, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of users.

```ts
db.select().from(users);
```

```sql
select "id", "name", "email", "emailVerified", "image", "password"
from "user"
```

**Count**: How many rows users has.

```ts
db.select({ count: count() }).from(users);
```

```sql
select count(*)
from "user"
```

**Find many**: The rows of users, through the relational query API.

```ts
db.query.users.findMany();
```

```sql
select "d0"."id" as "id", "d0"."name" as "name", "d0"."email" as "email", "d0"."emailVerified" as "emailVerified", "d0"."image" as "image", "d0"."password" as "password"
from "user" as "d0"
```

**Find by primary key**: One row of users, through the relational query API.

```ts
db.query.users.findFirst({ where: { id: "id" } });
```

```sql
select "d0"."id" as "id", "d0"."name" as "name", "d0"."email" as "email", "d0"."emailVerified" as "emailVerified", "d0"."image" as "image", "d0"."password" as "password"
from "user" as "d0"
where "d0"."id" = $1
limit $2
-- params: ["id",1]
```

**Insert**: Adds a row to users, giving the columns that have no default.

```ts
db.insert(users).values({ email: "email" });
```

```sql
insert into "user" ("id", "name", "email", "emailVerified", "image", "password")
values ($1, default, $2, default, default, default)
-- params: ["00000000-0000-4000-8000-000000000000","email"]
```

**Update**: Changes a row of users, found by its primary key.

```ts
db.update(users).set({ name: "name" }).where(eq(users.id, "id"));
```

```sql
update "user"
set "name" = $1
where "user"."id" = $2
-- params: ["name","id"]
```

**Delete**: Deletes a row of users, found by its primary key.

```ts
db.delete(users).where(eq(users.id, "id"));
```

```sql
delete from "user"
where "user"."id" = $1
-- params: ["id"]
```

</details>

<a id="table-accounts"></a>

### accounts

Table `account`.

![accounts and the tables it relates to](tables/accounts.svg)

| Column              | Type      | Constraints                         | Default | Description |
| ------------------- | --------- | ----------------------------------- | ------- | ----------- |
| `userId`            | `text`    | NOT NULL → [users.id](#table-users) |         |             |
| `type`              | `text`    | NOT NULL                            |         |             |
| `provider`          | `text`    | PK NOT NULL                         |         |             |
| `providerAccountId` | `text`    | PK NOT NULL                         |         |             |
| `refresh_token`     | `text`    | NULL                                |         |             |
| `access_token`      | `text`    | NULL                                |         |             |
| `expires_at`        | `integer` | NULL                                |         |             |
| `token_type`        | `text`    | NULL                                |         |             |
| `scope`             | `text`    | NULL                                |         |             |
| `id_token`          | `text`    | NULL                                |         |             |
| `session_state`     | `text`    | NULL                                |         |             |

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

- `userId` → [users.id](#table-users) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of accounts.

```ts
db.select().from(accounts);
```

```sql
select "userId", "type", "provider", "providerAccountId", "refresh_token", "access_token", "expires_at", "token_type", "scope", "id_token", "session_state"
from "account"
```

**Count**: How many rows accounts has.

```ts
db.select({ count: count() }).from(accounts);
```

```sql
select count(*)
from "account"
```

**Find many**: The rows of accounts, through the relational query API.

```ts
db.query.accounts.findMany();
```

```sql
select "d0"."userId" as "userId", "d0"."type" as "type", "d0"."provider" as "provider", "d0"."providerAccountId" as "providerAccountId", "d0"."refresh_token" as "refresh_token", "d0"."access_token" as "access_token", "d0"."expires_at" as "expires_at", "d0"."token_type" as "token_type", "d0"."scope" as "scope", "d0"."id_token" as "id_token", "d0"."session_state" as "session_state"
from "account" as "d0"
```

**Find by primary key**: One row of accounts, through the relational query API.

```ts
db.query.accounts.findFirst({
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

**Insert**: Adds a row to accounts, giving the columns that have no default.

```ts
db.insert(accounts).values({
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

**Update**: Changes a row of accounts, found by its primary key.

```ts
db.update(accounts)
  .set({ userId: "userId" })
  .where(
    and(eq(accounts.provider, "provider"), eq(accounts.providerAccountId, "providerAccountId")),
  );
```

```sql
update "account"
set "userId" = $1
where (("account"."provider" = $2) and ("account"."providerAccountId" = $3))
-- params: ["userId","provider","providerAccountId"]
```

**Delete**: Deletes a row of accounts, found by its primary key.

```ts
db.delete(accounts).where(
  and(eq(accounts.provider, "provider"), eq(accounts.providerAccountId, "providerAccountId")),
);
```

```sql
delete from "account"
where (("account"."provider" = $1) and ("account"."providerAccountId" = $2))
-- params: ["provider","providerAccountId"]
```

</details>

<a id="table-sessions"></a>

### sessions

Table `session`.

![sessions and the tables it relates to](tables/sessions.svg)

| Column         | Type        | Constraints                         | Default | Description |
| -------------- | ----------- | ----------------------------------- | ------- | ----------- |
| `sessionToken` | `text`      | PK NOT NULL                         |         |             |
| `userId`       | `text`      | NOT NULL → [users.id](#table-users) |         |             |
| `expires`      | `timestamp` | NOT NULL                            |         |             |

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

- `userId` → [users.id](#table-users) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of sessions.

```ts
db.select().from(sessions);
```

```sql
select "sessionToken", "userId", "expires"
from "session"
```

**Count**: How many rows sessions has.

```ts
db.select({ count: count() }).from(sessions);
```

```sql
select count(*)
from "session"
```

**Find many**: The rows of sessions, through the relational query API.

```ts
db.query.sessions.findMany();
```

```sql
select "d0"."sessionToken" as "sessionToken", "d0"."userId" as "userId", "d0"."expires" as "expires"
from "session" as "d0"
```

**Find by primary key**: One row of sessions, through the relational query API.

```ts
db.query.sessions.findFirst({ where: { sessionToken: "sessionToken" } });
```

```sql
select "d0"."sessionToken" as "sessionToken", "d0"."userId" as "userId", "d0"."expires" as "expires"
from "session" as "d0"
where "d0"."sessionToken" = $1
limit $2
-- params: ["sessionToken",1]
```

**Insert**: Adds a row to sessions, giving the columns that have no default.

```ts
db.insert(sessions).values({
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

**Update**: Changes a row of sessions, found by its primary key.

```ts
db.update(sessions).set({ userId: "userId" }).where(eq(sessions.sessionToken, "sessionToken"));
```

```sql
update "session"
set "userId" = $1
where "session"."sessionToken" = $2
-- params: ["userId","sessionToken"]
```

**Delete**: Deletes a row of sessions, found by its primary key.

```ts
db.delete(sessions).where(eq(sessions.sessionToken, "sessionToken"));
```

```sql
delete from "session"
where "session"."sessionToken" = $1
-- params: ["sessionToken"]
```

</details>

<a id="table-verificationTokens"></a>

### verificationTokens

Table `verificationToken`.

![verificationTokens](tables/verificationTokens.svg)

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

**Select**: Every row of verificationTokens.

```ts
db.select().from(verificationTokens);
```

```sql
select "identifier", "token", "expires"
from "verificationToken"
```

**Count**: How many rows verificationTokens has.

```ts
db.select({ count: count() }).from(verificationTokens);
```

```sql
select count(*)
from "verificationToken"
```

**Find many**: The rows of verificationTokens, through the relational query API.

```ts
db.query.verificationTokens.findMany();
```

```sql
select "d0"."identifier" as "identifier", "d0"."token" as "token", "d0"."expires" as "expires"
from "verificationToken" as "d0"
```

**Find by primary key**: One row of verificationTokens, through the relational query API.

```ts
db.query.verificationTokens.findFirst({ where: { identifier: "identifier", token: "token" } });
```

```sql
select "d0"."identifier" as "identifier", "d0"."token" as "token", "d0"."expires" as "expires"
from "verificationToken" as "d0"
where (("d0"."identifier" = $1) and ("d0"."token" = $2))
limit $3
-- params: ["identifier","token",1]
```

**Insert**: Adds a row to verificationTokens, giving the columns that have no default.

```ts
db.insert(verificationTokens).values({
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

**Update**: Changes a row of verificationTokens, found by its primary key.

```ts
db.update(verificationTokens)
  .set({ expires: new Date("2026-01-01T00:00:00.000Z") })
  .where(
    and(eq(verificationTokens.identifier, "identifier"), eq(verificationTokens.token, "token")),
  );
```

```sql
update "verificationToken"
set "expires" = $1
where (("verificationToken"."identifier" = $2) and ("verificationToken"."token" = $3))
-- params: ["2026-01-01T00:00:00.000Z","identifier","token"]
```

**Delete**: Deletes a row of verificationTokens, found by its primary key.

```ts
db.delete(verificationTokens).where(
  and(eq(verificationTokens.identifier, "identifier"), eq(verificationTokens.token, "token")),
);
```

```sql
delete from "verificationToken"
where (("verificationToken"."identifier" = $1) and ("verificationToken"."token" = $2))
-- params: ["identifier","token"]
```

</details>

<a id="table-authenticators"></a>

### authenticators

Table `authenticator`.

![authenticators and the tables it relates to](tables/authenticators.svg)

| Column                 | Type      | Constraints                            | Default | Description |
| ---------------------- | --------- | -------------------------------------- | ------- | ----------- |
| `credentialID`         | `text`    | PK UNIQUE NOT NULL                     |         |             |
| `userId`               | `text`    | PK NOT NULL → [users.id](#table-users) |         |             |
| `providerAccountId`    | `text`    | NOT NULL                               |         |             |
| `credentialPublicKey`  | `text`    | NOT NULL                               |         |             |
| `counter`              | `integer` | NOT NULL                               |         |             |
| `credentialDeviceType` | `text`    | NOT NULL                               |         |             |
| `credentialBackedUp`   | `boolean` | NOT NULL                               |         |             |
| `transports`           | `text`    | NULL                                   |         |             |

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

- `userId` → [users.id](#table-users) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of authenticators.

```ts
db.select().from(authenticators);
```

```sql
select "credentialID", "userId", "providerAccountId", "credentialPublicKey", "counter", "credentialDeviceType", "credentialBackedUp", "transports"
from "authenticator"
```

**Count**: How many rows authenticators has.

```ts
db.select({ count: count() }).from(authenticators);
```

```sql
select count(*)
from "authenticator"
```

**Find many**: The rows of authenticators, through the relational query API.

```ts
db.query.authenticators.findMany();
```

```sql
select "d0"."credentialID" as "credentialID", "d0"."userId" as "userId", "d0"."providerAccountId" as "providerAccountId", "d0"."credentialPublicKey" as "credentialPublicKey", "d0"."counter" as "counter", "d0"."credentialDeviceType" as "credentialDeviceType", "d0"."credentialBackedUp" as "credentialBackedUp", "d0"."transports" as "transports"
from "authenticator" as "d0"
```

**Find by primary key**: One row of authenticators, through the relational query API.

```ts
db.query.authenticators.findFirst({ where: { credentialID: "credentialID", userId: "userId" } });
```

```sql
select "d0"."credentialID" as "credentialID", "d0"."userId" as "userId", "d0"."providerAccountId" as "providerAccountId", "d0"."credentialPublicKey" as "credentialPublicKey", "d0"."counter" as "counter", "d0"."credentialDeviceType" as "credentialDeviceType", "d0"."credentialBackedUp" as "credentialBackedUp", "d0"."transports" as "transports"
from "authenticator" as "d0"
where (("d0"."credentialID" = $1) and ("d0"."userId" = $2))
limit $3
-- params: ["credentialID","userId",1]
```

**Insert**: Adds a row to authenticators, giving the columns that have no default.

```ts
db.insert(authenticators).values({
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

**Update**: Changes a row of authenticators, found by its primary key.

```ts
db.update(authenticators)
  .set({ providerAccountId: "providerAccountId" })
  .where(and(eq(authenticators.credentialID, "credentialID"), eq(authenticators.userId, "userId")));
```

```sql
update "authenticator"
set "providerAccountId" = $1
where (("authenticator"."credentialID" = $2) and ("authenticator"."userId" = $3))
-- params: ["providerAccountId","credentialID","userId"]
```

**Delete**: Deletes a row of authenticators, found by its primary key.

```ts
db.delete(authenticators).where(
  and(eq(authenticators.credentialID, "credentialID"), eq(authenticators.userId, "userId")),
);
```

```sql
delete from "authenticator"
where (("authenticator"."credentialID" = $1) and ("authenticator"."userId" = $2))
-- params: ["credentialID","userId"]
```

</details>

<a id="table-projects"></a>

### projects

Table `project`.

![projects and the tables it relates to](tables/projects.svg)

| Column         | Type        | Constraints                         | Default         | Description |
| -------------- | ----------- | ----------------------------------- | --------------- | ----------- |
| `id`           | `text`      | PK NOT NULL                         | `.$defaultFn()` |             |
| `name`         | `text`      | NOT NULL                            |                 |             |
| `userId`       | `text`      | NOT NULL → [users.id](#table-users) |                 |             |
| `json`         | `text`      | NOT NULL                            |                 |             |
| `height`       | `integer`   | NOT NULL                            |                 |             |
| `width`        | `integer`   | NOT NULL                            |                 |             |
| `thumbnailUrl` | `text`      | NULL                                |                 |             |
| `isTemplate`   | `boolean`   | NULL                                |                 |             |
| `isPro`        | `boolean`   | NULL                                |                 |             |
| `createdAt`    | `timestamp` | NOT NULL                            |                 |             |
| `updatedAt`    | `timestamp` | NOT NULL                            |                 |             |

#### Definition

```sql
CREATE TABLE "project" (
	"id" text PRIMARY KEY,
	"name" text NOT NULL,
	"userId" text NOT NULL,
	"json" text NOT NULL,
	"height" integer NOT NULL,
	"width" integer NOT NULL,
	"thumbnailUrl" text,
	"isTemplate" boolean,
	"isPro" boolean,
	"createdAt" timestamp NOT NULL,
	"updatedAt" timestamp NOT NULL
);

ALTER TABLE "project" ADD CONSTRAINT "project_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;
```

#### References

- `userId` → [users.id](#table-users) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of projects.

```ts
db.select().from(projects);
```

```sql
select "id", "name", "userId", "json", "height", "width", "thumbnailUrl", "isTemplate", "isPro", "createdAt", "updatedAt"
from "project"
```

**Count**: How many rows projects has.

```ts
db.select({ count: count() }).from(projects);
```

```sql
select count(*)
from "project"
```

**Find many**: The rows of projects, through the relational query API.

```ts
db.query.projects.findMany();
```

```sql
select "d0"."id" as "id", "d0"."name" as "name", "d0"."userId" as "userId", "d0"."json" as "json", "d0"."height" as "height", "d0"."width" as "width", "d0"."thumbnailUrl" as "thumbnailUrl", "d0"."isTemplate" as "isTemplate", "d0"."isPro" as "isPro", "d0"."createdAt" as "createdAt", "d0"."updatedAt" as "updatedAt"
from "project" as "d0"
```

**Find by primary key**: One row of projects, through the relational query API.

```ts
db.query.projects.findFirst({ where: { id: "id" } });
```

```sql
select "d0"."id" as "id", "d0"."name" as "name", "d0"."userId" as "userId", "d0"."json" as "json", "d0"."height" as "height", "d0"."width" as "width", "d0"."thumbnailUrl" as "thumbnailUrl", "d0"."isTemplate" as "isTemplate", "d0"."isPro" as "isPro", "d0"."createdAt" as "createdAt", "d0"."updatedAt" as "updatedAt"
from "project" as "d0"
where "d0"."id" = $1
limit $2
-- params: ["id",1]
```

**Insert**: Adds a row to projects, giving the columns that have no default.

```ts
db.insert(projects).values({
  name: "name",
  userId: "userId",
  json: "json",
  height: 1,
  width: 1,
  createdAt: new Date("2026-01-01T00:00:00.000Z"),
  updatedAt: new Date("2026-01-01T00:00:00.000Z"),
});
```

```sql
insert into "project" ("id", "name", "userId", "json", "height", "width", "thumbnailUrl", "isTemplate", "isPro", "createdAt", "updatedAt")
values ($1, $2, $3, $4, $5, $6, default, default, default, $7, $8)
-- params: ["00000000-0000-4000-8000-000000000000","name","userId","json",1,1,"2026-01-01T00:00:00.000Z","2026-01-01T00:00:00.000Z"]
```

**Update**: Changes a row of projects, found by its primary key.

```ts
db.update(projects).set({ name: "name" }).where(eq(projects.id, "id"));
```

```sql
update "project"
set "name" = $1
where "project"."id" = $2
-- params: ["name","id"]
```

**Delete**: Deletes a row of projects, found by its primary key.

```ts
db.delete(projects).where(eq(projects.id, "id"));
```

```sql
delete from "project"
where "project"."id" = $1
-- params: ["id"]
```

</details>

<a id="table-subscriptions"></a>

### subscriptions

Table `subscription`.

![subscriptions and the tables it relates to](tables/subscriptions.svg)

| Column             | Type        | Constraints                         | Default         | Description |
| ------------------ | ----------- | ----------------------------------- | --------------- | ----------- |
| `id`               | `text`      | PK NOT NULL                         | `.$defaultFn()` |             |
| `userId`           | `text`      | NOT NULL → [users.id](#table-users) |                 |             |
| `subscriptionId`   | `text`      | NOT NULL                            |                 |             |
| `customerId`       | `text`      | NOT NULL                            |                 |             |
| `priceId`          | `text`      | NOT NULL                            |                 |             |
| `status`           | `text`      | NOT NULL                            |                 |             |
| `currentPeriodEnd` | `timestamp` | NULL                                |                 |             |
| `createdAt`        | `timestamp` | NOT NULL                            |                 |             |
| `updatedAt`        | `timestamp` | NOT NULL                            |                 |             |

#### Definition

```sql
CREATE TABLE "subscription" (
	"id" text PRIMARY KEY,
	"userId" text NOT NULL,
	"subscriptionId" text NOT NULL,
	"customerId" text NOT NULL,
	"priceId" text NOT NULL,
	"status" text NOT NULL,
	"currentPeriodEnd" timestamp,
	"createdAt" timestamp NOT NULL,
	"updatedAt" timestamp NOT NULL
);

ALTER TABLE "subscription" ADD CONSTRAINT "subscription_userId_user_id_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE;
```

#### References

- `userId` → [users.id](#table-users) (many to one, on delete cascade)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of subscriptions.

```ts
db.select().from(subscriptions);
```

```sql
select "id", "userId", "subscriptionId", "customerId", "priceId", "status", "currentPeriodEnd", "createdAt", "updatedAt"
from "subscription"
```

**Count**: How many rows subscriptions has.

```ts
db.select({ count: count() }).from(subscriptions);
```

```sql
select count(*)
from "subscription"
```

**Find many**: The rows of subscriptions, through the relational query API.

```ts
db.query.subscriptions.findMany();
```

```sql
select "d0"."id" as "id", "d0"."userId" as "userId", "d0"."subscriptionId" as "subscriptionId", "d0"."customerId" as "customerId", "d0"."priceId" as "priceId", "d0"."status" as "status", "d0"."currentPeriodEnd" as "currentPeriodEnd", "d0"."createdAt" as "createdAt", "d0"."updatedAt" as "updatedAt"
from "subscription" as "d0"
```

**Find by primary key**: One row of subscriptions, through the relational query API.

```ts
db.query.subscriptions.findFirst({ where: { id: "id" } });
```

```sql
select "d0"."id" as "id", "d0"."userId" as "userId", "d0"."subscriptionId" as "subscriptionId", "d0"."customerId" as "customerId", "d0"."priceId" as "priceId", "d0"."status" as "status", "d0"."currentPeriodEnd" as "currentPeriodEnd", "d0"."createdAt" as "createdAt", "d0"."updatedAt" as "updatedAt"
from "subscription" as "d0"
where "d0"."id" = $1
limit $2
-- params: ["id",1]
```

**Insert**: Adds a row to subscriptions, giving the columns that have no default.

```ts
db.insert(subscriptions).values({
  userId: "userId",
  subscriptionId: "subscriptionId",
  customerId: "customerId",
  priceId: "priceId",
  status: "status",
  createdAt: new Date("2026-01-01T00:00:00.000Z"),
  updatedAt: new Date("2026-01-01T00:00:00.000Z"),
});
```

```sql
insert into "subscription" ("id", "userId", "subscriptionId", "customerId", "priceId", "status", "currentPeriodEnd", "createdAt", "updatedAt")
values ($1, $2, $3, $4, $5, $6, default, $7, $8)
-- params: ["00000000-0000-4000-8000-000000000000","userId","subscriptionId","customerId","priceId","status","2026-01-01T00:00:00.000Z","2026-01-01T00:00:00.000Z"]
```

**Update**: Changes a row of subscriptions, found by its primary key.

```ts
db.update(subscriptions).set({ userId: "userId" }).where(eq(subscriptions.id, "id"));
```

```sql
update "subscription"
set "userId" = $1
where "subscription"."id" = $2
-- params: ["userId","id"]
```

**Delete**: Deletes a row of subscriptions, found by its primary key.

```ts
db.delete(subscriptions).where(eq(subscriptions.id, "id"));
```

```sql
delete from "subscription"
where "subscription"."id" = $1
-- params: ["id"]
```

</details>
