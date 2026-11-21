import {
  integer,
  primaryKey,
  sqliteTable,
  text,
} from "drizzle-orm/sqlite-core"

// Auth.js account types. Upstream imports this from `@auth/core/adapters`.
type AdapterAccountType = "oauth" | "oidc" | "email" | "credentials" | "webauthn"

export function defineTables(
  schema: {
    usersTable?: ReturnType<typeof sqliteTable>
    accountsTable?: ReturnType<typeof sqliteTable>
    sessionsTable?: ReturnType<typeof sqliteTable>
    verificationTokensTable?: ReturnType<typeof sqliteTable>
    authenticatorsTable?: ReturnType<typeof sqliteTable>
  } = {}
) {
  const usersTable =
    schema.usersTable ??
    sqliteTable("user", {
      id: text("id")
        .primaryKey()
        .$defaultFn(() => crypto.randomUUID()),
      name: text("name"),
      email: text("email").unique(),
      emailVerified: integer("emailVerified", { mode: "timestamp_ms" }),
      image: text("image"),
    })

  const accountsTable =
    schema.accountsTable ??
    sqliteTable(
      "account",
      {
        userId: text("userId")
          .notNull()
          .references(() => usersTable.id, { onDelete: "cascade" }),
        type: text("type").$type<AdapterAccountType>().notNull(),
        provider: text("provider").notNull(),
        providerAccountId: text("providerAccountId").notNull(),
        refresh_token: text("refresh_token"),
        access_token: text("access_token"),
        expires_at: integer("expires_at"),
        token_type: text("token_type"),
        scope: text("scope"),
        id_token: text("id_token"),
        session_state: text("session_state"),
      },
      (account) => ({
        compositePk: primaryKey({
          columns: [account.provider, account.providerAccountId],
        }),
      })
    )

  const sessionsTable =
    schema.sessionsTable ??
    sqliteTable("session", {
      sessionToken: text("sessionToken").primaryKey(),
      userId: text("userId")
        .notNull()
        .references(() => usersTable.id, { onDelete: "cascade" }),
      expires: integer("expires", { mode: "timestamp_ms" }).notNull(),
    })

  const verificationTokensTable =
    schema.verificationTokensTable ??
    sqliteTable(
      "verificationToken",
      {
        identifier: text("identifier").notNull(),
        token: text("token").notNull(),
        expires: integer("expires", { mode: "timestamp_ms" }).notNull(),
      },
      (verficationToken) => ({
        compositePk: primaryKey({
          columns: [verficationToken.identifier, verficationToken.token],
        }),
      })
    )

  const authenticatorsTable =
    schema.authenticatorsTable ??
    sqliteTable(
      "authenticator",
      {
        credentialID: text("credentialID").notNull().unique(),
        userId: text("userId")
          .notNull()
          .references(() => usersTable.id, { onDelete: "cascade" }),
        providerAccountId: text("providerAccountId").notNull(),
        credentialPublicKey: text("credentialPublicKey").notNull(),
        counter: integer("counter").notNull(),
        credentialDeviceType: text("credentialDeviceType").notNull(),
        credentialBackedUp: integer("credentialBackedUp", {
          mode: "boolean",
        }).notNull(),
        transports: text("transports"),
      },
      (authenticator) => ({
        compositePK: primaryKey({
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
