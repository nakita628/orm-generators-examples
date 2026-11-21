#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The dated match table: "`user_abc123` on Anthropic is platform user #42,
/// valid from March 3", closed when it stops being true.
///
/// Dating is what makes a re-issued email address survivable: the leaver's link
/// closes at offboarding, the new hire gets a new link, and January's spend
/// stays with January's person. A correction closes the old row and opens a new
/// one; nothing is ever silently rewritten.
/// </summary>
public partial class IdentityMatch
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// The provider-side identity this link is about.
    /// </summary>
    public string DiscoveredPersonId { get; set; } = null!;

    /// <summary>
    /// The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
    /// the row and its dates remain, so the history of the link is not rewritten
    /// by the erasure either.
    /// </summary>
    public string? UserId { get; set; }

    /// <summary>
    /// directory_id | verified_email | human_confirmed. Exact verified-email
    /// equality is the primary evidence; a directory id corroborates but never
    /// stands alone (ADR-128 §12).
    /// </summary>
    public string EvidenceKind { get; set; } = null!;

    public DateTime ValidFrom { get; set; }

    /// <summary>
    /// Null means the link is open. Offboarding and corrections close it.
    /// </summary>
    public DateTime? ValidTo { get; set; }

    public DateTime CreatedAt { get; set; }
}
