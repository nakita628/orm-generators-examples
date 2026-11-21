#nullable enable

using System;

namespace Teable.Models;

/// <summary>
/// Task→field refset for activity refcounting (prevents merge double-count).
/// </summary>
public partial class ComputedTaskFieldRef
{
    public string TaskId { get; set; } = null!;

    public string FieldId { get; set; } = null!;

    public string TableId { get; set; } = null!;

    public string BaseId { get; set; } = null!;

    public bool WasProcessing { get; set; } = false;

    public DateTime CreatedAt { get; set; }
}
