package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/**
 * Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
 * `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
 * endpoint, which is why the migration backfills one row per existing client.
 * Both sides reference natural keys, matching the other oauth child tables.
 */
object OauthClientResourceTable : IdTable<String>("\"oauthClientResource\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val clientId =
        text("clientId").references(
            OauthClientTable.clientId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"oauthClientResource_clientId_fkey\"",
        )
    val resourceId =
        text("resourceId").references(
            OauthResourceTable.identifier,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"oauthClientResource_resourceId_fkey\"",
        )
    val metadata = jsonb<String>("metadata", { it }, { it }).nullable()
    val createdAt = pgTimestamp("createdAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"oauthClientResource_pkey\"")

    init {
        index("\"oauthClientResource_resourceId_idx\"", false, resourceId)
        uniqueIndex("\"oauthClientResource_clientId_resourceId_key\"", clientId, resourceId)
    }
}
