package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class ScimTokenEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ScimTokenEntity>(ScimTokenTable)

    var organizationId by ScimTokenTable.organizationId

    /**
     * D08: the SSO connection this token was issued for, and its ENTIRE write
     * authority — a token minted for one connection can never touch another
     * connection's people. No relation, matching `SsoConnection` itself: that
     * projection carries no foreign keys, so neither may a row pointing at it.
     *
     * Nullable only for the tokens that predate the column. They were minted
     * when "which organization" was the whole answer, so they keep the
     * organization-wide authority they were sold with; the backfill points
     * each at its organization's first connection where there is one. Minting
     * a NEW token without a connection is refused (`scim_connection_required`).
     */
    var connectionId by ScimTokenTable.connectionId

    /**
     * UNIQUE, and that is a tenancy boundary rather than tidiness.
     *
     * The lookup that turns a bearer token into an organization is keyed on
     * this column alone — it has to be, since the token is the only thing the
     * caller presents. While every token was 32 random bytes a collision was
     * impossible and the missing constraint cost nothing. An administrator may
     * now supply a value they already hold, so two organizations CAN choose the
     * same string, and without this the lookup would answer with whichever row
     * the planner reached first: one customer's directory provisioning and
     * deleting another customer's people.
     */
    var hashedToken by ScimTokenTable.hashedToken

    /**
     * How `hashedToken` was derived, so the two schemes can coexist while
     * tokens minted under the old one are still in use.
     *
     * `sha256` is a bare digest, which is correct for 32 random bytes and wrong
     * for a value a person chose — a database dump of those is a wordlist away
     * from live credentials. `hmac-sha256` is keyed on the deployment's own
     * secret, so a dump on its own is inert. Every new row is the latter.
     */
    var hashScheme by ScimTokenTable.hashScheme
    var description by ScimTokenTable.description
    var createdAt by ScimTokenTable.createdAt
    var lastUsedAt by ScimTokenTable.lastUsedAt
}
