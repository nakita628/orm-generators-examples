#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class OrgMember
{
    public string Id { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public OrgMemberRole Role { get; set; } = OrgMemberRole.Member;

    /// <summary>
    /// Who invited them (a User id, kept if that user is deleted). Null if not invited or joined before this was recorded.
    /// </summary>
    public string? InvitedById { get; set; }

    /// <summary>
    /// How they joined: created, invite, sso_jit, directory_sync or manual. Null if joined before this was recorded.
    /// </summary>
    public string? JoinMethod { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<RuntimeEnvironment> Environments { get; set; } = new List<RuntimeEnvironment>();
}
