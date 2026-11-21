#nullable enable

using System;

namespace TypeHero.Models;

public partial class ImageUpload
{
    /// <summary>
    /// the id of the db row
    /// </summary>
    public string Id { get; set; } = null!;

    /// <summary>
    /// the user id fk
    /// </summary>
    public string UserId { get; set; } = null!;

    /// <summary>
    /// The user that this is related to
    /// </summary>
    public virtual User User { get; set; } = null!;

    /// <summary>
    /// the url to the s3 uploaded assets
    /// </summary>
    public string Url { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
