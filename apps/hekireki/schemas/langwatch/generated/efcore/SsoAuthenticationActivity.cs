#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A successful callback through one SSO connection. Identifier.lastUsedAt
/// also advances during routing and resolution, so it cannot prove that a
/// member actually authenticated through the replacement before cutover.
/// </summary>
public partial class SsoAuthenticationActivity
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string ConnectionId { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public DateTime AuthenticatedAt { get; set; }

    /// <summary>
    /// The provider-side account the assertion named; null before #7633's arrival door recorded it.
    /// </summary>
    public string? ProviderAccountId { get; set; }
}
