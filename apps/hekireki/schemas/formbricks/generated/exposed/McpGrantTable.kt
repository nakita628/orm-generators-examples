package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
 * its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
 * a token is now bound to the resource it was approved for instead of to a static allow-list.
 * `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
 * What one user allows one MCP client to do (GLOSSARY.md "Grant", ADR-0041): its organization & workspace
 * access and an access level per permission. One per OAuth connection, keyed by user and client rather
 * than by `oauthConsent` row, because the provider can write duplicate consent rows for one connection
 * and settles them later. Every MCP request reads it; without one the request is refused. The client's
 * OAuth scopes stay on the consent and cap it.
 */
object McpGrantTable : IdTable<String>("\"McpGrant\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"McpGrant_userId_fkey\"",
        )
    val clientId =
        text("clientId").references(
            OauthClientTable.clientId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"McpGrant_clientId_fkey\"",
        )

    /**
     * Which organizations and workspaces the grant reaches (ENG-7153): All access, only the organizations
     * in `organizations` (and every workspace in them, later ones included), or only the workspaces in
     * `workspaces`. Enforced in SpiceDB through `can()` (ADR-0042); a value no row backs reaches nothing.
     */
    val access = pgEnum("access", "\"McpGrantAccess\"", McpGrantAccess.entries, McpGrantAccess::dbName).default(McpGrantAccess.ALL)

    /**
     * The access level per permission of `apps/web/modules/mcp/permissions.ts`: `{ "surveys": "write" }`.
     * A permission that is absent has No access.
     */
    val permissions = jsonb<String>("permissions", { it }, { it }).default("{}")
    override val primaryKey = PrimaryKey(id, name = "\"McpGrant_pkey\"")

    init {
        index("\"McpGrant_clientId_idx\"", false, clientId)
        uniqueIndex("\"McpGrant_userId_clientId_key\"", userId, clientId)
    }
}
