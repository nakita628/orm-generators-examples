package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class SsoDetailsEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SsoDetailsEntity>(SsoDetailsTable)

    var domain by SsoDetailsTable.domain
    val organizations by OrganizationEntity optionalReferrersOn OrganizationTable.ssoDetailsId
    var baseUserGroupId by SsoDetailsTable.baseUserGroupId
    var selfServiceGroupId by SsoDetailsTable.selfServiceGroupId
    var adminGroupId by SsoDetailsTable.adminGroupId
    var managerGroupId by SsoDetailsTable.managerGroupId
    var requireSsoLogin by SsoDetailsTable.requireSsoLogin
    var createdAt by SsoDetailsTable.createdAt
    var updatedAt by SsoDetailsTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SsoDetailsTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
