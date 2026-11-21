#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Qr
{
    public string Id { get; set; } = null!;

    public int Version { get; set; } = 0;

    public ErrorCorrection ErrorCorrection { get; set; } = ErrorCorrection.L;

    public virtual Asset? Asset { get; set; }

    public string? AssetId { get; set; }

    public virtual Kit? Kit { get; set; }

    public string? KitId { get; set; }

    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    public virtual Organization? Organization { get; set; }

    public string? OrganizationId { get; set; }

    public virtual PrintBatch? Batch { get; set; }

    public string? BatchId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<Scan> Scans { get; set; } = new List<Scan>();
}
