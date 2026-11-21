#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A person asked for a full seat the organization had no room for (modules/organization/specs/
/// seat-requests.feature). One row per organization and address; reopening reuses the row.
/// </summary>
public partial class SeatRequest
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// Lowercased address of the person asking.
    /// </summary>
    public string Email { get; set; } = null!;

    /// <summary>
    /// Null when only the address is known (a sign-in the identity provider refused).
    /// </summary>
    public string? UserId { get; set; }

    public string? Name { get; set; }

    /// <summary>
    /// OPEN | APPROVED | DECLINED
    /// </summary>
    public string Status { get; set; } = null!;

    /// <summary>
    /// pending | lite | idp_refused: how the person was left when the request opened.
    /// </summary>
    public string Source { get; set; } = null!;

    /// <summary>
    /// The identity provider also refused this person's sign-in.
    /// </summary>
    public bool IdpRefused { get; set; } = false;

    /// <summary>
    /// okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml
    /// </summary>
    public string? IdpKind { get; set; }

    /// <summary>
    /// When sales were told about this request; reset when it reopens.
    /// </summary>
    public DateTime? SalesNotifiedAt { get; set; }

    public DateTime OpenedAt { get; set; }

    public DateTime? ResolvedAt { get; set; }

    public string? ResolvedByUserId { get; set; }

    /// <summary>
    /// Set when an approval claims the request, cleared once its seat is given; an approval
    /// still pending long after is one whose seat never landed, and is reopened.
    /// </summary>
    public DateTime? SeatPendingSince { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
