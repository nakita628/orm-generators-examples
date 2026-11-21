#nullable enable

using System;

namespace Jitsu.Models;

public partial class UserPassword
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

    public string UserId { get; set; } = null!;

    public virtual UserProfile User { get; set; } = null!;

    public string Hash { get; set; } = null!;

    public bool ChangeAtNextLogin { get; set; } = false;
}
