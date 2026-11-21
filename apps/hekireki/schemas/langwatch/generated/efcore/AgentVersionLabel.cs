#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A deploy label ("production", "staging") pointing at one version of an agent.
/// </summary>
public partial class AgentVersionLabel
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string AgentId { get; set; } = null!;

    public string Label { get; set; } = null!;

    public string VersionId { get; set; } = null!;

    public string? CreatedById { get; set; }

    public string? UpdatedById { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
