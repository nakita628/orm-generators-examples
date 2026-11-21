package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Links surveys to specific trigger actions.
 * Defines which user actions should cause a survey to be displayed.
 * This is the connection table between Surveys and ActionClasses that determines
 * when and under what conditions a survey should be triggered.
 *
 * @property id - Unique identifier for the trigger
 * @property survey - The survey to be triggered
 * @property actionClass - The action that triggers the survey
 * @property createdAt - When the trigger was created
 * @property updatedAt - When the trigger was last modified
 */
object SurveyTriggerTable : IdTable<String>("\"SurveyTrigger\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val surveyId =
        reference(
            "surveyId",
            SurveyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"SurveyTrigger_surveyId_fkey\"",
        )
    val actionClassId =
        reference(
            "actionClassId",
            ActionClassTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"SurveyTrigger_actionClassId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"SurveyTrigger_pkey\"")

    init {
        uniqueIndex("\"SurveyTrigger_surveyId_actionClassId_key\"", surveyId, actionClassId)
    }
}
