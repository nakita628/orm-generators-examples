#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a quota configuration for a survey.
/// Defines response limits and conditions for quota management.
///
/// @property id - Unique identifier for the quota
/// @property survey - The survey this quota belongs to
/// @property name - Display name for the quota
/// @property limit - Maximum number of responses allowed
/// @property conditions - JSON object containing quota conditions
/// @property action - Action to take when quota is reached
/// @property endingCardId - Optional ending card to show when quota is reached
/// </summary>
public partial class SurveyQuota
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Survey Survey { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public string Name { get; set; } = null!;

    public int Limit { get; set; }

    /// <summary>
    /// [SurveyQuotaLogic]
    /// </summary>
    public string Logic { get; set; } = "{}";

    public SurveyQuotaAction Action { get; set; }

    public string? EndingCardId { get; set; }

    public virtual ICollection<ResponseQuotaLink> QuotaLinks { get; set; } = new List<ResponseQuotaLink>();

    public bool CountPartialSubmissions { get; set; } = false;
}
