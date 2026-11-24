#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot.
/// </summary>
public partial class AuthzOrganizationPolicy
{
    public string OrganizationId { get; set; } = null!;

    public bool BraidsOrganizationAdminsOnly { get; set; } = false;

    public bool BraidsEnabled { get; set; } = false;

    public string UpdatedByUserId { get; set; } = null!;

    public DateTime UpdatedAt { get; set; }
}
