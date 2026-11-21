#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// The join-request pipeline's projection (ADR-117, D12 - see
/// specs/identity/join-requests.feature). A pure event-truth head like
/// `Identifier` and `SsoConnection`: every column is fold-written, rows are
/// never deleted (a terminal state is a tombstone the admin panel stops
/// listing, not an absence), and a replay rebuilds it whole-row from the
/// event log.
///
/// The fold cursor lives on the row rather than in a sibling table, because
/// one row IS one aggregate here - a request.
///
/// No relation to Organization or User on purpose, matching the two heads
/// above: a projection is rebuilt from events, and a foreign key would make a
/// replay depend on relational state the log does not carry.
///
/// Membership is NOT here. An approval dispatches an attach on the grants
/// ledger, exactly as accepting an invitation does, and that ledger fact
/// carries `source: "join-request"` as its provenance - which is what puts a
/// join approval on the customer's audit page beside every other membership
/// change rather than in a history of its own.
///
/// The requester's address is not here either: the DOMAIN is the fact, and it
/// is the only part of an address a join decision may depend on.
/// </summary>
public partial class JoinRequest
{
    public string Id { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// The normalized domain the match was made on.
    /// </summary>
    public string Domain { get; set; } = null!;

    /// <summary>
    /// PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN
    /// </summary>
    public string State { get; set; } = null!;

    /// <summary>
    /// How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`.
    /// </summary>
    public string MatchedVia { get; set; } = null!;

    /// <summary>
    /// Where the request was made: `web` or `cli` (ADR-171 v6). A request from
    /// the terminal lands a Developer seat whatever the joiner seat says.
    /// </summary>
    public string Origin { get; set; } = "web";

    /// <summary>
    /// While PENDING: when the day-14 wake lapses it. Null once resolved.
    /// </summary>
    public DateTime? ExpiresAt { get; set; }

    public DateTime? ResolvedAt { get; set; }

    /// <summary>
    /// user | policy | invite - who ended it. `policy` is domain auto-join's
    /// own principal, deliberately not the one SSO auto-join uses.
    /// </summary>
    public string? ResolvedByType { get; set; }

    /// <summary>
    /// The admin's user id, `domain-auto`, or the invitation id.
    /// </summary>
    public string? ResolvedById { get; set; }

    /// <summary>
    /// user | invite-accepted, on a withdrawal.
    /// </summary>
    public string? WithdrawalCause { get; set; }

    /// <summary>
    /// The SSO connection a single sign-on arrival was admitted through; null otherwise.
    /// </summary>
    public string? ConnectionId { get; set; }

    public DateTime OccurredAt { get; set; }

    public string LastEventId { get; set; } = null!;

    public DateTime AcceptedAt { get; set; }

    public string ProjectionVersion { get; set; } = null!;

    /// <summary>
    /// Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    /// would make a replayed row differ from the row it rebuilds, which is
    /// exactly the parity this projection has to keep.
    /// </summary>
    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
