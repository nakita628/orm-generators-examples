#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// D08 — one connection's directory sync, as the `scim-sync` pipeline's fold
/// writes it.
///
/// Pure event truth: every column is fold-written, rows are never deleted,
/// and REVOKED is a tombstone rather than an absence — a dead letter that
/// vanished when a connection was torn down would be a removal nobody could
/// check afterwards.
///
/// No relation to Organization or SsoConnection, matching `SsoConnection`
/// itself: a projection is rebuilt from events, and a foreign key would make
/// a replay depend on relational state the log does not carry.
/// </summary>
public partial class ScimSyncState
{
    /// <summary>
    /// The aggregate: the connection's id, which IS its sync's id.
    /// </summary>
    public string Id { get; set; } = null!;

    public string ConnectionId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// TOKEN_ISSUED | SYNCING | ERROR | REVOKED
    /// </summary>
    public string State { get; set; } = null!;

    public DateTime? LastPushedAt { get; set; }

    /// <summary>
    /// The failure standing right now: { op, errorCode, attempts, retiredAtMs,
    /// userId, occurredAtMs }. A reason CODE and ids — never a token, never a
    /// secret, never an internal hostname.
    /// </summary>
    public string? LastFailure { get; set; }

    /// <summary>
    /// Every failure retired without ever being applied, newest last. The dead
    /// letters an administrator has to see; retiring never drops one.
    /// </summary>
    public string DeadLetters { get; set; } = "[]";

    /// <summary>
    /// revoke | teardown, once the sync has ended.
    /// </summary>
    public string? RevokedCause { get; set; }

    public DateTime OccurredAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public DateTime AcceptedAt { get; set; }

    public string ProjectionVersion { get; set; } = null!;

    /// <summary>
    /// Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    /// would make a replayed row differ from the row it rebuilds, which is
    /// exactly the parity this projection has to keep.
    /// </summary>
    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
