#nullable enable

using System;

namespace Shelf.Models;

public partial class UserUpdateRead
{
    public string Id { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual Update Update { get; set; } = null!;

    public string UpdateId { get; set; } = null!;

    public DateTime ReadAt { get; set; }
}
