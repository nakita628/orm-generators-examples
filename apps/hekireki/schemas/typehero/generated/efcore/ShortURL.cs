#nullable enable

using System;

namespace TypeHero.Models;

public partial class ShortURL
{
    public int Id { get; set; }

    /// <summary>
    /// the original full url that we want to shorten e.g. https://www.url.com
    /// </summary>
    public string OriginalUrl { get; set; } = null!;

    /// <summary>
    /// slug for the short url e.g. url.com/[shortUrlSlug]
    /// </summary>
    public string ShortUrlSlug { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// the time the short url will expire
    /// </summary>
    public DateTime? ExpiresAt { get; set; }

    /// <summary>
    /// the number of times the short url has redirected to the original url
    /// </summary>
    public int Clicks { get; set; } = 0;

    public string? UserId { get; set; }

    public virtual User? User { get; set; }
}
