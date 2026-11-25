package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** An organization enrolled in the weekly digest. No row: not enrolled, so nothing sends. */
class DigestCohortOrganizationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DigestCohortOrganizationEntity>(DigestCohortOrganizationTable)

    var enrolledBy by DigestCohortOrganizationTable.enrolledBy
    var enrolledAt by DigestCohortOrganizationTable.enrolledAt
}
