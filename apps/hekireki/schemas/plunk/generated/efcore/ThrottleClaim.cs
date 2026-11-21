#nullable enable

using System;

namespace Plunk.Models;

/// <summary>
/// One transactional notification held for one recipient (or for the whole project) by a
/// send with `delivery.throttle`. While it is held, the same notification is not sent again.
///
/// The claim is taken by inserting (or taking over an expired) row, so the unique constraint
/// decides the race between concurrent sends, like `IdempotencyKey`. It is held from the
/// moment the send is accepted, re-anchored to the real send time once SES has the message,
/// and released when the email never gets there (cancelled, or failed before sending).
/// </summary>
public partial class ThrottleClaim
{
    public string Id { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string Key { get; set; } = null!;

    public string RecipientKey { get; set; } = null!;

    public virtual Contact? Contact { get; set; }

    public string? ContactId { get; set; }

    public string? EmailId { get; set; }

    public virtual Workflow? Workflow { get; set; }

    public string? WorkflowId { get; set; }

    public int? WindowSeconds { get; set; }

    public DateTime? ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
