#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// An automation holding a Slack connection in use (ARCHITECTURE.md §3, the
/// claims ruling): a claimed connection is never deleted. `projectId` is the
/// claimant's project; `organizationId` is the connection's tenancy anchor.
/// No foreign key: slack's service refuses to delete a claimed connection.
/// </summary>
public partial class SlackConnectionClaim
{
    public string ConnectionId { get; set; } = null!;

    public string ClaimantId { get; set; } = null!;

    public string ClaimantLabel { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
