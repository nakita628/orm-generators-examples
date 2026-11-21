package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a user action that can trigger surveys.
 * Used to define points in the user journey where surveys can be shown.
 *
 * @property id - Unique identifier for the action
 * @property name - Display name of the action
 * @property type - Whether this is a code or no-code action
 * @property key - Unique identifier used in code implementation
 * @property noCodeConfig - Configuration for no-code setup
 * @property workspace - The workspace this action belongs to
 */
class ActionClassEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ActionClassEntity>(ActionClassTable)

    var createdAt by ActionClassTable.createdAt
    var updatedAt by ActionClassTable.updatedAt
    var name by ActionClassTable.name
    var description by ActionClassTable.description
    var type by ActionClassTable.type
    var key by ActionClassTable.key

    /** [ActionClassNoCodeConfig] */
    var noCodeConfig by ActionClassTable.noCodeConfig
    var workspace by WorkspaceEntity referencedOn ActionClassTable.workspaceId
    var workspaceId by ActionClassTable.workspaceId
    val surveyTriggers by SurveyTriggerEntity referrersOn SurveyTriggerTable.actionClassId
    val creativeTriggers by CreativeTriggerEntity referrersOn CreativeTriggerTable.actionClassId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ActionClassTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
