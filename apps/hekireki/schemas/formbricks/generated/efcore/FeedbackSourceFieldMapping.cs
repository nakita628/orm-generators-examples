#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
/// Maps source fields to Hub FeedbackRecord fields.
///
/// @property id - Unique identifier for the mapping
/// @property feedbackSource - The parent feedback source
/// @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
/// @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
/// @property staticValue - If set, use this value instead of reading from sourceFieldId
/// </summary>
public partial class FeedbackSourceFieldMapping
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public string FeedbackSourceId { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual FeedbackSource FeedbackSource { get; set; } = null!;

    public string SourceFieldId { get; set; } = null!;

    public string TargetFieldId { get; set; } = null!;

    public string? StaticValue { get; set; }
}
