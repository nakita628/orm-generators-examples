#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The same shape as DiscoveredPerson for non-human actors - Genie spaces,
/// Copilot bots, OpenAI projects - plus whatever provider-native metadata the
/// source published. Discovering an agent grants it nothing either.
/// </summary>
public partial class DiscoveredAgent
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string Provider { get; set; } = null!;

    /// <summary>
    /// Genie space / Copilot bot / OpenAI project id, verbatim.
    /// </summary>
    public string RawAgentId { get; set; } = null!;

    public string DisplayText { get; set; } = null!;

    /// <summary>
    /// Workspace, plan, and kin - provider-native descriptive fields only. Never
    /// credentials: this row is read by admin screens, not by any pull path.
    /// </summary>
    public string Metadata { get; set; } = null!;

    public DateTime FirstSeenAt { get; set; }

    public DateTime LastSeenAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
