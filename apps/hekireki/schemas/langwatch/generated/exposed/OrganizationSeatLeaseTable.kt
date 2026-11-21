package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Held by the one seat-taking write running on an organization (seat-requests.feature). A
 * row past expiresAt is free to take; a holder deletes its own row when its write is done.
 */
object OrganizationSeatLeaseTable : IdTable<String>("\"OrganizationSeatLease\"") {
    override val id = text("organizationId").entityId()

    /** Random per hold, so a release deletes only its own lease. */
    val holder = text("holder")
    val expiresAt = pgTimestamp("expiresAt", 3)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"OrganizationSeatLease_pkey\"")
}
