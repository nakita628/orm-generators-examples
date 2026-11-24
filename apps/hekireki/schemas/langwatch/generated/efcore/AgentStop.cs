#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
/// every target's outcome, folded from the stop's events.
/// </summary>
public partial class AgentStop
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string AgentId { get; set; } = null!;

    public string AgentName { get; set; } = null!;

    public string Status { get; set; } = null!;

    public string Severity { get; set; } = null!;

    public string Approval { get; set; } = null!;

    public string Trigger { get; set; } = null!;

    public string Reason { get; set; } = null!;

    public string? RequestedByUserId { get; set; }

    public string? RequestedByApiKeyId { get; set; }

    public string? DecidedByUserId { get; set; }

    public string? DecidedByApiKeyId { get; set; }

    public string? RestoredByUserId { get; set; }

    public string? RestoredByApiKeyId { get; set; }

    public string Targets { get; set; } = null!;

    public DateTime RequestedAt { get; set; }

    public DateTime? DecidedAt { get; set; }

    public DateTime? QuarantinedAt { get; set; }

    public DateTime? RestoreRequestedAt { get; set; }

    public DateTime? RestoredAt { get; set; }

    public string LastAppliedEventId { get; set; } = null!;

    public long LastAcceptedAt { get; set; }

    public long LastOccurredAt { get; set; }

    public string ProjectionVersion { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
