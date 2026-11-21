#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// A person's push notification switches. A missing row means every switch
/// is on, so no backfill is needed. The three category columns are named
/// after `PushCategory` in apps/webapp/app/modules/push/types.ts.
/// </summary>
public partial class NotificationPreference
{
    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public bool PushEnabled { get; set; } = true;

    public bool Bookings { get; set; } = true;

    public bool Audits { get; set; } = true;

    public bool StockAndReminders { get; set; } = true;

    public DateTime UpdatedAt { get; set; }
}
