#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
/// and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
/// so retention adds nothing to the survey row.
/// </summary>
public partial class RetentionSurveyRestore
{
    public virtual Survey Survey { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public DateTime RestoredAt { get; set; }
}
