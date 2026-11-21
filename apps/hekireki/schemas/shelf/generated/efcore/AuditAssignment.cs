#nullable enable

using System;

namespace Shelf.Models;

public partial class AuditAssignment
{
    public string Id { get; set; } = null!;

    public virtual AuditSession AuditSession { get; set; } = null!;

    public string AuditSessionId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public AuditAssignmentRole? Role { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
