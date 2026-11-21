# Schema docs

![ER diagram](er.svg)

<details>
<summary>The whole schema as SQL</summary>

```sql
CREATE TABLE `project1_account` (
	`userId` varchar(255) NOT NULL,
	`type` varchar(255) NOT NULL,
	`provider` varchar(255) NOT NULL,
	`providerAccountId` varchar(255) NOT NULL,
	`refresh_token` text,
	`access_token` text,
	`expires_at` int,
	`token_type` varchar(255),
	`scope` varchar(255),
	`id_token` text,
	`session_state` varchar(255),
	CONSTRAINT PRIMARY KEY(`provider`,`providerAccountId`)
);

CREATE TABLE `project1_post` (
	`id` bigint AUTO_INCREMENT PRIMARY KEY,
	`name` varchar(256),
	`createdById` varchar(255) NOT NULL,
	`createdAt` timestamp NOT NULL,
	`updatedAt` timestamp ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `project1_session` (
	`sessionToken` varchar(255) PRIMARY KEY,
	`userId` varchar(255) NOT NULL,
	`expires` timestamp NOT NULL
);

CREATE TABLE `project1_user` (
	`id` varchar(255) PRIMARY KEY,
	`name` varchar(255),
	`email` varchar(255) NOT NULL,
	`emailVerified` timestamp(3) DEFAULT (CURRENT_TIMESTAMP(3)),
	`image` varchar(255)
);

CREATE TABLE `project1_verification_token` (
	`identifier` varchar(255) NOT NULL,
	`token` varchar(255) NOT NULL,
	`expires` timestamp NOT NULL,
	CONSTRAINT PRIMARY KEY(`identifier`,`token`)
);

CREATE INDEX `account_user_id_idx` ON `project1_account` (`userId`);

CREATE INDEX `created_by_idx` ON `project1_post` (`createdById`);

CREATE INDEX `name_idx` ON `project1_post` (`name`);

CREATE INDEX `session_user_id_idx` ON `project1_session` (`userId`);

ALTER TABLE `project1_account` ADD CONSTRAINT `project1_account_userId_project1_user_id_fkey` FOREIGN KEY (`userId`) REFERENCES `project1_user`(`id`);

ALTER TABLE `project1_post` ADD CONSTRAINT `project1_post_createdById_project1_user_id_fkey` FOREIGN KEY (`createdById`) REFERENCES `project1_user`(`id`);

ALTER TABLE `project1_session` ADD CONSTRAINT `project1_session_userId_project1_user_id_fkey` FOREIGN KEY (`userId`) REFERENCES `project1_user`(`id`);
```

</details>

## Contents

| Table                                           | Name in the database          | Description |
| ----------------------------------------------- | ----------------------------- | ----------- |
| [posts](#table-posts)                           | `project1_post`               |             |
| [users](#table-users)                           | `project1_user`               |             |
| [accounts](#table-accounts)                     | `project1_account`            |             |
| [sessions](#table-sessions)                     | `project1_session`            |             |
| [verificationTokens](#table-verificationTokens) | `project1_verification_token` |             |

## Tables

<a id="table-posts"></a>

### posts

Table `project1_post`.

![posts and the tables it relates to](tables/posts.svg)

| Column        | Type        | Constraints                         | Default         | Description |
| ------------- | ----------- | ----------------------------------- | --------------- | ----------- |
| `id`          | `bigint`    | PK NOT NULL                         |                 |             |
| `name`        | `varchar`   | NULL                                |                 |             |
| `createdById` | `varchar`   | NOT NULL → [users.id](#table-users) |                 |             |
| `createdAt`   | `timestamp` | NOT NULL                            | `.$defaultFn()` |             |
| `updatedAt`   | `timestamp` | NULL                                |                 |             |

#### Definition

```sql
CREATE TABLE `project1_post` (
	`id` bigint AUTO_INCREMENT PRIMARY KEY,
	`name` varchar(256),
	`createdById` varchar(255) NOT NULL,
	`createdAt` timestamp NOT NULL,
	`updatedAt` timestamp ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX `created_by_idx` ON `project1_post` (`createdById`);

CREATE INDEX `name_idx` ON `project1_post` (`name`);

ALTER TABLE `project1_post` ADD CONSTRAINT `project1_post_createdById_project1_user_id_fkey` FOREIGN KEY (`createdById`) REFERENCES `project1_user`(`id`);
```

#### Indexes

- INDEX `createdById`
- INDEX `name`

#### References

- `createdById` → [users.id](#table-users) (many to one)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of posts.

```ts
db.select().from(posts);
```

```sql
select `id`, `name`, `createdById`, `createdAt`, `updatedAt`
from `project1_post`
```

**Count**: How many rows posts has.

```ts
db.select({ count: count() }).from(posts);
```

```sql
select count(*)
from `project1_post`
```

**Find many**: The rows of posts, through the relational query API.

```ts
db.query.posts.findMany();
```

```sql
select `d0`.`id` as `id`, `d0`.`name` as `name`, `d0`.`createdById` as `createdById`, `d0`.`createdAt` as `createdAt`, `d0`.`updatedAt` as `updatedAt`
from `project1_post` as `d0`
```

**Find by primary key**: One row of posts, through the relational query API.

```ts
db.query.posts.findFirst({ where: { id: 1 } });
```

```sql
select `d0`.`id` as `id`, `d0`.`name` as `name`, `d0`.`createdById` as `createdById`, `d0`.`createdAt` as `createdAt`, `d0`.`updatedAt` as `updatedAt`
from `project1_post` as `d0`
where `d0`.`id` = ?
limit ?
-- params: [1,1]
```

**Insert**: Adds a row to posts, giving the columns that have no default.

```ts
db.insert(posts).values({ id: 1, createdById: "createdById" });
```

```sql
insert into `project1_post` (`id`, `name`, `createdById`, `createdAt`, `updatedAt`)
values (?, default, ?, ?, default)
-- params: [1,"createdById","2026-01-01 00:00:00.000"]
```

**Update**: Changes a row of posts, found by its primary key.

```ts
db.update(posts).set({ name: "name" }).where(eq(posts.id, 1));
```

```sql
update `project1_post`
set `name` = ?
where `project1_post`.`id` = ?
-- params: ["name",1]
```

**Delete**: Deletes a row of posts, found by its primary key.

```ts
db.delete(posts).where(eq(posts.id, 1));
```

```sql
delete from `project1_post`
where `project1_post`.`id` = ?
-- params: [1]
```

</details>

<a id="table-users"></a>

### users

Table `project1_user`.

![users and the tables it relates to](tables/users.svg)

| Column          | Type        | Constraints | Default                                 | Description |
| --------------- | ----------- | ----------- | --------------------------------------- | ----------- |
| `id`            | `varchar`   | PK NOT NULL | `.$defaultFn()`                         |             |
| `name`          | `varchar`   | NULL        |                                         |             |
| `email`         | `varchar`   | NOT NULL    |                                         |             |
| `emailVerified` | `timestamp` | NULL        | ``.default(sql`CURRENT_TIMESTAMP(3)`)`` |             |
| `image`         | `varchar`   | NULL        |                                         |             |

#### Definition

```sql
CREATE TABLE `project1_user` (
	`id` varchar(255) PRIMARY KEY,
	`name` varchar(255),
	`email` varchar(255) NOT NULL,
	`emailVerified` timestamp(3) DEFAULT (CURRENT_TIMESTAMP(3)),
	`image` varchar(255)
);
```

#### Referenced by

- `id` → [posts.createdById](#table-posts) (one to many)
- `id` → [accounts.userId](#table-accounts) (one to many)
- `id` → [sessions.userId](#table-sessions) (one to many)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of users.

```ts
db.select().from(users);
```

```sql
select `id`, `name`, `email`, `emailVerified`, `image`
from `project1_user`
```

**Count**: How many rows users has.

```ts
db.select({ count: count() }).from(users);
```

```sql
select count(*)
from `project1_user`
```

**Find many**: The rows of users, through the relational query API.

```ts
db.query.users.findMany();
```

```sql
select `d0`.`id` as `id`, `d0`.`name` as `name`, `d0`.`email` as `email`, `d0`.`emailVerified` as `emailVerified`, `d0`.`image` as `image`
from `project1_user` as `d0`
```

**Find by primary key**: One row of users, through the relational query API.

```ts
db.query.users.findFirst({ where: { id: "id" } });
```

```sql
select `d0`.`id` as `id`, `d0`.`name` as `name`, `d0`.`email` as `email`, `d0`.`emailVerified` as `emailVerified`, `d0`.`image` as `image`
from `project1_user` as `d0`
where `d0`.`id` = ?
limit ?
-- params: ["id",1]
```

**Insert**: Adds a row to users, giving the columns that have no default.

```ts
db.insert(users).values({ email: "email" });
```

```sql
insert into `project1_user` (`id`, `name`, `email`, `emailVerified`, `image`)
values (?, default, ?, default, default)
-- params: ["00000000-0000-4000-8000-000000000000","email"]
```

**Update**: Changes a row of users, found by its primary key.

```ts
db.update(users).set({ name: "name" }).where(eq(users.id, "id"));
```

```sql
update `project1_user`
set `name` = ?
where `project1_user`.`id` = ?
-- params: ["name","id"]
```

**Delete**: Deletes a row of users, found by its primary key.

```ts
db.delete(users).where(eq(users.id, "id"));
```

```sql
delete from `project1_user`
where `project1_user`.`id` = ?
-- params: ["id"]
```

</details>

<a id="table-accounts"></a>

### accounts

Table `project1_account`.

![accounts and the tables it relates to](tables/accounts.svg)

| Column              | Type      | Constraints                         | Default | Description |
| ------------------- | --------- | ----------------------------------- | ------- | ----------- |
| `userId`            | `varchar` | NOT NULL → [users.id](#table-users) |         |             |
| `type`              | `varchar` | NOT NULL                            |         |             |
| `provider`          | `varchar` | PK NOT NULL                         |         |             |
| `providerAccountId` | `varchar` | PK NOT NULL                         |         |             |
| `refresh_token`     | `text`    | NULL                                |         |             |
| `access_token`      | `text`    | NULL                                |         |             |
| `expires_at`        | `int`     | NULL                                |         |             |
| `token_type`        | `varchar` | NULL                                |         |             |
| `scope`             | `varchar` | NULL                                |         |             |
| `id_token`          | `text`    | NULL                                |         |             |
| `session_state`     | `varchar` | NULL                                |         |             |

#### Definition

```sql
CREATE TABLE `project1_account` (
	`userId` varchar(255) NOT NULL,
	`type` varchar(255) NOT NULL,
	`provider` varchar(255) NOT NULL,
	`providerAccountId` varchar(255) NOT NULL,
	`refresh_token` text,
	`access_token` text,
	`expires_at` int,
	`token_type` varchar(255),
	`scope` varchar(255),
	`id_token` text,
	`session_state` varchar(255),
	CONSTRAINT PRIMARY KEY(`provider`,`providerAccountId`)
);

CREATE INDEX `account_user_id_idx` ON `project1_account` (`userId`);

ALTER TABLE `project1_account` ADD CONSTRAINT `project1_account_userId_project1_user_id_fkey` FOREIGN KEY (`userId`) REFERENCES `project1_user`(`id`);
```

#### Indexes

- PRIMARY KEY `provider`, `providerAccountId`
- INDEX `userId`

#### References

- `userId` → [users.id](#table-users) (many to one)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of accounts.

```ts
db.select().from(accounts);
```

```sql
select `userId`, `type`, `provider`, `providerAccountId`, `refresh_token`, `access_token`, `expires_at`, `token_type`, `scope`, `id_token`, `session_state`
from `project1_account`
```

**Count**: How many rows accounts has.

```ts
db.select({ count: count() }).from(accounts);
```

```sql
select count(*)
from `project1_account`
```

**Find many**: The rows of accounts, through the relational query API.

```ts
db.query.accounts.findMany();
```

```sql
select `d0`.`userId` as `userId`, `d0`.`type` as `type`, `d0`.`provider` as `provider`, `d0`.`providerAccountId` as `providerAccountId`, `d0`.`refresh_token` as `refresh_token`, `d0`.`access_token` as `access_token`, `d0`.`expires_at` as `expires_at`, `d0`.`token_type` as `token_type`, `d0`.`scope` as `scope`, `d0`.`id_token` as `id_token`, `d0`.`session_state` as `session_state`
from `project1_account` as `d0`
```

**Find by primary key**: One row of accounts, through the relational query API.

```ts
db.query.accounts.findFirst({
  where: { provider: "provider", providerAccountId: "providerAccountId" },
});
```

```sql
select `d0`.`userId` as `userId`, `d0`.`type` as `type`, `d0`.`provider` as `provider`, `d0`.`providerAccountId` as `providerAccountId`, `d0`.`refresh_token` as `refresh_token`, `d0`.`access_token` as `access_token`, `d0`.`expires_at` as `expires_at`, `d0`.`token_type` as `token_type`, `d0`.`scope` as `scope`, `d0`.`id_token` as `id_token`, `d0`.`session_state` as `session_state`
from `project1_account` as `d0`
where ((`d0`.`provider` = ?) and (`d0`.`providerAccountId` = ?))
limit ?
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
insert into `project1_account` (`userId`, `type`, `provider`, `providerAccountId`, `refresh_token`, `access_token`, `expires_at`, `token_type`, `scope`, `id_token`, `session_state`)
values (?, ?, ?, ?, default, default, default, default, default, default, default)
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
update `project1_account`
set `userId` = ?
where ((`project1_account`.`provider` = ?) and (`project1_account`.`providerAccountId` = ?))
-- params: ["userId","provider","providerAccountId"]
```

**Delete**: Deletes a row of accounts, found by its primary key.

```ts
db.delete(accounts).where(
  and(eq(accounts.provider, "provider"), eq(accounts.providerAccountId, "providerAccountId")),
);
```

```sql
delete from `project1_account`
where ((`project1_account`.`provider` = ?) and (`project1_account`.`providerAccountId` = ?))
-- params: ["provider","providerAccountId"]
```

</details>

<a id="table-sessions"></a>

### sessions

Table `project1_session`.

![sessions and the tables it relates to](tables/sessions.svg)

| Column         | Type        | Constraints                         | Default | Description |
| -------------- | ----------- | ----------------------------------- | ------- | ----------- |
| `sessionToken` | `varchar`   | PK NOT NULL                         |         |             |
| `userId`       | `varchar`   | NOT NULL → [users.id](#table-users) |         |             |
| `expires`      | `timestamp` | NOT NULL                            |         |             |

#### Definition

```sql
CREATE TABLE `project1_session` (
	`sessionToken` varchar(255) PRIMARY KEY,
	`userId` varchar(255) NOT NULL,
	`expires` timestamp NOT NULL
);

CREATE INDEX `session_user_id_idx` ON `project1_session` (`userId`);

ALTER TABLE `project1_session` ADD CONSTRAINT `project1_session_userId_project1_user_id_fkey` FOREIGN KEY (`userId`) REFERENCES `project1_user`(`id`);
```

#### Indexes

- INDEX `userId`

#### References

- `userId` → [users.id](#table-users) (many to one)

<details>
<summary>Queries and their SQL</summary>

**Select**: Every row of sessions.

```ts
db.select().from(sessions);
```

```sql
select `sessionToken`, `userId`, `expires`
from `project1_session`
```

**Count**: How many rows sessions has.

```ts
db.select({ count: count() }).from(sessions);
```

```sql
select count(*)
from `project1_session`
```

**Find many**: The rows of sessions, through the relational query API.

```ts
db.query.sessions.findMany();
```

```sql
select `d0`.`sessionToken` as `sessionToken`, `d0`.`userId` as `userId`, `d0`.`expires` as `expires`
from `project1_session` as `d0`
```

**Find by primary key**: One row of sessions, through the relational query API.

```ts
db.query.sessions.findFirst({ where: { sessionToken: "sessionToken" } });
```

```sql
select `d0`.`sessionToken` as `sessionToken`, `d0`.`userId` as `userId`, `d0`.`expires` as `expires`
from `project1_session` as `d0`
where `d0`.`sessionToken` = ?
limit ?
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
insert into `project1_session` (`sessionToken`, `userId`, `expires`)
values (?, ?, ?)
-- params: ["sessionToken","userId","2026-01-01 00:00:00.000"]
```

**Update**: Changes a row of sessions, found by its primary key.

```ts
db.update(sessions).set({ userId: "userId" }).where(eq(sessions.sessionToken, "sessionToken"));
```

```sql
update `project1_session`
set `userId` = ?
where `project1_session`.`sessionToken` = ?
-- params: ["userId","sessionToken"]
```

**Delete**: Deletes a row of sessions, found by its primary key.

```ts
db.delete(sessions).where(eq(sessions.sessionToken, "sessionToken"));
```

```sql
delete from `project1_session`
where `project1_session`.`sessionToken` = ?
-- params: ["sessionToken"]
```

</details>

<a id="table-verificationTokens"></a>

### verificationTokens

Table `project1_verification_token`.

![verificationTokens](tables/verificationTokens.svg)

| Column       | Type        | Constraints | Default | Description |
| ------------ | ----------- | ----------- | ------- | ----------- |
| `identifier` | `varchar`   | PK NOT NULL |         |             |
| `token`      | `varchar`   | PK NOT NULL |         |             |
| `expires`    | `timestamp` | NOT NULL    |         |             |

#### Definition

```sql
CREATE TABLE `project1_verification_token` (
	`identifier` varchar(255) NOT NULL,
	`token` varchar(255) NOT NULL,
	`expires` timestamp NOT NULL,
	CONSTRAINT PRIMARY KEY(`identifier`,`token`)
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
select `identifier`, `token`, `expires`
from `project1_verification_token`
```

**Count**: How many rows verificationTokens has.

```ts
db.select({ count: count() }).from(verificationTokens);
```

```sql
select count(*)
from `project1_verification_token`
```

**Find many**: The rows of verificationTokens, through the relational query API.

```ts
db.query.verificationTokens.findMany();
```

```sql
select `d0`.`identifier` as `identifier`, `d0`.`token` as `token`, `d0`.`expires` as `expires`
from `project1_verification_token` as `d0`
```

**Find by primary key**: One row of verificationTokens, through the relational query API.

```ts
db.query.verificationTokens.findFirst({ where: { identifier: "identifier", token: "token" } });
```

```sql
select `d0`.`identifier` as `identifier`, `d0`.`token` as `token`, `d0`.`expires` as `expires`
from `project1_verification_token` as `d0`
where ((`d0`.`identifier` = ?) and (`d0`.`token` = ?))
limit ?
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
insert into `project1_verification_token` (`identifier`, `token`, `expires`)
values (?, ?, ?)
-- params: ["identifier","token","2026-01-01 00:00:00.000"]
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
update `project1_verification_token`
set `expires` = ?
where ((`project1_verification_token`.`identifier` = ?) and (`project1_verification_token`.`token` = ?))
-- params: ["2026-01-01 00:00:00.000","identifier","token"]
```

**Delete**: Deletes a row of verificationTokens, found by its primary key.

```ts
db.delete(verificationTokens).where(
  and(eq(verificationTokens.identifier, "identifier"), eq(verificationTokens.token, "token")),
);
```

```sql
delete from `project1_verification_token`
where ((`project1_verification_token`.`identifier` = ?) and (`project1_verification_token`.`token` = ?))
-- params: ["identifier","token"]
```

</details>
