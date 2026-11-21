#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One agent in the organisation's governance registry, from any source governance reads
/// (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
/// machine-readable snapshot; `version` grows by one on every material change.
/// </summary>
public partial class AgentRegistryEntry
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// Where the agent came from and its id there, e.g. `langwatch:&lt;agentId&gt;`.
    /// </summary>
    public string SourceKey { get; set; } = null!;

    public string Source { get; set; } = null!;

    public string Status { get; set; } = null!;

    public int Version { get; set; }

    public string Fingerprint { get; set; } = null!;

    public string Entry { get; set; } = null!;

    public DateTime FirstSeenAt { get; set; }

    public DateTime? LastSeenAt { get; set; }

    public DateTime? RemovedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
