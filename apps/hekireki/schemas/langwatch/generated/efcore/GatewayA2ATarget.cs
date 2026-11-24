#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

public partial class GatewayA2ATarget
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string? ProjectId { get; set; }

    public string Name { get; set; } = null!;

    public string AgentName { get; set; } = null!;

    public string CardUrl { get; set; } = null!;

    public string UpstreamUrl { get; set; } = null!;

    public string? EncryptedUpstreamHeaders { get; set; }

    public bool RestrictToAllowList { get; set; } = false;

    public List<string>? AllowedVirtualKeyIds { get; set; } = new List<string>();

    public bool Blocked { get; set; } = false;

    public string? BlockReason { get; set; }

    public string CreatedById { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? ArchivedAt { get; set; }
}
