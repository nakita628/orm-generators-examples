package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class McpGrantEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, McpGrantEntity>(McpGrantTable)

    var createdAt by McpGrantTable.createdAt
    var updatedAt by McpGrantTable.updatedAt
    var userId by McpGrantTable.userId
    var user by UserEntity referencedOn McpGrantTable.userId
    var clientId by McpGrantTable.clientId
    var client by OauthClientEntity referencedOn McpGrantTable.clientId

    /**
     * Which organizations and workspaces the grant reaches (ENG-7153): All access, only the organizations
     * in `organizations` (and every workspace in them, later ones included), or only the workspaces in
     * `workspaces`. Enforced in SpiceDB through `can()` (ADR-0042); a value no row backs reaches nothing.
     */
    var access by McpGrantTable.access

    /**
     * The access level per permission of `apps/web/modules/mcp/permissions.ts`: `{ "surveys": "write" }`.
     * A permission that is absent has No access.
     */
    var permissions by McpGrantTable.permissions

    /** The selections behind `organizations` and `workspaces` access (ENG-7153). */
    val organizations by McpGrantOrganizationEntity referrersOn McpGrantOrganizationTable.grantId
    val workspaces by McpGrantWorkspaceEntity referrersOn McpGrantWorkspaceTable.grantId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(McpGrantTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
