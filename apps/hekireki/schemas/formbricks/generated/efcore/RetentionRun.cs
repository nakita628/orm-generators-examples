#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
/// while the run is in progress, or if it never finished.
/// </summary>
public partial class RetentionRun
{
    public string Id { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public RetentionEntity Entity { get; set; }

    public RetentionRunKind Kind { get; set; } = RetentionRunKind.Sweep;

    public DateTime StartedAt { get; set; }

    public DateTime? FinishedAt { get; set; }

    public int NotifiedCount { get; set; } = 0;

    /// <summary>
    /// Archived surveys, or deactivated members.
    /// </summary>
    public int ArchivedCount { get; set; } = 0;

    /// <summary>
    /// Deleted responses, or, on a purge run, deleted surveys.
    /// </summary>
    public int DeletedCount { get; set; } = 0;

    public int SkippedCount { get; set; } = 0;

    /// <summary>
    /// False when the run changed nothing. History hides those by default.
    /// </summary>
    public bool HasChanges { get; set; } = false;

    /// <summary>
    /// Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
    /// of rescanning the same first targets every night. Null when the scan reached the end.
    /// </summary>
    public string? ScanCursor { get; set; }

    public virtual ICollection<RetentionRunItem> Items { get; set; } = new List<RetentionRunItem>();
}
