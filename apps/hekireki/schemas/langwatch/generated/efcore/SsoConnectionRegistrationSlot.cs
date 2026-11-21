#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The one-statement pre-event registration lock. `SsoConnection` remains
/// event truth; this row only decides which command may append that truth.
/// A slot whose projected connection is terminal may be atomically replaced.
/// </summary>
public partial class SsoConnectionRegistrationSlot
{
    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// legacy | direct
    /// </summary>
    public string Kind { get; set; } = null!;

    public string ConnectionId { get; set; } = null!;

    public string? ReplacesConnectionId { get; set; }

    public string CommandId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
