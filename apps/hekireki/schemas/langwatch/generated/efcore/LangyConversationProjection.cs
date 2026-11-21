#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Rebuildable operational state folded from Langy's canonical event stream.
/// The source of truth remains the ClickHouse event log.
/// </summary>
public partial class LangyConversationProjection
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string ConversationId { get; set; } = null!;

    /// <summary>
    /// Owner; first-writer-wins from the first message's userId.
    /// </summary>
    public string UserId { get; set; } = null!;

    public string? Title { get; set; }

    /// <summary>
    /// derived | auto | user — governs auto-title precedence (user is sticky).
    /// </summary>
    public LangyProjectionTitleSource TitleSource { get; set; }

    /// <summary>
    /// Who started the conversation: `interactive` for the person in the panel, `run` for a
    /// module acting for them with nobody at the keyboard. A row folded before the column
    /// existed reads as interactive.
    /// </summary>
    public string Origin { get; set; } = "interactive";

    public string Status { get; set; } = null!;

    public bool IsShared { get; set; } = false;

    public double? SharedAt { get; set; }

    public string? SharedById { get; set; }

    public int MessageCount { get; set; } = 0;

    public double? LastActivityAt { get; set; }

    /// <summary>
    /// The turn in flight, or null when idle (busy guard reads this).
    /// </summary>
    public string? CurrentTurnId { get; set; }

    public string? LastError { get; set; }

    /// <summary>
    /// The model the latest accepted turn ran on; reopening seeds the picker from it.
    /// </summary>
    public string? LastModel { get; set; }

    /// <summary>
    /// ADR-048 shutdown-handoff resume token (opaque, worker-authored).
    /// </summary>
    public string? PendingHandoffToken { get; set; }

    public string? PendingHandoffTurnId { get; set; }

    /// <summary>
    /// Per-conversation HMAC key for worker stream frames; server-side only.
    /// </summary>
    public string? RunToken { get; set; }

    public double? ArchivedAt { get; set; }

    public double CreatedAt { get; set; }

    public double UpdatedAt { get; set; }

    /// <summary>
    /// Business time of the latest applied event.
    /// </summary>
    public double OccurredAt { get; set; }

    /// <summary>
    /// Event-log acceptance time. Together with lastEventId this is the cursor.
    /// </summary>
    public double AcceptedAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public string ProjectionVersion { get; set; } = null!;
}
