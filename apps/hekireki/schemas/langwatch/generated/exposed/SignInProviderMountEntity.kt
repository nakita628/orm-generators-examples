package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class SignInProviderMountEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SignInProviderMountEntity>(SignInProviderMountTable)

    /** The configured provider's name; null when the deployment signs in by email. */
    var provider by SignInProviderMountTable.provider
    var mounted by SignInProviderMountTable.mounted

    /** The last sso event applied: the row's idempotency guard and its version. */
    var lastAppliedEventId by SignInProviderMountTable.lastAppliedEventId
    var updatedAt by SignInProviderMountTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SignInProviderMountTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
