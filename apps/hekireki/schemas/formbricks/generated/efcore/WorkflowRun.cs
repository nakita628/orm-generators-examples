#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

public partial class WorkflowRun
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Workflow Workflow { get; set; } = null!;

    public string WorkflowId { get; set; } = null!;

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual WorkflowVersion? WorkflowVersion { get; set; }

    public string? WorkflowVersionId { get; set; }

    public virtual Response? Response { get; set; }

    public string? ResponseId { get; set; }

    public WorkflowRunStatus Status { get; set; } = WorkflowRunStatus.Queued;

    public string TriggerType { get; set; } = null!;

    public string? SurveyId { get; set; }

    public bool IsDryRun { get; set; } = false;

    public string? IdempotencyKey { get; set; }

    public int Attempt { get; set; } = 0;

    public DateTime? NextAttemptAt { get; set; }

    public DateTime? LastErrorAt { get; set; }

    /// <summary>
    /// [WorkflowTriggerRunPayload]
    /// </summary>
    public string TriggerPayload { get; set; } = null!;

    /// <summary>
    /// [WorkflowRunData]
    /// </summary>
    public string Data { get; set; } = "{}";

    public string? Error { get; set; }

    public DateTime? DispatchedAt { get; set; }

    public DateTime? StartedAt { get; set; }

    public DateTime? FinishedAt { get; set; }

    public virtual ICollection<WorkflowRunLog> Logs { get; set; } = new List<WorkflowRunLog>();

    /// <summary>
    /// The caller's IP for a run started by a Webhook Secret URL call; shown in run detail.
    /// </summary>
    public string? CallerIp { get; set; }

    /// <summary>
    /// When the run retention job cleared the run's content (ADR 0044): payload contents, run data, step
    /// input, output and errors, the run error and the caller IP. Null while the content is kept.
    /// </summary>
    public DateTime? ContentDeletedAt { get; set; }

    /// <summary>
    /// The retention period in days that applied when the content was cleared ("Details deleted after N days").
    /// </summary>
    public int? ContentRetentionDays { get; set; }

    /// <summary>
    /// The draft definition a Test run executed (ADR 0049), so its graph can be drawn after the draft
    /// changes. Null for live runs (they point at `workflowVersionId`) and for Test runs saved before it
    /// existed. Configuration like a version's definition, not run content: retention leaves it alone.
    /// [WorkflowExecutableDefinition]
    /// </summary>
    public string? TestRunDefinition { get; set; }
}
