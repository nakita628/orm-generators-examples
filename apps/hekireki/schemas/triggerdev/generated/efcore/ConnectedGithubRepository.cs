#nullable enable

using System;

namespace Trigger.Models;

public partial class ConnectedGithubRepository
{
    public string Id { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual GithubRepository Repository { get; set; } = null!;

    public string RepositoryId { get; set; } = null!;

    public string BranchTracking { get; set; } = null!;

    public bool PreviewDeploymentsEnabled { get; set; } = true;

    /// <summary>
    /// Skip queued builds once a newer commit is queued for the same environment or preview branch
    /// </summary>
    public bool SupersedeProductionBuilds { get; set; } = false;

    public bool SupersedeStagingBuilds { get; set; } = true;

    public bool SupersedePreviewBuilds { get; set; } = true;

    public string? BuildFilters { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
