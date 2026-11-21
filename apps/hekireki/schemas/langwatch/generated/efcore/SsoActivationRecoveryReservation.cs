#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A durable claim on one live recovery path while activation is between its
/// guard and projected ACTIVE fact. It has no expiry: only projection of the
/// matching activation, or an explicit cancellation, may consume it.
/// </summary>
public partial class SsoActivationRecoveryReservation
{
    public string CommandId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string ConnectionId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }
}
