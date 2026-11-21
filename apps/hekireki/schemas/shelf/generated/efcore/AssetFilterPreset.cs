#nullable enable

using System;

namespace Shelf.Models;

public partial class AssetFilterPreset
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string OwnerId { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string Query { get; set; } = null!;

    public bool Starred { get; set; } = false;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public virtual User Owner { get; set; } = null!;
}
