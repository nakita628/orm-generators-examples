package teable.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

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
class ComputedUpdateStageLedgerEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<ComputedUpdateStageLedgerEntity>(ComputedUpdateStageLedgerTable)

    val scopeId by ComputedUpdateStageLedgerTable.scopeId
    val kind by ComputedUpdateStageLedgerTable.kind
    val tableId by ComputedUpdateStageLedgerTable.tableId
    val recordId by ComputedUpdateStageLedgerTable.recordId
    var seq by ComputedUpdateStageLedgerTable.seq
}
