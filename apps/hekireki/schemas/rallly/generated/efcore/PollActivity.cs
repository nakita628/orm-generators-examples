#nullable enable

using System;

namespace Rallly.Models;

public partial class PollActivity
{
    public string Id { get; set; } = null!;

    public string PollId { get; set; } = null!;

    /// <summary>
    /// Validated by a zod union in the feature, not a DB enum: the activity vocabulary grows without migrations
    /// </summary>
    public string Type { get; set; } = null!;

    /// <summary>
    /// Subject ids are soft references (no FKs): activity must outlive its subjects — including account deletion — so refs persist as identity and the payload carries the display snapshot; whether a ref still resolves is the renderer's question
    /// </summary>
    public string? UserId { get; set; }

    public string? ParticipantId { get; set; }

    public string? InviteId { get; set; }

    public string? OptionId { get; set; }

    public string Payload { get; set; } = "{}";

    public DateTime CreatedAt { get; set; }

    public virtual Poll Poll { get; set; } = null!;
}
