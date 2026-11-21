#nullable enable

using System;

namespace Trigger.Models;

public partial class BulkActionItem
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// @deprecated not used in new BulkActions
    /// </summary>
    public string? FriendlyId { get; set; }

    public string GroupId { get; set; } = null!;

    public BulkActionType Type { get; set; }

    /// <summary>
    /// When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed.
    /// </summary>
    public BulkActionItemStatus Status { get; set; } = BulkActionItemStatus.Pending;

    /// <summary>
    /// The run that is the source of the action, e.g. when replaying this is the original run
    /// </summary>
    public virtual TaskRun SourceRun { get; set; } = null!;

    public string SourceRunId { get; set; } = null!;

    /// <summary>
    /// The run that's a result of the action, this will be set when the run has been created
    /// </summary>
    public virtual TaskRun? DestinationRun { get; set; }

    public string? DestinationRunId { get; set; }

    public string? Error { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
