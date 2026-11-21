#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Barcode
{
    public string Id { get; set; } = null!;

    public string Value { get; set; } = null!;

    public BarcodeType Type { get; set; } = BarcodeType.Code128;

    public virtual Asset? Asset { get; set; }

    public string? AssetId { get; set; }

    public virtual Asset? PreferredByAsset { get; set; }

    public virtual ICollection<Scan> Scans { get; set; } = new List<Scan>();

    public virtual Kit? Kit { get; set; }

    public string? KitId { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
