#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class UserOrganization
{
    public string Id { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public List<OrganizationRoles>? Roles { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// Secret, unguessable token for this member's iCal calendar-feed subscription
    /// in this workspace. Null until the user generates a feed; regenerating it
    /// rotates the value (invalidating any previously-shared subscription URL).
    /// `@unique` doubles as the lookup index for the public feed route.
    /// </summary>
    public string? CalendarTokenId { get; set; }
}
