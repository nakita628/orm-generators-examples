#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

public partial class Organization
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = "Personal";

    public OrganizationType Type { get; set; } = OrganizationType.Personal;

    public virtual User Owner { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public Currency Currency { get; set; } = Currency.Usd;

    public virtual ICollection<TeamMember> Members { get; set; } = new List<TeamMember>();

    public virtual ICollection<Asset> Assets { get; set; } = new List<Asset>();

    public virtual ICollection<AssetKit> AssetKits { get; set; } = new List<AssetKit>();

    public virtual ICollection<AssetLocation> AssetLocations { get; set; } = new List<AssetLocation>();

    public virtual ICollection<Scan> Scans { get; set; } = new List<Scan>();

    public virtual Image? Image { get; set; }

    public string? ImageId { get; set; }

    public bool EnabledSso { get; set; } = false;

    public string? SsoDetailsId { get; set; }

    public virtual SsoDetails? SsoDetails { get; set; }

    public virtual WorkingHours? WorkingHours { get; set; }

    public virtual BookingSettings? BookingSettings { get; set; }

    public bool SelfServiceCanSeeCustody { get; set; } = false;

    public bool SelfServiceCanSeeBookings { get; set; } = false;

    public bool BaseUserCanSeeCustody { get; set; } = false;

    public bool BaseUserCanSeeBookings { get; set; } = false;

    public bool ManagersCanEditAssets { get; set; } = false;

    public bool BarcodesEnabled { get; set; } = false;

    public DateTime? BarcodesEnabledAt { get; set; }

    public bool UsedBarcodeTrial { get; set; } = false;

    public bool AuditsEnabled { get; set; } = false;

    public DateTime? AuditsEnabledAt { get; set; }

    public bool UsedAuditTrial { get; set; } = false;

    public bool WorkspaceDisabled { get; set; } = false;

    public virtual ICollection<Location> Locations { get; set; } = new List<Location>();

    public virtual ICollection<Category> Categories { get; set; } = new List<Category>();

    public virtual ICollection<Tag> Tags { get; set; } = new List<Tag>();

    public virtual ICollection<Qr> QrCodes { get; set; } = new List<Qr>();

    public virtual ICollection<Barcode> Barcodes { get; set; } = new List<Barcode>();

    public virtual ICollection<Invite> Invites { get; set; } = new List<Invite>();

    public virtual ICollection<UserOrganization> UserOrganizations { get; set; } = new List<UserOrganization>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<CustomField> CustomFields { get; set; } = new List<CustomField>();

    public virtual ICollection<Image> Images { get; set; } = new List<Image>();

    public virtual ICollection<Booking> Bookings { get; set; } = new List<Booking>();

    public virtual ICollection<Kit> Kits { get; set; } = new List<Kit>();

    public virtual ICollection<AssetModel> AssetModels { get; set; } = new List<AssetModel>();

    public virtual ICollection<AssetIndexSettings> AssetIndexSettings { get; set; } = new List<AssetIndexSettings>();

    public virtual ICollection<AssetReminder> AssetReminders { get; set; } = new List<AssetReminder>();

    public virtual ICollection<AssetFilterPreset> AssetFilterPresets { get; set; } = new List<AssetFilterPreset>();

    public virtual ICollection<AuditSession> AuditSessions { get; set; } = new List<AuditSession>();

    public virtual ICollection<AuditImage> AuditImages { get; set; } = new List<AuditImage>();

    public virtual ICollection<AssetImport> AssetImports { get; set; } = new List<AssetImport>();

    public virtual ICollection<LinkPreview> LinkPreviews { get; set; } = new List<LinkPreview>();

    public virtual ICollection<PdfTemplate> PdfTemplates { get; set; } = new List<PdfTemplate>();

    public virtual ICollection<User> UsersWithLastSelected { get; set; } = new List<User>();

    public virtual ICollection<ScimToken> ScimTokens { get; set; } = new List<ScimToken>();

    public virtual ICollection<UserScimExternalId> ScimExternalIds { get; set; } = new List<UserScimExternalId>();

    public virtual ICollection<RoleChangeLog> RoleChangeLogs { get; set; } = new List<RoleChangeLog>();

    public virtual ICollection<TeamMemberNote> TeamMemberNotes { get; set; } = new List<TeamMemberNote>();

    public virtual ICollection<ActivityEvent> ActivityEvents { get; set; } = new List<ActivityEvent>();

    public bool HasSequentialIdsMigrated { get; set; } = false;

    public QrIdDisplayPreference QrIdDisplayPreference { get; set; } = QrIdDisplayPreference.QrId;

    public bool ShowShelfBranding { get; set; } = true;

    public bool ShowQrCodesOnPdfs { get; set; } = true;

    public string? CustomEmailFooter { get; set; }
}
