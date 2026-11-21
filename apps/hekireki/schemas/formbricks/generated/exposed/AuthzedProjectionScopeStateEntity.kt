package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * ENG-3282: per-scope readiness of a SpiceDB projection. A row with `readyAt` set means the scope was
 * backfilled and audited clean, so the evaluator may enforce it. Written only by the backfill CLI.
 */
class AuthzedProjectionScopeStateEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuthzedProjectionScopeStateEntity>(AuthzedProjectionScopeStateTable)

    var readyAt by AuthzedProjectionScopeStateTable.readyAt
    var readyBy by AuthzedProjectionScopeStateTable.readyBy
    var updatedAt by AuthzedProjectionScopeStateTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuthzedProjectionScopeStateTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
