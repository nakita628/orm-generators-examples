package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object PlaygroundConversationTable : IdTable<String>("\"PlaygroundConversation\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()

    /** The chat session ID used by the transport */
    val chatId = text("chatId")

    /** User-editable conversation title (auto-generated from first message) */
    val title = text("title").default("New conversation")

    /** Which agent this conversation is with */
    val agentSlug = text("agentSlug")

    /** The current active run backing this conversation (null if no run yet) */
    val runId =
        optReference(
            "runId",
            TaskRunTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PlaygroundConversation_runId_fkey\"",
        )

    /** The client data JSON used for this conversation */
    val clientData = jsonb<String>("clientData", { it }, { it }).nullable()

    /** Accumulated UIMessages from completed turns (for resume without stream replay) */
    val messages = jsonb<String>("messages", { it }, { it }).nullable()

    /** Last SSE event ID — resume from this position to avoid replaying old turns */
    val lastEventId = text("lastEventId").nullable()
    val projectId =
        reference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PlaygroundConversation_projectId_fkey\"",
        )
    val runtimeEnvironmentId =
        reference(
            "runtimeEnvironmentId",
            RuntimeEnvironmentTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PlaygroundConversation_runtimeEnvironmentId_fkey\"",
        )

    /** The user who started this conversation */
    val userId = text("userId")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"PlaygroundConversation_pkey\"")

    init {
        index(
            "PlaygroundConversation_runtimeEnvironmentId_agentSlug_updat_idx",
            false,
            functions =
                listOf(
                    indexExpression("\"runtimeEnvironmentId\""),
                    indexExpression("\"agentSlug\""),
                    indexExpression("\"updatedAt\" DESC"),
                ),
        )
        index("\"PlaygroundConversation_userId_runtimeEnvironmentId_idx\"", false, userId, runtimeEnvironmentId)
        uniqueIndex("\"PlaygroundConversation_chatId_runtimeEnvironmentId_key\"", chatId, runtimeEnvironmentId)
    }
}
