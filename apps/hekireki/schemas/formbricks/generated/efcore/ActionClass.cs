#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a user action that can trigger surveys.
/// Used to define points in the user journey where surveys can be shown.
///
/// @property id - Unique identifier for the action
/// @property name - Display name of the action
/// @property type - Whether this is a code or no-code action
/// @property key - Unique identifier used in code implementation
/// @property noCodeConfig - Configuration for no-code setup
/// @property workspace - The workspace this action belongs to
/// </summary>
public partial class ActionClass
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public ActionType Type { get; set; }

    public string? Key { get; set; }

    /// <summary>
    /// [ActionClassNoCodeConfig]
    /// </summary>
    public string? NoCodeConfig { get; set; }

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual ICollection<SurveyTrigger> SurveyTriggers { get; set; } = new List<SurveyTrigger>();

    public virtual ICollection<CreativeTrigger> CreativeTriggers { get; set; } = new List<CreativeTrigger>();
}
