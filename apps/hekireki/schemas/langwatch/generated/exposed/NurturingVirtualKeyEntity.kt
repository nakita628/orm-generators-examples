package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Each gateway virtual key, folded from gateway's vk_lifecycle facts: nurturing's own copy. The
 * creator comes from the created fact alone; the holder and purpose are each the latest non-null
 * value by the fact's time (stamped beside them); the state times keep their latest.
 */
class NurturingVirtualKeyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingVirtualKeyEntity>(NurturingVirtualKeyTable)

    var organizationId by NurturingVirtualKeyTable.organizationId
    var name by NurturingVirtualKeyTable.name
    var occurredAt by NurturingVirtualKeyTable.occurredAt
    var keyCreatedAt by NurturingVirtualKeyTable.keyCreatedAt
    var creatorUserId by NurturingVirtualKeyTable.creatorUserId
    var principalUserId by NurturingVirtualKeyTable.principalUserId
    var principalUserIdAt by NurturingVirtualKeyTable.principalUserIdAt
    var purpose by NurturingVirtualKeyTable.purpose
    var purposeAt by NurturingVirtualKeyTable.purposeAt
    var disabledAt by NurturingVirtualKeyTable.disabledAt
    var enabledAt by NurturingVirtualKeyTable.enabledAt
    var revokedAt by NurturingVirtualKeyTable.revokedAt
}
