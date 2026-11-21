#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Every ClickHouse `TenantId` this organization has ever written governance
/// rows under, appended on first use and never pruned.
///
/// The live `resolveGovProjectId` lookup cannot be this translation: it filters
/// `archivedAt: null` and so returns null forever once someone archives the
/// governance project, while the write path re-reads the same project by slug
/// with no such filter and keeps landing rows under the old tenant. The
/// organization would map to zero tenants on read and one on write - a
/// permanent split-brain in which the cost screen reports "no governance data"
/// and an erasure job erases nothing and reports success.
///
/// Reads resolve against the whole history; erasure walks all of it, because
/// personal data does not stop existing in a tenant that stopped being current.
/// </summary>
public partial class GovernanceTenantHistory
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// A hidden governance project id this organization has written rows under.
    /// </summary>
    public string TenantId { get; set; } = null!;

    public DateTime FirstUsedAt { get; set; }

    public DateTime LastUsedAt { get; set; }
}
