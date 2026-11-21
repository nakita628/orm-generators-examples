#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Role
{
    public string Id { get; set; } = null!;

    public Roles Name { get; set; } = Roles.User;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<User> Users { get; set; } = new List<User>();
}
