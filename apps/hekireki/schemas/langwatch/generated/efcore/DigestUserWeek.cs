#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A person's ISO week of coding-agent activity in one organization, folded from the agents' facts.
/// </summary>
public partial class DigestUserWeek
{
    public string Id { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string IsoWeek { get; set; } = null!;

    /// <summary>
    /// Sessions (a capped set of ids) and the agents' metric series, converged last-write-wins.
    /// </summary>
    public string State { get; set; } = null!;

    public DateTime OccurredAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public DateTime AcceptedAt { get; set; }

    public string ProjectionVersion { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
