package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** Idempotency marker for a source event consumed by a process definition. */
class ProcessManagerInboxEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ProcessManagerInboxEntity>(ProcessManagerInboxTable)

    var processName by ProcessManagerInboxTable.processName
    var projectId by ProcessManagerInboxTable.projectId
    var processKey by ProcessManagerInboxTable.processKey
    var tenantId by ProcessManagerInboxTable.tenantId

    /**
     * The raw source event id, kept UNINDEXED for diagnostics. A caller composes
     * this and its length is that domain's business, so it must never carry a
     * btree constraint — see `sourceEventKey`.
     */
    var sourceEventId by ProcessManagerInboxTable.sourceEventId

    /**
     * The fixed-width digest of `sourceEventId` the uniqueness is enforced on,
     * derived by the store (`stores/inboxKey.ts`). Postgres cannot index a btree
     * row past ~2704 bytes; keying on a digest makes the constraint the same
     * width no matter what any pipeline concatenates into an idempotency key.
     *
     * Nullable ONLY for the rolling deploy that introduces it — pods running the
     * previous build write rows without it, and a NOT NULL would fail their
     * commits outright. The store always supplies it.
     *
     * The trade, stated plainly: a NULL is not constrained by the unique index
     * (btree treats NULLs as distinct), so while an old writer is serving, the
     * one case that loses its guard is the same source event reaching two
     * different process keys at once — the store's own read-check plus its
     * per-processKey advisory lock cover everything else. A follow-up migration
     * adds NOT NULL once the fleet has cycled and closes it.
     */
    var sourceEventKey by ProcessManagerInboxTable.sourceEventKey
    var consumedAt by ProcessManagerInboxTable.consumedAt
}
