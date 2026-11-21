#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// When an organization was created, from organization's created event: nurturing's own copy for
/// the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
/// NurturingOrganization row.
/// </summary>
public partial class NurturingOrganizationCreated
{
    public string OrganizationId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }
}
