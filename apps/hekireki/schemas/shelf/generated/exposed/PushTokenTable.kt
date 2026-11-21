package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One Expo push token per installed companion app. A token identifies a
 * device, not a person: when someone else signs in on the same phone, the
 * row moves to them. Pushes go only to rows where `disabledAt` is null.
 * @see apps/webapp/app/modules/push/service.server.ts
 */
object PushTokenTable : IdTable<String>("\"PushToken\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val token = text("token").uniqueIndex("\"PushToken_token_key\"")
    val platform = pgEnum("platform", "\"DevicePlatform\"", DevicePlatform.entries, DevicePlatform::dbName)
    val deviceName = text("deviceName").nullable()
    val appVersion = text("appVersion").nullable()
    val lastSeenAt = pgTimestamp("lastSeenAt", 3)
    val disabledAt = pgTimestamp("disabledAt", 3).nullable()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PushToken_userId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"PushToken_pkey\"")

    init {
        index("\"PushToken_userId_idx\"", false, userId)
    }
}
