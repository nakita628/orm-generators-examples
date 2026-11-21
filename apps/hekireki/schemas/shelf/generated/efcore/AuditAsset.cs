#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class AuditAsset
{
    public string Id { get; set; } = null!;

    public virtual AuditSession AuditSession { get; set; } = null!;

    public string AuditSessionId { get; set; } = null!;

    public virtual Asset Asset { get; set; } = null!;

    public string AssetId { get; set; } = null!;

    public bool Expected { get; set; } = true;

    public AuditAssetStatus Status { get; set; } = AuditAssetStatus.Pending;

    public virtual User? ScannedBy { get; set; }

    public string? ScannedById { get; set; }

    public DateTime? ScannedAt { get; set; }

    public string? Metadata { get; set; }

    /// <summary>
    /// Units the audit expects to find, snapshotted when the row is created: the
    /// units in the audit's scope (placed at its locations, held in its kits, or
    /// held by its person; at most `Asset.quantity`), or `Asset.quantity` when
    /// none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
    /// which count as one.
    ///
    /// A snapshot, not a live read: stock that moves during the audit must not
    /// move the target the auditor is counting against.
    /// </summary>
    public int? ExpectedQuantity { get; set; }

    /// <summary>
    /// Units the auditor counted. Null until a quantity-tracked asset is
    /// counted, and always null for INDIVIDUAL assets. A FOUND row with a count
    /// below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
    /// was located), the gap is this number.
    ///
    /// Both columns are non-negative, enforced by CHECK constraints in raw SQL
    /// (migration `20261104120000_add_audit_asset_quantities`).
    /// </summary>
    public int? FoundQuantity { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<AuditScan> Scans { get; set; } = new List<AuditScan>();

    public virtual ICollection<AuditImage> Images { get; set; } = new List<AuditImage>();

    public virtual ICollection<AuditNote> Notes { get; set; } = new List<AuditNote>();
}
