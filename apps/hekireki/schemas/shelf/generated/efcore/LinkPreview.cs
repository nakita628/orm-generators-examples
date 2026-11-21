#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// What Shelf knows about one web link used in a workspace's custom fields,
/// keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
/// the link preview job (fetch columns) and people (customTitle).
/// </summary>
public partial class LinkPreview
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// `LinkKey` from @shelf/links, never a raw field value. At most 2048
    /// characters (CHECK constraint in the migration).
    /// </summary>
    public string Url { get; set; } = null!;

    public string? PageTitle { get; set; }

    /// <summary>
    /// Null until the first fetch finishes.
    /// </summary>
    public LinkPreviewOutcome? Outcome { get; set; }

    public int Attempts { get; set; } = 0;

    public DateTime? FetchedAt { get; set; }

    /// <summary>
    /// When the link may be fetched next; null means never again. A claim moves
    /// it a few minutes ahead, which is also the lease on a fetch in flight.
    /// </summary>
    public DateTime? NextFetchAt { get; set; } = DateTime.UtcNow;

    /// <summary>
    /// The name a person gave the link in this workspace; wins over pageTitle.
    /// </summary>
    public string? CustomTitle { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;
}
