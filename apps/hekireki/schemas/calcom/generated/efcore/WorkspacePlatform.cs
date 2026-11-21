#nullable enable

using System;
using System.Collections.Generic;

namespace Calcom.Models;

public partial class WorkspacePlatform
{
    public int Id { get; set; }

    /// <summary>
    /// @zod.string.min(1)
    /// </summary>
    public string Slug { get; set; } = null!;

    /// <summary>
    /// @zod.string.min(1)
    /// </summary>
    public string Name { get; set; } = null!;

    public string Description { get; set; } = null!;

    /// <summary>
    /// @zod.import(["import { serviceAccountKeySchema } from '../../zod-utils'"]).custom.use(serviceAccountKeySchema)
    /// </summary>
    public string DefaultServiceAccountKey { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public bool Enabled { get; set; } = false;

    public virtual ICollection<DelegationCredential> DelegationCredentials { get; set; } = new List<DelegationCredential>();
}
