package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
 * and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
 * so retention adds nothing to the survey row.
 */
object RetentionSurveyRestoreTable : IdTable<String>("\"RetentionSurveyRestore\"") {
    override val id =
        text("surveyId")
            .references(
                SurveyTable.id,
                onDelete = ReferenceOption.CASCADE,
                onUpdate = ReferenceOption.CASCADE,
                fkName = "\"RetentionSurveyRestore_surveyId_fkey\"",
            ).entityId()
    val restoredAt = pgTimestamp("restoredAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"RetentionSurveyRestore_pkey\"")
}
