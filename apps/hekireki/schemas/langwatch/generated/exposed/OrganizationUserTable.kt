package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object OrganizationUserTable : CompositeIdTable("\"OrganizationUser\"") {
    val userId = text("userId").entityId()
    val organizationId = text("organizationId").entityId()
    val role = pgEnum("role", "\"OrganizationUserRole\"", OrganizationUserRole.entries, OrganizationUserRole::dbName)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /**
     * Department assignment for spend attribution (the people lens).
     * Single-valued, org-scoped, accounting-only (never gates access).
     * Plain column ref (relationMode="prisma"); the service validates the
     * id belongs to this org. See departments.feature.
     */
    val departmentId = text("departmentId").nullable()

    /**
     * Set when an admin disables this membership to stay within the licensed
     * seat count. The person loses access to this organization only, keeps
     * their role, department and everything they did, and frees their seat.
     * Reversible: re-enabling restores access if a seat is available.
     * Seat counting excludes disabled memberships. See
     * seat-reconciliation.feature.
     */
    val disabledAt = pgTimestamp("disabledAt", 3).nullable()

    /**
     * Stable grant intent for an unfinished automatic SSO admission.
     * Cleared after the authoritative grant is observed; never inferred from missing access.
     */
    val pendingSsoGrantId = text("pendingSsoGrantId").nullable()

    /**
     * A new generation for each membership lifetime. Stamped USER grant facts
     * must match the current active membership before they can be projected;
     * removing and rejoining therefore cannot revive a delayed old attach.
     */
    val membershipStamp = text("membershipStamp").databaseDefault("gen_random_uuid()::text")

    /**
     * When this membership started holding the full seat it holds now; null means createdAt.
     * Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
     */
    val fullSeatSince = pgTimestamp("fullSeatSince", 3).nullable()

    /**
     * Set when admission wrote this membership disabled because no seat was free: the person
     * is waiting for one. Cleared when the membership is enabled.
     */
    val seatPendingSince = pgTimestamp("seatPendingSince", 3).nullable()
    override val primaryKey = PrimaryKey(userId, organizationId, name = "\"OrganizationUser_pkey\"")

    init {
        index("\"OrganizationUser_organizationId_idx\"", false, organizationId)
        index("\"OrganizationUser_userId_idx\"", false, userId)
        index("\"OrganizationUser_departmentId_idx\"", false, departmentId)
    }
}
