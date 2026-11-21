#nullable enable

using System;

namespace Shelf.Models;

public partial class CustomTierLimit
{
    public string Id { get; set; } = null!;

    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    public bool CanImportAssets { get; set; } = true;

    public bool CanExportAssets { get; set; } = true;

    public bool CanImportNRM { get; set; } = true;

    public bool CanHideShelfBranding { get; set; } = true;

    public bool CanUseReports { get; set; } = true;

    public bool CanUseRecurringReminders { get; set; } = true;

    public bool CanUsePushNotifications { get; set; } = true;

    public int MaxCustomFields { get; set; } = 1000;

    public int MaxOrganizations { get; set; } = 1;

    public bool IsEnterprise { get; set; } = false;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
