package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Why an outbox delivery failed, one row per FAILED attempt
 * (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
 * table grows with trouble, not with traffic. Rows die with their message:
 * the cascade delete is the retention story, riding the existing sweep.
 */
class ProcessManagerOutboxAttemptEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ProcessManagerOutboxAttemptEntity>(ProcessManagerOutboxAttemptTable)

    var outboxId by ProcessManagerOutboxAttemptTable.outboxId
    var projectId by ProcessManagerOutboxAttemptTable.projectId

    /** 1-based attempt number, as the dispatcher counts it. */
    var attempt by ProcessManagerOutboxAttemptTable.attempt
    var occurredAt by ProcessManagerOutboxAttemptTable.occurredAt
    var outcome by ProcessManagerOutboxAttemptTable.outcome
    var errorType by ProcessManagerOutboxAttemptTable.errorType

    /**
     * The safe failure diagnostic the dispatcher already computes for its
     * span — never a raw provider body.
     */
    var errorMessage by ProcessManagerOutboxAttemptTable.errorMessage
    var retryAfterMs by ProcessManagerOutboxAttemptTable.retryAfterMs
}
