package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class RoleEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, RoleEntity>(RoleTable)

    var name by RoleTable.name
    var createdAt by RoleTable.createdAt
    var updatedAt by RoleTable.updatedAt
    var users by UserEntity via RoleToUserTable

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(RoleTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
