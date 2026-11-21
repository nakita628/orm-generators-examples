#nullable enable

using System;

namespace LangWatch.Models;

public partial class VerificationToken
{
    public string Id { get; set; } = null!;

    public string Identifier { get; set; } = null!;

    /// <summary>
    /// Not unique alone: better-auth stores plain values here (a user id, a counter). See WEB-19201.
    /// </summary>
    public string Token { get; set; } = null!;

    public DateTime Expires { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
