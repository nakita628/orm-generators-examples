#nullable enable

using System;

namespace Shelf.Models;

public partial class AuditScan
{
    public string Id { get; set; } = null!;

    public virtual AuditSession AuditSession { get; set; } = null!;

    public string AuditSessionId { get; set; } = null!;

    public virtual AuditAsset? AuditAsset { get; set; }

    public string? AuditAssetId { get; set; }

    public virtual Asset? Asset { get; set; }

    public string? AssetId { get; set; }

    public virtual User? ScannedBy { get; set; }

    public string? ScannedById { get; set; }

    public string? Code { get; set; }

    /// <summary>
    /// What was scanned, recorded BY VALUE at scan time.
    ///
    /// `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset
    /// leaves this row pointing at nothing: no name, and no way to tell whether
    /// it had been expected. An audit is a historical record, and a record whose
    /// entries can silently lose their meaning is not one — so the two
    /// descriptive facts are snapshotted here rather than only referenced.
    ///
    /// Readers prefer the LIVE asset when it still exists (a rename should show
    /// the current name) and fall back to these.
    ///
    /// Null means the snapshot could not be taken: scans whose asset had already
    /// been deleted before these columns were backfilled. Nothing can recover
    /// those, and they degrade to the scanned `code`, which also survives. Every
    /// other row carries a snapshot, so null is never merely "old".
    /// </summary>
    public string? AssetTitle { get; set; }

    public bool? WasExpected { get; set; }

    public string? Metadata { get; set; }

    public DateTime ScannedAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
