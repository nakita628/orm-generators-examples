#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Records when a survey is shown to a user.
/// Tracks survey display history and response status.
///
/// @property id - Unique identifier for the display event
/// @property survey - The survey that was displayed
/// @property contact - The contact who saw the survey
/// @property response - The associated response if one exists
/// @property creativeId - The creative that opened the survey. No foreign key, so the record survives the creative's deletion.
/// </summary>
public partial class Display
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Survey Survey { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public virtual Contact? Contact { get; set; }

    public string? ContactId { get; set; }

    public virtual Response? Response { get; set; }

    public string? CreativeId { get; set; }
}
