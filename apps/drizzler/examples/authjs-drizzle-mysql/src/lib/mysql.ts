import {
  boolean,
  int,
  mysqlTable,
  primaryKey,
  timestamp,
  varchar,
} from "drizzle-orm/mysql-core"

// Auth.js account types. Upstream imports this from `@auth/core/adapters`.
type AdapterAccountType = "oauth" | "oidc" | "email" | "credentials" | "webauthn"

export function defineTables(
  schema: {
    usersTable?: ReturnType<typeof mysqlTable>
    accountsTable?: ReturnType<typeof mysqlTable>
    sessionsTable?: ReturnType<typeof mysqlTable>
    verificationTokensTable?: ReturnType<typeof mysqlTable>
    authenticatorsTable?: ReturnType<typeof mysqlTable>
  } = {}
) {
  const usersTable =
    schema.usersTable ??
    mysqlTable("user", {
      id: varchar("id", { length: 255 })
        .primaryKey()
        .$defaultFn(() => crypto.randomUUID()),
      name: varchar("name", { length: 255 }),
      email: varchar("email", { length: 255 }).unique(),
      emailVerified: timestamp("emailVerified", { mode: "date", fsp: 3 }),
      image: varchar("image", { length: 255 }),
    })

  const accountsTable =
    schema.accountsTable ??
    mysqlTable(
      "account",
      {
        userId: varchar("userId", { length: 255 })
          .notNull()
          .references(() => usersTable.id, { onDelete: "cascade" }),
        type: varchar("type", { length: 255 })
          .$type<AdapterAccountType>()
          .notNull(),
        provider: varchar("provider", { length: 255 }).notNull(),
        providerAccountId: varchar("providerAccountId", {
          length: 255,
        }).notNull(),
        refresh_token: varchar("refresh_token", { length: 255 }),
        access_token: varchar("access_token", { length: 255 }),
        expires_at: int("expires_at"),
        token_type: varchar("token_type", { length: 255 }),
        scope: varchar("scope", { length: 255 }),
        id_token: varchar("id_token", { length: 2048 }),
        session_state: varchar("session_state", { length: 255 }),
      },
      (account) => ({
        compositePk: primaryKey({
          columns: [account.provider, account.providerAccountId],
        }),
      })
    )

  const sessionsTable =
    schema.sessionsTable ??
    mysqlTable("session", {
      sessionToken: varchar("sessionToken", { length: 255 }).primaryKey(),
      userId: varchar("userId", { length: 255 })
        .notNull()
        .references(() => usersTable.id, { onDelete: "cascade" }),
      expires: timestamp("expires", { mode: "date" }).notNull(),
    })

  const verificationTokensTable =
    schema.verificationTokensTable ??
    mysqlTable(
      "verificationToken",
      {
        identifier: varchar("identifier", { length: 255 }).notNull(),
        token: varchar("token", { length: 255 }).notNull(),
        expires: timestamp("expires", { mode: "date" }).notNull(),
      },
      (verficationToken) => ({
        compositePk: primaryKey({
          columns: [verficationToken.identifier, verficationToken.token],
        }),
      })
    )

  const authenticatorsTable =
    schema.authenticatorsTable ??
    mysqlTable(
      "authenticator",
      {
        credentialID: varchar("credentialID", { length: 255 })
          .notNull()
          .unique(),
        userId: varchar("userId", { length: 255 })
          .notNull()
          .references(() => usersTable.id, { onDelete: "cascade" }),
        providerAccountId: varchar("providerAccountId", {
          length: 255,
        }).notNull(),
        credentialPublicKey: varchar("credentialPublicKey", {
          length: 255,
        }).notNull(),
        counter: int("counter").notNull(),
        credentialDeviceType: varchar("credentialDeviceType", {
          length: 255,
        }).notNull(),
        credentialBackedUp: boolean("credentialBackedUp").notNull(),
        transports: varchar("transports", { length: 255 }),
      },
      (authenticator) => ({
        compositePk: primaryKey({
          columns: [authenticator.userId, authenticator.credentialID],
        }),
      })
    )

  return {
    usersTable,
    accountsTable,
    sessionsTable,
    verificationTokensTable,
    authenticatorsTable,
  }
}

export const {
  usersTable,
  accountsTable,
  sessionsTable,
  verificationTokensTable,
  authenticatorsTable,
} = defineTables()
