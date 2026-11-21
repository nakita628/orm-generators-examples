#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class AuditSession
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public string? TargetId { get; set; }

    public AuditStatus Status { get; set; } = AuditStatus.Pending;

    public string? ScopeMeta { get; set; }

    public int ExpectedAssetCount { get; set; } = 0;

    public int FoundAssetCount { get; set; } = 0;

    public int MissingAssetCount { get; set; } = 0;

    public int UnexpectedAssetCount { get; set; } = 0;

    public DateTime? StartedAt { get; set; }

    public DateTime? DueDate { get; set; }

    public DateTime? CompletedAt { get; set; }

    public DateTime? CancelledAt { get; set; }

    public string? ActiveSchedulerReference { get; set; }

    public virtual User CreatedBy { get; set; } = null!;

    public string CreatedById { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual ICollection<AuditAssignment> Assignments { get; set; } = new List<AuditAssignment>();

    public virtual ICollection<AuditAsset> Assets { get; set; } = new List<AuditAsset>();

    public virtual ICollection<AuditScan> Scans { get; set; } = new List<AuditScan>();

    public virtual ICollection<AuditNote> Notes { get; set; } = new List<AuditNote>();

    public virtual ICollection<AuditImage> Images { get; set; } = new List<AuditImage>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
