#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A project's ISO week, folded from its owners' events. The id is the fold's aggregate key.
/// </summary>
public partial class DigestProjectWeek
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string IsoWeek { get; set; } = null!;

    /// <summary>
    /// The fold state: counts, sums, a fixed-bucket latency histogram, capped model and evaluator maps.
    /// </summary>
    public string State { get; set; } = null!;

    public DateTime OccurredAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public DateTime AcceptedAt { get; set; }

    public string ProjectionVersion { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
