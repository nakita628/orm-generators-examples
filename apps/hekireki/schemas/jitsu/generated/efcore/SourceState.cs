#nullable enable

using System;

namespace Jitsu.Models;

public partial class SourceState
{
    public string SyncId { get; set; } = null!;

    public string Stream { get; set; } = null!;

    public string State { get; set; } = null!;

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime Timestamp { get; set; }
}
