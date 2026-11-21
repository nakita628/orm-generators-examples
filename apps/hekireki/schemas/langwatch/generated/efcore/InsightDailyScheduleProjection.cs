#nullable enable

namespace LangWatch.Models;

/// <summary>
/// One person's daily insights run on one board in a project, folded from its
/// `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
/// replaying the event log. Epochs in milliseconds.
/// </summary>
public partial class InsightDailyScheduleProjection
{
    /// <summary>
    /// The schedule id, which is also its aggregate id: one per project, person and board.
    /// </summary>
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string UserId { get; set; } = null!;

    /// <summary>
    /// dashboard | template. A pointer only, never a relation: a run checks the board is there.
    /// </summary>
    public string BoardKind { get; set; } = null!;

    public string BoardId { get; set; } = null!;

    public string BoardName { get; set; } = null!;

    /// <summary>
    /// on | off | undecided.
    /// </summary>
    public string State { get; set; } = "undecided";

    /// <summary>
    /// board_deleted | skipped_days | member_removed | project_archived. Why the system turned
    /// it off. NULL while it is on, and when the person turned it off.
    /// </summary>
    public string? OffReason { get; set; }

    public int? Hour { get; set; }

    public string? Timezone { get; set; }

    public int? MaxInsights { get; set; }

    /// <summary>
    /// The last run that recorded an outcome. NULL until one did.
    /// </summary>
    public string? LastRunId { get; set; }

    public double? LastRunAt { get; set; }

    /// <summary>
    /// filed | nothing | failed | skipped
    /// </summary>
    public string? LastRunOutcome { get; set; }

    public string? LastRunReason { get; set; }

    public int? LastRunFiled { get; set; }

    /// <summary>
    /// How many insights the run found still true and renewed.
    /// </summary>
    public int? LastRunRenewed { get; set; }

    public string? LastRunConversationId { get; set; }

    /// <summary>
    /// The scheduled days in a row whose run ended skipped, and the slot of the newest counted.
    /// </summary>
    public int ConsecutiveSkips { get; set; } = 0;

    public double? SkipStreakSlot { get; set; }

    public double CreatedAt { get; set; }

    public double UpdatedAt { get; set; }

    public double OccurredAt { get; set; }

    public double AcceptedAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public string ProjectionVersion { get; set; } = null!;
}
