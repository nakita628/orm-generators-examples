package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/** Idempotency marker for a source event consumed by a process definition. */
object ProcessManagerInboxTable : IdTable<String>("\"ProcessManagerInbox\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val processName = text("processName")
    val projectId = text("projectId")
    val processKey = text("processKey")
    val tenantId = text("tenantId")

    /**
     * The raw source event id, kept UNINDEXED for diagnostics. A caller composes
     * this and its length is that domain's business, so it must never carry a
     * btree constraint — see `sourceEventKey`.
     */
    val sourceEventId = text("sourceEventId")

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
    val sourceEventKey = text("sourceEventKey").nullable()
    val consumedAt = pgTimestamptz("consumedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"ProcessManagerInbox_pkey\"")

    init {
        index("ProcessManagerInbox_processName_projectId_processKey_consum_idx", false, processName, projectId, processKey, consumedAt)
        index("\"ProcessManagerInbox_consumedAt_idx\"", false, consumedAt)
        uniqueIndex("\"ProcessManagerInbox_processName_projectId_sourceEventKey_key\"", processName, projectId, sourceEventKey)
    }
}
