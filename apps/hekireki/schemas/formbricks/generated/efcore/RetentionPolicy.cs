#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// One policy per organisation and kind of data: a notice `warnDays` before the action, and the action
/// `periodDays` after the target's clock. The action follows from the kind of data: responses are
/// deleted, surveys archived (the archive purge deletes them 30 days later, as it does every archived
/// survey), members deactivated. The dates themselves are never stored; they are computed by the shared
/// helper in `apps/web/modules/ee/data-retention/lib/schedule.ts`, so the UI, the emails and the sweep
/// always agree.
///
/// The allowed ranges are validated by the API. The migration repeats the ones that guard against
/// acting early or re-notifying (a notice of at least 14 days and shorter than the period, an enabled
/// policy always has `enabledAt`) and the structural ones (conditions on the surveys policy only) as
/// CHECK constraints, which Prisma cannot express.
/// </summary>
public partial class RetentionPolicy
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public RetentionEntity Entity { get; set; }

    public bool Enabled { get; set; } = false;

    /// <summary>
    /// When the current configuration took effect: the policy was switched on or unpaused, or changed
    /// in a way that could bring an action earlier (a shorter period or warning, different
    /// conditions). Lengthening does not move it. The warning always runs in full from here, so no
    /// change can delete data on the night it is made (ENG-3614). Also the clock for members with no
    /// recorded sign-in.
    /// </summary>
    public DateTime? EnabledAt { get; set; }

    /// <summary>
    /// How long before the action the notice goes out. The action never runs before the notice has
    /// been out for this long, so the warning always runs in full.
    /// </summary>
    public int WarnDays { get; set; } = 60;

    /// <summary>
    /// The period, counted from the target's clock, after which the policy acts: responses are
    /// deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
    /// </summary>
    public int PeriodDays { get; set; }

    /// <summary>
    /// Surveys policy only.
    /// </summary>
    public List<RetentionSurveyCondition>? Conditions { get; set; } = new List<RetentionSurveyCondition>();

    /// <summary>
    /// The last night the sweep's time ran out before it reached this policy's organisation, licensed.
    /// The organisation was due and the sweep was running, so it counts as reached for the gap that
    /// restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
    /// </summary>
    public DateTime? DeferredAt { get; set; }

    public virtual User? UpdatedBy { get; set; }

    public string? UpdatedById { get; set; }
}
