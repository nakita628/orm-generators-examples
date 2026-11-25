#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Each account nurturing erased, as a SHA-256 of its id and never the id: a gateway key fact
/// naming it, however late it lands, stores no creator or holder. Platform-wide, like the erasure.
/// </summary>
public partial class NurturingErasedUser
{
    public string UserIdHash { get; set; } = null!;

    public DateTime ErasedAt { get; set; }
}
