#nullable enable

namespace LangWatch.Models;

/// <summary>
/// One insight in its owner's inbox in a project, folded from its `insight` aggregate
/// (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
/// milliseconds.
/// </summary>
public partial class InsightProjection
{
    /// <summary>
    /// The insight id, which is also its aggregate id.
    /// </summary>
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string Title { get; set; } = null!;

    public string Body { get; set; } = null!;

    /// <summary>
    /// bad | watch | good
    /// </summary>
    public string Tone { get; set; } = null!;

    public string? Topic { get; set; }

    public int ValidDays { get; set; }

    public string? Lwql { get; set; }

    public string? SourceConversationId { get; set; }

    public string? SourceMessageId { get; set; }

    public string? FiledByUserId { get; set; }

    public double FiledAt { get; set; }

    /// <summary>
    /// Set by a later run that found the insight still true.
    /// </summary>
    public double? RenewedAt { get; set; }

    public double CreatedAt { get; set; }

    public double UpdatedAt { get; set; }

    public double OccurredAt { get; set; }

    public double AcceptedAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public string ProjectionVersion { get; set; } = null!;

    /// <summary>
    /// The board and widget it came from. A pointer only, never a relation: the row reads the
    /// same once either is deleted, so the names are kept as filed.
    /// </summary>
    public string? BoardId { get; set; }

    public string? BoardName { get; set; }

    public string? WidgetId { get; set; }

    public string? WidgetName { get; set; }

    /// <summary>
    /// chat | run
    /// </summary>
    public string FiledVia { get; set; } = "chat";

    /// <summary>
    /// The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step.
    /// </summary>
    public double? ReplayStart { get; set; }

    public double? ReplayEnd { get; set; }

    public int? ReplayGranularitySeconds { get; set; }

    /// <summary>
    /// The values in force when it was filed: `{ period, parameters }`.
    /// </summary>
    public string? ReplayContext { get; set; }

    /// <summary>
    /// Whose insight it is: the one person who reads it. NULL on a row folded before the
    /// column existed, which is owned by whoever filed it (`filedByUserId`).
    /// </summary>
    public string? OwnerUserId { get; set; }

    /// <summary>
    /// The stable key a daily run matches a repeated finding by, so it renews the insight
    /// and files no twin. NULL for an insight saved from a chat and on a row folded before
    /// the column existed.
    /// </summary>
    public string? FindingKey { get; set; }
}
