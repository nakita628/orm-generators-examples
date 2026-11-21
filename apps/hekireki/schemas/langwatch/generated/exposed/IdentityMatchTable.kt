package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The dated match table: "`user_abc123` on Anthropic is platform user #42,
 * valid from March 3", closed when it stops being true.
 *
 * Dating is what makes a re-issued email address survivable: the leaver's link
 * closes at offboarding, the new hire gets a new link, and January's spend
 * stays with January's person. A correction closes the old row and opens a new
 * one; nothing is ever silently rewritten.
 */
object IdentityMatchTable : IdTable<String>("\"IdentityMatch\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")

    /** The provider-side identity this link is about. */
    val discoveredPersonId = text("discoveredPersonId")

    /**
     * The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
     * the row and its dates remain, so the history of the link is not rewritten
     * by the erasure either.
     */
    val userId = text("userId").nullable()

    /**
     * directory_id | verified_email | human_confirmed. Exact verified-email
     * equality is the primary evidence; a directory id corroborates but never
     * stands alone (ADR-128 §12).
     */
    val evidenceKind = text("evidenceKind")
    val validFrom = pgTimestamp("validFrom", 3)

    /** Null means the link is open. Offboarding and corrections close it. */
    val validTo = pgTimestamp("validTo", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"IdentityMatch_pkey\"")

    init {
        index("\"IdentityMatch_organizationId_discoveredPersonId_idx\"", false, organizationId, discoveredPersonId)
        index("\"IdentityMatch_organizationId_userId_idx\"", false, organizationId, userId)
    }
}
