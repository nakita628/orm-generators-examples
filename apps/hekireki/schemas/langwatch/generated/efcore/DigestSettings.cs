#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The digest's one settings row. `armed` is the weekly send's master switch; it starts off.
/// </summary>
public partial class DigestSettings
{
    public string Id { get; set; } = "singleton";

    public bool Armed { get; set; } = false;

    public string? LastPlannedWeek { get; set; }

    public DateTime UpdatedAt { get; set; }
}
