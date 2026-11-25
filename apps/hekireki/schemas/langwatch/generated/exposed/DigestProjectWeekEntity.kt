package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** A project's ISO week, folded from its owners' events. The id is the fold's aggregate key. */
class DigestProjectWeekEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DigestProjectWeekEntity>(DigestProjectWeekTable)

    var projectId by DigestProjectWeekTable.projectId
    var organizationId by DigestProjectWeekTable.organizationId
    var isoWeek by DigestProjectWeekTable.isoWeek

    /** The fold state: counts, sums, a fixed-bucket latency histogram, capped model and evaluator maps. */
    var state by DigestProjectWeekTable.state
    var occurredAt by DigestProjectWeekTable.occurredAt
    var lastEventId by DigestProjectWeekTable.lastEventId
    var acceptedAt by DigestProjectWeekTable.acceptedAt
    var projectionVersion by DigestProjectWeekTable.projectionVersion
    var createdAt by DigestProjectWeekTable.createdAt
    var updatedAt by DigestProjectWeekTable.updatedAt
}
