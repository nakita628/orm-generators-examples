package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
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
object EmbeddedDataTable : IdTable<String>("\"EmbeddedData\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val key = text("key").nullable()
    val name = text("name")
    val description = text("description").nullable()
    val source1 = pgEnum("source", "\"EmbeddedDataSource\"", EmbeddedDataSource.entries, EmbeddedDataSource::dbName)
    val dataType =
        pgEnum("dataType", "\"EmbeddedDataType\"", EmbeddedDataType.entries, EmbeddedDataType::dbName).default(EmbeddedDataType.STRING)

    /** [EmbeddedDataDefaultValue] */
    val defaultValue = jsonb<String>("defaultValue", { it }, { it }).nullable()
    val locked = bool("locked").default(false)
    val surveyId = text("surveyId").nullable()
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"EmbeddedData_workspaceId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"EmbeddedData_pkey\"")

    init {
        index("\"EmbeddedData_surveyId_idx\"", false, surveyId)
        uniqueIndex("\"EmbeddedData_workspaceId_key_key\"", workspaceId, key)
        uniqueIndex("\"EmbeddedData_id_workspaceId_key\"", id, workspaceId)
        foreignKey(
            surveyId to SurveyTable.id,
            workspaceId to SurveyTable.workspaceId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            name = "\"EmbeddedData_surveyId_workspaceId_fkey\"",
        )
    }
}
