#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Every SCIM request we could attribute, and what we answered (ADR-126).
///
/// EVIDENCE, NOT TRUTH. A request authors nothing — an HTTP round trip that
/// was refused decided nothing — so this is a table and not an event: nothing
/// downstream derives from it, rebuilding the world from the log must not
/// depend on how often a provider retried a GET, and rows age out of a
/// retention window an event log would never have.
///
/// NO ROW FOR A REQUEST WE CANNOT ATTRIBUTE. A token that does not verify is
/// answered 401 and recorded nowhere: we do not know whose organization to
/// file it under, and a table unauthenticated traffic can write is a table
/// anybody on the internet can fill. What answers that case instead is
/// `ScimToken.lastUsedAt` being null. The 403 IS recorded — a lapsed plan is
/// a credential we recognize.
///
/// THE PAYLOAD RULE HOLDS. A method, a resource name, a status, a reason slug
/// and our own sentence. Never a token, never a header, never a provider's
/// raw message — the same reason `scim_apply_failed` carries a code and not a
/// message.
/// </summary>
public partial class ScimRequestLog
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    /// <summary>
    /// Null when the refusal resolved an organization and stopped before the
    /// connection — a lapsed plan is answered off the token alone.
    /// </summary>
    public string? ConnectionId { get; set; }

    public string Method { get; set; } = null!;

    /// <summary>
    /// The SCIM resource asked for ("Users", "Users/:id", "Groups"), never the
    /// raw path: a path carries ids and query strings, and this is read by
    /// people rather than matched by machines.
    /// </summary>
    public string Resource { get; set; } = null!;

    public int Status { get; set; }

    /// <summary>
    /// A stable slug, only on a refusal, so a reader branches on this and never
    /// on the sentence beside it.
    /// </summary>
    public string? Reason { get; set; }

    /// <summary>
    /// Our own short sentence, customer-safe.
    /// </summary>
    public string? Detail { get; set; }

    public DateTime OccurredAt { get; set; }
}
