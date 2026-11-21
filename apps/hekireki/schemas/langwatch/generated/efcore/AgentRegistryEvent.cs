#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The registry's append-only change feed: one row per registered, versioned, removed or
/// tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
/// fan-out as recorded in governance's outbox.
/// </summary>
public partial class AgentRegistryEvent
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public long Sequence { get; set; }

    public string Type { get; set; } = null!;

    public string EntryId { get; set; } = null!;

    public int? FromVersion { get; set; }

    public int? ToVersion { get; set; }

    public string Payload { get; set; } = null!;

    public DateTime OccurredAt { get; set; }

    public DateTime? DeliveryRecordedAt { get; set; }
}
