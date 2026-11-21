package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a webhook endpoint for receiving survey-related events.
 * Webhooks can be configured to receive notifications about response creation, updates, and completion.
 *
 * @property id - Unique identifier for the webhook
 * @property name - Optional display name for the webhook
 * @property url - The endpoint URL where events will be sent
 * @property source - Origin of the webhook (user, zapier, make, etc.)
 * @property workspace - Associated workspace
 * @property triggers - Types of events that trigger this webhook
 * @property surveyIds - List of surveys this webhook is monitoring
 */
object WebhookTable : IdTable<String>("\"Webhook\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name").nullable()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val url = text("url")
    val source1 = pgEnum("source", "\"WebhookSource\"", WebhookSource.entries, WebhookSource::dbName).default(WebhookSource.USER)
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Webhook_workspaceId_fkey\"",
        )
    val triggers =
        array("triggers", PgEnumColumnType("\"PipelineTriggers\"", PipelineTriggers.entries, PipelineTriggers::dbName)).nullable()
    val surveyIds = array("surveyIds", TextColumnType()).nullable()
    val secret = text("secret").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Webhook_pkey\"")

    init {
        index("\"Webhook_workspaceId_idx\"", false, workspaceId)
    }
}
