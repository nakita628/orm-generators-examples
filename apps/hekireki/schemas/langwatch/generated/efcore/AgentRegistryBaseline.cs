#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// An organisation's registry baseline: its first reconcile, written as already pushed so the
/// agents it already had are not announced. Every reconcile after it pushes its changes.
/// </summary>
public partial class AgentRegistryBaseline
{
    public string OrganizationId { get; set; } = null!;

    public DateTime BaselinedAt { get; set; }
}
