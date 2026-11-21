#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The two-factor plugin's own table (D06). Protocol state, and row-truth for
/// good (ADR-101 R12): the shared secret and the backup codes live here and
/// in no event, ever. Nothing in this repo reads or writes these columns -
/// the plugin owns them end to end, including the lockout counters, which is
/// why `failedVerificationCount` and `lockedUntil` are here rather than
/// rebuilt on the projection.
///
/// `backupCodes` holds an ENCRYPTED blob, not hashes. A backup code has to be
/// compared against what somebody types, and the plugin decrypts and compares;
/// hashing them would leave it unable to verify its own codes.
/// </summary>
public partial class TwoFactor
{
    public string Id { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string Secret { get; set; } = null!;

    public string BackupCodes { get; set; } = null!;

    public bool Verified { get; set; } = true;

    public int FailedVerificationCount { get; set; } = 0;

    public DateTime? LockedUntil { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
