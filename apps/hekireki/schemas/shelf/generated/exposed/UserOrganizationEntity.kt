package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class UserOrganizationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserOrganizationEntity>(UserOrganizationTable)

    var user by UserEntity referencedOn UserOrganizationTable.userId
    var userId by UserOrganizationTable.userId
    var organization by OrganizationEntity referencedOn UserOrganizationTable.organizationId
    var organizationId by UserOrganizationTable.organizationId
    var roles by UserOrganizationTable.roles
    var createdAt by UserOrganizationTable.createdAt
    var updatedAt by UserOrganizationTable.updatedAt

    /**
     * Secret, unguessable token for this member's iCal calendar-feed subscription
     * in this workspace. Null until the user generates a feed; regenerating it
     * rotates the value (invalidating any previously-shared subscription URL).
     * `@unique` doubles as the lookup index for the public feed route.
     */
    var calendarTokenId by UserOrganizationTable.calendarTokenId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(UserOrganizationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
