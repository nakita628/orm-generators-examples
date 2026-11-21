package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Whether a user signs in on the identifier path: identity's own switch, set when the
 * identifier backfill proves the user and pinned off by an operator's rollback. No row is the
 * legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
 */
class IdentitySignInSwitchEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, IdentitySignInSwitchEntity>(IdentitySignInSwitchTable)

    var onIdentifiers by IdentitySignInSwitchTable.onIdentifiers
    var pinned by IdentitySignInSwitchTable.pinned
    var switchedAt by IdentitySignInSwitchTable.switchedAt
    var lastAppliedEventId by IdentitySignInSwitchTable.lastAppliedEventId
    var updatedAt by IdentitySignInSwitchTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(IdentitySignInSwitchTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
