package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines follow-up actions for survey responses.
 * Enables automated actions based on specific survey response conditions.
 *
 * @property id - Unique identifier for the follow-up
 * @property survey - The associated survey
 * @property name - Display name for the follow-up
 * @property trigger - Conditions that activate the follow-up
 * @property action - Actions to take when triggered
 */
object SurveyFollowUpTable : IdTable<String>("\"SurveyFollowUp\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val surveyId =
        reference(
            "surveyId",
            SurveyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"SurveyFollowUp_surveyId_fkey\"",
        )
    val name = text("name")

    /** [SurveyFollowUpTrigger] */
    val trigger = jsonb<String>("trigger", { it }, { it })

    /** [SurveyFollowUpAction] */
    val action = jsonb<String>("action", { it }, { it })
    override val primaryKey = PrimaryKey(id, name = "\"SurveyFollowUp_pkey\"")
}
