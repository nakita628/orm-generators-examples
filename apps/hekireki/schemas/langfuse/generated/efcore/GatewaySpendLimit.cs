#nullable enable

using System;
using System.Collections.Generic;

namespace Langfuse.Models;

public partial class GatewaySpendLimit
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public GatewaySpendLimitScope Scope { get; set; }

    public string? ApiKeyId { get; set; }

    public virtual ApiKey? ApiKey { get; set; }

    public string? UserId { get; set; }

    public virtual User? User { get; set; }

    public GatewaySpendLimitInterval Interval { get; set; }

    public decimal LimitUsd { get; set; }

    public List<int>? AlertThresholds { get; set; } = new List<int> { 100 };

    public DateTime CountingStartedAt { get; set; }

    public string? OwnerUserId { get; set; }

    public virtual User? OwnerUser { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<GatewaySpendLimitAlert> Alerts { get; set; } = new List<GatewaySpendLimitAlert>();
}
