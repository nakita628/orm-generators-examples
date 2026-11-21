#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The mail log for seat requests and refused sign-ins: at most one row per organization,
/// kind, address and UTC day, each mailed at once.
/// </summary>
public partial class OrganizationAccessNotice
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// seat_request | idp_refused
    /// </summary>
    public string Kind { get; set; } = null!;

    public string Email { get; set; } = null!;

    public string? Name { get; set; }

    public string? Idp { get; set; }

    /// <summary>
    /// UTC day the notice was recorded on, YYYY-MM-DD.
    /// </summary>
    public string DayKey { get; set; } = null!;

    /// <summary>
    /// Whether the seat request it announces was also refused by the identity provider.
    /// </summary>
    public bool IdpRefused { get; set; } = false;

    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// Set once the notice was handed to the mailer; null if a send never started.
    /// </summary>
    public DateTime? MailedAt { get; set; }
}
