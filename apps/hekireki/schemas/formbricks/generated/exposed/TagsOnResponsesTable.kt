package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Junction table linking tags to responses.
 * Enables many-to-many relationship between tags and responses.
 *
 * @property response - The tagged response
 * @property tag - The tag applied to the response
 */
object TagsOnResponsesTable : CompositeIdTable("\"TagsOnResponses\"") {
    val responseId =
        reference(
            "responseId",
            ResponseTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TagsOnResponses_responseId_fkey\"",
        )
    val tagId =
        reference(
            "tagId",
            TagTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TagsOnResponses_tagId_fkey\"",
        )
    override val primaryKey = PrimaryKey(responseId, tagId, name = "\"TagsOnResponses_pkey\"")

    init {
        addIdColumn(responseId)
        addIdColumn(tagId)
    }
}
