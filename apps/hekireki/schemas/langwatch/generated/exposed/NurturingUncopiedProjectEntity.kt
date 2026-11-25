package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * retired: unread since v3.28.0. The trace-day folds count a trace with no placement copy, so no
 * retry clock reads or writes this table; an operator step may drop it once old writers are gone.
 */
class NurturingUncopiedProjectEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingUncopiedProjectEntity>(NurturingUncopiedProjectTable)

    var firstSeenAt by NurturingUncopiedProjectTable.firstSeenAt
}
