#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// The two-step verification pipeline's projection (D06 - see
/// specs/identity/mfa-and-session-shape.feature). One row per PERSON, because
/// one person has one setup however many organizations they belong to.
///
/// A pure event-truth head like `SsoConnection`: every column is fold-written,
/// rows are never deleted (DISABLED is a tombstone, so history survives being
/// turned off), and a replay rebuilds it whole-row from the event log. The
/// cursor columns are inline because the fold produces exactly one row per
/// aggregate - the `Identifier` fan-out needed its own cursor table, this
/// does not.
///
/// Every column here is a lifecycle fact. There is deliberately nowhere in
/// this model a secret or a backup code could be put: `consumedBackupCodeIndexes`
/// holds POSITIONS, which is what lets "how many are left" be answered
/// without the projection ever knowing a code.
/// </summary>
public partial class MfaEnrollment
{
    /// <summary>
    /// The user. Also the aggregate id and the tenant - one identity history
    /// per person (ADR-029 §4).
    /// </summary>
    public string UserId { get; set; } = null!;

    /// <summary>
    /// The current enrollment, or NULL for somebody who never started one.
    /// </summary>
    public string? EnrollmentId { get; set; }

    /// <summary>
    /// totp. Never sms.
    /// </summary>
    public string? Method { get; set; }

    /// <summary>
    /// NONE | PENDING | ENABLED | EXPIRED | DISABLED
    /// </summary>
    public string State { get; set; } = null!;

    public DateTime? EnrolledAt { get; set; }

    public DateTime? ConfirmedAt { get; set; }

    public DateTime? ExpiredAt { get; set; }

    public DateTime? DisabledAt { get; set; }

    /// <summary>
    /// password+totp | org-admin - who ended it, and on whose authority.
    /// </summary>
    public string? DisabledVia { get; set; }

    public int BackupCodeCount { get; set; } = 0;

    /// <summary>
    /// Positions already spent, ascending. Never codes.
    /// </summary>
    public List<int>? ConsumedBackupCodeIndexes { get; set; } = new List<int>();

    public int FailedCount { get; set; } = 0;

    public DateTime OccurredAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public DateTime AcceptedAt { get; set; }

    public string ProjectionVersion { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
