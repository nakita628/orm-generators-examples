package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One Expo push token per installed companion app. A token identifies a
 * device, not a person: when someone else signs in on the same phone, the
 * row moves to them. Pushes go only to rows where `disabledAt` is null.
 * @see apps/webapp/app/modules/push/service.server.ts
 */
class PushTokenEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PushTokenEntity>(PushTokenTable)

    var token by PushTokenTable.token
    var platform by PushTokenTable.platform
    var deviceName by PushTokenTable.deviceName
    var appVersion by PushTokenTable.appVersion
    var lastSeenAt by PushTokenTable.lastSeenAt
    var disabledAt by PushTokenTable.disabledAt
    var user by UserEntity referencedOn PushTokenTable.userId
    var userId by PushTokenTable.userId
    var createdAt by PushTokenTable.createdAt
    var updatedAt by PushTokenTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PushTokenTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
