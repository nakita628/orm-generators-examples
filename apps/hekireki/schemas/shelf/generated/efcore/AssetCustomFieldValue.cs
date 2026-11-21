#nullable enable

using System;

namespace Shelf.Models;

public partial class AssetCustomFieldValue
{
    public string Id { get; set; } = null!;

    public string Value { get; set; } = null!;

    public virtual Asset Asset { get; set; } = null!;

    public string AssetId { get; set; } = null!;

    public virtual CustomField CustomField { get; set; } = null!;

    public string CustomFieldId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
