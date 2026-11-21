#nullable enable

namespace Formbricks.Models;

/// <summary>
/// What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
/// never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
/// `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
/// </summary>
public partial class RetentionRunItem
{
    public string Id { get; set; } = null!;

    public virtual RetentionRun Run { get; set; } = null!;

    public string RunId { get; set; } = null!;

    public RetentionTargetType TargetType { get; set; }

    public string TargetId { get; set; } = null!;

    public string? TargetName { get; set; }

    public RetentionRunItemAction Action { get; set; }

    /// <summary>
    /// 1, except for grouped response deletions, where it is the number of responses deleted.
    /// </summary>
    public int Count { get; set; } = 1;

    /// <summary>
    /// Who was emailed, for `notified` rows.
    /// </summary>
    public string? Recipient { get; set; }

    public RetentionSkipReason? SkipReason { get; set; }
}
