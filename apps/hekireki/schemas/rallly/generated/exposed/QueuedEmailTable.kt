package rallly.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * An email waiting to be sent by the queue, one row per recipient. The
 * email is built when it is sent, from the row its subject points at.
 */
object QueuedEmailTable : IdTable<String>("queued_emails") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()

    /**
     * Which handler builds the email. A string rather than an enum so a new
     * kind needs no migration; the claim only takes kinds the build knows.
     */
    val kind = text("kind")

    /**
     * The row the email is about, interpreted by the kind. Not a foreign key:
     * kinds point at different tables, and a subject deleted before the send
     * makes its handler skip the email.
     */
    val subjectId = text("subject_id")

    /**
     * Groups the emails one action queued, so the action can send its own
     * first batch at once and its progress can be read back.
     */
    val batchId = text("batch_id")

    /**
     * The account whose action queued the email. Banning it skips whatever
     * is still pending.
     */
    val userId =
        reference(
            "user_id",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "queued_emails_user_id_fkey",
        )
    val status =
        pgEnum("status", "\"queued_email_status\"", QueuedEmailStatus.entries, QueuedEmailStatus::dbName).default(QueuedEmailStatus.PENDING)

    /** Not claimed before this time. Unset sends at the next run. */
    val sendAfter = pgTimestamp("send_after", 3).nullable()
    val attempts = integer("attempts").default(0)
    val claimedAt = pgTimestamp("claimed_at", 3).nullable()
    val sentAt = pgTimestamp("sent_at", 3).nullable()
    val lastError = text("last_error").nullable()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "queued_emails_pkey")

    init {
        index("queued_emails_status_created_at_idx", false, status, createdAt)
        index("queued_emails_status_updated_at_idx", false, status, updatedAt)
        index("queued_emails_batch_id_idx", false, batchId)
        index("queued_emails_user_id_status_idx", false, userId, status)
    }
}
