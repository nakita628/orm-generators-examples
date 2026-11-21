#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Held by the one seat-taking write running on an organization (seat-requests.feature). A
/// row past expiresAt is free to take; a holder deletes its own row when its write is done.
/// </summary>
public partial class OrganizationSeatLease
{
    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// Random per hold, so a release deletes only its own lease.
    /// </summary>
    public string Holder { get; set; } = null!;

    public DateTime ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
