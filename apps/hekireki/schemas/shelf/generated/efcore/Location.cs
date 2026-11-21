#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Location
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public string? Address { get; set; }

    public double? Latitude { get; set; }

    public double? Longitude { get; set; }

    public string? ImageUrl { get; set; }

    public string? ThumbnailUrl { get; set; }

    public virtual Image? Image { get; set; }

    public string? ImageId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual ICollection<AssetLocation> AssetLocations { get; set; } = new List<AssetLocation>();

    public virtual ICollection<Kit> Kits { get; set; } = new List<Kit>();

    public virtual ICollection<LocationNote> Notes { get; set; } = new List<LocationNote>();

    public virtual ICollection<Custody> Custodies { get; set; } = new List<Custody>();

    public virtual ICollection<ConsumptionLog> ConsumptionLogs { get; set; } = new List<ConsumptionLog>();

    public string? ParentId { get; set; }

    public virtual Location? Parent { get; set; }

    public virtual ICollection<Location> Children { get; set; } = new List<Location>();
}
