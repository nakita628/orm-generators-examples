#nullable enable

using System;

namespace Jitsu.Models;

public partial class GlobalProps
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public string? Value { get; set; }
}
