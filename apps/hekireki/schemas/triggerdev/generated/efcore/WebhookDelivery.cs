#nullable enable

using System;

namespace Trigger.Models;

public partial class WebhookDelivery
{
    public string Id { get; set; } = null!;

    public string FriendlyId { get; set; } = null!;

    public string WebhookEndpointId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string RuntimeEnvironmentId { get; set; } = null!;

    public RuntimeEnvironmentType EnvironmentType { get; set; }

    public string ExternalDeliveryId { get; set; } = null!;

    public string IdempotencyKey { get; set; } = null!;

    public string? RunId { get; set; }

    public WebhookDeliveryStatus Status { get; set; } = WebhookDeliveryStatus.Pending;

    /// <summary>
    /// Set from the x-trigger-test ingress header; marks console/test-send deliveries so the list can filter them.
    /// </summary>
    public bool IsTest { get; set; } = false;

    public string? ParsedEvent { get; set; }

    public string? Headers { get; set; }

    public string? RawBodyHash { get; set; }

    public string? ErrorMessage { get; set; }

    public string? FilterReason { get; set; }

    /// <summary>
    /// Per-target results: WebhookDeliveryTargetResult[] ({ id, type, status, reason?, runId?, error? }).
    /// </summary>
    public string Targets { get; set; } = "[]";

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? ProcessedAt { get; set; }

    /// <summary>
    /// Retention class in days: the LIST partition key, stamped at ingest from the org's plan and encoded in the id.
    /// </summary>
    public int RetentionDays { get; set; } = 30;
}
