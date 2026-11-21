package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class KitCustodyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, KitCustodyEntity>(KitCustodyTable)

    var custodian by TeamMemberEntity referencedOn KitCustodyTable.custodianId
    var custodianId by KitCustodyTable.custodianId
    var kit by KitEntity referencedOn KitCustodyTable.kitId
    var kitId by KitCustodyTable.kitId
    val inheritedCustody by CustodyEntity optionalReferrersOn CustodyTable.kitCustodyId
    var createdAt by KitCustodyTable.createdAt
    var updatedAt by KitCustodyTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(KitCustodyTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
