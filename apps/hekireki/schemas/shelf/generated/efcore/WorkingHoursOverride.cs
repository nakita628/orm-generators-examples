#nullable enable

using System;

namespace Shelf.Models;

public partial class WorkingHoursOverride
{
    public string Id { get; set; } = null!;

    public DateOnly Date { get; set; }

    public bool IsOpen { get; set; } = false;

    public string? OpenTime { get; set; }

    public string? CloseTime { get; set; }

    public string? Reason { get; set; }

    public virtual WorkingHours WorkingHours { get; set; } = null!;

    public string WorkingHoursId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
