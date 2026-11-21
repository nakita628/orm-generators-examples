package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class CustodyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CustodyEntity>(CustodyTable)

    var custodian by TeamMemberEntity referencedOn CustodyTable.teamMemberId
    var teamMemberId by CustodyTable.teamMemberId
    var asset by AssetEntity referencedOn CustodyTable.assetId
    var assetId by CustodyTable.assetId
    var kitCustody by KitCustodyEntity optionalReferencedOn CustodyTable.kitCustodyId
    var kitCustodyId by CustodyTable.kitCustodyId
    var quantity by CustodyTable.quantity
    var location by LocationEntity optionalReferencedOn CustodyTable.locationId
    var locationId by CustodyTable.locationId
    var sourceUnknown by CustodyTable.sourceUnknown
    var createdAt by CustodyTable.createdAt
    var updatedAt by CustodyTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(CustodyTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
