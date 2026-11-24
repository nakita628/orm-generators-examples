#nullable enable

using System;

namespace Documenso.Models;

/// <summary>
/// @zod.import(["import { ZAdminRuleContextSchema, ZAdminRuleActionRunsSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
/// </summary>
public partial class AdminRuleMatch
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public AdminRuleMode Mode { get; set; }

    /// <summary>
    /// [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
    /// </summary>
    public string Trigger { get; set; } = null!;

    public string? OrganisationId { get; set; }

    public string? EnvelopeId { get; set; }

    /// <summary>
    /// [AdminRuleContext] @zod.custom.use(ZAdminRuleContextSchema)
    /// </summary>
    public string EventSnapshot { get; set; } = null!;

    /// <summary>
    /// [AdminRuleActionRuns] @zod.custom.use(ZAdminRuleActionRunsSchema)
    /// </summary>
    public string ActionsRun { get; set; } = null!;

    public string RuleId { get; set; } = null!;

    public virtual AdminRule Rule { get; set; } = null!;

    public int? ActorUserId { get; set; }

    public virtual User? ActorUser { get; set; }
}
