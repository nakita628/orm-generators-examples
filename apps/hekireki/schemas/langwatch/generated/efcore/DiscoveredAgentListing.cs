#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One source's presence fact for one discovered agent. An agent is gone only when every
/// source that listed it marked it gone; an agent no source listed (older sightings) never is.
/// </summary>
public partial class DiscoveredAgentListing
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string IngestionSourceId { get; set; } = null!;

    public string DiscoveredAgentId { get; set; } = null!;

    public DateTime LastListedAt { get; set; }

    /// <summary>
    /// Set when this source's listing stopped naming the agent; cleared when it names it again.
    /// </summary>
    public DateTime? GoneAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
