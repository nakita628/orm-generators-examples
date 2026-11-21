package teable.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Durable per-stage state for budget-staged computed updates, keyed by the
 * continuation chain's root task id (its scope). Rows are written once per
 * record and never copied between continuation tasks:
 * - kind 'excluded': processed-target exclusion ledger for partial batches;
 * - kind 'frontier': seq-ordered queue of sources whose outgoing propagation
 * is not finished (self-referential generations + migrated explicit seeds);
 * - kind 'consumed': retired frontier sources preserved for deferred edge
 * chunks, handed to the continuation as seeds at stage completion.
 * Cleared when the chain's stage completes or the chain dead-letters.
 */
object ComputedUpdateStageLedgerTable : CompositeIdTable("computed_update_stage_ledger") {
    val scopeId = text("scope_id").entityId()
    val kind = text("kind").entityId()
    val tableId = text("table_id").entityId()
    val recordId = text("record_id").entityId()
    val seq = long("seq").default(0L)
    override val primaryKey = PrimaryKey(scopeId, kind, tableId, recordId, name = "computed_update_stage_ledger_pkey")

    init {
        index("computed_update_stage_ledger_scope_id_kind_seq_idx", false, scopeId, kind, seq)
    }
}
