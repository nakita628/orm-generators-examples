#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

public partial class Workflow
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public WorkflowStatus Status { get; set; } = WorkflowStatus.Draft;

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual User? Creator { get; set; }

    public string? CreatedBy { get; set; }

    /// <summary>
    /// [WorkflowDefinition]
    /// </summary>
    public string Definition { get; set; } = null!;

    public virtual ICollection<WorkflowVersion> Versions { get; set; } = new List<WorkflowVersion>();

    public virtual ICollection<WorkflowRun> Runs { get; set; } = new List<WorkflowRun>();

    public virtual ICollection<WorkflowSecret> Secrets { get; set; } = new List<WorkflowSecret>();

    public virtual ICollection<WorkflowSample> Samples { get; set; } = new List<WorkflowSample>();

    /// <summary>
    /// Webhook Secret URL token (ADRs 0048, 0057): SHA-256 hash to check calls, and an encrypted copy so
    /// the editor can show the URL again. Never part of the definition, its versions or its copies.
    /// </summary>
    public string? WebhookTokenHash { get; set; }

    public string? WebhookTokenEncrypted { get; set; }
}
