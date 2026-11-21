#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Stores one-time verification tokens used by Auth.js adapter flows.
/// </summary>
public partial class VerificationToken
{
    public string Identifier { get; set; } = null!;

    public string Token { get; set; } = null!;

    public DateTime Expires { get; set; }
}
