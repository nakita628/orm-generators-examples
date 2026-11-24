package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object OAuthClientTable : IdTable<String>("\"OAuthClient\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val kind = pgEnum("kind", "\"OAuthClientKind\"", OAuthClientKind.entries, OAuthClientKind::dbName)
    val clientId = text("clientId").uniqueIndex("\"OAuthClient_clientId_key\"")
    val metadataUrl = text("metadataUrl").nullable()
    val displayName = text("displayName").nullable()
    val selfReportedName = text("selfReportedName").nullable()
    val redirectUris = array("redirectUris", TextColumnType()).nullable()
    val dedupeKey = text("dedupeKey").nullable().uniqueIndex("\"OAuthClient_dedupeKey_key\"")
    val metadataExpiresAt = pgTimestamp("metadataExpiresAt", 3).nullable()
    val lastUsedAt = pgTimestamp("lastUsedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"OAuthClient_pkey\"")

    init {
        index("\"OAuthClient_kind_lastUsedAt_idx\"", false, kind, lastUsedAt)
    }
}
