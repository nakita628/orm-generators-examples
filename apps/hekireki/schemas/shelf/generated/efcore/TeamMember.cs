#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class TeamMember
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual ICollection<Custody> Custodies { get; set; } = new List<Custody>();

    public virtual ICollection<Invite> ReceivedInvites { get; set; } = new List<Invite>();

    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    public virtual ICollection<KitCustody> KitCustodies { get; set; } = new List<KitCustody>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? DeletedAt { get; set; }

    public virtual ICollection<Booking> Bookings { get; set; } = new List<Booking>();

    public virtual ICollection<AssetReminder> AssetReminders { get; set; } = new List<AssetReminder>();

    public virtual ICollection<ConsumptionLog> ConsumptionLogs { get; set; } = new List<ConsumptionLog>();

    public virtual ICollection<TeamMemberNote> TeamMemberNotes { get; set; } = new List<TeamMemberNote>();

    public virtual ICollection<BookingSettings> AlwaysNotifyBookingSettings { get; set; } = new List<BookingSettings>();

    public virtual ICollection<Booking> BookingNotifications { get; set; } = new List<Booking>();
}
