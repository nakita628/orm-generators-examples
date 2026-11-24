#nullable enable

namespace LangWatch.Models;

public partial class RetentionApplyStatus
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public RetentionPolicyScopeType ScopeType { get; set; }

    public string ScopeId { get; set; } = null!;

    public double? ChangedAt { get; set; }

    public double? AppliedAt { get; set; }

    public double? RaisedAt { get; set; }

    public double? ShortenedAt { get; set; }

    public double CreatedAt { get; set; }

    public double UpdatedAt { get; set; }

    public double OccurredAt { get; set; }

    public double AcceptedAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public string ProjectionVersion { get; set; } = null!;
}
