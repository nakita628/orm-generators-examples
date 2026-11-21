#nullable enable

using System;

namespace LangWatch.Models;

public partial class AuditLog
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public string? UserId { get; set; }

    public string? ProjectId { get; set; }

    public string? OrganizationId { get; set; }

    public string Action { get; set; } = null!;

    public string? Args { get; set; }

    public string? Error { get; set; }

    public string? IpAddress { get; set; }

    public string? UserAgent { get; set; }

    public string? Metadata { get; set; }

    /// <summary>
    /// WHO REALLY DID IT, when that is not `userId`.
    ///
    /// Under an impersonation the session's `user.id` is the SUBJECT — the
    /// customer whose access an operator is borrowing — so a row naming only
    /// `userId` filed the operator's act against the customer. An operator
    /// granting a permanent way into somebody's organization appeared, in the
    /// only durable record of it, to be that somebody.
    ///
    /// Null on an ordinary request, where the two would be the same person and
    /// saying so twice adds nothing. ADR-118 / D06: every decision names both.
    /// </summary>
    public string? ActorUserId { get; set; }

    public string? TargetKind { get; set; }

    public string? TargetId { get; set; }

    public string? Before { get; set; }

    public string? After { get; set; }

    /// <summary>
    /// A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2).
    /// </summary>
    public string? IdempotencyKey { get; set; }
}
