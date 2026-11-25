package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** An organization enrolled in the weekly digest. No row: not enrolled, so nothing sends. */
object DigestCohortOrganizationTable : IdTable<String>("\"DigestCohortOrganization\"") {
    override val id = text("organizationId").entityId()
    val enrolledBy = text("enrolledBy")
    val enrolledAt = pgTimestamp("enrolledAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DigestCohortOrganization_pkey\"")
}
