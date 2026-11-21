#nullable enable

using System;

namespace Shelf.Models;

public partial class TierLimit
{
    public TierId Id { get; set; }

    public virtual Tier? Tier { get; set; }

    public bool CanImportAssets { get; set; } = false;

    public bool CanExportAssets { get; set; } = false;

    public bool CanImportNRM { get; set; } = false;

    public bool CanHideShelfBranding { get; set; } = false;

    public bool CanUseReports { get; set; } = false;

    public bool CanUseRecurringReminders { get; set; } = false;

    public bool CanUsePushNotifications { get; set; } = false;

    public int MaxCustomFields { get; set; } = 0;

    public int MaxOrganizations { get; set; } = 1;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
