#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// A reason the environment variables "Needs attention" page shows a key in one environment, which
/// the project's team chose to ignore. An ignored EMPTY goes when the value is next saved.
/// </summary>
public partial class EnvironmentVariableIgnore
{
    public string Id { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual RuntimeEnvironment Environment { get; set; } = null!;

    public string EnvironmentId { get; set; } = null!;

    /// <summary>
    /// For MISSING, a variable of the project with no value in this environment.
    /// </summary>
    public virtual EnvironmentVariable Variable { get; set; } = null!;

    public string VariableId { get; set; } = null!;

    public EnvironmentVariableIgnoreReason Reason { get; set; }

    public virtual User? CreatedBy { get; set; }

    public string? CreatedById { get; set; }

    public DateTime CreatedAt { get; set; }
}
