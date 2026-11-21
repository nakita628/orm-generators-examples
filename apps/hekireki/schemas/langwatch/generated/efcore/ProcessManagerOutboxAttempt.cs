#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Why an outbox delivery failed, one row per FAILED attempt
/// (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
/// table grows with trouble, not with traffic. Rows die with their message:
/// the cascade delete is the retention story, riding the existing sweep.
/// </summary>
public partial class ProcessManagerOutboxAttempt
{
    public string Id { get; set; } = null!;

    public string OutboxId { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// 1-based attempt number, as the dispatcher counts it.
    /// </summary>
    public int Attempt { get; set; }

    public DateTime OccurredAt { get; set; }

    public ProcessManagerOutboxAttemptOutcome Outcome { get; set; }

    public string ErrorType { get; set; } = null!;

    /// <summary>
    /// The safe failure diagnostic the dispatcher already computes for its
    /// span — never a raw provider body.
    /// </summary>
    public string ErrorMessage { get; set; } = null!;

    public int? RetryAfterMs { get; set; }

    public virtual ProcessManagerOutbox Outbox { get; set; } = null!;
}
