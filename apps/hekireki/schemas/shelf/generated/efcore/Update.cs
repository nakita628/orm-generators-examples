#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Update
{
    public string Id { get; set; } = null!;

    public string Title { get; set; } = null!;

    public string Content { get; set; } = null!;

    public string? Url { get; set; }

    public string? ImageUrl { get; set; }

    public DateTime PublishDate { get; set; }

    public UpdateStatus Status { get; set; } = UpdateStatus.Draft;

    public List<OrganizationRoles>? TargetRoles { get; set; }

    public int ClickCount { get; set; } = 0;

    public int ViewCount { get; set; } = 0;

    public virtual ICollection<UserUpdateRead> UserReads { get; set; } = new List<UserUpdateRead>();

    public virtual User CreatedBy { get; set; } = null!;

    public string CreatedById { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
