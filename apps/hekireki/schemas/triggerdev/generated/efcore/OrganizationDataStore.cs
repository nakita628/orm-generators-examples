#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

/// <summary>
/// Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
/// Multiple organizations can share a single data store row via organizationIds.
/// </summary>
public partial class OrganizationDataStore
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// Human-readable unique key (e.g. "hipaa-clickhouse-us-east")
    /// </summary>
    public string Key { get; set; } = null!;

    /// <summary>
    /// Organization IDs that use this data store
    /// </summary>
    public List<string>? OrganizationIds { get; set; }

    public DataStoreKind Kind { get; set; }

    /// <summary>
    /// Versioned config JSON. Structure is discriminated by the top-level `version` field.
    /// </summary>
    public string Config { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
