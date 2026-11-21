#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The per-ATTEMPT delivery log for BOTH webhook channels: the endpoints
/// platform's batches and the automations channel's trigger fires. One row per
/// HTTP attempt; `dispatchId` groups every attempt of one logical send (the
/// value of the X-LangWatch-Delivery-Id or X-LangWatch-Event-Id header). A slim
/// facts table: outcome, status, latency, a capped error message, and a
/// truncated copy of a failure response. Our request content (URL, headers,
/// body) is never stored. Pruned after 30 days.
///
/// The two channels used to keep separate tables with the same columns and two
/// prune implementations. They are one table with a discriminator now, so the
/// tenancy columns of the channel that did not write a row are null.
/// </summary>
public partial class WebhookEndpointDelivery
{
    public string Id { get; set; } = null!;

    public WebhookDeliveryChannel Channel { get; set; } = WebhookDeliveryChannel.Platform;

    /// <summary>
    /// Platform tenancy. Null on automations rows.
    /// </summary>
    public string? OrganizationId { get; set; }

    public string? EndpointId { get; set; }

    public virtual WebhookEndpoint? Endpoint { get; set; }

    /// <summary>
    /// Automations tenancy. Null on platform rows.
    /// </summary>
    public string? ProjectId { get; set; }

    public virtual Project? Project { get; set; }

    public string? TriggerId { get; set; }

    public virtual Trigger? Trigger { get; set; }

    public string DispatchId { get; set; } = null!;

    /// <summary>
    /// 1-based attempt and the batch's envelope count. Both null on automations
    /// rows, which record neither: one trigger fire is one event, and the
    /// channel's sender does not thread its attempt number into the log.
    /// </summary>
    public int? Attempt { get; set; }

    public int? EventCount { get; set; }

    public int? ResponseStatus { get; set; }

    public int? LatencyMs { get; set; }

    public string? Error { get; set; }

    public string? Response { get; set; }

    public WebhookDeliveryOutcome Outcome { get; set; }

    public DateTime FiredAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
