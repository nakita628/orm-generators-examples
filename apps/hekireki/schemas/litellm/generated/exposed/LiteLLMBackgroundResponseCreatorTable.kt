package litellm.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object LiteLLMBackgroundResponseCreatorTable : IdTable<String>("\"LiteLLM_BackgroundResponseCreator\"") {
    override val id = text("response_id").entityId()
    val upstreamUrl = text("upstream_url")
    val createContext = jsonb<String>("create_context", { it }, { it })
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val pollLeaseExpiresAt = pgTimestamp("poll_lease_expires_at", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"LiteLLM_BackgroundResponseCreator_pkey\"")

    init {
        index("idx_background_response_creator_created_at", false, createdAt)
    }
}
