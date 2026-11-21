#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// *
/// * Error group state is used to track when a user has interacted with an error (ignored/resolved)
/// * The actual error data is in ClickHouse.
/// </summary>
public partial class ErrorGroupState
{
    public string Id { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// *
    /// * You can ignore/resolve an error across all environments, or specific ones
    /// </summary>
    public virtual RuntimeEnvironment? Environment { get; set; }

    public string? EnvironmentId { get; set; }

    public string TaskIdentifier { get; set; } = null!;

    public string ErrorFingerprint { get; set; } = null!;

    public ErrorGroupStatus Status { get; set; } = ErrorGroupStatus.Unresolved;

    /// <summary>
    /// *
    /// * Error is ignored until this date
    /// </summary>
    public DateTime? IgnoredUntil { get; set; }

    /// <summary>
    /// *
    /// * Error is ignored until this occurrence rate
    /// </summary>
    public int? IgnoredUntilOccurrenceRate { get; set; }

    /// <summary>
    /// *
    /// * Error is ignored until this total occurrences
    /// </summary>
    public int? IgnoredUntilTotalOccurrences { get; set; }

    /// <summary>
    /// Total occurrence count at the time the error was ignored (from ClickHouse).
    /// Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
    /// </summary>
    public long? IgnoredAtOccurrenceCount { get; set; }

    /// <summary>
    /// *
    /// * Error was ignored at this date
    /// </summary>
    public DateTime? IgnoredAt { get; set; }

    /// <summary>
    /// *
    /// * Reason for ignoring the error
    /// </summary>
    public string? IgnoredReason { get; set; }

    /// <summary>
    /// *
    /// * User who ignored the error
    /// </summary>
    public string? IgnoredByUserId { get; set; }

    /// <summary>
    /// *
    /// * Error was resolved at this date
    /// </summary>
    public DateTime? ResolvedAt { get; set; }

    /// <summary>
    /// *
    /// * Error was resolved in this version
    /// </summary>
    public string? ResolvedInVersion { get; set; }

    /// <summary>
    /// *
    /// * User who resolved the error
    /// </summary>
    public string? ResolvedBy { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
