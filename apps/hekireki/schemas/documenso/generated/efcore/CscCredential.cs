#nullable enable

using System;

namespace Documenso.Models;

/// <summary>
/// Per-recipient cached credential metadata for the configured Cloud Signature
/// Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
/// symmetric-encrypted service-scope access token. One row per recipient on a
/// CSC envelope; absent for SES recipients.
/// </summary>
public partial class CscCredential
{
    public string Id { get; set; } = null!;

    public string ProviderId { get; set; } = null!;

    public string CredentialId { get; set; } = null!;

    /// <summary>
    /// TSP-validated certificate chain (length-prefixed). Nullable until the
    /// service-scope OAuth callback successfully retrieves and validates it.
    /// </summary>
    public byte[]? CertCache { get; set; }

    /// <summary>
    /// Algorithm metadata derived from the CSC `credentials/info` response.
    /// Held as flat scalars so each is independently queryable.
    /// </summary>
    public string SignatureAlgorithm { get; set; } = null!;

    public string KeyType { get; set; } = null!;

    public string DigestAlgorithm { get; set; } = null!;

    public int? KeyLenBits { get; set; }

    public string? SignAlgoParams { get; set; }

    /// <summary>
    /// Service-scope access token, symmetric-encrypted at the application layer
    /// before persistence. Nullable until the service-scope OAuth callback runs.
    /// </summary>
    public byte[]? ServiceTokenCiphertext { get; set; }

    public DateTime? ServiceTokenExpiresAt { get; set; }

    public int RecipientId { get; set; }

    public virtual Recipient Recipient { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
