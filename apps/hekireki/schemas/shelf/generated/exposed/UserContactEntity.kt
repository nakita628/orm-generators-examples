package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class UserContactEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserContactEntity>(UserContactTable)

    var phone by UserContactTable.phone
    var street by UserContactTable.street
    var city by UserContactTable.city
    var stateProvince by UserContactTable.stateProvince
    var zipPostalCode by UserContactTable.zipPostalCode
    var countryRegion by UserContactTable.countryRegion
    var user by UserEntity referencedOn UserContactTable.userId
    var userId by UserContactTable.userId
    var createdAt by UserContactTable.createdAt
    var updatedAt by UserContactTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(UserContactTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
