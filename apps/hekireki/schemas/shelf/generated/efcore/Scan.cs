#nullable enable

using System;

namespace Shelf.Models;

public partial class Scan
{
    public string Id { get; set; } = null!;

    public string? Latitude { get; set; }

    public string? Longitude { get; set; }

    public string? UserAgent { get; set; }

    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    public virtual Qr? Qr { get; set; }

    public string? QrId { get; set; }

    public string? RawQrId { get; set; }

    public ScanCodeType CodeType { get; set; } = ScanCodeType.Qr;

    public string? Code { get; set; }

    public virtual Barcode? Barcode { get; set; }

    public string? BarcodeId { get; set; }

    public virtual Asset? Asset { get; set; }

    public string? AssetId { get; set; }

    public virtual Kit? Kit { get; set; }

    public string? KitId { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public ScanSource? Source { get; set; }

    public bool ManuallyGenerated { get; set; } = false;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
