#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage.
/// </summary>
public partial class DatasetVersion
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string DatasetId { get; set; } = null!;

    public int Version { get; set; }

    /// <summary>
    /// A name given by hand; a named version is never thinned by retention.
    /// </summary>
    public string? Name { get; set; }

    /// <summary>
    /// "autosave", "manual", "experiment_run", "restore" or "initial".
    /// </summary>
    public string Trigger { get; set; } = null!;

    /// <summary>
    /// SHA-256 over the columns and the rows in order; equal content saves no new version.
    /// </summary>
    public string ContentHash { get; set; } = null!;

    public string ColumnTypes { get; set; } = null!;

    public int RowCount { get; set; }

    public long SizeBytes { get; set; }

    public int ChunkCount { get; set; }

    public int? RestoredFromVersion { get; set; }

    public string? AuthorId { get; set; }

    /// <summary>
    /// "user", "api_key" or "system": who saved the version when authorId is null.
    /// </summary>
    public string? AuthorKind { get; set; }

    /// <summary>
    /// Set when an experiment pins this version; a pinned version is never thinned.
    /// </summary>
    public DateTime? PinnedAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
