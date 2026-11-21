#nullable enable

using System;

namespace Shelf.Models;

public partial class TeamMemberNote
{
    public string Id { get; set; } = null!;

    public string Content { get; set; } = null!;

    public NoteType Type { get; set; } = NoteType.Comment;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    public virtual TeamMember TeamMember { get; set; } = null!;

    public string TeamMemberId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;
}
