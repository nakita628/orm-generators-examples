package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Each gateway virtual key, folded from gateway's vk_lifecycle facts: nurturing's own copy. The
 * creator comes from the created fact alone; the holder and purpose are each the latest non-null
 * value by the fact's time (stamped beside them); the state times keep their latest.
 */
object NurturingVirtualKeyTable : IdTable<String>("\"NurturingVirtualKey\"") {
    override val id = text("virtualKeyId").entityId()
    val organizationId = text("organizationId")
    val name = text("name")
    val occurredAt = pgTimestamp("occurredAt", 3)
    val keyCreatedAt = pgTimestamp("keyCreatedAt", 3).nullable()
    val creatorUserId = text("creatorUserId").nullable()
    val principalUserId = text("principalUserId").nullable()
    val principalUserIdAt = pgTimestamp("principalUserIdAt", 3).nullable()
    val purpose = text("purpose").nullable()
    val purposeAt = pgTimestamp("purposeAt", 3).nullable()
    val disabledAt = pgTimestamp("disabledAt", 3).nullable()
    val enabledAt = pgTimestamp("enabledAt", 3).nullable()
    val revokedAt = pgTimestamp("revokedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"NurturingVirtualKey_pkey\"")

    init {
        index("\"NurturingVirtualKey_creatorUserId_idx\"", false, creatorUserId)
        index("\"NurturingVirtualKey_principalUserId_idx\"", false, principalUserId)
    }
}
