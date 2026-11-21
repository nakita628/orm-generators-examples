package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/** The languages a creative's texts are written in. Mirrors `SurveyLanguage`; exactly one is the default. */
object CreativeLanguageTable : CompositeIdTable("\"CreativeLanguage\"") {
    val languageId =
        reference(
            "languageId",
            LanguageTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CreativeLanguage_languageId_fkey\"",
        )
    val creativeId =
        reference(
            "creativeId",
            CreativeTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CreativeLanguage_creativeId_fkey\"",
        )
    val default = bool("default").default(false)
    val enabled = bool("enabled").default(true)
    override val primaryKey = PrimaryKey(languageId, creativeId, name = "\"CreativeLanguage_pkey\"")

    init {
        addIdColumn(languageId)
        addIdColumn(creativeId)
        index("\"CreativeLanguage_creativeId_idx\"", false, creativeId)
    }
}
