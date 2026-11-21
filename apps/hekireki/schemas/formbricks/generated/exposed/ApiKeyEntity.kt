package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Represents enhanced API authentication keys with organization-level ownership.
 * Used for authenticating API requests to Formbricks with more granular permissions.
 *
 * @property id - Unique identifier for the API key
 * @property label - Optional descriptive name for the key
 * @property hashedKey - Securely stored API key
 * @property organization - The organization this key belongs to
 * @property createdBy - User ID who created this key
 * @property lastUsedAt - Timestamp of last usage
 * @property apiKeyWorkspaces - Workspaces this key has access to
 */
class ApiKeyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ApiKeyEntity>(ApiKeyTable)

    var createdAt by ApiKeyTable.createdAt
    var createdBy by ApiKeyTable.createdBy
    var lastUsedAt by ApiKeyTable.lastUsedAt
    var label by ApiKeyTable.label
    var hashedKey by ApiKeyTable.hashedKey
    var lookupHash by ApiKeyTable.lookupHash
    var organizationId by ApiKeyTable.organizationId
    var organization by OrganizationEntity referencedOn ApiKeyTable.organizationId
    val apiKeyWorkspaces by ApiKeyWorkspaceEntity referrersOn ApiKeyWorkspaceTable.apiKeyId

    /** [OrganizationAccess] */
    var organizationAccess by ApiKeyTable.organizationAccess
}
