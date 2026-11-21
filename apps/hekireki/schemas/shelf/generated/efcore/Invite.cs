#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Invite
{
    public string Id { get; set; } = null!;

    public virtual User Inviter { get; set; } = null!;

    public string InviterId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual User? InviteeUser { get; set; }

    public string? InviteeUserId { get; set; }

    public virtual TeamMember InviteeTeamMember { get; set; } = null!;

    public string TeamMemberId { get; set; } = null!;

    public string InviteeEmail { get; set; } = null!;

    public InviteStatuses Status { get; set; } = InviteStatuses.Pending;

    public string InviteCode { get; set; } = null!;

    public List<OrganizationRoles>? Roles { get; set; }

    public string? InviteMessage { get; set; }

    public DateTime ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
