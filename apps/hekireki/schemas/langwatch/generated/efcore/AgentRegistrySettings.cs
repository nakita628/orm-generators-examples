#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// An organisation's agent registry settings. No row reads as the defaults.
/// </summary>
public partial class AgentRegistrySettings
{
    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// Days without traffic after which a listed agent reads as dormant.
    /// </summary>
    public int DormantAfterDays { get; set; } = 30;

    public DateTime UpdatedAt { get; set; }
}
