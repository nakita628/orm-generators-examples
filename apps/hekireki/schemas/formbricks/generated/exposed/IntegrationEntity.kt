package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Represents third-party service integrations.
 * Enables data flow between Formbricks and external services.
 *
 * @property id - Unique identifier for the integration
 * @property type - The service being integrated (Google Sheets, Notion, etc.)
 * @property workspace - The workspace this integration belongs to
 * @property config - Service-specific configuration details
 */
class IntegrationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, IntegrationEntity>(IntegrationTable)

    var type by IntegrationTable.type

    /** [IntegrationConfig] */
    var config by IntegrationTable.config
    var workspace by WorkspaceEntity referencedOn IntegrationTable.workspaceId
    var workspaceId by IntegrationTable.workspaceId
}
