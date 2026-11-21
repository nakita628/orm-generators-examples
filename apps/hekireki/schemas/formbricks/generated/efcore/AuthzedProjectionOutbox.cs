#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Durable PostgreSQL queue for rebuilding SpiceDB relationships after an authorization source changes.
/// Source-table triggers insert these rows in the same transaction as the mutation. Identifiers are
/// operationally sensitive and must never be included in logs, metrics, or CLI output.
///
/// This model declares NO indexes on purpose. All three live in
/// `migration/20260818120000_add_authzed_projection_outbox/migration.sql`, because each one is
/// partial (`WHERE "processedAt" IS NULL AND ...`) and Prisma's `@@index` has no `where`. Declaring
/// approximate non-partial copies here would be worse than declaring none: `prisma db push` would
/// see a name match with a different definition, drop and recreate them without the predicate, and
/// the migration's `CREATE INDEX IF NOT EXISTS` would then skip them — leaving a silently wrong
/// shape that carries a week of delivered rows on every hot-path lookup. Deployments are unaffected
/// either way (`prisma migrate deploy` never reads this file); a dev who runs `db push` restores
/// them by rerunning the migration, which is the same contract its triggers already live under.
/// </summary>
public partial class AuthzedProjectionOutbox
{
    public string Id { get; set; } = null!;

    public string TargetType { get; set; } = null!;

    public string PrimaryId { get; set; } = null!;

    public string? SecondaryId { get; set; }

    public bool IsRevocation { get; set; } = false;

    /// <summary>
    /// Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried.
    /// </summary>
    public int Attempts { get; set; } = 0;

    /// <summary>
    /// Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
    /// that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
    /// </summary>
    public int PermanentFailures { get; set; } = 0;

    public DateTime AvailableAt { get; set; }

    public DateTime? LeasedAt { get; set; }

    public DateTime? LeaseExpiresAt { get; set; }

    public string? LeaseOwner { get; set; }

    public DateTime? ProcessedAt { get; set; }

    public DateTime? DeadLetteredAt { get; set; }

    public DateTime? LastAttemptAt { get; set; }

    public string? LastErrorCode { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
