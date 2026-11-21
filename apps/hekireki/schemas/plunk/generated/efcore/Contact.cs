#nullable enable

using System;
using System.Collections.Generic;

namespace Plunk.Models;

public partial class Contact
{
    public string Id { get; set; } = null!;

    public string Email { get; set; } = null!;

    public string? Data { get; set; }

    public bool Subscribed { get; set; } = true;

    /// <summary>
    /// When set, `subscribed = false` is temporary rather than permanent: the snooze
    /// sweep resubscribes the contact once this passes. Suppression itself is still
    /// carried entirely by `subscribed`, so every send path filters snoozed contacts
    /// out without knowing this column exists.
    ///
    /// Cleared by every other write to `subscribed` — including the bounce and
    /// complaint paths, so a suppressed address can never be woken back up.
    /// </summary>
    public DateTime? SnoozedUntil { get; set; }

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual ICollection<Email> Emails { get; set; } = new List<Email>();

    public virtual ICollection<WorkflowExecution> WorkflowExecutions { get; set; } = new List<WorkflowExecution>();

    public virtual ICollection<Event> Events { get; set; } = new List<Event>();

    public virtual ICollection<SegmentMembership> SegmentMemberships { get; set; } = new List<SegmentMembership>();

    public virtual ICollection<ThrottleClaim> ThrottleClaims { get; set; } = new List<ThrottleClaim>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
