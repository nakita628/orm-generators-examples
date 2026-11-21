#nullable enable

using System;

namespace Shelf.Models;

public partial class ScimToken
{
    public string Id { get; set; } = null!;

    public string TokenHash { get; set; } = null!;

    public string Label { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string CreatedById { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime? LastUsedAt { get; set; }
}
