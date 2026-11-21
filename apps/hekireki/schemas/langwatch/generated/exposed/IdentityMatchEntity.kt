package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * The dated match table: "`user_abc123` on Anthropic is platform user #42,
 * valid from March 3", closed when it stops being true.
 *
 * Dating is what makes a re-issued email address survivable: the leaver's link
 * closes at offboarding, the new hire gets a new link, and January's spend
 * stays with January's person. A correction closes the old row and opens a new
 * one; nothing is ever silently rewritten.
 */
class IdentityMatchEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, IdentityMatchEntity>(IdentityMatchTable)

    var organizationId by IdentityMatchTable.organizationId

    /** The provider-side identity this link is about. */
    var discoveredPersonId by IdentityMatchTable.discoveredPersonId

    /**
     * The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
     * the row and its dates remain, so the history of the link is not rewritten
     * by the erasure either.
     */
    var userId by IdentityMatchTable.userId

    /**
     * directory_id | verified_email | human_confirmed. Exact verified-email
     * equality is the primary evidence; a directory id corroborates but never
     * stands alone (ADR-128 §12).
     */
    var evidenceKind by IdentityMatchTable.evidenceKind
    var validFrom by IdentityMatchTable.validFrom

    /** Null means the link is open. Offboarding and corrections close it. */
    var validTo by IdentityMatchTable.validTo
    var createdAt by IdentityMatchTable.createdAt
}
