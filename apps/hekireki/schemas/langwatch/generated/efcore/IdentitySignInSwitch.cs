#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Whether a user signs in on the identifier path: identity's own switch, set when the
/// identifier backfill proves the user and pinned off by an operator's rollback. No row is the
/// legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
/// </summary>
public partial class IdentitySignInSwitch
{
    public string UserId { get; set; } = null!;

    public bool OnIdentifiers { get; set; }

    public bool Pinned { get; set; } = false;

    public DateTime SwitchedAt { get; set; }

    public string? LastAppliedEventId { get; set; }

    public DateTime UpdatedAt { get; set; }
}
