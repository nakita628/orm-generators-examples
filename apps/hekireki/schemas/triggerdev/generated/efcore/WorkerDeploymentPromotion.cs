#nullable enable

namespace Trigger.Models;

public partial class WorkerDeploymentPromotion
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// This is the promotion label, e.g. "current"
    /// </summary>
    public string Label { get; set; } = null!;

    public virtual WorkerDeployment Deployment { get; set; } = null!;

    public string DeploymentId { get; set; } = null!;

    public virtual RuntimeEnvironment Environment { get; set; } = null!;

    public string EnvironmentId { get; set; } = null!;
}
