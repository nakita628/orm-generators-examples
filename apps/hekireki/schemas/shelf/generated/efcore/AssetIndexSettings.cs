#nullable enable

using System;

namespace Shelf.Models;

public partial class AssetIndexSettings
{
    public string Id { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public AssetIndexMode Mode { get; set; } = AssetIndexMode.Simple;

    public string Columns { get; set; } = "[ {\"name\": \"id\", \"visible\": true, \"position\": 0}, {\"name\": \"status\", \"visible\": true, \"position\": 1}, {\"name\": \"description\", \"visible\": true, \"position\": 2}, {\"name\": \"valuation\", \"visible\": true, \"position\": 3}, {\"name\": \"createdAt\", \"visible\": true, \"position\": 4}, {\"name\": \"category\", \"visible\": true, \"position\": 5}, {\"name\": \"tags\", \"visible\": true, \"position\": 6}, {\"name\": \"location\", \"visible\": true, \"position\": 7}, {\"name\": \"kit\", \"visible\": true, \"position\": 8}, {\"name\": \"custody\", \"visible\": true, \"position\": 9} ]";

    public bool FreezeColumn { get; set; } = true;

    public bool ShowAssetImage { get; set; } = true;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
