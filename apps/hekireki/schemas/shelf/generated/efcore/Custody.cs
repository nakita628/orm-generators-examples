#nullable enable

using System;

namespace Shelf.Models;

public partial class Custody
{
    public string Id { get; set; } = null!;

    public virtual TeamMember Custodian { get; set; } = null!;

    public string TeamMemberId { get; set; } = null!;

    public virtual Asset Asset { get; set; } = null!;

    public string AssetId { get; set; } = null!;

    public virtual KitCustody? KitCustody { get; set; }

    public string? KitCustodyId { get; set; }

    public int Quantity { get; set; } = 1;

    public virtual Location? Location { get; set; }

    public string? LocationId { get; set; }

    public bool SourceUnknown { get; set; } = false;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
