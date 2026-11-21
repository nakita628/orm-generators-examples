#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Nurturing's copy of each personal project's owner, from organization's and project's facts.
/// The first owner recorded stays: a personal project's owner never changes.
/// </summary>
public partial class NurturingPersonalProject
{
    public string ProjectId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string OwnerUserId { get; set; } = null!;
}
