#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

/// <summary>
/// Schedules can be attached to tasks to trigger them on a schedule
/// </summary>
public partial class TaskSchedule
{
    public string Id { get; set; } = null!;

    public ScheduleType Type { get; set; } = ScheduleType.Imperative;

    /// <summary>
    /// users see this as `id`. They start with schedule_
    /// </summary>
    public string FriendlyId { get; set; } = null!;

    /// <summary>
    /// a reference to a task (not a foreign key because it's across versions)
    /// </summary>
    public string TaskIdentifier { get; set; } = null!;

    /// <summary>
    /// can be provided and we won't create another with the same key
    /// </summary>
    public string DeduplicationKey { get; set; } = null!;

    public bool UserProvidedDeduplicationKey { get; set; } = false;

    /// <summary>
    /// the CRON pattern
    /// </summary>
    public string GeneratorExpression { get; set; } = null!;

    public string GeneratorDescription { get; set; } = "";

    public ScheduleGeneratorType GeneratorType { get; set; } = ScheduleGeneratorType.Cron;

    /// <summary>
    /// These are IANA format string, or the default "UTC". E.g. "America/New_York"
    /// </summary>
    public string Timezone { get; set; } = "UTC";

    public int? WindowDurationSeconds { get; set; }

    public int? WindowPercentage { get; set; }

    public int? DefaultWindowDurationSeconds { get; set; }

    public int? MinimumWindowDurationSeconds { get; set; }

    /// <summary>
    /// Can be provided by the user then accessed inside a run
    /// </summary>
    public string? ExternalId { get; set; }

    /// <summary>
    /// Instances of the schedule that are active
    /// </summary>
    public virtual ICollection<TaskScheduleInstance> Instances { get; set; } = new List<TaskScheduleInstance>();

    /// <summary>
    /// @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up.
    /// </summary>
    public DateTime? LastRunTriggeredAt { get; set; }

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public bool Active { get; set; } = true;
}
