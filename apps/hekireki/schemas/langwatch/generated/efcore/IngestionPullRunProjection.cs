#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Rebuildable per-source pull checkpoint and operator-facing run status.
/// </summary>
public partial class IngestionPullRunProjection
{
    public string Id { get; set; } = null!;

    public string SourceId { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public bool Enabled { get; set; } = false;

    public string? Cron { get; set; }

    public string? Cursor { get; set; }

    public double? LastRunAt { get; set; }

    public string? LastRunOutcome { get; set; }

    public int LastRunEventCount { get; set; } = 0;

    public string? LastRunError { get; set; }

    public string? LastRunErrorCode { get; set; }

    public int ConsecutiveErrors { get; set; } = 0;

    /// <summary>
    /// When a run last completed without erroring. Backfilled by replay, which
    /// is why the projection version moved with this column.
    /// </summary>
    public double? LastSuccessAt { get; set; }

    /// <summary>
    /// Identity of the run whose outcome this row reflects. Outcomes from a
    /// superseded run arrive with a smaller value and are ignored, so a late
    /// completion cannot regress Cursor or the status fields.
    /// </summary>
    public double? LastRunScheduledFor { get; set; }

    /// <summary>
    /// ADR-128 - how far the last run read, in epoch milliseconds, and whether
    /// it reached the end. The fold writes both; the IngestionSource row mirrors
    /// them for the screens. Null read-through means no run has reported one.
    /// Null completeness reads as unknown, never as either answer, which is what
    /// every row written before this existed holds.
    ///
    /// Nullable on purpose, so replay is not required to make this migration
    /// safe: a row that has never seen a run carrying these stays null and every
    /// reader already handles unknown.
    /// Spec: specs/governance/ingestion-source-health.feature
    /// </summary>
    public double? LastReadThroughAt { get; set; }

    public string? LastRunCompleteness { get; set; }

    /// <summary>
    /// The last time somebody asked this source to enumerate its agents, and how
    /// that ask ended. A listing is a different scope on the same connection as
    /// the pull above: a credential that cannot enumerate a directory still pulls
    /// cost perfectly, so no health, coverage or completeness reader may treat a
    /// refusal here as the source failing to pull.
    ///
    /// Count is set ONLY when the outcome is listed, and stays null on a refusal.
    /// Zero is a real answer meaning the provider returned an empty list, so a
    /// zero default would manufacture that answer for a source that refused.
    /// Reason is set only on a refusal and is a plain string, never an enum and
    /// never a foreign key, because the log outlives the vocabulary. Status is
    /// nullable even when refused, because a refusal we generated ourselves has
    /// no HTTP status, and it is never rendered to a customer.
    ///
    /// The two `...ListingAt` columns hold the instant the ASK was accepted, not
    /// the instant its answer was recorded, and the fold orders outcomes by them
    /// so a superseded listing cannot overwrite the one that replaced it. The
    /// name predates that job and is kept because renaming a column is a
    /// migration; the fold's own field docs are the statement of record.
    /// </summary>
    public double? LastAgentsListingAt { get; set; }

    public string? LastAgentsListingOutcome { get; set; }

    public int? LastAgentsListingCount { get; set; }

    public string? LastAgentsListingReason { get; set; }

    public int? LastAgentsListingStatus { get; set; }

    /// <summary>
    /// The same for people, in its own set of columns rather than sharing the
    /// agents set with a kind discriminator. A source can be refused for people
    /// and fine for agents, and one shared set holds only the most recent listing
    /// of either kind, so syncing people would erase an agents refusal and the
    /// agents page would go back to not knowing after briefly knowing.
    ///
    /// LastPeopleDirectoryCount is everyone the provider's directory named.
    /// LastPeopleWithheldCount is how many of those this deployment does not hold
    /// because erasure suppression removed them: a SUBSET of the directory count,
    /// never an addition to it. Adding the two counts the same people twice.
    /// Subtracting is the valid arithmetic, and the surviving total is that
    /// subtraction rather than a third column that could disagree.
    ///
    /// The withheld count is safe to show as a CURRENT figure and never as a
    /// series: it moving from zero to one at a known moment says an erasure
    /// happened then, which on a small tenant identifies the person as surely as
    /// a name would. So no trend line, no history drawer, no per-run export, and
    /// never beside a per-person list.
    /// </summary>
    public double? LastPeopleListingAt { get; set; }

    public string? LastPeopleListingOutcome { get; set; }

    public int? LastPeopleDirectoryCount { get; set; }

    public int? LastPeopleWithheldCount { get; set; }

    public string? LastPeopleListingReason { get; set; }

    public int? LastPeopleListingStatus { get; set; }

    public double CreatedAt { get; set; }

    public double UpdatedAt { get; set; }

    public double OccurredAt { get; set; }

    public double AcceptedAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public string ProjectionVersion { get; set; } = null!;
}
