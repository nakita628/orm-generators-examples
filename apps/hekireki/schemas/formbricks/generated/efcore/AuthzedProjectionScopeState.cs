#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// ENG-3282: per-scope readiness of a SpiceDB projection. A row with `readyAt` set means the scope was
/// backfilled and audited clean, so the evaluator may enforce it. Written only by the backfill CLI.
/// </summary>
public partial class AuthzedProjectionScopeState
{
    public string Scope { get; set; } = null!;

    public DateTime? ReadyAt { get; set; }

    public string? ReadyBy { get; set; }

    public DateTime UpdatedAt { get; set; }
}
