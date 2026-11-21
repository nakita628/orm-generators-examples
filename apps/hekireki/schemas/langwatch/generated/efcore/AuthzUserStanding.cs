#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant").
/// </summary>
public partial class AuthzUserStanding
{
    public string UserId { get; set; } = null!;

    public DateTime? DeactivatedAt { get; set; }

    public DateTime? ErasedAt { get; set; }

    public DateTime StandingChangedAt { get; set; }
}
