#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A department is a single-valued accounting dimension, separate from
/// RBAC, for rolling up AI spend by org unit (Engineering, Marketing,
/// ...). Assigned to OrganizationUser, Team, and Project via their
/// `departmentId` column. Pure accounting: never grants or restricts
/// access. See specs/ai-gateway/governance/departments.feature.
/// </summary>
public partial class Department
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string Name { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? ArchivedAt { get; set; }
}
