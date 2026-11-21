#nullable enable

using System;

namespace Trigger.Models;

public partial class WorkerInstance
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// For example "worker-1"
    /// </summary>
    public string Name { get; set; } = null!;

    /// <summary>
    /// If managed, it will default to the name, e.g. "worker-1"
    /// If unmanged, it will be prefixed with the deployment ID e.g. "deploy-123-worker-1"
    /// </summary>
    public string ResourceIdentifier { get; set; } = null!;

    public string? Metadata { get; set; }

    public virtual WorkerInstanceGroup WorkerGroup { get; set; } = null!;

    public string WorkerGroupId { get; set; } = null!;

    public virtual Organization? Organization { get; set; }

    public string? OrganizationId { get; set; }

    public virtual Project? Project { get; set; }

    public string? ProjectId { get; set; }

    public virtual RuntimeEnvironment? Environment { get; set; }

    public string? EnvironmentId { get; set; }

    public virtual WorkerDeployment? Deployment { get; set; }

    public string? DeploymentId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? LastDequeueAt { get; set; }

    public DateTime? LastHeartbeatAt { get; set; }
}
