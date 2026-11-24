#nullable enable

using System;

namespace Documenso.Models;

/// <summary>
/// @zod.import(["import { ZAdminTagTypeSchema } from '@documenso/lib/types/admin-tags';"])
/// </summary>
public partial class AdminTag
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// [AdminTagType] @zod.custom.use(ZAdminTagTypeSchema)
    /// </summary>
    public string Type { get; set; } = null!;

    public string Description { get; set; } = null!;

    public int? UserId { get; set; }

    public virtual User? User { get; set; }

    public string? OrganisationId { get; set; }

    public virtual Organisation? Organisation { get; set; }

    public int? CreatedByUserId { get; set; }

    public virtual User? CreatedByUser { get; set; }

    public string? AdminRuleId { get; set; }

    public virtual AdminRule? AdminRule { get; set; }
}
