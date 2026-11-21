#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Category
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public string Color { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<Asset> Assets { get; set; } = new List<Asset>();

    public virtual ICollection<Kit> Kits { get; set; } = new List<Kit>();

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual ICollection<CustomField> CustomFields { get; set; } = new List<CustomField>();

    public virtual ICollection<AssetModel> AssetModelDefaults { get; set; } = new List<AssetModel>();

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;
}
