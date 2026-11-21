#nullable enable

using System;

namespace Shelf.Models;

public partial class BookingNote
{
    public string Id { get; set; } = null!;

    public string Content { get; set; } = null!;

    public NoteType Type { get; set; } = NoteType.Comment;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    public virtual Booking Booking { get; set; } = null!;

    public string BookingId { get; set; } = null!;
}
