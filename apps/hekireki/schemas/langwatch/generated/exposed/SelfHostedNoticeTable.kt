package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
 * report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
 * tenant: every install receives every published notice that has not expired.
 */
object SelfHostedNoticeTable : IdTable<String>("\"SelfHostedNotice\"") {
    override val id = text("id").entityId()

    /** `info`, `warning` or `security`. */
    val severity = text("severity")
    val title = text("title")
    val body = text("body")

    /** An https link to read more; null where the notice names none. */
    val url = text("url").nullable()

    /** The release that fixes what the notice is about; an install on it or newer hides it. */
    val fixedIn = text("fixedIn").nullable()

    /** `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent. */
    val status = text("status")
    val expiresAt = pgTimestamp("expiresAt", 3).nullable()
    val publishedAt = pgTimestamp("publishedAt", 3).nullable()

    /** The staff member who published it, by user id. */
    val publishedBy = text("publishedBy").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SelfHostedNotice_pkey\"")
}
