package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Records of previously-valid API keys that are still accepted for authentication
 * during a grace window after rotation. Extend or end the grace period by updating `expiresAt`.
 */
class RevokedApiKeyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, RevokedApiKeyEntity>(RevokedApiKeyTable)

    var apiKey by RevokedApiKeyTable.apiKey
    var runtimeEnvironment by RuntimeEnvironmentEntity referencedOn RevokedApiKeyTable.runtimeEnvironmentId
    var runtimeEnvironmentId by RevokedApiKeyTable.runtimeEnvironmentId
    var expiresAt by RevokedApiKeyTable.expiresAt
    var createdAt by RevokedApiKeyTable.createdAt
}
