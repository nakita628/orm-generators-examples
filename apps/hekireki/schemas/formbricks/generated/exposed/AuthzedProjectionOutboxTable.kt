package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit
import java.util.UUID

/**
 * Durable PostgreSQL queue for rebuilding SpiceDB relationships after an authorization source changes.
 * Source-table triggers insert these rows in the same transaction as the mutation. Identifiers are
 * operationally sensitive and must never be included in logs, metrics, or CLI output.
 *
 * This model declares NO indexes on purpose. All three live in
 * `migration/20260818120000_add_authzed_projection_outbox/migration.sql`, because each one is
 * partial (`WHERE "processedAt" IS NULL AND ...`) and Prisma's `@@index` has no `where`. Declaring
 * approximate non-partial copies here would be worse than declaring none: `prisma db push` would
 * see a name match with a different definition, drop and recreate them without the predicate, and
 * the migration's `CREATE INDEX IF NOT EXISTS` would then skip them — leaving a silently wrong
 * shape that carries a week of delivered rows on every hot-path lookup. Deployments are unaffected
 * either way (`prisma migrate deploy` never reads this file); a dev who runs `db push` restores
 * them by rerunning the migration, which is the same contract its triggers already live under.
 */
object AuthzedProjectionOutboxTable : IdTable<String>("\"AuthzedProjectionOutbox\"") {
    override val id = text("id").clientDefault { UUID.randomUUID().toString() }.entityId()
    val targetType = text("targetType")
    val primaryId = text("primaryId")
    val secondaryId = text("secondaryId").nullable()
    val isRevocation = bool("isRevocation").default(false)

    /** Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried. */
    val attempts = integer("attempts").default(0)

    /**
     * Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
     * that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
     */
    val permanentFailures = integer("permanentFailures").default(0)
    val availableAt = pgTimestamp("availableAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val leasedAt = pgTimestamp("leasedAt", 3).nullable()
    val leaseExpiresAt = pgTimestamp("leaseExpiresAt", 3).nullable()
    val leaseOwner = text("leaseOwner").nullable()
    val processedAt = pgTimestamp("processedAt", 3).nullable()
    val deadLetteredAt = pgTimestamp("deadLetteredAt", 3).nullable()
    val lastAttemptAt = pgTimestamp("lastAttemptAt", 3).nullable()
    val lastErrorCode = text("lastErrorCode").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AuthzedProjectionOutbox_pkey\"")
}
