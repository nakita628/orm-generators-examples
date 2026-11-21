#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class WorkingHours
{
    public string Id { get; set; } = null!;

    public bool Enabled { get; set; } = false;

    public string WeeklySchedule { get; set; } = "{\"0\":{\"isOpen\":false},\"1\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"2\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"3\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"4\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"5\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"6\":{\"isOpen\":false}}";

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual ICollection<WorkingHoursOverride> Overrides { get; set; } = new List<WorkingHoursOverride>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
