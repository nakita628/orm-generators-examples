#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Whether an organization's permission checks run on the authorization engine: authz's own
/// switch, set when the engine migration proves the organization and pinned off by an operator's
/// rollback. No row is the legacy path. Spec: modules/authz/specs/engine-switch.feature.
/// </summary>
public partial class AuthzEngineSwitch
{
    public string OrganizationId { get; set; } = null!;

    public bool OnEngine { get; set; }

    public bool Pinned { get; set; } = false;

    public DateTime SwitchedAt { get; set; }

    public string? LastAppliedEventId { get; set; }

    public DateTime UpdatedAt { get; set; }
}
