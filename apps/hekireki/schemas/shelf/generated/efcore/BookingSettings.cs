#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class BookingSettings
{
    public string Id { get; set; } = null!;

    public int BufferStartTime { get; set; } = 0;

    public bool TagsRequired { get; set; } = false;

    public int? MaxBookingLength { get; set; }

    public bool MaxBookingLengthSkipClosedDays { get; set; } = false;

    public bool AutoArchiveBookings { get; set; } = false;

    public int AutoArchiveDays { get; set; } = 2;

    public bool AutoArchiveExpiredReservations { get; set; } = false;

    public bool RequireExplicitCheckinForAdmin { get; set; } = false;

    public bool RequireExplicitCheckinForSelfService { get; set; } = false;

    public bool RequireExplicitCheckoutForAdmin { get; set; } = false;

    public bool RequireExplicitCheckoutForSelfService { get; set; } = false;

    public bool CountKitsAsSingleUnit { get; set; } = false;

    public bool NotifyBookingCreator { get; set; } = true;

    public bool NotifyAdminsOnNewBooking { get; set; } = true;

    public virtual ICollection<TeamMember> AlwaysNotifyTeamMembers { get; set; } = new List<TeamMember>();

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
