#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Defines follow-up actions for survey responses.
/// Enables automated actions based on specific survey response conditions.
///
/// @property id - Unique identifier for the follow-up
/// @property survey - The associated survey
/// @property name - Display name for the follow-up
/// @property trigger - Conditions that activate the follow-up
/// @property action - Actions to take when triggered
/// </summary>
public partial class SurveyFollowUp
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Survey Survey { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public string Name { get; set; } = null!;

    /// <summary>
    /// [SurveyFollowUpTrigger]
    /// </summary>
    public string Trigger { get; set; } = null!;

    /// <summary>
    /// [SurveyFollowUpAction]
    /// </summary>
    public string Action { get; set; } = null!;
}
