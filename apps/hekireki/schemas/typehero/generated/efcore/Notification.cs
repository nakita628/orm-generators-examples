#nullable enable

using System;

namespace TypeHero.Models;

public partial class Notification
{
    public int Id { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public bool IsRead { get; set; } = false;

    public NotificationType Type { get; set; }

    public virtual User ToUser { get; set; } = null!;

    public string ToUserId { get; set; } = null!;

    public virtual User FromUser { get; set; } = null!;

    public string FromUserId { get; set; } = null!;

    public int? CommentId { get; set; }

    public virtual Comment? Comment { get; set; }

    /// <summary>
    /// could be comment text or solution title to avoid having to join to grab this data
    /// </summary>
    public string? Content { get; set; }

    /// <summary>
    /// can be used to replace the blurb (eg: mention you in a comment)
    /// </summary>
    public string? Blurb { get; set; }

    /// <summary>
    /// use this instead of building the url on the fly (avoids a lot of joins and perf hits to build a string)
    /// </summary>
    public string Url { get; set; } = null!;

    public virtual Challenge? Challenge { get; set; }

    public int? ChallengeId { get; set; }

    public virtual SharedSolution? SharedSolution { get; set; }

    public int? SharedSolutionId { get; set; }
}
