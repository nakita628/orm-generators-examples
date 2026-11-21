package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines the possible attributes that can be assigned to contacts.
 * Acts as a schema for contact attributes within a workspace.
 *
 * @property id - Unique identifier for the attribute key
 * @property isUnique - Whether the attribute must have unique values across contacts
 * @property key - The attribute identifier used in the system
 * @property name - Display name for the attribute
 * @property type - Whether this is a default or custom attribute
 * @property dataType - The data type of the attribute (string, number, date)
 * @property workspace - The workspace this attribute belongs to
 */
object ContactAttributeKeyTable : IdTable<String>("\"ContactAttributeKey\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val isUnique = bool("isUnique").default(false)
    val key = text("key")
    val name = text("name").nullable()
    val description = text("description").nullable()
    val type =
        pgEnum("type", "\"ContactAttributeType\"", ContactAttributeType.entries, ContactAttributeType::dbName).default(
            ContactAttributeType.CUSTOM,
        )
    val dataType =
        pgEnum("dataType", "\"ContactAttributeDataType\"", ContactAttributeDataType.entries, ContactAttributeDataType::dbName).default(
            ContactAttributeDataType.STRING,
        )
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ContactAttributeKey_workspaceId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"ContactAttributeKey_pkey\"")

    init {
        index("\"ContactAttributeKey_workspaceId_created_at_idx\"", false, workspaceId, createdAt)
        uniqueIndex("\"ContactAttributeKey_key_workspaceId_key\"", key, workspaceId)
    }
}
