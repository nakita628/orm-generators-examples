#nullable enable

using System;

namespace Jitsu.Models;

public partial class ConnectorPackage
{
    public string Id { get; set; } = null!;

    public string PackageId { get; set; } = null!;

    public string PackageType { get; set; } = "airbyte";

    public string? Meta { get; set; }

    public byte[]? LogoSvg { get; set; }

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime UpdatedAt { get; set; }
}
