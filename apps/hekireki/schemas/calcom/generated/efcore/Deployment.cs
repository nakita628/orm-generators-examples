#nullable enable

using System;

namespace Calcom.Models;

public partial class Deployment
{
    /// <summary>
    /// This is a single row table, so we use a fixed id
    /// </summary>
    public int Id { get; set; } = 1;

    public string? Logo { get; set; }

    /// <summary>
    /// @zod.import(["import { DeploymentTheme } from '../../zod-utils'"]).custom.use(DeploymentTheme)
    /// </summary>
    public string? Theme { get; set; }

    public string? LicenseKey { get; set; }

    public string? SignatureTokenEncrypted { get; set; }

    public DateTime? AgreedLicenseAt { get; set; }
}
