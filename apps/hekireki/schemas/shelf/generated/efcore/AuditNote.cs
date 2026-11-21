#nullable enable

using System;

namespace Shelf.Models;

public partial class AuditNote
{
    public string Id { get; set; } = null!;

    public string Content { get; set; } = null!;

    public NoteType Type { get; set; } = NoteType.Comment;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    public virtual AuditSession AuditSession { get; set; } = null!;

    public string AuditSessionId { get; set; } = null!;

    public virtual AuditAsset? AuditAsset { get; set; }

    public string? AuditAssetId { get; set; }
}
