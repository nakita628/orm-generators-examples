package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a quota configuration for a survey.
 * Defines response limits and conditions for quota management.
 *
 * @property id - Unique identifier for the quota
 * @property survey - The survey this quota belongs to
 * @property name - Display name for the quota
 * @property limit - Maximum number of responses allowed
 * @property conditions - JSON object containing quota conditions
 * @property action - Action to take when quota is reached
 * @property endingCardId - Optional ending card to show when quota is reached
 */
object SurveyQuotaTable : IdTable<String>("\"SurveyQuota\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val surveyId =
        reference(
            "surveyId",
            SurveyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"SurveyQuota_surveyId_fkey\"",
        )
    val name = text("name")
    val limit = integer("limit")

    /** [SurveyQuotaLogic] */
    val logic = jsonb<String>("logic", { it }, { it }).default("{}")
    val action = pgEnum("action", "\"SurveyQuotaAction\"", SurveyQuotaAction.entries, SurveyQuotaAction::dbName)
    val endingCardId = text("endingCardId").nullable()
    val countPartialSubmissions = bool("countPartialSubmissions").default(false)
    override val primaryKey = PrimaryKey(id, name = "\"SurveyQuota_pkey\"")

    init {
        uniqueIndex("\"SurveyQuota_surveyId_name_key\"", surveyId, name)
    }
}
