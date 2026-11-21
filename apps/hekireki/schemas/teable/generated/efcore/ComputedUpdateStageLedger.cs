#nullable enable

namespace Teable.Models;

/// <summary>
/// Durable per-stage state for budget-staged computed updates, keyed by the
/// continuation chain's root task id (its scope). Rows are written once per
/// record and never copied between continuation tasks:
/// - kind 'excluded': processed-target exclusion ledger for partial batches;
/// - kind 'frontier': seq-ordered queue of sources whose outgoing propagation
/// is not finished (self-referential generations + migrated explicit seeds);
/// - kind 'consumed': retired frontier sources preserved for deferred edge
/// chunks, handed to the continuation as seeds at stage completion.
/// Cleared when the chain's stage completes or the chain dead-letters.
/// </summary>
public partial class ComputedUpdateStageLedger
{
    public string ScopeId { get; set; } = null!;

    public string Kind { get; set; } = null!;

    public string TableId { get; set; } = null!;

    public string RecordId { get; set; } = null!;

    public long Seq { get; set; } = 0L;
}
