#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Postgres-authoritative one-turn admission for a conversation. PREPARING
/// claims are leased so a crash before the durable start event can recover;
/// COMMITTED claims are released only by the matching terminal event.
/// </summary>
public partial class LangyActiveTurn
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string ConversationId { get; set; } = null!;

    public string TurnId { get; set; } = null!;

    public string RequestId { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public string Status { get; set; } = null!;

    public string LeaseOwner { get; set; } = null!;

    public DateTime LeaseExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
