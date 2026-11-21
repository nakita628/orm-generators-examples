package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot. */
class AuthzOrganizationPolicyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuthzOrganizationPolicyEntity>(AuthzOrganizationPolicyTable)

    var braidsOrganizationAdminsOnly by AuthzOrganizationPolicyTable.braidsOrganizationAdminsOnly
    var updatedByUserId by AuthzOrganizationPolicyTable.updatedByUserId
    var updatedAt by AuthzOrganizationPolicyTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuthzOrganizationPolicyTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
