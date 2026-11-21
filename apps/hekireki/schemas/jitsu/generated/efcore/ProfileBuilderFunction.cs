#nullable enable

using System;

namespace Jitsu.Models;

public partial class ProfileBuilderFunction
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

    public string ProfileBuilderId { get; set; } = null!;

    public virtual ProfileBuilder ProfileBuilder { get; set; } = null!;

    public string FunctionId { get; set; } = null!;

    public virtual ConfigurationObject Function { get; set; } = null!;
}
