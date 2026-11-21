#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
/// storage files that can't be removed inside the database transaction. A row is written in the same
/// transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
/// queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
/// </summary>
public partial class DeletionCleanup
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DeletionCleanupKind Kind { get; set; }

    /// <summary>
    /// Scoping and logs. Not foreign keys: the rows they name are usually gone already.
    /// </summary>
    public string OrganizationId { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    /// <summary>
    /// Hub kinds: the organisation's feedback directories, read inside the delete's transaction.
    /// </summary>
    public List<string>? TenantIds { get; set; } = new List<string>();

    /// <summary>
    /// `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response.
    /// </summary>
    public List<string>? ResponseIds { get; set; } = new List<string>();

    /// <summary>
    /// `storageFiles`: the object keys or URLs to delete.
    /// </summary>
    public List<string>? FileKeys { get; set; } = new List<string>();

    public int Attempts { get; set; } = 0;

    /// <summary>
    /// Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup.
    /// </summary>
    public DateTime NextAttemptAt { get; set; }

    /// <summary>
    /// The last failure: a short code and message, never record or file content.
    /// </summary>
    public string? LastError { get; set; }
}
