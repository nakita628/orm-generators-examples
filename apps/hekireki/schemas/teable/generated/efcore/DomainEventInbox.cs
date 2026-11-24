#nullable enable

using System;

namespace Teable.Models;

public partial class DomainEventInbox
{
    public string ConsumerId { get; set; } = null!;

    public string EventId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }
}
