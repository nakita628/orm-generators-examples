package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** A person's ISO week of coding-agent activity in one organization, folded from the agents' facts. */
class DigestUserWeekEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DigestUserWeekEntity>(DigestUserWeekTable)

    var userId by DigestUserWeekTable.userId
    var organizationId by DigestUserWeekTable.organizationId
    var isoWeek by DigestUserWeekTable.isoWeek

    /** Sessions (a capped set of ids) and the agents' metric series, converged last-write-wins. */
    var state by DigestUserWeekTable.state
    var occurredAt by DigestUserWeekTable.occurredAt
    var lastEventId by DigestUserWeekTable.lastEventId
    var acceptedAt by DigestUserWeekTable.acceptedAt
    var projectionVersion by DigestUserWeekTable.projectionVersion
    var createdAt by DigestUserWeekTable.createdAt
    var updatedAt by DigestUserWeekTable.updatedAt
}
