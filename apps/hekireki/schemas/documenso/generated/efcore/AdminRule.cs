#nullable enable

using System;
using System.Collections.Generic;

namespace Documenso.Models;

/// <summary>
/// @zod.import(["import { ZAdminRuleDefinitionSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
/// </summary>
public partial class AdminRule
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public AdminRuleMode Mode { get; set; }

    /// <summary>
    /// [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
    /// </summary>
    public string Trigger { get; set; } = null!;

    /// <summary>
    /// [AdminRuleDefinition] @zod.custom.use(ZAdminRuleDefinitionSchema)
    /// </summary>
    public string Definition { get; set; } = null!;

    public int Priority { get; set; }

    public bool StopOnMatch { get; set; }

    /// <summary>
    /// Escape hatch: organisations this rule is skipped for. Only meaningful on triggers that have an
    /// organisation
    /// </summary>
    public List<string>? ExemptOrganisationIds { get; set; } = new List<string>();

    public virtual ICollection<AdminRuleMatch> Matches { get; set; } = new List<AdminRuleMatch>();

    public virtual ICollection<AdminTag> AdminTags { get; set; } = new List<AdminTag>();
}
