package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines one Embedded Data field: the extra context stored on a response that isn't an
 * answer to a question (brand, plan, page type, a computed score, ...).
 *
 * A field is either **local** — used by a single survey, which is how existing variables and
 * hidden fields are migrated — or **shared**, meaning it is part of the workspace library and can
 * be linked by many surveys. There is no stored flag for that: `surveyId` is the owner and `key`
 * is the library name, so a local field has `surveyId` set and `key` null, and a shared field the
 * other way round. Deriving it rather than storing a boolean removes a column that could drift
 * out of step with the two fields it stands for.
 *
 * @property key - Library name, unique per workspace and immutable. Null for local fields
 * @property name - Human-readable display label
 * @property source - Where the value comes from (computed | ingested). `reserved` is part of the
 * shared enum but never stored as a row: those fields are a static catalog in code
 * @property dataType - Declared value type, used for coercion at ingest
 * @property defaultValue - Initial value (computed) or fallback when nothing arrives (ingested)
 * @property locked - Ingested only: ignore external writes and use the default instead
 * @property survey - The owning survey for a local field; cascades so nothing is orphaned
 */
class EmbeddedDataEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, EmbeddedDataEntity>(EmbeddedDataTable)

    var createdAt by EmbeddedDataTable.createdAt
    var updatedAt by EmbeddedDataTable.updatedAt
    var key by EmbeddedDataTable.key
    var name by EmbeddedDataTable.name
    var description by EmbeddedDataTable.description
    var source by EmbeddedDataTable.source1
    var dataType by EmbeddedDataTable.dataType

    /** [EmbeddedDataDefaultValue] */
    var defaultValue by EmbeddedDataTable.defaultValue
    var locked by EmbeddedDataTable.locked
    var surveyId by EmbeddedDataTable.surveyId
    var workspace by WorkspaceEntity referencedOn EmbeddedDataTable.workspaceId
    var workspaceId by EmbeddedDataTable.workspaceId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(EmbeddedDataTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
