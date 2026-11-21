#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Idempotency marker for a source event consumed by a process definition.
/// </summary>
public partial class ProcessManagerInbox
{
    public string Id { get; set; } = null!;

    public string ProcessName { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string ProcessKey { get; set; } = null!;

    public string TenantId { get; set; } = null!;

    /// <summary>
    /// The raw source event id, kept UNINDEXED for diagnostics. A caller composes
    /// this and its length is that domain's business, so it must never carry a
    /// btree constraint — see `sourceEventKey`.
    /// </summary>
    public string SourceEventId { get; set; } = null!;

    /// <summary>
    /// The fixed-width digest of `sourceEventId` the uniqueness is enforced on,
    /// derived by the store (`stores/inboxKey.ts`). Postgres cannot index a btree
    /// row past ~2704 bytes; keying on a digest makes the constraint the same
    /// width no matter what any pipeline concatenates into an idempotency key.
    ///
    /// Nullable ONLY for the rolling deploy that introduces it — pods running the
    /// previous build write rows without it, and a NOT NULL would fail their
    /// commits outright. The store always supplies it.
    ///
    /// The trade, stated plainly: a NULL is not constrained by the unique index
    /// (btree treats NULLs as distinct), so while an old writer is serving, the
    /// one case that loses its guard is the same source event reaching two
    /// different process keys at once — the store's own read-check plus its
    /// per-processKey advisory lock cover everything else. A follow-up migration
    /// adds NOT NULL once the fleet has cycled and closes it.
    /// </summary>
    public string? SourceEventKey { get; set; }

    public DateTime ConsumedAt { get; set; }
}
