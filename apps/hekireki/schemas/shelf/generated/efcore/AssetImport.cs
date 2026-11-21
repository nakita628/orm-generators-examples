#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// One CSV content import. Holds the parsed rows from preview until the import
/// completes, so the job imports exactly the file the user previewed.
/// A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
/// active import per workspace. Prisma cannot express it; it lives in migration
/// 20261015120100_asset_import_one_active_per_workspace.
/// </summary>
public partial class AssetImport
{
    public string Id { get; set; } = null!;

    public AssetImportStatus Status { get; set; } = AssetImportStatus.Draft;

    public AssetImportPhase Phase { get; set; } = AssetImportPhase.Prepare;

    public string? FileName { get; set; }

    public int TotalRows { get; set; }

    /// <summary>
    /// Rows committed so far. Advanced in the same transaction as each chunk.
    /// </summary>
    public int ProcessedRows { get; set; } = 0;

    /// <summary>
    /// Parsed rows (`StoredImportRow[]`). Cleared when the import completes.
    /// </summary>
    public string? Rows { get; set; }

    /// <summary>
    /// Name to id maps written by the PREPARE step (`ResolvedImportEntities`).
    /// </summary>
    public string? Resolved { get; set; }

    public bool CanUseBarcodes { get; set; } = false;

    /// <summary>
    /// User-facing reason the import failed. Null unless status is FAILED.
    /// </summary>
    public string? ErrorMessage { get; set; }

    /// <summary>
    /// Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500.
    /// </summary>
    public string? Warnings { get; set; }

    /// <summary>
    /// Exact number of warnings, including any beyond the stored cap.
    /// </summary>
    public int WarningCount { get; set; } = 0;

    /// <summary>
    /// A step holds the import until this time. Null when no step is running.
    /// </summary>
    public DateTime? LockedUntil { get; set; }

    public DateTime? ConfirmedAt { get; set; }

    public DateTime? StartedAt { get; set; }

    public DateTime? CompletedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User CreatedBy { get; set; } = null!;

    public string CreatedById { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;
}
