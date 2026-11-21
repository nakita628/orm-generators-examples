package plunk.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit
import java.util.UUID

object ContactTable : IdTable<String>("contacts") {
    override val id = text("id").clientDefault { UUID.randomUUID().toString() }.entityId()
    val email = text("email")
    val data = jsonb<String>("data", { it }, { it }).nullable()
    val subscribed = bool("subscribed").default(true)

    /**
     * When set, `subscribed = false` is temporary rather than permanent: the snooze
     * sweep resubscribes the contact once this passes. Suppression itself is still
     * carried entirely by `subscribed`, so every send path filters snoozed contacts
     * out without knowing this column exists.
     *
     * Cleared by every other write to `subscribed` — including the bounce and
     * complaint paths, so a suppressed address can never be woken back up.
     */
    val snoozedUntil = pgTimestamp("snoozedUntil", 3).nullable()
    val projectId =
        reference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"contacts_projectId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "contacts_pkey")

    init {
        index("\"contacts_projectId_idx\"", false, projectId)
        index("\"contacts_projectId_subscribed_idx\"", false, projectId, subscribed)
        index("\"contacts_snoozedUntil_idx\"", false, snoozedUntil)
        index("contacts_data_idx", false, functions = listOf(indexExpression("\"data\" jsonb_ops")), indexType = "gin")
        uniqueIndex("\"contacts_projectId_email_key\"", projectId, email)
    }
}
