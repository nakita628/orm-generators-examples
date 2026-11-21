#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Links a creative to the actions that show it (`showOn: action`). Mirrors `SurveyTrigger`.
/// </summary>
public partial class CreativeTrigger
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Creative Creative { get; set; } = null!;

    public string CreativeId { get; set; } = null!;

    public virtual ActionClass ActionClass { get; set; } = null!;

    public string ActionClassId { get; set; } = null!;
}
