package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** A person's opt-out of the digest. No row, or no `unsubscribedAt`: subscribed. */
object DigestSubscriptionTable : IdTable<String>("\"DigestSubscription\"") {
    override val id = text("userId").entityId()
    val unsubscribedAt = pgTimestamp("unsubscribedAt", 3).nullable()

    /** The send whose unsubscribe link or one-click header the person used. */
    val unsubscribedFromSendId = text("unsubscribedFromSendId").nullable()
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DigestSubscription_pkey\"")
}
