#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Stores the active password reset token for a user.
/// Tokens are opaque, hashed at rest, revocable, and single-use.
/// </summary>
public partial class PasswordResetToken
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string TokenHash { get; set; } = null!;

    public DateTime ExpiresAt { get; set; }

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;
}
