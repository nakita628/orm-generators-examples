package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class VerificationTokenEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, VerificationTokenEntity>(VerificationTokenTable)

    var identifier by VerificationTokenTable.identifier

    /** Not unique alone: better-auth stores plain values here (a user id, a counter). See WEB-19201. */
    var token by VerificationTokenTable.token
    var expires by VerificationTokenTable.expires
    var createdAt by VerificationTokenTable.createdAt
    var updatedAt by VerificationTokenTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(VerificationTokenTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
