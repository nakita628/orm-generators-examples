#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Asset
{
    public string Id { get; set; } = null!;

    public string Title { get; set; } = null!;

    public string? Description { get; set; }

    public string? MainImage { get; set; }

    public string? ThumbnailImage { get; set; }

    public DateTime? MainImageExpiration { get; set; }

    public AssetStatus Status { get; set; } = AssetStatus.Available;

    public double? Valuation { get; set; }

    public bool AvailableToBook { get; set; } = true;

    public string? SequentialId { get; set; }

    public DateTime? ArchivedAt { get; set; }

    public AssetType Type { get; set; } = AssetType.Individual;

    public string? AssetModelId { get; set; }

    public virtual AssetModel? AssetModel { get; set; }

    public int? Quantity { get; set; }

    public int? MinQuantity { get; set; }

    public ConsumptionType? ConsumptionType { get; set; }

    public string? UnitOfMeasure { get; set; }

    public DateTime? LowStockNotifiedAt { get; set; }

    public DateTime? StockLedgerStartedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Category? Category { get; set; }

    public string? CategoryId { get; set; }

    public virtual ICollection<Custody> Custody { get; set; } = new List<Custody>();

    public virtual ICollection<Note> Notes { get; set; } = new List<Note>();

    public virtual ICollection<Qr> QrCodes { get; set; } = new List<Qr>();

    public virtual ICollection<Barcode> Barcodes { get; set; } = new List<Barcode>();

    public virtual ICollection<ReportFound> Reports { get; set; } = new List<ReportFound>();

    public virtual ICollection<Tag> Tags { get; set; } = new List<Tag>();

    public virtual ICollection<AssetCustomFieldValue> CustomFields { get; set; } = new List<AssetCustomFieldValue>();

    public virtual ICollection<BookingAsset> BookingAssets { get; set; } = new List<BookingAsset>();

    public virtual ICollection<ConsumptionLog> ConsumptionLogs { get; set; } = new List<ConsumptionLog>();

    public virtual ICollection<AssetReminder> Reminders { get; set; } = new List<AssetReminder>();

    public virtual ICollection<AuditAsset> AuditEntries { get; set; } = new List<AuditAsset>();

    public virtual ICollection<AuditScan> AuditScans { get; set; } = new List<AuditScan>();

    public virtual ICollection<Scan> Scans { get; set; } = new List<Scan>();

    public virtual ICollection<AssetKit> AssetKits { get; set; } = new List<AssetKit>();

    public virtual ICollection<AssetLocation> AssetLocations { get; set; } = new List<AssetLocation>();

    public string? PreferredBarcodeId { get; set; }

    public virtual Barcode? PreferredBarcode { get; set; }
}
