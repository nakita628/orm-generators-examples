#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Nurturing's read model: which organization a project belongs to, from project's created event.
/// The per-email usage places a project and lists an organization's projects from these rows only.
/// </summary>
public partial class NurturingProject
{
    public string ProjectId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;
}
