#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
/// report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
/// tenant: every install receives every published notice that has not expired.
/// </summary>
public partial class SelfHostedNotice
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// `info`, `warning` or `security`.
    /// </summary>
    public string Severity { get; set; } = null!;

    public string Title { get; set; } = null!;

    public string Body { get; set; } = null!;

    /// <summary>
    /// An https link to read more; null where the notice names none.
    /// </summary>
    public string? Url { get; set; }

    /// <summary>
    /// The release that fixes what the notice is about; an install on it or newer hides it.
    /// </summary>
    public string? FixedIn { get; set; }

    /// <summary>
    /// `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent.
    /// </summary>
    public string Status { get; set; } = null!;

    public DateTime? ExpiresAt { get; set; }

    public DateTime? PublishedAt { get; set; }

    /// <summary>
    /// The staff member who published it, by user id.
    /// </summary>
    public string? PublishedBy { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
