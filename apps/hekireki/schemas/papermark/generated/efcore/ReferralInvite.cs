#nullable enable

using System;

namespace Papermark.Models;

public partial class ReferralInvite
{
    public string Id { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string? TeamId { get; set; }

    public virtual Team? Team { get; set; }

    /// <summary>
    /// *
    /// * Whose cap and discount this invite counts against — the fund's team, or the
    /// * sender alone. One non-null key rather than a pair of nullable ones: on a
    /// * nullable teamId every personal invite is distinct to Postgres, so the two
    /// * keys guarded different scopes and a member could not invite a company both
    /// * for their fund and for themselves.
    /// </summary>
    public string ScopeId { get; set; } = null!;

    public string Email { get; set; } = null!;

    public string? ReferralLink { get; set; }

    public DateTime SentAt { get; set; }
}
