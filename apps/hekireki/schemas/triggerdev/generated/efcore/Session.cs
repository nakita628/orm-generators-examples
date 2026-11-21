#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

/// <summary>
/// Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
/// The row is essentially static — no status, no counters, no pointers. No
/// foreign keys: project/runtimeEnvironment/organization ids are plain
/// scalar columns (matches TaskRun pattern). List-style queries are served
/// from ClickHouse, not Postgres, so only point-lookup indexes live here.
/// </summary>
public partial class Session
{
    public string Id { get; set; } = null!;

    public string FriendlyId { get; set; } = null!;

    /// <summary>
    /// User-supplied identifier scoped to the environment. Used for
    /// idempotent upsert and for resolving sessions via the public API.
    /// </summary>
    public string? ExternalId { get; set; }

    /// <summary>
    /// Plain string — intentionally not an enum.
    /// </summary>
    public string Type { get; set; } = null!;

    /// <summary>
    /// Denormalized scoping columns — no FK relations.
    /// </summary>
    public string ProjectId { get; set; } = null!;

    public string RuntimeEnvironmentId { get; set; } = null!;

    public RuntimeEnvironmentType EnvironmentType { get; set; }

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// Task this session triggers runs against. Required — Sessions are
    /// task-bound: creating a session also triggers its first run, and
    /// every subsequent re-trigger uses this same identifier.
    /// </summary>
    public string TaskIdentifier { get; set; } = null!;

    /// <summary>
    /// Trigger config used for every run this session schedules. Shape
    /// (validated at the route layer, opaque to the DB):
    /// { basePayload: object, machine?: string, queue?: string,
    /// tags?: string[], maxAttempts?: number,
    /// idleTimeoutInSeconds?: number }
    /// `basePayload` carries the customer's client-data; runtime fields
    /// (chatId, messages, trigger) are merged at trigger time.
    /// </summary>
    public string TriggerConfig { get; set; } = null!;

    /// <summary>
    /// Whether this session was created from the Test/playground UI rather
    /// than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
    /// </summary>
    public bool IsTest { get; set; } = false;

    public List<string>? Tags { get; set; } = new List<string>();

    public string? Metadata { get; set; }

    /// <summary>
    /// Live run pointer — non-FK so run deletion never cascades. Can lag
    /// reality; the `.in/append` handler re-checks the snapshot status
    /// before reusing it.
    /// </summary>
    public string? CurrentRunId { get; set; }

    /// <summary>
    /// Monotonic counter used for optimistic locking on `currentRunId`
    /// swaps. Bumped atomically alongside any update that changes
    /// `currentRunId`.
    /// </summary>
    public int CurrentRunVersion { get; set; } = 0;

    /// <summary>
    /// Terminal markers — written once, never flipped back.
    /// </summary>
    public DateTime? ClosedAt { get; set; }

    public string? ClosedReason { get; set; }

    public DateTime? ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// S2 basin where this session's stream pair lives. Stamped at create
    /// time from `Organization.streamBasinName` so reads can resolve the
    /// basin without joining org. Null when the org has no per-org basin
    /// (OSS, or pre-backfill); reads fall back to the global basin.
    /// </summary>
    public string? StreamBasinName { get; set; }

    /// <summary>
    /// Storage URI (with protocol prefix) for this session's chat.agent
    /// snapshot blob. Set on first snapshot write. Null = pre-column session,
    /// fall back to computed default path.
    /// </summary>
    public string? ChatSnapshotStoragePath { get; set; }

    public virtual ICollection<SessionRun> Runs { get; set; } = new List<SessionRun>();
}
