#nullable enable

using System;

namespace Jitsu.Models;

public partial class KvStoreEntry
{
    public string Key { get; set; } = null!;

    public string Value { get; set; } = null!;

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime? ExpiresAt { get; set; }
}
