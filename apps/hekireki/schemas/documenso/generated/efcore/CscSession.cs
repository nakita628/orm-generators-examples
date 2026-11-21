#nullable enable

using System;

namespace Documenso.Models;

/// <summary>
/// Per-recipient transient sign-time state held across the credential-scope
/// OAuth round-trip. The `recipientId @unique` constraint caps each recipient
/// to at most one in-flight session — re-clicking Sign UPSERTs the row.
/// Cleaned up transitively when the Recipient is deleted.
/// @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
/// </summary>
public partial class CscSession
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
    /// flows through the Recipient cascade.
    /// </summary>
    public string EnvelopeId { get; set; } = null!;

    /// <summary>
    /// Pinned at prep time; stable across overlay re-renders so the captured
    /// signedAttrs digest matches the bytes the TSP signs.
    /// </summary>
    public DateTime SigningTime { get; set; }

    /// <summary>
    /// Contract between prep and sign: an array of
    /// `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
    /// `documentDataId` pins the immutable bytes that produced the captured
    /// digest so sign-time can re-derive the same hash against the same content.
    /// [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
    /// </summary>
    public string ItemsJson { get; set; } = null!;

    /// <summary>
    /// Populated by the credential-scope OAuth callback. Both nullable at row
    /// creation; set together when the SAD is exchanged.
    /// </summary>
    public byte[]? EncryptedSad { get; set; }

    public DateTime? SadExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public int RecipientId { get; set; }

    public virtual Recipient Recipient { get; set; } = null!;
}
