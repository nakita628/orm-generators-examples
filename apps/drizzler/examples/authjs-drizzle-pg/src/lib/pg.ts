import {
  boolean,
  integer,
  pgTable,
  primaryKey,
  text,
  timestamp,
} from "drizzle-orm/pg-core"

// Auth.js account types. Upstream imports this from `@auth/core/adapters`.
type AdapterAccountType = "oauth" | "oidc" | "email" | "credentials" | "webauthn"

export function defineTables(
  schema: {
    usersTable?: ReturnType<typeof pgTable>
    accountsTable?: ReturnType<typeof pgTable>
    sessionsTable?: ReturnType<typeof pgTable>
    verificationTokensTable?: ReturnType<typeof pgTable>
    authenticatorsTable?: ReturnType<typeof pgTable>
  } = {}
) {
  const usersTable =
    schema.usersTable ??
    pgTable("user", {
      id: text("id")
        .primaryKey()
        .$defaultFn(() => crypto.randomUUID()),
      name: text("name"),
      email: text("email").unique(),
      emailVerified: timestamp("emailVerified", { mode: "date" }),
      image: text("image"),
    })

  const accountsTable =
    schema.accountsTable ??
    pgTable(
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
    pgTable("session", {
      sessionToken: text("sessionToken").primaryKey(),
      userId: text("userId")
        .notNull()
        .references(() => usersTable.id, { onDelete: "cascade" }),
      expires: timestamp("expires", { mode: "date" }).notNull(),
    })

  const verificationTokensTable =
    schema.verificationTokensTable ??
    pgTable(
      "verificationToken",
      {
        identifier: text("identifier").notNull(),
        token: text("token").notNull(),
        expires: timestamp("expires", { mode: "date" }).notNull(),
      },
      (verificationToken) => ({
        compositePk: primaryKey({
          columns: [verificationToken.identifier, verificationToken.token],
        }),
      })
    )

  const authenticatorsTable =
    schema.authenticatorsTable ??
    pgTable(
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
        credentialBackedUp: boolean("credentialBackedUp").notNull(),
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
