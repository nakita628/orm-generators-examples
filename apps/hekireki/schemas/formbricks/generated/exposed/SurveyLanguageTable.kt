package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Links surveys to their supported languages.
 * Manages which languages are available for each survey.
 *
 * @property language - The supported language
 * @property survey - The associated survey
 * @property default - Whether this is the default language
 * @property enabled - Whether this language is currently active
 */
object SurveyLanguageTable : CompositeIdTable("\"SurveyLanguage\"") {
    val languageId =
        reference(
            "languageId",
            LanguageTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"SurveyLanguage_languageId_fkey\"",
        )
    val surveyId =
        reference(
            "surveyId",
            SurveyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"SurveyLanguage_surveyId_fkey\"",
        )
    val default = bool("default").default(false)
    val enabled = bool("enabled").default(true)
    override val primaryKey = PrimaryKey(languageId, surveyId, name = "\"SurveyLanguage_pkey\"")

    init {
        addIdColumn(languageId)
        addIdColumn(surveyId)
        index("\"SurveyLanguage_surveyId_idx\"", false, surveyId)
    }
}
