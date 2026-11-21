#nullable enable

using System;
using System.Collections.Generic;

namespace Plunk.Models;

public partial class Email
{
    public string Id { get; set; } = null!;

    public virtual Contact Contact { get; set; } = null!;

    public string ContactId { get; set; } = null!;

    public string? ToName { get; set; }

    public string Subject { get; set; } = null!;

    public string Body { get; set; } = null!;

    public string From { get; set; } = null!;

    public string? FromName { get; set; }

    public string? ReplyTo { get; set; }

    public string? Headers { get; set; }

    public string? Attachments { get; set; }

    public string? MessageId { get; set; }

    public EmailSourceType SourceType { get; set; }

    public virtual Template? Template { get; set; }

    public string? TemplateId { get; set; }

    public virtual Campaign? Campaign { get; set; }

    public string? CampaignId { get; set; }

    public virtual WorkflowExecution? WorkflowExecution { get; set; }

    public string? WorkflowExecutionId { get; set; }

    public virtual WorkflowStepExecution? WorkflowStepExecution { get; set; }

    public string? WorkflowStepExecutionId { get; set; }

    public EmailStatus Status { get; set; } = EmailStatus.Pending;

    /// <summary>
    /// When a transactional email sent with `delivery.at` is due. Only meaningful while the
    /// status is SCHEDULED; the delayed queue job is the timer and the scheduled-email sweep
    /// re-queues any row still SCHEDULED past this time. The partial index on SCHEDULED rows
    /// is built by the background index builder, not a migration (see background-index-builder.ts).
    /// </summary>
    public DateTime? ScheduledFor { get; set; }

    public DateTime? SentAt { get; set; }

    public DateTime? DeliveredAt { get; set; }

    public DateTime? OpenedAt { get; set; }

    public DateTime? ClickedAt { get; set; }

    public DateTime? BouncedAt { get; set; }

    public DateTime? ComplainedAt { get; set; }

    /// <summary>
    /// When the recipient opted out from this email (unsubscribe or snooze). Only set for
    /// opt-outs attributed to a campaign email and counted in `Campaign.unsubscribedCount`,
    /// in the same transaction, so the stamped rows are the list behind that figure. Indexed
    /// by a partial index built in the background (see background-index-builder.ts).
    /// </summary>
    public DateTime? UnsubscribedAt { get; set; }

    public int Opens { get; set; } = 0;

    public int Clicks { get; set; } = 0;

    /// <summary>
    /// The address handed to SES belonged to the mailbox simulator
    /// (`@simulator.amazonses.com`), stamped at send time from the address actually used --
    /// which is the recipient override when one is set, not necessarily `contact.email`.
    ///
    /// AWS itself excludes simulator mail from bounce rates, complaint rates, sending quotas
    /// and reputation metrics, so Plunk's security rates exclude it from both the numerator
    /// and the denominator. The events themselves are untouched: the webhook fires, the
    /// activity feed shows the bounce, the contact is still suppressed.
    /// </summary>
    public bool Simulated { get; set; } = false;

    public string? Error { get; set; }

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual ICollection<Event> Events { get; set; } = new List<Event>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
