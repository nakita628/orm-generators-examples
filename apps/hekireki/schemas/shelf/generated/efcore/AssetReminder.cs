#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class AssetReminder
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string Message { get; set; } = null!;

    public DateTime AlertDateTime { get; set; }

    public string? ActiveSchedulerReference { get; set; }

    public ReminderRecurrenceUnit? RecurrenceUnit { get; set; }

    public int? RecurrenceInterval { get; set; }

    public string? RecurrenceTimezone { get; set; }

    public DateTime? RecurrenceEndsAt { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Asset Asset { get; set; } = null!;

    public string AssetId { get; set; } = null!;

    public virtual ICollection<TeamMember> TeamMembers { get; set; } = new List<TeamMember>();

    public virtual User CreatedBy { get; set; } = null!;

    public string CreatedById { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
