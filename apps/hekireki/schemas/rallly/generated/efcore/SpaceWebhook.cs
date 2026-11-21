#nullable enable

using System;
using System.Collections.Generic;

namespace Rallly.Models;

public partial class SpaceWebhook
{
    public string Id { get; set; } = null!;

    public string SpaceId { get; set; } = null!;

    public string Url { get; set; } = null!;

    /// <summary>
    /// AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation
    /// </summary>
    public string Secret { get; set; } = null!;

    /// <summary>
    /// Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations
    /// </summary>
    public List<string>? Events { get; set; }

    public bool Enabled { get; set; } = true;

    /// <summary>
    /// The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape
    /// </summary>
    public string Version { get; set; } = null!;

    /// <summary>
    /// Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history
    /// </summary>
    public DateTime Cursor { get; set; }

    public DateTime? LastDeliveredAt { get; set; }

    public DateTime? LastFailedAt { get; set; }

    /// <summary>
    /// Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap
    /// </summary>
    public int ConsecutiveFailures { get; set; } = 0;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Space Space { get; set; } = null!;

    public virtual ICollection<WebhookDelivery> Deliveries { get; set; } = new List<WebhookDelivery>();
}
