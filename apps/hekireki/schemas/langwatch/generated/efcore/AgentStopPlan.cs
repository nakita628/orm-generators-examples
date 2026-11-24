#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One agent's stop plan version: bumped by every save, so a save over a plan someone else saved
/// since it was read is refused. No row means the plan was never saved (version 0).
/// </summary>
public partial class AgentStopPlan
{
    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// The AgentRegistryEntry the plan belongs to.
    /// </summary>
    public string AgentId { get; set; } = null!;

    public int Version { get; set; }

    public DateTime UpdatedAt { get; set; }
}
