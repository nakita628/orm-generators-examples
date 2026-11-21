#nullable enable

using System;

namespace Trigger.Models;

public partial class OrgMemberInvite
{
    public string Id { get; set; } = null!;

    public string Token { get; set; } = null!;

    public string Email { get; set; } = null!;

    public OrgMemberRole Role { get; set; } = OrgMemberRole.Member;

    /// <summary>
    /// Optional RBAC role to assign on invite acceptance. When set, the
    /// accept-invite flow calls the loaded RBAC plugin's setUserRole with
    /// this id after creating the OrgMember. Null = legacy behaviour, the
    /// runtime fallback derives the role from `role` above.
    ///
    /// Plain text (not an FK) — the RBAC plugin's RbacRole table lives on
    /// a separate schema (Drizzle, not Prisma) so we can't model the FK
    /// here. Validation happens at write time (action) and read time
    /// (acceptInvite).
    /// </summary>
    public string? RbacRoleId { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual User Inviter { get; set; } = null!;

    public string InviterId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
