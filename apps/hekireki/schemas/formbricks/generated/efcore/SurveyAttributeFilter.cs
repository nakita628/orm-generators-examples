#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Defines targeting rules for surveys based on contact attributes.
/// Used to show surveys only to contacts matching specific criteria.
///
/// @property id - Unique identifier for the filter
/// @property attributeKey - The contact attribute to filter on
/// @property survey - The survey being filtered
/// @property condition - The comparison operator to use
/// @property value - The value to compare against
/// </summary>
public partial class SurveyAttributeFilter
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ContactAttributeKey AttributeKey { get; set; } = null!;

    public string AttributeKeyId { get; set; } = null!;

    public virtual Survey Survey { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public SurveyAttributeFilterCondition Condition { get; set; }

    public string Value { get; set; } = null!;
}
