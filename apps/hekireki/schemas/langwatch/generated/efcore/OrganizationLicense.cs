#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The licence an organization activated and its dates, owned by licensing
/// (round 37 D6). A row whose key is null is a cleared licence. An organization
/// with no row is read from its Organization columns until the copy step runs.
/// </summary>
public partial class OrganizationLicense
{
    public string OrganizationId { get; set; } = null!;

    public string? LicenseKey { get; set; }

    public DateTime? ExpiresAt { get; set; }

    public DateTime? ValidatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
