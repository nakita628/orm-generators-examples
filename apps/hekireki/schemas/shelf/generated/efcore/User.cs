#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class User
{
    public string Id { get; set; } = null!;

    public string Email { get; set; } = null!;

    public string Username { get; set; } = null!;

    public string? FirstName { get; set; }

    public string? LastName { get; set; }

    public string? DisplayName { get; set; }

    public string? ProfilePicture { get; set; }

    public bool UsedFreeTrial { get; set; } = false;

    public bool Onboarded { get; set; } = false;

    public string? CustomerId { get; set; }

    public bool Sso { get; set; } = false;

    public bool CreatedWithInvite { get; set; } = false;

    public bool SkipSubscriptionCheck { get; set; } = false;

    public bool HasUnpaidInvoice { get; set; } = false;

    public bool WarnForNoPaymentMethod { get; set; } = false;

    public virtual Organization? LastSelectedOrganization { get; set; }

    public string? LastSelectedOrganizationId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? DeletedAt { get; set; }

    public DateTime? LastMobileActiveAt { get; set; }

    public DateFormatPreference? DateFormat { get; set; }

    public TimeFormatPreference? TimeFormat { get; set; }

    public WeekStartPreference? WeekStart { get; set; }

    public string? TimeZone { get; set; }

    public virtual ICollection<Asset> Assets { get; set; } = new List<Asset>();

    public virtual ICollection<Category> Categories { get; set; } = new List<Category>();

    public virtual ICollection<Note> Notes { get; set; } = new List<Note>();

    public virtual ICollection<BookingNote> BookingNotes { get; set; } = new List<BookingNote>();

    public virtual ICollection<AuditNote> AuditNotes { get; set; } = new List<AuditNote>();

    public virtual ICollection<LocationNote> LocationNotes { get; set; } = new List<LocationNote>();

    public virtual ICollection<Qr> QrCodes { get; set; } = new List<Qr>();

    public virtual ICollection<Scan> Scans { get; set; } = new List<Scan>();

    public virtual ICollection<Tag> Tags { get; set; } = new List<Tag>();

    public virtual ICollection<Role> Roles { get; set; } = new List<Role>();

    public virtual ICollection<Location> Locations { get; set; } = new List<Location>();

    public virtual ICollection<Image> Images { get; set; } = new List<Image>();

    public virtual ICollection<Organization> Organizations { get; set; } = new List<Organization>();

    public virtual ICollection<CustomField> CustomFields { get; set; } = new List<CustomField>();

    public virtual ICollection<Invite> SentInvites { get; set; } = new List<Invite>();

    public virtual ICollection<Invite> ReceivedInvites { get; set; } = new List<Invite>();

    public virtual ICollection<TeamMember> TeamMembers { get; set; } = new List<TeamMember>();

    public virtual ICollection<UserOrganization> UserOrganizations { get; set; } = new List<UserOrganization>();

    public virtual ICollection<Booking> Bookings { get; set; } = new List<Booking>();

    public virtual ICollection<Booking> Custodies { get; set; } = new List<Booking>();

    public virtual ICollection<Booking> ApprovedBookings { get; set; } = new List<Booking>();

    public virtual ICollection<Kit> CreatedKits { get; set; } = new List<Kit>();

    public TierId TierId { get; set; } = TierId.Free;

    public virtual Tier Tier { get; set; } = null!;

    public virtual ICollection<AssetReminder> AssetReminders { get; set; } = new List<AssetReminder>();

    public virtual UserContact? Contact { get; set; }

    public virtual UserBusinessIntel? BusinessIntel { get; set; }

    public virtual ICollection<AuditSession> AuditsCreated { get; set; } = new List<AuditSession>();

    public virtual ICollection<AssetImport> AssetImports { get; set; } = new List<AssetImport>();

    public virtual ICollection<AuditAssignment> AuditAssignments { get; set; } = new List<AuditAssignment>();

    public virtual ICollection<AuditAsset> AuditAssetsScanned { get; set; } = new List<AuditAsset>();

    public virtual ICollection<AuditScan> AuditScans { get; set; } = new List<AuditScan>();

    public virtual ICollection<AuditImage> AuditImagesUploaded { get; set; } = new List<AuditImage>();

    public string? ReferralSource { get; set; }

    public virtual ICollection<AssetIndexSettings> AssetIndexSettings { get; set; } = new List<AssetIndexSettings>();

    public virtual ICollection<PartialBookingCheckin> PartialCheckins { get; set; } = new List<PartialBookingCheckin>();

    public virtual ICollection<PartialBookingCheckout> PartialCheckouts { get; set; } = new List<PartialBookingCheckout>();

    public virtual ICollection<MobileAuthCode> MobileAuthCodes { get; set; } = new List<MobileAuthCode>();

    public virtual ICollection<PushToken> PushTokens { get; set; } = new List<PushToken>();

    public virtual NotificationPreference? NotificationPreference { get; set; }

    public virtual ICollection<Update> CreatedUpdates { get; set; } = new List<Update>();

    public virtual ICollection<UserUpdateRead> UpdateReads { get; set; } = new List<UserUpdateRead>();

    public virtual CustomTierLimit? CustomTierLimit { get; set; }

    public virtual ICollection<AssetFilterPreset> AssetFilterPresets { get; set; } = new List<AssetFilterPreset>();

    public virtual ICollection<AssetModel> AssetModels { get; set; } = new List<AssetModel>();

    public virtual ICollection<ConsumptionLog> ConsumptionLogs { get; set; } = new List<ConsumptionLog>();

    public virtual ICollection<UserScimExternalId> ScimExternalIds { get; set; } = new List<UserScimExternalId>();

    public virtual ICollection<RoleChangeLog> RoleChangesReceived { get; set; } = new List<RoleChangeLog>();

    public virtual ICollection<RoleChangeLog> RoleChangesPerformed { get; set; } = new List<RoleChangeLog>();

    public virtual ICollection<TeamMemberNote> TeamMemberNotesAuthored { get; set; } = new List<TeamMemberNote>();
}
