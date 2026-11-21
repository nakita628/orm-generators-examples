#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// A survey's response attachments built into one ZIP in object storage by a background job and
/// downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
/// Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
/// survey deletes it with the survey's storage prefix.
///
/// A partial unique index in the migration allows one queued or running export per requester and
/// survey, so a double click returns the export already in flight.
/// </summary>
public partial class AttachmentExport
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Survey Survey { get; set; } = null!;

    public string SurveyId { get; set; } = null!;

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual User? RequestedBy { get; set; }

    public string? RequestedById { get; set; }

    public AttachmentExportStatus Status { get; set; } = AttachmentExportStatus.Queued;

    /// <summary>
    /// The response filter at request time, a `TResponseFilterCriteria`; null exports every response.
    /// </summary>
    public string? FilterCriteria { get; set; }

    /// <summary>
    /// IANA zone the archive's folder names are stamped in: the org's display zone at request time.
    /// </summary>
    public string TimeZone { get; set; } = null!;

    /// <summary>
    /// Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
    /// older pod dropped as unknown during a rolling upgrade.
    /// </summary>
    public int DispatchCount { get; set; } = 0;

    /// <summary>
    /// The pre-flight's counts, shown while the archive builds.
    /// </summary>
    public int FileCount { get; set; }

    public int ResponseCount { get; set; }

    /// <summary>
    /// What the archive actually holds, once built.
    /// </summary>
    public int? IncludedFileCount { get; set; }

    public long? TotalBytes { get; set; }

    public bool Truncated { get; set; } = false;

    public string? StorageKey { get; set; }

    /// <summary>
    /// A short code (`stale`, `storage_upload_failed`, …), never a message or file name.
    /// </summary>
    public string? ErrorCode { get; set; }

    public DateTime? StartedAt { get; set; }

    public DateTime? FinishedAt { get; set; }

    public DateTime? ExpiresAt { get; set; }
}
