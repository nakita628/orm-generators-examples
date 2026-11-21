#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// Bulk actions, like canceling and replaying runs
/// </summary>
public partial class BulkActionGroup
{
    public string Id { get; set; } = null!;

    public string FriendlyId { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// If this is set then it's a V2 Bulk Action that supports queries
    /// </summary>
    public virtual RuntimeEnvironment? Environment { get; set; }

    public string? EnvironmentId { get; set; }

    public BulkActionType Type { get; set; }

    /// <summary>
    /// When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed.
    /// </summary>
    public BulkActionStatus Status { get; set; } = BulkActionStatus.Pending;

    /// <summary>
    /// The query that will be executed to get the runs to process (if null, we are passing run ids directly)
    /// </summary>
    public string? QueryName { get; set; }

    /// <summary>
    /// The params that will be passed to the query
    /// </summary>
    public string? Params { get; set; }

    /// <summary>
    /// Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions.
    /// </summary>
    public string? DedupeKey { get; set; }

    /// <summary>
    /// The cursor that will be passed to the query (null for the first page)
    /// </summary>
    public string? Cursor { get; set; }

    /// <summary>
    /// The number of runs that have been processed successfully
    /// </summary>
    public int SuccessCount { get; set; } = 0;

    /// <summary>
    /// The number of runs that have been failed
    /// </summary>
    public int FailureCount { get; set; } = 0;

    /// <summary>
    /// The total number of runs that will be processed
    /// </summary>
    public int TotalCount { get; set; } = 0;

    public BulkActionNotificationType CompletionNotification { get; set; } = BulkActionNotificationType.None;

    /// <summary>
    /// The userId who did the bulk action
    /// </summary>
    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    /// <summary>
    /// The name of the bulk action
    /// </summary>
    public string? Name { get; set; }

    /// <summary>
    /// The time the bulk action was completed
    /// </summary>
    public DateTime? CompletedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
