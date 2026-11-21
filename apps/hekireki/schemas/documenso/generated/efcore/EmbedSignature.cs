#nullable enable

using System;

namespace Documenso.Models;

/// <summary>
/// `key` is sha256(organisationId:lower(recipientEmail):clientUserId) so neither
/// the recipient email nor the embedder's identifier is stored in plaintext.
/// </summary>
public partial class EmbedSignature
{
    public string Key { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string OrganisationId { get; set; } = null!;

    public virtual Organisation Organisation { get; set; } = null!;

    /// <summary>
    /// Last writer. Not a relation so the row survives document deletion and
    /// remains useful for abuse tracing.
    /// </summary>
    public int RecipientId { get; set; }

    public string Signature { get; set; } = null!;
}
