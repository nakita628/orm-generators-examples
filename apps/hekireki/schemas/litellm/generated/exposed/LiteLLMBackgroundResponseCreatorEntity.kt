package litellm.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class LiteLLMBackgroundResponseCreatorEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LiteLLMBackgroundResponseCreatorEntity>(LiteLLMBackgroundResponseCreatorTable)

    var upstreamUrl by LiteLLMBackgroundResponseCreatorTable.upstreamUrl
    var createContext by LiteLLMBackgroundResponseCreatorTable.createContext
    var createdAt by LiteLLMBackgroundResponseCreatorTable.createdAt
    var pollLeaseExpiresAt by LiteLLMBackgroundResponseCreatorTable.pollLeaseExpiresAt
}
