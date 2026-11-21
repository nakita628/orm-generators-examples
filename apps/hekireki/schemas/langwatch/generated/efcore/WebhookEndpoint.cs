#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// Org-anchored outbound webhook endpoint (the webhook platform, enterprise
/// only). The signing secret is AES-256-GCM encrypted at rest (shared
/// encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
/// once, at create or roll time. `enabledEvents` holds registry event types
/// or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
/// 72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
/// </summary>
public partial class WebhookEndpoint
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    /// <summary>
    /// Where this endpoint delivers. Everything above the last hop is shared:
    /// same batching, same retry ladder, same delivery log, same signature over
    /// the same bytes.
    /// </summary>
    public WebhookDestinationKind DestinationKind { get; set; } = WebhookDestinationKind.Http;

    /// <summary>
    /// The receiver URL. Present on `http` endpoints and null on the others;
    /// the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
    /// </summary>
    public string? Url { get; set; }

    /// <summary>
    /// The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
    /// region and the owning account are read off it rather than configured
    /// separately, so they can never disagree with the queue.
    /// </summary>
    public string? SqsQueueUrl { get; set; }

    /// <summary>
    /// The role we assume to write to the queue, and the ExternalId that proves
    /// the request came from us rather than from anyone who learned the role's
    /// name. The recommended credential mode.
    /// </summary>
    public string? SqsRoleArn { get; set; }

    public string? SqsExternalId { get; set; }

    /// <summary>
    /// A static key pair for the queue, encrypted at rest exactly like the
    /// signing secret. Both halves are present or both are absent: a half-filled
    /// pair is what the AWS SDK reads as "we have credentials" and stops looking
    /// with, so the CHECK constraint refuses it.
    /// </summary>
    public string? SqsAccessKeyId { get; set; }

    public string? SqsSecretAccessKeyEncrypted { get; set; }

    public string SecretEncrypted { get; set; } = null!;

    /// <summary>
    /// How deliveries are formatted and signed. Null: the batch envelope signed
    /// `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
    /// only for endpoints migrated from governance anomaly destinations.
    /// </summary>
    public string? SignatureScheme { get; set; }

    /// <summary>
    /// Set only by governance's anomaly destination migration: a create that
    /// repeats it in the same organization answers the endpoint it made. An
    /// archived endpoint gives its key up, so a retry after cleanup creates anew.
    /// </summary>
    public string? IdempotencyKey { get; set; }

    /// <summary>
    /// The secret this endpoint signed with before the last roll. Deliveries
    /// carry a second `v1=` computed from it until it expires, so a receiver
    /// swaps secrets on its own schedule instead of dropping deliveries in the
    /// gap. Null once the window closes or if the secret was never rolled.
    /// </summary>
    public string? PreviousSecretEncrypted { get; set; }

    /// <summary>
    /// When the previous secret stops being signed with and accepted.
    /// </summary>
    public DateTime? PreviousSecretExpiresAt { get; set; }

    public List<string>? EnabledEvents { get; set; } = new List<string>();

    public WebhookEndpointStatus Status { get; set; } = WebhookEndpointStatus.Active;

    /// <summary>
    /// Why the endpoint is disabled: "auto_failures_72h" or "manual".
    /// </summary>
    public string? DisabledReason { get; set; }

    public DateTime? DisabledAt { get; set; }

    /// <summary>
    /// Start of the current consecutive-failure streak; cleared on any 2xx.
    /// 72 hours of unbroken failures from this instant auto-disables.
    /// </summary>
    public DateTime? FailingSince { get; set; }

    /// <summary>
    /// Delivery controls, each server-bounded (see webhookDeliveryControls
    /// constants): envelopes per POST, coalescing window before a partial
    /// batch ships, and concurrent POSTs to this receiver.
    /// </summary>
    public int MaxBatchSize { get; set; } = 100;

    public int MaxBatchDelayMs { get; set; } = 250;

    public int MaxInFlight { get; set; } = 4;

    /// <summary>
    /// Self-hosted only: deliveries accept the receiver's self-signed TLS
    /// certificate. Off by default; the hosted product ignores it.
    /// </summary>
    public bool AllowSelfSignedCertificate { get; set; } = false;

    public DateTime? LastSuccessAt { get; set; }

    public DateTime? LastFailureAt { get; set; }

    public DateTime? ArchivedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<WebhookEndpointDelivery> Deliveries { get; set; } = new List<WebhookEndpointDelivery>();
}
