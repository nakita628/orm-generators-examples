package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
 * report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
 * tenant: every install receives every published notice that has not expired.
 */
class SelfHostedNoticeEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SelfHostedNoticeEntity>(SelfHostedNoticeTable)

    /** `info`, `warning` or `security`. */
    var severity by SelfHostedNoticeTable.severity
    var title by SelfHostedNoticeTable.title
    var body by SelfHostedNoticeTable.body

    /** An https link to read more; null where the notice names none. */
    var url by SelfHostedNoticeTable.url

    /** The release that fixes what the notice is about; an install on it or newer hides it. */
    var fixedIn by SelfHostedNoticeTable.fixedIn

    /** `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent. */
    var status by SelfHostedNoticeTable.status
    var expiresAt by SelfHostedNoticeTable.expiresAt
    var publishedAt by SelfHostedNoticeTable.publishedAt

    /** The staff member who published it, by user id. */
    var publishedBy by SelfHostedNoticeTable.publishedBy
    var createdAt by SelfHostedNoticeTable.createdAt
    var updatedAt by SelfHostedNoticeTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SelfHostedNoticeTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
