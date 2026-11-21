#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class WorkerInstanceGroup
{
    public string Id { get; set; } = null!;

    public WorkerInstanceGroupType Type { get; set; }

    /// <summary>
    /// For example "us-east-1"
    /// </summary>
    public string Name { get; set; } = null!;

    /// <summary>
    /// If managed, it will default to the name, e.g. "us-east-1"
    /// If unmanged, it will be prefixed with the project ID e.g. "project_1-us-east-1"
    /// </summary>
    public string MasterQueue { get; set; } = null!;

    /// <summary>
    /// "N. Virginia, USA", "Frankfurt, Germany", etc. Used for display purposes
    /// </summary>
    public string? Description { get; set; }

    public bool Hidden { get; set; } = false;

    public virtual WorkerGroupToken Token { get; set; } = null!;

    public string TokenId { get; set; } = null!;

    public virtual ICollection<WorkerInstance> Workers { get; set; } = new List<WorkerInstance>();

    public virtual ICollection<BackgroundWorker> BackgroundWorkers { get; set; } = new List<BackgroundWorker>();

    public virtual ICollection<Project> DefaultForProjects { get; set; } = new List<Project>();

    public virtual Organization? Organization { get; set; }

    public string? OrganizationId { get; set; }

    public virtual Project? Project { get; set; }

    public string? ProjectId { get; set; }

    public string? CloudProvider { get; set; }

    /// <summary>
    /// "usa", "europe", etc. Used like a pseudo enum for things like flags
    /// </summary>
    public string? Location { get; set; }

    public string? StaticIPs { get; set; }

    public WorkloadType WorkloadType { get; set; } = WorkloadType.Container;

    /// <summary>
    /// Geo region; container + MICROVM groups for one geo share it. Set-once once it has runs.
    /// </summary>
    public string? Region { get; set; }

    /// <summary>
    /// When true, runs enqueued to this worker queue may skip the intermediate queue
    /// and be pushed directly to the worker queue when concurrency is available.
    /// </summary>
    public bool EnableFastPath { get; set; } = false;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
