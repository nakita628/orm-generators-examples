#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One target of a registry agent's kill switch stop plan (enterprise/modules/governance/specs/
/// agent-kill-switch.feature): a Databricks serving endpoint, a Bedrock agent alias or a gateway
/// virtual key. A plan write replaces every row of the agent at once.
/// </summary>
public partial class AgentStopPlanTarget
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// The AgentRegistryEntry the plan belongs to.
    /// </summary>
    public string AgentId { get; set; } = null!;

    public int Position { get; set; }

    public string Label { get; set; } = null!;

    public string Target { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
