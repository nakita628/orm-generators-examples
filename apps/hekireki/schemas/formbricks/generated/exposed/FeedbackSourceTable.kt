package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A feedback source — connects external data into Unify Feedback for record creation.
 * Naming note: "FeedbackSource" is the Unify Feedback domain term; the word "Connector"
 * is reserved for the Workflows domain (see docs/development/standards/practices/naming-conventions.mdx).
 *
 * @property id - Unique identifier for the feedback source
 * @property name - Display name for the feedback source
 * @property type - Type of feedback source (formbricks_survey, csv, …)
 * @property status - Current state of the feedback source (active, paused, error)
 * @property workspace - The workspace this feedback source belongs to
 * @property feedbackDirectory - Hub directory the feedback source writes into
 * @property formbricksMappings - Element mappings for Formbricks feedback sources
 * @property fieldMappings - Field mappings for other feedback source types
 */
object FeedbackSourceTable : IdTable<String>("\"FeedbackSource\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val name = text("name")
    val type = pgEnum("type", "\"FeedbackSourceType\"", FeedbackSourceType.entries, FeedbackSourceType::dbName)
    val status =
        pgEnum("status", "\"FeedbackSourceStatus\"", FeedbackSourceStatus.entries, FeedbackSourceStatus::dbName).default(
            FeedbackSourceStatus.ACTIVE,
        )

    /**
     * Defaults to completedOnly so a new source agrees with the live ingestion path, which only
     * fires on responseFinished. Existing rows take the default; nothing is re-imported.
     */
    val importMode =
        pgEnum("import_mode", "\"FeedbackSourceImportMode\"", FeedbackSourceImportMode.entries, FeedbackSourceImportMode::dbName).default(
            FeedbackSourceImportMode.COMPLETED_ONLY,
        )
    val elementScope =
        pgEnum("elementScope", "\"FeedbackSourceElementScope\"", FeedbackSourceElementScope.entries, FeedbackSourceElementScope::dbName)
            .default(FeedbackSourceElementScope.SPECIFIC)
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"FeedbackSource_workspaceId_fkey\"",
        )
    val feedbackDirectoryId =
        reference(
            "feedbackDirectoryId",
            FeedbackDirectoryTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"FeedbackSource_feedbackDirectoryId_fkey\"",
        )
    val lastSyncAt = pgTimestamp("last_sync_at", 3).nullable()
    val createdBy =
        optReference(
            "created_by",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"FeedbackSource_created_by_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"FeedbackSource_pkey\"")

    init {
        index("\"FeedbackSource_type_idx\"", false, type)
        index("\"FeedbackSource_feedbackDirectoryId_workspaceId_idx\"", false, feedbackDirectoryId, workspaceId)
        uniqueIndex("\"FeedbackSource_id_workspaceId_key\"", id, workspaceId)
        uniqueIndex("\"FeedbackSource_workspaceId_name_key\"", workspaceId, name)
        foreignKey(
            feedbackDirectoryId to FeedbackDirectoryWorkspaceTable.feedbackDirectoryId,
            workspaceId to FeedbackDirectoryWorkspaceTable.workspaceId,
            onDelete = ReferenceOption.NO_ACTION,
            onUpdate = ReferenceOption.CASCADE,
            name = "\"FeedbackSource_feedbackDirectoryId_workspaceId_fkey\"",
        )
    }
}
