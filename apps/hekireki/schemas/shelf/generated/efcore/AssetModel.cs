#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

/// <summary>
/// *
/// * A template/grouping entity for assets.
/// * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
/// * and provides default field values when creating new assets from the model.
/// </summary>
public partial class AssetModel
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    /// <summary>
    /// Public URL (bucket `files`) of the model's cover image. Rendered by any
    /// asset of this model that has no image of its own. Public — and therefore
    /// never-expiring and cacheable — like Location.imageUrl; see
    /// resolveAssetImage() for the cascade.
    /// </summary>
    public string? Image { get; set; }

    /// <summary>
    /// Public URL of the 108px thumbnail generated alongside `image`. Used at
    /// list sizes so inheriting assets never render the full-size original.
    /// </summary>
    public string? ThumbnailImage { get; set; }

    public string? DefaultCategoryId { get; set; }

    public virtual Category? DefaultCategory { get; set; }

    public double? DefaultValuation { get; set; }

    /// <summary>
    /// Default custom field values. Like the category and valuation defaults,
    /// each one fills the matching field only when it is empty, at the moment
    /// an asset joins this model. Nothing stays linked afterwards.
    /// </summary>
    public virtual ICollection<AssetModelCustomFieldDefault> CustomFieldDefaults { get; set; } = new List<AssetModelCustomFieldDefault>();

    public virtual ICollection<Asset> Assets { get; set; } = new List<Asset>();

    public virtual ICollection<BookingModelRequest> ModelRequests { get; set; } = new List<BookingModelRequest>();

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual User CreatedBy { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
