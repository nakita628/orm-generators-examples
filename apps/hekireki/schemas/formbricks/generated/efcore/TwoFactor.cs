#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
/// User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
/// at rest by Better Auth using BETTER_AUTH_SECRET.
/// </summary>
public partial class TwoFactor
{
    public string Id { get; set; } = null!;

    public string Secret { get; set; } = null!;

    public string BackupCodes { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public bool Verified { get; set; } = true;

    /// <summary>
    /// Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts.
    /// </summary>
    public int FailedVerificationCount { get; set; } = 0;

    /// <summary>
    /// Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit.
    /// </summary>
    public DateTime? LockedUntil { get; set; }
}
