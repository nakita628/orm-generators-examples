package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * A project nurturing saw traces of with no placement copy, and when it first did: every worker
 * retries its trace batches only within one window from then, across restarts. Deleted when the
 * copy lands.
 */
class NurturingUncopiedProjectEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingUncopiedProjectEntity>(NurturingUncopiedProjectTable)

    var firstSeenAt by NurturingUncopiedProjectTable.firstSeenAt
}
