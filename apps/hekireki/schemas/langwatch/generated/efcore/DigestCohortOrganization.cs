#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// An organization enrolled in the weekly digest. No row: not enrolled, so nothing sends.
/// </summary>
public partial class DigestCohortOrganization
{
    public string OrganizationId { get; set; } = null!;

    public string EnrolledBy { get; set; } = null!;

    public DateTime EnrolledAt { get; set; }
}
