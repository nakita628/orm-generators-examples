#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Operational round-robin position for ADR-123's domain re-proof sweep.
///
/// This is deliberately not a column on `SsoConnection`: that row is a
/// whole-row event projection and its fold is its only writer, while a
/// healthy re-read states no fact. Keeping the scheduling cursor beside the
/// projection preserves both rules and still lets a fixed-size sweep revisit
/// every connection, including one whose proof has started wavering.
/// </summary>
public partial class SsoConnectionReproofCursor
{
    public string ConnectionId { get; set; } = null!;

    public DateTime LastReproofAt { get; set; }

    public virtual SsoConnection Connection { get; set; } = null!;
}
