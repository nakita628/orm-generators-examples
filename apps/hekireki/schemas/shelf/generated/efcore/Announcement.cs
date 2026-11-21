#nullable enable

using System;

namespace Shelf.Models;

public partial class Announcement
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string Content { get; set; } = null!;

    public string? Link { get; set; }

    public string? LinkText { get; set; }

    public bool Published { get; set; } = false;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
