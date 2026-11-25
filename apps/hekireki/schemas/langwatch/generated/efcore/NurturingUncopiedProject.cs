#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// retired: unread since v3.28.0. The trace-day folds count a trace with no placement copy, so no
/// retry clock reads or writes this table; an operator step may drop it once old writers are gone.
/// </summary>
public partial class NurturingUncopiedProject
{
    public string ProjectId { get; set; } = null!;

    public DateTime FirstSeenAt { get; set; }
}
