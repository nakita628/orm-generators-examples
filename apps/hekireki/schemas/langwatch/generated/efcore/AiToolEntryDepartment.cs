#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Per-department scope binding for AiToolEntry. Empty set = org-wide.
/// Non-empty = entry visible only to members of those departments (an
/// OrganizationUser whose `departmentId` is in the set). Mirrors
/// AiToolEntryTeam's shape. `departmentId` is a plain column ref
/// (relationMode="prisma"), matching how Department is referenced
/// elsewhere; the service validates the id belongs to the org.
/// </summary>
public partial class AiToolEntryDepartment
{
    public string Id { get; set; } = null!;

    public string EntryId { get; set; } = null!;

    public string DepartmentId { get; set; } = null!;

    public virtual AiToolEntry Entry { get; set; } = null!;
}
