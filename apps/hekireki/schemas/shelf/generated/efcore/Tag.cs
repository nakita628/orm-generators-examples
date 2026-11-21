#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Tag
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public string? Color { get; set; }

    public List<TagUseFor>? UseFor { get; set; } = new List<TagUseFor> { TagUseFor.Asset };

    public virtual ICollection<Asset> Assets { get; set; } = new List<Asset>();

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual ICollection<Booking> Bookings { get; set; } = new List<Booking>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
