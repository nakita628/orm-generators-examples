#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Per-call data kept beside a workflow (ADR 0048): one row per source, `trigger` for the Webhook
/// trigger's last payload (plus the last call's time and IP), or an HTTP request step id for its last
/// Test request response. Written on every call, so it lives here and never bumps `Workflow.updatedAt`.
/// </summary>
public partial class WorkflowSample
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Workflow Workflow { get; set; } = null!;

    public string WorkflowId { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public string Source { get; set; } = null!;

    /// <summary>
    /// [WorkflowSampleData]
    /// </summary>
    public string Sample { get; set; } = null!;

    public DateTime? LastCalledAt { get; set; }

    public string? LastCallerIp { get; set; }
}
