#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Tier
{
    public TierId Id { get; set; }

    public string Name { get; set; } = null!;

    public virtual ICollection<User> Subscribers { get; set; } = new List<User>();

    public TierId? TierLimitId { get; set; }

    public virtual TierLimit? TierLimit { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
