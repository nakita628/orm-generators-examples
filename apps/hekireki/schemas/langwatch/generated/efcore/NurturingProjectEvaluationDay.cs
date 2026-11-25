#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Each project and UTC day with a completed evaluation, from evaluation's lifecycle_completed,
/// written insert-ignore. Replaces the `evaluation` rows of NurturingProjectUsageView.
/// </summary>
public partial class NurturingProjectEvaluationDay
{
    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// UTC `YYYY-MM-DD`.
    /// </summary>
    public string Day { get; set; } = null!;
}
