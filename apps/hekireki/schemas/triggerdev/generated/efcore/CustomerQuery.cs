#nullable enable

using System;

namespace Trigger.Models;

public partial class CustomerQuery
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The TSQL query text that was executed
    /// </summary>
    public string Query { get; set; } = null!;

    /// <summary>
    /// The scope of the query (determines which tenant IDs were used)
    /// </summary>
    public CustomerQueryScope Scope { get; set; }

    /// <summary>
    /// Query execution statistics from ClickHouse
    /// </summary>
    public string Stats { get; set; } = null!;

    /// <summary>
    /// AI-generated title summarizing the query
    /// </summary>
    public string? Title { get; set; }

    /// <summary>
    /// Where the query originated from
    /// </summary>
    public CustomerQuerySource Source { get; set; } = CustomerQuerySource.Dashboard;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Project? Project { get; set; }

    public string? ProjectId { get; set; }

    public virtual RuntimeEnvironment? Environment { get; set; }

    public string? EnvironmentId { get; set; }

    /// <summary>
    /// Optional user who executed the query (null for API calls)
    /// </summary>
    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    /// <summary>
    /// Time filter settings used when the query was executed
    /// Period like "7d", "24h", etc.
    /// </summary>
    public string? FilterPeriod { get; set; }

    /// <summary>
    /// Custom start date (ISO 8601)
    /// </summary>
    public DateTime? FilterFrom { get; set; }

    /// <summary>
    /// Custom end date (ISO 8601)
    /// </summary>
    public DateTime? FilterTo { get; set; }

    public DateTime CreatedAt { get; set; }
}
