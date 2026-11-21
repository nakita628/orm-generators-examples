#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One dated person → department link (ADR-128 §13, #7882).
///
/// `OrganizationUser.departmentId` stays the CURRENT assignment every screen
/// reads; these rows are its history, appended by the same `assignUser` write
/// so the two cannot drift. They exist for the cost reads: January's spend
/// stays with January's department, so a read resolves against the link that
/// was open on the cost's day, not against today's pointer. Members with no
/// link on a day resolve to "unassigned", never to an error.
///
/// Same shape as IdentityMatch's dated links, same one-open-link guard: the
/// partial unique index in the migration rejects a second open row per
/// (organizationId, userId) with SQLSTATE 23505.
/// </summary>
public partial class DepartmentMembershipHistory
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public string DepartmentId { get; set; } = null!;

    public DateTime ValidFrom { get; set; }

    /// <summary>
    /// Null means the link is open. Reassignment and clearing close it.
    /// </summary>
    public DateTime? ValidTo { get; set; }

    public DateTime CreatedAt { get; set; }
}
