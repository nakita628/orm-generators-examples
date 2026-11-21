#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Rebuildable per-turn operational render document.
/// </summary>
public partial class LangyConversationTurnProjection
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string ConversationId { get; set; } = null!;

    public string TurnId { get; set; } = null!;

    /// <summary>
    /// LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
    /// is a PROJECTION — one writer (the fold), rebuildable from the event log,
    /// and nothing queries or indexes on it. An enum bought no selectivity here,
    /// only typo protection, and charged a migration for every new status that
    /// Postgres then could not remove. The one definition lives in
    /// `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
    /// </summary>
    public string Status { get; set; } = null!;

    /// <summary>
    /// LangyMessagePart[] — the user question that opened the turn.
    /// </summary>
    public string QuestionParts { get; set; } = null!;

    /// <summary>
    /// LangyMessagePart[] — the agent's whole final answer.
    /// </summary>
    public string AnswerParts { get; set; } = null!;

    /// <summary>
    /// LangyTurnToolCall[] — tool-call lifecycle audit in initiation order.
    /// </summary>
    public string ToolCalls { get; set; } = null!;

    /// <summary>
    /// LangyPlanItemData[] | null — null means the turn never kept a plan.
    /// </summary>
    public string? Plan { get; set; }

    public string? Error { get; set; }

    public double? StartedAt { get; set; }

    public double? EndedAt { get; set; }

    public double CreatedAt { get; set; }

    public double UpdatedAt { get; set; }

    public double OccurredAt { get; set; }

    public double AcceptedAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public string ProjectionVersion { get; set; } = null!;
}
