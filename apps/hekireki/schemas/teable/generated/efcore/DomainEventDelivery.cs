#nullable enable

using System;

namespace Teable.Models;

public partial class DomainEventDelivery
{
    public string Id { get; set; } = null!;

    public string EventId { get; set; } = null!;

    public string ConsumerId { get; set; } = null!;

    public string Status { get; set; } = null!;

    public int Attempts { get; set; } = 0;

    public int MaxAttempts { get; set; } = 12;

    public string? LeaseToken { get; set; }

    public DateTime? LeaseExpiresAt { get; set; }

    public DateTime NextAttemptAt { get; set; }

    public string? LastError { get; set; }

    public DateTime CreatedAt { get; set; }
}
