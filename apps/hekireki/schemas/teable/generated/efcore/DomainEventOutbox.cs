#nullable enable

using System;

namespace Teable.Models;

public partial class DomainEventOutbox
{
    public string Id { get; set; } = null!;

    public string BaseId { get; set; } = null!;

    public string? TableId { get; set; }

    public string MessageName { get; set; } = null!;

    public int SchemaVersion { get; set; }

    public string? AggregateId { get; set; }

    public string Payload { get; set; } = null!;

    public int PayloadBytes { get; set; }

    public int CatalogGeneration { get; set; }

    public string RequiredConsumers { get; set; } = null!;

    public string? BindingId { get; set; }

    public int? StorageEpoch { get; set; }

    public bool Unpublished { get; set; } = true;

    public string? Settled { get; set; }

    public DateTime? SettledAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
