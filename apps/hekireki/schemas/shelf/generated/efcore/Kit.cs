#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Kit
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public KitStatus Status { get; set; } = KitStatus.Available;

    public string? Image { get; set; }

    public DateTime? ImageExpiration { get; set; }

    public virtual ICollection<AssetKit> AssetKits { get; set; } = new List<AssetKit>();

    public virtual KitCustody? Custody { get; set; }

    public virtual ICollection<Qr> QrCodes { get; set; } = new List<Qr>();

    public virtual ICollection<Barcode> Barcodes { get; set; } = new List<Barcode>();

    public virtual ICollection<ReportFound> Reports { get; set; } = new List<ReportFound>();

    public virtual ICollection<Scan> Scans { get; set; } = new List<Scan>();

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual User CreatedBy { get; set; } = null!;

    public string CreatedById { get; set; } = null!;

    public virtual Category? Category { get; set; }

    public string? CategoryId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Location? Location { get; set; }

    public string? LocationId { get; set; }
}
