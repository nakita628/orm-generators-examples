#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A name-text seen on a provider's cost or audit rows: "the provider told us
/// someone called m.silva exists". One row per provider-scoped identity.
///
/// This is the unit of GDPR erasure (ADR-128 §9), and deliberately so: most of
/// these people have no LangWatch login at all - contractors, seat holders who
/// never signed in, anyone whose email a provider put on a cost row - so there
/// is no platform user to key on and a user-deletion event can never fire for
/// them.
/// </summary>
public partial class DiscoveredPerson
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// "anthropic" | "openai" | "databricks" | ... - the provider that named them.
    /// </summary>
    public string Provider { get; set; } = null!;

    /// <summary>
    /// What the provider calls them, verbatim. The join key onto the money rows'
    /// RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
    /// the row addressable for spend attribution without keeping the identifier.
    /// </summary>
    public string RawActorId { get; set; } = null!;

    /// <summary>
    /// Name or email text as seen. Pseudonymized alongside rawActorId on erasure.
    /// </summary>
    public string DisplayText { get; set; } = null!;

    /// <summary>
    /// person | service_account - determined deterministically per ADR-128 §10
    /// (a bare UUID in Databricks `executed_by` is a service account, never an
    /// agent), never guessed from the shape of a name.
    /// </summary>
    public string Kind { get; set; } = null!;

    /// <summary>
    /// The department the provider's directory filed this person under, verbatim.
    /// An OBSERVATION, not the `Department` entity: it is free text a tenant owns
    /// on their side, it exists for people who hold no LangWatch account, and
    /// nothing rolls spend up by it. Spend attribution stays on `Department`, via
    /// the membership the match engine proves (directoryDepartmentSync.service).
    /// Null means no directory has named one - never "no department", because a
    /// blank field in the tenant's directory and a person no directory lists are
    /// the same null here.
    /// </summary>
    public string? Department { get; set; }

    public DateTime FirstSeenAt { get; set; }

    public DateTime LastSeenAt { get; set; }

    /// <summary>
    /// Set when this row's identity has been erased. The row survives so its
    /// spend stays attributable to *someone*; the identifier does not.
    /// </summary>
    public DateTime? ErasedAt { get; set; }

    /// <summary>
    /// Set alongside `erasedAt` and cleared once the daily cost rows have been
    /// removed and their rebuild requested. Non-null means the ClickHouse half
    /// of the erasure did not finish — a crash between the delete and the
    /// replay leaves days short by the erased amount, and without this marker a
    /// re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
    /// did no work. The re-run resumes at that step instead.
    /// </summary>
    public DateTime? MoneyRowsPendingAt { get; set; }

    /// <summary>
    /// The earliest day to rebuild from, recorded BEFORE the rows are deleted.
    /// Once they are gone there is nothing left to ask which days they were on,
    /// so a resume with no record of this could only replay everything or
    /// nothing. Null means nothing was rebuildable — either no day carried the
    /// identifier, or every one of them is older than the event log keeps.
    /// </summary>
    public string? MoneyRebuildSince { get; set; }

    /// <summary>
    /// ADR-128 §12 conflict rule: automatic linking has stopped for this
    /// identity because evidence pointed at two candidates, or contradicted an
    /// existing link. Stored HERE rather than on the suggestion rows, because a
    /// halt on auto-linking that a recompute or a restart clears is not a halt.
    /// </summary>
    public DateTime? SuspendedAt { get; set; }

    /// <summary>
    /// What tripped the suspension, for the human who reviews it.
    /// </summary>
    public string? SuspendedReason { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
