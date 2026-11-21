package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class TierEntity(
    id: EntityID<TierId>,
) : Entity<TierId>(id) {
    companion object : EntityClass<TierId, TierEntity>(TierTable)

    var name by TierTable.name
    val subscribers by UserEntity referrersOn UserTable.tierId
    var tierLimitId by TierTable.tierLimitId
    var tierLimit by TierLimitEntity optionalReferencedOn TierTable.tierLimitId
    var createdAt by TierTable.createdAt
    var updatedAt by TierTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TierTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
