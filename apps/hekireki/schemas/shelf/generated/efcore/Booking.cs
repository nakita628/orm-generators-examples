#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Booking
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public BookingStatus Status { get; set; } = BookingStatus.Draft;

    public string? Description { get; set; } = "";

    public string? ActiveSchedulerReference { get; set; }

    public virtual User Creator { get; set; } = null!;

    public string CreatorId { get; set; } = null!;

    public virtual User? CustodianUser { get; set; }

    public string? CustodianUserId { get; set; }

    public virtual TeamMember? CustodianTeamMember { get; set; }

    public string? CustodianTeamMemberId { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual ICollection<BookingAsset> BookingAssets { get; set; } = new List<BookingAsset>();

    public virtual ICollection<ConsumptionLog> ConsumptionLogs { get; set; } = new List<ConsumptionLog>();

    public virtual ICollection<BookingNote> Notes { get; set; } = new List<BookingNote>();

    public virtual ICollection<BookingModelRequest> ModelRequests { get; set; } = new List<BookingModelRequest>();

    public virtual ICollection<Tag> Tags { get; set; } = new List<Tag>();

    public virtual ICollection<PartialBookingCheckin> PartialCheckins { get; set; } = new List<PartialBookingCheckin>();

    public virtual ICollection<PartialBookingCheckout> PartialCheckouts { get; set; } = new List<PartialBookingCheckout>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime From { get; set; }

    public DateTime To { get; set; }

    public DateTime? OriginalFrom { get; set; }

    public DateTime? OriginalTo { get; set; }

    public DateTime? AutoArchivedAt { get; set; }

    public bool ArchivedWithoutCheckin { get; set; } = false;

    public string? CancellationReason { get; set; }

    public DateTime? ApprovedAt { get; set; }

    public virtual User? ApprovedBy { get; set; }

    public string? ApprovedById { get; set; }

    public virtual ICollection<TeamMember> NotificationRecipients { get; set; } = new List<TeamMember>();
}
