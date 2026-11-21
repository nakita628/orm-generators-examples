#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// A custom field value an AssetModel gives the assets that join it.
///
/// `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
/// asset is a plain copy. Rows cascade away with their model or field.
/// </summary>
public partial class AssetModelCustomFieldDefault
{
    public string Id { get; set; } = null!;

    public string Value { get; set; } = null!;

    public virtual AssetModel AssetModel { get; set; } = null!;

    public string AssetModelId { get; set; } = null!;

    public string CustomFieldId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
