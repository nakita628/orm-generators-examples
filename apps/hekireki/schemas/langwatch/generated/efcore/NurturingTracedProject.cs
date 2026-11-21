#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Nurturing's copy of each project that received a real trace, from trace's and project's facts.
/// </summary>
public partial class NurturingTracedProject
{
    public string ProjectId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;
}
