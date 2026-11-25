package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** One person's digest for one ISO week, accepted by our own mailer. The id is `userId:isoWeek`. */
object DigestSendTable : IdTable<String>("\"DigestSend\"") {
    override val id = text("id").entityId()
    val userId = text("userId")
    val organizationId = text("organizationId")
    val isoWeek = text("isoWeek")
    val templateId = text("templateId")
    val sentAt = pgTimestamp("sentAt", 3)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DigestSend_pkey\"")

    init {
        index("\"DigestSend_userId_idx\"", false, userId)
    }
}
