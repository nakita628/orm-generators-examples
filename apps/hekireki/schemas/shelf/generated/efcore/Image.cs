#nullable enable

using System;

namespace Shelf.Models;

public partial class Image
{
    public string Id { get; set; } = null!;

    public string ContentType { get; set; } = null!;

    public string? AltText { get; set; }

    public byte[] Blob { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Location? Location { get; set; }

    public virtual Organization? Organization { get; set; }

    public virtual Organization OwnerOrg { get; set; } = null!;

    public string OwnerOrgId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;
}
