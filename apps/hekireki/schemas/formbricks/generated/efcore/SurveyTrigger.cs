#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Links surveys to specific trigger actions.
/// Defines which user actions should cause a survey to be displayed.
/// This is the connection table between Surveys and ActionClasses that determines
/// when and under what conditions a survey should be triggered.
///
/// @property id - Unique identifier for the trigger
/// @property survey - The survey to be triggered
/// @property actionClass - The action that triggers the survey
/// @property createdAt - When the trigger was created
/// @property updatedAt - When the trigger was last modified
/// </summary>
public partial class SurveyTrigger
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Survey Survey { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public virtual ActionClass ActionClass { get; set; } = null!;

    public string ActionClassId { get; set; } = null!;
}
