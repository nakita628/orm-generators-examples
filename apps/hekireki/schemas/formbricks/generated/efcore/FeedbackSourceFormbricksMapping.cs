#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Maps survey elements to Hub FeedbackRecords for Formbricks feedback sources.
/// Each row represents one element that will create FeedbackRecords when answered.
///
/// @property id - Unique identifier for the mapping
/// @property feedbackSource - The parent feedback source
/// @property survey - The survey containing the element
/// @property elementId - The element ID within the survey (from blocks[].elements[].id)
/// @property hubFieldType - The field_type to use in Hub (text, nps, rating, etc.)
/// @property customFieldLabel - Optional override for the element headline as field_label in Hub
/// </summary>
public partial class FeedbackSourceFormbricksMapping
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public string FeedbackSourceId { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual FeedbackSource FeedbackSource { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public virtual Survey Survey { get; set; } = null!;

    public string ElementId { get; set; } = null!;

    public HubFieldType HubFieldType { get; set; }

    public string? CustomFieldLabel { get; set; }
}
