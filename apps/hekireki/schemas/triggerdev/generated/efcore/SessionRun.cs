#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// Historical record of every run a Session has owned. Append-only —
/// rows are inserted on each `ensureRunForSession` claim, never updated.
/// Lets us reconstruct the run timeline of a chat for debugging /
/// dashboard surfaces. The relation cascades on Session delete (tied to
/// the session lifecycle) but `runId` is a plain string column with no
/// FK to TaskRun so run pruning is independent.
/// </summary>
public partial class SessionRun
{
    public string Id { get; set; } = null!;

    public string SessionId { get; set; } = null!;

    /// <summary>
    /// TaskRun.id (no FK — runs may be archived independently of session history)
    /// </summary>
    public string RunId { get; set; } = null!;

    /// <summary>
    /// One of: "initial" | "continuation" | "upgrade" | "manual".
    /// Plain string for forward-compat with future trigger reasons.
    /// </summary>
    public string Reason { get; set; } = null!;

    public DateTime TriggeredAt { get; set; }

    public virtual Session Session { get; set; } = null!;
}
