#nullable enable

using System;

namespace Rallly.Models;

public partial class PollInvite
{
    public string Id { get; set; } = null!;

    public string PollId { get; set; } = null!;

    public string Email { get; set; } = null!;

    /// <summary>
    /// Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access
    /// </summary>
    public string Token { get; set; } = null!;

    public DateTime? OpenedAt { get; set; }

    public DateTime? RemindedAt { get; set; }

    /// <summary>
    /// Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token
    /// </summary>
    public DateTime? RevokedAt { get; set; }

    /// <summary>
    /// Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite
    /// </summary>
    public string? ParticipantId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Poll Poll { get; set; } = null!;

    public virtual Participant? Participant { get; set; }
}
