package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

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
class AuthzedProjectionOutboxEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuthzedProjectionOutboxEntity>(AuthzedProjectionOutboxTable)

    var targetType by AuthzedProjectionOutboxTable.targetType
    var primaryId by AuthzedProjectionOutboxTable.primaryId
    var secondaryId by AuthzedProjectionOutboxTable.secondaryId
    var isRevocation by AuthzedProjectionOutboxTable.isRevocation

    /** Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried. */
    var attempts by AuthzedProjectionOutboxTable.attempts

    /**
     * Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
     * that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
     */
    var permanentFailures by AuthzedProjectionOutboxTable.permanentFailures
    var availableAt by AuthzedProjectionOutboxTable.availableAt
    var leasedAt by AuthzedProjectionOutboxTable.leasedAt
    var leaseExpiresAt by AuthzedProjectionOutboxTable.leaseExpiresAt
    var leaseOwner by AuthzedProjectionOutboxTable.leaseOwner
    var processedAt by AuthzedProjectionOutboxTable.processedAt
    var deadLetteredAt by AuthzedProjectionOutboxTable.deadLetteredAt
    var lastAttemptAt by AuthzedProjectionOutboxTable.lastAttemptAt
    var lastErrorCode by AuthzedProjectionOutboxTable.lastErrorCode
    var createdAt by AuthzedProjectionOutboxTable.createdAt
    var updatedAt by AuthzedProjectionOutboxTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuthzedProjectionOutboxTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
