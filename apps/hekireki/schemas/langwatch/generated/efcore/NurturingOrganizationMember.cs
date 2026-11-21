#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Nurturing's own copy of each organization's members, from the owners' member facts (§9):
/// organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
/// and member_disabled leave, member_enabled returns. Each column keeps its latest time.
/// </summary>
public partial class NurturingOrganizationMember
{
    public string OrganizationId { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public DateTime? JoinedAt { get; set; }

    public DateTime? LeftAt { get; set; }

    public DateTime? ReturnedAt { get; set; }
}
