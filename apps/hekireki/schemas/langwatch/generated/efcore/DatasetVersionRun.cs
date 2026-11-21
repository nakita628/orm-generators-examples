#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// An experiment run that read one dataset version; a version a run read is never thinned.
/// </summary>
public partial class DatasetVersionRun
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string DatasetId { get; set; } = null!;

    public string VersionId { get; set; } = null!;

    public string RunId { get; set; } = null!;

    public string? ExperimentId { get; set; }

    public DateTime CreatedAt { get; set; }
}
