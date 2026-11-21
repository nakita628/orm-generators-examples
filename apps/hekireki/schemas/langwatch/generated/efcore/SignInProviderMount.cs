#nullable enable

using System;

namespace LangWatch.Models;

public partial class SignInProviderMount
{
    /// <summary>
    /// Fixed, so this table holds one row: licensing's copy of sso's sign-in
    /// provider fact (ADR-179).
    /// </summary>
    public string Id { get; set; } = "self";

    /// <summary>
    /// The configured provider's name; null when the deployment signs in by email.
    /// </summary>
    public string? Provider { get; set; }

    public bool Mounted { get; set; }

    /// <summary>
    /// The last sso event applied: the row's idempotency guard and its version.
    /// </summary>
    public string LastAppliedEventId { get; set; } = null!;

    public DateTime UpdatedAt { get; set; }
}
