package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One dated person → department link (ADR-128 §13, #7882).
 *
 * `OrganizationUser.departmentId` stays the CURRENT assignment every screen
 * reads; these rows are its history, appended by the same `assignUser` write
 * so the two cannot drift. They exist for the cost reads: January's spend
 * stays with January's department, so a read resolves against the link that
 * was open on the cost's day, not against today's pointer. Members with no
 * link on a day resolve to "unassigned", never to an error.
 *
 * Same shape as IdentityMatch's dated links, same one-open-link guard: the
 * partial unique index in the migration rejects a second open row per
 * (organizationId, userId) with SQLSTATE 23505.
 */
object DepartmentMembershipHistoryTable : IdTable<String>("\"DepartmentMembershipHistory\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val userId = text("userId")
    val departmentId = text("departmentId")
    val validFrom = pgTimestamp("validFrom", 3)

    /** Null means the link is open. Reassignment and clearing close it. */
    val validTo = pgTimestamp("validTo", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DepartmentMembershipHistory_pkey\"")

    init {
        index("DepartmentMembershipHistory_organizationId_userId_validFrom_idx", false, organizationId, userId, validFrom)
    }
}
