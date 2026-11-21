#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// SCIM's own copy of an organization's SSO connections: a peer fold over
/// identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
/// reducer, one row per connection. SCIM reads this instead of asking identity.
/// Rebuilt from identity's event log by a projection replay, so no relation:
/// a foreign key would make a replay depend on state the log does not carry.
/// </summary>
public partial class ScimSsoConnectionView
{
    /// <summary>
    /// The aggregate: identity's connection id.
    /// </summary>
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt).
    /// </summary>
    public string Folded { get; set; } = null!;

    /// <summary>
    /// The event ids of the last delivery folded in, so a redelivery folds once.
    /// </summary>
    public List<string>? AppliedEventIds { get; set; } = new List<string>();

    /// <summary>
    /// A row of another version reads as undecodable and is re-folded from the log.
    /// </summary>
    public string ProjectionVersion { get; set; } = null!;

    /// <summary>
    /// Event-derived, so a replayed row equals the row it rebuilds.
    /// </summary>
    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
