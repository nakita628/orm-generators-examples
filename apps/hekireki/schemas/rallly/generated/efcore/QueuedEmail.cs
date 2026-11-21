#nullable enable

using System;

namespace Rallly.Models;

/// <summary>
/// An email waiting to be sent by the queue, one row per recipient. The
/// email is built when it is sent, from the row its subject points at.
/// </summary>
public partial class QueuedEmail
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// Which handler builds the email. A string rather than an enum so a new
    /// kind needs no migration; the claim only takes kinds the build knows.
    /// </summary>
    public string Kind { get; set; } = null!;

    /// <summary>
    /// The row the email is about, interpreted by the kind. Not a foreign key:
    /// kinds point at different tables, and a subject deleted before the send
    /// makes its handler skip the email.
    /// </summary>
    public string SubjectId { get; set; } = null!;

    /// <summary>
    /// Groups the emails one action queued, so the action can send its own
    /// first batch at once and its progress can be read back.
    /// </summary>
    public string BatchId { get; set; } = null!;

    /// <summary>
    /// The account whose action queued the email. Banning it skips whatever
    /// is still pending.
    /// </summary>
    public string UserId { get; set; } = null!;

    public QueuedEmailStatus Status { get; set; } = QueuedEmailStatus.Pending;

    /// <summary>
    /// Not claimed before this time. Unset sends at the next run.
    /// </summary>
    public DateTime? SendAfter { get; set; }

    public int Attempts { get; set; } = 0;

    public DateTime? ClaimedAt { get; set; }

    public DateTime? SentAt { get; set; }

    public string? LastError { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User User { get; set; } = null!;
}
