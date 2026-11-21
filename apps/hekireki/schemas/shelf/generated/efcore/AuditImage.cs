#nullable enable

using System;

namespace Shelf.Models;

public partial class AuditImage
{
    public string Id { get; set; } = null!;

    public string ImageUrl { get; set; } = null!;

    public string? ThumbnailUrl { get; set; }

    public string? Description { get; set; }

    public virtual AuditSession AuditSession { get; set; } = null!;

    public string AuditSessionId { get; set; } = null!;

    public virtual AuditAsset? AuditAsset { get; set; }

    public string? AuditAssetId { get; set; }

    public virtual User? UploadedBy { get; set; }

    public string? UploadedById { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// When the photo was taken, from its EXIF data (or the phone app's copy of
    /// it). Null when the photo carried no usable date. EXIF can be edited, so
    /// this is the best signal available, not proof.
    /// </summary>
    public DateTime? CapturedAt { get; set; }

    /// <summary>
    /// How the photo reached Shelf. Null only for photos uploaded before this
    /// field existed; a phone upload with no or an unknown source is PHONE.
    /// </summary>
    public AuditImageCaptureSource? CaptureSource { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
