#nullable enable

using System;

namespace Jitsu.Models;

public partial class WorkspaceOptions
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime UpdatedAt { get; set; }

    public string WorkspaceId { get; set; } = null!;

    public virtual Workspace Workspace { get; set; } = null!;

    public string Namespace { get; set; } = null!;

    public string Value { get; set; } = null!;
}
