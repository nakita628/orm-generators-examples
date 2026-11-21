#nullable enable

using System;

namespace Trigger.Models;

public partial class IntegrationDeployment
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// For example Vercel
    /// </summary>
    public string IntegrationName { get; set; } = null!;

    /// <summary>
    /// External ID
    /// </summary>
    public string IntegrationDeploymentId { get; set; } = null!;

    public string CommitSHA { get; set; } = null!;

    public string? DeploymentId { get; set; }

    /// <summary>
    /// External deployment status
    /// </summary>
    public string? Status { get; set; }

    public virtual WorkerDeployment? WorkerDeployment { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
