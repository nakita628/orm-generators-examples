#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class SsoDetails
{
    public string Id { get; set; } = null!;

    public string Domain { get; set; } = null!;

    public virtual ICollection<Organization> Organizations { get; set; } = new List<Organization>();

    public string? BaseUserGroupId { get; set; }

    public string? SelfServiceGroupId { get; set; }

    public string? AdminGroupId { get; set; }

    public string? ManagerGroupId { get; set; }

    public bool RequireSsoLogin { get; set; } = true;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
