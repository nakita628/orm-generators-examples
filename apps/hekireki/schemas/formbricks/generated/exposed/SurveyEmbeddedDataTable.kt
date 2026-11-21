package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Links a survey to an Embedded Data field and records the key that field's value is stored
 * under inside that survey.
 *
 * `workspaceId` is part of both foreign keys on purpose: `surveyId` and `embeddedDataId` alone are
 * independent, so nothing would stop a survey in one workspace linking a field defined in another.
 * Referencing `(id, workspaceId)` on both sides makes a cross-workspace pair unrepresentable, and
 * puts the workspace scope every query needs on the row instead of behind a join through Survey.
 *
 * @property storageKey - The key the value lives under: a cuid for computed fields, the field
 * name for ingested ones. Migrated fields keep their original key, which is what lets existing
 * recall tokens, logic conditions and stored responses keep resolving untouched.
 *
 * @property order - Position within this survey's field list, dense from 0. Lives on the link
 * rather than on `EmbeddedData` because a shared definition sits at a different position in each
 * survey that links it. Deliberately has no default: every writer knows the position it is writing,
 * so a missing one should be a type error rather than a row that silently sorts to the front.
 */
object SurveyEmbeddedDataTable : IdTable<String>("\"SurveyEmbeddedData\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val workspaceId = text("workspaceId")
    val surveyId = text("surveyId")
    val embeddedDataId = text("embeddedDataId")
    val storageKey = text("storageKey")
    val order = integer("order")
    override val primaryKey = PrimaryKey(id, name = "\"SurveyEmbeddedData_pkey\"")

    init {
        index("\"SurveyEmbeddedData_embeddedDataId_idx\"", false, embeddedDataId)
        index("\"SurveyEmbeddedData_surveyId_order_idx\"", false, surveyId, order)
        uniqueIndex("\"SurveyEmbeddedData_surveyId_embeddedDataId_key\"", surveyId, embeddedDataId)
        uniqueIndex("\"SurveyEmbeddedData_surveyId_storageKey_key\"", surveyId, storageKey)
        foreignKey(
            surveyId to SurveyTable.id,
            workspaceId to SurveyTable.workspaceId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            name = "\"SurveyEmbeddedData_surveyId_workspaceId_fkey\"",
        )
        foreignKey(
            embeddedDataId to EmbeddedDataTable.id,
            workspaceId to EmbeddedDataTable.workspaceId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            name = "\"SurveyEmbeddedData_embeddedDataId_workspaceId_fkey\"",
        )
    }
}
