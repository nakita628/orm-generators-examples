package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Why an outbox delivery failed, one row per FAILED attempt
 * (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
 * table grows with trouble, not with traffic. Rows die with their message:
 * the cascade delete is the retention story, riding the existing sweep.
 */
object ProcessManagerOutboxAttemptTable : IdTable<String>("\"ProcessManagerOutboxAttempt\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val outboxId = text("outboxId")
    val projectId = text("projectId")

    /** 1-based attempt number, as the dispatcher counts it. */
    val attempt = integer("attempt")
    val occurredAt = pgTimestamptz("occurredAt", 3)
    val outcome =
        pgEnum(
            "outcome",
            "\"ProcessManagerOutboxAttemptOutcome\"",
            ProcessManagerOutboxAttemptOutcome.entries,
            ProcessManagerOutboxAttemptOutcome::dbName,
        )
    val errorType = text("errorType")

    /**
     * The safe failure diagnostic the dispatcher already computes for its
     * span — never a raw provider body.
     */
    val errorMessage = text("errorMessage")
    val retryAfterMs = integer("retryAfterMs").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"ProcessManagerOutboxAttempt_pkey\"")

    init {
        index("\"ProcessManagerOutboxAttempt_outboxId_attempt_idx\"", false, outboxId, attempt)
    }
}
